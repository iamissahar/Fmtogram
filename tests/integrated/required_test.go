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

var parsemode = []string{types.HTML, types.Markdown, types.MarkdownV2}

type testcase struct {
	msg              *formatter.Message
	prm              formatter.IParameters
	ch               formatter.IChat
	ph               formatter.IPhoto
	vd               formatter.IVideo
	an               formatter.IAnimation
	get              formatter.IGet
	withoutKB        bool
	onlyInline       bool
	withoutBCon      bool
	withoutThreads   bool
	withoutParseMode bool
	withoutSettings  bool
	withoutEnding    bool
}

func (tc *testcase) init() {
	tc.msg = formatter.CreateEmpltyMessage()
	tc.prm = tc.msg.NewParameters()
	tc.ch = tc.msg.NewChat()
	tc.ph = tc.msg.NewPhoto()
	tc.vd = tc.msg.NewVideo()
	tc.an = tc.msg.NewAnimation()
	tc.get = tc.msg.NewIGet()
	types.DEBUG = true
	logs.DebugOrInfo()
}

func initBools(tc *testcase, i int) {
	matrix := [][]bool{{true, false, false}, {false, false, false}, {true, true, true},
		{true, true, true}, {true, false, true}, {true, false, false}, {true, false, false}}

	tc.withoutKB = matrix[0][i]
	tc.onlyInline = matrix[1][i]
	tc.withoutBCon = matrix[2][i]
	tc.withoutThreads = matrix[3][i]
	tc.withoutParseMode = matrix[4][i]
	tc.withoutSettings = matrix[5][i]
	tc.withoutEnding = matrix[6][i]
}

func sendMessageCommon(i int, t *testing.T) *testcase {
	tc := new(testcase)
	tc.init()
	initBools(tc, i)
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
		tc.replyKb(t)
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
		tc.inlineKb(t)
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
		tc.forceKb(t)
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
}

func forwardMessage(t *testing.T) {
	t.Run("Required", fmsgReq)
	t.Run("AllParameters", fmsgAllPrm)
}

func forwardMessages(t *testing.T) {

}

func TestForward(t *testing.T) {
	t.Run("Message", forwardMessage)
	t.Run("Messages", forwardMessages)
}

func TestMe(t *testing.T) {
	var offset int
	var err error
	types.BotID = testbotdata.Token
	if err = executer.GetOffset(&offset); err != nil {
		t.Fatal(err)
	}
	for err != nil {
		err = executer.GetOffset(&offset)
		time.Sleep(time.Second / 10)
	}
	tg := new(types.Telegram)
	if err = executer.GetUpdates(tg, &offset); err != nil {
		t.Fatal(err)
	}
}
