package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func videonoteStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vdn.GetResponse())
	t.Log(ch.GetResponse())
}

func videonoteStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.VideoNote)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vdn.GetResponse())
	t.Log(ch.GetResponse())
}

func videonoteStorageUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(11); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vdn.GetResponse())
}

func videonoteStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(11); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.VideoNote)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vdn.GetResponse())
}

func videonoteStorageUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(11); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createReplyKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.VideoNote)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vdn.GetResponse())
}

func videonoteTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("DQACAgIAAxkDAAIDRmcpRSgQWnun_irRK7THUS51PuBEAAJeXAAC43ZRSUQtjaW6gJjyNgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vdn.GetResponse())
	t.Log(ch.GetResponse())
}

func videonoteTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("DQACAgIAAxkDAAIDRmcpRSgQWnun_irRK7THUS51PuBEAAJeXAAC43ZRSUQtjaW6gJjyNgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.VideoNote)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vdn.GetResponse())
	t.Log(ch.GetResponse())
}

func videonoteTelegramUnReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("DQACAgIAAxkDAAIDRmcpRSgQWnun_irRK7THUS51PuBEAAJeXAAC43ZRSUQtjaW6gJjyNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(11); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vdn.GetResponse())
}

func videonoteTelegramUnReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("DQACAgIAAxkDAAIDRmcpRSgQWnun_irRK7THUS51PuBEAAJeXAAC43ZRSUQtjaW6gJjyNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(11); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.VideoNote)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vdn.GetResponse())
}

func videonoteTelegramUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("DQACAgIAAxkDAAIDRmcpRSgQWnun_irRK7THUS51PuBEAAJeXAAC43ZRSUQtjaW6gJjyNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(11); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createReplyKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.VideoNote)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vdn.GetResponse())
}

func videonoteFunctional(t *testing.T) {
	t.Parallel()
	t.Run("VideoNoteStorageReqWithoutName", videonoteStorageReqWithoutName)
	t.Run("VideoNoteStorageReqWithName", videonoteStorageReqWithName)
	t.Run("VideoNoteStorageUnreqWithoutName", videonoteStorageUnreqWithoutName)
	t.Run("VideoNoteStorageUnreqWithName", videonoteStorageUnreqWithName)
	t.Run("VideoNoteStorageUnreqReplyKB", videonoteStorageUnreqReplyKB)
	t.Run("VideoNoteTelegramReqWithoutName", videonoteTelegramReqWithoutName)
	t.Run("VideoNoteTelegramReqWithName", videonoteTelegramReqWithName)
	t.Run("VideoNoteTelegramUnReqWithoutName", videonoteTelegramUnReqWithoutName)
	t.Run("VideoNoteTelegramUnReqWithName", videonoteTelegramUnReqWithName)
	t.Run("VideoNoteTelegramUnreqReplyKB", videonoteTelegramUnreqReplyKB)
}

func videonoteStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
}

func videonoteStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(444); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func videonoteSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteVideoNoteStorage("media_test/black.mp4"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage("media_test/tel-aviv.jpg"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func videonoteSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteStOK)
	t.Run("OKs", videonoteStOKs)
	t.Run("Code10", videonoteSt10)
	t.Run("Code12", videonoteSt12)
	t.Run("Code20", videonoteSt20)
}

func videonoteTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("aklsdlkasdkla"); err != nil {
		t.Fatal(err)
	}
}

func videonoteTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("aklsdlkasdkla"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(444); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func videonoteTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("aklsdlkasdkla"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteVideoNoteTelegram("aklsdlkasdkla"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteTgOK)
	t.Run("OKs", videonoteTgOKs)
	t.Run("Code10", videonoteTg10)
	t.Run("Code20", videonoteTg20)
}

func videonoteIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func videonoteIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(444); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func videonoteInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteVideoNoteInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteIntOK)
	t.Run("OKs", videonoteIntOKs)
	t.Run("Code10", videonoteInt10)
	t.Run("Code20", videonoteInt20)
}

func videonoteDurOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
}

func videonoteDur10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteDuration(123); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteDur20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteDuration(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteDuration(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteDurOK)
	t.Run("Code10", videonoteDur10)
	t.Run("Code20", videonoteDur20)
}

func videonoteLenOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteLength(123); err != nil {
		t.Fatal(err)
	}
}

func videonoteLen10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteLength(123); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteLength(123); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteLen20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteLength(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteLength(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteLenOK)
	t.Run("Code10", videonoteLen10)
	t.Run("Code20", videonoteLen20)
}

func videonoteThumbStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func videonoteThumbSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}
func videonoteThumbSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailStorage("media_test/musk.mp4"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func videonoteThumbSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteThumbnailStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteThumbStOK)
	t.Run("Code10", videonoteThumbSt10)
	t.Run("Code12", videonoteThumbSt12)
	t.Run("Code20", videonoteThumbSt20)
}

func videonoteThumbTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailTelegram("asdasdas"); err != nil {
		t.Fatal(err)
	}
}

func videonoteThumbTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailTelegram("asdasdas"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteThumbnailTelegram("asdasdas"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteThumbTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteThumbnailTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteThumbTgOK)
	t.Run("Code10", videonoteThumbTg10)
	t.Run("Code20", videonoteThumbTg20)
}

func videonoteThumbIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func videonoteThumbInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vdn.WriteThumbnailInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videonoteThumbInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := vdn.WriteThumbnailInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videonoteThumbnailInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", videonoteThumbIntOK)
	t.Run("Code10", videonoteThumbInt10)
	t.Run("Code20", videonoteThumbInt20)
}

func videonoteCode21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vdn := msg.NewVideoNote()
	if err := msg.AddVideoNote(vdn); err.Error() != code21 {
		t.Fatal(err)
	}
}

func videonoteUnit(t *testing.T) {
	t.Parallel()
	t.Run("WriteVideoNoteStorage", videonoteStorage)
	t.Run("WriteVideoNoteTelegram", videonoteTelegram)
	t.Run("WriteVideoNoteInternet", videonoteInternet)
	t.Run("WriteDuration", videonoteDuration)
	t.Run("WriteLength", videonoteLength)
	t.Run("WriteThumbnailStorage", videonoteThumbnailStorage)
	t.Run("WriteThumbnailTelegram", videonoteThumbnailTelegram)
	t.Run("WriteThumbnailInternet", videonoteThumbnailInternet)
	t.Run("Code21", videonoteCode21)
}

func TestSendVideoNote(t *testing.T) {
	t.Run("Functional", videonoteFunctional)
	t.Run("Unit", videonoteUnit)
}
