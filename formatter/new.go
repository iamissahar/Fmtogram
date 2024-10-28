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
	return &animation{Type: "video"}
}

// Creates a new message(inf) stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewMessage() IMSGInformation {
	logs.NewInterfaceCreated("Message")
	return &information{}
}

// Creates a new chat stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewChat() IChat {
	logs.NewInterfaceCreated("Chat")
	return &chat{}
}

// Creates a new inline-keyboard stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewInlineKeyboard() IInline {
	logs.NewInterfaceCreated("Inline Keyboard")
	return &inline{}
}

// Creates a new reply-keyboard stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (msg *Message) NewReplyKeyboard() IReply {
	logs.NewInterfaceCreated("Reply Keyboard")
	return &reply{}
}
