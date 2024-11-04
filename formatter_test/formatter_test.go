package formatter_test

import (
	"testing"

	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func createInlineKeyboard(msg *formatter.Message, t *testing.T) {
	kb := msg.NewInlineKeyboard()
	kb.Set([]int{1, 2})
	button00, err := kb.NewButton(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	button10, err := kb.NewButton(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	button11, err := kb.NewButton(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	button00.WriteString("⭐")
	button00.WriteURL("https://www.youtube.com/watch?v=xsjfuPlrjT0")

	button10.WriteString("Hello!")
	button10.WriteCallbackData("Hello X2")

	button11.WriteString("Sheesh")
	button11.WriteCallbackData("TAMAM")

	msg.AddInlineKeyboard(kb)
}

func createChat(msg *formatter.Message) formatter.IChat {
	chat := msg.NewChat()
	chat.WriteChatID(738070596)
	msg.AddChat(chat)
	return chat
}

func TestGetMe(t *testing.T) {
	getme, err := executer.GetMe()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(getme)
	if !getme.Ok {
		t.Fatal("getme.Ok is false")
	} else if getme.Result == nil {
		t.Fatal("getme.Result is nil")
	}
	t.Log(getme.Result)
}

func TestGetUpdates(t *testing.T) {
	offset := 0
	err := executer.GetOffset(&offset)
	if err != nil {
		t.Log(offset)
		t.Fatal(err)
	}
	t.Log(offset)
	offset++

	tg := new(types.Telegram)
	err = executer.GetUpdates(tg, &offset)
	if err != nil {
		t.Fatal(err)
	}
	if !tg.Ok {
		t.Fatal("tg.Ok is false")
	} else if len(tg.Result) < 1 {
		t.Log("No new messages")
	}

	t.Log(tg.Result[0].Message.Text)
	t.Log(tg.Result[0].Message.MessageID)
	t.Log(tg.Result[0].Message.From)
	t.Log(tg.Result[0].Message.Chat)

}

func TestSendMessage(t *testing.T) {
	t.Run("IMSGInformation IChat", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()
		inf := msg.NewMessage()
		inf.WriteString("Shalom")

		ch := createChat(msg)

		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("IMSGInformation IChat IInline", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()
		inf := msg.NewMessage()
		inf.WriteString("<b>SHALOM</b> TO Y'<b>ALL</b>")
		inf.WriteParseMode(types.HTML)

		ch := createChat(msg)

		createInlineKeyboard(msg, t)

		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(ch.GetResponse())
	})
}

func TestForwardMessage(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(501)

		msg.AddMethod(methods.ForwardMessage)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(501)
		inf.WriteDisableNotification()
		inf.WriteProtectContent()

		msg.AddMethod(methods.ForwardMessage)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(ch.GetResponse())
	})
}

func TestForwardMessages(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})

		msg.AddMethod(methods.ForwardMessages)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetMessageIDs())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})
		inf.WriteDisableNotification()
		inf.WriteProtectContent()

		msg.AddMethod(methods.ForwardMessages)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetMessageIDs())
	})
}

func TestCopyMessage(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(507)

		msg.AddMethod(methods.CopyMessage)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetMessageIDs())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(507)
		inf.WriteCaption("<b>SHALOM</b> Y'ALL!!")
		inf.WriteParseMode(types.HTML)
		f := make([]*types.MessageEntity, 1)
		f[0] = &types.MessageEntity{Type: "/start@jobs_bot"}
		inf.WriteEntities(f)
		inf.WriteDisableNotification()
		inf.WriteProtectContent()
		inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

		createInlineKeyboard(msg, t)

		msg.AddMethod(methods.CopyMessage)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetMessageIDs())
	})
}

func TestCopyMessages(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})

		msg.AddMethod(methods.CopyMessages)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetMessageIDs())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})
		inf.WriteDisableNotification()
		inf.WriteProtectContent()

		msg.AddMethod(methods.CopyMessages)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetMessageIDs())
	})
}

