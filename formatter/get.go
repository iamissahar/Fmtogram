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

func (g *get) Bot() types.User {
	return g.bot
}

func (g *get) Replyed() IGet {
	return g.replyed
}

func (g *get) ForwardOrigin() types.MessageOrigin {
	return g.msgOrigin
}
