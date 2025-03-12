package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"

	"github.com/iamissahar/Fmtogram/formatter/methods"
	"github.com/iamissahar/Fmtogram/types"
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

func checkAndAdd(values *url.Values, tofvalfield reflect.StructField, value string) {
	tag := tofvalfield.Tag.Get("field")
	if tag != "" {
		values.Add(tag, value)
	}
}

func handleValue(values *url.Values, valfield reflect.Value, tofvalfield reflect.StructField) {
	ifempty := tofvalfield.Tag.Get("ifempty")
	switch valfield.Kind() {
	case reflect.String:
		str := valfield.String()
		if ifempty == "true" {
			checkAndAdd(values, tofvalfield, str)
		} else if ifempty == "false" && str != "" {
			checkAndAdd(values, tofvalfield, str)
		}
	case reflect.Int:
		integer := valfield.Int()
		if ifempty == "true" {
			checkAndAdd(values, tofvalfield, fmt.Sprint(integer))
		} else if ifempty == "false" && integer != 0 {
			checkAndAdd(values, tofvalfield, fmt.Sprint(integer))
		}
	case reflect.Float64:
		float := valfield.Float()
		if ifempty == "true" {
			checkAndAdd(values, tofvalfield, fmt.Sprintf("%f", float))
		} else if ifempty == "false" && float != 0 {
			checkAndAdd(values, tofvalfield, fmt.Sprintf("%f", float))
		}
	case reflect.Bool:
		boolean := valfield.Bool()
		if ifempty == "true" {
			checkAndAdd(values, tofvalfield, fmt.Sprintf("%t", boolean))
		} else if ifempty == "false" && boolean {
			checkAndAdd(values, tofvalfield, fmt.Sprintf("%t", boolean))
		}
	case reflect.Slice, reflect.Array:
		if valfield.Kind() != reflect.Array {
			if !valfield.IsNil() {
				data, err := json.Marshal(valfield.Interface())
				if err == nil {
					checkAndAdd(values, tofvalfield, string(data))
				}
			}
		} else {
			data, err := json.Marshal(valfield.Interface())
			if err == nil {
				checkAndAdd(values, tofvalfield, string(data))
			}
		}
	case reflect.Struct:
		data, err := json.Marshal(valfield.Interface())
		if err == nil {
			checkAndAdd(values, tofvalfield, string(data))
		}
	case reflect.Pointer:
		if !valfield.IsNil() {
			valfield = valfield.Elem()
			handleValue(values, valfield, tofvalfield)
		}
	case reflect.Interface:
		if !valfield.IsNil() {
			data, err := json.Marshal(valfield.Interface())
			if err == nil {
				checkAndAdd(values, tofvalfield, string(data))
			}
		}
	}
}

func reflectStructToQuery(values *url.Values, val reflect.Value, tofval reflect.Type, queryString *string) {
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			valfield := val.Field(i)
			tofvalfield := tofval.Field(i)
			handleValue(values, valfield, tofvalfield)
		}
		*queryString = values.Encode()
	} else if val.Kind() == reflect.Pointer {
		if !val.IsNil() {
			val = val.Elem()
			tofval = val.Type()
			reflectStructToQuery(values, val, tofval, queryString)
		}
	} else if val.Kind() == reflect.Map {
		for _, k := range val.MapKeys() {
			val := val.MapIndex(k)
			key := fmt.Sprintf("%v", k.Interface())
			switch val.Kind() {
			case reflect.Interface:
				if !val.IsNil() {
					data, err := json.Marshal(val.Interface())
					if err == nil {
						values.Add(key, string(data))
						*queryString = values.Encode()
					}
				}
			default:
				panic("unexpected type of map value")
			}
		}
	}
}

func putUrlValues(data interface{}) (string, error) {
	var (
		values      = url.Values{}
		queryString string
		err         error
	)
	if data != nil {
		val := reflect.ValueOf(data)
		tofval := reflect.TypeOf(data)
		reflectStructToQuery(&values, val, tofval, &queryString)
	}
	return queryString, err
}

func uniteStrings(query [14]string) (string, error) {
	var (
		val, v = url.Values{}, url.Values{}
		err    error
	)
	for i := 0; (i < len(query)) && (err == nil); i++ {
		if query[i] != "" {
			if v, err = url.ParseQuery(query[i]); err == nil {
				for key, values := range v {
					for _, value := range values {
						val.Add(key, value)
					}
				}
			}
		}
	}
	return val.Encode(), err
}

