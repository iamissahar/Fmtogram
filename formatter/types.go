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
	kb          handlerKB
	loc         *location
	con         *contact
	poll        *poll
	contentType string
	writer      *multipart.Writer
	mh          *mediaHolder
	buf         *bytes.Buffer
	method      string
	notchange   bool
}

type Message struct {
	fm *formatter
}