func TestSendPhoto(t *testing.T) {
	t.Run("Required fields, without naming method and photo from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ph := msg.NewPhoto()
		ph.WritePhotoStorage("media_test/photo1.jpg")

		msg.AddPhoto(ph)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(ph.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("Required fields with naming method and photo from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ph := msg.NewPhoto()
		ph.WritePhotoStorage("media_test/photo1.jpg")

		msg.AddMethod(methods.Photo)
		msg.AddPhoto(ph)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(ph.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and photo from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ph := msg.NewPhoto()
		ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph.WriteCaption("<b>SHALOM!</b> Y'ALL!")
		ph.WriteParseMode(types.HTML)
		f := make([]*types.MessageEntity, 1)
		f[0] = &types.MessageEntity{Type: "/start@jobs_bot"}
		ph.WriteCaptionEntities(f)
		ph.WriteShowCaptionAboveMedia()
		ph.WriteHasSpoiler()

		inf := msg.NewMessage()
		inf.WriteDisableNotification()
		inf.WriteProtectContent()
		inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

		createInlineKeyboard(msg, t)

		msg.AddMessage(inf)
		msg.AddPhoto(ph)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(ph.GetResponse())
	})
}

func TestSendAudio(t *testing.T) {
	t.Run("Required fields without naming method and audio from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ad := msg.NewAudio()
		ad.WriteAudioStorage("media_test/sound.mp3")

		msg.AddAudio(ad)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(ad.GetResponse())
	})
	t.Run("Required fields with naming method and audio from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ad := msg.NewAudio()
		ad.WriteAudioStorage("media_test/sound.mp3")

		msg.AddMethod(methods.Audio)
		msg.AddAudio(ad)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(ad.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and audio from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ad := msg.NewAudio()
		ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE")
		ad.WriteCaption("<b>SELAM</b> ALEYKUM!")
		ad.WriteParseMode(types.HTML)
		f := make([]*types.MessageEntity, 1)
		f[0] = &types.MessageEntity{Type: "/start@jobs_bot"}
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
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(ad.GetResponse())
	})
}

func TestSendVideo(t *testing.T) {
	t.Run("Required fields without naming method and video from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vd := msg.NewVideo()
		vd.WriteVideoStorage("media_test/musk.mp4")

		msg.AddVideo(vd)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(vd.GetResponse())
	})
	t.Run("Required fields with naming method and video from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

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
	})
	t.Run("Required and unrequired fields without naming method and video from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vd := msg.NewVideo()

		vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd.WriteCaption("<b>SELAM</b> ALEYKUM!")
		vd.WriteParseMode(types.HTML)
		f := make([]*types.MessageEntity, 1)
		f[0] = &types.MessageEntity{Type: "/start@jobs_bot"}
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

		msg.AddMessage(inf)
		msg.AddVideo(vd)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(vd.GetResponse())
	})
}

func TestSendDocument(t *testing.T) {
	t.Run("Required fields without naming method and document from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		dc := msg.NewDocument()
		dc.WriteDocumentStorage("media_test/Resume.pdf")

		msg.AddDocument(dc)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(dc.GetResponse())
	})
	t.Run("Required fields with naming method and document from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		dc := msg.NewDocument()
		dc.WriteDocumentStorage("media_test/Resume.pdf")

		msg.AddDocument(dc)
		msg.AddMethod(methods.Document)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(dc.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and document from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		dc := msg.NewDocument()

		dc.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ")
		dc.WriteCaption("<b>SELAM</b> ALEYKUM!")
		dc.WriteParseMode(types.HTML)
		f := make([]*types.MessageEntity, 1)
		f[0] = &types.MessageEntity{Type: "/start@jobs_bot"}
		dc.WriteCaptionEntities(f)
		dc.WriteDisableContentTypeDetection()

		inf := msg.NewMessage()
		inf.WriteDisableNotification()
		inf.WriteProtectContent()
		inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

		createInlineKeyboard(msg, t)

		msg.AddMessage(inf)
		msg.AddDocument(dc)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(dc.GetResponse())
	})
}

