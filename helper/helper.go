package helper

import (
	"github.com/iamissahar/Fmtogram/executer"
	"github.com/iamissahar/Fmtogram/types"
)

func IsWorking() (*types.User, error) {
	bot := new(types.User)
	getme, err := executer.GetMe()
	if err == nil {
		bot = getme.Result
	}
	return bot, err
}

func GetUserID(tg *types.Telegram) int {
	var userID int
	if tg.Result[0].Message.From.ID != 0 {
		userID = tg.Result[0].Message.From.ID
	} else if tg.Result[0].CallbackQuery.From.ID != 0 {
		userID = tg.Result[0].CallbackQuery.From.ID
	}
	return userID
}

func GetChatID(tg *types.Telegram) int {
	var chatID int
	if tg.Result[0].Message.Chat.ID != 0 {
		chatID = tg.Result[0].Message.Chat.ID
	}
	return chatID
}
