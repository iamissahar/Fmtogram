package fmtogram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"

	"github.com/iamissahar/Fmtogram/methods"
)

// func CreateMessage(botID string) *Message {
// 	m := new(Message)
// 	m.fm = new(formatter)
// 	m.fm.ch = &chat{}
// 	m.fm.prm = &information{}
// 	m.fm.mh = &mediaHolder{}
// 	m.fm.sticker = &stickerHolder{}
// 	m.fm.token = botID
// 	return m
// }

// func CreateEmpltyMessage() *Message {
// 	m := new(Message)
// 	m.fm = new(formatter)
// 	m.fm.ch = &chat{}
// 	m.fm.prm = &information{}
// 	m.fm.mh = &mediaHolder{}
// 	m.fm.sticker = &stickerHolder{}
// 	m.fm.token = BotID
// 	return m
// }

func CreateWebHook() IWebHook {
	return &webhook{}
}

func (wh *webhook) send(method, query string) error {
	var (
		resp *http.Response
		body []byte
	)
	if query != "" {
		query = "?" + query
	}
	url := fmt.Sprint(TelegramAPI, "bot", wh.token, "/", method, query)
	req, err := http.NewRequest(http.MethodPost, url, wh.buf)
	req.Header.Set("Content-Type", wh.contenttype)
	if err == nil {
		cl := http.Client{}
		resp, err = cl.Do(req)
	}
	if err == nil {
		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		if err == nil {
			errResp := new(Error)
			err = json.Unmarshal(body, errResp)
			if !errResp.Ok {
				wh.status, wh.errorcode, wh.errormsg = errResp.Ok, errResp.ErrorCode, errResp.Description
				err = code22()
			} else if err == nil {
				if method == "setWebhook" || method == "deleteWebhook" {
					sr := new(SimpleResponse)
					err = json.Unmarshal(body, sr)
					if err == nil {
						wh.status = sr.Ok
						if !sr.Result {
							err = code22()
						}
					}
				} else {
					whi := new(WebhookResponse)
					err = json.Unmarshal(body, whi)
					wh.status, wh.info = whi.Ok, whi.Result
				}
			}
		}
	}
	return err
}

func (wh *webhook) Set() error {
	var (
		body  []byte
		err   error
		query string
	)
	wh.buf = bytes.NewBuffer(nil)
	wr := multipart.NewWriter(wh.buf)
	if wh.Certificate != "" {
		if err = wh.certificateJob(wr); err == nil {
			query, err = putUrlValues(wh)
			if err == nil {
				err = wr.Close()
			}
		}
	} else {
		body, err = json.Marshal(wh)
		wh.buf.Write(body)
		wh.contenttype = "application/json"
	}
	if err == nil {
		err = wh.send("setWebhook", query)
	}
	return err
}

func (wh *webhook) Delete() error {
	var query string
	if wh.DropPendingUpdates {
		val := url.Values{}
		val.Add("drop_pending_updates", "true")
		query = val.Encode()
	}
	return wh.send("deleteWebhook", query)
}

