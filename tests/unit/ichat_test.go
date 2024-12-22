package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
)

type chatT struct {
	name          string
	str           string
	integer       int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type chTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(chtc chTestContainer, ch formatter.IChat, i int) *chatT
}

func putChatWriteBusinessConnectionID(chtc chTestContainer, ch formatter.IChat, i int) *chatT {
	return &chatT{chtc.name, chtc.inputStr[i], 0, ch.WriteBusinessConnectionID, chtc.isExpectedErr[i], chtc.codeErr[i]}
}

func putChatWriteChatID(chtc chTestContainer, ch formatter.IChat, i int) *chatT {
	return &chatT{chtc.name, "", chtc.inputInt[i], ch.WriteChatID, chtc.isExpectedErr[i], chtc.codeErr[i]}
}

func putChatWriteChatName(chtc chTestContainer, ch formatter.IChat, i int) *chatT {
	return &chatT{chtc.name, chtc.inputStr[i], 0, ch.WriteChatName, chtc.isExpectedErr[i], chtc.codeErr[i]}
}

func putChatWriteFromChatID(chtc chTestContainer, ch formatter.IChat, i int) *chatT {
	return &chatT{chtc.name, "", chtc.inputInt[i], ch.WriteFromChatID, chtc.isExpectedErr[i], chtc.codeErr[i]}
}

func putChatWriteFromChatName(chtc chTestContainer, ch formatter.IChat, i int) *chatT {
	return &chatT{chtc.name, chtc.inputStr[i], 0, ch.WriteFromChatName, chtc.isExpectedErr[i], chtc.codeErr[i]}
}

func (chtc *chTestContainer) writeBusinessConnectionID() {
	chtc.name = "(IChat).WriteBusinessConnectionID()"
	chtc.inputStr = []string{"iljksasdkljaskljd", "", "1203891o3l1k", "as;'das;'d;'asd;'"}
	chtc.isExpectedErr = []bool{false, true, false, true}
	chtc.codeErr = []string{"", "20", "", "10"}
	chtc.amount, chtc.until = 4, 2
	chtc.buildF = putChatWriteBusinessConnectionID
}

func (chtc *chTestContainer) writeChatID() {
	chtc.name = "(IChat).WriteChatID()"
	chtc.inputInt = []int{10230, 0, 2214, 666623}
	chtc.isExpectedErr = []bool{false, true, false, true}
	chtc.codeErr = []string{"", "20", "", "10"}
	chtc.amount, chtc.until = 4, 2
	chtc.buildF = putChatWriteChatID
}

func (chtc *chTestContainer) writeChatName() {
	chtc.name = "(IChat).WriteChatName()"
	chtc.inputStr = []string{"Name", "", ":noname", "something like a name"}
	chtc.isExpectedErr = []bool{false, true, false, true}
	chtc.codeErr = []string{"", "20", "", "10"}
	chtc.amount, chtc.until = 4, 2
	chtc.buildF = putChatWriteChatName
}

func (chtc *chTestContainer) writeFromChatID() {
	chtc.name = "(IChat).WriteFromChatID()"
	chtc.inputInt = []int{10230, 0, 2214, 666623}
	chtc.isExpectedErr = []bool{false, true, false, true}
	chtc.codeErr = []string{"", "20", "", "10"}
	chtc.amount, chtc.until = 4, 2
	chtc.buildF = putChatWriteFromChatID
}

func (chtc *chTestContainer) writeFromChatName() {
	chtc.name = "(IChat).WriteFromChatName()"
	chtc.inputStr = []string{"Name", "", ":noname", "something like a name"}
	chtc.isExpectedErr = []bool{false, true, false, true}
	chtc.codeErr = []string{"", "20", "", "10"}
	chtc.amount, chtc.until = 4, 2
	chtc.buildF = putChatWriteFromChatName
}

func (ch *chatT) callStrF(f func(string) error, t *testing.T) {
	if !ch.isExpectedErr {
		if err := f(ch.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(ch.str); err.Error() != ch.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ch *chatT) callIntF(f func(int) error, t *testing.T) {
	if !ch.isExpectedErr {
		if err := f(ch.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(ch.integer); err.Error() != ch.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ch *chatT) callBoolF(f func() error, t *testing.T) {
	if !ch.isExpectedErr {
		if err := f(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(); err.Error() != ch.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ch *chatT) startTest(part string, i int, t *testing.T) {
	switch f := ch.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, ch.name, ch.codeErr, ch.str, ch.isExpectedErr, i)
		ch.callStrF(f, t)
	case func(int) error:
		printTestLog(part, ch.name, ch.codeErr, ch.integer, ch.isExpectedErr, i)
		ch.callIntF(f, t)
	case func() error:
		printTestLog(part, ch.name, ch.codeErr, true, ch.isExpectedErr, i)
		ch.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (chtc *chTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var ch formatter.IChat
	a, b := make([]UnitTest, chtc.until), make([]UnitTest, chtc.amount-chtc.until)
	for i, j := 0, 0; i < chtc.amount; i++ {
		if i < chtc.until {
			ch = msg.NewChat()
			a[i] = chtc.buildF(*chtc, ch, i)
		} else {
			if j%2 == 0 {
				ch = msg.NewChat()
			}
			b[j] = chtc.buildF(*chtc, ch, i)
			j++
		}
	}
	return a, b
}

func (chtc *chTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	chacontainer, doublecontainer := chtc.createTestArrays(msg)
	check(chacontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestChatWriteBusinessConnectionID(t *testing.T) {
	chtc := new(chTestContainer)
	chtc.writeBusinessConnectionID()
	msg := formatter.CreateEmpltyMessage()
	chtc.mainLogic(msg, t)
}

func TestChatWriteChatID(t *testing.T) {
	chtc := new(chTestContainer)
	chtc.writeChatID()
	msg := formatter.CreateEmpltyMessage()
	chtc.mainLogic(msg, t)
}

func TestChatWriteChatName(t *testing.T) {
	chtc := new(chTestContainer)
	chtc.writeChatName()
	msg := formatter.CreateEmpltyMessage()
	chtc.mainLogic(msg, t)
}

func TestChatWriteFromChatID(t *testing.T) {
	chtc := new(chTestContainer)
	chtc.writeFromChatID()
	msg := formatter.CreateEmpltyMessage()
	chtc.mainLogic(msg, t)
}

func TestChatWriteFromChatName(t *testing.T) {
	chtc := new(chTestContainer)
	chtc.writeFromChatName()
	msg := formatter.CreateEmpltyMessage()
	chtc.mainLogic(msg, t)
}
