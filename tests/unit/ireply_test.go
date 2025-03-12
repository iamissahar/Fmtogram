package unit

import (
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
)

type replyT struct {
	name          string
	str           string
	integer       [2]int
	array         []int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type rpTestContainer struct {
	name          string
	inputStr      []string
	inputInt      [][2]int
	inputArr      [][]int
	isExpectedErr []bool
	firststep     func(_ formatter.IReply, plan []int) error
	setplan       [][]int
	codeErr       []string
	amount, until int
	buildF        func(rptc rpTestContainer, rp formatter.IReply, i int) *replyT
}

func putReplyNewButton(rptc rpTestContainer, rp formatter.IReply, i int) *replyT {
	return &replyT{rptc.name, "", rptc.inputInt[i], nil, rp.NewButton, rptc.isExpectedErr[i], rptc.codeErr[i]}
}

func putReplySet(rptc rpTestContainer, rp formatter.IReply, i int) *replyT {
	return &replyT{rptc.name, "", [2]int{}, rptc.inputArr[i], rp.Set, rptc.isExpectedErr[i], rptc.codeErr[i]}
}

func putReplyWriteIsPersistent(rptc rpTestContainer, rp formatter.IReply, i int) *replyT {
	return &replyT{rptc.name, "", [2]int{}, nil, rp.WriteIsPersistent, rptc.isExpectedErr[i], rptc.codeErr[i]}
}

func putReplyWriteResizeKeyboard(rptc rpTestContainer, rp formatter.IReply, i int) *replyT {
	return &replyT{rptc.name, "", [2]int{}, nil, rp.WriteResizeKeyboard, rptc.isExpectedErr[i], rptc.codeErr[i]}
}

func putReplyWriteInputFieldPlaceholder(rptc rpTestContainer, rp formatter.IReply, i int) *replyT {
	return &replyT{rptc.name, rptc.inputStr[i], [2]int{}, nil, rp.WriteInputFieldPlaceholder, rptc.isExpectedErr[i], rptc.codeErr[i]}
}

func putReplyWriteRemove(rptc rpTestContainer, rp formatter.IReply, i int) *replyT {
	return &replyT{rptc.name, "", [2]int{}, nil, rp.WriteRemove, rptc.isExpectedErr[i], rptc.codeErr[i]}
}

func (rptc *rpTestContainer) newButton() {
	rptc.name = "(IReply).NewButton()"
	rptc.inputInt = [][2]int{{2, 4}, {-1, -6}, {11, 55}, {2, 0}, {1, 1}, {1, 1}}
	rptc.firststep = formatter.IReply.Set
	rptc.setplan = [][]int{nil, nil, {4, 3}, {4, 6, 1}, {3, 3}, nil}
	rptc.isExpectedErr = []bool{true, true, true, false, false, true}
	rptc.codeErr = []string{"1", "20", "1", "", "", "10"}
	rptc.amount, rptc.until = 6, 4
	rptc.buildF = putReplyNewButton
}

func (rptc *rpTestContainer) set() {
	rptc.name = "(IReply).Set()"
	rptc.inputArr = [][]int{{4, 3}, nil, {0, 0}, {0}, {4, 6, 1}, {3, 3}}
	rptc.isExpectedErr = []bool{false, true, true, true, false, true}
	rptc.codeErr = []string{"", "20", "20", "20", "", "10"}
	rptc.amount, rptc.until = 6, 4
	rptc.buildF = putReplySet
}

func (rptc *rpTestContainer) writeIsPersistent() {
	rptc.name = "(IReply).WriteIsPersistent"
	rptc.isExpectedErr = []bool{false, false, true}
	rptc.codeErr = []string{"", "", "10"}
	rptc.amount, rptc.until = 3, 1
	rptc.buildF = putReplyWriteIsPersistent
}

func (rptc *rpTestContainer) writeResizeKeyboard() {
	rptc.name = "(IReply).WriteResizeKeyboard"
	rptc.isExpectedErr = []bool{false, false, true}
	rptc.codeErr = []string{"", "", "10"}
	rptc.amount, rptc.until = 3, 1
	rptc.buildF = putReplyWriteResizeKeyboard
}

func (rptc *rpTestContainer) writeOneTimeKeyboard() {
	rptc.name = "(IReply).WriteOneTimeKeyboard"
	rptc.isExpectedErr = []bool{false, false, true}
	rptc.codeErr = []string{"", "", "10"}
	rptc.amount, rptc.until = 3, 1
	rptc.buildF = putReplyWriteResizeKeyboard
}

func (rptc *rpTestContainer) writeInputFieldPlaceholder() {
	rptc.name = "(IReply).WriteInputFieldPlaceholder"
	rptc.inputStr = []string{"asdlals;d", "", "14rsdffsd", "something"}
	rptc.isExpectedErr = []bool{false, true, false, true}
	rptc.codeErr = []string{"", "20", "", "10"}
	rptc.amount, rptc.until = 4, 2
	rptc.buildF = putReplyWriteInputFieldPlaceholder
}

func (rptc *rpTestContainer) writeRemove() {
	rptc.name = "(IReply).WriteRemove"
	rptc.isExpectedErr = []bool{false, false, true}
	rptc.codeErr = []string{"", "", "10"}
	rptc.amount, rptc.until = 3, 1
	rptc.buildF = putReplyWriteRemove
}

func (rp *replyT) callStrF(f func(string) error, t *testing.T) {
	if !rp.isExpectedErr {
		if err := f(rp.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rp.str); err.Error() != rp.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rp *replyT) callDIntF(f func(int, int) (formatter.IReplyButton, error), t *testing.T) {
	if !rp.isExpectedErr {
		if _, err := f(rp.integer[0], rp.integer[1]); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if _, err := f(rp.integer[0], rp.integer[1]); err.Error() != rp.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rp *replyT) callSliceF(f func([]int) error, t *testing.T) {
	if !rp.isExpectedErr {
		if err := f(rp.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rp.array); err.Error() != rp.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rp *replyT) callBoolF(f func() error, t *testing.T) {
	if !rp.isExpectedErr {
		if err := f(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(); err.Error() != rp.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rp *replyT) startTest(part string, i int, t *testing.T) {
	switch f := rp.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, rp.name, rp.codeErr, rp.str, rp.isExpectedErr, i, t)
		rp.callStrF(f, t)
	case func(int, int) (formatter.IReplyButton, error):
		printTestLog(part, rp.name, rp.codeErr, rp.integer, rp.isExpectedErr, i, t)
		rp.callDIntF(f, t)
	case func([]int) error:
		printTestLog(part, rp.name, rp.codeErr, rp.array, rp.isExpectedErr, i, t)
		rp.callSliceF(f, t)
	case func() error:
		printTestLog(part, rp.name, rp.codeErr, true, rp.isExpectedErr, i, t)
		rp.callBoolF(f, t)
	default:
		t.Logf("%T", f)
		t.Fatal("unexpected type of tested function")
	}
}

func (rptc *rpTestContainer) createTestArrays(msg *formatter.Message, t *testing.T) ([]UnitTest, []UnitTest) {
	var rp formatter.IReply
	var err error
	a, b := make([]UnitTest, rptc.until), make([]UnitTest, rptc.amount-rptc.until)
	for i, j := 0, 0; i < rptc.amount; i++ {
		if i < rptc.until {
			if rp, err = msg.NewKeyboard().WriteReply(); err != nil {
				t.Fatal(err)
			}
			a[i] = rptc.buildF(*rptc, rp, i)
		} else {
			if j%2 == 0 {
				if rp, err = msg.NewKeyboard().WriteReply(); err != nil {
					t.Fatal(err)
				}
			}
			b[j] = rptc.buildF(*rptc, rp, i)
			j++
		}
		if (len(rptc.setplan) > i) && (rptc.setplan[i] != nil) {
			rptc.firststep(rp, rptc.setplan[i])
		}
	}
	return a, b
}

func (rptc *rpTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	rpumentcontainer, doublecontainer := rptc.createTestArrays(msg, t)
	check(rpumentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestReplyNewButton(t *testing.T) {
	rptc := new(rpTestContainer)
	rptc.newButton()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}

func TestReplySet(t *testing.T) {
	t.Parallel()
	rptc := new(rpTestContainer)
	rptc.set()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}

func TestReplyWriteIsPersistent(t *testing.T) {
	t.Parallel()
	rptc := new(rpTestContainer)
	rptc.writeIsPersistent()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}

func TestReplyWriteResizeKeyboard(t *testing.T) {
	t.Parallel()
	rptc := new(rpTestContainer)
	rptc.writeResizeKeyboard()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}

func TestReplyWriteOneTimeKeyboard(t *testing.T) {
	t.Parallel()
	rptc := new(rpTestContainer)
	rptc.writeOneTimeKeyboard()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}

func TestReplyWriteInputFieldPlaceholder(t *testing.T) {
	t.Parallel()
	rptc := new(rpTestContainer)
	rptc.writeInputFieldPlaceholder()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}

func TestReplyWriteRemove(t *testing.T) {
	t.Parallel()
	rptc := new(rpTestContainer)
	rptc.writeRemove()
	msg := formatter.CreateEmpltyMessage()
	rptc.mainLogic(msg, t)
}
