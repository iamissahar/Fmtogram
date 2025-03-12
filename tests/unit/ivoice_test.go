package unit

import (
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
)

type voiceT struct {
	name          string
	str           string
	integer       int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type vcTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(vctc vcTestContainer, vc formatter.IVoice, i int) *voiceT
}

func putWriteVoiceStorage(vctc vcTestContainer, vc formatter.IVoice, i int) *voiceT {
	return &voiceT{vctc.name, vctc.inputStr[i], 0, vc.WriteVoiceStorage, vctc.isExpectedErr[i], vctc.codeErr[i]}
}

func putWriteVoiceTelegram(vctc vcTestContainer, vc formatter.IVoice, i int) *voiceT {
	return &voiceT{vctc.name, vctc.inputStr[i], 0, vc.WriteVoiceTelegram, vctc.isExpectedErr[i], vctc.codeErr[i]}
}

func putWriteVoiceInternet(vctc vcTestContainer, vc formatter.IVoice, i int) *voiceT {
	return &voiceT{vctc.name, vctc.inputStr[i], 0, vc.WriteVoiceInternet, vctc.isExpectedErr[i], vctc.codeErr[i]}
}

func putWriteVoiceDuration(vctc vcTestContainer, vc formatter.IVoice, i int) *voiceT {
	return &voiceT{vctc.name, "", vctc.inputInt[i], vc.WriteDuration, vctc.isExpectedErr[i], vctc.codeErr[i]}
}

func (vctc *vcTestContainer) writeVoiceStorage() {
	vctc.name = "(IVoice).WriteVoiceStorage"
	vctc.inputStr = []string{"../media/prichinatryski.mp4", "", "../media/tel-aviv.jpg", "../media/prichinatryski.mp4", "../media/prichinatryski.mp4"}
	vctc.isExpectedErr = []bool{false, true, true, false, true}
	vctc.codeErr = []string{"", "20", "12", "", "10"}
	vctc.amount, vctc.until = 5, 3
	vctc.buildF = putWriteVoiceStorage
}

func (vctc *vcTestContainer) writeVoiceTelegram() {
	vctc.name = "(IVoice).WriteVoiceTelegram"
	vctc.inputStr = []string{"LKASDKLASKLDLKASDLKASD", "", "A:KLSDK:LASL:DKA:SLDAS:", ":LASDK:LAKL:SDKLADA"}
	vctc.isExpectedErr = []bool{false, true, false, true}
	vctc.codeErr = []string{"", "20", "", "10"}
	vctc.amount, vctc.until = 4, 2
	vctc.buildF = putWriteVoiceTelegram
}

func (vctc *vcTestContainer) writeVoiceInternet() {
	vctc.name = "(IVoice).WriteVoiceInternet"
	vctc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	vctc.isExpectedErr = []bool{false, true, false, true}
	vctc.codeErr = []string{"", "20", "", "10"}
	vctc.amount, vctc.until = 4, 2
	vctc.buildF = putWriteVoiceInternet
}

func (vctc *vcTestContainer) writeDuration() {
	vctc.name = "(IVoice).WriteDuration"
	vctc.inputInt = []int{123155, 0, 1231566, 7}
	vctc.isExpectedErr = []bool{false, true, false, true}
	vctc.codeErr = []string{"", "20", "", "10"}
	vctc.amount, vctc.until = 4, 2
	vctc.buildF = putWriteVoiceDuration
}

func (vc *voiceT) callStrF(testedF func(string) error, t *testing.T) {
	if !vc.isExpectedErr {
		if err := testedF(vc.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vc.str); err.Error() != vc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vc *voiceT) callIntF(testedF func(int) error, t *testing.T) {
	if !vc.isExpectedErr {
		if err := testedF(vc.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vc.integer); err.Error() != vc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vc *voiceT) startTest(part string, i int, t *testing.T) {
	switch f := vc.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, vc.name, vc.codeErr, vc.str, vc.isExpectedErr, i, t)
		vc.callStrF(f, t)
	case func(int) error:
		printTestLog(part, vc.name, vc.codeErr, vc.integer, vc.isExpectedErr, i, t)
		vc.callIntF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (vctc *vcTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var vc formatter.IVoice
	a, b := make([]UnitTest, vctc.until), make([]UnitTest, vctc.amount-vctc.until)
	for i, j := 0, 0; i < vctc.amount; i++ {
		if i < vctc.until {
			vc = msg.NewVoice()
			a[i] = vctc.buildF(*vctc, vc, i)
		} else {
			if j%2 == 0 {
				vc = msg.NewVoice()
			}
			b[j] = vctc.buildF(*vctc, vc, i)
			j++
		}
	}
	return a, b
}

func mainVoiceLogic(msg *formatter.Message, antc vcTestContainer, t *testing.T) {
	animationcontainer, doublecontainer := antc.createTestArrays(msg)
	check(animationcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteVoiceStorage(t *testing.T) {
	t.Parallel()
	antc := new(vcTestContainer)
	antc.writeVoiceStorage()
	msg := formatter.CreateEmpltyMessage()
	mainVoiceLogic(msg, *antc, t)
}

func TestWriteVoiceTelegram(t *testing.T) {
	t.Parallel()
	antc := new(vcTestContainer)
	antc.writeVoiceTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainVoiceLogic(msg, *antc, t)
}

func TestWriteVoiceInternet(t *testing.T) {
	t.Parallel()
	antc := new(vcTestContainer)
	antc.writeVoiceInternet()
	msg := formatter.CreateEmpltyMessage()
	mainVoiceLogic(msg, *antc, t)
}

func TestWriteVoiceDuration(t *testing.T) {
	t.Parallel()
	antc := new(vcTestContainer)
	antc.writeDuration()
	msg := formatter.CreateEmpltyMessage()
	mainVoiceLogic(msg, *antc, t)
}
