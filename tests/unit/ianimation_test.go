package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
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
	buildF        func(vdtc anTestContainer, vd formatter.IAnimation, i int) *animationT
}

func putWriteAnimationStorage(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteAnimationStorage, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationTelegram(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteAnimationTelegram, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationInternet(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteAnimationInternet, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationDuration(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", antc.inputInt[i], an.WriteDuration, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationWidth(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", antc.inputInt[i], an.WriteWidth, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationHeight(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", antc.inputInt[i], an.WriteHeight, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationThumbnailStorage(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteThumbnailStorage, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationThumbnailTelegram(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteThumbnailTelegram, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationThumbnailInternet(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, antc.inputStr[i], 0, an.WriteThumbnailInternet, antc.isExpectedErr[i], antc.codeErr[i]}
}

func putWriteAnimationHasSpoiler(antc anTestContainer, an formatter.IAnimation, i int) *animationT {
	return &animationT{antc.name, "", 0, an.WriteHasSpoiler, antc.isExpectedErr[i], antc.codeErr[i]}
}

func (antc *anTestContainer) writeAnimationStorage() {
	antc.name = "(IAnimation).WriteAnimationStorage"
	antc.inputStr = []string{"../media_test/prichinatryski.mp4", "", "../media_test/tel-aviv.jpg", "../media_test/prichinatryski.mp4", "../media_test/prichinatryski.mp4"}
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

func (antc *anTestContainer) writeThumbnailStorage() {
	antc.name = "(IAnimation).WriteThumbnailStorage"
	antc.inputStr = []string{"../media_test/tel-aviv.jpg", "", "../media_test/sound.mp3", "../media_test/tel-aviv.jpg", "../media_test/tel-aviv.jpg"}
	antc.isExpectedErr = []bool{false, true, true, false, true}
	antc.codeErr = []string{"", "20", "12", "", "10"}
	antc.amount, antc.until = 5, 3
	antc.buildF = putWriteAnimationThumbnailStorage
}

func (antc *anTestContainer) writeThumbnailTelegram() {
	antc.name = "(IAnimation).WriteThumbnailTelegram"
	antc.inputStr = []string{"ASL:KDKAOL:SLK:@#$!:L", "", "A:LSKDKL:ASK:DLASKL:DASD", ")()*()#I@!K:OLAS:KLDAS:L"}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationThumbnailTelegram
}

func (antc *anTestContainer) writeThumbnailInternet() {
	antc.name = "(IAnimation).WriteThumbnailInternet"
	antc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", ")()*()#I@!K:OLAS:KLDAS:L"}
	antc.isExpectedErr = []bool{false, true, false, true}
	antc.codeErr = []string{"", "20", "", "10"}
	antc.amount, antc.until = 4, 2
	antc.buildF = putWriteAnimationThumbnailInternet
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
		printTestLog(part, an.name, an.codeErr, an.str, an.isExpectedErr, i)
		an.callStrF(f, t)
	case func(int) error:
		printTestLog(part, an.name, an.codeErr, an.integer, an.isExpectedErr, i)
		an.callIntF(f, t)
	case func() error:
		printTestLog(part, an.name, an.codeErr, true, an.isExpectedErr, i)
		an.callBoolF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (antc *anTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var an formatter.IAnimation
	a, b := make([]UnitTest, antc.until), make([]UnitTest, antc.amount-antc.until)
	for i, j := 0, 0; i < antc.amount; i++ {
		if i < antc.until {
			an = msg.NewAnimation()
			a[i] = antc.buildF(*antc, an, i)
		} else {
			if j%2 == 0 {
				an = msg.NewAnimation()
			}
			b[j] = antc.buildF(*antc, an, i)
			j++
		}
	}
	return a, b
}

func mainAnimationLogic(msg *formatter.Message, antc anTestContainer, t *testing.T) {
	animationcontainer, doublecontainer := antc.createTestArrays(msg)
	check(animationcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteAnimationStorage(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeAnimationStorage()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationTelegram(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeAnimationTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationInternet(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeAnimationInternet()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationDuration(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeDuration()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationWidth(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeWidth()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationHeight(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeHeight()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationThumbnailStorage(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeThumbnailStorage()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationThumbnailTelegram(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeThumbnailTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationThumbnailInternet(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeThumbnailInternet()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}

func TestWriteAnimationHasSpoiler(t *testing.T) {
	antc := new(anTestContainer)
	antc.writeHasSpoiler()
	msg := formatter.CreateEmpltyMessage()
	mainAnimationLogic(msg, *antc, t)
}
