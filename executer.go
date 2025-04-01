package fmtogram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	none        int    = -1
	defaultType string = "application/json"
)

var (
	timeout   int = 7
	limit     int = 10
	byDefault     = []string{"message", "edited_message", "channel_post", "edited_channel_post", "business_connection",
		"business_message", "edited_business_message", "deleted_business_messages", "message_reaction", "message_reaction_count",
		"inline_query", "chosen_inline_result", "callback_query", "shipping_query", "pre_checkout_query", "purchased_paid_media", "poll",
		"poll_answer", "my_chat_member", "chat_member", "chat_join_request", "chat_boost", "removed_chat_boost"}
)

type entry struct {
	UserId int
	Chu    chan *Telegram
	Chb    chan *Telegram
}

type regTable struct {
	Reg []entry
}

func code50() error {
	err := new(FME)
	err.Code = 50
	err.String = "The value of fmtogram.BasicSettings structure cannot be nil"
	return err
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

func (reg *regTable) seeker(chatID int) int {
	var i int
	index := none
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

func (reg *regTable) newIndex() (newIndex int) {
	newIndex = len(reg.Reg)
	reg.Reg = append(reg.Reg, entry{})
	return newIndex
}

func getUpdates(tg *Telegram, offset *int, bs *BasicSettings) {
	var (
		values      = url.Values{}
		alupd, body []byte
		err         error
	)
	if bs.AllowedUpdates == nil {
		bs.AllowedUpdates = byDefault
	}
	if bs.Timeout == nil {
		bs.Timeout = &timeout
	}
	if bs.Limit == nil {
		bs.Limit = &limit
	}

	alupd, err = json.Marshal(bs.AllowedUpdates)
	if err == nil {
		values.Add("offset", fmt.Sprint(*offset))
		values.Add("allowed_updates", fmt.Sprint(alupd))
		values.Add("timeout", fmt.Sprint(*bs.Timeout))
		values.Add("limit", fmt.Sprint(*bs.Limit))
		url := TelegramAPI + "bot" + bs.Token + "/" + "getUpdates" + "?" + values.Encode()
		body, err = sendRequest(bytes.NewBuffer(nil), url, "application/json", "GET")
		if err == nil {
			err = json.Unmarshal(body, tg)
		}
	}
	if err != nil {
		if bs.ErrorHandler != nil {
			bs.ErrorHandler <- err
		} else {
			panic(err)
		}
	}
}

func getOffset(offset *int, bs *BasicSettings) {
	var values = url.Values{}
	tg := new(Telegram)
	values.Add("limit", "1")
	url := fmt.Sprint(TelegramAPI, fmt.Sprintf("bot%s", bs.Token), "/",
		"getUpdates", "?", values.Encode())
	body, err := sendRequest(bytes.NewBuffer(nil), url, "application/json", "GET")
	if err == nil {
		err = json.Unmarshal(body, tg)
		if err == nil {

			if len(tg.Result) > 0 {
				*offset = tg.Result[0].UpdateID
			}

		}
	}
	if err != nil {
		if bs.ErrorHandler != nil {
			bs.ErrorHandler <- err
		} else {
			panic(err)
		}
	}

}

func GetMe() (*GetMee, error) {
	getme := new(GetMee)
	url := fmt.Sprint(
		TelegramAPI,
		fmt.Sprintf("bot%s", BotID), "/", "getMe")
	body, err := sendRequest(bytes.NewBuffer(nil), url, defaultType, "GET")
	if err == nil {
		err = json.Unmarshal(body, getme)
	}
	return getme, err
}

func DownloadFile(filepath, savepath string) error {
	var file *os.File
	url := fmt.Sprint(
		TelegramAPI,
		fmt.Sprintf("file/bot%s", BotID), "/", filepath)
	body, err := sendRequest(bytes.NewBuffer(nil), url, defaultType, "GET")
	if err == nil {
		file, err = os.Create(savepath)
		if err == nil {
			defer file.Close()
			_, err = file.Write(body)
		}
	}
	return err
}
