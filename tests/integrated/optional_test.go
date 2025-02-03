package integrated

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

func send(msg *formatter.Message, t *testing.T) {
	if _, err := msg.Send(); err != nil {
		t.Log(err)
	}
}

func (tc *testcase) txtAndPrm(gotext func(string) error, goentities func([]*types.MessageEntity) error, do bool, t *testing.T) {
	var err error
	if err = gotext("There's a text"); err != nil {
		t.Fatal(err)
	}
	if do {
		if err = goentities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
			t.Fatal(err)
		}
	}
}

func (tc *testcase) begining(t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteBusinessConnectionID("BusinessConnectionID"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageThreadID(1111); err != nil {
		t.Fatal(err)
	}
}

func addKb(msg *formatter.Message, kb formatter.IKeyboard, t *testing.T) {
	if err := msg.AddKeyboard(kb); err != nil {
		t.Fatal(err)
	}
}

func (tc *testcase) ending(t *testing.T) {
	var err error
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	// if err = tc.prm.WriteAllowPaidBroadcast(); err != nil {
	// 	t.Fatal(err)
	// }
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func addString(rpbtn formatter.IReplyButton, inbtn formatter.IInlineButton, what int, t *testing.T) {
	var err error
	if what == 0 {
		if err = rpbtn.WriteString("????"); err != nil {
			t.Fatal(err)
		}
	} else if what == 1 {
		if err = inbtn.WriteString("????"); err != nil {
			t.Fatal(err)
		}
	}
}

func setRpKb(set []int, btns []formatter.IReplyButton, kb formatter.IKeyboard, t *testing.T) {
	rp, err := kb.WriteReply()
	if err != nil {
		t.Fatal(err)
	}
	if err = rp.Set(set); err != nil {
		t.Fatal(err)
	}
	if err = rp.WriteInputFieldPlaceholder("placeholder"); err != nil {
		t.Fatal(err)
	}
	if err = rp.WriteIsPersistent(); err != nil {
		t.Fatal(err)
	}
	if err = rp.WriteOneTimeKeyboard(); err != nil {
		t.Fatal(err)
	}
	if err = rp.WriteRemove(); err != nil {
		t.Fatal(err)
	}
	if err = rp.WriteSelective(); err != nil {
		t.Fatal(err)
	}
	for i, j := 0, 0; i < len(set); i++ {
		for k := 0; k < set[i]; k++ {
			if btns[j], err = rp.NewButton(i, k); err != nil {
				t.Fatal(err)
			}
			j++
		}
	}
}

func createRpKBAndData(btns []formatter.IReplyButton) []interface{} {
	return []interface{}{btns[0].WriteRequestChat, btns[1].WriteRequestContact, btns[2].WriteRequestLocation, btns[3].WriteRequestPoll,
		btns[4].WriteRequestUsers, btns[5].WriteWebApp}
}

func (*testcase) replyKb(tc *testcase, t *testing.T) {
	var btn, btn1, btn2, btn3, btn4, btn5 formatter.IReplyButton
	var err error
	kb := tc.msg.NewKeyboard()
	btns := []formatter.IReplyButton{btn, btn1, btn2, btn3, btn4, btn5}
	setRpKb([]int{2, 4}, btns, kb, t)
	functions := createRpKBAndData(btns)
	for i, f := range functions {
		addString(btns[i], nil, 0, t)
		switch ff := f.(type) {
		case func() error:
			if err = ff(); err != nil {
				t.Fatal(err)
			}
		case func(*types.KeyboardButtonRequestChat) error:
			if err = ff(&types.KeyboardButtonRequestChat{RequestID: 738070596}); err != nil {
				t.Fatal(err)
			}
		case func(*types.KeyboardButtonRequestUsers) error:
			if err = ff(&types.KeyboardButtonRequestUsers{}); err != nil {
				t.Fatal(err)
			}
		case func(*types.KeyboardButtonPollType) error:
			if err = ff(&types.KeyboardButtonPollType{Type: "regular"}); err != nil {
				t.Fatal(err)
			}
		case func(*types.WebAppInfo) error:
			if err = ff(&types.WebAppInfo{Url: "https://youtube.com"}); err != nil {
				t.Fatal(err)
			}
		}
	}
	addKb(tc.msg, kb, t)
	send(tc.msg, t)
}

// I'm not testing WritePay and WriteCallbackGame
func createInKBAndData(btns []formatter.IInlineButton) ([]interface{}, []string) {
	functions := []interface{}{
		btns[0].WriteCallbackData,
		btns[1].WriteLoginUrl,
		btns[2].WriteSwitchInlineQuery,
		btns[3].WriteSwitchInlineQueryChosenChat,
		btns[4].WriteSwitchInlineQueryCurrentChat,
		btns[5].WriteURL,
		btns[6].WriteWebApp,
	}
	data := []string{"callback data", "", "", "https://youtube.com"}
	return functions, data
}

func setInKb(set []int, btns []formatter.IInlineButton, kb formatter.IKeyboard, t *testing.T) {
	in, err := kb.WriteInline()
	if err != nil {
		t.Fatal(err)
	}
	if err = in.Set(set); err != nil {
		t.Fatal(err)
	}
	for i, j := 0, 0; i < len(set); i++ {
		for k := 0; k < set[i]; k++ {
			if btns[j], err = in.NewButton(i, k); err != nil {
				t.Fatal(err)
			}
			j++
		}
	}
}

func (*testcase) inlineKb(tc *testcase, t *testing.T) {
	var btn, btn1, btn2, btn3, btn4, btn5, btn6 formatter.IInlineButton
	var err error
	kb := tc.msg.NewKeyboard()
	btns := []formatter.IInlineButton{btn, btn1, btn2, btn3, btn4, btn5, btn6}
	set := []int{2, 4, 1}
	setInKb(set, btns, kb, t)
	functions, data := createInKBAndData(btns)
	j := 0
	for i, f := range functions {
		addString(nil, btns[i], 1, t)
		switch ff := f.(type) {
		case func(string) error:
			if err = ff(data[j]); err != nil {
				t.Fatal(err)
			}
			j++
		case func(*types.CallbackGame) error:
			if err = ff(&types.CallbackGame{}); err != nil {
				t.Fatal(err)
			}
		case func(*types.LoginUrl) error:
			if err = ff(&types.LoginUrl{Url: "https://youtube.com", ForwardText: "Yep, that's me"}); err != nil {
				t.Fatal(err)
			}
		case func(*types.SwitchInlineQueryChosenChat) error:
			if err = ff(&types.SwitchInlineQueryChosenChat{Query: "query", AllowUserChats: true}); err != nil {
				t.Fatal(err)
			}
		case func(*types.WebAppInfo) error:
			if err = ff(&types.WebAppInfo{Url: "https://youtube.com"}); err != nil {
				t.Fatal(err)
			}
		}
	}
	addKb(tc.msg, kb, t)
	send(tc.msg, t)
}

func (*testcase) forceKb(tc *testcase, t *testing.T) {
	kb := tc.msg.NewKeyboard()
	frp, err := kb.WriteForceReply()
	if err != nil {
		t.Fatal(err)
	}
	if err = frp.WriteForceReply(); err != nil {
		t.Fatal(err)
	}
	if err = frp.WriteInputFieldPlaceholder("placeholder"); err != nil {
		t.Fatal(err)
	}
	if err = frp.WriteSelective(); err != nil {
		t.Fatal(err)
	}
	addKb(tc.msg, kb, t)
	send(tc.msg, t)
}

// func (tc *testcase) ending(t *testing.T) {
// 	if err := tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if !tc.withoutEnding {
// 		withoutKB(tc.prm, t)
// 		if !tc.withoutKB {
// 			if tc.onlyInline {
// 				withInlineKB(tc.msg, t)
// 			} else {
// 				withReplyKB(tc.msg, t)
// 				withInlineKB(tc.msg, t)
// 				withForceReplyKB(tc.msg, t)
// 			}
// 		}
// 	} else {
// 		send(tc.msg, t)
// 	}
// }
