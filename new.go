package fmtogram

func NewMessage() *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{}
	m.fm.prm = &information{}
	m.fm.mh = &mediaHolder{}
	m.fm.sticker = &stickerHolder{}
	m.fm.g = &get{}
	return m
}

func NewPhoto() IPhoto {
	return &photo{Type: "photo"}
}

func NewDocument() IDocument {
	return &document{Type: "document"}
}

func NewAudio() IAudio {
	return &audio{Type: "audio"}
}

func NewVideo() IVideo {
	return &video{Type: "video"}
}

func NewAnimation() IAnimation {
	return &animation{Type: "animation"}
}

func NewVoice() IVoice {
	return &voice{Type: "voice"}
}

func NewVideoNote() IVideoNote {
	return &videonote{}
}

func NewLocation() ILocation {
	return &location{}
}

func NewContact() IContact {
	return &contact{}
}

func NewPoll() IPoll {
	return &poll{}
}

func NewSticker() ISticker {
	return &sticker{}
}

func NewKeyboard() IKeyboard {
	return &keyboard{}
}

func NewPayment() IPayment {
	return &payment{}
}

func NewGift() IGift {
	return &gift{}
}

func (msg *Message) GetResults() IGet {
	msg.fm.g = &get{}
	return msg.fm.g
}