func (mh *mediaHolder) single(wr *multipart.Writer, buf *bytes.Buffer, contenttype *string) (string, error) {
	var (
		queryStr string
		err      error
		jsn      []byte
	)
	if mh.amount != 0 {
		if mh.atLeastOnce {
			if err = mh.storage[mh.i-1].proccessFile(wr, contenttype, false); err == nil {
				queryStr, err = putUrlValues(mh.storage[mh.i-1])
			}
		} else {
			if jsn, err = json.Marshal(mh.storage[mh.i-1]); err == nil {
				buf.Write(jsn)
			}
		}
	}
	return queryStr, err
}

func (mh *mediaHolder) multiple(wr *multipart.Writer, buf *bytes.Buffer, contenttype *string) (string, error) {
	var (
		queryStr string
		err      error
		jsn      []byte
	)
	if mh.amount != 0 {
		group := make([]interface{}, mh.i)
		mh.m = make(map[string]interface{}, 1)
		for i := 0; (i < len(mh.storage)) && (err == nil); i++ {
			if mh.storage[i] != nil {
				if mh.atLeastOnce {
					err = mh.storage[i].proccessFile(wr, contenttype, true)
				}
				group[i] = mh.storage[i]
			}
		}
		mh.m["media"] = group
		if mh.atLeastOnce {
			queryStr, err = putUrlValues(mh.m)
		} else {
			if jsn, err = json.Marshal(mh.m); err == nil {
				buf.Write(jsn)
			}
		}
	}
	return queryStr, err
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

func (sh *stickerHolder) single(wr *multipart.Writer, buf *bytes.Buffer, contenttype *string) (string, error) {
	var (
		queryStr string
		err      error
		jsn      []byte
	)
	if sh.amount != 0 {
		if sh.atLeastOnce {
			err = sh.storage[sh.i-1].proccessFile(wr, contenttype, false)
		} else {
			if jsn, err = json.Marshal(sh.storage[sh.i-1]); err == nil {
				buf.Write(jsn)
			}
		}

	}
	return queryStr, err
}

func (sh *stickerHolder) multiple(wr *multipart.Writer, buf *bytes.Buffer, contenttype *string) (string, error) {
	var (
		err      error
		queryStr string
		jsn      []byte
	)
	if sh.amount != 0 {
		group := make([]interface{}, sh.i)
		sh.m = make(map[string]interface{}, 1)
		for i := 0; (i < len(sh.storage)) && (err == nil); i++ {
			if sh.storage[i] != nil {
				if sh.atLeastOnce {
					err = sh.storage[i].proccessFile(wr, contenttype, true)
				}
				group[i] = sh.storage[i]
			}
		}
		sh.m["stickers"] = group
		if sh.atLeastOnce {
			queryStr, err = putUrlValues(sh.m)
		} else {
			if jsn, err = json.Marshal(sh.m); err == nil {
				buf.Write(jsn)
			}
		}
	}
	return queryStr, err
}

func travers(msg *Message) (string, error) {
	var (
		fields = [12]interface{}{msg.fm.inf, msg.fm.ch, msg.fm.kb, msg.fm.loc, msg.fm.con, msg.fm.poll,
			msg.fm.link, msg.fm.forum, msg.fm.bot, msg.fm.inlinemode, msg.fm.payment, msg.fm.game}
		media    = [2]handleFiles{msg.fm.mh, msg.fm.sticker}
		err      error
		queryStr string
		cntr     int
		queryArr [len(fields) + len(media)]string
		jsn      []byte
	)

	for i := 0; i < len(fields) && err == nil; i++ {
		if fields[i] != nil {
			switch fields[i].(type) {
			case *information:
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
			}
			if msg.fm.mh.atLeastOnce || msg.fm.sticker.atLeastOnce {
				queryArr[cntr], err = putUrlValues(fields[i])
				cntr++
			} else {
				if jsn, err = json.Marshal(fields[i]); err == nil {
					msg.fm.buf.Write(jsn)
				}
			}
		}
	}

	for i := 0; i < len(media) && err == nil; i++ {
		if media[i] != nil {
			if (msg.fm.mh.amount == 1 && msg.fm.method != methods.PaidMedia) ||
				(msg.fm.sticker.amount == 1 && msg.fm.method != methods.CreateNewStickerSet) {
				queryArr[cntr], err = media[i].single(msg.fm.writer, msg.fm.buf, &msg.fm.contentType)
				cntr++
			} else {
				queryArr[cntr], err = media[i].multiple(msg.fm.writer, msg.fm.buf, &msg.fm.contentType)
				cntr++
			}
		}
	}

	if err == nil {
		queryStr, err = uniteStrings(queryArr)
	}
	return queryStr, err
}

func makeRequest(msg *Message) (string, error) {
	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.writer = multipart.NewWriter(msg.fm.buf)

	if msg.fm.contentType == "" {
		msg.fm.contentType = "application/json"
	}
	return travers(msg)
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
	if msg.fm.method == methods.PaidMedia {
		msg.fm.g.photo = t.Result.PaidMedia.PaidMedia[0].Photo
		msg.fm.g.video = t.Result.PaidMedia.PaidMedia[0].Video
	} else {
		msg.fm.g.photo = t.Result.Photo
		msg.fm.g.video = t.Result.Video
	}
	msg.fm.g.audio = t.Result.Audio
	msg.fm.g.document = t.Result.Document
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
	if t.Result[0].ReplyToMessage != nil {
		msg.fm.g.replyed = &get{chat: t.Result[0].ReplyToMessage.Chat, sender: t.Result[0].From,
			date: t.Result[0].ReplyToMessage.Date, msgID: t.Result[0].ReplyToMessage.MessageID,
			msgOrigin: t.Result[0].ReplyToMessage.ForwardOrigin, photo: t.Result[0].ReplyToMessage.Photo,
			audio: t.Result[0].ReplyToMessage.Audio, document: t.Result[0].ReplyToMessage.Document,
			video: t.Result[0].ReplyToMessage.Video, anim: t.Result[0].ReplyToMessage.Animation,
			voice: t.Result[0].ReplyToMessage.Voice, vdn: t.Result[0].ReplyToMessage.VideoNote,
			paid: t.Result[0].ReplyToMessage.PaidMedia.PaidMedia[0], poll: t.Result[0].ReplyToMessage.Poll,
			dice: t.Result[0].ReplyToMessage.Dice}
	}
	msg.fm.g.msgs = t.Result
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
	case *types.MessageIDResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.msgID = t.Result.MessageID
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
		msg.fm.g.mg = new(mediagroup)
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
		if t.Result != nil && t.Result.G != nil {
			msg.fm.g.gifts = t.Result.G
		}
	case *types.WepAppMsgResponse:
		msg.fm.g.status = t.Ok
		if t.Result != nil {
			msg.fm.g.str = t.Result.InlineMessageID
		}
	case *types.PreparedInlineMessageResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.prepinlmsg = t.Result
	case *types.StringResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.str = t.Result
	case *types.IntResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.integer = &t.Result
	case *types.StarTransactionResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.startrans = t.Result
	case *types.GameHighScoresResponse:
		msg.fm.g.status = t.Ok
		msg.fm.g.score = t.Result
	}
	return err
}

