package unit

import (
	"strings"
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type stickerT struct {
	name          string
	str           string
	array         []string
	maskPos       *types.MaskPosition
	ent           []*types.MessageEntity
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type stTestContainer struct {
	name          string
	inputStr      []string
	inputArr      [][]string
	maskPos       []*types.MaskPosition
	Entities      [][]*types.MessageEntity
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(st formatter.ISticker, s *stickerT, i int)
}

func (sttc *stTestContainer) defaultStikerT(i int) *stickerT {
	return &stickerT{name: sttc.name, isExpectedErr: sttc.isExpectedErr[i], codeErr: sttc.codeErr[i]}
}

func (sttc *stTestContainer) stickerStorage(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteStickerStorage
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) stickerTelegram(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteStickerTelegram
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) stickerInternet(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteStickerInternet
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) associatedEmoji(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteAssociatedEmoji
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) associatedEmojies(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteAssociatedEmojies
	s.array = sttc.inputArr[i]
}

func (sttc *stTestContainer) emojiID(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteEmojiID
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) emojiIDs(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteEmojiIDs
	s.array = sttc.inputArr[i]
}

func (sttc *stTestContainer) stickerFormat(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteFormat
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) position(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WritePosition
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) oldSticker(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteOldSticker
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) keywords(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteKeywords
	s.array = sttc.inputArr[i]
}

func (sttc *stTestContainer) maskPosition(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteMaskPosition
	s.maskPos = sttc.maskPos[i]
}

func (sttc *stTestContainer) thumbnailStorage(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteThumbnailStorage
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) thumbnailTelegram(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteThumbnailTelegram
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) thumbnailInternet(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteThumbnailInternet
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) thumbnailFormat(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteThumbnailFormat
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) giftID(st formatter.ISticker, s *stickerT, i int) {
	s.testedFunc = st.WriteGiftID
	s.str = sttc.inputStr[i]
}

func (sttc *stTestContainer) writeStickerStorage() {
	sttc.name = "(ISticker).WriteStickerStorage()"
	sttc.inputStr = []string{"../media_test/sticker.webp", "", "../media_test/Resume.pdf", "../media_test/sticker.webp", "../media_test/sticker.webp"}
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

func (sttc *stTestContainer) writeAssociatedEmojies() {
	sttc.name = "(ISticker).WriteAssociatedEmojies()"
	sttc.inputArr = [][]string{{"a sticker"}, nil, {"something like a sticker", "", ""}, {"not a sticker"}, {"st"}}
	sttc.isExpectedErr = []bool{false, true, true, false, true}
	sttc.codeErr = []string{"", "20", "5", "", "10"}
	sttc.amount, sttc.until = 5, 3
	sttc.buildF = sttc.associatedEmojies
}

func (sttc *stTestContainer) writeEmojiID() {
	sttc.name = "(ISticker).WriteEmojiID()"
	sttc.inputStr = []string{"a sticker id", "", "not a sticker id", "stID"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.emojiID
}

func (sttc *stTestContainer) writeEmojiIDs() {
	sttc.name = "(ISticker).WriteEmojiIDs()"
	sttc.inputArr = [][]string{{"a sticker id"}, nil, {"something like a sticker id", "", ""}, {"not a sticker id"}, {"stID"}}
	sttc.isExpectedErr = []bool{false, true, true, false, true}
	sttc.codeErr = []string{"", "20", "5", "", "10"}
	sttc.amount, sttc.until = 5, 3
	sttc.buildF = sttc.emojiIDs
}

func (sttc *stTestContainer) writeFormat() {
	sttc.name = "(ISticker).WriteFormat()"
	sttc.inputStr = []string{"static", "", "animated", "video", "not a format of a stiker", "video", "video"}
	sttc.isExpectedErr = []bool{false, true, false, false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "", "20", "", "10"}
	sttc.amount, sttc.until = 7, 5
	sttc.buildF = sttc.stickerFormat
}

func (sttc *stTestContainer) writePosition() {
	sttc.name = "(ISticker).WritePosition()"
	sttc.inputStr = []string{"0123124", "0", "", "8384281", "as;klddl;asd;lg", "0123124", "0123124"}
	sttc.isExpectedErr = []bool{false, false, true, true, true, false, true}
	sttc.codeErr = []string{"", "", "20", "20", "20", "", "10"}
	sttc.amount, sttc.until = 3, 1
	sttc.buildF = sttc.position
}

func (sttc *stTestContainer) writeOldSticker() {
	sttc.name = "(ISticker).WriteOldSticker()"
	sttc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.oldSticker
}

func (sttc *stTestContainer) writeKeywords() {
	sttc.name = "(ISticker).WriteKeywords()"
	sttc.inputArr = [][]string{{"asdasdasdad"}, {""},
		nil,
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		{strings.Repeat(":", 64)},
		{strings.Repeat(":", 65)},
		{"ASKLJDAKSJLFJKLASJKLD"}, {"AS:>DLKAWIO$*(AS)(ED()ASD)*()"}}
	sttc.isExpectedErr = []bool{false, false,
		false,
		false,
		true,
		false,
		true,
		false, true}
	sttc.codeErr = []string{"", "",
		"",
		"",
		"20",
		"",
		"5",
		"", "10"}
	sttc.amount, sttc.until = 9, 7
	sttc.buildF = sttc.keywords
}

func (sttc *stTestContainer) writeMaskPosition() {
	sttc.name = "(ISticker).WriteMaskPosition()"
	sttc.maskPos = []*types.MaskPosition{{Point: "1231"}, nil, {Point: "1231"}, {Point: "1231"}}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.maskPosition
}

func (sttc *stTestContainer) writeThumbnailStorage() {
	sttc.name = "(ISticker).WriteThumbnailStorage()"
	sttc.inputStr = []string{"../media_test/cookie.png",
		"../media_test/sticker.webp",
		"../media_test/prichinatryski.webm", "",
		"../media_test/tel-aviv.jpg",
		"../media_test/cookie.png", "../media_test/cookie.png"}
	sttc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	sttc.codeErr = []string{"", "", "", "20", "12", "", "10"}
	sttc.amount, sttc.until = 7, 5
	sttc.buildF = sttc.thumbnailStorage
}

func (sttc *stTestContainer) writeThumbnailTelegram() {
	sttc.name = "(ISticker).WriteThumbnailTelegram()"
	sttc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.thumbnailTelegram
}

func (sttc *stTestContainer) writeThumbnailInternet() {
	sttc.name = "(ISticker).WriteThumbnailInternet()"
	sttc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.thumbnailInternet
}

func (sttc *stTestContainer) writeThumbnailFormat() {
	sttc.name = "(ISticker).WriteThumbnailFormat()"
	sttc.inputStr = []string{"static", "animated", "video", "", "ASKLJDAKSJLFJKLASJKLD", "animated", "animated"}
	sttc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	sttc.codeErr = []string{"", "", "", "20", "20", "", "10"}
	sttc.amount, sttc.until = 7, 5
	sttc.buildF = sttc.thumbnailFormat
}

func (sttc *stTestContainer) writeGiftID() {
	sttc.name = "(ISticker).WriteGiftID()"
	sttc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	sttc.isExpectedErr = []bool{false, true, false, true}
	sttc.codeErr = []string{"", "20", "", "10"}
	sttc.amount, sttc.until = 4, 2
	sttc.buildF = sttc.giftID
}

func (st *stickerT) startTest(part string, i int, t *testing.T) {
	switch f := st.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, st.name, st.codeErr, st.str, st.isExpectedErr, i)
		checkError(f(st.str), st.isExpectedErr, st.codeErr, t)
	case func([]string) error:
		printTestLog(part, st.name, st.codeErr, st.array, st.isExpectedErr, i)
		checkError(f(st.array), st.isExpectedErr, st.codeErr, t)
	case func(*types.MaskPosition) error:
		printTestLog(part, st.name, st.codeErr, st.maskPos, st.isExpectedErr, i)
		checkError(f(st.maskPos), st.isExpectedErr, st.codeErr, t)
	case func([]*types.MessageEntity) error:
		printTestLog(part, st.name, st.codeErr, st.ent, st.isExpectedErr, i)
		checkError(f(st.ent), st.isExpectedErr, st.codeErr, t)
	case func() error:
		printTestLog(part, st.name, st.codeErr, true, st.isExpectedErr, i)
		checkError(f(), st.isExpectedErr, st.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (sttc *stTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var st formatter.ISticker
	a, b := make([]UnitTest, sttc.until), make([]UnitTest, sttc.amount-sttc.until)
	for i, j := 0, 0; i < sttc.amount; i++ {
		s := sttc.defaultStikerT(i)
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

func TestStickerWriteStickerStorage(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeStickerStorage()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteStickerTelegram(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeStickerTelegram()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteStickerInternet(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeStickerInternet()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteAssociatedEmoji(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeAssociatedEmoji()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteAssociatedEmojies(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeAssociatedEmojies()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteEmojiID(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeEmojiID()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteEmojiIDs(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeEmojiIDs()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWroteFormat(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeFormat()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStcikerWritePosition(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writePosition()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteOldSticker(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeOldSticker()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteKeywords(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeKeywords()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteMaskPosition(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeMaskPosition()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStikerWriteThumbnailStorage(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeThumbnailStorage()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStikerWriteThumbnailTelegra(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeThumbnailTelegram()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStikerWriteThumbnailInternet(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeThumbnailInternet()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStikerWriteGiftID(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeThumbnailFormat()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}

func TestStickerWriteGiftID(t *testing.T) {
	sttc := new(stTestContainer)
	sttc.writeGiftID()
	msg := formatter.CreateEmpltyMessage()
	sttc.mainLogic(msg, t)
}