func (wh *webhook) Info() (*WebhookInfo, error) {
	return wh.info, wh.send("getWebhookInfo", "")
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

func uniteStrings(query [15]string) (string, error) {
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
		fields = [13]interface{}{msg.fm.prm, msg.fm.ch, msg.fm.kb, msg.fm.loc, msg.fm.con, msg.fm.poll,
			msg.fm.link, msg.fm.forum, msg.fm.bot, msg.fm.inlinemode, msg.fm.payment, msg.fm.game, msg.fm.gift}
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
					if msg.fm.prm.AdminRights != nil {
						msg.fm.prm.IsAnonymous = msg.fm.prm.AdminRights.IsAnonymous
						msg.fm.prm.CanManageChat = msg.fm.prm.AdminRights.CanManageChat
						msg.fm.prm.CanDeleteMessages = msg.fm.prm.AdminRights.CanDeleteMessages
						msg.fm.prm.CanManageVideoChats = msg.fm.prm.AdminRights.CanManageVideoChats
						msg.fm.prm.CanRestrictMembers = msg.fm.prm.AdminRights.CanRestrictMembers
						msg.fm.prm.CanPromoteMembers = msg.fm.prm.AdminRights.CanPromoteMembers
						msg.fm.prm.CanChangeInfo = msg.fm.prm.AdminRights.CanChangeInfo
						msg.fm.prm.CanInviteUsers = msg.fm.prm.AdminRights.CanInviteUsers
						msg.fm.prm.CanPostStories = msg.fm.prm.AdminRights.CanPostStories
						msg.fm.prm.CanEditStories = msg.fm.prm.AdminRights.CanEditStories
						msg.fm.prm.CanDeleteStories = msg.fm.prm.AdminRights.CanDeleteStories
						msg.fm.prm.CanPostMessages = msg.fm.prm.AdminRights.CanPostMessages
						msg.fm.prm.CanEditMessages = msg.fm.prm.AdminRights.CanEditMessages
						msg.fm.prm.CanPinMessages = msg.fm.prm.AdminRights.CanPinMessages
						msg.fm.prm.CanManageTopics = msg.fm.prm.AdminRights.CanManageTopics
						msg.fm.prm.AdminRights = nil
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

func handlerMessage(msg *Message, t *MessageResponse) {
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
		msg.fm.g.stickers = []*Sticker{t.Result.Sticker}
	}
	msg.fm.g.msg = t.Result
}

func handlerMediaGroup(msg *Message, t *MediaGroupResponse) {
	msg.fm.g.status = t.Ok
	msg.fm.g.mg.id = t.Result[0].MediaGroupID
	msg.fm.g.sender = t.Result[0].From
	msg.fm.g.date = t.Result[0].Date
	msg.fm.g.chat = t.Result[0].Chat
	if t.Result[0].ReplyToMessage != nil {
		g := &get{chat: t.Result[0].ReplyToMessage.Chat, sender: t.Result[0].From,
			date: t.Result[0].ReplyToMessage.Date, msgID: t.Result[0].ReplyToMessage.MessageID,
			msgOrigin: t.Result[0].ReplyToMessage.ForwardOrigin, photo: t.Result[0].ReplyToMessage.Photo,
			audio: t.Result[0].ReplyToMessage.Audio, document: t.Result[0].ReplyToMessage.Document,
			video: t.Result[0].ReplyToMessage.Video, anim: t.Result[0].ReplyToMessage.Animation,
			voice: t.Result[0].ReplyToMessage.Voice, vdn: t.Result[0].ReplyToMessage.VideoNote,
			poll: t.Result[0].ReplyToMessage.Poll, dice: t.Result[0].ReplyToMessage.Dice}
		if t.Result[0].ReplyToMessage.PaidMedia != nil {
			g.paid = t.Result[0].ReplyToMessage.PaidMedia.PaidMedia[0]
		}
		msg.fm.g.replyed = g
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
	if msg.fm.g != nil {
		switch t := msg.fm.tgr.(type) {
		case *SimpleResponse:
			msg.fm.g.status = t.Result
		case *MessageIDResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.msgID = t.Result.MessageID
		case *MessageIDsResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.msgIDs = make([]int, len(t.Result))
			for i, id := range t.Result {
				msg.fm.g.msgIDs[i] = id.MessageID
			}
		case *MessageResponse:
			handlerMessage(msg, t)
		case *InviteLinkResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.invlink = t.Result
		case *MediaGroupResponse:
			msg.fm.g.mg = new(mediagroup)
			handlerMediaGroup(msg, t)
		case *UserProfilePhotosResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.uprph = t.Result
		case *FileResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.file = t.Result
		case *ChatFullInfoResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.chatinfo = t.Result
		case *ChatMembersResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.members = t.Result
		case *ChatMemberResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.members = make([]*ChatMember, 1)
			msg.fm.g.members[0] = t.Result
		case *StickersResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.stickers = make([]*Sticker, len(t.Result))
			msg.fm.g.stickers = t.Result
		case *ForumResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.forum = t.Result
		case *UserBoostsResponse:
			msg.fm.g.status = t.Ok
			if t.Result != nil {
				msg.fm.g.boosts = t.Result.Boosts
			}
		case *BusinessConResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.bconn = t.Result
		case *BotCommandResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.botcomm = t.Result
		case *BotNameResponse:
			msg.fm.g.status = t.Ok
			if t.Result != nil {
				msg.fm.g.str = t.Result.Name
			}
		case *BotDescriptionResponse:
			msg.fm.g.status = t.Ok
			if t.Result != nil {
				msg.fm.g.str = t.Result.Description
			}
		case *BotShorDescriptionResponse:
			msg.fm.g.status = t.Ok
			if t.Result != nil {
				msg.fm.g.str = t.Result.ShortDescription
			}
		case *MenuButtonResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.menuButton = t.Result
		case *AdminRightResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.admin = t.Result
		case *PollResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.poll = t.Result
		case *StickerSetResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.stickerset = t.Result
		case *GiftsResponse:
			msg.fm.g.status = t.Ok
			if t.Result != nil && t.Result.G != nil {
				msg.fm.g.gifts = t.Result.G
			}
		case *WepAppMsgResponse:
			msg.fm.g.status = t.Ok
			if t.Result != nil {
				msg.fm.g.str = t.Result.InlineMessageID
			}
		case *PreparedInlineMessageResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.prepinlmsg = t.Result
		case *StringResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.str = t.Result
		case *IntResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.integer = &t.Result
		case *StarTransactionResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.startrans = t.Result
		case *GameHighScoresResponse:
			msg.fm.g.status = t.Ok
			msg.fm.g.score = t.Result
		}
	}
	return err
}

func sendRequest2(msg *Message, query, token string) error {
	var resp *http.Response
	var body []byte
	var err error
	var req *http.Request
	if !msg.fm.mh.atLeastOnce && !msg.fm.sticker.atLeastOnce {
		err = uniteEverything(msg)
	} else {
		err = msg.fm.writer.Close()
	}
	if msg.fm.g != nil {
		msg.fm.g.request = fmt.Sprintf("content-type=%s\nmethod=%s\nhttp-method=%s\nbot-token=%s\n",
			msg.fm.contentType, msg.fm.method, msg.fm.httpMethod, token)
	}
	if msg.fm.contentType == "application/json" {
		if msg.fm.g != nil {
			msg.fm.g.request = fmt.Sprintf("%s\n%s", msg.fm.g.request, msg.fm.buf.String())
		}
	} else {
		if msg.fm.g != nil {
			msg.fm.g.request = fmt.Sprintf("%s\nurl=%s", msg.fm.g.request, query)
		}
	}

	if err == nil {
		if query != "" {
			query = "?" + query
		}
		url := fmt.Sprint(TelegramAPI, "bot", token, "/", msg.fm.method, query)
		req, err = http.NewRequest(msg.fm.httpMethod, url, msg.fm.buf)
		req.Header.Set("Content-Type", msg.fm.contentType)

		if err == nil {
			client := http.Client{}
			resp, err = client.Do(req)

			if err == nil {
				defer resp.Body.Close()
				body, err = io.ReadAll(resp.Body)

				if err == nil {
					errResp := new(Error)
					err = json.Unmarshal(body, errResp)
					if msg.fm.g != nil {
						msg.fm.g.response = string(body)
					}
					if !errResp.Ok {
						if msg.fm.g != nil {
							msg.fm.g.status, msg.fm.g.errorCode, msg.fm.g.errorMsg = errResp.Ok, errResp.ErrorCode, errResp.Description
						}
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

func (bs *BasicSettings) Send(msg *Message, chatID interface{}) error {
	var (
		err   error
		query string
	)
	switch c := chatID.(type) {
	case int, string:
		msg.fm.ch.ID = c
	default:
		err = code20()
	}
	if msg.fm.tgr == nil {
		msg.fm.tgr = new(MessageResponse)
	}
	if msg.fm.method == "" {
		msg.fm.method = methods.Message
	}
	if err == nil {
		if query, err = makeRequest(msg); err == nil {
			if msg.fm.tgr == nil {
				panic("!~")
			}
			err = sendRequest2(msg, query, bs.Token)
		}
	}

	return err
}
