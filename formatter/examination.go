package formatter

import (
	"bytes"
	"log"
	"os"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/types"
)

type checking interface {
	defaultString(string) bool
	defautBool(bool) bool
	defaultInt(int) bool
}

func (*photo) defaultString(data string) bool {
	return data == ""
}

func (*photo) defautBool(data bool) bool {
	return !data
}

func (*photo) defaultInt(data int) bool {
	return data == 0
}

func (*video) defaultString(data string) bool {
	return data == ""
}

func (*video) defautBool(data bool) bool {
	return !data
}

func (*video) defaultInt(data int) bool {
	return data == 0
}

func (*audio) defaultString(data string) bool {
	return data == ""
}

func (*audio) defautBool(data bool) bool {
	return !data
}

func (*audio) defaultInt(data int) bool {
	return data == 0
}

func (*document) defaultString(data string) bool {
	return data == ""
}

func (*document) defautBool(data bool) bool {
	return !data
}

func (*document) defaultInt(data int) bool {
	return data == 0
}

func (*animation) defaultString(data string) bool {
	return data == ""
}

func (*animation) defautBool(data bool) bool {
	return !data
}

func (*animation) defaultInt(data int) bool {
	return data == 0
}

func (*voice) defaultString(data string) bool {
	return data == ""
}

func (*voice) defautBool(data bool) bool {
	return !data
}

func (*voice) defaultInt(data int) bool {
	return data == 0
}

func (*videonote) defaultString(data string) bool {
	return data == ""
}

func (*videonote) defautBool(data bool) bool {
	return !data
}

func (*videonote) defaultInt(data int) bool {
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

func requiredAnimationData(an *animation, num int) error {
	var err error
	if an.Animation == "" {
		err = errors.MissedRequiredField(interfaceAnimation, "Write{Storage,Telgram,Internet}Animation{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredVoiceData(vc *voice, num int) error {
	var err error
	if vc.Voice == "" {
		err = errors.MissedRequiredField(interfaceVoice, "Write{Storage,Telgram,Internet}Voice{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredVideoNoteData(vdn *videonote, num int) error {
	var err error
	if vdn.VideoNote == "" {
		err = errors.MissedRequiredField(interfaceVideoNote, "Write{Storage,Telgram,Internet}VideoNote{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func headerGiver(path string) ([]byte, error) {
	header := make([]byte, 12)
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		_, err = file.Read(header)
	}
	return header, err
}

// func (an *sticker) detectStickerType() error {
// 	file, err := os.Open(an.Animation)
// 	if err == nil {
// 		defer file.Close()

// 		header := make([]byte, 12)
// 		_, err = file.Read(header)

// 		if err == nil {
// 			if ((bytes.HasPrefix(header, []byte{0x52, 0x49, 0x46, 0x46})) && (bytes.Contains(header[8:], []byte("WEBP")))) ||
// 				(bytes.HasPrefix(header, []byte{0x54, 0x47, 0x53, 0x33})) || (bytes.HasPrefix(header, []byte{0x1A, 0x45, 0xDF, 0xA3})) {
// 				err = nil
// 			} else {
// 				err = errors.UnsupportedTypeOfFile("GIF or H.264/MPEG-4 AVC video without sound")
// 			}
// 		}
// 	}

// 	return err
// }

func (vd *video) checkTypeOfFile() error {
	header, err := headerGiver(vd.Video)
	if err == nil {
		if (bytes.HasPrefix(header, []byte{0x00, 0x00, 0x00, 0x20})) && (bytes.Contains(header[4:], []byte("ftyp"))) {
			err = nil
		} else {
			err = errors.UnsupportedTypeOfFile("MPEG4 videos")
		}
	}

	return err
}

func (ad *audio) checkTypeOfFile() error {
	header, err := headerGiver(ad.Audio)
	if err == nil {
		if (bytes.HasPrefix(header, []byte{0x66, 0x74, 0x79, 0x70})) && (bytes.Contains(header[8:], []byte("mp4"))) {
			err = nil
		} else {
			err = errors.UnsupportedTypeOfFile(".MP3 or .M4A format")
		}
	}

	return err
}

func (an *animation) checkTypeOfFile() error {
	header, err := headerGiver(an.Animation)
	if err == nil {
		if (bytes.HasPrefix(header, []byte{0x00, 0x00, 0x00, 0x18})) && (bytes.Contains(header[4:], []byte("ftyp"))) {
			err = nil
		} else {
			err = errors.UnsupportedTypeOfFile("GIF or H.264/MPEG-4 AVC video without sound")
		}
	}

	return err
}

func (vc *voice) checkTypeOfFile() error {
	header, err := headerGiver(vc.Voice)
	if err == nil {
		if bytes.HasPrefix(header, []byte{0x4F, 0x67, 0x67, 0x53}) || // OGG
			bytes.HasPrefix(header, []byte{0xFF, 0xFB}) || // MP3
			(bytes.HasPrefix(header, []byte{0x66, 0x74, 0x79, 0x70}) && // M4A
				bytes.Contains(header[8:], []byte("mp4"))) {
			err = nil
		} else {
			err = errors.UnsupportedTypeOfFile(".OGG file encoded with OPUS, or in .MP3 format, or in .M4A format")
		}
	}
	return err
}

func (vdn *videonote) checkTypeOfFile() error {
	header, err := headerGiver(vdn.VideoNote)
	if err == nil {
		if (bytes.HasPrefix(header, []byte{0x00, 0x00, 0x00, 0x20})) && (bytes.Contains(header[4:], []byte("ftyp"))) {
			err = nil
		} else {
			err = errors.UnsupportedTypeOfFile("MPEG4 videos")
		}
	}

	return err
}
