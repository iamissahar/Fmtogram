package unit

import (
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
)

type inlineT struct {
	name          string
	integer       [2]int
	array         []int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type inTestContainer struct {
	name          string
	inputInt      [][2]int
	inputArr      [][]int
	isExpectedErr []bool
	firststep     func(_ formatter.IInline, plan []int) error
	setplan       [][]int
	codeErr       []string
	amount, until int
	buildF        func(intc inTestContainer, in formatter.IInline, i int) *inlineT
}

func putInlineNewButton(intc inTestContainer, in formatter.IInline, i int) *inlineT {
	return &inlineT{intc.name, intc.inputInt[i], nil, in.NewButton, intc.isExpectedErr[i], intc.codeErr[i]}
}

func putInlineSet(intc inTestContainer, in formatter.IInline, i int) *inlineT {
	return &inlineT{intc.name, [2]int{}, intc.inputArr[i], in.Set, intc.isExpectedErr[i], intc.codeErr[i]}
}

func (intc *inTestContainer) newButton() {
	intc.name = "(IIline).NewButton()"
	intc.inputInt = [][2]int{{2, 4}, {-1, -6}, {11, 55}, {2, 0}, {1, 1}, {1, 1}}
	intc.firststep = formatter.IInline.Set
	intc.setplan = [][]int{nil, nil, {4, 3}, {4, 6, 1}, {3, 3}, nil}
	intc.isExpectedErr = []bool{true, true, true, false, false, true}
	intc.codeErr = []string{"1", "20", "1", "", "", "10"}
	intc.amount, intc.until = 6, 4
	intc.buildF = putInlineNewButton
}

func (intc *inTestContainer) set() {
	intc.name = "(IIline).Set()"
	intc.inputArr = [][]int{{4, 3}, nil, {0, 0}, {0}, {4, 6, 1}, {3, 3}}
	intc.isExpectedErr = []bool{false, true, true, true, false, true}
	intc.codeErr = []string{"", "20", "20", "20", "", "10"}
	intc.amount, intc.until = 6, 4
	intc.buildF = putInlineSet
}

func (in *inlineT) callDIntF(f func(int, int) (formatter.IInlineButton, error), t *testing.T) {
	if !in.isExpectedErr {
		if _, err := f(in.integer[0], in.integer[1]); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if _, err := f(in.integer[0], in.integer[1]); err.Error() != in.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (in *inlineT) callSliceF(f func([]int) error, t *testing.T) {
	if !in.isExpectedErr {
		if err := f(in.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(in.array); err.Error() != in.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (in *inlineT) startTest(part string, i int, t *testing.T) {
	switch f := in.testedFunc.(type) {
	case func(int, int) (formatter.IInlineButton, error):
		printTestLog(part, in.name, in.codeErr, in.integer, in.isExpectedErr, i, t)
		in.callDIntF(f, t)
	case func([]int) error:
		printTestLog(part, in.name, in.codeErr, in.array, in.isExpectedErr, i, t)
		in.callSliceF(f, t)
	default:
		t.Logf("%T", f)
		t.Fatal("unexpected type of tested function")
	}
}

func (intc *inTestContainer) createTestArrays(msg *formatter.Message, t *testing.T) ([]UnitTest, []UnitTest) {
	var in formatter.IInline
	var err error
	a, b := make([]UnitTest, intc.until), make([]UnitTest, intc.amount-intc.until)
	for i, j := 0, 0; i < intc.amount; i++ {
		if i < intc.until {
			if in, err = msg.NewKeyboard().WriteInline(); err != nil {
				t.Fatal(err)
			}
			a[i] = intc.buildF(*intc, in, i)
		} else {
			if j%2 == 0 {
				if in, err = msg.NewKeyboard().WriteInline(); err != nil {
					t.Fatal(err)
				}
			}
			b[j] = intc.buildF(*intc, in, i)
			j++
		}
		if (len(intc.setplan) > i) && (intc.setplan[i] != nil) {
			intc.firststep(in, intc.setplan[i])
		}
	}
	return a, b
}

func (intc *inTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	inumentcontainer, doublecontainer := intc.createTestArrays(msg, t)
	check(inumentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestInlineNewButtton(t *testing.T) {
	t.Parallel()
	intc := new(inTestContainer)
	intc.newButton()
	msg := formatter.CreateEmpltyMessage()
	intc.mainLogic(msg, t)
}

func TestInlineSet(t *testing.T) {
	t.Parallel()
	intc := new(inTestContainer)
	intc.set()
	msg := formatter.CreateEmpltyMessage()
	intc.mainLogic(msg, t)
}
