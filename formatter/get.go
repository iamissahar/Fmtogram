package formatter

import "github.com/l1qwie/Fmtogram/types"

func (g *get) Status() bool {
	return g.status
}

func (g *get) Error() (int, string) {
	return g.errorCode, g.errorMsg
}

func (g *get) Chat() types.Chat {
	return g.chat
}

func (g *get) User() types.User {
	return g.user
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

func (g *get) Bot() types.User {
	return g.bot
}

func (g *get) Replyed() IGet {
	return g.replyed
}

func (g *get) ForwardOrigin() types.MessageOrigin {
	return g.msgOrigin
}

func (g *get) Photo() []types.PhotoSize {
	return g.photo
}

func (g *get) Audio() types.Audio {
	return g.audio
}

func (g *get) Document() types.Document {
	return g.document
}

func (g *get) Video() types.Video {
	return g.video
}

func (g *get) Animation() types.Animation {
	return g.anim
}
