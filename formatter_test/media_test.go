package formatter_test

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	code10 string = "10"
	code20 string = "20"
	code12 string = "12"
	code21 string = "21"
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
		panic(err)
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

func sendPhotoMediaRequired(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	createChat(msg, t)
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.MisMedia()
		if err.Error() != "[ERROR] You're supposed to add a media interface if you specified a media method.\n"+
			"Add a media interface according to your method, or do not specify any method.\n"+
			"Here are all available media interfaces:\n\n"+
			"┌──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐\n"+
			"│  formatter.IPhoto    │  formatter.IAudio    │  formatter.IVideo    │  formatter.IDocument │\n"+
			"├──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤\n"+
			"│  formatter.IAnimation│  formatter.IVoice    │  formatter.IVideoNote│                      │\n"+
			"└──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘" {
			t.Fatal(err)
		}
	}
}

func photoRequired(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	createChat(msg, t)
	ph := msg.NewPhoto()

	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.MissedRequiredField(interfacePhoto, "Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}", num, 0, true, false)
		if err.Error() != `[ERROR] Required data is missing:

┌───────────────┬──────────────────────────────────────────────────────────────┐
│ Interface     │ Required                                                     │
├───────────────┼──────────────────────────────────────────────────────────────┤
│ IPhoto        │ Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}        │
└───────────────┴──────────────────────────────────────────────────────────────┘
Interface was added as 0` {
			t.Fatal(err)
		}
	}
}

func photoAndChatRequired(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := msg.NewChat()
	ph := msg.NewPhoto()

	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddChat(ch); err.Error() != code21 {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	// 	_, err := msg.Send()
	// 	if err != nil {
	// 		// fmerrors.MissedRequiredField("IChat", "WriteChat{ID/Name}()", 0, 0, false, false)
	// 		if err.Error() != `[ERROR] Required data is missing:

	// ┌───────────────┬──────────────────────────────────────────────────────────────┐
	// │ Interface     │ Required                                                     │
	// ├───────────────┼──────────────────────────────────────────────────────────────┤
	// │ IPhoto        │ Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}        │
	// └───────────────┴──────────────────────────────────────────────────────────────┘
	//
	//	Interface was added as 0` {
	//				t.Fatal(err)
	//			}
	//	}
}

func photoWithoutChat(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ph := msg.NewPhoto()
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Photo)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.ChatIDIsMissed()
		if err.Error() != "[ERROR]\tYou are required to provide the data in WriteChat{ID,Name}" {
			t.Fatal(err)
		}
	}
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

func photoStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", photoStOK)
	t.Run("ALO", photoStOKs)
	t.Run("Code10", photoStCode10)
	t.Run("Code12", photoStCode12)
	t.Run("Code20", photoStCode20)
	// 21
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
	if err := ph.WriteCaptionEntities(f); err.Error() != "5" {
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
	t.Run("MediaRequired", sendPhotoMediaRequired)
	t.Run("PhotoRequired", photoRequired)
	t.Run("PhotoAndChatRequired", photoAndChatRequired)
	t.Run("PhotoWithoutChat", photoWithoutChat)
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
}

func TestSendPhoto(t *testing.T) {
	t.Run("Functional", photoFunctional)
	t.Run("Unit", photoUnit)
}

func audioStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}

	msg.AddAudio(ad)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(ad.GetResponse())
}

func audioStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}

	msg.AddMethod(methods.Audio)
	msg.AddAudio(ad)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(ad.GetResponse())
}

func audioStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("?"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("A SOUND"); err != nil {
		t.Fatal(err)
	}
	inf := msg.NewMessage()
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	msg.AddAudio(ad)
	msg.AddParameters(inf)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ad.GetResponse())
}

func audioStorageUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("?"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("A SOUND"); err != nil {
		t.Fatal(err)
	}
	inf := msg.NewMessage()
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}

	createReplyKeyboard(msg, t)

	msg.AddAudio(ad)
	msg.AddParameters(inf)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ad.GetResponse())
}

func audioStorageUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("?"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("A SOUND"); err != nil {
		t.Fatal(err)
	}
	inf := msg.NewMessage()
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	msg.AddAudio(ad)
	msg.AddParameters(inf)
	msg.AddMethod(methods.Audio)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ad.GetResponse())
}

func audioTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}

	msg.AddAudio(ad)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(ad.GetResponse())
}

func audioTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}

	msg.AddAudio(ad)
	msg.AddMethod(methods.Audio)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(ad.GetResponse())
}

func audioTelegramUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("?"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("A SOUND"); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewMessage()
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	msg.AddAudio(ad)
	msg.AddParameters(inf)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ad.GetResponse())
}

func audioTelegramUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("?"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("A SOUND"); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewMessage()
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}

	createReplyKeyboard(msg, t)

	msg.AddAudio(ad)
	msg.AddParameters(inf)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ad.GetResponse())
}

