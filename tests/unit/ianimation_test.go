package unit

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type animationT struct {
	name          string
	str           string
	integer       int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type anTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(vdtc anTestContainer, vd fmtogram.IAnimation, i int) *animationT
}

func putWriteAnimationStorage(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteAnimationStorage, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationTelegram(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteAnimationTelegram, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationInternet(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteAnimationInternet, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationDuration(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", antc.inputInt[i], an.WriteDuration, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationWidth(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", antc.inputInt[i], an.WriteWidth, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationHeight(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", antc.inputInt[i], an.WriteHeight, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationThumbnail(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteThumbnail, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationThumbnailID(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteThumbnailID, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationHasSpoiler(antc anTestContainer, an fmtogram.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", 0, an.WriteHasSpoiler, antc.isExpectedErr[i], antc.codeErr[i]}
}

func (antc *anTestContainer) writeAnimationStorage() {
	antc.name = "(IAnimation).WriteAnimationStorage"
	antc.inputStr = []string{"../media/prichinatryski.mp4", "", "../media/tel-aviv.jpg", "../media/prichinatryski.mp4", "../media/prichinatryski.mp4"}
	antc.isExpectedErr = []bool{false, true, true, false, true}
	antc.codeErr = []string{"", "20", "12", "", "10"}
	antc.amount, antc.until = 5, 3
	antc.buildF = putWriteAnimationStorage
}

func (antc *anTestContainer) writeAnimationTelegram() {
	antc.name = "(IAnimation).WriteAnimationTelegram"
	antc.inputStr = []string{"LKASDKLASKLDLKASDLKASD", "", "A:KLSDK:LASL:DKA:SLDAS:", ":LASDK:LAKL:SDKLADA"}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationTelegram
}

func (antc *anTestContainer) writeAnimationInternet() {
	antc.name = "(IAnimation).WriteAnimationInternet"
	antc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationInternet
}

func (antc *anTestContainer) writeDuration() {
	antc.name = "(IAnimation).WriteDuration"
	antc.inputInt = []int{123, 0, 12351, 56875}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationDuration
}

func (antc *anTestContainer) writeWidth() {
	antc.name = "(IAnimation).WriteWidth"
	antc.inputInt = []int{123, 0, 12351, 56875}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationWidth
}

func (antc *anTestContainer) writeHeight() {
	antc.name = "(IAnimation).WriteHeight"
	antc.inputInt = []int{123, 0, 12351, 56875}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationHeight
}

func (antc *anTestContainer) writeThumbnail() {
	antc.name = "(IAnimation).WriteThumbnail"
	antc.inputStr = []string{"../media/tel-aviv.jpg", "", "../media/sound.mp3", "../media/tel-aviv.jpg", "../media/tel-aviv.jpg"}
	antc.isExpectedErr = []bool{false, true, true, false, true}
	antc.codeErr = []string{"", "20", "12", "", "10"}
	antc.amount, antc.until = 5, 3
	antc.buildF = putWriteAnimationThumbnail
}

func (antc *anTestContainer) writeThumbnailID() {
	antc.name = "(IAnimation).WriteThumbnail"
	antc.inputStr = []string{"ASL:KDKAOL:SLK:@#$!:L", "", "A:LSKDKL:ASK:DLASKL:DASD", ")()*()#I@!K:OLAS:KLDAS:L"}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationThumbnailID
}

func (antc *anTestContainer) writeHasSpoiler() {
	antc.name = "(IAnimation).WriteHasSpoiler"
	antc.isExpectedErr = []bool{false, false, true}
	antc.codeErr = []string{"", "", "10"}
	antc.amount, antc.until = 3, 1
	antc.buildF = putWriteAnimationHasSpoiler
}

func (an *animationT) callStrF(testedF func(string) error, t *testing.T) {
	if !an.isExpectedErr {
		if err := testedF(an.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(an.str); err.Error() != an.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (an *animationT) callIntF(testedF func(int) error, t *testing.T) {
	if !an.isExpectedErr {
		if err := testedF(an.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(an.integer); err.Error() != an.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (an *animationT) callBoolF(testedF func() error, t *testing.T) {
	if !an.isExpectedErr {
		if err := testedF(); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(); err.Error() != an.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (an *animationT) startTest(part string, i int, t *testing.T) {
	switch f := an.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, an.name, an.codeErr, an.str, an.isExpectedErr, i, t)
		an.callStrF(f, t)
	case func(int) error:
		printTestLog(part, an.name, an.codeErr, an.integer, an.isExpectedErr, i, t)
		an.callIntF(f, t)
	case func() error:
		printTestLog(part, an.name, an.codeErr, true, an.isExpectedErr, i, t)
		an.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (antc *anTestContainer) createTestArrays(msg *fmtogram.Message) ([]UnitTest, []UnitTest) {
	var an fmtogram.IAnimation
	a, b := make([]UnitTest, antc.until), make([]UnitTest, antc.amount-antc.until)
	for i, j := 0, 0; i < antc.amount; i++ {
		if i < antc.until {
			an = fmtogram.NewAnimation()
			a[i] = antc.buildF(*antc, an, i)
		} else {
			if j%2 == 0 {
				an = fmtogram.NewAnimation()
			}
			b[j] = antc.buildF(*antc, an, i)
			j++
		}
	}
	return a, b
}

func mainAnimationLogic(msg *fmtogram.Message, antc anTestContainer, t *testing.T) {
	animationcontainer, doublecontainer := antc.createTestArrays(msg)
	check(animationcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteAnimationStorage(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeAnimationStorage()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationTelegram(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeAnimationTelegram()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationInternet(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeAnimationInternet()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationDuration(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeDuration()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationWidth(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeWidth()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationHeight(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeHeight()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationThumbnailStorage(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeThumbnail()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationThumbnailTelegram(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeThumbnailID()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationHasSpoiler(t *testing.T) {
	t.Parallel()
	antc := new(anTestContainer)
	antc.writeHasSpoiler()
	msg := fmtogram.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}
