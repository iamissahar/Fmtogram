package formatter

import (
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

// Added a photo interface to the message you're building
func (msg *Message) AddPhoto(ph IPhoto) error {
	var err error
	if p, ok := ph.(*photo); ok {
		if p.Photo != "" {
			if compatibilityCheck(msg, constPhoto) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if p.gottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if msg.fm.method == "" || msg.fm.method == methods.Photo {
						if msg.fm.mh.i == 0 {
							msg.fm.method = methods.Photo
							msg.fm.tgr = new(types.MessageResponse)
						} else {
							msg.fm.method = methods.MediaGroup
							msg.fm.tgr = new(types.MediaGroupResponse)
						}
					}
					msg.fm.mh.storage[msg.fm.mh.i] = p
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfacePhoto)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
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
			if compatibilityCheck(msg, constVideo) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if v.videoGottenFrom == Storage || v.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if msg.fm.method == "" || msg.fm.method == methods.Video {
						if msg.fm.mh.i == 0 {
							msg.fm.method = methods.Video
							msg.fm.tgr = new(types.MessageResponse)
							v.StartTimestamp = &msg.fm.inf.VideoStartTimestamp
						} else {
							msg.fm.method = methods.MediaGroup
							msg.fm.tgr = new(types.MediaGroupResponse)
						}
					}
					msg.fm.mh.storage[msg.fm.mh.i] = v
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfaceVideo)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
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
			if compatibilityCheck(msg, constDoc) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if d.documentGottenFrom == Storage || d.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if msg.fm.mh.i == 0 {
						msg.fm.method = methods.Document
						msg.fm.tgr = new(types.MessageResponse)
					} else {
						msg.fm.method = methods.MediaGroup
						msg.fm.tgr = new(types.MediaGroupResponse)
					}
					msg.fm.mh.storage[msg.fm.mh.i] = d
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfaceDocument)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
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
			if compatibilityCheck(msg, constAnim) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if a.animationGottenFrom == Storage || a.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					msg.fm.method = methods.Animation
					msg.fm.tgr = new(types.MessageResponse)
					msg.fm.mh.storage[msg.fm.mh.i] = a
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfaceAnimation)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
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
			if compatibilityCheck(msg, constAudio) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if a.audioGottenFrom == Storage || a.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if msg.fm.mh.i == 0 {
						msg.fm.method = methods.Audio
						msg.fm.tgr = new(types.MessageResponse)
					} else {
						msg.fm.method = methods.MediaGroup
						msg.fm.tgr = new(types.MediaGroupResponse)
					}
					msg.fm.mh.storage[msg.fm.mh.i] = a
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfaceAudio)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
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
			if compatibilityCheck(msg, constVoice) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if v.gottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					msg.fm.method = methods.Voice
					msg.fm.tgr = new(types.MessageResponse)
					msg.fm.mh.storage[msg.fm.mh.i] = v
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfaceVoice)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
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
			if compatibilityCheck(msg, constVideoNote) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if v.videoGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					msg.fm.method = methods.VideoNote
					msg.fm.tgr = new(types.MessageResponse)
					msg.fm.mh.storage[msg.fm.mh.i] = v
					msg.fm.mh.i++
					msg.fm.mh.amount++
					logs.InterfaceSaved(interfaceVideoNote)
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

