package unit

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type stickerT struct {
	name          string
	str           string
	array         []string
	maskPos       *fmtogram.MaskPosition
	ent           []*fmtogram.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type stTestContainer struct {
	name          string
	inputStr      []string
	inputArr      [][]string
	maskPos       []*fmtogram.MaskPosition
	Entities      [][]*fmtogram.MessageEntity
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(st fmtogram.ISticker, s *stickerT, i int)
}

func (sttc *stTestContainer) defaultStikerT(i int) *stickerT {
	return &stickerT{name: sttc.name, isExpectedErr: sttc.isExpectedErr[i], codeErr: sttc.codeErr[i]}
}

func (sttc *stTestContainer) stickerStorage(st fmtogram.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteStickerStorage
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) stickerTelegram(st fmtogram.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteStickerTelegram
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) stickerInternet(st fmtogram.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteStickerInternet
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) associatedEmoji(st fmtogram.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteAssociatedEmoji
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) writeStickerStorage() {
	sttc.name = "(ISticker).WriteStickerStorage()"
	sttc.inputStr = []string{"../media/sticker.webp", "", "../media/Resume.pdf", "../media/sticker.webp", "../media/sticker.webp"}
	sttc.isExpectedErr = []bool{false, true, true, false, true}
	sttc.codeErr = []string{"", "20", "12", "", "10"}
	sttc.amount, sttc.until = 5, 3
	sttc.buildF = sttc.stickerStorage
}

func (sttc *stTestContainer) writeStickerTelegram() {
	sttc.name = "(ISticker).WriteStickerTelegram()"
	sttc.inputStr = []string{"ASLKJ:DKJALSDKLJA", "", "ASKDALJKSLJK%AS#$()", "ASKJDLAIOSUDAS*U()"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.stickerTelegram
}

func (sttc *stTestContainer) writeStickerInternet() {
	sttc.name = "(ISticker).WriteStickerInternet()"
	sttc.inputStr = []string{"ASLKJ:DKJALSDKLJA", "", "ASKDALJKSLJK%AS#$()", "ASKJDLAIOSUDAS*U()"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.stickerInternet
}

func (sttc *stTestContainer) writeAssociatedEmoji() {
	sttc.name = "(ISticker).WriteAssociatedEmoji()"
	sttc.inputStr = []string{"a sticker", "", "not a sticker", "st"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.associatedEmoji
}

func (st *stickerT) startTest(part string, i int, t *testing.T) {
	switch f := st.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, st.name, st.codeErr, st.str, st.isExpectedErr, i, t)
		checkError(f(st.str), st.isExpectedErr, st.codeErr, t)
	case func([]string) error:
		printTestLog(part, st.name, st.codeErr, st.array, st.isExpectedErr, i, t)
		checkError(f(st.array), st.isExpectedErr, st.codeErr, t)
	case func(*fmtogram.MaskPosition) error:
		printTestLog(part, st.name, st.codeErr, st.maskPos, st.isExpectedErr, i, t)
		checkError(f(st.maskPos), st.isExpectedErr, st.codeErr, t)
	case func([]*fmtogram.MessageEntity) error:
		printTestLog(part, st.name, st.codeErr, st.ent, st.isExpectedErr, i, t)
		checkError(f(st.ent), st.isExpectedErr, st.codeErr, t)
	case func() error:
		printTestLog(part, st.name, st.codeErr, true, st.isExpectedErr, i, t)
		checkError(f(), st.isExpectedErr, st.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (sttc *stTestContainer) createTestArrays(msg *fmtogram.Message) ([]UnitTest, []UnitTest) {
	var st fmtogram.ISticker
	a, b := make([]UnitTest, sttc.until), make([]UnitTest, sttc.amount-sttc.until)
	for i, j := 0, 0; i < sttc.amount; i++ {
		s := sttc.defaultStikerT(i)
		if i < sttc.until {
			sttc.buildF(fmtogram.NewSticker(), s, i)
			a[i] = s
		} else {
			if j%2 == 0 {
				st = fmtogram.NewSticker()
			}
			sttc.buildF(st, s, i)
			b[j] = s
			j++
		}
	}
	return a, b
}

func (sttc *stTestContainer) mainLogic(msg *fmtogram.Message, t *testing.T) {
	stcontainer, doublecontainer := sttc.createTestArrays(msg)
	check(stcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestStickerWriteStickerStorage(t *testing.T) {
	t.Parallel()
	sttc := new(stTestContainer)
	sttc.writeStickerStorage()
	msg := fmtogram.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteStickerTelegram(t *testing.T) {
	t.Parallel()
	sttc := new(stTestContainer)
	sttc.writeStickerTelegram()
	msg := fmtogram.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteStickerInternet(t *testing.T) {
	t.Parallel()
	sttc := new(stTestContainer)
	sttc.writeStickerInternet()
	msg := fmtogram.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteAssociatedEmoji(t *testing.T) {
	t.Parallel()
	sttc := new(stTestContainer)
	sttc.writeAssociatedEmoji()
	msg := fmtogram.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}
