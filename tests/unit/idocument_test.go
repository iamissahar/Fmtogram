package unit

import (
	"fmt"
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type documentT struct {
	name          string
	str           string
	integer       int
	array         []*types.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type docTestContainer struct {
	name          string
	inputStr      []string
	inputArr      [][]*types.MessageEntity
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(doctc docTestContainer, doc formatter.IDocument, i int) *documentT
}

func putWriteDocumentStorage(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteDocumentStorage, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentTelegram(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteDocumentTelegram, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentInternet(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteDocumentInternet, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentCaption(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteCaption, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentCaptionEntities(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, "", 0, doctc.inputArr[i], doc.WriteCaptionEntities, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentDisableContentTypeDetection(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, "", 0, nil, doc.WriteDisableContentTypeDetection, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentParseMode(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteParseMode, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentThumbnailStorage(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteThumbnailStorage, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentThumbnailTelegram(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteThumbnailTelegram, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentThumbnailInternet(doctc docTestContainer, doc formatter.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], 0, nil, doc.WriteThumbnailInternet, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func (doctc *docTestContainer) writeDocumentStorage() {
	doctc.name = "(IDocument).WriteDocumentStorage"
	doctc.inputStr = []string{"../media_test/Resume.pdf", "", "../media_test/sound.mp3", "../media_test/Resume.pdf", "../media_test/Resume.pdf"}
	doctc.isExpectedErr = []bool{false, true, false, false, true}
	doctc.codeErr = []string{"", "20", "", "", "10"}
	doctc.amount, doctc.until = 5, 3
	doctc.buildF = putWriteDocumentStorage
}

func (doctc *docTestContainer) writeDocumentTelegram() {
	doctc.name = "(IDocument).WriteDocumentTelegram"
	doctc.inputStr = []string{"A:SKLDKL:ASDKL:ASL:KDA:KL", "", "ASLKDJAKLSDJKLASJKLDLK", "ASKLJDJKLASLDJKALSJKD"}
	doctc.isExpectedErr = []bool{false, true, false, true}
	doctc.codeErr = []string{"", "20", "", "10"}
	doctc.amount, doctc.until = 4, 2
	doctc.buildF = putWriteDocumentTelegram
}

func (doctc *docTestContainer) writeDocumentInternet() {
	doctc.name = "(IDocument).WriteDocumentInternet"
	doctc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	doctc.isExpectedErr = []bool{false, true, false, true}
	doctc.codeErr = []string{"", "20", "", "10"}
	doctc.amount, doctc.until = 4, 2
	doctc.buildF = putWriteDocumentInternet
}

func (doctc *docTestContainer) writeCaption() {
	doctc.name = "(IDocument).WriteCaption"
	doctc.inputStr = []string{"SHALOM!", "", "Just a small caption", "something like a caption"}
	doctc.isExpectedErr = []bool{false, true, false, true}
	doctc.codeErr = []string{"", "20", "", "10"}
	doctc.amount, doctc.until = 4, 2
	doctc.buildF = putWriteDocumentCaption
}

func (doctc *docTestContainer) writeCaptionEntities() {
	doctc.name = "(IDocument).WriteCaptionEntities"
	doctc.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	doctc.isExpectedErr = []bool{false, true, true, false, true}
	doctc.codeErr = []string{"", "20", "5", "", "10"}
	doctc.amount, doctc.until = 5, 3
	doctc.buildF = putWriteDocumentCaptionEntities
}

func (doctc *docTestContainer) writeDisableContentTypeDetection() {
	doctc.name = "(IDocument).WriteDisableContentTypeDetection"
	doctc.isExpectedErr = []bool{false, false, true}
	doctc.codeErr = []string{"", "", "10"}
	doctc.amount, doctc.until = 3, 1
	doctc.buildF = putWriteDocumentDisableContentTypeDetection
}

func (doctc *docTestContainer) writeParseMode() {
	doctc.name = "(IDocument).WriteParseMode"
	doctc.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	doctc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	doctc.codeErr = []string{"", "", "", "20", "20", "", "10"}
	doctc.amount, doctc.until = 7, 5
	doctc.buildF = putWriteDocumentParseMode
}

func (doctc *docTestContainer) writeThumbnailStorage() {
	doctc.name = "(IDocument).WriteThumbnailStorage"
	doctc.inputStr = []string{"../media_test/tel-aviv.jpg", "", "../media_test/sound.mp3", "../media_test/tel-aviv.jpg", "../media_test/tel-aviv.jpg"}
	doctc.isExpectedErr = []bool{false, true, true, false, true}
	doctc.codeErr = []string{"", "20", "12", "", "10"}
	doctc.amount, doctc.until = 5, 3
	doctc.buildF = putWriteDocumentThumbnailStorage
}

func (doctc *docTestContainer) writeThumbnailTelegram() {
	doctc.name = "(IDocument).WriteThumbnailTelegram"
	doctc.inputStr = []string{"ASL:KDKAOL:SLK:@#$!:L", "", "A:LSKDKL:ASK:DLASKL:DASD", ")()*()#I@!K:OLAS:KLDAS:L"}
	doctc.isExpectedErr = []bool{false, true, false, true}
	doctc.codeErr = []string{"", "20", "", "10"}
	doctc.amount, doctc.until = 4, 2
	doctc.buildF = putWriteDocumentThumbnailTelegram
}

func (doctc *docTestContainer) writeThumbnailInternet() {
	doctc.name = "(IDocument).WriteThumbnailInternet"
	doctc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", ")()*()#I@!K:OLAS:KLDAS:L"}
	doctc.isExpectedErr = []bool{false, true, false, true}
	doctc.codeErr = []string{"", "20", "", "10"}
	doctc.amount, doctc.until = 4, 2
	doctc.buildF = putWriteDocumentThumbnailInternet
}

func (doc *documentT) callStrF(testedF func(string) error, t *testing.T) {
	if !doc.isExpectedErr {
		if err := testedF(doc.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(doc.str); err.Error() != doc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (doc *documentT) callSliceF(testedF func([]*types.MessageEntity) error, t *testing.T) {
	if !doc.isExpectedErr {
		if err := testedF(doc.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(doc.array); err.Error() != doc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (doc *documentT) callBoolF(testedF func() error, t *testing.T) {
	if !doc.isExpectedErr {
		if err := testedF(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(); err.Error() != doc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (doc *documentT) startTest(part string, i int, t *testing.T) {
	t.Logf(part, fmt.Sprintf(logMsg, doc.name, doc.isExpectedErr, doc.codeErr, i))
	switch f := doc.testedFunc.(type) {
	case func(string) error:
		doc.callStrF(f, t)
	case func([]*types.MessageEntity) error:
		doc.callSliceF(f, t)
	case func() error:
		doc.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (doctc *docTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var doc formatter.IDocument
	a, b := make([]UnitTest, doctc.until), make([]UnitTest, doctc.amount-doctc.until)
	for i, j := 0, 0; i < doctc.amount; i++ {
		if i < doctc.until {
			doc = msg.NewDocument()
			a[i] = doctc.buildF(*doctc, doc, i)
		} else {
			if j%2 == 0 {
				doc = msg.NewDocument()
			}
			b[j] = doctc.buildF(*doctc, doc, i)
			j++
		}
	}
	return a, b
}

func mainDocumentLogic(msg *formatter.Message, doctc docTestContainer, t *testing.T) {
	documentcontainer, doublecontainer := doctc.createTestArrays(msg)
	check(documentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteDocumentStorage(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeDocumentStorage()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentTelegram(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeDocumentTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentInternet(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeDocumentInternet()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentCaption(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeCaption()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentCaptionEntities(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeCaptionEntities()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentDisableContentTypeDetection(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeDisableContentTypeDetection()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentParseMode(t *testing.T) {
	doctc := new(docTestContainer)
	doctc.writeParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentThumbnailStorage(t *testing.T) {
	vdtc := new(docTestContainer)
	vdtc.writeThumbnailStorage()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *vdtc, t)
}

func TestWriteDocumentThumbnailTelegram(t *testing.T) {
	vdtc := new(docTestContainer)
	vdtc.writeThumbnailTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *vdtc, t)
}

func TestWriteDocumentThumbnailInternet(t *testing.T) {
	vdtc := new(docTestContainer)
	vdtc.writeThumbnailInternet()
	msg := formatter.CreateEmpltyMessage()
	mainDocumentLogic(msg, *vdtc, t)
}
