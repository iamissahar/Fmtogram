package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type inlineModeT struct {
	name          string
	str           string
	integer       int
	but           *types.InlineQueryResultsButton
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type inmTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	inputButton   []*types.InlineQueryResultsButton
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(inmode formatter.IInlineMode, in *inlineModeT, i int)
}

func (inmtc *inmTestContainer) byDefault(i int) *inlineModeT {
	return &inlineModeT{name: inmtc.name, isExpectedErr: inmtc.isExpectedErr[i], codeErr: inmtc.codeErr[i]}
}

func (inmtc *inmTestContainer) queryID(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteQueryID
	in.str = inmtc.inputStr[i]
}

func (inmtc *inmTestContainer) webAppQueryID(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteWebAppQueryID
	in.str = inmtc.inputStr[i]
}

func (inmtc *inmTestContainer) result(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteResult
}

func (inmtc *inmTestContainer) results(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteResults
	in.integer = inmtc.inputInt[i]
}

func (inmtc *inmTestContainer) cacheTime(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteCacheTime
	in.integer = inmtc.inputInt[i]
}

func (inmtc *inmTestContainer) isPersonal(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteIsPersonal
}

func (inmtc *inmTestContainer) nextOffset(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteNextOffset
	in.str = inmtc.inputStr[i]
}

func (inmtc *inmTestContainer) button(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteButton
	in.but = inmtc.inputButton[i]
}

func (inmtc *inmTestContainer) allowUsersChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteAllowUserChats
}

func (inmtc *inmTestContainer) allowBotChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteAllowUserChats
}

func (inmtc *inmTestContainer) allowGroupChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteAllowUserChats
}

func (inmtc *inmTestContainer) allowChannelChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
	in.testedFunc = inmode.WriteAllowUserChats
}

func (inmtc *inmTestContainer) writeQueryID() {
	inmtc.name = "(IInlineMode).WriteQueryID()"
	inmtc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	inmtc.isExpectedErr = []bool{false, true, false, true}
	inmtc.codeErr = []string{"", "20", "", "10"}
	inmtc.amount, inmtc.until = 4, 2
	inmtc.buildF = inmtc.queryID
}

func (inmtc *inmTestContainer) writeWebAppQueryID() {
	inmtc.name = "(IInlineMode).WriteWebAppQueryID()"
	inmtc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	inmtc.isExpectedErr = []bool{false, true, false, true}
	inmtc.codeErr = []string{"", "20", "", "10"}
	inmtc.amount, inmtc.until = 4, 2
	inmtc.buildF = inmtc.webAppQueryID
}

func (inmtc *inmTestContainer) writeResult() {
	inmtc.name = "(IInlineMode).WriteResult()"
	inmtc.isExpectedErr = []bool{false, false, true}
	inmtc.codeErr = []string{"", "", "10"}
	inmtc.amount, inmtc.until = 3, 1
	inmtc.buildF = inmtc.result
}

func (inmtc *inmTestContainer) writeResults() {
	inmtc.name = "(IInlineMode).WriteResults()"
	inmtc.inputInt = []int{0, -12124, -1, 1231, 34341234, 645, 1}
	inmtc.isExpectedErr = []bool{true, true, true, false, false, false, true}
	inmtc.codeErr = []string{"20", "20", "20", "", "", "", "10"}
	inmtc.amount, inmtc.until = 7, 5
	inmtc.buildF = inmtc.results
}

func (inmtc *inmTestContainer) writeCacheTime() {
	inmtc.name = "(IInlineMode).WriteCacheTime()"
	inmtc.inputInt = []int{0, -12124, -1, 1231, 34341234, 645, 1}
	inmtc.isExpectedErr = []bool{true, true, true, false, false, false, true}
	inmtc.codeErr = []string{"20", "20", "20", "", "", "", "10"}
	inmtc.amount, inmtc.until = 7, 5
	inmtc.buildF = inmtc.cacheTime
}

func (inmtc *inmTestContainer) writeIsPersonal() {
	inmtc.name = "(IInlineMode).WriteIsPersonal()"
	inmtc.isExpectedErr = []bool{false, false, true}
	inmtc.codeErr = []string{"", "", "10"}
	inmtc.amount, inmtc.until = 3, 1
	inmtc.buildF = inmtc.isPersonal
}

func (inmtc *inmTestContainer) writeNextOffset() {
	inmtc.name = "(IInlineMode).WriteNextOffset()"
	inmtc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	inmtc.isExpectedErr = []bool{false, false, false, true}
	inmtc.codeErr = []string{"", "", "", "10"}
	inmtc.amount, inmtc.until = 4, 2
	inmtc.buildF = inmtc.nextOffset
}

