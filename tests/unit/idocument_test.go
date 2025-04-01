package unit

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type documentT struct {
	name          string
	str           string
	array         []*fmtogram.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type docTestContainer struct {
	name          string
	inputStr      []string
	inputArr      [][]*fmtogram.MessageEntity
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT
}

func putWriteDocumentStorage(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteDocumentStorage, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentTelegram(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteDocumentTelegram, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentInternet(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteDocumentInternet, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentCaption(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteCaption, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentCaptionEntities(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, "", doctc.inputArr[i], doc.WriteCaptionEntities, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentDisableContentTypeDetection(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, "", nil, doc.WriteDisableContentTypeDetection, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentParseMode(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteParseMode, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentThumbnail(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteThumbnail, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func putWriteDocumentThumbnailID(doctc docTestContainer, doc fmtogram.IDocument, i int) *documentT {
	return &documentT{doctc.name, doctc.inputStr[i], nil, doc.WriteThumbnailID, doctc.isExpectedErr[i], doctc.codeErr[i]}
}

func (doctc *docTestContainer) writeDocumentStorage() {
	doctc.name = "(IDocument).WriteDocumentStorage"
	doctc.inputStr = []string{"../media/Resume.pdf", "", "../media/sound.mp3", "../media/Resume.pdf", "../media/Resume.pdf"}
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
	doctc.inputArr = [][]*fmtogram.MessageEntity{
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
	doctc.inputStr = []string{fmtogram.HTML, fmtogram.Markdown, fmtogram.MarkdownV2, "", "something else", fmtogram.HTML, fmtogram.Markdown}
	doctc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	doctc.codeErr = []string{"", "", "", "20", "20", "", "10"}
	doctc.amount, doctc.until = 7, 5
	doctc.buildF = putWriteDocumentParseMode
}

func (doctc *docTestContainer) writeThumbnail() {
	doctc.name = "(IDocument).WriteThumbnail"
	doctc.inputStr = []string{"../media/tel-aviv.jpg", "", "../media/sound.mp3", "../media/tel-aviv.jpg", "../media/tel-aviv.jpg"}
	doctc.isExpectedErr = []bool{false, true, true, false, true}
	doctc.codeErr = []string{"", "20", "12", "", "10"}
	doctc.amount, doctc.until = 5, 3
	doctc.buildF = putWriteDocumentThumbnail
}

func (doctc *docTestContainer) writeThumbnailID() {
	doctc.name = "(IDocument).WriteThumbnailID"
	doctc.inputStr = []string{".:KLKL:ASK:LDKL:ASD:KLK:LL:KKL:SAHJDAHSDGHAD", "", "ASDKJKJASDHJKASDJ", ".asdaasfaDJKLASJHKDA"}
	doctc.isExpectedErr = []bool{false, true, false, true}
	doctc.codeErr = []string{"", "20", "", "10"}
	doctc.amount, doctc.until = 4, 2
	doctc.buildF = putWriteDocumentThumbnailID
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

func (doc *documentT) callSliceF(testedF func([]*fmtogram.MessageEntity) error, t *testing.T) {
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
	switch f := doc.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, doc.name, doc.codeErr, doc.str, doc.isExpectedErr, i, t)
		doc.callStrF(f, t)
	case func([]*fmtogram.MessageEntity) error:
		printTestLog(part, doc.name, doc.codeErr, doc.array, doc.isExpectedErr, i, t)
		doc.callSliceF(f, t)
	case func() error:
		printTestLog(part, doc.name, doc.codeErr, true, doc.isExpectedErr, i, t)
		doc.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (doctc *docTestContainer) createTestArrays(msg *fmtogram.Message) ([]UnitTest, []UnitTest) {
	var doc fmtogram.IDocument
	a, b := make([]UnitTest, doctc.until), make([]UnitTest, doctc.amount-doctc.until)
	for i, j := 0, 0; i < doctc.amount; i++ {
		if i < doctc.until {
			doc = fmtogram.NewDocument()
			a[i] = doctc.buildF(*doctc, doc, i)
		} else {
			if j%2 == 0 {
				doc = fmtogram.NewDocument()
			}
			b[j] = doctc.buildF(*doctc, doc, i)
			j++
		}
	}
	return a, b
}

func mainDocumentLogic(msg *fmtogram.Message, doctc docTestContainer, t *testing.T) {
	documentcontainer, doublecontainer := doctc.createTestArrays(msg)
	check(documentcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteDocumentStorage(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeDocumentStorage()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentTelegram(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeDocumentTelegram()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentInternet(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeDocumentInternet()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentCaption(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeCaption()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentCaptionEntities(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeCaptionEntities()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentDisableContentTypeDetection(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeDisableContentTypeDetection()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentParseMode(t *testing.T) {
	t.Parallel()
	doctc := new(docTestContainer)
	doctc.writeParseMode()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *doctc, t)
}

func TestWriteDocumentThumbnail(t *testing.T) {
	t.Parallel()
	vdtc := new(docTestContainer)
	vdtc.writeThumbnail()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *vdtc, t)
}

func TestWriteDocumentThumbnailID(t *testing.T) {
	t.Parallel()
	vdtc := new(docTestContainer)
	vdtc.writeThumbnailID()
	msg := fmtogram.CreateEmpltyMessage()
	mainDocumentLogic(msg, *vdtc, t)
}
