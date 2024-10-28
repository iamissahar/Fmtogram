package fmtogram

import (
	"fmt"
	"time"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/helper"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

var StartFunc func(*types.Telegram, *types.Telegram, *formatter.Message)

func firstStep(offset *int) {
	logs.GetOffset()

	err := executer.GetOffset(offset)
	for err != nil {
		err = executer.GetOffset(offset)
		time.Sleep(time.Second / 10)
	}

	logs.GottenOffset()
}

func queue(reg *executer.RegTable, tg *types.Telegram, chatID int, index *int) {
	if *index != executer.None {
		logs.AlreadyInQueue()

		reg.Reg[*index].Chu <- tg
	} else {
		logs.FirstTimeUser()

		*index = reg.NewIndex()
		reg.Reg[*index].UserId = chatID
		reg.Reg[*index].Chu = make(chan *types.Telegram, 1)
		reg.Reg[*index].Chu <- tg

		fmt.Println(reg.Reg[*index])
	}
}

func pullResponse(reg *executer.RegTable) {
	var offset int
	firstStep(&offset)
	for {
		tg := new(types.Telegram)
		err := executer.GetUpdates(tg, &offset)
		if len(tg.Result) != 0 && err == nil {
			logs.FoundSomeIformation()

			chatID := helper.GetUserID(tg)
			index := reg.Seeker(chatID)

			queue(reg, tg, chatID, &index)
			go worker(reg.Reg[index].Chu, reg.Reg[index].Chb)

			offset = offset + 1
		} else if err != nil {
			errors.ErrorInUserData(err)
		}
		time.Sleep(time.Second / 10)
	}
}

func worker(input chan *types.Telegram, returned chan *types.Telegram) {
	tg := new(types.Telegram)
	tgReturned := new(types.Telegram)

	for len(input) > 0 {
		tg = <-input
		if len(returned) > 0 {
			logs.ReturnedIsntEmply()
			tgReturned = <-returned
		} else {
			logs.ReturnedIsEmply()
		}
		logs.CallDeveloperFunc()
		mes := formatter.CreateMessage(tg)
		StartFunc(tg, tgReturned, mes)
	}
}

func Start() {
	defer logs.TurnOff()
	logs.DebugOrInfo()
	logs.TurnOn()

	reg := new(executer.RegTable)
	pullResponse(reg)
}