func TestSendMediaGroup(t *testing.T) {
	t.Run("Required fields without naming method and 1 photo and 1 video from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ph := msg.NewPhoto()
		ph.WritePhotoStorage("media_test/photo1.jpg")

		vd := msg.NewVideo()
		vd.WriteVideoStorage("media_test/musk.mp4")

		msg.AddPhoto(ph)
		msg.AddVideo(vd)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(ph.GetResponse())
		t.Log(vd.GetResponse())
	})
	t.Run("Required fields without naming method and 3 audio (1xstorage, 2xtelegram)", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ad := msg.NewAudio()
		ad1 := msg.NewAudio()
		ad2 := msg.NewAudio()

		ad.WriteAudioStorage("media_test/sound.mp3")
		ad1.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE")
		ad2.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE")

		msg.AddAudio(ad)
		msg.AddAudio(ad1)
		msg.AddAudio(ad2)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(ad.GetResponse())
		t.Log(ad1.GetResponse())
		t.Log(ad2.GetResponse())
	})
	t.Run("Required fields without naming method and 3 audio (1xstorage, 2xtelegram)", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ad := msg.NewAudio()
		ad1 := msg.NewAudio()
		ad2 := msg.NewAudio()

		ad.WriteAudioStorage("media_test/sound.mp3")
		ad1.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE")
		ad2.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE")

		msg.AddAudio(ad)
		msg.AddAudio(ad1)
		msg.AddAudio(ad2)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(ad.GetResponse())
		t.Log(ad1.GetResponse())
		t.Log(ad2.GetResponse())
	})
	t.Run("Required fields without naming method and 3 document (1xstorage, 2xtelegram)", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		dc := msg.NewDocument()
		dc1 := msg.NewDocument()
		dc2 := msg.NewDocument()

		dc.WriteDocumentStorage("media_test/sound.mp3")
		dc1.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ")
		dc2.WriteDocumentTelegram("BQACAgIAAxkDAAICYmcmSintJKtk4ZCtE_H7nD3QoIp8AALrWwACWrQxSa6dehGI5tWYNgQ")

		msg.AddDocument(dc)
		msg.AddDocument(dc1)
		msg.AddDocument(dc2)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(dc.GetResponse())
		t.Log(dc1.GetResponse())
		t.Log(dc2.GetResponse())
	})
	t.Run("Required fields without naming method and 10 photos (2xstorage, 8xtelegram)", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		ph := msg.NewPhoto()
		ph1 := msg.NewPhoto()
		ph2 := msg.NewPhoto()
		ph3 := msg.NewPhoto()
		ph4 := msg.NewPhoto()
		ph5 := msg.NewPhoto()
		ph6 := msg.NewPhoto()
		ph7 := msg.NewPhoto()
		ph8 := msg.NewPhoto()
		ph9 := msg.NewPhoto()

		ph.WritePhotoStorage("media_test/photo1.jpg")
		ph1.WritePhotoStorage("media_test/photo1.jpg")
		ph2.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph3.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph4.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph5.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph6.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph7.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph8.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")
		ph9.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE")

		msg.AddPhoto(ph)
		msg.AddPhoto(ph1)
		msg.AddPhoto(ph2)
		msg.AddPhoto(ph3)
		msg.AddPhoto(ph4)
		msg.AddPhoto(ph5)

		msg.AddPhoto(ph6)
		msg.AddPhoto(ph7)
		msg.AddPhoto(ph8)
		msg.AddPhoto(ph9)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(ph.GetResponse())
		t.Log(ph1.GetResponse())
		t.Log(ph2.GetResponse())
		t.Log(ph3.GetResponse())
		t.Log(ph4.GetResponse())
		t.Log(ph5.GetResponse())
		t.Log(ph6.GetResponse())
		t.Log(ph7.GetResponse())
		t.Log(ph8.GetResponse())
		t.Log(ph9.GetResponse())
	})
	t.Run("Required fields without naming method and 10 videos (2xstorage 8xtelegram)", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vd := msg.NewVideo()
		vd1 := msg.NewVideo()
		vd2 := msg.NewVideo()
		vd3 := msg.NewVideo()
		vd4 := msg.NewVideo()
		vd5 := msg.NewVideo()
		vd6 := msg.NewVideo()
		vd7 := msg.NewVideo()
		vd8 := msg.NewVideo()
		vd9 := msg.NewVideo()

		vd.WriteVideoStorage("media_test/musk.mp4")
		vd1.WriteVideoStorage("media_test/musk.mp4")
		vd2.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd3.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd4.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd5.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd6.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd7.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd8.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")
		vd9.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE")

		msg.AddVideo(vd)
		msg.AddVideo(vd1)
		msg.AddVideo(vd2)
		msg.AddVideo(vd3)
		msg.AddVideo(vd4)
		msg.AddVideo(vd5)
		msg.AddVideo(vd6)
		msg.AddVideo(vd7)
		msg.AddVideo(vd8)
		msg.AddVideo(vd9)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(vd.GetResponse())
		t.Log(vd1.GetResponse())
		t.Log(vd2.GetResponse())
		t.Log(vd3.GetResponse())
		t.Log(vd4.GetResponse())
		t.Log(vd5.GetResponse())
		t.Log(vd6.GetResponse())
		t.Log(vd7.GetResponse())
		t.Log(vd8.GetResponse())
		t.Log(vd9.GetResponse())
	})
}

