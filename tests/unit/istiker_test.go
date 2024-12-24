package unit

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
)

type stickerT struct {
	name          string
	str           string
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type stTestContainer struct {
	name          string
	inputStr      []string
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(st formatter.ISticker, s *stickerT, i int)
}

func (sttc *stTestContainer) defaultParamT(i int) *stickerT {
	return &stickerT{name: sttc.name, isExpectedErr: sttc.isExpectedErr[i], codeErr: sttc.codeErr[i]}
}
func (sttc *stTestContainer) sentName(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteSetName
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) writeSetName() {
	sttc.name = "(ISticker).WriteSetName()"
	sttc.inputStr = []string{"K:LASDKASKL:DAK:LSD", "", "Name", "there isn't a name"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.sentName
}

func (st *stickerT) startTest(part string, i int, t *testing.T) {
	switch f := st.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, st.name, st.codeErr, st.str, st.isExpectedErr, i)
		checkError(f(st.str), st.isExpectedErr, st.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (sttc *stTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var st formatter.ISticker
	a, b := make([]UnitTest, sttc.until), make([]UnitTest, sttc.amount-sttc.until)
	for i, j := 0, 0; i < sttc.amount; i++ {
		s := sttc.defaultParamT(i)
		if i < sttc.until {
			sttc.buildF(msg.NewSticker(), s, i)
			a[i] = s
		} else {
			if j%2 == 0 {
				st = msg.NewSticker()
			}
			sttc.buildF(st, s, i)
			b[j] = s
			j++
		}
	}
	return a, b
}

func (sttc *stTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	stcontainer, doublecontainer := sttc.createTestArrays(msg)
	check(stcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestStickerWriteSetName(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeSetName()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}
