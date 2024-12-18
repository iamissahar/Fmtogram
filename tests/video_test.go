package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func videoStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}

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
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
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

	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}

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

	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
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
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}

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
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
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

	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}

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

	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}

	inf := msg.NewParameters()
	if err := inf.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteString("<b>SELAM</b> ALEYKUM!"); err != nil {
		t.Fatal(err)
	}
	if err := inf.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := inf.WriteEntities(f); err != nil {
		t.Fatal(err)
	}

	createInlineKeyboard(msg, t)

	if err := msg.AddParameters(inf); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Video)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vd.GetResponse())
}

func videoFunctional(t *testing.T) {
	t.Run("VideoStorageReqWithoutName", videoStorageReqWithoutName)
	t.Run("VideoStorageReqWithName", videoStorageReqWithName)
	t.Run("VideoStorageUnreqWithoutName", videoStorageUnreqWithoutName)
	t.Run("VideoStorageUnreqWithName", videoStorageUnreqWithName)
	t.Run("VideoTelegramReqWithoutName", videoTelegramReqWithoutName)
	t.Run("VideoTelegramReqWithName", videoTelegramReqWithName)
	t.Run("VideoTelegramUnreqWithoutName", videoTelegramUnreqWithoutName)
	t.Run("VideoTelegramUnreqWithName", videoTelegramUnreqWithName)
}

func videoStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
}

func videoStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteCaption("?A?SDASLK"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := vd.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
}

func videoStCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoStCode12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/photo1.jpg"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func videoStCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoStOK)
	t.Run("OKs", videoStOKs)
	t.Run("Code10", videoStCode10)
	t.Run("Code12", videoStCode12)
	t.Run("Code20", videoStCode20)
}

func videoTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}
}

func videoTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteCaption("?A?SDASLK"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := vd.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
}

func videoTgCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoTgCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoTG(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoTgOK)
	t.Run("Oks", videoTgOKs)
	t.Run("Code10", videoTgCode10)
	t.Run("Code20", videoTgCode20)
}

func videoIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func videoIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteCaption("?A?SDASLK"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := vd.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(490); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(886); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
}

func videoIntCode10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteVideoInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoIntCode20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoIntOK)
	t.Run("OKs", videoIntOKs)
	t.Run("Code10", videoIntCode10)
	t.Run("Code20", videoIntCode20)
}

func videoCaptOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
}

func videoCapt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteCaption("SHALOM"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoCapt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteCaption(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoCaption(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoCaptOK)
	t.Run("Code10", videoCapt10)
	t.Run("Code20", videoCapt20)
}

func videoPrsmOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
}

func videoPrsm10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteParseMode(types.Markdown); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteParseMode(types.MarkdownV2); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoPrsm20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteParseMode("?ASDASK:L"); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoParseMode(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoPrsmOK)
	t.Run("Code10", videoPrsm10)
	t.Run("Code20", videoPrsm20)
}

func videoCaptEntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := vd.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
}

func videoCaptEnt5(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	f := make([]*types.MessageEntity, 3)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := vd.WriteCaptionEntities(f); err.Error() != code5 {
		t.Fatal(err)
	}
}

func videoCaptEnt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := vd.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteCaptionEntities(f); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoCaptEnt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteCaptionEntities([]*types.MessageEntity{}); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoCaptionEntities(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoCaptEntOK)
	t.Run("Code5", videoCaptEnt5)
	t.Run("Code10", videoCaptEnt10)
	t.Run("Code20", videoCaptEnt20)
}

func videoHSpoilerOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
}

func videoHSpoiler10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoHasSpoiler(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoHSpoilerOK)
	t.Run("Code10", videoHSpoiler10)
}

func videoHeightOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteHeight(124); err != nil {
		t.Fatal(err)
	}
}

func videoHeight10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteHeight(124); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(124); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoHeight20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteHeight(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoHeight(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoHeightOK)
	t.Run("Code10", videoHeight10)
	t.Run("Code20", videoHeight20)
}

func videoWidthOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteWidth(124); err != nil {
		t.Fatal(err)
	}
}

func videoWidth10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteWidth(124); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(124); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoWidth20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteWidth(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoWidth(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoWidthOK)
	t.Run("Code10", videoWidth10)
	t.Run("Code20", videoWidth20)
}

func videoSCAMOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func videoSCAM10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteShowCaptionAboveMedia(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoShowCaptionAboveMedia(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoSCAMOK)
	t.Run("Code10", videoSCAM10)
}

func videoSupStrK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
}

func videoSupStr10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteSupportStreaming(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportStreaming(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoSupportsStreaming(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoSupStrK)
	t.Run("Code10", videoSupStr10)
}

func videoThumbStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func videoThumbSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoThumbSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailStorage("media_test/Resume.pdf"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func videoThumbSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoThumbnailStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoThumbStOK)
	t.Run("Code10", videoThumbSt10)
	t.Run("Code12", videoThumbSt12)
	t.Run("Code20", videoThumbSt20)
}

func videoThumbTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailTelegram("askldklas1233"); err != nil {
		t.Fatal(err)
	}
}

func videoThumbTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailTelegram("askldklas1233"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailTelegram("askldklas1233"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoThumbTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoThumbnailTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoThumbTgOK)
	t.Run("Code10", videoThumbTg10)
	t.Run("Code20", videoThumbTg20)
}

func videoThumbIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func videoThumbInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func videoThumbInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteThumbnailInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func videoThumbnailInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", videoThumbIntOK)
	t.Run("Code10", videoThumbInt10)
	t.Run("Code20", videoThumbInt20)
}

func videoCode21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := msg.AddVideo(vd); err.Error() != code21 {
		t.Fatal(err)
	}
}

func videoUnit(t *testing.T) {
	t.Parallel()
	t.Run("WriteVideoStorage", videoStorage)
	t.Run("WriteVideoTelegram", videoTG)
	t.Run("WriteVideoInternet", videoInternet)
	t.Run("WriteCaption", videoCaption)
	t.Run("WriteParseMode", videoParseMode)
	t.Run("WriteCaptionEntities", videoCaptionEntities)
	t.Run("WriteHasSpoiler", videoHasSpoiler)
	t.Run("WriteHeight", videoHeight)
	t.Run("WriteWidth", videoWidth)
	t.Run("WriteShowCaptionAboveMedia", videoShowCaptionAboveMedia)
	t.Run("WriteSupportStreaming", videoSupportsStreaming)
	t.Run("WriteThumbnailStorage", videoThumbnailStorage)
	t.Run("WriteThumbnailTelegram", videoThumbnailTelegram)
	t.Run("WriteThumbnailInternet", videoThumbnailInternet)
	t.Run("Code21", videoCode21)
}

func TestSendVideo(t *testing.T) {
	t.Run("Functionl", videoFunctional)
	t.Run("Unit", videoUnit)
}
