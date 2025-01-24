package formatter

import (
	"bytes"
	"mime/multipart"
)

type mediaHolder struct {
	storage     [10]handlerMedia
	i           int
	amount      int
	atLeastOnce bool
}

type formatter struct {
	inf         *information
	ch          *chat
	kb          *keyboard
	loc         *location
	con         *contact
	poll        *poll
	link        *link
	sticker     *sticker
	forum       *forum
	bot         *bot
	inlinemode  *inlinemode
	payment     *payment
	game        *game
	contentType string
	writer      *multipart.Writer
	mh          *mediaHolder
	buf         *bytes.Buffer
	method      string
	notchange   bool
	g           *get
	tgr         interface{}
}

type Message struct {
	fm *formatter
}
