package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Added a photo interface to the message you're building
func (msg *Message) AddPhoto(ph IPhoto) error {
	var err error
	if p, ok := ph.(*photo); ok {
		if p.Photo != "" {
			if p.gottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = p
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfacePhoto)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added a video interface to the message you're building
func (msg *Message) AddVideo(vd IVideo) error {
	var err error
	if v, ok := vd.(*video); ok {
		if v.Video != "" {
			if v.videoGottenFrom == Storage || v.thumbnailGottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = v
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfaceVideo)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added a document interface to the message you're building
func (msg *Message) AddDocument(dc IDocument) error {
	var err error
	if d, ok := dc.(*document); ok {
		if d.Document != "" {
			if d.documentGottenFrom == Storage || d.thumbnailGottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = d
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfaceDocument)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added an animation interface to the message you're building
func (msg *Message) AddAnimation(an IAnimation) error {
	var err error
	if a, ok := an.(*animation); ok {
		if a.Animation != "" {
			if a.animationGottenFrom == Storage || a.thumbnailGottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = a
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfaceAnimation)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added an audio interface to the message you're building
func (msg *Message) AddAudio(ad IAudio) error {
	var err error
	if a, ok := ad.(*audio); ok {
		if a.Audio != "" {
			if a.audioGottenFrom == Storage || a.thumbnailGottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = a
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfaceAudio)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added a voice interface to the message you're building
func (msg *Message) AddVoice(vc IVoice) error {
	var err error
	if v, ok := vc.(*voice); ok {
		if v.Voice != "" {
			if v.gottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = v
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfaceVoice)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added a voice interface to the message you're building
func (msg *Message) AddVideoNote(vdn IVideoNote) error {
	var err error
	if v, ok := vdn.(*videonote); ok {
		if v.VideoNote != "" {
			if v.videoGottenFrom == Storage {
				msg.fm.mh.atLeastOnce = true
			}
			msg.fm.mh.storage[msg.fm.mh.i] = v
			msg.fm.mh.i++
			msg.fm.mh.amount++
			logs.InterfaceSaved(interfaceVideoNote)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added a message interface to the message you're building
func (msg *Message) AddParameters(param IParameters) error {
	var err error
	if p, ok := param.(*information); ok {
		msg.fm.inf = p
		logs.InterfaceSaved(interfaceInf)
	} else {
		err = code20()
	}
	return err
}

// Added a chat interface to the message you're building
func (msg *Message) AddChat(ch IChat) error {
	var err error
	if c, ok := ch.(*chat); ok {
		if c.ID != nil {
			msg.fm.ch = c
			logs.InterfaceSaved(interfaceChat)
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Add a reply-keyboard interface to the message you're building
func (msg *Message) AddReplyKeyboard(replykb IReply) error {
	var err error
	if kb, ok := replykb.(*reply); ok {
		if kb.Keyboard != nil {
			for i := 0; (i < len(kb.Keyboard.Keyboard)) && (err == nil); i++ {
				for j := 0; (j < len(kb.Keyboard.Keyboard[i])) && (err == nil); j++ {
					if kb.Keyboard.Keyboard[i][j].Text == "" {
						err = code21()
					}
				}
			}
			if err == nil {
				msg.fm.kb = kb
				logs.InterfaceSaved(interfaceReplyKB)
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Add a inline-keyboard interface to the message you're building
func (msg *Message) AddInlineKeyboard(inlinekb IInline) error {
	var err error
	if kb, ok := inlinekb.(*inline); ok {
		if kb.Keyboard != nil {
			for i := 0; (i < len(kb.Keyboard.InlineKeyboard)) && (err == nil); i++ {
				for j := 0; (j < len(kb.Keyboard.InlineKeyboard[i])) && (err == nil); j++ {
					if kb.Keyboard.InlineKeyboard[i][j].Text == "" {
						err = code21()
					}
				}
			}
			if err == nil {
				msg.fm.kb = kb
				logs.InterfaceSaved(interfaceInKB)
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddMethod(method string) {
	msg.fm.method = method
	msg.fm.notchange = true
	logs.MethodSaved(method)
}
