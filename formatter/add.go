package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Added a photo to the message you're building
func (msg *Message) AddPhoto(ph IPhoto) {
	p := ph.(*photo)
	if p.GottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = p
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Photo")
}

// Added a video to the message you're building
func (msg *Message) AddVideo(vd IVideo) {
	v := vd.(*video)
	if v.VideoGottenFrom == Storage || v.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = v
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Video")
}

// Added a document to the message you're building
func (msg *Message) AddDocument(dc IDocument) {
	d := dc.(*document)
	if d.DocumentGottenFrom == Storage || d.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = d
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Document")
}

// Added an audio to the message you're building
func (msg *Message) AddAudio(ad IAudio) {
	a := ad.(*audio)
	if a.AudioGottenFrom == Storage || a.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = a
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Audio")
}

// Added a message to the message you're building
func (msg *Message) AddMessage(inf IMSGInformation) {
	msg.fm.inf = inf.(*information)
	logs.InterfaceSaved("Message")
}

// Added a chat to the message you're building
func (msg *Message) AddChat(ch IChat) {
	msg.fm.ch = ch.(*chat)
	logs.InterfaceSaved("Chat")
}

// Add a reply-keyboard to the message you're building
func (msg *Message) AddReplyKeyboard(kb IReply) {
	msg.fm.kb = kb.(*reply)
	logs.InterfaceSaved("Reply Keyboard")
}

// Add a inline-keyboard to the message you're building
func (msg *Message) AddInlineKeyboard(kb IInline) {
	msg.fm.kb = kb.(*inline)
	logs.InterfaceSaved("Reply Keyboard")
}

func (msg *Message) AddMethod(method string) {
	msg.fm.method = method
	msg.fm.notchange = true
	logs.MethodSaved(method)
}
