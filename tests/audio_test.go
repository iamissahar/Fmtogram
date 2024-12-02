package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

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
}

func audioStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err, err != nil)
	}
}

func audioStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ad.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(11); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("SHALOM"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("YA PERDOL'E"); err != nil {
		t.Fatal(err)
	}
}

func audioStCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioStCode12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sticker.webp"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func audioStCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioStOK)
	t.Run("ALO", audioStOKs)
	t.Run("Code10", audioStCode10)
	t.Run("Code12", audioStCode12)
	t.Run("Code20", audioStCode20)
}

func audioTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err, err != nil)
	}
}

func audioTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ad.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(11); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("SHALOM"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("YA PERDOL'E"); err != nil {
		t.Fatal(err)
	}
}

func audioTgCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioTgCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioTG(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioTgOK)
	t.Run("OKs", audioTgOKs)
	t.Run("Code10", audioTgCode10)
	t.Run("Code20", audioTgCode20)
}

func audioIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func audioIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteCaption("something"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ad.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(11); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("SHALOM"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("YA PERDOL'E"); err != nil {
		t.Fatal(err)
	}
}

func audioIntCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteAudioInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioIntCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioIntOK)
	t.Run("OKs", audioIntOKs)
	t.Run("Code10", audioIntCode10)
	t.Run("Code20", audioIntCode20)
}

func audioCaptOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
}

func audioCapt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteCaption("SHALOM"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioCapt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteCaption(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioCaption(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioCaptOK)
	t.Run("Code10", audioCapt10)
	t.Run("Code20", audioCapt20)
}

func audioPmOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
}

func audioPm10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteParseMode(types.MarkdownV2); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioPm20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteParseMode("???"); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioParseMode(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioPmOK)
	t.Run("Code10", audioPm10)
	t.Run("Code20", audioPm20)
}

func audioCaptEntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ad.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
}

func audioCaptEnt5(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	f := make([]*types.MessageEntity, 3)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ad.WriteCaptionEntities(f); err.Error() != code5 {
		t.Fatal(err)
	}
}

func audioCaptEnt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := ad.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteCaptionEntities(f); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioCaptEnt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteCaptionEntities([]*types.MessageEntity{}); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioCaptionEntities(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioCaptEntOK)
	t.Run("Code5", audioCaptEnt5)
	t.Run("Code10", audioCaptEnt10)
	t.Run("Code20", audioCaptEnt20)
}

func audioDurOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
}

func audioDur10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(123); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioDur20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteDuration(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioDuration(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioDurOK)
	t.Run("Code10", audioDur10)
	t.Run("Code20", audioDur20)
}

func audioPerfOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WritePerformer("!@#!"); err != nil {
		t.Fatal(err)
	}
}

func audioPerf10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WritePerformer("!@#!"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("!@#!"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioPerf20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WritePerformer(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioPerformer(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioPerfOK)
	t.Run("Code10", audioPerf10)
	t.Run("Code20", audioPerf20)
}

func audioThumbStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func audioThumbSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioThumbSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailStorage("media_test/sound.mp3"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func audioThumbSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioThumbnailStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioThumbStOK)
	t.Run("Code10", audioThumbSt10)
	t.Run("Code12", audioThumbSt12)
	t.Run("Code20", audioThumbSt20)
}

func audioThumbTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailTelegram("asl;dkaksldakl"); err != nil {
		t.Fatal(err)
	}
}

func audioThumbTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailTelegram("asl;dkaksldakl"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailTelegram("asl;dkaksldakl"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioThumbTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioThumbnailTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioThumbTgOK)
	t.Run("Code10", audioThumbTg10)
	t.Run("Code20", audioThumbTg20)
}

func audioThumbIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func audioThumbInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioThumbInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteThumbnailInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioThumbnailInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioThumbIntOK)
	t.Run("Code10", audioThumbInt10)
	t.Run("Code20", audioThumbInt20)
}

func audioTitleOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteTitle("AS?DA?"); err != nil {
		t.Fatal(err)
	}
}

func audioTitle10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteTitle("AS?DA?"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("AS?DA?"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func audioTitle20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteTitle(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func audioTitle(t *testing.T) {
	t.Parallel()
	t.Run("OK", audioTitleOK)
	t.Run("Code10", audioTitle10)
	t.Run("Code20", audioTitle20)
}

func audioCode21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := msg.AddAudio(ad); err.Error() != code21 {
		t.Fatal(err)
	}
}

func audioUnit(t *testing.T) {
	t.Parallel()
	t.Run("WriteAudioStorage", audioStorage)
	t.Run("WriteAudioTelegram", audioTG)
	t.Run("WriteAudioInternet", audioInternet)
	t.Run("WriteCaption", audioCaption)
	t.Run("WriteParseMode", audioParseMode)
	t.Run("WriteCaptionEntities", audioCaptionEntities)
	t.Run("WriteDuration", audioDuration)
	t.Run("WritePerformer", audioPerformer)
	t.Run("WriteThumbnailStorage", audioThumbnailStorage)
	t.Run("WriteThumbnailTelegram", audioThumbnailTelegram)
	t.Run("WriteThumbnailInternet", audioThumbnailInternet)
	t.Run("WriteTitle", audioTitle)
	t.Run("Code21", audioCode21)
}

func TestSendAudio(t *testing.T) {
	t.Run("Functional", audioFunctional)
	t.Run("Unit", audioUnit)
}
