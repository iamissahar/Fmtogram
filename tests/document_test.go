package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func docStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dc.GetResponse())
	t.Log(ch.GetResponse())
}

func docStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Document)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dc.GetResponse())
	t.Log(ch.GetResponse())
}

func docStorageUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
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
	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(dc.GetResponse())
}

func docStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
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
	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Document)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(dc.GetResponse())
}

func docStorageUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
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
	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Document)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(dc.GetResponse())
}

func docTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dc.GetResponse())
	t.Log(ch.GetResponse())
}

func docTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Document)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dc.GetResponse())
	t.Log(ch.GetResponse())
}

func docTelegramUnReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
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
	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(dc.GetResponse())
}

func docTelegramUnReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
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
	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Document)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(dc.GetResponse())
}

func docTelegramUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()

	ch := createChat(msg, t)

	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
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
	if err := msg.AddDocument(dc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Document)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(dc.GetResponse())
}

func documentFunctional(t *testing.T) {
	t.Parallel()
	t.Run("DocumentStorageReqWithoutName", docStorageReqWithoutName)
	t.Run("DocumentStorageReqWithName", docStorageReqWithName)
	t.Run("DocumentStorageUnreqWithoutName", docStorageUnreqWithoutName)
	t.Run("DocumentStorageUnreqWithName", docStorageUnreqWithName)
	t.Run("DocumentStorageUnreqReplyKB", docStorageUnreqReplyKB)
	t.Run("DocumentTelegramReqWithoutName", docTelegramReqWithoutName)
	t.Run("DocumentTelegramReqWithName", docTelegramReqWithName)
	t.Run("DocumentTelegramUnReqWithoutName", docTelegramUnReqWithoutName)
	t.Run("DocumentTelegramUnReqWithName", docTelegramUnReqWithName)
	t.Run("DocumentTelegramUnreqReplyKB", docTelegramUnreqReplyKB)
}

func docStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
}

func docStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := dc.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteParseMode(types.MarkdownV2); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteThumbnailStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
}

func docSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDocumentStorage("media_test/Resume.pdf"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", docStOK)
	t.Run("OKs", docStOKs)
	t.Run("Code10", docSt10)
	t.Run("Code20", docSt20)
}

func docTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}
}

func docTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := dc.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteParseMode(types.MarkdownV2); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteThumbnailStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
}

func docTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", docTgOK)
	t.Run("OKs", docTgOKs)
	t.Run("Code10", docTg10)
	t.Run("Code20", docTg20)
}

func docIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func docIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteCaption("SHALOM"); err != nil {
		t.Fatal(err)
	}
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := dc.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteParseMode(types.MarkdownV2); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteThumbnailStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
}

func docInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDocumentInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDocumentInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", docIntOK)
	t.Run("OKs", docIntOKs)
	t.Run("Code10", docInt10)
	t.Run("Code20", docInt20)
}

func docCaptOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteCaption("SHALOM <3"); err != nil {
		t.Fatal(err)
	}
}

func docCapt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteCaption("SHALOM <3"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteCaption("SHALOM <3"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docCapt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteCaption(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docCaption(t *testing.T) {
	t.Parallel()
	t.Run("OK", docCaptOK)
	t.Run("Code10", docCapt10)
	t.Run("Code20", docCapt20)
}

func docCaptEntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := dc.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
}

func docCaptEnt5(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	f := make([]*types.MessageEntity, 3)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := dc.WriteCaptionEntities(f); err.Error() != code5 {
		t.Fatal(err)
	}
}

func docCaptEnt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	f := make([]*types.MessageEntity, 1)
	f[0] = &types.MessageEntity{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}
	if err := dc.WriteCaptionEntities(f); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteCaptionEntities(f); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docCaptEnt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteCaptionEntities([]*types.MessageEntity{}); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docCaptionEntities(t *testing.T) {
	t.Parallel()
	t.Run("OK", docCaptEntOK)
	t.Run("Code5", docCaptEnt5)
	t.Run("Code10", docCaptEnt10)
	t.Run("Code20", docCaptEnt20)
}

func docDisContDetOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
}

func docDisContDet10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteDisableContentTypeDetection(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docDisContDet(t *testing.T) {
	t.Parallel()
	t.Run("OK", docDisContDetOK)
	t.Run("Code10", docDisContDet10)
}

func docPrsmOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
}

func docPrsm10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteParseMode(types.Markdown); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docPrsm20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteParseMode("ASDKLDAKLS"); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docParsemode(t *testing.T) {
	t.Parallel()
	t.Run("OK", docPrsmOK)
	t.Run("Code10", docPrsm10)
	t.Run("Code20", docPrsm20)
}

func docThumbStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func docThumbSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docThumbSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailStorage("media_test/black.mp4"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func docThumbSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docThumbStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", docThumbStOK)
	t.Run("Code10", docThumbSt10)
	t.Run("Code12", docThumbSt12)
	t.Run("Code20", docThumbSt20)
}

func docThumbTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailTelegram("aklsjdkl;jasd"); err != nil {
		t.Fatal(err)
	}
}

func docThumbTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailTelegram("aklsjdkl;jasd"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteThumbnailTelegram("aklsjdkl;jasd"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docThumbTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docThumbTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", docThumbTgOK)
	t.Run("Code10", docThumbTg10)
	t.Run("Code20", docThumbTg20)
}

func docThumbIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func docThumbInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := dc.WriteThumbnailInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func docThumbInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := dc.WriteThumbnailInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func docThumbInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", docThumbIntOK)
	t.Run("Code10", docThumbInt10)
	t.Run("Code20", docThumbInt20)
}

func docCode21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	dc := msg.NewDocument()
	if err := msg.AddDocument(dc); err.Error() != code21 {
		t.Fatal(err)
	}
}

func documentUnit(t *testing.T) {
	t.Parallel()
	t.Run("WriteDocumentStorage", docStorage)
	t.Run("WriteDocumentTelegram", docTelegram)
	t.Run("WriteDocumentInternet", docInternet)
	t.Run("WriteCaption", docCaption)
	t.Run("WriteCaptionEntities", docCaptionEntities)
	t.Run("WriteDisableContentTypeDetection", docDisContDet)
	t.Run("WriteParseMode", docParsemode)
	t.Run("WriteThumbnailStorage", docThumbStorage)
	t.Run("WriteThumbnailTelegram", docThumbTelegram)
	t.Run("WriteThumbnailInternet", docThumbInternet)
	t.Run("Code21", docCode21)
}

func TestSendDocument(t *testing.T) {
	t.Run("Functional", documentFunctional)
	t.Run("Unit", documentUnit)
}
