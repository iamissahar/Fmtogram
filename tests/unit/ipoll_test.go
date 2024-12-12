package unit

import (
	"strings"
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type pollT struct {
	name          string
	str           string
	integer       int
	array         []*types.MessageEntity
	array2        []*types.PollOption
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type pollTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	inputArr      [][]*types.MessageEntity
	inputArr2     [][]*types.PollOption
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT
}

func putWritePollQuestion(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, polltc.inputStr[i], 0, nil, nil, poll.WriteQuestion, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollQuestionParseMode(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, polltc.inputStr[i], 0, nil, nil, poll.WriteQuestionParseMode, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollQuestionEntities(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", 0, polltc.inputArr[i], nil, poll.WriteQuestionEntities, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollOptions(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", 0, nil, polltc.inputArr2[i], poll.WriteOptions, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollAnonymous(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", 0, nil, nil, poll.WriteAnonymous, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollType(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, polltc.inputStr[i], 0, nil, nil, poll.WriteType, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollAllowMultipleAnswers(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", 0, nil, nil, poll.WriteAllowMultipleAnswers, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollCorrectOptionID(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, polltc.inputStr[i], 0, nil, nil, poll.WriteCorrectOptionID, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollExplanation(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, polltc.inputStr[i], 0, nil, nil, poll.WriteExplanation, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollExplanationParseMode(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, polltc.inputStr[i], 0, nil, nil, poll.WriteExplanationParseMode, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollExplanationEntities(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", 0, polltc.inputArr[i], nil, poll.WriteExplanationEntities, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollOpenPeriod(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", polltc.inputInt[i], nil, nil, poll.WriteOpenPeriod, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollCloseDate(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", polltc.inputInt[i], nil, nil, poll.WriteCloseDate, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func putWritePollClosed(polltc pollTestContainer, poll formatter.IPoll, i int) *pollT {
	return &pollT{polltc.name, "", 0, nil, nil, poll.WriteClosed, polltc.isExpectedErr[i], polltc.codeErr[i]}
}

func (pollct *pollTestContainer) writeQuestion() {
	pollct.name = "(IPoll).WriteQuestion"
	pollct.inputStr = []string{"Who's gonna go?", "", strings.Repeat("g", 301), strings.Repeat("h", 300), "pollybody there?", "Are we flying to Bali tomorrow?"}
	pollct.isExpectedErr = []bool{false, true, true, false, false, true}
	pollct.codeErr = []string{"", "20", "20", "", "", "10"}
	pollct.amount, pollct.until = 6, 4
	pollct.buildF = putWritePollQuestion
}

func (pollct *pollTestContainer) writeQuestionParseMode() {
	pollct.name = "(IPoll).WriteQuestionParseMode"
	pollct.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	pollct.isExpectedErr = []bool{false, false, false, true, true, false, true}
	pollct.codeErr = []string{"", "", "", "20", "20", "", "10"}
	pollct.amount, pollct.until = 7, 5
	pollct.buildF = putWritePollQuestionParseMode
}

func (pollct *pollTestContainer) writeQuestionEntities() {
	pollct.name = "(IPoll).WriteQuestionEntities"
	pollct.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	pollct.isExpectedErr = []bool{false, true, true, false, true}
	pollct.codeErr = []string{"", "20", "5", "", "10"}
	pollct.amount, pollct.until = 5, 3
	pollct.buildF = putWritePollQuestionEntities
}

func (pollct *pollTestContainer) writeOptions() {
	pollct.name = "(IPoll).WriteOptions"
	pollct.inputArr2 = [][]*types.PollOption{
		{{Text: "Option 1"}, {Text: "Option 2"}, {Text: "Option 3"}},
		{},
		{{Text: "Option 3"}, nil, nil},
		{{Text: "Option 1"}, {Text: "Option 2"}, {Text: "Option 3"}}, {{Text: "Option 1"}, {Text: "Option 2"}, {Text: "Option 3"}}}
	pollct.isExpectedErr = []bool{false, true, true, false, true}
	pollct.codeErr = []string{"", "20", "5", "", "10"}
	pollct.amount, pollct.until = 5, 3
	pollct.buildF = putWritePollOptions
}

func (pollct *pollTestContainer) writeAnonymous() {
	pollct.name = "(IPoll).WriteAnonymous"
	pollct.isExpectedErr = []bool{false, false, true}
	pollct.codeErr = []string{"", "", "10"}
	pollct.amount, pollct.until = 3, 1
	pollct.buildF = putWritePollAnonymous
}

func (pollct *pollTestContainer) writeType() {
	pollct.name = "(IPoll).WriteType"
	pollct.inputStr = []string{"quiz", "", "regular", "al;sdl;asdl;", "1231512345", "quiz", "regular"}
	pollct.isExpectedErr = []bool{false, true, false, true, true, false, true}
	pollct.codeErr = []string{"", "20", "", "20", "20", "", "10"}
	pollct.amount, pollct.until = 7, 5
	pollct.buildF = putWritePollType
}

func (pollct *pollTestContainer) writeAllowMultipleAnswers() {
	pollct.name = "(IPoll).WriteAllowMultipleAnswers"
	pollct.isExpectedErr = []bool{false, false, true}
	pollct.codeErr = []string{"", "", "10"}
	pollct.amount, pollct.until = 3, 1
	pollct.buildF = putWritePollAllowMultipleAnswers
}

func (pollct *pollTestContainer) writeCorrectOptionID() {
	pollct.name = "(IPoll).WriteCorrectOptionID"
	pollct.inputStr = []string{"05", "", "9999", "al;sdl;asdl;", "010", "0245", "02"}
	pollct.isExpectedErr = []bool{false, true, true, true, false, false, true}
	pollct.codeErr = []string{"", "20", "20", "20", "", "", "10"}
	pollct.amount, pollct.until = 7, 5
	pollct.buildF = putWritePollCorrectOptionID
}

func (pollct *pollTestContainer) writeExplanation() {
	pollct.name = "(IPoll).WriteExplanation"
	pollct.inputStr = []string{"Ohh!", "", "Yes, it's me!", strings.Repeat("y", 200), strings.Repeat("y", 201), "0245", "02"}
	pollct.isExpectedErr = []bool{false, true, false, false, true, false, true}
	pollct.codeErr = []string{"", "20", "", "", "20", "", "10"}
	pollct.amount, pollct.until = 7, 5
	pollct.buildF = putWritePollExplanation
}

func (pollct *pollTestContainer) writeExplanationParseMode() {
	pollct.name = "(IPoll).WriteExplanationParseMode"
	pollct.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	pollct.isExpectedErr = []bool{false, false, false, true, true, false, true}
	pollct.codeErr = []string{"", "", "", "20", "20", "", "10"}
	pollct.amount, pollct.until = 7, 5
	pollct.buildF = putWritePollExplanationParseMode
}

func (pollct *pollTestContainer) writeExplanationEntities() {
	pollct.name = "(IPoll).WriteExplanationEntities"
	pollct.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	pollct.isExpectedErr = []bool{false, true, true, false, true}
	pollct.codeErr = []string{"", "20", "5", "", "10"}
	pollct.amount, pollct.until = 5, 3
	pollct.buildF = putWritePollExplanationEntities
}

func (pollct *pollTestContainer) writeOpenPeriod() {
	pollct.name = "(IPoll).WriteOpenPeriod"
	pollct.inputInt = []int{21304, 0, 4, 5, 599, 601, 600, 600, 66}
	pollct.isExpectedErr = []bool{true, true, true, false, false, true, false, false, true}
	pollct.codeErr = []string{"20", "20", "20", "", "", "20", "", "", "10"}
	pollct.amount, pollct.until = 9, 7
	pollct.buildF = putWritePollOpenPeriod
}

func (pollct *pollTestContainer) writeCloseDate() {
	pollct.name = "(IPoll).WriteCloseDate"
	pollct.inputInt = []int{21304, 0, 4, 5, 599, 601, 600, 600, 66}
	pollct.isExpectedErr = []bool{true, true, true, false, false, true, false, false, true}
	pollct.codeErr = []string{"20", "20", "20", "", "", "20", "", "", "10"}
	pollct.amount, pollct.until = 9, 7
	pollct.buildF = putWritePollCloseDate
}

func (pollct *pollTestContainer) writeClosed() {
	pollct.name = "(IPoll).WriteClosed"
	pollct.isExpectedErr = []bool{false, false, true}
	pollct.codeErr = []string{"", "", "10"}
	pollct.amount, pollct.until = 3, 1
	pollct.buildF = putWritePollClosed
}

func (poll *pollT) callStrF(testedF func(string) error, t *testing.T) {
	if !poll.isExpectedErr {
		if err := testedF(poll.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(poll.str); err.Error() != poll.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (poll *pollT) callIntF(testedF func(int) error, t *testing.T) {
	if !poll.isExpectedErr {
		if err := testedF(poll.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(poll.integer); err.Error() != poll.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (poll *pollT) callSliceF(testedF func([]*types.MessageEntity) error, t *testing.T) {
	if !poll.isExpectedErr {
		if err := testedF(poll.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(poll.array); err.Error() != poll.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (poll *pollT) callSliceF2(testedF func([]*types.PollOption) error, t *testing.T) {
	if !poll.isExpectedErr {
		if err := testedF(poll.array2); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(poll.array2); err.Error() != poll.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (poll *pollT) callBoolF(testedF func() error, t *testing.T) {
	if !poll.isExpectedErr {
		if err := testedF(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(); err.Error() != poll.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (poll *pollT) startTest(part string, i int, t *testing.T) {
	switch f := poll.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, poll.name, poll.codeErr, poll.str, poll.isExpectedErr, i)
		poll.callStrF(f, t)
	case func(int) error:
		printTestLog(part, poll.name, poll.codeErr, poll.integer, poll.isExpectedErr, i)
		poll.callIntF(f, t)
	case func([]*types.MessageEntity) error:
		printTestLog(part, poll.name, poll.codeErr, poll.array, poll.isExpectedErr, i)
		poll.callSliceF(f, t)
	case func([]*types.PollOption) error:
		printTestLog(part, poll.name, poll.codeErr, poll.array2, poll.isExpectedErr, i)
		poll.callSliceF2(f, t)
	case func() error:
		printTestLog(part, poll.name, poll.codeErr, true, poll.isExpectedErr, i)
		poll.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (polltc *pollTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var poll formatter.IPoll
	a, b := make([]UnitTest, polltc.until), make([]UnitTest, polltc.amount-polltc.until)
	for i, j := 0, 0; i < polltc.amount; i++ {
		if i < polltc.until {
			poll = msg.NewPoll()
			a[i] = polltc.buildF(*polltc, poll, i)
		} else {
			if j%2 == 0 {
				poll = msg.NewPoll()
			}
			b[j] = polltc.buildF(*polltc, poll, i)
			j++
		}
	}
	return a, b
}

func mainPollLogic(msg *formatter.Message, polltc pollTestContainer, t *testing.T) {
	pollimationcontainer, doublecontainer := polltc.createTestArrays(msg)
	check(pollimationcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWritePollQuestion(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeQuestion()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollQuestionParseMode(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeQuestionParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollQuestionEntities(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeQuestionEntities()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollOptions(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeOptions()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollAnonymous(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeAnonymous()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollType(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeType()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollAllowMultipleAnswers(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeAllowMultipleAnswers()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollCorrectOptionID(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeCorrectOptionID()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollExplanation(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeExplanation()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollExplanationParseMode(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeExplanationParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollExplanationEntities(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeExplanationEntities()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollOpenPeriod(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeOpenPeriod()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollCloseDate(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeCloseDate()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}

func TestWritePollClosed(t *testing.T) {
	pollct := new(pollTestContainer)
	pollct.writeClosed()
	msg := formatter.CreateEmpltyMessage()
	mainPollLogic(msg, *pollct, t)
}