func TestSendAnimation(t *testing.T) {
	t.Run("Required fields without naming method and animation from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		an := msg.NewAnimation()
		an.WriteAnimationStorage("media_test/prichinatryski.mp4")

		msg.AddAnimation(an)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(an.GetResponse())
	})
	t.Run("Required fields with naming method and animation from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		an := msg.NewAnimation()
		an.WriteAnimationStorage("media_test/prichinatryski.mp4")

		msg.AddAnimation(an)
		msg.AddMethod(methods.Animation)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(an.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and animation from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		an := msg.NewAnimation()

		an.WriteAnimationTelegram("CgACAgIAAxkDAAIDK2co1j66v-5uxwohfxzKv0MhLZntAAL1WQAC43ZJSRKRpvz2ZWV_NgQ")
		an.WriteDuration(3)
		an.WriteHasSpoiler()
		an.WriteHeight(16)
		an.WriteWidth(88)

		inf := msg.NewMessage()
		inf.WriteString("<b>SELAM</b> ALEYKUM!")
		inf.WriteParseMode(types.HTML)
		inf.WriteShowCaptionAboveMedia()
		inf.WriteDisableNotification()
		inf.WriteProtectContent()
		inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

		createInlineKeyboard(msg, t)

		msg.AddMessage(inf)
		msg.AddAnimation(an)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(an.GetResponse())
	})
}

func TestSendVoice(t *testing.T) {
	t.Run("Required fields without naming method and Voice from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		an := msg.NewVoice()
		an.WriteVoiceStorage("media_test/dimaJOSKAproNATO.ogg")

		msg.AddVoice(an)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(an.GetResponse())
	})
	t.Run("Required fields with naming method and Voice from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		an := msg.NewVoice()
		an.WriteVoiceStorage("media_test/dimaJOSKAproNATO.ogg")

		msg.AddVoice(an)
		msg.AddMethod(methods.Voice)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(an.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and Voice from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vc := msg.NewVoice()

		vc.WriteVoiceTelegram("AwACAgIAAxkDAAIDO2cpOUinskD9C8TqK2VytX6vThYDAAL7WwAC43ZRSU1OLw_dlkGCNgQ")
		f := make([]*types.MessageEntity, 1)
		f[0] = &types.MessageEntity{Type: "/start@jobs_bot"}
		vc.WriteDuration(3)

		inf := msg.NewMessage()
		inf.WriteString("SHALOM! <b>Y'ALL</b>")
		inf.WriteParseMode(types.HTML)
		inf.WriteDisableNotification()
		inf.WriteProtectContent()
		inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

		createInlineKeyboard(msg, t)

		msg.AddMessage(inf)
		msg.AddVoice(vc)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(vc.GetResponse())
	})
}

func TestSendVideoNote(t *testing.T) {
	t.Run("Required fields without naming method and voice-note from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vdn := msg.NewVideoNote()
		vdn.WriteVideoNoteStorage("media_test/black.mp4")

		msg.AddVideoNote(vdn)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(vdn.GetResponse())
	})
	t.Run("Required fields with naming method and voice-note from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vdn := msg.NewVideoNote()
		vdn.WriteVideoNoteStorage("media_test/black.mp4")

		msg.AddVideoNote(vdn)

		msg.AddMethod(methods.VideoNote)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(vdn.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and voice-note from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := createChat(msg)

		vdn := msg.NewVideoNote()

		vdn.WriteVideoNoteTelegram("DQACAgIAAxkDAAIDRmcpRSgQWnun_irRK7THUS51PuBEAAJeXAAC43ZRSUQtjaW6gJjyNgQ")
		vdn.WriteDuration(3)
		vdn.WriteLength(8)

		inf := msg.NewMessage()
		inf.WriteString("<b>SELAM</b> ALEYKUM!")
		inf.WriteParseMode(types.HTML)
		inf.WriteDisableNotification()
		inf.WriteProtectContent()
		inf.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596})

		createInlineKeyboard(msg, t)

		msg.AddMessage(inf)
		msg.AddVideoNote(vdn)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(vdn.GetResponse())
	})
}

func init() {
	types.BotID = testbotdata.Token
}
