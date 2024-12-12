package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
)

type videonoteT struct {
	name          string
	str           string
	integer       int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type vdnTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(vdtc vdnTestContainer, vd formatter.IVideoNote, i int) *videonoteT
}

func putWriteVideoNoteStorage(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, vdntc.inputStr[i], 0, vdn.WriteVideoNoteStorage, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteTelegram(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, vdntc.inputStr[i], 0, vdn.WriteVideoNoteTelegram, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteInternet(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, vdntc.inputStr[i], 0, vdn.WriteVideoNoteInternet, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteDuration(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, "", vdntc.inputInt[i], vdn.WriteDuration, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteLength(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, "", vdntc.inputInt[i], vdn.WriteLength, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteThumbnailStorage(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, vdntc.inputStr[i], 0, vdn.WriteThumbnailStorage, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteThumbnailTelegram(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, vdntc.inputStr[i], 0, vdn.WriteThumbnailTelegram, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func putWriteVideoNoteThumbnailInternet(vdntc vdnTestContainer, vdn formatter.IVideoNote, i int) *videonoteT {
	return &videonoteT{vdntc.name, vdntc.inputStr[i], 0, vdn.WriteThumbnailInternet, vdntc.isExpectedErr[i], vdntc.codeErr[i]}
}

func (vdntc *vdnTestContainer) writeVideoNoteStorage() {
	vdntc.name = "(IVideoNote).WriteVideoNoteStorage"
	vdntc.inputStr = []string{"../media_test/prichinatryski.mp4", "", "../media_test/tel-aviv.jpg", "../media_test/prichinatryski.mp4", "../media_test/prichinatryski.mp4"}
	vdntc.isExpectedErr = []bool{false, true, true, false, true}
	vdntc.codeErr = []string{"", "20", "12", "", "10"}
	vdntc.amount, vdntc.until = 5, 3
	vdntc.buildF = putWriteVideoNoteStorage
}

func (vdntc *vdnTestContainer) writeVideoNoteTelegram() {
	vdntc.name = "(IVideoNote).WriteVideoNoteTelegram"
	vdntc.inputStr = []string{"LKASDKLASKLDLKASDLKASD", "", "A:KLSDK:LASL:DKA:SLDAS:", ":LASDK:LAKL:SDKLADA"}
	vdntc.isExpectedErr = []bool{false, true, false, true}
	vdntc.codeErr = []string{"", "20", "", "10"}
	vdntc.amount, vdntc.until = 4, 2
	vdntc.buildF = putWriteVideoNoteTelegram
}

func (vdntc *vdnTestContainer) writeVideoNoteInternet() {
	vdntc.name = "(IVideoNote).WriteVideoNoteInternet"
	vdntc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", "https://youtube.com"}
	vdntc.isExpectedErr = []bool{false, true, false, true}
	vdntc.codeErr = []string{"", "20", "", "10"}
	vdntc.amount, vdntc.until = 4, 2
	vdntc.buildF = putWriteVideoNoteInternet
}

func (vdntc *vdnTestContainer) writeDuration() {
	vdntc.name = "(IVideoNote).WriteDuration"
	vdntc.inputInt = []int{238094029, 0, 68986, 33}
	vdntc.isExpectedErr = []bool{false, true, false, true}
	vdntc.codeErr = []string{"", "20", "", "10"}
	vdntc.amount, vdntc.until = 4, 2
	vdntc.buildF = putWriteVideoNoteDuration
}

func (vdntc *vdnTestContainer) writeLength() {
	vdntc.name = "(IVideoNote).WriteLength"
	vdntc.inputInt = []int{238094029, 0, 68986, 33}
	vdntc.isExpectedErr = []bool{false, true, false, true}
	vdntc.codeErr = []string{"", "20", "", "10"}
	vdntc.amount, vdntc.until = 4, 2
	vdntc.buildF = putWriteVideoNoteLength
}

func (vdntc *vdnTestContainer) writeThumbnailStorage() {
	vdntc.name = "(IVideoNote).WriteThumbnailStorage"
	vdntc.inputStr = []string{"../media_test/tel-aviv.jpg", "", "../media_test/sound.mp3", "../media_test/tel-aviv.jpg", "../media_test/tel-aviv.jpg"}
	vdntc.isExpectedErr = []bool{false, true, true, false, true}
	vdntc.codeErr = []string{"", "20", "12", "", "10"}
	vdntc.amount, vdntc.until = 5, 3
	vdntc.buildF = putWriteVideoNoteThumbnailStorage
}

func (vdntc *vdnTestContainer) writeThumbnailTelegram() {
	vdntc.name = "(IVideoNote).WriteThumbnailTelegram"
	vdntc.inputStr = []string{"ASL:KDKAOL:SLK:@#$!:L", "", "A:LSKDKL:ASK:DLASKL:DASD", ")()*()#I@!K:OLAS:KLDAS:L"}
	vdntc.isExpectedErr = []bool{false, true, false, true}
	vdntc.codeErr = []string{"", "20", "", "10"}
	vdntc.amount, vdntc.until = 4, 2
	vdntc.buildF = putWriteVideoNoteThumbnailTelegram
}

func (vdntc *vdnTestContainer) writeThumbnailInternet() {
	vdntc.name = "(IVideoNote).WriteThumbnailInternet"
	vdntc.inputStr = []string{"https://youtube.com", "", "https://youtube.com", ")()*()#I@!K:OLAS:KLDAS:L"}
	vdntc.isExpectedErr = []bool{false, true, false, true}
	vdntc.codeErr = []string{"", "20", "", "10"}
	vdntc.amount, vdntc.until = 4, 2
	vdntc.buildF = putWriteVideoNoteThumbnailInternet
}

func (vdn *videonoteT) callStrF(testedF func(string) error, t *testing.T) {
	if !vdn.isExpectedErr {
		if err := testedF(vdn.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vdn.str); err.Error() != vdn.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vdn *videonoteT) callIntF(testedF func(int) error, t *testing.T) {
	if !vdn.isExpectedErr {
		if err := testedF(vdn.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(vdn.integer); err.Error() != vdn.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (vdn *videonoteT) startTest(part string, i int, t *testing.T) {
	switch f := vdn.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, vdn.name, vdn.codeErr, vdn.str, vdn.isExpectedErr, i)
		vdn.callStrF(f, t)
	case func(int) error:
		printTestLog(part, vdn.name, vdn.codeErr, vdn.integer, vdn.isExpectedErr, i)
		vdn.callIntF(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (vdntc *vdnTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var vdn formatter.IVideoNote
	a, b := make([]UnitTest, vdntc.until), make([]UnitTest, vdntc.amount-vdntc.until)
	for i, j := 0, 0; i < vdntc.amount; i++ {
		if i < vdntc.until {
			vdn = msg.NewVideoNote()
			a[i] = vdntc.buildF(*vdntc, vdn, i)
		} else {
			if j%2 == 0 {
				vdn = msg.NewVideoNote()
			}
			b[j] = vdntc.buildF(*vdntc, vdn, i)
			j++
		}
	}
	return a, b
}

func mainVideoNoteLogic(msg *formatter.Message, vdntc vdnTestContainer, t *testing.T) {
	videonotecontainer, doublecontainer := vdntc.createTestArrays(msg)
	check(videonotecontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteVideoNoteStorage(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeVideoNoteStorage()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteTelegram(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeVideoNoteTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteInternet(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeVideoNoteInternet()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteDuration(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeDuration()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteLength(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeLength()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteThumbnailStorage(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeThumbnailStorage()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteThumbnailTelegram(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeThumbnailTelegram()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}

func TestWriteVideoNoteThumbnailInternet(t *testing.T) {
	vdntc := new(vdnTestContainer)
	vdntc.writeThumbnailInternet()
	msg := formatter.CreateEmpltyMessage()
	mainVideoNoteLogic(msg, *vdntc, t)
}
