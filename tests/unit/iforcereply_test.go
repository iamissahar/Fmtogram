package unit

import (
	"strings"
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
)

type forceReplyT struct {
	name          string
	str           string
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type frpTestContainer struct {
	name          string
	inputStr      []string
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(frptc frpTestContainer, frp formatter.IForceReply, i int) *forceReplyT
}

func putFReplyWriteForceReply(frptc frpTestContainer, frp formatter.IForceReply, i int) *forceReplyT {
	return &forceReplyT{frptc.name, "", frp.WriteForceReply, frptc.isExpectedErr[i], frptc.codeErr[i]}
}

func putFReplyWriteInputFieldPlaceholder(frptc frpTestContainer, frp formatter.IForceReply, i int) *forceReplyT {
	return &forceReplyT{frptc.name, frptc.inputStr[i], frp.WriteInputFieldPlaceholder, frptc.isExpectedErr[i], frptc.codeErr[i]}
}

func putFReplyWriteSelective(frptc frpTestContainer, frp formatter.IForceReply, i int) *forceReplyT {
	return &forceReplyT{frptc.name, "", frp.WriteSelective, frptc.isExpectedErr[i], frptc.codeErr[i]}
}

func (frptc *frpTestContainer) writeForceReply() {
	frptc.name = "(IForceReply).WriteForceReply()"
	frptc.isExpectedErr = []bool{false, false, true}
	frptc.codeErr = []string{"", "", "10"}
	frptc.amount, frptc.until = 3, 1
	frptc.buildF = putFReplyWriteForceReply
}

func (frptc *frpTestContainer) writeInputFieldPlaceholder() {
	frptc.name = "(IForceReply).WriteInputFieldPlaceholder()"
	frptc.inputStr = []string{"askld", "", strings.Repeat("2", 64), strings.Repeat("2", 65), "asdad", "something"}
	frptc.isExpectedErr = []bool{false, true, false, true, false, true}
	frptc.codeErr = []string{"", "20", "", "20", "", "10"}
	frptc.amount, frptc.until = 6, 4
	frptc.buildF = putFReplyWriteInputFieldPlaceholder
}

func (frptc *frpTestContainer) writeSelective() {
	frptc.name = "(IForceReply).WriteSelective()"
	frptc.isExpectedErr = []bool{false, false, true}
	frptc.codeErr = []string{"", "", "10"}
	frptc.amount, frptc.until = 3, 1
	frptc.buildF = putFReplyWriteSelective
}

func (frp *forceReplyT) startTest(part string, i int, t *testing.T) {
	switch f := frp.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, frp.name, frp.codeErr, frp.str, frp.isExpectedErr, i, t)
		checkError(f(frp.str), frp.isExpectedErr, frp.codeErr, t)
	case func() error:
		printTestLog(part, frp.name, frp.codeErr, true, frp.isExpectedErr, i, t)
		checkError(f(), frp.isExpectedErr, frp.codeErr, t)
	default:
		t.Fatalf("unexpected type of tested function: %T", f)
	}
}

func (frptc *frpTestContainer) createTestArrays(msg *formatter.Message, t *testing.T) ([]UnitTest, []UnitTest) {
	var kb formatter.IForceReply
	var err error
	a, b := make([]UnitTest, frptc.until), make([]UnitTest, frptc.amount-frptc.until)
	for i, j := 0, 0; i < frptc.amount; i++ {
		if i < frptc.until {
			if kb, err = msg.NewKeyboard().WriteForceReply(); err != nil {
				t.Fatal(err)
			}
			a[i] = frptc.buildF(*frptc, kb, i)
		} else {
			if j%2 == 0 {
				if kb, err = msg.NewKeyboard().WriteForceReply(); err != nil {
					t.Fatal(err)
				}
			}
			b[j] = frptc.buildF(*frptc, kb, i)
			j++
		}
	}
	return a, b
}

func (frptc *frpTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	rpumentcontainer, doublecontainer := frptc.createTestArrays(msg, t)
	check(rpumentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestForceReplyWriteForceReply(t *testing.T) {
	t.Parallel()
	frptc := new(frpTestContainer)
	frptc.writeForceReply()
	msg := formatter.CreateEmpltyMessage()
	frptc.mainLogic(msg, t)
}

func TestForceReplyWriteInputFieldPlaceholder(t *testing.T) {
	t.Parallel()
	frptc := new(frpTestContainer)
	frptc.writeInputFieldPlaceholder()
	msg := formatter.CreateEmpltyMessage()
	frptc.mainLogic(msg, t)
}

func TestForceReplyWriteSelective(t *testing.T) {
	t.Parallel()
	frptc := new(frpTestContainer)
	frptc.writeSelective()
	msg := formatter.CreateEmpltyMessage()
	frptc.mainLogic(msg, t)
}
