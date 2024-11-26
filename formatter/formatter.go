package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/l1qwie/Fmtogram/fmerrors"
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

func buildConstRules() [7][7]int {
	var matrix [7][7]int

	matrix[constPhoto][constPhoto] = allowed
	matrix[constPhoto][constVideo] = allowed
	matrix[constAudio][constAudio] = allowed
	matrix[constVideo][constPhoto] = allowed
	matrix[constVideo][constVideo] = allowed
	matrix[constDoc][constDoc] = allowed
	return matrix
}

func compatibilityСheck(msg *Message, matrix [7][7]int, current int) error {
	var err error

	currentname, currentconst := msg.fm.mh.storage[current].nameAndConst()

	for j := 0; j < current; j++ {

		name, num := msg.fm.mh.storage[j].nameAndConst()
		if (matrix[currentconst][num] != allowed) || (msg.fm.method == methods.PaidMedia) {
			err = fmerrors.ImpossibleCombination(currentname, name, msg.fm.method)
		}
	}
	return err
}

func casePhoto(msg *Message, ph *photo, i int, matrix [7][7]int, object *string) error {
	err := requiredPhotoData(ph, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.Photo
		}
		*object = interfacePhoto
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func caseAudio(msg *Message, ad *audio, i int, matrix [7][7]int, object *string) error {
	err := requiredAudioData(ad, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.Audio
		}
		*object = interfaceAudio
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func caseVideo(msg *Message, vd *video, i int, matrix [7][7]int, object *string) error {
	err := requiredVideoData(vd, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.Video
		}
		*object = interfaceVideo
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func caseDocument(msg *Message, dc *document, i int, matrix [7][7]int, object *string) error {
	err := requiredDocumentData(dc, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.Document
		}
		*object = interfaceDocument
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func caseAnimation(msg *Message, an *animation, i int, matrix [7][7]int, object *string) error {
	err := requiredAnimationData(an, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.Animation
		}
		*object = interfaceAnimation
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func caseVoice(msg *Message, vc *voice, i int, matrix [7][7]int, object *string) error {
	err := requiredVoiceData(vc, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.Voice
		}
		*object = interfaceVoice
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func caseVideoNote(msg *Message, vdn *videonote, i int, matrix [7][7]int, object *string) error {
	err := requiredVideoNoteData(vdn, i)
	if err == nil {
		if !msg.fm.notchange {
			msg.fm.method = methods.VideoNote
		}
		*object = interfaceVideoNote
		err = compatibilityСheck(msg, matrix, i)
	}
	return err
}

func requiredMedia(msg *Message, tgr *interface{}, object *string) error {
	var err error

	matrix := buildConstRules()
	for i, j := 0, 0; (i < len(msg.fm.mh.storage)) && (err == nil); i++ {
		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				err = casePhoto(msg, m, i, matrix, object)
			case *audio:
				err = caseAudio(msg, m, i, matrix, object)
			case *video:
				err = caseVideo(msg, m, i, matrix, object)
			case *document:
				err = caseDocument(msg, m, i, matrix, object)
			case *animation:
				err = caseAnimation(msg, m, i, matrix, object)
			case *voice:
				err = caseVoice(msg, m, i, matrix, object)
			case *videonote:
				err = caseVideoNote(msg, m, i, matrix, object)
			}
			if err == nil {
				if j >= 1 {
					if !msg.fm.notchange {
						msg.fm.method = methods.MediaGroup
					}
					*object = "any of those you have mentioned: "
					for j := 0; j < i; j++ {
						name, _ := msg.fm.mh.storage[j].nameAndConst()
						*object = fmt.Sprintf("%s/%s", *object, name)
					}
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

	if (msg.fm.method == "") && (msg.fm.inf.Text == "") {
		err = fmerrors.MissedRequiredField("IMSGInformation", "WriteString()", 0, 0, false, false)
	}
	if _, ok := methods.Media[msg.fm.method]; (ok) && (msg.fm.inf.Text != "") {
		if msg.fm.mh.i > 1 {
			logs.DataMightBeLost("IMSGInformation", "WriteString()", msg.fm.inf.Text, object, "WriteCaption()")
		} else {
			msg.fm.inf.Caption, msg.fm.inf.Text = msg.fm.inf.Text, ""
			msg.fm.inf.CaptionEntities, msg.fm.inf.Entities = msg.fm.inf.Entities, nil
		}
	}
	fmt.Println("METHOD: ", msg.fm.method, msg.fm.inf.StarCount)
	if msg.fm.method == "" {
		msg.fm.method = methods.Message
		msg.fm.contentType = "application/json"
		*tgr = new(types.TelegramResponse)
	} else if (msg.fm.method == methods.ForwardMessage) || (msg.fm.method == methods.ForwardMessages) || (msg.fm.method == methods.CopyMessage) || (msg.fm.method == methods.CopyMessages) {
		if msg.fm.ch.FromChatID == nil {
			err = fmerrors.MissedRequiredField("IChat", "WriteFromChat{ID/Name}()", 0, 0, false, false)
		}
		if (msg.fm.inf.MessageID == 0) && ((msg.fm.method == methods.ForwardMessage) || (msg.fm.method == methods.CopyMessage)) {
			err = fmerrors.MissedRequiredField("IMSGInformation", "WriteMessageID()", 0, 0, false, false)
		}
		if (msg.fm.inf.MessageIDs == nil) && ((msg.fm.method == methods.ForwardMessages) || (msg.fm.method == methods.CopyMessages)) {
			err = fmerrors.MissedRequiredField("IMSGInformation", "WriteMessageIDs()", 0, 0, false, false)
		}
		msg.fm.contentType = "application/json"
		if (msg.fm.method == methods.ForwardMessage) || (msg.fm.method == methods.CopyMessage) {
			*tgr = new(types.TelegramResponse)
		} else {
			*tgr = new(types.TelegramMessageIDs)
		}
	} else if (msg.fm.method == methods.PaidMedia) && (msg.fm.inf.StarCount == 0) {
		err = fmerrors.MissedRequiredField("IMSGInformation", "WriteStarCount()", 0, 0, false, false)
	} else if ((msg.fm.method == methods.Video) || (msg.fm.method == methods.Photo)) && (msg.fm.inf.StarCount != 0) {
		msg.fm.method = methods.PaidMedia
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
						err = fmerrors.MissedRequiredField("IInlineButton", "WriteString()", j, i, false, true)
					}
				}
			}
		case *reply:
			for i := 0; i < len(kb.Keyboard.Keyboard) && err == nil; i++ {
				for j := 0; j < len(kb.Keyboard.Keyboard[i]) && err == nil; j++ {
					if kb.Keyboard.Keyboard[i][j].Text == "" {
						err = fmerrors.MissedRequiredField("IReplyButton", "WriteString()", j, i, false, true)
					}
				}
			}
		}
	}
	return err
}

func mediaPart(msg *Message) error {
	var err error

	if (msg.fm.mh.amount == 1) && (msg.fm.method != methods.PaidMedia) {
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
			err = msg.fm.kb.multipartFields(msg.fm.writer)
		} else {

			bytes, err = msg.fm.kb.jsonFields()

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

	if msg.fm.ch.ID == nil {
		err = fmerrors.ChatIDIsMissed()
	} else {
		if msg.fm.mh.atLeastOnce {
			err = msg.fm.ch.multipartFields(msg.fm.writer)
		} else {

			bytes, err = msg.fm.ch.createJson()

			if err == nil {
				msg.fm.buf.Write(bytes)
			}
		}
	}

	return err
}

func uniteEverything(msg *Message) error {
	var err error
	var mergedJSON []byte

	decoder := json.NewDecoder(msg.fm.buf)
	result := make(map[string]interface{})

	for err == nil {
		var data map[string]interface{}
		err = decoder.Decode(&data)
		if err != io.EOF {
			if err == nil {
				for k, v := range data {
					result[k] = v
				}
			} else {
				err = fmt.Errorf("error decoding JSON: %s", err)
			}
		}
	}
	if err == io.EOF {
		mergedJSON, err = json.Marshal(result)
		if err != nil {
			err = fmt.Errorf("error serializing merged JSON: %s", err)
		}
		msg.fm.buf = bytes.NewBuffer(nil)
		msg.fm.buf.Write(mergedJSON)
	}
	if err == io.EOF {
		err = nil
	}
	return err
}

func preMediaCheck(msg *Message) error {
	var found bool
	var err error
	if _, ok := methods.Media[msg.fm.method]; ok {
		for i := 0; i < len(msg.fm.mh.storage) && !found; i++ {
			if msg.fm.mh.storage[i] != nil {
				found = true
			}
		}
		if !found {
			err = fmerrors.MisMedia()
		}
	}
	return err
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

	if err = preMediaCheck(msg); err == nil {
		if msg.fm.mh.amount > 0 {
			err = requiredMedia(msg, tgr, &object)
		} else {
			shouldSkip[0] = true
		}

		if err == nil {
			if err = requiredMessage(msg, tgr, object); err == nil {
				err = requiredKeyboard(msg)
			}
		}
	}

	for i := 0; (i < len(doPlan)) && (err == nil); i++ {
		if !shouldSkip[i] {
			err = doPlan[i](msg)
		}
	}

	if err == nil {
		if !msg.fm.mh.atLeastOnce {
			err = uniteEverything(msg)
		} else {
			err = msg.fm.writer.Close()
		}
	}
	return err
}

func distributorTelegram(msg *Message, t *types.TelegramMediaGroup) {
	for i := 0; i < len(msg.fm.mh.storage); i++ {

		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				for j := 0; j < len(t.Result[i].Photo); j++ {
					m.response[j] = types.PhotoSize{
						FileID:       t.Result[i].Photo[j].FileID,
						FileSize:     t.Result[i].Photo[j].FileSize,
						FileUniqueID: t.Result[i].Photo[j].FileUniqueID,
						Width:        t.Result[i].Photo[j].Width,
						Height:       t.Result[i].Photo[j].Height,
					}
				}
			case *video:
				m.response = *t.Result[i].Video
			case *audio:
				m.response = *t.Result[i].Audio
			case *document:
				m.response = *t.Result[i].Document
			case *animation:
				m.response = *t.Result[i].Animation
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
					if msg.fm.method == methods.PaidMedia {
						m.response[j] = types.PhotoSize{
							FileID:       t.Result.PaidMedia.PaidMedia[i].Photo[j].FileID,
							FileSize:     t.Result.PaidMedia.PaidMedia[i].Photo[j].FileSize,
							FileUniqueID: t.Result.PaidMedia.PaidMedia[i].Photo[j].FileUniqueID,
							Width:        t.Result.PaidMedia.PaidMedia[i].Photo[j].Width,
							Height:       t.Result.PaidMedia.PaidMedia[i].Photo[j].Height,
						}
					} else {
						m.response[j] = types.PhotoSize{
							FileID:       t.Result.Photo[j].FileID,
							FileSize:     t.Result.Photo[j].FileSize,
							FileUniqueID: t.Result.Photo[j].FileUniqueID,
							Width:        t.Result.Photo[j].Width,
							Height:       t.Result.Photo[j].Height,
						}
					}
				}
			case *video:
				if msg.fm.method == methods.PaidMedia {
					m.response = *t.Result.PaidMedia.PaidMedia[i].Video
				} else {
					m.response = *t.Result.Video
				}
			case *audio:
				m.response = *t.Result.Audio
			case *document:
				m.response = *t.Result.Document
			case *animation:
				m.response = *t.Result.Animation
			case *videonote:
				m.response = *t.Result.VideoNote
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
			err = fmerrors.TelegramError(t.ErrorCode, t.Description)
		} else {
			distributorTelegram(msg, t)
		}
	case *types.TelegramResponse:
		if t.ErrorCode != 0 {
			err = fmerrors.TelegramError(t.ErrorCode, t.Description)
		} else {
			distributorTelegramResponse(msg, t)
		}
	case *types.TelegramMessageIDs:
		if t.ErrorCode != 0 {
			err = fmerrors.TelegramError(t.ErrorCode, t.Description)
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
	log.Print(msg.fm.method)
	// log.Print(msg.fm.contentType)

	url := fmt.Sprint(types.TelegramAPI, "bot", testbotdata.Token, "/", msg.fm.method)

	req, err := http.NewRequest("POST", url, msg.fm.buf)
	req.Header.Set("Content-Type", msg.fm.contentType)

	if err != nil {
		err = fmerrors.CantMakeRequest(err)
	} else {

		client := http.Client{}
		resp, err = client.Do(req)

		if err != nil {
			err = fmerrors.CantGetResponse(err)
		} else {

			defer resp.Body.Close()
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				err = fmerrors.CantReadResponse(err)
			} else {

				fmt.Println(string(body))

				err = json.Unmarshal(body, tgr)
				if err != nil {
					err = fmerrors.CantUnmarshal(err)
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
