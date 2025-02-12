package formatter

import "github.com/l1qwie/Fmtogram/types"

func (g *get) Status() bool {
	return g.status
}

func (g *get) Error() (int, string) {
	return g.errorCode, g.errorMsg
}

func (g *get) Chat() *types.Chat {
	return g.chat
}

func (g *get) Date() int {
	return g.date
}

func (g *get) MessageID() int {
	return g.msgID
}

func (g *get) MessageIDs() []int {
	return g.msgIDs
}

func (g *get) Sender() *types.User {
	return g.sender
}

func (g *get) Replyed() IGet {
	return g.replyed
}

func (g *get) ForwardOrigin() *types.MessageOrigin {
	return g.msgOrigin
}

func (g *get) Photo() []*types.PhotoSize {
	return g.photo
}

func (g *get) Audio() *types.Audio {
	return g.audio
}

func (g *get) Document() *types.Document {
	return g.document
}

func (g *get) Video() *types.Video {
	return g.video
}

func (g *get) Animation() *types.Animation {
	return g.anim
}

func (g *get) Voice() *types.Voice {
	return g.voice
}

func (g *get) VideoNote() *types.VideoNote {
	return g.vdn
}

func (g *get) PaidMedia() *types.PaidMedia {
	return g.paid
}

func (g *get) MediaGroupID() string {
	return g.mg.id
}

func (g *get) Photos() [][]*types.PhotoSize {
	return g.mg.photos
}

func (g *get) Videos() []*types.Video {
	return g.mg.videos
}

func (g *get) Audios() []*types.Audio {
	return g.mg.audios
}

func (g *get) Documents() []*types.Document {
	return g.mg.docs
}

func (g *get) Poll() *types.Poll {
	return g.poll
}

func (g *get) Dice() *types.Dice {
	return g.dice
}

func (g *get) ProfilePhotos() *types.UserProfilePhotos {
	return g.uprph
}

func (g *get) File() *types.File {
	return g.file
}

func (g *get) Stickers() []*types.Sticker {
	return g.stickers
}

func (g *get) Gifts() []*types.Gift {
	return g.gifts
}

func (g *get) Message() *types.Message {
	return g.msg
}

func (g *get) String() string {
	return g.str
}

func (g *get) InviteLink() *types.ChatInviteLink {
	return g.invlink
}

func (g *get) ChatInfo() *types.ChatFullInfo {
	return g.chatinfo
}

func (g *get) Members() []*types.ChatMember {
	return g.members
}

func (g *get) Integer() *int {
	return g.integer
}

func (g *get) Forum() *types.ForumTopic {
	return g.forum
}