func sendRequest(msg *Message, query string) error {
	var resp *http.Response
	var body []byte
	var err error
	var req *http.Request
	if !msg.fm.mh.atLeastOnce && !msg.fm.sticker.atLeastOnce {
		err = uniteEverything(msg)
	} else {
		err = msg.fm.writer.Close()
	}

	msg.fm.g.request = fmt.Sprintf("content-type=%s\nmethod=%s\nhttp-method=%s\nbot-token=%s\n",
		msg.fm.contentType, msg.fm.method, msg.fm.httpMethod, msg.fm.token)
	if msg.fm.contentType == "application/json" {
		msg.fm.g.request = fmt.Sprintf("%s\n%s", msg.fm.g.request, msg.fm.buf.String())
	} else {
		msg.fm.g.request = fmt.Sprintf("%s\nurl=%s", msg.fm.g.request, query)
	}

	if err == nil {
		if query != "" {
			query = "?" + query
		}
		url := fmt.Sprint(types.TelegramAPI, "bot", msg.fm.token, "/", msg.fm.method, query)
		req, err = http.NewRequest(msg.fm.httpMethod, url, msg.fm.buf)
		req.Header.Set("Content-Type", msg.fm.contentType)

		if err == nil {
			client := http.Client{}
			resp, err = client.Do(req)

			if err == nil {
				defer resp.Body.Close()
				body, err = io.ReadAll(resp.Body)

				if err == nil {
					errResp := new(types.Error)
					err = json.Unmarshal(body, errResp)
					msg.fm.g.response = string(body)
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
	}
	return err
}

func (msg *Message) Send() error {
	var err error
	var query string

	if query, err = makeRequest(msg); err == nil {
		if msg.fm.tgr == nil {
			panic("!~")
		}
		err = sendRequest(msg, query)
	}

	return err
}
