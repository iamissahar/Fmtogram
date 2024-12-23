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
	buildF        func(chtc chTestContainer, ch formatter.IChat, c *chatT, i int)
}

func (chtc chTestContainer) defaultChatT(i int) *chatT {
	return &chatT{name: chtc.name, isExpectedErr: chtc.isExpectedErr[i], codeErr: chtc.codeErr[i]}
}

func putChatWriteBusinessConnectionID(chtc chTestContainer, ch formatter.IChat, c *chatT, i int) {
	c.testedFunc = ch.WriteBusinessConnectionID
	c.str = chtc.inputStr[i]
}

func putChatWriteChatID(chtc chTestContainer, ch formatter.IChat, c *chatT, i int) {
	c.testedFunc = ch.WriteChatID
	c.integer = chtc.inputInt[i]
}

func putChatWriteChatName(chtc chTestContainer, ch formatter.IChat, c *chatT, i int) {
	c.testedFunc = ch.WriteChatName
	c.str = chtc.inputStr[i]
}

func putChatWriteFromChatID(chtc chTestContainer, ch formatter.IChat, c *chatT, i int) {
	c.testedFunc = ch.WriteFromChatID
	c.integer = chtc.inputInt[i]
}

func putChatWriteFromChatName(chtc chTestContainer, ch formatter.IChat, c *chatT, i int) {
	c.testedFunc = ch.WriteFromChatName
	c.str = chtc.inputStr[i]
}

func putChatWriteSenderChatID(chtc chTestContainer, ch formatter.IChat, c *chatT, i int) {
	c.testedFunc = ch.WriteSenderChatID
	c.integer = chtc.inputInt[i]
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

func (chtc *chTestContainer) writeSenderChatID() {
	chtc.name = "(IChat).WriteSenderChatID()"
	chtc.inputInt = []int{10230, 0, 2214, 666623}
	chtc.isExpectedErr = []bool{false, true, false, true}
	chtc.codeErr = []string{"", "20", "", "10"}
	chtc.amount, chtc.until = 4, 2
	chtc.buildF = putChatWriteSenderChatID
}

func (ch *chatT) startTest(part string, i int, t *testing.T) {
	switch f := ch.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, ch.name, ch.codeErr, ch.str, ch.isExpectedErr, i)
		checkError(f(ch.str), ch.isExpectedErr, ch.codeErr, t)
	case func(int) error:
		printTestLog(part, ch.name, ch.codeErr, ch.integer, ch.isExpectedErr, i)
		checkError(f(ch.integer), ch.isExpectedErr, ch.codeErr, t)
	case func() error:
		printTestLog(part, ch.name, ch.codeErr, true, ch.isExpectedErr, i)
		checkError(f(), ch.isExpectedErr, ch.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (chtc *chTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var ch formatter.IChat
	a, b := make([]UnitTest, chtc.until), make([]UnitTest, chtc.amount-chtc.until)
	for i, j := 0, 0; i < chtc.amount; i++ {
		c := chtc.defaultChatT(i)
		if i < chtc.until {
			ch = msg.NewChat()
			chtc.buildF(*chtc, ch, c, i)
			a[i] = c
		} else {
			if j%2 == 0 {
				ch = msg.NewChat()
			}
			chtc.buildF(*chtc, ch, c, i)
			b[j] = c
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

func TestChatWriteSenderChatID(t *testing.T) {
	chtc := new(chTestContainer)
	chtc.writeSenderChatID()
	msg := formatter.CreateEmpltyMessage()
	chtc.mainLogic(msg, t)
}