func sendAudioMediaRequired(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	createChat(msg, t)
	msg.AddMethod(methods.Audio)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.MisMedia()
		if err.Error() != "[ERROR] You're supposed to add a media interface if you specified a media method.\n"+
			"Add a media interface according to your method, or do not specify any method.\n"+
			"Here are all available media interfaces:\n\n"+
			"┌──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐\n"+
			"│  formatter.IPhoto    │  formatter.IAudio    │  formatter.IVideo    │  formatter.IDocument │\n"+
			"├──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤\n"+
			"│  formatter.IAnimation│  formatter.IVoice    │  formatter.IVideoNote│                      │\n"+
			"└──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘" {
			t.Fatal(err)
		}
	}
}

func audioTelegramUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	ad := msg.NewAudio()
	ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE")
	ad.WriteCaption("<b>SELAM</b> ALEYKUM!")
	ad.WriteParseMode(types.HTML)
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	ad.WriteCaptionEntities(f)
	ad.WriteDuration(3)
	ad.WritePerformer("?")
	ad.WriteTitle("A SOUND")

	inf := msg.NewMessage()
	inf.WriteDisableNotification()
	inf.WriteProtectContent()
	inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

	createInlineKeyboard(msg, t)

	msg.AddAudio(ad)
	msg.AddParameters(inf)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(ad.GetResponse())
}

func audioRequired(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	createChat(msg, t)
	ph := msg.NewPhoto()

	msg.AddPhoto(ph)
	msg.AddMethod(methods.Audio)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.MissedRequiredField(interfacePhoto, "Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}", num, 0, true, false)
		if err.Error() != `[ERROR] Required data is missing:

┌───────────────┬──────────────────────────────────────────────────────────────┐
│ Interface     │ Required                                                     │
├───────────────┼──────────────────────────────────────────────────────────────┤
│ IAudio        │ Write{Storage,Telgram,Internet}Audio{FilePath/ID/URL}        │
└───────────────┴──────────────────────────────────────────────────────────────┘
Interface was added as 0` {
			t.Fatal(err)
		}
	}
}

func audioAndChatRequired(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := msg.NewChat()
	ad := msg.NewAudio()

	msg.AddAudio(ad)
	msg.AddChat(ch)
	msg.AddMethod(methods.Audio)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.MissedRequiredField("IChat", "WriteChat{ID/Name}()", 0, 0, false, false)
		if err.Error() != `[ERROR] Required data is missing:

┌───────────────┬──────────────────────────────────────────────────────────────┐
│ Interface     │ Required                                                     │
├───────────────┼──────────────────────────────────────────────────────────────┤
│ IAudio        │ Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}        │
└───────────────┴──────────────────────────────────────────────────────────────┘
Interface was added as 0` {
			t.Fatal(err)
		}
	}
}

func audioWithoutChat(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ad := msg.NewAudio()
	ad.WriteAudioStorage("media_test/photo1.jpg")
	msg.AddAudio(ad)
	msg.AddMethod(methods.Audio)

	_, err := msg.Send()
	if err != nil {
		// fmerrors.ChatIDIsMissed()
		if err.Error() != "[ERROR]\tYou are required to provide the data in WriteChat{ID,Name}" {
			t.Fatal(err)
		}
	}
}

func audioFunctional(t *testing.T) {
	t.Parallel()
	t.Run("AudioStorageReqWithoutName", audioStorageReqWithoutName)
	t.Run("AudioStorageReqWithName", audioStorageReqWithName)
	t.Run("AudioStorageUnreqWithoutName", audioStorageUnreqWithoutName)
	t.Run("AudioStorageUnreqWithName", audioStorageUnreqWithName)
	t.Run("AudioStorageUnreqReplyKB", audioStorageUnreqReplyKB)
	t.Run("AudioTelegramReqWithoutName", audioTelegramReqWithoutName)
	t.Run("AudioTelegramReqWithName", audioTelegramReqWithName)
	t.Run("audioTelegramUnreqWithoutName", audioTelegramUnreqWithoutName)
	t.Run("AudioTelegramUnreqWithName", audioTelegramUnreqWithName)
	t.Run("AudioTelegramUnreqReplyKB", audioTelegramUnreqReplyKB)
	t.Run("MediaRequired", sendAudioMediaRequired)
	t.Run("AudioRequired", audioRequired)
	t.Run("AudioAndChatRequired", audioAndChatRequired)
	t.Run("AudioWithoutChat", audioWithoutChat)
}

