package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Creates a new photo stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewPhoto() IPhoto {
	logs.NewInterface("IPhoto")
	return &photo{Type: "photo"}
}

// Creates a new document stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewDocument() IDocument {
	logs.NewInterfaceCreated("Document")
	return &document{Type: "document"}
}

// Creates a new audio stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewAudio() IAudio {
	logs.NewInterfaceCreated("Audio")
	return &audio{Type: "audio"}
}

// Creates a new video stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewVideo() IVideo {
	logs.NewInterfaceCreated("Video")
	return &video{Type: "video"}
}

func (msg *Message) NewAnimation() IAnimation {
	logs.NewInterfaceCreated("Video")
	return &animation{Type: "animation"}
}

func (msg *Message) NewVoice() IVoice {
	logs.NewInterfaceCreated(interfaceVoice)
	return &voice{Type: "voice"}
}

func (msg *Message) NewVideoNote() IVideoNote {
	logs.NewInterfaceCreated(interfaceVideoNote)
	return &videonote{Type: "video_note"}
}

func (msg *Message) NewLocation() ILocation {
	logs.NewInterfaceCreated(interfaceLocation)
	return &location{}
}

func (msg *Message) NewContact() IContact {
	logs.NewInterfaceCreated(interfaceContact)
	return &contact{}
}

func (msg *Message) NewPoll() IPoll {
	logs.NewInterfaceCreated(interfacePoll)
	return &poll{}
}

func (msg *Message) NewLink() ILink {
	logs.NewInterfaceCreated(interfaceLink)
	return &link{}
}

func (msg *Message) NewSticker() ISticker {
	logs.NewInterfaceCreated(interfaceSticker)
	return &sticker{}
}

func (msg *Message) NewForum() IForum {
	logs.NewInterfaceCreated(interfaceForum)
	return &forum{}
}

func (msg *Message) NewBot() IBot {
	logs.NewInterfaceCreated(interfaceBot)
	return &bot{}
}

// Creates a new message(inf) stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewParameters() IParameters {
	logs.NewInterfaceCreated(interfaceParam)
	return &information{}
}

// Creates a new chat stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewChat() IChat {
	logs.NewInterfaceCreated(interfaceChat)
	return &chat{}
}

func (msg *Message) NewKeyboard() IKeyboard {
	logs.NewInterfaceCreated(interfaceKeyboard)
	return &keyboard{}
}

func (msg *Message) NewInlineMode() IInlineMode {
	logs.NewInterfaceCreated(interfaceInlineMode)
	return &inlinemode{}
}

func (msg *Message) NewPayment() IPayment {
	logs.NewInterfaceCreated(interfacePayment)
	return &payment{}
}

func (msg *Message) NewGame() IGame {
	logs.NewInterfaceCreated(interfaceGame)
	return &game{}
}