// Added a location interface to the message you're building
func (msg *Message) AddLocation(loc ILocation) error {
	var err error
	if l, ok := loc.(*location); ok {
		if msg.fm.loc == nil {
			if (l.Latitude != 0) && (l.Longitude) != 0 {
				msg.fm.loc = l
				if msg.fm.loc.Title == "" {
					msg.fm.method = methods.Location
				} else {
					msg.fm.method = methods.Venue
				}
				msg.fm.tgr = new(types.MessageResponse)
				logs.InterfaceSaved(interfaceLocation)
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddContact(con IContact) error {
	var err error
	if c, ok := con.(*contact); ok {
		if msg.fm.con == nil {
			if (c.PhoneNumber != "") && (c.FirstName != "") {
				msg.fm.con = c
				msg.fm.method = methods.Contact
				msg.fm.tgr = new(types.MessageResponse)
				logs.InterfaceSaved(interfaceContact)
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddPoll(p IPoll) error {
	var err error
	if pp, ok := p.(*poll); ok {
		if msg.fm.poll == nil {
			if (pp.Question != "") && (pp.Options != nil) {
				msg.fm.poll = pp
				msg.fm.method = methods.Poll
				msg.fm.tgr = new(types.MessageResponse)
				logs.InterfaceSaved(interfacePoll)
			} else {
				err = code21()
			}
		} else {
			err = code10()
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
		if isDefaultParams(msg.fm.inf) {
			msg.fm.inf = p
			if p.StarCount != 0 {
				msg.fm.method = methods.PaidMedia
				msg.fm.tgr = new(types.MessageResponse)
			} else if p.Emoji != "" {
				msg.fm.method = methods.Dice
				msg.fm.tgr = new(types.MessageResponse)
			} else if p.Action != "" {
				msg.fm.method = methods.ChatAction
				msg.fm.tgr = new(types.SimpleResponse)
			} else if msg.fm.method == "" {
				msg.fm.method = methods.Message
				msg.fm.tgr = new(types.MessageResponse)
			}
			logs.InterfaceSaved(interfaceParam)
		} else {
			err = code10()
		}
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
			if isDefaultChat(msg.fm.ch) {
				msg.fm.ch = c
				logs.InterfaceSaved(interfaceChat)
			} else {
				err = code10()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddKeyboard(kb IKeyboard) error {
	var err error
	if k, ok := kb.(*keyboard); ok {
		if err = isNill(k.Keyboard); err != nil {
			if err = k.Keyboard.isOK(); err == nil {
				if msg.fm.kb == nil {
					msg.fm.kb = k
					logs.InterfaceSaved(interfaceKeyboard)
				} else {
					err = code10()
				}
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddMethod(method string) error {
	var err error
	var ok bool
	var data interface{}
	if data, ok = simpleResponse[method]; ok {
		msg.fm.tgr = data
	} else if data, ok = messageResponse[method]; ok {
		msg.fm.tgr = data
	} else if data, ok = messageIDsReponse[method]; ok {
		msg.fm.tgr = data
	} else if data, ok = inviteLinkResponse[method]; ok {
		msg.fm.tgr = data
	} else if method == methods.MediaGroup {
		msg.fm.tgr = new(types.MediaGroupResponse)
	} else if method == methods.UserProfilePhotos {
		msg.fm.tgr = new(types.UserProfilePhotosResponse)
	} else if (method == methods.File) || (method == methods.UploadSticker) {
		msg.fm.tgr = new(types.FileResponse)
	} else if method == methods.GetChat {
		msg.fm.tgr = new(types.ChatFullInfoResponse)
	} else if method == methods.GetAdmins {
		msg.fm.tgr = new(types.ChatMembersResponse)
	} else if method == methods.GetMember {
		msg.fm.tgr = new(types.ChatMemberResponse)
	} else if (method == methods.GetForumIconStickers) || (method == methods.GetEmojiStickers) {
		msg.fm.tgr = new(types.StickersResponse)
	} else if method == methods.CreateForumTopic {
		msg.fm.tgr = new(types.ForumResponse)
	} else if method == methods.GetUsersBoosts {
		msg.fm.tgr = new(types.UserBoostsResponse)
	} else if method == methods.GetBusinessConnection {
		msg.fm.tgr = new(types.BusinessConResponse)
	} else if method == methods.GetMyCommands {
		msg.fm.tgr = new(types.BotCommandResponse)
	} else if method == methods.GetMyName {
		msg.fm.tgr = new(types.BotNameResponse)
	} else if method == methods.GetMyDescription {
		msg.fm.tgr = new(types.BotDescriptionResponse)
	} else if method == methods.GetMyShortDescription {
		msg.fm.tgr = new(types.BotShorDescriptionResponse)
	} else if method == methods.GetChatMenuButton {
		msg.fm.tgr = new(types.MenuButtonResponse)
	} else if method == methods.GetMyDefaultAdministratorRights {
		msg.fm.tgr = new(types.AdminRightResponse)
	} else if method == methods.StopPoll {
		msg.fm.tgr = new(types.PollResponse)
	} else if method == methods.GetStickerSet {
		msg.fm.tgr = new(types.StickerSetResponse)
	} else if method == methods.GetAvailableGifts {
		msg.fm.tgr = new(types.GiftsResponse)
	} else if method == methods.AnswerWebAppQuery {
		msg.fm.tgr = new(types.WepAppMsgResponse)
	} else if method == methods.SavePreparedInlineMessage {
		msg.fm.tgr = new(types.PreparedInlineMessageResponse)
	} else if method == methods.ExportInviteLink || method == methods.CreateInvoiceLink {
		msg.fm.tgr = new(types.StringResponse)
	} else if method == methods.GetStarTransactions {
		msg.fm.tgr = new(types.StarTransactionResponse)
	} else if method == methods.GetGameHighScores {
		msg.fm.tgr = new(types.GameHighScoresResponse)
	} else if method == methods.GetMemberCount {
		msg.fm.tgr = new(types.IntResponse)
	} else {
		err = code07()
	}
	if err == nil {
		msg.fm.method = method
		msg.fm.notchange = true
		logs.MethodSaved(method)
	}
	return err
}

func (msg *Message) AddLink(alink ILink) error {
	var err error
	if l, ok := alink.(*link); ok {
		if msg.fm.link == nil {
			msg.fm.link = l
			logs.InterfaceSaved(interfaceLink)
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddSticker(st ISticker) error {
	var err error
	if s, ok := st.(*sticker); ok {
		if msg.fm.sticker.storage[msg.fm.sticker.i] == nil {
			if (s.Sticker != "") || (s.SetName != "") {
				if s.stickerGottenFrom == Storage {
					msg.fm.sticker.atLeastOnce = true
				}
				msg.fm.sticker.storage[msg.fm.sticker.i] = s
				msg.fm.sticker.i++
				if s.GiftID != "" {
					msg.fm.method = methods.Gift
				} else {
					msg.fm.method = methods.Sticker
				}
				logs.InterfaceSaved(interfaceSticker)
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddForum(frm IForum) error {
	var err error
	if f, ok := frm.(*forum); ok {
		if msg.fm.forum == nil {
			msg.fm.forum = f
			logs.InterfaceSaved(interfaceForum)
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddBot(b IBot) error {
	var err error
	if bb, ok := b.(*bot); ok {
		if msg.fm.bot == nil {
			if ((bb.Name != nil) && (*bb.Name == "" && bb.Language == nil)) ||
				((bb.Description != nil) && (*bb.Description == "" && bb.Language == nil)) ||
				((bb.ShortDescription != nil) && (*bb.ShortDescription == "" && bb.Language == nil)) {
				err = code25()
			} else {
				msg.fm.bot = bb
				logs.InterfaceSaved(interfaceBot)
			}
		} else {
			err = code10()
		}
	}
	return err
}

func (msg *Message) AddInlineMode(inmode IInlineMode) error {
	var err error
	if in, ok := inmode.(*inlinemode); ok {
		if msg.fm.inlinemode == nil {
			msg.fm.inlinemode = in
			logs.InterfaceSaved(interfaceInlineMode)
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddPayment(pay IPayment) error {
	var err error
	if p, ok := pay.(*payment); ok {
		if msg.fm.payment == nil {
			if (p.Title != "") && (p.Description != "") && (p.Payload != "") && (p.Currency != "") && (p.Prices != nil) {
				msg.fm.payment = p
				msg.fm.method = methods.Invoice
				logs.InterfaceSaved(interfacePayment)
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) AddGame(gg IGame) error {
	var err error
	if g, ok := gg.(*game); ok {
		if msg.fm.game == nil {
			if g.ShortName != "" {
				msg.fm.game = g
				msg.fm.method = methods.Game
				logs.InterfaceSaved(interfacePayment)
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}
