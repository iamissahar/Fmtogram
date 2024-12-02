package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func voiceStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vc.GetResponse())
	t.Log(ch.GetResponse())
}

func voiceStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Voice)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vc.GetResponse())
	t.Log(ch.GetResponse())
}

func voiceStorageUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err != nil {
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
	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vc.GetResponse())
}

func voiceStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err != nil {
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
	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Voice)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vc.GetResponse())
}

func voiceStorageUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/black.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err != nil {
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
	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Voice)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vc.GetResponse())
}

func voiceTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AwACAgIAAxkDAAIDO2cpOUinskD9C8TqK2VytX6vThYDAAL7WwAC43ZRSU1OLw_dlkGCNgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vc.GetResponse())
	t.Log(ch.GetResponse())
}

func voiceTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AwACAgIAAxkDAAIDO2cpOUinskD9C8TqK2VytX6vThYDAAL7WwAC43ZRSU1OLw_dlkGCNgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Voice)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vc.GetResponse())
	t.Log(ch.GetResponse())
}

func voiceTelegramUnReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AwACAgIAAxkDAAIDO2cpOUinskD9C8TqK2VytX6vThYDAAL7WwAC43ZRSU1OLw_dlkGCNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err != nil {
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
	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vc.GetResponse())
}

func voiceTelegramUnReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AwACAgIAAxkDAAIDO2cpOUinskD9C8TqK2VytX6vThYDAAL7WwAC43ZRSU1OLw_dlkGCNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err != nil {
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
	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Voice)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vc.GetResponse())
}

func voiceTelegramUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AwACAgIAAxkDAAIDO2cpOUinskD9C8TqK2VytX6vThYDAAL7WwAC43ZRSU1OLw_dlkGCNgQ"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err != nil {
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
	if err := msg.AddVoice(vc); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Voice)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(vc.GetResponse())
}

func voiceFunctional(t *testing.T) {
	t.Parallel()
	t.Run("VoiceStorageReqWithoutName", voiceStorageReqWithoutName)
	t.Run("VoiceStorageReqWithName", voiceStorageReqWithName)
	t.Run("VoiceStorageUnreqWithoutName", voiceStorageUnreqWithoutName)
	t.Run("VoiceStorageUnreqWithName", voiceStorageUnreqWithName)
	t.Run("VoiceStorageUnreqReplyKB", voiceStorageUnreqReplyKB)
	t.Run("VoiceTelegramReqWithoutName", voiceTelegramReqWithoutName)
	t.Run("VoiceTelegramReqWithName", voiceTelegramReqWithName)
	t.Run("VoiceTelegramUnReqWithoutName", voiceTelegramUnReqWithoutName)
	t.Run("VoiceTelegramUnReqWithName", voiceTelegramUnReqWithName)
	t.Run("VoiceTelegramUnreqReplyKB", voiceTelegramUnreqReplyKB)
}

func voiceStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/dimaJOSKAproNATO.ogg"); err != nil {
		t.Fatal(err)
	}
}

func voiceStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/dimaJOSKAproNATO.ogg"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
}

func voiceSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/dimaJOSKAproNATO.ogg"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteVoiceStorage("media_test/dimaJOSKAproNATO.ogg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func voiceSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage("media_test/photo1.jpg"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func voiceSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func voiceStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", voiceStOK)
	t.Run("OKs", voiceStOKs)
	t.Run("Code10", voiceSt10)
	t.Run("Code12", voiceSt12)
	t.Run("Code20", voiceSt20)
}

func voiceTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("LASdl;aL:DASD"); err != nil {
		t.Fatal(err)
	}
}

func voiceTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("LASdl;aL:DASD"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
}

func voiceTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("LASdl;aL:DASD"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteVoiceTelegram("LASdl;aL:DASD"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func voiceTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func voiceTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", voiceTgOK)
	t.Run("OKs", voiceTgOKs)
	t.Run("Code10", voiceTg10)
	t.Run("Code20", voiceTg20)
}

func voiceIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func voiceIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
}

func voiceInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteVoiceInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func voiceInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteVoiceInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func voiceInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", voiceIntOK)
	t.Run("OKs", voiceIntOKs)
	t.Run("Code10", voiceInt10)
	t.Run("Code20", voiceInt20)
}

func voiceDurOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
}

func voiceDur10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := vc.WriteDuration(124); err.Error() != code10 {
		t.Fatal(err)
	}
}

func voiceDur20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vc := msg.NewVoice()
	if err := vc.WriteDuration(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func voiceDuration(t *testing.T) {
	t.Parallel()
	t.Run("OK", voiceDurOK)
	t.Run("Code10", voiceDur10)
	t.Run("Code20", voiceDur20)
}

func voiceUnit(t *testing.T) {
	t.Parallel()
	t.Run("WriteVoiceStorage", voiceStorage)
	t.Run("WriteVoiceTelegram", voiceTelegram)
	t.Run("WriteVoiceInternet", voiceInternet)
	t.Run("WriteDuration", voiceDuration)
}

func TestSendVoice(t *testing.T) {
	t.Run("Functional", voiceFunctional)
	t.Run("Unit", voiceUnit)
}
