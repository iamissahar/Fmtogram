package formatter_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func sendRequest(method, url, contentType string, buf []byte) []byte {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", contentType)
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
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

		chat := msg.NewChat()
		chat.WriteChatID(738070596)

		msg.AddMessage(inf)
		msg.AddChat(chat)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(chat.GetResponse())
	})
	t.Run("IMSGInformation IChat IInline", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()
		inf := msg.NewMessage()
		inf.WriteString("<b>SHALOM</b> TO Y'<b>ALL</b>")
		inf.WriteParseMode(types.HTML)

		chat := msg.NewChat()
		chat.WriteChatID(738070596)

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

		msg.AddMessage(inf)
		msg.AddChat(chat)
		msg.AddInlineKeyboard(kb)

		_, err = msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(chat.GetResponse())
	})
}

func TestForwardMessage(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(501)

		msg.AddMethod(methods.ForwardMessage)

		msg.AddMessage(inf)
		msg.AddChat(ch)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(501)
		inf.WriteDisableNotification()
		inf.WriteProtectContent()

		msg.AddMethod(methods.ForwardMessage)

		msg.AddMessage(inf)
		msg.AddChat(ch)

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

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})

		msg.AddMethod(methods.ForwardMessages)

		msg.AddMessage(inf)
		msg.AddChat(ch)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetMessageIDs())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})
		inf.WriteDisableNotification()
		inf.WriteProtectContent()

		msg.AddMethod(methods.ForwardMessages)

		msg.AddMessage(inf)
		msg.AddChat(ch)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetMessageIDs())
	})
}

func TestCopyMessage(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageID(507)

		msg.AddMethod(methods.CopyMessage)

		msg.AddMessage(inf)
		msg.AddChat(ch)

		_, err := msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetMessageIDs())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
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

		msg.AddMethod(methods.CopyMessage)

		msg.AddMessage(inf)
		msg.AddChat(ch)
		msg.AddInlineKeyboard(kb)

		_, err = msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(inf.GetMessageIDs())
	})
}

func TestCopyMessages(t *testing.T) {
	t.Run("Required fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})

		msg.AddMethod(methods.CopyMessages)
		msg.AddChat(ch)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(inf.GetMessageIDs())
	})
	t.Run("Required and unrequired fields", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)
		ch.WriteFromChatID(1051812255)

		inf := msg.NewMessage()
		inf.WriteMessageIDs([]int{501, 507, 508, 509})
		inf.WriteDisableNotification()
		inf.WriteProtectContent()

		msg.AddMethod(methods.CopyMessages)
		msg.AddChat(ch)
		msg.AddMessage(inf)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(inf.GetMessageIDs())
	})
}

func TestSendPhoto(t *testing.T) {
	t.Run("Required fields, without naming method and photo from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)

		ph := msg.NewPhoto()
		ph.WritePhotoStorage("media_test/photo1.jpg")

		msg.AddPhoto(ph)
		msg.AddChat(ch)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(ph.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("Required fields with naming method and photo from storage", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)

		ph := msg.NewPhoto()
		ph.WritePhotoStorage("media_test/photo1.jpg")

		msg.AddMethod(methods.Photo)
		msg.AddPhoto(ph)
		msg.AddChat(ch)

		_, err := msg.Send()
		if err != nil {
			panic(err)
		}

		t.Log(ph.GetResponse())
		t.Log(ch.GetResponse())
	})
	t.Run("Required and unrequired fields without naming method and photo from telegram", func(t *testing.T) {
		msg := formatter.CreateEmpltyMessage()

		ch := msg.NewChat()
		ch.WriteChatID(738070596)

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

		msg.AddChat(ch)
		msg.AddMessage(inf)
		msg.AddPhoto(ph)
		msg.AddInlineKeyboard(kb)

		_, err = msg.Send()
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

		ch := msg.NewChat()
		ch.WriteChatID(738070596)

		ad := msg.NewAudio()
		ad.WriteAudioStorage("media_test/sound.mp3")

		msg.AddChat(ch)
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

		ch := msg.NewChat()
		ch.WriteChatID(738070596)

		ad := msg.NewAudio()
		ad.WriteAudioStorage("media_test/sound.mp3")

		msg.AddMethod(methods.Audio)
		msg.AddChat(ch)
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

		ch := msg.NewChat()
		ch.WriteChatID(738070596)

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

		msg.AddChat(ch)
		msg.AddMessage(inf)
		msg.AddAudio(ad)
		msg.AddInlineKeyboard(kb)

		_, err = msg.Send()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(ch.GetResponse())
		t.Log(inf.GetResponse())
		t.Log(ad.GetResponse())
	})
}

func init() {
	types.BotID = testbotdata.Token
}
