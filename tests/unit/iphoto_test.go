package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type photoT struct {
	name          string
	str           string
	array         []*types.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type phTestContainer struct {
	name          string
	inputStr      []string
	inputArr      [][]*types.MessageEntity
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(at phTestContainer, ph formatter.IPhoto, i int) *photoT
}

func putWritePhotoStorage(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, pt.inputStr[i], nil, ph.WritePhotoStorage, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoTelegram(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, pt.inputStr[i], nil, ph.WritePhotoTelegram, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoInternet(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, pt.inputStr[i], nil, ph.WritePhotoInternet, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoCaption(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, pt.inputStr[i], nil, ph.WriteCaption, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoParseMode(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, pt.inputStr[i], nil, ph.WriteParseMode, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoCaptionEntities(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, "", pt.inputArr[i], ph.WriteCaptionEntities, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoShowCaptionAboveMedia(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, "", nil, ph.WriteShowCaptionAboveMedia, pt.isExpectedErr[i], pt.codeErr[i]}
}

func putWritePhotoHasSpoiler(pt phTestContainer, ph formatter.IPhoto, i int) *photoT {
	return &photoT{pt.name, "", nil, ph.WriteHasSpoiler, pt.isExpectedErr[i], pt.codeErr[i]}
}

func (t *phTestContainer) writePhotoStorage() {
	t.name = "(IPhoto).WritePhotoStorage"
	t.inputStr = []string{"../media_test/tel-aviv.jpg", "", "../media_test/sound.mp3", "../media_test/tel-aviv.jpg", "../media_test/photo1.jpg"}
	t.isExpectedErr = []bool{false, true, true, false, true}
	t.codeErr = []string{"", "20", "12", "", "10"}
	t.amount, t.until = 5, 3
	t.buildF = putWritePhotoStorage
}

func (t *phTestContainer) writePhotoTelegram() {
	t.name = "(IPhoto).WritePhotoTelegram"
	t.inputStr = []string{"AS:LKDAKL:SD:KLASD", "", "SDK:LASKLJDJKLASKLGJAKLSJD", "something"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWritePhotoTelegram
}

func (t *phTestContainer) writePhotoInternet() {
	t.name = "(IPhoto).WritePhotoInternet"
	t.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWritePhotoInternet
}

func (t *phTestContainer) writeCaption() {
	t.name = "(IPhoto).WriteCaption"
	t.inputStr = []string{"Caption", "", "SHALOM!", "PRIVET"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWritePhotoCaption
}

func (t *phTestContainer) writeParseMode() {
	t.name = "(IPhoto).WriteParseMode"
	t.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	t.isExpectedErr = []bool{false, false, false, true, true, false, true}
	t.codeErr = []string{"", "", "", "20", "20", "", "10"}
	t.amount, t.until = 7, 5
	t.buildF = putWritePhotoParseMode
}

func (t *phTestContainer) writeCaptionEntities() {
	t.name = "(IPhoto).WriteCaptionEntities"
	t.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	t.isExpectedErr = []bool{false, true, true, false, true}
	t.codeErr = []string{"", "20", "5", "", "10"}
	t.amount, t.until = 5, 3
	t.buildF = putWritePhotoCaptionEntities
}

func (t *phTestContainer) writeShowCaptionAboveMedia() {
	t.name = "(IPhoto).WriteShowCaptionAboveMedia"
	t.isExpectedErr = []bool{false, false, true}
	t.codeErr = []string{"", "", "10"}
	t.amount, t.until = 3, 1
	t.buildF = putWritePhotoShowCaptionAboveMedia
}

func (t *phTestContainer) writeHasSpoiler() {
	t.name = "(IPhoto).WriteHasSpoiler"
	t.isExpectedErr = []bool{false, false, true}
	t.codeErr = []string{"", "", "10"}
	t.amount, t.until = 3, 1
	t.buildF = putWritePhotoHasSpoiler
}

func (ph *photoT) callBoolF(testedFunc func() error, t *testing.T) {
	if !ph.isExpectedErr {
		if err := testedFunc(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedFunc(); err.Error() != ph.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ph *photoT) callStrF(testedFunc func(string) error, data string, t *testing.T) {
	if !ph.isExpectedErr {
		if err := testedFunc(data); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedFunc(data); err.Error() != ph.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ph *photoT) callSliceF(testedFunc func([]*types.MessageEntity) error, data []*types.MessageEntity, t *testing.T) {
	if !ph.isExpectedErr {
		if err := testedFunc(data); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedFunc(data); err.Error() != ph.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ph *photoT) startTest(part string, i int, t *testing.T) {
	switch f := ph.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, ph.name, ph.codeErr, ph.str, ph.isExpectedErr, i)
		ph.callStrF(f, ph.str, t)
	case func([]*types.MessageEntity) error:
		printTestLog(part, ph.name, ph.codeErr, ph.array, ph.isExpectedErr, i)
		ph.callSliceF(f, ph.array, t)
	case func() error:
		printTestLog(part, ph.name, ph.codeErr, true, ph.isExpectedErr, i)
		ph.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (pt *phTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var ph formatter.IPhoto
	a, b := make([]UnitTest, pt.until), make([]UnitTest, pt.amount-pt.until)
	for i, j := 0, 0; i < pt.amount; i++ {
		if i < pt.until {
			ph = msg.NewPhoto()
			a[i] = pt.buildF(*pt, ph, i)
		} else {
			if j%2 == 0 {
				ph = msg.NewPhoto()
			}
			b[j] = pt.buildF(*pt, ph, i)
			j++
		}
	}
	return a, b
}

func mainPhotoLogic(msg *formatter.Message, pt phTestContainer, t *testing.T) {
	photoholder, doublephoto := pt.createTestArrays(msg)
	check(photoholder, t)
	doubleCheck(doublephoto, t)
}

func TestWritePhotoStorage(t *testing.T) {
	pt := new(phTestContainer)
	pt.writePhotoStorage()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}

func TestWritePhotoTelegram(t *testing.T) {
	pt := new(phTestContainer)
	pt.writePhotoTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}

func TestWritePhotoInternet(t *testing.T) {
	pt := new(phTestContainer)
	pt.writePhotoInternet()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}

func TestWritePhotoCaption(t *testing.T) {
	pt := new(phTestContainer)
	pt.writeCaption()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}

func TestWritePhotoParseMode(t *testing.T) {
	pt := new(phTestContainer)
	pt.writeParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}

func TestWritePhotoCaptionEntities(t *testing.T) {
	pt := new(phTestContainer)
	pt.writeCaptionEntities()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}

func TestWritePhotoShowCaptionAboveMedia(t *testing.T) {
	pt := new(phTestContainer)
	pt.writeShowCaptionAboveMedia()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}
func TestWritePhotoHasSpoiler(t *testing.T) {
	pt := new(phTestContainer)
	pt.writeHasSpoiler()
	msg := formatter.CreateEmpltyMessage()
	mainPhotoLogic(msg, *pt, t)
}
