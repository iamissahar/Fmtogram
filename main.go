package fmtogram

import (
	"fmt"
	"time"
)

type BasicSettings struct {
	// Required. The function that is going to be called every time a bot takes some updates
	StartFunc func(upd *Update, bs *BasicSettings)

	// Required. The token of a bot
	Token string

	// Optionl. By default any kind of updates are recieved
	AllowedUpdates []string

	// Optional. By default the value of this field is 7 seconds.
	// Change it only in case you want shorted pulling or longer. Might be 0
	Timeout *int

	// Optional. By default the value of this field is 10. Discribes the limit of pulling updates
	Limit *int

	// Optional, but recomended. If error appreas during the proccess of getting updates, it's going to be sent in this channel.
	// Note: if the channel is nil the proccess is going to panic
	ErrorHandler chan error
}

// func queue(reg *regTable, tg *Telegram, chatID int, index *int) {
// 	if *index != none {
// 		logs.AlreadyInQueue()
// 		reg.Reg[*index].Chu <- tg
// 	} else {
// 		logs.FirstTimeUser()
// 		*index = reg.newIndex()
// 		reg.Reg[*index].UserId = chatID
// 		reg.Reg[*index].Chu = make(chan *Telegram, 1)
// 		reg.Reg[*index].Chu <- tg
// 	}
// }

// func getChatID(upd *Update, bs *BasicSettings) int {
// 	var err error
// 	chatid := 0
// 	if upd != nil {
// 		if upd.Msg != nil && upd.Msg.Ch != nil {
// 			chatid = upd.Msg.Ch.ID
// 		} else if upd.EditedMsg != nil && upd.EditedMsg.Ch != nil {
// 			chatid = upd.EditedMsg.Ch.ID
// 		} else if upd.ChanPost != nil && upd.ChanPost.Ch != nil {
// 			chatid = upd.ChanPost.Ch.ID
// 		} else if upd.EditedChanPost != nil && upd.EditedChanPost.Ch != nil {
// 			chatid = upd.EditedChanPost.Ch.ID
// 		} else if upd.BusinessConn != nil && upd.BusinessConn.User != nil {
// 			chatid = upd.BusinessConn.User.ID
// 		} else if upd.BusinessMsg != nil && upd.BusinessMsg.Ch != nil {
// 			chatid = upd.BusinessMsg.Ch.ID
// 		} else if upd.EditedBusinessMsg != nil && upd.EditedBusinessMsg.Ch != nil {
// 			chatid = upd.EditedBusinessMsg.Ch.ID
// 		} else if upd.DeletedBusinessMessage != nil && upd.DeletedBusinessMessage.Ch != nil {
// 			chatid = upd.DeletedBusinessMessage.Ch.ID
// 		} else if upd.MessageR != nil && upd.MessageR.Ch != nil {
// 			chatid = upd.MessageR.Ch.ID
// 		} else if upd.MessageRcount != nil && upd.MessageRcount.Ch != nil {
// 			chatid = upd.MessageRcount.Ch.ID
// 		} else if upd.InlineQ != nil && upd.InlineQ.From != nil {
// 			chatid = upd.InlineQ.From.ID
// 		} else if upd.ChosenInlineR != nil && upd.ChosenInlineR.From != nil {
// 			chatid = upd.ChosenInlineR.From.ID
// 		} else if upd.CallbackQ != nil && upd.CallbackQ.From != nil {
// 			chatid = upd.CallbackQ.From.ID
// 		} else if upd.ShippingQ != nil && upd.ShippingQ.From != nil {
// 			chatid = upd.ShippingQ.From.ID
// 		} else if upd.PreCheckoutQ != nil && upd.PreCheckoutQ.From != nil {
// 			chatid = upd.PreCheckoutQ.From.ID
// 		} else if upd.PaidMedia != nil && upd.PaidMedia.From != nil {
// 			chatid = upd.PaidMedia.From.ID
// 		} else if upd.PollUpd != nil {
// 			chatid = 0
// 		} else if upd.PollAnswr != nil && upd.PollAnswr.VChat != nil {
// 			chatid = upd.PollAnswr.VChat.ID
// 		} else if upd.MyChatMem != nil && upd.MyChatMem.Ch != nil {
// 			chatid = upd.MyChatMem.Ch.ID
// 		} else if upd.ChatMem != nil && upd.ChatMem.Ch != nil {
// 			chatid = upd.ChatMem.Ch.ID
// 		} else if upd.ChatJoinReq != nil && upd.ChatJoinReq.Ch != nil {
// 			chatid = upd.ChatJoinReq.Ch.ID
// 		} else if upd.ChatBoostUpd != nil && upd.ChatBoostUpd.Ch != nil {
// 			chatid = upd.ChatBoostUpd.Ch.ID
// 		} else if upd.RemovedChatBoostUpd != nil && upd.RemovedChatBoostUpd.Ch != nil {
// 			chatid = upd.RemovedChatBoostUpd.Ch.ID
// 		} else {
// 			err = code52()
// 			if bs.ErrorHandler != nil {
// 				bs.ErrorHandler <- err
// 			} else {
// 				panic(err)
// 			}
// 		}
// 	} else {
// 		err = code53()
// 		if bs.ErrorHandler != nil {
// 			bs.ErrorHandler <- err
// 		} else {
// 			panic(err)
// 		}
// 	}
// 	return chatid
// }

func worker(tg *Telegram, bs *BasicSettings, workdone chan struct{}) {
	for _, upd := range tg.Result {
		fmt.Println(upd)
		fmt.Println(bs)
		fmt.Println(bs.StartFunc)
		bs.StartFunc(upd, bs)
		workdone <- struct{}{}
	}
}

func waiter(tg *Telegram, offset *int, workdone chan struct{}) {
	var i int
	for range workdone {
		i++
		if i == len(tg.Result) {
			*offset = tg.Result[len(tg.Result)-1].UpdateID + 1
			break
		}
	}
	close(workdone)
}

func pullResponse(bs *BasicSettings) {
	var offset, i int
	for offset == 0 {
		getOffset(&offset, bs)
	}
	for {
		tg := new(Telegram)
		getUpdates(tg, &offset, bs)
		workdone := make(chan struct{})
		if (len(tg.Result) != 0) && (len(tg.Result) > i) {
			go worker(tg, bs, workdone)
			go waiter(tg, &offset, workdone)
		}
		time.Sleep(time.Second)
	}
}

// Entry of every bots' work. BasicSetting cannot be nil, otherwise the function panics
func (bs *BasicSettings) Start() {
	if bs != nil {
		go pullResponse(bs)
	} else {
		panic(code50())
	}
}
