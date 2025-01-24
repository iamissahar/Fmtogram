package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type inlineButtonT struct {
	name          string
	str           string
	game          *types.CallbackGame
	login         *types.LoginUrl
	sw            *types.SwitchInlineQueryChosenChat
	webapp        *types.WebAppInfo
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type inbTestContainer struct {
	name          string
	inputStr      []string
	inputGames    []*types.CallbackGame
	inputLogins   []*types.LoginUrl
	inputSWs      []*types.SwitchInlineQueryChosenChat
	inputWebApp   []*types.WebAppInfo
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT
}

func putIButtonWriteCallbackData(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, inbtc.inputStr[i], nil, nil, nil, nil, inb.WriteCallbackData, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteCallbackGame(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, "", inbtc.inputGames[i], nil, nil, nil, inb.WriteCallbackGame, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteLoginUrl(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, "", nil, inbtc.inputLogins[i], nil, nil, inb.WriteLoginUrl, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWritePay(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, "", nil, nil, nil, nil, inb.WritePay, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteString(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, inbtc.inputStr[i], nil, nil, nil, nil, inb.WriteString, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteSwitchInlineQuery(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, inbtc.inputStr[i], nil, nil, nil, nil, inb.WriteSwitchInlineQuery, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteSwitchInlineQueryChosenChat(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, "", nil, nil, inbtc.inputSWs[i], nil, inb.WriteSwitchInlineQueryChosenChat, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteSwitchInlineQueryCurrentChat(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, inbtc.inputStr[i], nil, nil, nil, nil, inb.WriteSwitchInlineQueryCurrentChat, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteURL(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, inbtc.inputStr[i], nil, nil, nil, nil, inb.WriteURL, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func putIButtonWriteWebApp(inbtc inbTestContainer, inb formatter.IInlineButton, i int) *inlineButtonT {
	return &inlineButtonT{inbtc.name, "", nil, nil, nil, inbtc.inputWebApp[i], inb.WriteWebApp, inbtc.isExpectedErr[i], inbtc.codeErr[i]}
}

func (inbtc *inbTestContainer) writeCallbackData() {
	inbtc.name = "(IInlineButton).WriteCallbackData()"
	inbtc.inputStr = []string{"Soemthing", "", "yaya", "tututu"}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteCallbackData
}

func (inbtc *inbTestContainer) writeCallbackGame() {
	inbtc.name = "(IInlineButton).WriteCallbackGame()"
	inbtc.inputGames = []*types.CallbackGame{{}, nil, {}, {}}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteCallbackGame
}

func (inbtc *inbTestContainer) writeLoginUrl() {
	inbtc.name = "(IInlineButton).WriteLoginUrl()"
	inbtc.inputLogins = []*types.LoginUrl{{Url: "https://youtube.com"}, nil, {Url: "https://youtube.com"}, {Url: "https://youtube.com"}}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteLoginUrl
}

func (inbtc *inbTestContainer) writePay() {
	inbtc.name = "(IInlineButton).WritePay()"
	inbtc.isExpectedErr = []bool{false, false, true}
	inbtc.codeErr = []string{"", "", "10"}
	inbtc.amount, inbtc.until = 3, 1
	inbtc.buildF = putIButtonWritePay
}

func (inbtc *inbTestContainer) writeString() {
	inbtc.name = "(IInlineButton).WriteString()"
	inbtc.inputStr = []string{"string", "", "nothing", "a"}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteString
}

func (inbtc *inbTestContainer) writeSwitchInlineQuery() {
	inbtc.name = "(IInlineButton).WriteSwitchInlineQuery()"
	inbtc.inputStr = []string{"string", "", "nothing", "a"}
	inbtc.isExpectedErr = []bool{false, false, false, true}
	inbtc.codeErr = []string{"", "", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteSwitchInlineQuery
}

func (inbtc *inbTestContainer) writeSwitchInlineQueryChosenChat() {
	inbtc.name = "(IInlineButton).WriteSwitchInlineQueryChosenChat()"
	inbtc.inputSWs = []*types.SwitchInlineQueryChosenChat{{Query: "ALO!?"}, nil, {Query: "ALO!?"}, {Query: "ALO!?"}}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteSwitchInlineQueryChosenChat
}

func (inbtc *inbTestContainer) writeSwitchInlineQueryCurrentChat() {
	inbtc.name = "(IInlineButton).WriteSwitchInlineQueryCurrentChat()"
	inbtc.inputStr = []string{"string", "", "nothing", "a"}
	inbtc.isExpectedErr = []bool{false, false, false, true}
	inbtc.codeErr = []string{"", "", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteSwitchInlineQueryCurrentChat
}

func (inbtc *inbTestContainer) writeURL() {
	inbtc.name = "(IInlineButton).WriteURL()"
	inbtc.inputStr = []string{"https://youtube.com", "", "https://amazon.com", "https://t.me/l1qwie"}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteURL
}

func (inbtc *inbTestContainer) writeWebApp() {
	inbtc.name = "(IInlineButton).WriteWebApp()"
	inbtc.inputWebApp = []*types.WebAppInfo{{Url: "https://amazon.com"}, nil, {Url: "https://amazon.com"}, {Url: "https://amazon.com"}}
	inbtc.isExpectedErr = []bool{false, true, false, true}
	inbtc.codeErr = []string{"", "20", "", "10"}
	inbtc.amount, inbtc.until = 4, 2
	inbtc.buildF = putIButtonWriteWebApp
}

func (inb *inlineButtonT) startTest(part string, i int, t *testing.T) {
	switch f := inb.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, inb.name, inb.codeErr, inb.str, inb.isExpectedErr, i)
		checkError(f(inb.str), inb.isExpectedErr, inb.codeErr, t)
	case func(*types.CallbackGame) error:
		printTestLog(part, inb.name, inb.codeErr, inb.game, inb.isExpectedErr, i)
		checkError(f(inb.game), inb.isExpectedErr, inb.codeErr, t)
	case func(*types.LoginUrl) error:
		printTestLog(part, inb.name, inb.codeErr, inb.login, inb.isExpectedErr, i)
		checkError(f(inb.login), inb.isExpectedErr, inb.codeErr, t)
	case func(*types.SwitchInlineQueryChosenChat) error:
		printTestLog(part, inb.name, inb.codeErr, inb.sw, inb.isExpectedErr, i)
		checkError(f(inb.sw), inb.isExpectedErr, inb.codeErr, t)
	case func(*types.WebAppInfo) error:
		printTestLog(part, inb.name, inb.codeErr, inb.webapp, inb.isExpectedErr, i)
		checkError(f(inb.webapp), inb.isExpectedErr, inb.codeErr, t)
	case func() error:
		printTestLog(part, inb.name, inb.codeErr, true, inb.isExpectedErr, i)
		checkError(f(), inb.isExpectedErr, inb.codeErr, t)
	default:
		t.Fatalf("unexpected type of tested function: %T", f)
	}
}

func (inbtc *inbTestContainer) createTestArrays(msg *formatter.Message, t *testing.T) ([]UnitTest, []UnitTest) {
	var btn formatter.IInlineButton
	a, b := make([]UnitTest, inbtc.until), make([]UnitTest, inbtc.amount-inbtc.until)
	for i, j := 0, 0; i < inbtc.amount; i++ {
		if i < inbtc.until {
			if kb, err := msg.NewKeyboard().WriteInline(); err != nil {
				t.Fatal(err)
			} else {
				if err = kb.Set([]int{1, 2, 3}); err != nil {
					t.Fatal(err)
				}
				if btn, err = kb.NewButton(1, 1); err != nil {
					t.Fatal(err)
				}
			}
			a[i] = inbtc.buildF(*inbtc, btn, i)
		} else {
			if j%2 == 0 {
				if kb, err := msg.NewKeyboard().WriteInline(); err != nil {
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
			b[j] = inbtc.buildF(*inbtc, btn, i)
			j++
		}
	}
	return a, b
}

func (inbtc *inbTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	rpumentcontainer, doublecontainer := inbtc.createTestArrays(msg, t)
	check(rpumentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestIButtonWriteCallbackData(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeCallbackData()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteCallbackGame(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeCallbackGame()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButttonWriteLoginUrl(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeLoginUrl()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWritePay(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writePay()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteString(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeString()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteSwitchInlineQuery(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeSwitchInlineQuery()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteSwitchInlineQueryChosenChat(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeSwitchInlineQueryChosenChat()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteSwitchInlineQueryCurrentChat(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeSwitchInlineQueryCurrentChat()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteURL(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeURL()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}

func TestIButtonWriteWebApp(t *testing.T) {
	inbtc := new(inbTestContainer)
	inbtc.writeWebApp()
	msg := formatter.CreateEmpltyMessage()
	inbtc.mainLogic(msg, t)
}
