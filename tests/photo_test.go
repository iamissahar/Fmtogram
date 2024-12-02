package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func photoStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ph.GetResponse())
	t.Log(ch.GetResponse())
}

func photoStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}

	msg.AddMethod(methods.Photo)
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ph.GetResponse())
	t.Log(ch.GetResponse())
}

func photoStorageUnreqWithoutName(t *testing.T) {
	var err error
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err = ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err = ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewMessage()
	if err = inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err = inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err = inf.WriteString("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}

	_, err = msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ph.GetResponse())
}

func photoStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewMessage()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
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
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ph.GetResponse())
}

func photoStorageUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewMessage()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
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
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ph.GetResponse())
}

func photoTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ph.GetResponse())
	t.Log(ch.GetResponse())
}

func photoTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		panic(err)
	}

	t.Log(ph.GetResponse())
	t.Log(ch.GetResponse())
}

func photoTelegramUnReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	inf := msg.NewMessage()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
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
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ph.GetResponse())
}

func photoTelegramUnReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaption("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewMessage()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ph.GetResponse())
}

func photoTelegramUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaption("<b>SHALOM!</b> Y'ALL!"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewMessage()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}

	createReplyKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ph.GetResponse())
}

func photoStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err, err != nil)
	}
}

func photoStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func photoStCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoStCode12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/black.mp4"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func photoStCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func photoCode21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := msg.AddPhoto(ph); err.Error() != code21 {
		t.Fatal(err)
	}
}

func photoStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoStOK)
	t.Run("ALO", photoStOKs)
	t.Run("Code10", photoStCode10)
	t.Run("Code12", photoStCode12)
	t.Run("Code20", photoStCode20)
}

func photoTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("ASKLDAKLSDKLASLKD123asl"); err != nil {
		t.Fatal(err)
	}
}

func photoTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("ASKLDAKLSDKLASLKD123asl"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func photoTgCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("ASKLDAKLSDKLASLKD123aslg"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WritePhotoTelegram("ASKLDAKLSDKLASLKD123aslpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoTgCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func photoTG(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoTgOK)
	t.Run("OKs", photoTgOKs)
	t.Run("Code10", photoTgCode10)
	t.Run("Code20", photoTgCode20)
}

func photoIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func photoIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func photoIntCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WritePhotoInternet("https://youtube.com22"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoIntCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func photoInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoIntOK)
	t.Run("OKs", photoIntOKs)
	t.Run("Code10", photoIntCode10)
	t.Run("Code20", photoIntCode20)
}

func photoCaptOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
}

func photoCapt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaption("something"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoCapt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteCaption(""); err.Error() != code20 {
		t.Fatal(err)
	}

}

func photoCaption(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoCaptOK)
	t.Run("Code10", photoCapt10)
	t.Run("Code20", photoCapt20)
}

func photoHTML(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
}

func photoMarkdown(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteParseMode(types.Markdown); err != nil {
		t.Fatal(err)
	}
}

func photoMarkdownV2(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteParseMode(types.MarkdownV2); err != nil {
		t.Fatal(err)
	}
}

func photoPmOK(t *testing.T) {
	t.Parallel()
	t.Run("types.HTML", photoHTML)
	t.Run("types.Markdown", photoMarkdown)
	t.Run("types.MarkdownV2", photoMarkdownV2)
}

func photoPm10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteParseMode(types.MarkdownV2); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.MarkdownV2); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoPm20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteParseMode("ASKLDL"); err.Error() != code20 {
		t.Fatal(err)
	}
}

func photoParseMode(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoPmOK)
	t.Run("Code10", photoPm10)
	t.Run("Code20", photoPm20)
}

func photoCaptEntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
}

func photoCaptEnt5(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	f := make([]*types.MessageEntity, 1)
	if err := ph.WriteCaptionEntities(f); err.Error() != code5 {
		t.Fatal(err)
	}
}

func photoCaptEnt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ph.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteCaptionEntities(f); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoCaptEnt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteCaptionEntities([]*types.MessageEntity{}); err.Error() != code20 {
		t.Fatal(err)
	}
}

func photoCaptionEntities(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoCaptEntOK)
	t.Run("Code5", photoCaptEnt5)
	t.Run("Code10", photoCaptEnt10)
	t.Run("Code20", photoCaptEnt20)
}

func photoAmOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func photoAm10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteShowCaptionAboveMedia(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoAboveMedia(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoAmOK)
	t.Run("Code10", photoAm10)
}

func photoHsK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
}

func photoHs10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func photoHasSpoiler(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoHsK)
	t.Run("Code10", photoHs10)
}

func photoFunctional(t *testing.T) {
	t.Parallel()
	t.Run("PhotoStorageReqWithoutName", photoStorageReqWithoutName)
	t.Run("PhotoStorageReqWithName", photoStorageReqWithName)
	t.Run("PhotoStorageUnreqWithoutName", photoStorageUnreqWithoutName)
	t.Run("PhotoStorageUnreqWithName", photoStorageUnreqWithName)
	t.Run("PhotoStorageUnreqReplyKB", photoStorageUnreqReplyKB)
	t.Run("PhotoTelegramReqWithoutName", photoTelegramReqWithoutName)
	t.Run("PhotoTelegramReqWithName", photoTelegramReqWithName)
	t.Run("PhotoTelegramUnReqWithoutName", photoTelegramUnReqWithoutName)
	t.Run("PhotoTelegramUnReqWithName", photoTelegramUnReqWithName)
	t.Run("PhotoTelegramUnreqReplyKB", photoTelegramUnreqReplyKB)
}

func photoUnit(t *testing.T) {
	t.Parallel()
	t.Run("WritePhotoStorage", photoStorage)
	t.Run("WritePhotoTelegram", photoTG)
	t.Run("WritePhotoInternet", photoInternet)
	t.Run("WriteCaption", photoCaption)
	t.Run("WriteParseMode", photoParseMode)
	t.Run("WriteCaptionEntities", photoCaptionEntities)
	t.Run("WriteShowCaptionAboveMedia", photoAboveMedia)
	t.Run("WriteHasSpoiler", photoHasSpoiler)
	t.Run("Code21", photoCode21)
}

func TestSendPhoto(t *testing.T) {
	t.Run("Functional", photoFunctional)
	t.Run("Unit", photoUnit)
}
