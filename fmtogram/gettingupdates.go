package fmtogram

import (
	"reflect"
	"time"

	"github.com/iamissahar/Fmtogram/formatter"
	"github.com/iamissahar/Fmtogram/types"
)

func (msg *Message) ID() int {
	return msg.MessageID
}

func (msg *Message) Sender() *types.User {
	return msg.From
}

func (msg *Message) Chat() *types.Chat {
	return msg.Ch
}

func (msg *Message) Date() (unix int, t time.Time) {
	return msg.DateUnix, time.Unix(int64(msg.DateUnix), 0).UTC()
}

func (msg *Message) Reply() formatter.IMessage {
	return msg.ReplyToMessage
}

func (msg *Message) Text() string {
	return msg.Txt
}

func (msg *Message) Audio() *types.Audio {
	return msg.Ad
}

func (msg *Message) Document() *types.Document {
	return msg.Doc
}

func (msg *Message) PaidMedia() *types.PaidMediaInfo {
	return msg.PaidM
}

func (msg *Message) Photo() []*types.PhotoSize {
	return msg.Ph
}

func (msg *Message) Sticker() *types.Sticker {
	return msg.Stckr
}

func (msg *Message) Story() *types.Story {
	return msg.Stry
}

func (msg *Message) Video() *types.Video {
	return msg.Vd
}

func (msg *Message) VideoNote() *types.VideoNote {
	return msg.Vdn
}

func (msg *Message) Voice() *types.Voice {
	return msg.Vc
}

func (msg *Message) Message() *types.Message {
	m := new(types.Message)
	msgVal := reflect.ValueOf(msg)
	anotherMsgVal := reflect.ValueOf(m).Elem()

	for i := 0; i < msgVal.NumField(); i++ {
		anotherMsgVal.Field(i).Set(msgVal.Field(i))
	}
	return m
}

func (bc *BusinessConnection) ID() string {
	return bc.BID
}

func (bc *BusinessConnection) Sender() *types.User {
	return bc.User
}

func (bc *BusinessConnection) SenderChatID() int64 {
	return bc.UserChatID
}

func (bc *BusinessConnection) Date() (unix int, t time.Time) {
	return bc.BDate, time.Unix(int64(bc.BDate), 0).UTC()
}

func (bc *BusinessConnection) CanReply() bool {
	return bc.BCanReply
}

func (bc *BusinessConnection) IsEnabled() bool {
	return bc.BIsEnabled
}

func (bmd *BusinessMessagesDeleted) ConnectionID() string {
	return bmd.BusinessConnectionID
}

func (bmd *BusinessMessagesDeleted) Chat() *types.Chat {
	return bmd.Ch
}

func (bmd *BusinessMessagesDeleted) MessageIDs() []int {
	return bmd.MsgIDs
}

func (mru *MessageReactionUpdated) Chat() *types.Chat {
	return mru.Ch
}

func (mru *MessageReactionUpdated) MessageID() int {
	return mru.MsgID
}

func (mru *MessageReactionUpdated) Sender() *types.User {
	return mru.User
}

func (mru *MessageReactionUpdated) ActorChat() *types.Chat {
	return mru.ActorCh
}

func (mru *MessageReactionUpdated) Date() (unix int, t time.Time) {
	return mru.DateUnix, time.Unix(int64(mru.DateUnix), 0).UTC()
}

func (mru *MessageReactionUpdated) OldReaction() []*types.ReactionType {
	return mru.OldReact
}

func (mru *MessageReactionUpdated) NewReaction() []*types.ReactionType {
	return mru.NewReact
}

func (mrcu *MessageReactionCountUpdated) Chat() *types.Chat {
	return mrcu.Ch
}

func (mrcu *MessageReactionCountUpdated) MessageID() int {
	return mrcu.MsgID
}

func (mrcu *MessageReactionCountUpdated) Date() (unix int, t time.Time) {
	return mrcu.DateUnix, time.Unix(int64(mrcu.DateUnix), 0).UTC()
}

func (mrcu *MessageReactionCountUpdated) Reactions() []*types.ReactionCount {
	return mrcu.React
}

func (iq *InlineQuery) ID() string {
	return iq.IqID
}

func (iq *InlineQuery) Sender() *types.User {
	return iq.From
}

func (iq *InlineQuery) Query() string {
	return iq.Q
}

func (iq *InlineQuery) Offset() string {
	return iq.IqOffset
}

func (iq *InlineQuery) ChatType() string {
	return iq.IqChatType
}

func (iq *InlineQuery) Location() *types.Location {
	return iq.IqLocation
}

func (cir *ChosenInlineResult) ID() string {
	return cir.ResultID
}

func (cir *ChosenInlineResult) Sender() *types.User {
	return cir.From
}

func (cir *ChosenInlineResult) Location() *types.Location {
	return cir.Loc
}

func (cir *ChosenInlineResult) InlineMessageID() string {
	return cir.InMsgID
}

func (cir *ChosenInlineResult) Query() string {
	return cir.Q
}

//
//
//

func (up *Update) ID() int {
	return up.UpdateID
}

func (up *Update) Message() formatter.IMessage {
	return up.Msg
}

func (up *Update) EditedMessage() formatter.IMessage {
	return up.EditedMsg
}

func (up *Update) ChannelPost() formatter.IMessage {
	return up.ChanPost
}

func (up *Update) EditedChannelPost() formatter.IMessage {
	return up.EditedChanPost
}

func (up *Update) BusinessConnection() formatter.IBusiness {
	return up.BusinessConn
}

func (up *Update) BusinessMessage() formatter.IMessage {
	return up.BusinessMsg
}

func (up *Update) EditedBusinessMessage() formatter.IMessage {
	return up.EditedBusinessMsg
}

func (up *Update) DeletedBusinessMessages() formatter.IBusinessMessage {
	return up.DeletedBusinessMessage
}

func (up *Update) MessageReaction() formatter.IMessageReaction {
	return up.MessageR
}

func (up *Update) MessageReactionCount() formatter.IMessageReactionCount {
	return up.MessageRcount
}

func (up *Update) InlineQuery() formatter.IInlineQuery {
	return up.InlineQ
}

func (up *Update) ChosenInlineResult() formatter.IInlineResult {
	return up.ChosenInlineR
}

func (up *Update) CallbackQuery() formatter.ICallbackQuery {
	return up.CallbackQ
}

func (up *Update) ShippingQuery() formatter.IShippingQuery {
	return up.ShippingQ
}

func (up *Update) PreCheckoutQuery() formatter.IPreCheckoutQuery {
	return up.PreCheckoutQ
}

func (up *Update) PurchasedPaidMedia() formatter.IPaidMedia {
	return up.PollUpd
}

func (up *Update) Poll() formatter.IPollUpdate {
	return up.PollAnswr
}

func (up *Update) PollAnswer() formatter.IPollAnswer {
	return up.MyChatMem
}

func (up *Update) MyChatMember() formatter.IChatMember {
	return up.ChatMem
}

func (up *Update) ChatMember() formatter.IChatMember {
	return up.ChatMem
}

func (up *Update) ChatJoinRequest() formatter.IJoinRequest {
	return up.ChatJoinReq
}

func (up *Update) ChatBoost() formatter.IBoostUpdate {
	return up.ChatBoostUpd
}

func (up *Update) RemovedChatBoost() formatter.IBoostRemove {
	return up.RemovedChatBoostUpd
}