// func audioUnit(t *testing.T) {
// 	t.Parallel()
// 	t.Run("WriteAudioStorage", audioStorage)
// 	t.Run("WriteAudioTelegram", audioTG)
// 	t.Run("WriteAudioInternet", audioInternet)
// 	t.Run("WriteCaption", audioCaption)
// 	t.Run("WriteParseMode", audioParseMode)
// 	t.Run("WriteCaptionEntities", audioCaptionEntities)
// 	t.Run("WriteDuration", audioDuration)
// 	t.Run("WritePerformer", audioPerformer)
// 	t.Run("WriteThumbnailStorage", audioThumbnailStorage)
// 	t.Run("WriteThumbnailTelegram", audioThumbnailTelegram)
// 	t.Run("WriteThumbnailInternet", audioThumbnailInternet)
// 	t.Run("WriteTitle", audioTitle)
// }

func TestSendAudio(t *testing.T) {
	t.Run("Functional", audioFunctional)
	// t.Run("Unit", audioUnit)
}

func videoStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()
	vd.WriteVideoStorage("media_test/musk.mp4")

	msg.AddVideo(vd)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(vd.GetResponse())
}

func videoStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()
	vd.WriteVideoStorage("media_test/musk.mp4")

	msg.AddVideo(vd)
	msg.AddMethod(methods.Video)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(vd.GetResponse())
}

func videoStorageUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()

	vd.WriteVideoStorage("media_test/musk.mp4")
	vd.WriteCaption("<b>SELAM</b> ALEYKUM!")
	vd.WriteParseMode(types.HTML)
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	vd.WriteCaptionEntities(f)
	vd.WriteDuration(3)
	vd.WriteHeight(490)
	vd.WriteWidth(886)
	vd.WriteSupportsStreaming()
	vd.WriteHasSpoiler()

	inf := msg.NewMessage()
	inf.WriteDisableNotification()
	inf.WriteProtectContent()
	inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

	createInlineKeyboard(msg, t)

	msg.AddParameters(inf)
	msg.AddVideo(vd)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vd.GetResponse())
}

func videoStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()

	vd.WriteVideoStorage("media_test/musk.mp4")
	vd.WriteCaption("<b>SELAM</b> ALEYKUM!")
	vd.WriteParseMode(types.HTML)
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	vd.WriteCaptionEntities(f)
	vd.WriteDuration(3)
	vd.WriteHeight(490)
	vd.WriteWidth(886)
	vd.WriteSupportsStreaming()
	vd.WriteHasSpoiler()

	inf := msg.NewMessage()
	inf.WriteDisableNotification()
	inf.WriteProtectContent()
	inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

	createInlineKeyboard(msg, t)

	msg.AddParameters(inf)
	msg.AddVideo(vd)
	msg.AddMethod(methods.Video)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vd.GetResponse())
}

func videoTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()
	vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")

	msg.AddVideo(vd)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(vd.GetResponse())
}

func videoTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()
	vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")

	msg.AddVideo(vd)
	msg.AddMethod(methods.Video)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(vd.GetResponse())
}

func videoTelegramUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()

	vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
	vd.WriteCaption("<b>SELAM</b> ALEYKUM!")
	vd.WriteParseMode(types.HTML)
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	vd.WriteCaptionEntities(f)
	vd.WriteDuration(3)
	vd.WriteHeight(490)
	vd.WriteWidth(886)
	vd.WriteSupportsStreaming()
	vd.WriteHasSpoiler()

	inf := msg.NewMessage()
	inf.WriteDisableNotification()
	inf.WriteProtectContent()
	inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

	createInlineKeyboard(msg, t)

	msg.AddParameters(inf)
	msg.AddVideo(vd)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vd.GetResponse())
}

func videoTelegramUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()

	vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
	vd.WriteCaption("<b>SELAM</b> ALEYKUM!")
	vd.WriteParseMode(types.HTML)
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	vd.WriteCaptionEntities(f)
	vd.WriteDuration(3)
	vd.WriteHeight(490)
	vd.WriteWidth(886)
	vd.WriteSupportsStreaming()
	vd.WriteHasSpoiler()

	inf := msg.NewMessage()
	inf.WriteDisableNotification()
	inf.WriteProtectContent()
	inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

	createInlineKeyboard(msg, t)

	msg.AddParameters(inf)
	msg.AddVideo(vd)
	msg.AddMethod(methods.Video)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vd.GetResponse())
}

func TestSendVideo(t *testing.T) {
	t.Run("VideoStorageReqWithoutName()", videoStorageReqWithoutName)
	t.Run("VideoStorageReqWithName()", videoStorageReqWithName)
	t.Run("VideoStorageUnreqWithoutName()", videoStorageUnreqWithoutName)
	t.Run("VideoStorageUnreqWithName()", videoStorageUnreqWithName)
	t.Run("VideoTelegramReqWithoutName()", videoTelegramReqWithoutName)
	t.Run("VideoTelegramReqWithName()", videoTelegramReqWithName)
	t.Run("VideoTelegramUnreqWithoutName()", videoTelegramUnreqWithoutName)
	t.Run("VideoTelegramUnreqWithName()", videoTelegramUnreqWithName)
}
