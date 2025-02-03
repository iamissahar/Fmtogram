package integrated

import (
	"testing"
	"time"

	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func (tc *testcase) init() {
	tc.msg = formatter.CreateEmpltyMessage()
	tc.prm = tc.msg.NewParameters()
	tc.ch = tc.msg.NewChat()
	tc.ph = tc.msg.NewPhoto()
	tc.ad = tc.msg.NewAudio()
	tc.dc = tc.msg.NewDocument()
	tc.vd = tc.msg.NewVideo()
	tc.an = tc.msg.NewAnimation()
	tc.vc = tc.msg.NewVoice()
	tc.vdn = tc.msg.NewVideoNote()
	tc.get = tc.msg.NewIGet()
	types.DEBUG = true
	logs.DebugOrInfo()
}

func sendMessageCommon(i int, t *testing.T) *testcase {
	tc := new(testcase)
	tc.init()
	tc.begining(t)
	tc.msg.AddChat(tc.ch)
	return tc
}

func sendMessageReq(t *testing.T) {
	tc := sendMessageCommon(0, t)
	if err := tc.prm.WriteString("There's a text"); err != nil {
		t.Fatal(err)
	}
	if err := tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func sendMessageCommonForKb(t *testing.T, do bool) *testcase {
	tc := sendMessageCommon(1, t)
	tc.prm.WriteLinkPreviewOptions(&types.LinkPreviewOptions{IsDisabled: false, URL: "https://youtube.com"})
	tc.txtAndPrm(tc.prm.WriteString, tc.prm.WriteEntities, do, t)
	return tc
}

func sendMessageReplyKb(t *testing.T) {
	var tc *testcase
	for i := 0; i < 4; i++ {
		if i == 3 {
			tc = sendMessageCommonForKb(t, true)
		} else {
			tc = sendMessageCommonForKb(t, false)
			tc.prm.WriteParseMode(parsemode[i])
		}
		tc.ending(t)
		tc.replyKb(tc, t)
		if code, msg := tc.get.Error(); code != 0 {
			t.Fatal(msg)
		}
	}
}

func sendMessageInlineKb(t *testing.T) {
	var tc *testcase
	for i := 0; i < 4; i++ {
		if i == 3 {
			tc = sendMessageCommonForKb(t, true)
		} else {
			tc = sendMessageCommonForKb(t, false)
			tc.prm.WriteParseMode(parsemode[i])
		}
		tc.ending(t)
		tc.inlineKb(tc, t)
		if code, msg := tc.get.Error(); code != 0 {
			t.Fatal(msg)
		}
	}
}

func sendMessageForceKb(t *testing.T) {
	var tc *testcase
	for i := 0; i < 4; i++ {
		if i == 3 {
			tc = sendMessageCommonForKb(t, true)
		} else {
			tc = sendMessageCommonForKb(t, false)
			tc.prm.WriteParseMode(parsemode[i])
		}
		tc.ending(t)
		tc.forceKb(tc, t)
		if code, msg := tc.get.Error(); code != 0 {
			t.Fatal(msg)
		}
	}
}

func sendMessageAll(t *testing.T) {
	t.Run("ReplyKB", sendMessageReplyKb)
	t.Run("InlineKB", sendMessageInlineKb)
	t.Run("ForceKB", sendMessageForceKb)
}

func TestSendMessage(t *testing.T) {
	t.Run("Required", sendMessageReq)
	t.Run("AllParameters", sendMessageAll)
}

func fmsgReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageID(3966); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.ForwardMessage); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func fmsgAllPrm(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	// if err = tc.prm.WriteMessageThreadID(1234); err != nil {
	// 	t.Fatal(err)
	// }
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageID(3966); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.ForwardMessage); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func forwardMessage(t *testing.T) {
	t.Run("Required", fmsgReq)
	t.Run("AllParameters", fmsgAllPrm)
}

func fmsgsReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageIDs([]int{4107, 4108}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.ForwardMessages); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func fmsgsAllPrm(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageIDs([]int{4107, 4108}); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.ForwardMessages); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func forwardMessages(t *testing.T) {
	t.Run("Required", fmsgsReq)
	t.Run("AllParameters", fmsgsAllPrm)
}

func TestForward(t *testing.T) {
	t.Run("1Message", forwardMessage)
	t.Run("LotsMessages", forwardMessages)
}

func copyMsgReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageID(4108); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.CopyMessage); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func commonCopyMessage(tc *testcase, t *testing.T) {
	var err error
	if err = tc.prm.WriteString("There is some kind of text"); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageID(4108); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.CopyMessage); err != nil {
		t.Fatal(err)
	}
}

func copyMsgInKb(t *testing.T) {
	var err error
	for i := 0; i < 4; i++ {
		tc := new(testcase)
		tc.init()
		if i == 3 {
			if err = tc.prm.WriteEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
				t.Fatal(err)
			}
		} else {
			if err = tc.prm.WriteParseMode(parsemode[i]); err != nil {
				t.Fatal(err)
			}
		}
		commonCopyMessage(tc, t)
		tc.inlineKb(tc, t)
		send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 {
			t.Fatal(msg)
		}
	}
}

func copyMsgReplyKb(t *testing.T) {
	var err error
	for i := 0; i < 4; i++ {
		tc := new(testcase)
		tc.init()
		if i == 3 {
			if err = tc.prm.WriteEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
				t.Fatal(err)
			}
		} else {
			if err = tc.prm.WriteParseMode(parsemode[i]); err != nil {
				t.Fatal(err)
			}
		}
		commonCopyMessage(tc, t)
		tc.replyKb(tc, t)
		send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 {
			t.Fatal(msg)
		}
	}
}

func copyMsgForceKb(t *testing.T) {
	var err error
	for i := 0; i < 4; i++ {
		tc := new(testcase)
		tc.init()
		if i == 3 {
			if err = tc.prm.WriteEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
				t.Fatal(err)
			}
		} else {
			if err = tc.prm.WriteParseMode(parsemode[i]); err != nil {
				t.Fatal(err)
			}
		}
		commonCopyMessage(tc, t)
		tc.forceKb(tc, t)
		send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 {
			t.Fatal(msg)
		}
	}
}

func copyMsgAll(t *testing.T) {
	t.Run("InKb", copyMsgInKb)
	t.Run("ReplyKb", copyMsgReplyKb)
	t.Run("ForceKb", copyMsgForceKb)
}

func copyMessage(t *testing.T) {
	t.Run("Required", copyMsgReq)
	t.Run("All", copyMsgAll)
}

func copyMsgsReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageIDs([]int{4105, 4108}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.CopyMessages); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func copyMsgsAll(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteFromChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageIDs([]int{4105, 4108}); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteRemoveCaption(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.CopyMessages); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 {
		t.Fatal(msg)
	}
}

func copyMessages(t *testing.T) {
	t.Run("Required", copyMsgsReq)
	t.Run("All", copyMsgsAll)
}

func TestCopy(t *testing.T) {
	t.Run("1Message", copyMessage)
	t.Run("LotsMessages", copyMessages)
}
func TestMe(t *testing.T) {
	var offset int
	var err error
	if err = executer.GetOffset(&offset, types.BotID); err != nil {
		t.Fatal(err)
	}
	for err != nil {
		err = executer.GetOffset(&offset, types.BotID)
		time.Sleep(time.Second / 10)
	}
	tg := new(types.Telegram)
	if err = executer.GetUpdates(tg, &offset, types.BotID); err != nil {
		t.Fatal(err)
	}
}

func init() {
	types.BotID = testbotdata.Token
}
