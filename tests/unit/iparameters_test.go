package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type paramsT struct {
	name          string
	integer       int
	str           string
	array         []*types.MessageEntity
	arrayInt      []int
	link          *types.LinkPreviewOptions
	replyP        *types.ReplyParameters
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type prmTestContainer struct {
	name          string
	inputInt      []int
	inputStr      []string
	inputArrInt   [][]int
	inputArr      [][]*types.MessageEntity
	inputLink     []*types.LinkPreviewOptions
	inputReplyP   []*types.ReplyParameters
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT
}

func putParamWriteDisableNotification(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, nil, nil, nil, prm.WriteDisableNotification, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteEntities(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", prmtc.inputArr[i], nil, nil, nil, prm.WriteEntities, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteLinkPreviewOptions(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, nil, prmtc.inputLink[i], nil, prm.WriteLinkPreviewOptions, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteMessageEffectID(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, prmtc.inputStr[i], nil, nil, nil, nil, prm.WriteMessageEffectID, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteMessageThreadID(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, prmtc.inputInt[i], "", nil, nil, nil, nil, prm.WriteMessageThreadID, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteMessageID(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, prmtc.inputInt[i], "", nil, nil, nil, nil, prm.WriteMessageID, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteMessageIDs(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, prmtc.inputArrInt[i], nil, nil, prm.WriteMessageIDs, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteCaption(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, prmtc.inputStr[i], nil, nil, nil, nil, prm.WriteCaption, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteParseMode(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, prmtc.inputStr[i], nil, nil, nil, nil, prm.WriteParseMode, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteProtectContent(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, nil, nil, nil, prm.WriteProtectContent, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteString(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, prmtc.inputStr[i], nil, nil, nil, nil, prm.WriteString, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteShowCaptionAboveMedia(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, nil, nil, nil, prm.WriteShowCaptionAboveMedia, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteReplyParameters(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, nil, nil, prmtc.inputReplyP[i], prm.WriteReplyParameters, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteAllowPaidBroadcast(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, "", nil, nil, nil, nil, prm.WriteAllowPaidBroadcast, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWriteStarCount(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, prmtc.inputInt[i], "", nil, nil, nil, nil, prm.WriteStarCount, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func putParamWritePayload(prmtc prmTestContainer, prm formatter.IParameters, i int) *paramsT {
	return &paramsT{prmtc.name, 0, prmtc.inputStr[i], nil, nil, nil, nil, prm.WritePayload, prmtc.isExpectedErr[i], prmtc.codeErr[i]}
}

func (prmtc *prmTestContainer) writeDisableNotification() {
	prmtc.name = "(IParameters).WriteDisableNotification"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteDisableNotification
}

func (prmtc *prmTestContainer) writeEntities() {
	prmtc.name = "(IParameters).WriteEntities"
	prmtc.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	prmtc.isExpectedErr = []bool{false, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "5", "", "10"}
	prmtc.amount, prmtc.until = 5, 3
	prmtc.buildF = putParamWriteEntities
}

func (prmtc *prmTestContainer) writeLinkPreviewOptions() {
	prmtc.name = "(IParameters).WriteLinkPreviewOptions"
	prmtc.inputLink = []*types.LinkPreviewOptions{{IsDisabled: true, URL: "https://youtube.com", PreferLargeMedia: true, ShowAboveText: true},
		nil,
		{IsDisabled: true, URL: "https://youtube.com", PreferLargeMedia: true, ShowAboveText: true}, {IsDisabled: true, URL: "https://youtube.com", PreferLargeMedia: true, ShowAboveText: true}}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteLinkPreviewOptions
}

func (prmtc *prmTestContainer) writeMessageEffectID() {
	prmtc.name = "(IParameters).WriteMessageEffectID"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteMessageEffectID
}

func (prmtc *prmTestContainer) writeMessageThreadID() {
	prmtc.name = "(IParameters).WriteMessageThreadID"
	prmtc.inputInt = []int{12331, 0, 222, 12345}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteMessageThreadID
}

func (prmtc *prmTestContainer) writeMessageID() {
	prmtc.name = "(IParameters).WriteMessageID"
	prmtc.inputInt = []int{55, 0, 12356, 12345}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteMessageID
}

func (prmtc *prmTestContainer) writeMessageIDs() {
	prmtc.name = "(IParameters).WriteMessageIDs"
	prmtc.inputArrInt = [][]int{{55, 777, 12356, 12345},
		{},
		{55, 777, 0, 12345},
		{55, 777, 12356, 12345}, {55, 777, 12356, 12345}}
	prmtc.isExpectedErr = []bool{false, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "5", "", "10"}
	prmtc.amount, prmtc.until = 5, 3
	prmtc.buildF = putParamWriteMessageIDs
}

func (prmtc *prmTestContainer) writeCaption() {
	prmtc.name = "(IParameters).WriteCaption"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteCaption
}

func (prmtc *prmTestContainer) writeParseMode() {
	prmtc.name = "(IParameters).WriteParseMode"
	prmtc.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	prmtc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	prmtc.codeErr = []string{"", "", "", "20", "20", "", "10"}
	prmtc.amount, prmtc.until = 7, 5
	prmtc.buildF = putParamWriteParseMode
}

func (prmtc *prmTestContainer) writeProtectContent() {
	prmtc.name = "(IParameters).WriteProtectContent"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteProtectContent
}

func (prmtc *prmTestContainer) writeString() {
	prmtc.name = "(IParameters).WriteString"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteString
}

func (prmtc *prmTestContainer) writeShowCaptionAboveMedia() {
	prmtc.name = "(IParameters).WriteShowCaptionAboveMedia"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteShowCaptionAboveMedia
}

func (prmtc *prmTestContainer) writeReplyParameters() {
	prmtc.name = "(IParameters).WriteReplyParameters"
	prmtc.inputReplyP = []*types.ReplyParameters{{MessageID: 999, Quote: "I am from the government and I am here to help"},
		nil,
		{MessageID: 999, Quote: "I am from the government and I am here to help"}, {MessageID: 999, Quote: "I am from the government and I am here to help"}}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteReplyParameters
}

func (prmtc *prmTestContainer) writeAllowPaidBroadcast() {
	prmtc.name = "(IParameters).WriteAllowPaidBroadcast"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteAllowPaidBroadcast
}

func (prmtc *prmTestContainer) writeStarCount() {
	prmtc.name = "(IParameters).WriteStarCount"
	prmtc.inputInt = []int{55, 0, 99, 45}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteStarCount
}

func (prmtc *prmTestContainer) writePayload() {
	prmtc.name = "(IParameters).WritePayload"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWritePayload
}

func (prm *paramsT) callStrF(f func(string) error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(prm.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(prm.str); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) callIntF(f func(int) error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(prm.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(prm.integer); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) callSliceF(f func([]*types.MessageEntity) error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(prm.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(prm.array); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) callSliceIntF(f func([]int) error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(prm.arrayInt); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(prm.arrayInt); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) callBoolF(f func() error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) callLinkF(f func(*types.LinkPreviewOptions) error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(prm.link); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(prm.link); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) callReplyPrmF(f func(*types.ReplyParameters) error, t *testing.T) {
	if !prm.isExpectedErr {
		if err := f(prm.replyP); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := f(prm.replyP); err.Error() != prm.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (prm *paramsT) startTest(part string, i int, t *testing.T) {
	switch f := prm.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, prm.name, prm.codeErr, prm.str, prm.isExpectedErr, i)
		prm.callStrF(f, t)
	case func(int) error:
		printTestLog(part, prm.name, prm.codeErr, prm.integer, prm.isExpectedErr, i)
		prm.callIntF(f, t)
	case func([]*types.MessageEntity) error:
		printTestLog(part, prm.name, prm.codeErr, prm.array, prm.isExpectedErr, i)
		prm.callSliceF(f, t)
	case func([]int) error:
		printTestLog(part, prm.name, prm.codeErr, prm.arrayInt, prm.isExpectedErr, i)
		prm.callSliceIntF(f, t)
	case func() error:
		printTestLog(part, prm.name, prm.codeErr, true, prm.isExpectedErr, i)
		prm.callBoolF(f, t)
	case func(*types.LinkPreviewOptions) error:
		printTestLog(part, prm.name, prm.codeErr, prm.link, prm.isExpectedErr, i)
		prm.callLinkF(f, t)
	case func(*types.ReplyParameters) error:
		printTestLog(part, prm.name, prm.codeErr, prm.replyP, prm.isExpectedErr, i)
		prm.callReplyPrmF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (prmtc *prmTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var prm formatter.IParameters
	a, b := make([]UnitTest, prmtc.until), make([]UnitTest, prmtc.amount-prmtc.until)
	for i, j := 0, 0; i < prmtc.amount; i++ {
		if i < prmtc.until {
			prm = msg.NewParameters()
			a[i] = prmtc.buildF(*prmtc, prm, i)
		} else {
			if j%2 == 0 {
				prm = msg.NewParameters()
			}
			b[j] = prmtc.buildF(*prmtc, prm, i)
			j++
		}
	}
	return a, b
}

func mainParametersLogic(msg *formatter.Message, prmtc prmTestContainer, t *testing.T) {
	prmcontainer, doublecontainer := prmtc.createTestArrays(msg)
	check(prmcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestParamWriteDisableNotification(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeDisableNotification()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteEntities(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeEntities()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteLinkPreviewOptions(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeLinkPreviewOptions()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageEffectID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageEffectID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageThreadID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageThreadID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageIDs(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageIDs()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteCaption(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeCaption()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteParseMode(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteProtectContent(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeProtectContent()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteString(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeString()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteShowCaptionAboveMedia(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeShowCaptionAboveMedia()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteReplyParameters(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeReplyParameters()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteAllowPaidBroadcast(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeAllowPaidBroadcast()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteStarCount(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeStarCount()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWritePayload(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writePayload()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}
