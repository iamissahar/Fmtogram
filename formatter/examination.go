package formatter

import (
	"log"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/types"
)

type checking interface {
	defaultString(string) bool
	defautBool(bool) bool
	defaultInt(int) bool
}

func (ph *photo) defaultString(data string) bool {
	return data == ""
}

func (ph *photo) defautBool(data bool) bool {
	return !data
}

func (ph *photo) defaultInt(data int) bool {
	return data == 0
}

func (ph *video) defaultString(data string) bool {
	return data == ""
}

func (ph *video) defautBool(data bool) bool {
	return !data
}

func (ph *video) defaultInt(data int) bool {
	return data == 0
}

func (ph *audio) defaultString(data string) bool {
	return data == ""
}

func (ph *audio) defautBool(data bool) bool {
	return !data
}

func (ph *audio) defaultInt(data int) bool {
	return data == 0
}

func (ph *document) defaultString(data string) bool {
	return data == ""
}

func (ph *document) defautBool(data bool) bool {
	return !data
}

func (ph *document) defaultInt(data int) bool {
	return data == 0
}

func isItEmply(ch checking, work int, data interface{}) bool {
	var ok bool
	switch d := data.(type) {
	case string:
		if work == checkString {
			ok = ch.defaultString(d)
		}
	case []*types.MessageEntity:
		if work == checkArray {
			ok = len(d) == 0
		}
	case bool:
		if work == checkBool {
			ok = ch.defautBool(d)
		}
	case int:
		if work == checkInt {
			ok = ch.defaultInt(d)
		}
	default:
		log.Print("SOMETHING WENT WRONG!")
	}

	return ok
}

func requiredPhotoData(ph *photo, num int) error {
	var err error
	if ph.Photo == "" {
		err = errors.MissedRequiredField(interfacePhoto, "Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredVideoData(vd *video, num int) error {
	var err error
	if vd.Video == "" {
		err = errors.MissedRequiredField(interfaceVideo, "Write{Storage,Telgram,Internet}Video{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredAudioData(ad *audio, num int) error {
	var err error
	if ad.Audio == "" {
		err = errors.MissedRequiredField(interfaceAudio, "Write{Storage,Telgram,Internet}Audio{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredDocumentData(dc *document, num int) error {
	var err error
	if dc.Document == "" {
		err = errors.MissedRequiredField(interfaceDocument, "Write{Storage,Telgram,Internet}Document{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}
