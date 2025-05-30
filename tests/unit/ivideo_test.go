package unit

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type videoT struct {
	name          string
	str           string
	integer       int
	array         []*fmtogram.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type vdTestContainer struct {
	name          string
	inputStr      []string
	inputArr      [][]*fmtogram.MessageEntity
	inputInt      []int
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT
}

func putWriteVideoStorage(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteVideoStorage, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoTelegram(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteVideoTelegram, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoInternet(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteVideoInternet, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoCaption(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteCaption, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoCaptionEntities(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", 0, vdtc.inputArr[i], vd.WriteCaptionEntities, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoDuration(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", vdtc.inputInt[i], nil, vd.WriteDuration, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoHasSpoiler(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", 0, nil, vd.WriteHasSpoiler, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoHeight(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", vdtc.inputInt[i], nil, vd.WriteHeight, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoParseMode(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteParseMode, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoShowCaptionAboveMedia(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", 0, nil, vd.WriteShowCaptionAboveMedia, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoSupportStreaming(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", 0, nil, vd.WriteSupportsStreaming, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoThumbnail(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteThumbnail, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoThumbnailID(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteThumbnailID, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteVideoWidth(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, "", vdtc.inputInt[i], nil, vd.WriteWidth, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteCoverStorage(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteCoverStorage, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteCoverTelegram(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteCoverTelegram, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func putWriteCoverInternet(vdtc vdTestContainer, vd fmtogram.IVideo, i int) *videoT {
	return &videoT{vdtc.name, vdtc.inputStr[i], 0, nil, vd.WriteCoverInternet, vdtc.isExpectedErr[i], vdtc.codeErr[i]}
}

func (vdtc *vdTestContainer) writeVideoStorage() {
	vdtc.name = "(IVideo).WriteVideoStorage"
	vdtc.inputStr = []string{"../media/black.mp4", "", "../media/sound.mp3", "../media/musk.mp4", "../media/black.mp4"}
	vdtc.isExpectedErr = []bool{false, true, true, false, true}
	vdtc.codeErr = []string{"", "20", "12", "", "10"}
	vdtc.amount, vdtc.until = 5, 3
	vdtc.buildF = putWriteVideoStorage
}

func (vdtc *vdTestContainer) writeVideoTelegram() {
	vdtc.name = "(IVideo).WriteVideoTelegram"
	vdtc.inputStr = []string{"AS:LKDAKL:SD:KLASD", "", "SDK:LASKLJDJKLASKLGJAKLSJD", "something"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteVideoTelegram
}

func (vdtc *vdTestContainer) writeVideoInternet() {
	vdtc.name = "(IVideo).WriteVideoInternet"
	vdtc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteVideoInternet
}

func (vdtc *vdTestContainer) writeCaption() {
	vdtc.name = "(IVideo).WriteCaption"
	vdtc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteVideoCaption
}

func (vdtc *vdTestContainer) writeCaptionEntities() {
	vdtc.name = "(IVideo).WriteCaptionEntities"
	vdtc.inputArr = [][]*fmtogram.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	vdtc.isExpectedErr = []bool{false, true, true, false, true}
	vdtc.codeErr = []string{"", "20", "5", "", "10"}
	vdtc.amount, vdtc.until = 5, 3
	vdtc.buildF = putWriteVideoCaptionEntities
}

func (vdtc *vdTestContainer) writeDuration() {
	vdtc.name = "(IVideo).WriteDuration"
	vdtc.inputInt = []int{1231125, 0, 45, -123451561, 68923904, 67}
	vdtc.isExpectedErr = []bool{false, true, false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "20", "", "10"}
	vdtc.amount, vdtc.until = 6, 4
	vdtc.buildF = putWriteVideoDuration
}

func (vdtc *vdTestContainer) writeHasSpoiler() {
	vdtc.name = "(IVideo).WriteHasSpoiler"
	vdtc.isExpectedErr = []bool{false, false, true}
	vdtc.codeErr = []string{"", "", "10"}
	vdtc.amount, vdtc.until = 3, 1
	vdtc.buildF = putWriteVideoHasSpoiler
}

func (vdtc *vdTestContainer) writeHeight() {
	vdtc.name = "(IVideo).WriteHeight"
	vdtc.inputInt = []int{1231125, 0, 45, -123451561, 68923904, 67}
	vdtc.isExpectedErr = []bool{false, true, false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "20", "", "10"}
	vdtc.amount, vdtc.until = 6, 4
	vdtc.buildF = putWriteVideoHeight
}

func (vdtc *vdTestContainer) writeParseMode() {
	vdtc.name = "(IVideo).WriteParseMode"
	vdtc.inputStr = []string{fmtogram.HTML, fmtogram.Markdown, fmtogram.MarkdownV2, "", "something else", fmtogram.HTML, fmtogram.Markdown}
	vdtc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	vdtc.codeErr = []string{"", "", "", "20", "20", "", "10"}
	vdtc.amount, vdtc.until = 7, 5
	vdtc.buildF = putWriteVideoParseMode
}

func (vdtc *vdTestContainer) writeShowCaptionAboveMedia() {
	vdtc.name = "(IVideo).WriteShowCaptionAboveMedia"
	vdtc.isExpectedErr = []bool{false, false, true}
	vdtc.codeErr = []string{"", "", "10"}
	vdtc.amount, vdtc.until = 3, 1
	vdtc.buildF = putWriteVideoShowCaptionAboveMedia
}

func (vdtc *vdTestContainer) writeSupportStreaming() {
	vdtc.name = "(IVideo).WriteSupportStreaming"
	vdtc.isExpectedErr = []bool{false, false, true}
	vdtc.codeErr = []string{"", "", "10"}
	vdtc.amount, vdtc.until = 3, 1
	vdtc.buildF = putWriteVideoSupportStreaming
}

func (vdtc *vdTestContainer) writeThumbnail() {
	vdtc.name = "(IVideo).WriteThumbnail"
	vdtc.inputStr = []string{"../media/tel-aviv.jpg", "", "../media/sound.mp3", "../media/tel-aviv.jpg", "../media/tel-aviv.jpg"}
	vdtc.isExpectedErr = []bool{false, true, true, false, true}
	vdtc.codeErr = []string{"", "20", "12", "", "10"}
	vdtc.amount, vdtc.until = 5, 3
	vdtc.buildF = putWriteVideoThumbnail
}

func (vdtc *vdTestContainer) writeThumbnailID() {
	vdtc.name = "(IVideo).WriteThumbnail"
	vdtc.inputStr = []string{"AS:LKDAKL:SD:KLASD", "", "SDK:LASKLJDJKLASKLGJAKLSJD", "something"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteVideoThumbnailID
}

func (vdtc *vdTestContainer) writeWidth() {
	vdtc.name = "(IVideo).WriteWidth"
	vdtc.inputInt = []int{1231125, 0, 45, -123451561, 68923904, 67}
	vdtc.isExpectedErr = []bool{false, true, false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "20", "", "10"}
	vdtc.amount, vdtc.until = 6, 4
	vdtc.buildF = putWriteVideoWidth
}

func (vdtc *vdTestContainer) writeCoverStorage() {
	vdtc.name = "(IVideo).WriteCoverStorage()"
	vdtc.inputStr = []string{"ASKL:DKL:ASDKL:A", "", "ASL:DA", "KSDKLKSD"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteCoverStorage
}

func (vdtc *vdTestContainer) writeCoverTelegram() {
	vdtc.name = "(IVideo).WriteCoverTelegram()"
	vdtc.inputStr = []string{"ASKL:DKL:ASDKL:A", "", "ASL:DA", "KSDKLKSD"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteCoverTelegram
}

func (vdtc *vdTestContainer) writeCoverInternet() {
	vdtc.name = "(IVideo).WriteCoverInternet()"
	vdtc.inputStr = []string{"ASKL:DKL:ASDKL:A", "", "ASL:DA", "KSDKLKSD"}
	vdtc.isExpectedErr = []bool{false, true, false, true}
	vdtc.codeErr = []string{"", "20", "", "10"}
	vdtc.amount, vdtc.until = 4, 2
	vdtc.buildF = putWriteCoverInternet
}

func (vd *videoT) callStrF(testedF func(string) error, t *testing.T) {
	if !vd.isExpectedErr {
		if err := testedF(vd.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vd.str); err.Error() != vd.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vd *videoT) callIntF(testedF func(int) error, t *testing.T) {
	if !vd.isExpectedErr {
		if err := testedF(vd.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vd.integer); err.Error() != vd.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vd *videoT) callSliceF(testedF func([]*fmtogram.MessageEntity) error, t *testing.T) {
	if !vd.isExpectedErr {
		if err := testedF(vd.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vd.array); err.Error() != vd.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vd *videoT) callBoolF(testedF func() error, t *testing.T) {
	if !vd.isExpectedErr {
		if err := testedF(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(); err.Error() != vd.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vd *videoT) startTest(part string, i int, t *testing.T) {
	switch f := vd.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, vd.name, vd.codeErr, vd.str, vd.isExpectedErr, i, t)
		vd.callStrF(f, t)
	case func(int) error:
		printTestLog(part, vd.name, vd.codeErr, vd.integer, vd.isExpectedErr, i, t)
		vd.callIntF(f, t)
	case func([]*fmtogram.MessageEntity) error:
		printTestLog(part, vd.name, vd.codeErr, vd.array, vd.isExpectedErr, i, t)
		vd.callSliceF(f, t)
	case func() error:
		printTestLog(part, vd.name, vd.codeErr, true, vd.isExpectedErr, i, t)
		vd.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (vdtc *vdTestContainer) createTestArrays(msg *fmtogram.Message) ([]UnitTest, []UnitTest) {
	var vd fmtogram.IVideo
	a, b := make([]UnitTest, vdtc.until), make([]UnitTest, vdtc.amount-vdtc.until)
	for i, j := 0, 0; i < vdtc.amount; i++ {
		if i < vdtc.until {
			vd = fmtogram.NewVideo()
			a[i] = vdtc.buildF(*vdtc, vd, i)
		} else {
			if j%2 == 0 {
				vd = fmtogram.NewVideo()
			}
			b[j] = vdtc.buildF(*vdtc, vd, i)
			j++
		}
	}
	return a, b
}

func mainVideoLogic(msg *fmtogram.Message, vdtc vdTestContainer, t *testing.T) {
	videocontainer, doublecontainer := vdtc.createTestArrays(msg)
	check(videocontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteVideoStorage(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeVideoStorage()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoTelegram(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeVideoTelegram()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoInternet(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeVideoInternet()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoCaption(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeCaption()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoCaptionEntities(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeCaptionEntities()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoDuration(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeDuration()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoHasSpoiler(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeHasSpoiler()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoHeight(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeHeight()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoParseMode(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeParseMode()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoShowCaptionAboveMedia(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeShowCaptionAboveMedia()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoSupportStreaming(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeSupportStreaming()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoThumbnail(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeThumbnail()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoThumbnailID(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeThumbnailID()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteVideoWidth(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeWidth()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteCoverStorage(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeCoverStorage()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteCoverTelegram(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeCoverTelegram()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}

func TestWriteCoverInternet(t *testing.T) {
	t.Parallel()
	vdtc := new(vdTestContainer)
	vdtc.writeCoverInternet()
	msg := fmtogram.CreateEmpltyMessage()
	mainVideoLogic(msg, *vdtc, t)
}
