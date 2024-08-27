package executer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	None        int    = -1
	defaultType string = "application/json"
)

type Entry struct {
	UserId int
	Chu    chan *types.Telegram
	Chb    chan *types.Telegram
}

type RegTable struct {
	Reg []Entry
}

func sendRequest(buf *bytes.Buffer, url, contentType, method string) ([]byte, error) {
	var body []byte
	response := new(http.Response)
	request, err := http.NewRequest(method, url, buf)
	if err == nil {
		request.Header.Set("Content-Type", contentType)
		client := &http.Client{}
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
			body, err = io.ReadAll(response.Body)
		}
	}
	return body, err
}

func Send(buf *bytes.Buffer, function, contenttype string, unmarshal bool) *types.Telegram {
	tg := new(types.Telegram)
	url := fmt.Sprintf("%sbot%s/%s", types.TelegramAPI, types.BotID, function)
	body, err := sendRequest(buf, url, contenttype, "POST")
	if err == nil && unmarshal {
		err = json.Unmarshal(body, tg)
	}
	if err != nil {
		errors.RequestError(err)
	}
	return tg
}

func (reg *RegTable) Seeker(chatID int) (index int) {
	var i int
	index = None
	if len(reg.Reg) != 0 {
		for i < len(reg.Reg) && reg.Reg[i].UserId != chatID {
			i++
		}
		if i < len(reg.Reg) {
			if reg.Reg[i].UserId == chatID {
				index = i
			}
		}
	}
	return index
}

func (reg *RegTable) NewIndex() (newIndex int) {
	newIndex = len(reg.Reg)
	reg.Reg = append(reg.Reg, Entry{})
	return newIndex
}

func GetUpdates(tg *types.Telegram, offset *int) error {
	url := fmt.Sprint(
		types.TelegramAPI, "/",
		fmt.Sprintf("bot%s", types.BotID), "/",
		"getUpdates", "?", "limit=1", fmt.Sprintf("offset=%d", *offset))
	body, err := sendRequest(nil, url, "application/json", "GET")
	if err == nil {
		err = json.Unmarshal(body, tg)
	}
	return err
}

func GetOffset(offset *int) error {
	tg := new(types.Telegram)
	url := fmt.Sprint(
		types.TelegramAPI, "/",
		fmt.Sprintf("bot%s", types.BotID), "/",
		"getUpdates", "?", "limit=1")
	body, err := sendRequest(nil, url, "application/json", "GET")
	if err == nil {

		err = json.Unmarshal(body, tg)
		if err == nil {

			if len(tg.Result) > 0 {
				*offset = tg.Result[0].UpdateID
			} else {
				err = errors.UpdatesMisstakes("Telegram returned an empty data")
			}

		}
	}
	return err
}

func GetMe() (*types.GetMe, error) {
	getme := new(types.GetMe)
	url := fmt.Sprint(
		types.TelegramAPI, "/",
		fmt.Sprintf("bot%s", types.BotID), "/", "getMe")
	body, err := sendRequest(nil, url, defaultType, "GET")
	if err == nil {
		err = json.Unmarshal(body, getme)
	}
	return getme, err
}
