package integrated

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

func addString(rpbtn fmtogram.IReplyButton, inbtn fmtogram.IInlineButton, what int, t *testing.T) {
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

func setRpKb(set []int, btns []fmtogram.IReplyButton, kb fmtogram.IKeyboard, t *testing.T) {
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

func createRpKBAndData(btns []fmtogram.IReplyButton) []interface{} {
	return []interface{}{btns[0].WriteRequestChat, btns[1].WriteRequestContact, btns[2].WriteRequestLocation, btns[3].WriteRequestPoll,
		btns[4].WriteRequestUsers, btns[5].WriteWebApp}
}

func replyKb(tc *testcase, t *testing.T) {
	var btn, btn1, btn2, btn3, btn4, btn5 fmtogram.IReplyButton
	var err error
	kb := fmtogram.NewKeyboard()
	btns := []fmtogram.IReplyButton{btn, btn1, btn2, btn3, btn4, btn5}
	setRpKb([]int{2, 4}, btns, kb, t)
	functions := createRpKBAndData(btns)
	for i, f := range functions {
		addString(btns[i], nil, 0, t)
		switch ff := f.(type) {
		case func() error:
			if err = ff(); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.KeyboardButtonRequestChat) error:
			if err = ff(&fmtogram.KeyboardButtonRequestChat{RequestID: 738070596}); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.KeyboardButtonRequestUsers) error:
			if err = ff(&fmtogram.KeyboardButtonRequestUsers{}); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.KeyboardButtonPollType) error:
			if err = ff(&fmtogram.KeyboardButtonPollType{Type: "regular"}); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.WebAppInfo) error:
			if err = ff(&fmtogram.WebAppInfo{Url: "https://youtube.com"}); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err = tc.msg.WriteKeyboard(kb); err != nil {
		t.Fatal(err)
	}
}

// I'm not testing WritePay and WriteCallbackGame
func createInKBAndData(btns []fmtogram.IInlineButton) ([]interface{}, []string) {
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

func setInKb(set []int, btns []fmtogram.IInlineButton, kb fmtogram.IKeyboard, t *testing.T) {
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

func inlineKb(tc *testcase, t *testing.T) {
	var btn, btn1, btn2, btn3, btn4, btn5, btn6 fmtogram.IInlineButton
	var err error
	kb := fmtogram.NewKeyboard()
	btns := []fmtogram.IInlineButton{btn, btn1, btn2, btn3, btn4, btn5, btn6}
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
		case func(*fmtogram.CallbackGame) error:
			if err = ff(&fmtogram.CallbackGame{}); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.LoginUrl) error:
			if err = ff(&fmtogram.LoginUrl{Url: "https://youtube.com", ForwardText: "Yep, that's me"}); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.SwitchInlineQueryChosenChat) error:
			if err = ff(&fmtogram.SwitchInlineQueryChosenChat{Query: "query", AllowUserChats: true}); err != nil {
				t.Fatal(err)
			}
		case func(*fmtogram.WebAppInfo) error:
			if err = ff(&fmtogram.WebAppInfo{Url: "https://youtube.com"}); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err = tc.msg.WriteKeyboard(kb); err != nil {
		t.Fatal(err)
	}
}

func forceKb(tc *testcase, t *testing.T) {
	kb := fmtogram.NewKeyboard()
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
	if err = tc.msg.WriteKeyboard(kb); err != nil {
		t.Fatal(err)
	}
}
