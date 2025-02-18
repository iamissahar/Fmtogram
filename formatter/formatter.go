package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func CreateMessage(tg *types.Telegram, botID string) *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{ID: tg.Result[0].Message.Chat.ID}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	m.fm.sticker = &stickerHolder{}
	m.fm.token = botID
	return m
}

func CreateEmpltyMessage() *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	m.fm.sticker = &stickerHolder{}
	m.fm.token = types.BotID
	return m
}

func uniqueMedia(msg *Message) error {
	var err error
	var mediajson []byte
	if msg.fm.mh.atLeastOnce {
		if err = msg.fm.mh.storage[msg.fm.mh.i-1].multipartFields(msg.fm.writer, nil, 0, false); err == nil {
			msg.fm.contentType = msg.fm.writer.FormDataContentType()
		}
	} else {
		if mediajson, err = msg.fm.mh.storage[msg.fm.mh.i-1].jsonFileds(); err == nil {
			msg.fm.buf.Write(mediajson)
		}
	}
	return err
}

func mediaGroup(msg *Message) error {
	var (
		err    error
		jsbody []byte
	)
	group := make([]interface{}, msg.fm.mh.amount)
	mediaMap := make(map[string]interface{}, 1)
	if !msg.fm.mh.atLeastOnce {
		for i := 0; i < len(msg.fm.mh.storage); i++ {

			if msg.fm.mh.storage[i] != nil {
				group[i] = msg.fm.mh.storage[i]
			}
		}

		mediaMap["media"] = group
		jsbody, err = json.Marshal(mediaMap)
		if err == nil {
			msg.fm.buf.Write(jsbody)
		}
	} else {

		for i := 0; i < len(msg.fm.mh.storage) && err == nil; i++ {
			if msg.fm.mh.storage[i] != nil {
				err = msg.fm.mh.storage[i].multipartFields(msg.fm.writer, &group, i, true)
			}
		}
		if err == nil {
			err = putGroup(msg.fm.writer, group, "media")
		}
		msg.fm.contentType = msg.fm.writer.FormDataContentType()
	}
	return err
}

func mediaPart(msg *Message) error {
	var err error
	if msg.fm.mh.amount != 0 {
		if (msg.fm.mh.amount == 1) && (msg.fm.method != methods.PaidMedia) {
			err = uniqueMedia(msg)
		} else {
			err = mediaGroup(msg)
		}
	}
	return err
}

