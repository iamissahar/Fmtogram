package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func animStorageReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(an.GetResponse())
	t.Log(ch.GetResponse())
}

func animStorageReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Animation)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(an.GetResponse())
	t.Log(ch.GetResponse())
}

func animStorageUnreqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(12); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(55); err != nil {
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
	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(an.GetResponse())
}

func animStorageUnreqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(12); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(55); err != nil {
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
	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Animation)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(an.GetResponse())
}

func animStorageUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(12); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(55); err != nil {
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
	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Animation)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(an.GetResponse())
}

func animTelegramReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(an.GetResponse())
	t.Log(ch.GetResponse())
}

func animTelegramReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}

	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Animation)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(an.GetResponse())
	t.Log(ch.GetResponse())
}

func animTelegramUnReqWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(12); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(55); err != nil {
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
	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(an.GetResponse())
}

func animTelegramUnReqWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(12); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(55); err != nil {
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
	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Animation)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(an.GetResponse())
}

func animTelegramUnreqReplyKB(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)

	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(124); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(12); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(55); err != nil {
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
	if err := msg.AddAnimation(an); err != nil {
		t.Fatal(err)
	}
	msg.AddMethod(methods.Animation)

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch.GetResponse())
	t.Log(inf.GetResponse())
	t.Log(an.GetResponse())
}

func animFunctional(t *testing.T) {
	t.Parallel()
	t.Run("AnimationStorageReqWithoutName", animStorageReqWithoutName)
	t.Run("AnimationStorageReqWithName", animStorageReqWithName)
	t.Run("AnimationStorageUnreqWithoutName", animStorageUnreqWithoutName)
	t.Run("AnimationStorageUnreqWithName", animStorageUnreqWithName)
	t.Run("AnimationStorageUnreqReplyKB", animStorageUnreqReplyKB)
	t.Run("AnimationTelegramReqWithoutName", animTelegramReqWithoutName)
	t.Run("AnimationTelegramReqWithName", animTelegramReqWithName)
	t.Run("AnimationTelegramUnReqWithoutName", animTelegramUnReqWithoutName)
	t.Run("AnimationTelegramUnReqWithName", animTelegramUnReqWithName)
	t.Run("AnimationTelegramUnreqReplyKB", animTelegramUnreqReplyKB)
}

func animStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}
}

func animStOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(44); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(11); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func animSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteAnimationStorage("media_test/prichinatryski.mp4"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage("media_test/dimaJOSKAproNATO.ogg"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func animSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", animStOK)
	t.Run("OKs", animStOKs)
	t.Run("Code10", animSt10)
	t.Run("Code12", animSt12)
	t.Run("Code20", animSt20)
}

func animTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}
}

func animTgOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(44); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(11); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func animTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", animTgOK)
	t.Run("OKs", animTgOKs)
	t.Run("Code10", animTg10)
	t.Run("Code20", animTg20)
}

func animIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func animIntOKs(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(123); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(44); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(11); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func animInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteAnimationInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteAnimationInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", animIntOK)
	t.Run("OKs", animIntOKs)
	t.Run("Code10", animInt10)
	t.Run("Code20", animInt20)
}

func animDurOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteDuration(1354); err != nil {
		t.Fatal(err)
	}
}

func animDur10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteDuration(1354); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteDuration(1354); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animDur20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteDuration(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationDuration(t *testing.T) {
	t.Parallel()
	t.Run("OK", animDurOK)
	t.Run("Code10", animDur10)
	t.Run("Code20", animDur20)
}

func animWidthOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteWidth(1245); err != nil {
		t.Fatal(err)
	}
}

func animWidth10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteWidth(1245); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteWidth(1245); err.Error() != code10 {
		t.Fatal(err)
	}
}
func animWidth20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteWidth(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationWidth(t *testing.T) {
	t.Parallel()
	t.Run("OK", animWidthOK)
	t.Run("Code10", animWidth10)
	t.Run("Code20", animWidth20)
}

func animHeightOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteHeight(55); err != nil {
		t.Fatal(err)
	}
}

func animHeight10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteHeight(55); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHeight(55); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animHeight20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteHeight(0); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationHeight(t *testing.T) {
	t.Parallel()
	t.Run("OK", animHeightOK)
	t.Run("Code10", animHeight10)
	t.Run("Code20", animHeight20)
}

func animThumbStOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
}

func animThumbSt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteThumbnailStorage("media_test/tel-aviv.jpg"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animThumbSt12(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailStorage("media_test/musk.mp4"); err.Error() != code12 {
		t.Fatal(err)
	}
}

func animThumbSt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailStorage(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationThumbStorage(t *testing.T) {
	t.Parallel()
	t.Run("OK", animThumbStOK)
	t.Run("Code10", animThumbSt10)
	t.Run("Code12", animThumbSt12)
	t.Run("Code20", animThumbSt20)
}

func animThumbTgOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailTelegram("alkmsdkl;asdkl"); err != nil {
		t.Fatal(err)
	}
}

func animThumbTg10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailTelegram("alkmsdkl;asdkl"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteThumbnailTelegram("alkmsdkl;asdkl"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animThumbTg20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailTelegram(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationThumbTelegram(t *testing.T) {
	t.Parallel()
	t.Run("OK", animThumbTgOK)
	t.Run("Code10", animThumbTg10)
	t.Run("Code20", animThumbTg20)
}

func animThumbIntOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
}

func animThumbInt10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailInternet("https://youtube.com"); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteThumbnailInternet("https://youtube.com"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animThumbInt20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteThumbnailInternet(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func animationThumbInternet(t *testing.T) {
	t.Parallel()
	t.Run("OK", animThumbIntOK)
	t.Run("Code10", animThumbInt10)
	t.Run("Code20", animThumbInt20)
}

func animHSpoilerOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
}

func animHSpoiler10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := an.WriteHasSpoiler(); err.Error() != code10 {
		t.Fatal(err)
	}
}

func animationHSpoiler(t *testing.T) {
	t.Parallel()
	t.Run("OK", animHSpoilerOK)
	t.Run("Code10", animHSpoiler10)
}

func animationCode21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	an := msg.NewAnimation()
	if err := msg.AddAnimation(an); err.Error() != code21 {
		t.Fatal(err)
	}
}

func animUnit(t *testing.T) {
	t.Parallel()
	t.Run("WriteAnimationStorage", animationStorage)
	t.Run("WriteAnimationTelegram", animationTelegram)
	t.Run("WriteAnimationInternet", animationInternet)
	t.Run("WriteDuration", animationDuration)
	t.Run("WriteWidth", animationWidth)
	t.Run("WriteHeight", animationHeight)
	t.Run("WriteThumbnailStorage", animationThumbStorage)
	t.Run("WriteThumbnailTelegram", animationThumbTelegram)
	t.Run("WriteThumbnailInternet", animationThumbInternet)
	t.Run("WriteHasSpoiler", animationHSpoiler)
	t.Run("Code21", animationCode21)
}

func TestSendAnimation(t *testing.T) {
	t.Run("Functional", animFunctional)
	t.Run("Unit", animUnit)
}