func (inmtc *inmTestContainer) writeButton() {
	inmtc.name = "(IInlineMode).WriteButton()"
	inmtc.inputButton = []*types.InlineQueryResultsButton{{Text: "nothing", WebApp: nil, StartParameter: "?"},
		{Text: "something"}, nil,
		{Text: "nothing", WebApp: nil}, {Text: "absolutely nothing", StartParameter: ""},
		{Text: "nothing", WebApp: nil, StartParameter: "?"}, {Text: "nothing", WebApp: nil, StartParameter: "?"}}
	inmtc.isExpectedErr = []bool{false, true, true, true, true, false, true}
	inmtc.codeErr = []string{"", "20", "20", "20", "20", "", "10"}
	inmtc.amount, inmtc.until = 7, 5
	inmtc.buildF = inmtc.button
}

func (inmtc *inmTestContainer) writeAllowUserChats() {
	inmtc.name = "(IInlineMode).WriteAllowUserChats()"
	inmtc.isExpectedErr = []bool{false, false, true}
	inmtc.codeErr = []string{"", "", "10"}
	inmtc.amount, inmtc.until = 3, 1
	inmtc.buildF = inmtc.allowUsersChats
}

func (inmtc *inmTestContainer) writeAllowBotChats() {
	inmtc.name = "(IInlineMode).WriteAllowBotChats()"
	inmtc.isExpectedErr = []bool{false, false, true}
	inmtc.codeErr = []string{"", "", "10"}
	inmtc.amount, inmtc.until = 3, 1
	inmtc.buildF = inmtc.allowBotChats
}

func (inmtc *inmTestContainer) writeAllowGroupChats() {
	inmtc.name = "(IInlineMode).WriteAllowGroupChats()"
	inmtc.isExpectedErr = []bool{false, false, true}
	inmtc.codeErr = []string{"", "", "10"}
	inmtc.amount, inmtc.until = 3, 1
	inmtc.buildF = inmtc.allowGroupChats
}

func (inmtc *inmTestContainer) writeAllowChannelChats() {
	inmtc.name = "(IInlineMode).WriteAllowChannelChats()"
	inmtc.isExpectedErr = []bool{false, false, true}
	inmtc.codeErr = []string{"", "", "10"}
	inmtc.amount, inmtc.until = 3, 1
	inmtc.buildF = inmtc.allowChannelChats
}

func (inm *inlineModeT) startTest(part string, i int, t *testing.T) {
	switch f := inm.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, inm.name, inm.codeErr, inm.str, inm.isExpectedErr, i)
		checkError(f(inm.str), inm.isExpectedErr, inm.codeErr, t)
	case func(int) error:
		printTestLog(part, inm.name, inm.codeErr, inm.integer, inm.isExpectedErr, i)
		checkError(f(inm.integer), inm.isExpectedErr, inm.codeErr, t)
	case func(*types.InlineQueryResultsButton) error:
		printTestLog(part, inm.name, inm.codeErr, inm.but, inm.isExpectedErr, i)
		checkError(f(inm.but), inm.isExpectedErr, inm.codeErr, t)
	case func() (formatter.IResult, error):
		printTestLog(part, inm.name, inm.codeErr, true, inm.isExpectedErr, i)
		_, err := f()
		checkError(err, inm.isExpectedErr, inm.codeErr, t)
	case func(int) ([]formatter.IResult, error):
		printTestLog(part, inm.name, inm.codeErr, inm.integer, inm.isExpectedErr, i)
		_, err := f(inm.integer)
		checkError(err, inm.isExpectedErr, inm.codeErr, t)
	case func() error:
		printTestLog(part, inm.name, inm.codeErr, true, inm.isExpectedErr, i)
		checkError(f(), inm.isExpectedErr, inm.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (inmtc *inmTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var inm formatter.IInlineMode
	a, b := make([]UnitTest, inmtc.until), make([]UnitTest, inmtc.amount-inmtc.until)
	for i, j := 0, 0; i < inmtc.amount; i++ {
		s := inmtc.byDefault(i)
		if i < inmtc.until {
			inmtc.buildF(msg.NewInlineMode(), s, i)
			a[i] = s
		} else {
			if j%2 == 0 {
				inm = msg.NewInlineMode()
			}
			inmtc.buildF(inm, s, i)
			b[j] = s
			j++
		}
	}
	return a, b
}

func (inmtc *inmTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	inmcontainer, doublecontainer := inmtc.createTestArrays(msg)
	check(inmcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestInlineModeWriteQueryID(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeQueryID()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteWebAppQueryID(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeWebAppQueryID()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteResult(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeResult()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteResults(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeResults()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteCacheTime(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeCacheTime()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteIsPersonal(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeIsPersonal()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteNextOffset(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeNextOffset()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteButton(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeButton()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteAllowUserChats(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeAllowUserChats()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteAllowBotChats(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeAllowBotChats()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteAllowGroupChats(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeAllowGroupChats()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}

func TestInlineModeWriteAllowChannelChats(t *testing.T) {
	inmtc := new(inmTestContainer)
	inmtc.writeAllowChannelChats()
	msg := formatter.CreateEmpltyMessage()
	inmtc.mainLogic(msg, t)
}
