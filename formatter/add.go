package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Added a photo interface to the message you're building
func (msg *Message) AddPhoto(ph IPhoto) {
	p := ph.(*photo)
	if p.GottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = p
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfacePhoto)
}

// Added a video interface to the message you're building
func (msg *Message) AddVideo(vd IVideo) {
	v := vd.(*video)
	if v.VideoGottenFrom == Storage || v.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = v
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfaceVideo)
}

// Added a document interface to the message you're building
func (msg *Message) AddDocument(dc IDocument) {
	d := dc.(*document)
	if d.DocumentGottenFrom == Storage || d.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = d
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfaceDocument)
}

// Added an animation interface to the message you're building
func (msg *Message) AddAnimation(an IAnimation) {
	a := an.(*animation)
	if a.AnimationGottenFrom == Storage || a.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = a
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfaceAnimation)
}

// Added an audio interface to the message you're building
func (msg *Message) AddAudio(ad IAudio) {
	a := ad.(*audio)
	if a.AudioGottenFrom == Storage || a.ThumbnailGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = a
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfaceAudio)
}

// Added a voice interface to the message you're building
func (msg *Message) AddVoice(vc IVoice) {
	v := vc.(*voice)
	if v.gottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = v
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfaceVoice)
}

// Added a voice interface to the message you're building
func (msg *Message) AddVideoNote(vdn IVideoNote) {
	v := vdn.(*videonote)
	if v.videoGottenFrom == Storage {
		msg.fm.mh.atLeastOnce = true
	}
	msg.fm.mh.storage[msg.fm.mh.i] = v
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved(interfaceVideoNote)
}

// Added a message interface to the message you're building
func (msg *Message) AddMessage(inf IMSGInformation) {
	msg.fm.inf = inf.(*information)
	logs.InterfaceSaved(interfaceInf)
}

// Added a chat interface to the message you're building
func (msg *Message) AddChat(ch IChat) {
	msg.fm.ch = ch.(*chat)
	logs.InterfaceSaved(interfaceChat)
}

// Add a reply-keyboard interface to the message you're building
func (msg *Message) AddReplyKeyboard(kb IReply) {
	msg.fm.kb = kb.(*reply)
	logs.InterfaceSaved(interfaceReplyKB)
}

// Add a inline-keyboard interface to the message you're building
func (msg *Message) AddInlineKeyboard(kb IInline) {
	msg.fm.kb = kb.(*inline)
	logs.InterfaceSaved(interfaceInKB)
}

func (msg *Message) AddMethod(method string) {
	msg.fm.method = method
	msg.fm.notchange = true
	logs.MethodSaved(method)
}
