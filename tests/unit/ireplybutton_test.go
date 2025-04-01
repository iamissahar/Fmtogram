package unit

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type replyButtonT struct {
	name          string
	str           string
	rchat         *fmtogram.KeyboardButtonRequestChat
	rpoll         *fmtogram.KeyboardButtonPollType
	rusers        *fmtogram.KeyboardButtonRequestUsers
	webapp        *fmtogram.WebAppInfo
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type rpbTestContainer struct {
	name          string
	inputStr      []string
	inputRChat    []*fmtogram.KeyboardButtonRequestChat
	inputRPoll    []*fmtogram.KeyboardButtonPollType
	inputRUsers   []*fmtogram.KeyboardButtonRequestUsers
	inputWebApp   []*fmtogram.WebAppInfo
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT
}

func putRButtonWriteRequestChat(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, "", rpbtc.inputRChat[i], nil, nil, nil, rpb.WriteRequestChat, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func putRButtonWriteRequestContact(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, "", nil, nil, nil, nil, rpb.WriteRequestContact, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func putRButtonWriteRequestLocation(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, "", nil, nil, nil, nil, rpb.WriteRequestLocation, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func putRButtonWriteRequestPoll(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, "", nil, rpbtc.inputRPoll[i], nil, nil, rpb.WriteRequestPoll, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func putRButtonWriteRequestUsers(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, "", nil, nil, rpbtc.inputRUsers[i], nil, rpb.WriteRequestUsers, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func putRButtonWriteString(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, rpbtc.inputStr[i], nil, nil, nil, nil, rpb.WriteString, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func putRButtonWriteWebApp(rpbtc rpbTestContainer, rpb fmtogram.IReplyButton, i int) *replyButtonT {
	return &replyButtonT{rpbtc.name, "", nil, nil, nil, rpbtc.inputWebApp[i], rpb.WriteWebApp, rpbtc.isExpectedErr[i], rpbtc.codeErr[i]}
}

func (rpbtc *rpbTestContainer) writeRequestChat() {
	rpbtc.name = "(IReplyButton).WriteRequestChat()"
	rpbtc.inputRChat = []*fmtogram.KeyboardButtonRequestChat{{RequestID: 777, RequestTitle: true},
		nil,
		{RequestID: 777, RequestTitle: true}, {RequestID: 777, RequestTitle: true}}
	rpbtc.isExpectedErr = []bool{false, true, false, true}
	rpbtc.codeErr = []string{"", "20", "", "10"}
	rpbtc.amount, rpbtc.until = 4, 2
	rpbtc.buildF = putRButtonWriteRequestChat
}

func (rpbtc *rpbTestContainer) writeRequestContact() {
	rpbtc.name = "(IReplyButton).WriteRequestContact()"
	rpbtc.isExpectedErr = []bool{false, false, true}
	rpbtc.codeErr = []string{"", "", "10"}
	rpbtc.amount, rpbtc.until = 3, 1
	rpbtc.buildF = putRButtonWriteRequestContact
}

func (rpbtc *rpbTestContainer) writeRequestLocation() {
	rpbtc.name = "(IReplyButton).WriteRequestLocation()"
	rpbtc.isExpectedErr = []bool{false, false, true}
	rpbtc.codeErr = []string{"", "", "10"}
	rpbtc.amount, rpbtc.until = 3, 1
	rpbtc.buildF = putRButtonWriteRequestLocation
}

func (rpbtc *rpbTestContainer) writeRequestPoll() {
	rpbtc.name = "(IReplyButton).WriteRequestPoll()"
	rpbtc.inputRPoll = []*fmtogram.KeyboardButtonPollType{{Type: "alo"},
		nil,
		{Type: "Bu ya"}, {Type: "nu ti"}}
	rpbtc.isExpectedErr = []bool{false, true, false, true}
	rpbtc.codeErr = []string{"", "20", "", "10"}
	rpbtc.amount, rpbtc.until = 4, 2
	rpbtc.buildF = putRButtonWriteRequestPoll
}

func (rpbtc *rpbTestContainer) writeRequestUsers() {
	rpbtc.name = "(IReplyButton).WriteRequestUsers()"
	rpbtc.inputRUsers = []*fmtogram.KeyboardButtonRequestUsers{{RequestID: 12315, RequestPhoto: true},
		nil,
		{RequestID: 12315, RequestPhoto: true}, {RequestID: 12315, RequestPhoto: true}}
	rpbtc.isExpectedErr = []bool{false, true, false, true}
	rpbtc.codeErr = []string{"", "20", "", "10"}
	rpbtc.amount, rpbtc.until = 4, 2
	rpbtc.buildF = putRButtonWriteRequestUsers
}

func (rpbtc *rpbTestContainer) writeString() {
	rpbtc.name = "(IReplyButton).WriteString()"
	rpbtc.inputStr = []string{"a;lsdl;kassdkl;", "", "a;lskdk;lasgk;la", "234i9023u8i9o45i2ojk"}
	rpbtc.isExpectedErr = []bool{false, true, false, true}
	rpbtc.codeErr = []string{"", "20", "", "10"}
	rpbtc.amount, rpbtc.until = 4, 2
	rpbtc.buildF = putRButtonWriteString
}

func (rpbtc *rpbTestContainer) writeWebApp() {
	rpbtc.name = "(IReplyButton).WriteWebApp()"
	rpbtc.inputWebApp = []*fmtogram.WebAppInfo{{Url: "https://youtube.com"},
		nil,
		{Url: "https://youtube.com"}, {Url: "https://youtube.com"}}
	rpbtc.isExpectedErr = []bool{false, true, false, true}
	rpbtc.codeErr = []string{"", "20", "", "10"}
	rpbtc.amount, rpbtc.until = 4, 2
	rpbtc.buildF = putRButtonWriteWebApp
}

func (rpb *replyButtonT) callStrF(f func(string) error, t *testing.T) {
	if !rpb.isExpectedErr {
		if err := f(rpb.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rpb.str); err.Error() != rpb.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rpb *replyButtonT) callRChatF(f func(*fmtogram.KeyboardButtonRequestChat) error, t *testing.T) {
	if !rpb.isExpectedErr {
		if err := f(rpb.rchat); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rpb.rchat); err.Error() != rpb.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rpb *replyButtonT) callRPollF(f func(*fmtogram.KeyboardButtonPollType) error, t *testing.T) {
	if !rpb.isExpectedErr {
		if err := f(rpb.rpoll); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rpb.rpoll); err.Error() != rpb.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rpb *replyButtonT) callRUsersF(f func(*fmtogram.KeyboardButtonRequestUsers) error, t *testing.T) {
	if !rpb.isExpectedErr {
		if err := f(rpb.rusers); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rpb.rusers); err.Error() != rpb.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rpb *replyButtonT) callWebAppF(f func(*fmtogram.WebAppInfo) error, t *testing.T) {
	if !rpb.isExpectedErr {
		if err := f(rpb.webapp); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(rpb.webapp); err.Error() != rpb.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rpb *replyButtonT) callBoolF(f func() error, t *testing.T) {
	if !rpb.isExpectedErr {
		if err := f(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(); err.Error() != rpb.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (rpb *replyButtonT) startTest(part string, i int, t *testing.T) {
	switch f := rpb.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, rpb.name, rpb.codeErr, rpb.str, rpb.isExpectedErr, i, t)
		rpb.callStrF(f, t)
	case func(*fmtogram.KeyboardButtonRequestChat) error:
		printTestLog(part, rpb.name, rpb.codeErr, rpb.rchat, rpb.isExpectedErr, i, t)
		rpb.callRChatF(f, t)
	case func(*fmtogram.KeyboardButtonPollType) error:
		printTestLog(part, rpb.name, rpb.codeErr, rpb.rpoll, rpb.isExpectedErr, i, t)
		rpb.callRPollF(f, t)
	case func(*fmtogram.KeyboardButtonRequestUsers) error:
		printTestLog(part, rpb.name, rpb.codeErr, rpb.rusers, rpb.isExpectedErr, i, t)
		rpb.callRUsersF(f, t)
	case func(*fmtogram.WebAppInfo) error:
		printTestLog(part, rpb.name, rpb.codeErr, rpb.webapp, rpb.isExpectedErr, i, t)
		rpb.callWebAppF(f, t)
	case func() error:
		printTestLog(part, rpb.name, rpb.codeErr, true, rpb.isExpectedErr, i, t)
		rpb.callBoolF(f, t)
	default:
		t.Logf("%T", f)
		t.Fatal("unexpected type of tested function")
	}
}

func (rpbtc *rpbTestContainer) createTestArrays(msg *fmtogram.Message, t *testing.T) ([]UnitTest, []UnitTest) {
	var btn fmtogram.IReplyButton
	a, b := make([]UnitTest, rpbtc.until), make([]UnitTest, rpbtc.amount-rpbtc.until)
	for i, j := 0, 0; i < rpbtc.amount; i++ {
		if i < rpbtc.until {
			if kb, err := fmtogram.NewKeyboard().WriteReply(); err != nil {
				t.Fatal(err)
			} else {
				if err = kb.Set([]int{1, 2, 3}); err != nil {
					t.Fatal(err)
				}
				if btn, err = kb.NewButton(1, 1); err != nil {
					t.Fatal(err)
				}
			}
			a[i] = rpbtc.buildF(*rpbtc, btn, i)
		} else {
			if j%2 == 0 {
				if kb, err := fmtogram.NewKeyboard().WriteReply(); err != nil {
					t.Fatal(err)
				} else {
					if err = kb.Set([]int{1, 2, 3}); err != nil {
						t.Fatal(err)
					}
					if btn, err = kb.NewButton(1, 1); err != nil {
						t.Fatal(err)
					}
				}
			}
			b[j] = rpbtc.buildF(*rpbtc, btn, i)
			j++
		}
	}
	return a, b
}

func (rpbtc *rpbTestContainer) mainLogic(msg *fmtogram.Message, t *testing.T) {
	rpumentcontainer, doublecontainer := rpbtc.createTestArrays(msg, t)
	check(rpumentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestRButtonWriteRequestChat(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeRequestChat()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}

func TestRButtonWriteRequestContact(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeRequestContact()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}

func TestRButtonWriteRequestLocation(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeRequestLocation()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}

func TestRButtonWriteRequestPoll(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeRequestPoll()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}

func TestRButtonWriteRequestUsers(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeRequestUsers()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}

func TestRButtonWriteString(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeString()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}

func TestRButtonWriteWebApp(t *testing.T) {
	t.Parallel()
	rpbtc := new(rpbTestContainer)
	rpbtc.writeWebApp()
	msg := fmtogram.CreateEmpltyMessage()
	rpbtc.mainLogic(msg, t)
}
