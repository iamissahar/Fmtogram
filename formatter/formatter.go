package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func CreateMessage(tg *types.Telegram) *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{ID: tg.Result[0].Message.Chat.ID}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	return m
}

func CreateEmpltyMessage() *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	return m
}

func requiredMedia(msg *Message, tgr *interface{}, object *string) error {
	var err error
	var storage [4]bool
	var obj2 string

	if msg.fm.ch.ID == 0 {
		err = errors.MissedRequiredField("IChat", "WriteChat{ID/Name}()", 0, 0, false, false)
	}
	for i, j := 0, 0; (i < len(msg.fm.mh.storage)) && (err == nil); i++ {
		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				requiredPhotoData(m, i)
				if !msg.fm.notchange {
					msg.fm.method = methods.Photo
				}
				*object = "IPhoto"
				if storage[1] || storage[3] {
					if storage[1] {
						obj2 = "Audio"
					} else {
						obj2 = "Document"
					}
					err = errors.ImpossibleCombination("Photo", obj2)
				} else {
					storage[0] = true
				}
			case *audio:
				requiredAudioData(m, i)
				if !msg.fm.notchange {
					msg.fm.method = methods.Audio
				}
				*object = "IAudio"
				if storage[0] || storage[2] || storage[3] {
					if storage[0] {
						obj2 = "Photo"
					} else if storage[2] {
						obj2 = "Video"
					} else {
						obj2 = "Document"
					}
					err = errors.ImpossibleCombination("Audio", obj2)
				} else {
					storage[1] = true
				}
			case *video:
				requiredVideoData(m, i)
				if !msg.fm.notchange {
					msg.fm.method = methods.Video
				}
				*object = "IVideo"
				if storage[1] || storage[3] {
					if storage[1] {
						obj2 = "Audio"
					} else {
						obj2 = "Document"
					}
					err = errors.ImpossibleCombination("Video", obj2)
				} else {
					storage[2] = true
				}
			case *document:
				requiredDocumentData(m, i)
				if !msg.fm.notchange {
					msg.fm.method = methods.Document
				}
				*object = "IDocument"
				if storage[0] || storage[1] || storage[2] {
					if storage[0] {
						obj2 = "Photo"
					} else if storage[1] {
						obj2 = "Audio"
					} else {
						obj2 = "Video"
					}
					err = errors.ImpossibleCombination("Document", obj2)
				} else {
					storage[3] = true
				}
			}
			if err == nil {
				if j >= 1 {
					if !msg.fm.notchange {
						msg.fm.method = methods.MediaGroup
					}
					*object = "any of those you have mentioned (IPhoto/IVideo/IDocument/IAudio)"
					*tgr = new(types.TelegramMediaGroup)
				} else {
					*tgr = new(types.TelegramResponse)
				}
				j++
			}
		}
	}
	return err
}

func requiredMessage(msg *Message, tgr *interface{}, object string) error {
	var err error

	if msg.fm.ch.ID == 0 {
		err = errors.MissedRequiredField("IChat", "WriteChat{ID/Name}()", 0, 0, false, false)
	}
	if (msg.fm.method == "") && (msg.fm.inf.Text == "") {
		err = errors.MissedRequiredField("IMSGInformation", "WriteString()", 0, 0, false, false)
	}
	if _, ok := methods.Media[msg.fm.method]; (ok) && (msg.fm.inf.Text != "") {
		logs.DataMightBeLost("IMSGInformation", "WriteString()", msg.fm.inf.Text, object, "WriteCaption()")
	}
	if msg.fm.method == "" {
		msg.fm.method = methods.Message
		msg.fm.contentType = "application/json"
		*tgr = new(types.TelegramResponse)
	} else if (msg.fm.method == methods.ForwardMessage) || (msg.fm.method == methods.ForwardMessages) || (msg.fm.method == methods.CopyMessage) || (msg.fm.method == methods.CopyMessages) {
		if msg.fm.ch.FromChatID == nil {
			err = errors.MissedRequiredField("IChat", "WriteFromChat{ID/Name}()", 0, 0, false, false)
		}
		if (msg.fm.inf.MessageID == 0) && ((msg.fm.method == methods.ForwardMessage) || (msg.fm.method == methods.CopyMessage)) {
			err = errors.MissedRequiredField("IMSGInformation", "WriteMessageID()", 0, 0, false, false)
		}
		if (msg.fm.inf.MessageIDs == nil) && ((msg.fm.method == methods.ForwardMessages) || (msg.fm.method == methods.CopyMessages)) {
			err = errors.MissedRequiredField("IMSGInformation", "WriteMessageIDs()", 0, 0, false, false)
		}
		msg.fm.contentType = "application/json"
		if (msg.fm.method == methods.ForwardMessage) || (msg.fm.method == methods.CopyMessage) {
			*tgr = new(types.TelegramResponse)
		} else {
			*tgr = new(types.TelegramMessageIDs)
		}
	}
	return err
}