func messagePart(msg *Message) error {
	var err error
	var bytes []byte
	if msg.fm.mh.atLeastOnce || msg.fm.sticker.atLeastOnce {
		err = msg.fm.inf.multipartFields(msg.fm.writer)
	} else {
		if msg.fm.method == methods.PromoteMember {
			if msg.fm.inf.AdminRights != nil {
				msg.fm.inf.IsAnonymous = msg.fm.inf.AdminRights.IsAnonymous
				msg.fm.inf.CanManageChat = msg.fm.inf.AdminRights.CanManageChat
				msg.fm.inf.CanDeleteMessages = msg.fm.inf.AdminRights.CanDeleteMessages
				msg.fm.inf.CanManageVideoChats = msg.fm.inf.AdminRights.CanManageVideoChats
				msg.fm.inf.CanRestrictMembers = msg.fm.inf.AdminRights.CanRestrictMembers
				msg.fm.inf.CanPromoteMembers = msg.fm.inf.AdminRights.CanPromoteMembers
				msg.fm.inf.CanChangeInfo = msg.fm.inf.AdminRights.CanChangeInfo
				msg.fm.inf.CanInviteUsers = msg.fm.inf.AdminRights.CanInviteUsers
				msg.fm.inf.CanPostStories = msg.fm.inf.AdminRights.CanPostStories
				msg.fm.inf.CanEditStories = msg.fm.inf.AdminRights.CanEditStories
				msg.fm.inf.CanDeleteStories = msg.fm.inf.AdminRights.CanDeleteStories
				msg.fm.inf.CanPostMessages = msg.fm.inf.AdminRights.CanPostMessages
				msg.fm.inf.CanEditMessages = msg.fm.inf.AdminRights.CanEditMessages
				msg.fm.inf.CanPinMessages = msg.fm.inf.AdminRights.CanPinMessages
				msg.fm.inf.CanManageTopics = msg.fm.inf.AdminRights.CanManageTopics
				msg.fm.inf.AdminRights = nil
			}
		}
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
		if msg.fm.mh.atLeastOnce || msg.fm.sticker.atLeastOnce {
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

	if msg.fm.mh.atLeastOnce || msg.fm.sticker.atLeastOnce {
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
	if msg.fm.link != nil {
		err = writeToBuf(msg, msg.fm.link)
	}
	if msg.fm.forum != nil {
		err = writeToBuf(msg, msg.fm.forum)
	}
	if msg.fm.bot != nil {
		err = writeToBuf(msg, msg.fm.bot)
	}
	if msg.fm.inlinemode != nil {
		err = writeToBuf(msg, msg.fm.inlinemode)
	}
	if msg.fm.payment != nil {
		err = writeToBuf(msg, msg.fm.payment)
	}
	if msg.fm.game != nil {
		err = writeToBuf(msg, msg.fm.game)
	}
	return err
}

func uniqueSticker(msg *Message) error {
	var err error
	if msg.fm.sticker.atLeastOnce {
		err = msg.fm.sticker.storage[0].multipartFields(msg.fm.writer, nil, 0, false)
	} else {
		err = writeToBuf(msg, msg.fm.sticker.storage[0])
	}
	return err
}

func arrayOfStickers(msg *Message, group, finalHolder *[]interface{}, stickerMap *map[string]interface{}) error {
	const stickers string = "stickers"
	var err error
	for i := 0; i < msg.fm.sticker.i; i++ {
		st := msg.fm.sticker.storage[i]
		if msg.fm.sticker.atLeastOnce {
			err = msg.fm.sticker.storage[i].multipartFields(msg.fm.writer, group, i, true)
		} else {
			(*finalHolder)[i] = st
		}

	}
	if err == nil {
		if msg.fm.sticker.atLeastOnce {
			err = putGroup(msg.fm.writer, *group, stickers)
			msg.fm.contentType = msg.fm.writer.FormDataContentType()
		} else {
			(*stickerMap)[stickers] = finalHolder
			err = writeToBuf(msg, stickerMap)
		}
	}
	return err
}

func stickerPart(msg *Message) error {
	var err error
	group := make([]interface{}, msg.fm.sticker.i)
	finalHolder := make([]interface{}, msg.fm.sticker.i)
	stickerMap := make(map[string]interface{}, 1)

	if msg.fm.method == methods.Sticker {
		err = uniqueSticker(msg)
	} else {
		err = arrayOfStickers(msg, &group, &finalHolder, &stickerMap)
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
	var shouldSkip [6]bool
	bypass := [6]func(*Message) error{mediaPart, messagePart, keyboardPart, chatPart, othersPart, stickerPart}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.writer = multipart.NewWriter(msg.fm.buf)

	for i := 0; (i < len(bypass)) && (err == nil); i++ {
		if !shouldSkip[i] {
			err = bypass[i](msg)
		}
	}

	if err == nil {
		if !msg.fm.mh.atLeastOnce && !msg.fm.sticker.atLeastOnce {
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

func handlerMessage(msg *Message, t *types.MessageResponse) {
	msg.fm.g.status = t.Ok
	msg.fm.g.msgID = t.Result.MessageID
	msg.fm.g.sender = t.Result.From
	msg.fm.g.date = t.Result.Date
	msg.fm.g.chat = t.Result.Chat
	if t.Result.ReplyToMessage != nil {
		msg.fm.g.replyed = &get{chat: t.Result.ReplyToMessage.Chat, sender: t.Result.From,
			date: t.Result.ReplyToMessage.Date, msgID: t.Result.ReplyToMessage.MessageID,
			msgOrigin: t.Result.ReplyToMessage.ForwardOrigin, photo: t.Result.ReplyToMessage.Photo,
			audio: t.Result.ReplyToMessage.Audio, document: t.Result.ReplyToMessage.Document,
			video: t.Result.ReplyToMessage.Video, anim: t.Result.ReplyToMessage.Animation,
			voice: t.Result.ReplyToMessage.Voice, vdn: t.Result.ReplyToMessage.VideoNote,
			poll: t.Result.ReplyToMessage.Poll, dice: t.Result.ReplyToMessage.Dice}
		if t.Result.ReplyToMessage.PaidMedia != nil {
			msg.fm.g.replyed.paid = t.Result.ReplyToMessage.PaidMedia.PaidMedia[0]
			if t.Result.PaidMedia != nil {
				msg.fm.g.paid = t.Result.PaidMedia.PaidMedia[0]
			}
		}
	}
	msg.fm.g.msgOrigin = t.Result.ForwardOrigin
	msg.fm.g.photo = t.Result.Photo
	msg.fm.g.audio = t.Result.Audio
	msg.fm.g.document = t.Result.Document
	msg.fm.g.video = t.Result.Video
	msg.fm.g.anim = t.Result.Animation
	msg.fm.g.voice = t.Result.Voice
	msg.fm.g.vdn = t.Result.VideoNote
	msg.fm.g.poll = t.Result.Poll
	msg.fm.g.dice = t.Result.Dice
	if t.Result.Sticker != nil {
		msg.fm.g.stickers = []*types.Sticker{t.Result.Sticker}
	}
	msg.fm.g.msg = t.Result
}

func handlerMediaGroup(msg *Message, t *types.MediaGroupResponse) {
	msg.fm.g.status = t.Ok
	msg.fm.g.mg.id = t.Result[0].MediaGroupID
	msg.fm.g.sender = t.Result[0].From
	msg.fm.g.date = t.Result[0].Date
	msg.fm.g.chat = t.Result[0].Chat
	msg.fm.g.replyed = &get{chat: t.Result[0].ReplyToMessage.Chat, sender: t.Result[0].From,
		date: t.Result[0].ReplyToMessage.Date, msgID: t.Result[0].ReplyToMessage.MessageID,
		msgOrigin: t.Result[0].ReplyToMessage.ForwardOrigin, photo: t.Result[0].ReplyToMessage.Photo,
		audio: t.Result[0].ReplyToMessage.Audio, document: t.Result[0].ReplyToMessage.Document,
		video: t.Result[0].ReplyToMessage.Video, anim: t.Result[0].ReplyToMessage.Animation,
		voice: t.Result[0].ReplyToMessage.Voice, vdn: t.Result[0].ReplyToMessage.VideoNote,
		paid: t.Result[0].ReplyToMessage.PaidMedia.PaidMedia[0], poll: t.Result[0].ReplyToMessage.Poll,
		dice: t.Result[0].ReplyToMessage.Dice}
	msg.fm.g.msgIDs = make([]int, len(t.Result))
	for i, m := range t.Result {
		msg.fm.g.msgIDs[i] = m.MessageID
		if m.Photo != nil {
			msg.fm.g.mg.photos = append(msg.fm.g.mg.photos, m.Photo)
		}
		if m.Audio != nil {
			msg.fm.g.mg.audios = append(msg.fm.g.mg.audios, m.Audio)
		}
		if m.Document != nil {
			msg.fm.g.mg.docs = append(msg.fm.g.mg.docs, m.Document)
		}
		if m.Video != nil {
			msg.fm.g.mg.videos = append(msg.fm.g.mg.videos, m.Video)
		}
	}
}

func responseDecoder(msg *Message) error {
	var err error
	switch t := msg.fm.tgr.(type) {
	case *types.SimpleResponse:
		msg.fm.g.status = t.Result
	case *types.MessageIDsResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.msgIDs = make([]int, len(t.Result))
		for i, id := range t.Result {
			msg.fm.g.msgIDs[i] = id.MessageID
		}
	case *types.MessageResponse:
		handlerMessage(msg, t)
	case *types.InviteLinkResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.invlink = t.Result
	case *types.MediaGroupResponse:
		handlerMediaGroup(msg, t)
	case *types.UserProfilePhotosResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.uprph = t.Result
	case *types.FileResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.file = t.Result
	case *types.ChatFullInfoResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.chatinfo = t.Result
	case *types.ChatMembersResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.members = t.Result
	case *types.ChatMemberResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.members = make([]*types.ChatMember, 1)
		msg.fm.g.members[0] = t.Result
	case *types.StickersResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.stickers = make([]*types.Sticker, len(t.Result))
		msg.fm.g.stickers = t.Result
	case *types.ForumResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.forum = t.Result
	case *types.UserBoostsResponse:
		msg.fm.g.status = t.Ok
		if t.Result != nil {
			msg.fm.g.boosts = t.Result.Boosts
		}
	case *types.BusinessConResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.bconn = t.Result
	case *types.BotCommandResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.botcomm = t.Result
	case *types.BotNameResponse:
		msg.fm.g.status = t.Ok
		if t.Result != nil {
			msg.fm.g.str = t.Result.Name
		}
	case *types.BotDescriptionResponse:
		msg.fm.g.status = t.Ok
		if t.Result != nil {
			msg.fm.g.str = t.Result.Description
		}
	case *types.BotShorDescriptionResponse:
		msg.fm.g.status = t.Ok
		if t.Result != nil {
			msg.fm.g.str = t.Result.ShortDescription
		}
	case *types.MenuButtonResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.menuButton = t.Result
	case *types.AdminRightResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.admin = t.Result
	case *types.PollResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.poll = t.Result
	case *types.StickerSetResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.stickerset = t.Result
	case *types.GiftsResponse:
		msg.fm.g.status = t.Ok
		// msg.fm.g.gift = t.Result
	case *types.WepAppMsgResponse:
		msg.fm.g.status = t.Ok
		// msg.fm.g.wepappmsg = t.Result
	case *types.PreparedInlineMessageResponse:
		msg.fm.g.status = t.Ok
		// msg.fm.g.inmsg = t.Result
	case *types.StringResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.str = t.Result
	case *types.IntResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.integer = &t.Result
	case *types.StarTransactionResponse:
		msg.fm.g.status = t.Ok
		// msg.fm.g.startrans = t.Result
	case *types.GameHighScoresResponse:
		msg.fm.g.status = t.Ok
		// msg.fm.g.score = t.Result
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
	url := fmt.Sprint(types.TelegramAPI, "bot", msg.fm.token, "/", msg.fm.method)
	req, err := http.NewRequest("POST", url, msg.fm.buf)
	req.Header.Set("Content-Type", msg.fm.contentType)

	if err == nil {
		client := http.Client{}
		resp, err = client.Do(req)

		if err == nil {
			defer resp.Body.Close()
			body, err = io.ReadAll(resp.Body)

			if err == nil {
				fmt.Println(string(body))
				errResp := new(types.Error)
				err = json.Unmarshal(body, errResp)
				if !errResp.Ok {
					msg.fm.g.status, msg.fm.g.errorCode, msg.fm.g.errorMsg = errResp.Ok, errResp.ErrorCode, errResp.Description
					err = code22()
				} else if err == nil {
					if err = json.Unmarshal(body, msg.fm.tgr); err == nil {
						err = responseDecoder(msg)
					}
				}
			}
		}
	}
	return err
}

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
