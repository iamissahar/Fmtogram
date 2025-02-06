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
	"github.com/l1qwie/Fmtogram/types"
)

func CreateMessage(tg *types.Telegram, botID string) *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{ID: tg.Result[0].Message.Chat.ID}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	m.fm.token = botID
	return m
}

func CreateEmpltyMessage() *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	m.fm.token = types.BotID
	return m
}

// func casePhoto(msg *Message, ph *photo, i int, matrix [7][7]int, object *string) error {
// 	err := requiredPhotoData(ph, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.Photo
// 		}
// 		*object = interfacePhoto
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func caseAudio(msg *Message, ad *audio, i int, matrix [7][7]int, object *string) error {
// 	err := requiredAudioData(ad, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.Audio
// 		}
// 		*object = interfaceAudio
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func caseVideo(msg *Message, vd *video, i int, matrix [7][7]int, object *string) error {
// 	err := requiredVideoData(vd, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.Video
// 		}
// 		*object = interfaceVideo
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func caseDocument(msg *Message, dc *document, i int, matrix [7][7]int, object *string) error {
// 	err := requiredDocumentData(dc, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.Document
// 		}
// 		*object = interfaceDocument
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func caseAnimation(msg *Message, an *animation, i int, matrix [7][7]int, object *string) error {
// 	err := requiredAnimationData(an, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.Animation
// 		}
// 		*object = interfaceAnimation
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func caseVoice(msg *Message, vc *voice, i int, matrix [7][7]int, object *string) error {
// 	err := requiredVoiceData(vc, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.Voice
// 		}
// 		*object = interfaceVoice
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func caseVideoNote(msg *Message, vdn *videonote, i int, matrix [7][7]int, object *string) error {
// 	err := requiredVideoNoteData(vdn, i)
// 	if err == nil {
// 		if !msg.fm.notchange {
// 			msg.fm.method = methods.VideoNote
// 		}
// 		*object = interfaceVideoNote
// 		err = compatibilityСheck(msg, matrix, i)
// 	}
// 	return err
// }

// func requiredMedia(msg *Message, tgr *interface{}, object *string) error {
// 	var err error

// 	for i, j := 0, 0; (i < len(msg.fm.mh.storage)) && (err == nil); i++ {
// 		if msg.fm.mh.storage[i] != nil {
// 			switch m := msg.fm.mh.storage[i].(type) {
// 			case *photo:
// 				err = casePhoto(msg, m, i, matrix, object)
// 			case *audio:
// 				err = caseAudio(msg, m, i, matrix, object)
// 			case *video:
// 				err = caseVideo(msg, m, i, matrix, object)
// 			case *document:
// 				err = caseDocument(msg, m, i, matrix, object)
// 			case *animation:
// 				err = caseAnimation(msg, m, i, matrix, object)
// 			case *voice:
// 				err = caseVoice(msg, m, i, matrix, object)
// 			case *videonote:
// 				err = caseVideoNote(msg, m, i, matrix, object)
// 			}
// 			if err == nil {
// 				if j >= 1 {
// 					if !msg.fm.notchange {
// 						msg.fm.method = methods.MediaGroup
// 					}
// 					*object = "any of those you have mentioned: "
// 					for j := 0; j < i; j++ {
// 						name, _ := msg.fm.mh.storage[j].nameAndConst()
// 						*object = fmt.Sprintf("%s/%s", *object, name)
// 					}
// 					*tgr = new(types.TelegramMediaGroup)
// 				} else {
// 					*tgr = new(types.TelegramResponse)
// 				}
// 				j++
// 			}
// 		}
// 	}
// 	return err
// }

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
		if bytes, err = msg.fm.inf.createJSON(); err == nil {
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
			if bytes, err = msg.fm.kb.jsonFields(); err == nil {
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
		if bytes, err = msg.fm.ch.createJson(); err == nil {
			msg.fm.buf.Write(bytes)
		}
	}

	return err
}

func writeToBuf(msg *Message, structer interface{}) error {
	var err error
	var body []byte
	if body, err = json.Marshal(structer); err == nil {
		msg.fm.buf.Write(body)
	}
	return err
}

func othersPart(msg *Message) error {
	var err error
	if msg.fm.loc != nil {
		err = writeToBuf(msg, msg.fm.loc)
	}
	if msg.fm.con != nil {
		err = writeToBuf(msg, msg.fm.con)
	}
	if msg.fm.poll != nil {
		err = writeToBuf(msg, msg.fm.poll)
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

func makeRequest(msg *Message) error {
	var err error
	var shouldSkip [5]bool
	// chat part must be skipped by default
	// it isn't skipped now, because of tests
	doPlan := [5]func(*Message) error{mediaPart, messagePart, keyboardPart, chatPart, othersPart}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.writer = multipart.NewWriter(msg.fm.buf)

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
	if msg.fm.contentType == "" {
		msg.fm.contentType = "application/json"
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
				if t.Result.VideoNote != nil {
					m.response = *t.Result.VideoNote
				} else {
					m.response = types.VideoNote{
						FileID:       t.Result.Video.FileID,
						FileUniqueID: t.Result.Video.FileUniqueID,
						Duration:     t.Result.Video.Duration,
						Thumbnail:    t.Result.Video.Thumbnail,
						FileSize:     t.Result.Video.FileSize,
					}
				}
			case *voice:
				if t.Result.Voice != nil {
					m.response = *t.Result.Voice
				}
			}
		}
	}
	if t.Result.Chat != nil {
		msg.fm.g.chat = *t.Result.Chat
	}
	if t.Result.From != nil {
		msg.fm.g.user = *t.Result.From
		msg.fm.inf.response = *t.Result.From
	}
	if msg.fm.loc != nil {
		if t.Result.Venue != nil {
			msg.fm.loc.response = *t.Result.Venue
		} else if t.Result.Location != nil {
			msg.fm.loc.response = types.Venue{Location: t.Result.Location}
		}
	}
	if msg.fm.con != nil {
		if t.Result.Contact != nil {
			msg.fm.con.response = *t.Result.Contact
		}
	}
	msg.fm.g.msgID = t.Result.MessageID
	msg.fm.g.date = t.Result.Date
	msg.fm.g.status = t.Ok
	msg.fm.inf.responseMessageIDs = append(msg.fm.inf.responseMessageIDs, t.Result.MessageID)
}

func responseDecoder(msg *Message) error {
	var err error
	switch t := msg.fm.tgr.(type) {
	case *types.TelegramMediaGroup:
		if t.ErrorCode != 0 {
			err = fmerrors.TelegramError(t.ErrorCode, t.Description)
			if msg.fm.g != nil {
				msg.fm.g.status, msg.fm.g.errorCode, msg.fm.g.errorMsg = t.Ok, t.ErrorCode, t.Description
			}
		} else {
			distributorTelegram(msg, t)
		}
	case *types.TelegramResponse:
		if t.ErrorCode != 0 {
			err = fmerrors.TelegramError(t.ErrorCode, t.Description)
			if msg.fm.g != nil {
				msg.fm.g.status, msg.fm.g.errorCode, msg.fm.g.errorMsg = t.Ok, t.ErrorCode, t.Description
			}
		} else {
			distributorTelegramResponse(msg, t)
		}
	case *types.TelegramMessageIDs:
		if t.ErrorCode != 0 {
			err = fmerrors.TelegramError(t.ErrorCode, t.Description)
			if msg.fm.g != nil {
				msg.fm.g.status, msg.fm.g.errorCode, msg.fm.g.errorMsg = t.Ok, t.ErrorCode, t.Description
			}
		} else {
			msg.fm.inf.responseMessageIDs = make([]int, len(t.Result))
			for i, v := range t.Result {
				msg.fm.inf.responseMessageIDs[i] = v.MessageID
			}
		}
	}
	return err
}

func sendRequest(msg *Message) error {
	var resp *http.Response
	var body []byte

	if msg.fm.contentType == "application/json" {
		log.Print(msg.fm.buf.String())
	}
	log.Print(msg.fm.method)
	// log.Print(msg.fm.contentType)

	url := fmt.Sprint(types.TelegramAPI, "bot", msg.fm.token, "/", msg.fm.method)

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

				err = json.Unmarshal(body, msg.fm.tgr)
				if err != nil {
					err = fmerrors.CantUnmarshal(err)
				} else {
					err = responseDecoder(msg)
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

	if err = makeRequest(msg); err == nil {
		if msg.fm.tgr == nil {
			panic("!~")
		}
		err = sendRequest(msg)
	}

	return msg.fm.tgr, err
}