func requiredKeyboard(msg *Message) error {
	var err error
	if msg.fm.kb != nil {
		switch kb := msg.fm.kb.(type) {
		case *inline:
			for i := 0; i < len(kb.Keyboard.InlineKeyboard) && err == nil; i++ {
				for j := 0; j < len(kb.Keyboard.InlineKeyboard[i]) && err == nil; j++ {
					if kb.Keyboard.InlineKeyboard[i][j].Text == "" {
						err = errors.MissedRequiredField("IInlineButton", "WriteString()", j, i, false, true)
					}
					if kb.Keyboard.InlineKeyboard[i][j].One > 1 {
						err = errors.TooMuchData("IInlineButton", j, i)
					}
				}
			}
		case *reply:
			//
			//
		}
	}
	return err
}

func mediaPart(msg *Message) error {
	var err error

	if msg.fm.mh.amount == 1 {
		err = uniqueMedia(msg)
	} else {
		err = mediaGroup(msg)
	}
	return err
}

func messagePart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.mh.atLeastOnce {
		err = msg.fm.inf.multipartFields(msg.fm.writer)
	} else {

		bytes, err = msg.fm.inf.createJSON()

		if err == nil {
			msg.fm.buf.Write(bytes)
		}
	}
	return err
}

func keyboardPart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.kb != nil {

		if msg.fm.mh.atLeastOnce {
			err = msg.fm.kb.MultipartFields(msg.fm.writer)
		} else {

			bytes, err = msg.fm.kb.JsonFields()

			if err == nil {
				msg.fm.buf.Write(bytes)
			}
		}

	}
	return err
}

func chatPart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.mh.atLeastOnce {
		err = msg.fm.ch.multipartFields(msg.fm.writer)
	} else {

		bytes, err = msg.fm.ch.createJson()

		if err == nil {
			msg.fm.buf.Write(bytes)
		}
	}

	return err
}

// not a final result
func uniteEverything(msg *Message) error {
	decoder := json.NewDecoder(msg.fm.buf)

	result := make(map[string]interface{})

	for {
		var data map[string]interface{}
		if err := decoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error decoding JSON: %s", err)
		}

		for k, v := range data {
			result[k] = v
		}
	}

	mergedJSON, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error serializing merged JSON: %s", err)
	}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.buf.Write(mergedJSON)
	return nil
}

func makeRequest(msg *Message, tgr *interface{}) error {
	var err error
	var object string
	var shouldSkip [4]bool
	// chat part must be skipped by default
	// it isn't skipped now, because of tests
	doPlan := [4]func(*Message) error{mediaPart, messagePart, keyboardPart, chatPart}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.writer = multipart.NewWriter(msg.fm.buf)

	if msg.fm.mh.amount > 0 {
		err = requiredMedia(msg, tgr, &object)
	} else {
		shouldSkip[0] = true
	}

	if err == nil {
		err = requiredMessage(msg, tgr, object)
	}
	if err == nil {
		err = requiredKeyboard(msg)
	}

	for i := 0; (i < len(doPlan)) && (err == nil); i++ {
		if !shouldSkip[i] {
			err = doPlan[i](msg)
		}
	}
	if err == nil && !msg.fm.mh.atLeastOnce {
		err = uniteEverything(msg)
	} else if err == nil && msg.fm.mh.atLeastOnce {
		err = msg.fm.writer.Close()
	}
	return err
}

