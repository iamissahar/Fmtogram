package fmtogram

import (
	"bytes"
	"mime/multipart"
)

type mediaHolder struct {
	amount      int
	storage     [10]handlerMedia
	i           int
	atLeastOnce bool
	m           map[string]interface{}
}

type stickerHolder struct {
	amount      int
	storage     [50]*sticker
	i           int
	atLeastOnce bool
	m           map[string]interface{}
}

type formatter struct {
	prm         *information
	ch          *chat
	kb          *keyboard
	loc         *location
	con         *contact
	poll        *poll
	link        *link
	sticker     *stickerHolder
	forum       *forum
	bot         *bot
	inlinemode  *inlinemode
	payment     *payment
	game        *game
	gift        *gift
	contentType string
	writer      *multipart.Writer
	mh          *mediaHolder
	buf         *bytes.Buffer
	method      string
	httpMethod  string
	notchange   bool
	g           *get
	tgr         interface{}
	token       string
}

type Message struct {
	fm *formatter
}
