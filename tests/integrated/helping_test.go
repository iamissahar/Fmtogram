package integrated

import (
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
	"github.com/iamissahar/Fmtogram/types"
)

func send(msg *formatter.Message, t *testing.T) {
	var err error
	if err = msg.Send(); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
}

func addKb(msg *formatter.Message, kb formatter.IKeyboard) {
	if err := msg.AddKeyboard(kb); err != nil {
		panic(err)
	}
}

func addString(rpbtn formatter.IReplyButton, inbtn formatter.IInlineButton, what int) {
	var err error
	if what == 0 {
		if err = rpbtn.WriteString("????"); err != nil {
			panic(err)
		}
	} else if what == 1 {
		if err = inbtn.WriteString("????"); err != nil {
			panic(err)
		}
	}
}

func setRpKb(set []int, btns []formatter.IReplyButton, kb formatter.IKeyboard) {
	rp, err := kb.WriteReply()
	if err != nil {
		panic(err)
	}
	if err = rp.Set(set); err != nil {
		panic(err)
	}
	if err = rp.WriteInputFieldPlaceholder("placeholder"); err != nil {
		panic(err)
	}
	if err = rp.WriteIsPersistent(); err != nil {
		panic(err)
	}
	if err = rp.WriteOneTimeKeyboard(); err != nil {
		panic(err)
	}
	if err = rp.WriteRemove(); err != nil {
		panic(err)
	}
	if err = rp.WriteSelective(); err != nil {
		panic(err)
	}
	for i, j := 0, 0; i < len(set); i++ {
		for k := 0; k < set[i]; k++ {
			if btns[j], err = rp.NewButton(i, k); err != nil {
				panic(err)
			}
			j++
		}
	}
}

func createRpKBAndData(btns []formatter.IReplyButton) []interface{} {
	return []interface{}{btns[0].WriteRequestChat, btns[1].WriteRequestContact, btns[2].WriteRequestLocation, btns[3].WriteRequestPoll,
		btns[4].WriteRequestUsers, btns[5].WriteWebApp}
}

func replyKb(tc *testcase) {
	var btn, btn1, btn2, btn3, btn4, btn5 formatter.IReplyButton
	var err error
	kb := tc.msg.NewKeyboard()
	btns := []formatter.IReplyButton{btn, btn1, btn2, btn3, btn4, btn5}
	setRpKb([]int{2, 4}, btns, kb)
	functions := createRpKBAndData(btns)
	for i, f := range functions {
		addString(btns[i], nil, 0)
		switch ff := f.(type) {
		case func() error:
			if err = ff(); err != nil {
				panic(err)
			}
		case func(*types.KeyboardButtonRequestChat) error:
			if err = ff(&types.KeyboardButtonRequestChat{RequestID: 738070596}); err != nil {
				panic(err)
			}
		case func(*types.KeyboardButtonRequestUsers) error:
			if err = ff(&types.KeyboardButtonRequestUsers{}); err != nil {
				panic(err)
			}
		case func(*types.KeyboardButtonPollType) error:
			if err = ff(&types.KeyboardButtonPollType{Type: "regular"}); err != nil {
				panic(err)
			}
		case func(*types.WebAppInfo) error:
			if err = ff(&types.WebAppInfo{Url: "https://youtube.com"}); err != nil {
				panic(err)
			}
		}
	}
	addKb(tc.msg, kb)
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

func setInKb(set []int, btns []formatter.IInlineButton, kb formatter.IKeyboard) {
	in, err := kb.WriteInline()
	if err != nil {
		panic(err)
	}
	if err = in.Set(set); err != nil {
		panic(err)
	}
	for i, j := 0, 0; i < len(set); i++ {
		for k := 0; k < set[i]; k++ {
			if btns[j], err = in.NewButton(i, k); err != nil {
				panic(err)
			}
			j++
		}
	}
}

func inlineKb(tc *testcase) {
	var btn, btn1, btn2, btn3, btn4, btn5, btn6 formatter.IInlineButton
	var err error
	kb := tc.msg.NewKeyboard()
	btns := []formatter.IInlineButton{btn, btn1, btn2, btn3, btn4, btn5, btn6}
	set := []int{2, 4, 1}
	setInKb(set, btns, kb)
	functions, data := createInKBAndData(btns)
	j := 0
	for i, f := range functions {
		addString(nil, btns[i], 1)
		switch ff := f.(type) {
		case func(string) error:
			if err = ff(data[j]); err != nil {
				panic(err)
			}
			j++
		case func(*types.CallbackGame) error:
			if err = ff(&types.CallbackGame{}); err != nil {
				panic(err)
			}
		case func(*types.LoginUrl) error:
			if err = ff(&types.LoginUrl{Url: "https://youtube.com", ForwardText: "Yep, that's me"}); err != nil {
				panic(err)
			}
		case func(*types.SwitchInlineQueryChosenChat) error:
			if err = ff(&types.SwitchInlineQueryChosenChat{Query: "query", AllowUserChats: true}); err != nil {
				panic(err)
			}
		case func(*types.WebAppInfo) error:
			if err = ff(&types.WebAppInfo{Url: "https://youtube.com"}); err != nil {
				panic(err)
			}
		}
	}
	addKb(tc.msg, kb)
}

func forceKb(tc *testcase) {
	kb := tc.msg.NewKeyboard()
	frp, err := kb.WriteForceReply()
	if err != nil {
		panic(err)
	}
	if err = frp.WriteForceReply(); err != nil {
		panic(err)
	}
	if err = frp.WriteInputFieldPlaceholder("placeholder"); err != nil {
		panic(err)
	}
	if err = frp.WriteSelective(); err != nil {
		panic(err)
	}
	addKb(tc.msg, kb)
}