func distributorTelegram(msg *Message, t *types.TelegramMediaGroup) {
	for i := 0; i < len(msg.fm.mh.storage); i++ {

		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				for j := 0; j < len(t.Result[0].Photo); j++ {
					m.response[j] = types.PhotoSize{
						FileID:       t.Result[0].Photo[j].FileID,
						FileSize:     t.Result[0].Photo[j].FileSize,
						FileUniqueID: t.Result[0].Photo[j].FileUniqueID,
						Width:        t.Result[0].Photo[j].Width,
						Height:       t.Result[0].Photo[j].Height,
					}
				}
			case *video:
				m.response = *t.Result[0].Video
			case *audio:
				m.response = *t.Result[0].Audio
			case *document:
				m.response = *t.Result[0].Document
			}
		}
	}
	msg.fm.ch.response = *t.Result[0].Chat
	msg.fm.inf.response = *t.Result[0].From

}

func distributorTelegramResponse(msg *Message, t *types.TelegramResponse) {
	for i := 0; i < len(msg.fm.mh.storage); i++ {

		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				for j := 0; j < len(t.Result.Photo); j++ {
					m.response[j] = types.PhotoSize{
						FileID:       t.Result.Photo[j].FileID,
						FileSize:     t.Result.Photo[j].FileSize,
						FileUniqueID: t.Result.Photo[j].FileUniqueID,
						Width:        t.Result.Photo[j].Width,
						Height:       t.Result.Photo[j].Height,
					}
				}
			case *video:
				m.response = *t.Result.Video
			case *audio:
				m.response = *t.Result.Audio
			case *document:
				m.response = *t.Result.Document
			}
		}
	}
	if t.Result.Chat != nil {
		msg.fm.ch.response = *t.Result.Chat
	}
	if t.Result.From != nil {
		msg.fm.inf.response = *t.Result.From
	}
	msg.fm.inf.responseMessageIDs = append(msg.fm.inf.responseMessageIDs, t.Result.MessageID)
}

func responseDecoder(msg *Message, tgr interface{}) error {
	var err error
	switch t := tgr.(type) {
	case *types.TelegramMediaGroup:
		if t.ErrorCode != 0 {
			err = errors.TelegramError(t.ErrorCode, t.Description)
		} else {
			distributorTelegram(msg, t)
		}
	case *types.TelegramResponse:
		if t.ErrorCode != 0 {
			err = errors.TelegramError(t.ErrorCode, t.Description)
		} else {
			distributorTelegramResponse(msg, t)
		}
	case *types.TelegramMessageIDs:
		if t.ErrorCode != 0 {
			err = errors.TelegramError(t.ErrorCode, t.Description)
		} else {
			msg.fm.inf.responseMessageIDs = make([]int, len(t.Result))
			for i, v := range t.Result {
				msg.fm.inf.responseMessageIDs[i] = v.MessageID
			}
		}
	}
	return err
}

func sendRequest(msg *Message, tgr interface{}) error {
	var resp *http.Response
	var body []byte

	// log.Print(msg.fm.buf.String())
	// log.Print(msg.fm.method)
	// log.Print(msg.fm.contentType)

	url := fmt.Sprint(types.TelegramAPI, "bot", testbotdata.Token, "/", msg.fm.method)

	// log.Print(url)

	req, err := http.NewRequest("POST", url, msg.fm.buf)
	req.Header.Set("Content-Type", msg.fm.contentType)

	if err != nil {
		err = errors.CantMakeRequest(err)
	} else {

		client := http.Client{}
		resp, err = client.Do(req)

		if err != nil {
			err = errors.CantGetResponse(err)
		} else {

			defer resp.Body.Close()
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				err = errors.CantReadResponse(err)
			} else {

				// fmt.Println(string(body))

				err = json.Unmarshal(body, tgr)
				if err != nil {
					err = errors.CantUnmarshal(err)
				} else {
					err = responseDecoder(msg, tgr)
				}
			}
		}
	}
	return err
}

// Sends the message you created. May send you an error if you made a mistake during the build process.
// Also sends a new pointer to the structure. This is a quick response from Telegram to your message. It collects some data
// about the message sent. For example: if you want to send a photo, this structure will have data about it from Telegram (id, size, etc.)
func (msg *Message) Send() (interface{}, error) {
	var err error
	var tgr interface{}

	if err = makeRequest(msg, &tgr); err == nil {
		if tgr == nil {
			panic("!~")
		}
		err = sendRequest(msg, tgr)
	}

	return tgr, err
}
