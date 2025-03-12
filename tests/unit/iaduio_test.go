package unit

import (
	"strings"
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
	"github.com/iamissahar/Fmtogram/types"
)

const (
	errMsg string = "unexpected error: %s"
	logMsg string = "\tFunction under test: %s\n\tInput: %v (%T)\n\tHaving error: %v\n\tError code: %s\n\tIndex in an array: %d"
)

type audio struct {
	name          string
	str           string
	integer       int
	array         []*types.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type audiotest struct {
	name          string
	inputStr      []string
	inputInt      []int
	inputArr      [][]*types.MessageEntity
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(at audiotest, ad formatter.IAudio, i int) *audio
}

func putWriteAudioStorage(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteAudioStorage, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteAudioTelegram(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteAudioTelegram, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteAudioInternet(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteAudioInternet, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteCaption(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteCaption, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteCaptionEntities(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, "", 0, at.inputArr[i], ad.WriteCaptionEntities, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteDuration(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, "", at.inputInt[i], nil, ad.WriteDuration, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteParseMode(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteParseMode, at.isExpectedErr[i], at.codeErr[i]}
}

func putWritePerformer(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WritePerformer, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteThumbnail(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteThumbnail, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteThumbnailID(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteThumbnailID, at.isExpectedErr[i], at.codeErr[i]}
}

func putWriteTitle(at audiotest, ad formatter.IAudio, i int) *audio {
	return &audio{at.name, at.inputStr[i], 0, nil, ad.WriteTitle, at.isExpectedErr[i], at.codeErr[i]}
}

func (t *audiotest) writeAudioStorage() {
	t.name = "(IAudio).WriteAudioStorage"
	t.inputStr = []string{"../media/sound.mp3", "", "../media/tel-aviv.jpg", "../media/sound.mp3", "../media/sound.mp3"}
	t.isExpectedErr = []bool{false, true, true, false, true}
	t.codeErr = []string{"", "20", "12", "", "10"}
	t.amount, t.until = 5, 3
	t.buildF = putWriteAudioStorage
}

func (t *audiotest) writeAudioTelegram() {
	t.name = "(IAudio).WriteAudioTelegram"
	t.inputStr = []string{":LASDKLASLKDAL:SD", "", "KLJASD!@#**&(OLED)()(AAASDASDIKLP)", "AKJSDKJASDJKKJAJKSGFIU"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWriteAudioTelegram
}

func (t *audiotest) writeAudioInternet() {
	t.name = "(IAudio).WriteAudioInternet"
	t.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "AKJSDKJASDJKKJAJKSGFIU"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWriteAudioInternet
}

func (t *audiotest) writeCaption() {
	t.name = "(IAudio).WriteCaption"
	t.inputStr = []string{"ASDASDASFALK:DL:AS", "", strings.Repeat("a", 1024), strings.Repeat("a", 1025), "Something like a small text", "ALO<#?"}
	t.isExpectedErr = []bool{false, true, false, true, false, true}
	t.codeErr = []string{"", "20", "", "20", "", "10"}
	t.amount, t.until = 5, 3
	t.buildF = putWriteCaption
}

func (t *audiotest) writeCaptionEntities() {
	t.name = "(IAudio).WriteCaptionEntities"
	t.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	t.isExpectedErr = []bool{false, true, true, false, true}
	t.codeErr = []string{"", "20", "5", "", "10"}
	t.amount, t.until = 5, 3
	t.buildF = putWriteCaptionEntities
}

func (t *audiotest) writeDuration() {
	t.name = "(IAudio).WriteDuration"
	t.inputInt = []int{123125, 0, 111, 222}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWriteDuration
}

func (t *audiotest) writeParseMode() {
	t.name = "(IAudio).WriteParseMode"
	t.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	t.isExpectedErr = []bool{false, false, false, true, true, false, true}
	t.codeErr = []string{"", "", "", "20", "20", "", "10"}
	t.amount, t.until = 7, 5
	t.buildF = putWriteParseMode
}

func (t *audiotest) writePerformer() {
	t.name = "(IAudio).WritePerformer"
	t.inputStr = []string{"AKLSDKLASDKLALD", "", "something", "somehow"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWritePerformer
}

func (t *audiotest) writeThumbnail() {
	t.name = "(IAudio).WriteThumbnail"
	t.inputStr = []string{"../media/tel-aviv.jpg", "", "../media/sound.mp3", "../media/tel-aviv.jpg", "../media/tel-aviv.jpg"}
	t.isExpectedErr = []bool{false, true, true, false, true}
	t.codeErr = []string{"", "20", "12", "", "10"}
	t.amount, t.until = 5, 3
	t.buildF = putWriteThumbnail
}

func (t *audiotest) writeThumbnailID() {
	t.name = "(IAudio).WriteThumbnailID"
	t.inputStr = []string{"AO:KSDKLASLDK:LAKSDA", "", "ASL:KDKL:ASD:KLASDASD", "A:OLD:IOLASOIP:FDIKOA:KLKL::KOL"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWriteThumbnailID
}

func (t *audiotest) writeTitle() {
	t.name = "(IAudio).WriteTitle"
	t.inputStr = []string{"Shalom", "", "Selam", "Shto to na slovyanskom"}
	t.isExpectedErr = []bool{false, true, false, true}
	t.codeErr = []string{"", "20", "", "10"}
	t.amount, t.until = 4, 2
	t.buildF = putWriteTitle
}

func (ad *audio) callStrF(testedFunc func(string) error, t *testing.T) {
	if !ad.isExpectedErr {
		if err := testedFunc(ad.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedFunc(ad.str); err.Error() != ad.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ad *audio) callIntF(testedFunc func(int) error, t *testing.T) {
	if !ad.isExpectedErr {
		if err := testedFunc(ad.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedFunc(ad.integer); err.Error() != ad.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ad *audio) callSliceF(testedFunc func([]*types.MessageEntity) error, t *testing.T) {
	if !ad.isExpectedErr {
		if err := testedFunc(ad.array); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedFunc(ad.array); err.Error() != ad.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (ad *audio) startTest(part string, i int, t *testing.T) {
	switch f := ad.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, ad.name, ad.codeErr, ad.str, ad.isExpectedErr, i, t)
		ad.callStrF(f, t)
	case func(int) error:
		printTestLog(part, ad.name, ad.codeErr, ad.integer, ad.isExpectedErr, i, t)
		ad.callIntF(f, t)
	case func([]*types.MessageEntity) error:
		printTestLog(part, ad.name, ad.codeErr, ad.array, ad.isExpectedErr, i, t)
		ad.callSliceF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (at *audiotest) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var ad formatter.IAudio
	a, b := make([]UnitTest, at.until), make([]UnitTest, at.amount-at.until)
	for i, j := 0, 0; i < at.amount; i++ {
		if i < at.until {
			ad = msg.NewAudio()
			a[i] = at.buildF(*at, ad, i)
		} else {
			if j%2 == 0 {
				ad = msg.NewAudio()
			}
			b[j] = at.buildF(*at, ad, i)
			j++
		}
	}
	return a, b
}

func mainAudioLogic(msg *formatter.Message, at audiotest, t *testing.T) {
	audioholder, doubleaudio := at.createTestArrays(msg)
	check(audioholder, t)
	doubleCheck(doubleaudio, t)
}

func TestWriteAudioStorage(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeAudioStorage()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestWriteAudioTelegram(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeAudioTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestWriteAudioInternet(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeAudioInternet()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteCaption(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeCaption()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteEntitiesCaption(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeCaptionEntities()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteDuration(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeDuration()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteParseMode(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWritePerformer(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writePerformer()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteThumbnail(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeThumbnail()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteThumbnailID(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeThumbnailID()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}

func TestAudioWriteTitle(t *testing.T) {
	t.Parallel()
	at := new(audiotest)
	at.writeTitle()
	msg := formatter.CreateEmpltyMessage()
	mainAudioLogic(msg, *at, t)
}
