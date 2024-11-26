package formatter

import (
	"bytes"
	"os"

	"github.com/l1qwie/Fmtogram/fmerrors"
)

func mpeg4(buf []byte) bool {
	return len(buf) >= 8 && bytes.HasPrefix(buf[4:], []byte("ftyp"))
}

func jpeg(buf []byte) bool {
	return bytes.HasPrefix(buf, []byte{0xFF, 0xD8, 0xFF})
}

func png(buf []byte) bool {
	return bytes.HasPrefix(buf, []byte{0x89, 0x50, 0x4E, 0x47})
}

func mp3(buf []byte) bool {
	return len(buf) >= 3 && (bytes.HasPrefix(buf, []byte{0xFF, 0xFB}) || bytes.HasPrefix(buf, []byte{0xFF, 0xF3}))
}

func m4a(buf []byte) bool {
	return len(buf) >= 8 && bytes.HasPrefix(buf[4:], []byte("ftyp")) &&
		(bytes.Contains(buf, []byte("isom")) || bytes.Contains(buf, []byte("mp42")) || bytes.Contains(buf, []byte("avc1")))
}

func gif(buf []byte) bool {
	return len(buf) >= 6 && bytes.HasPrefix(buf, []byte("GIF8"))
}

func h264mpeg4AVC(buf []byte) bool {
	return len(buf) >= 8 && bytes.HasPrefix(buf[4:], []byte("ftyp")) &&
		(bytes.Contains(buf, []byte("isom")) || bytes.Contains(buf, []byte("mp42")) || bytes.Contains(buf, []byte("avc1")))
}

func h264(buf []byte) bool {
	return bytes.HasPrefix(buf, []byte{0x00, 0x00, 0x00, 0x01}) || bytes.HasPrefix(buf, []byte{0x00, 0x00, 0x01})
}

func oggWithOpus(buf []byte) bool {
	return (len(buf) >= 8 && bytes.HasPrefix(buf, []byte("OggS"))) && (bytes.HasPrefix(buf[4:], []byte("OpusHead")))
}

func headerGiver(header *[]byte, path string) error {
	var n int
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		n, err = file.Read(*header)
	}
	*header = (*header)[:n]
	return err
}

func (ph *photo) isCorrectType(path string) error {
	var err error
	buf := make([]byte, 8)
	if err = headerGiver(&buf, path); err == nil {
		if !jpeg(buf) && !png(buf) {
			err = code12()
		}
	}
	return err
}

func (vd *video) isCorrectType(path string) error {
	var err error
	buf := make([]byte, 12)
	if err = headerGiver(&buf, path); err == nil {
		if !mpeg4(buf) {
			err = code12()
		}
	}
	return err
}

func isThumbnailCorrectType(path string) error {
	var err error
	buf := make([]byte, 8)
	if err = headerGiver(&buf, path); err == nil {
		if !jpeg(buf) {
			err = code12()
		}
	}
	return err
}

func (ad *audio) isCorrectType(path string) error {
	var err error
	buf := make([]byte, 12)
	if err = headerGiver(&buf, path); err == nil {
		if !mp3(buf) && !m4a(buf) {
			err = code12()
		}
	}
	return err
}

func (an *animation) isCorrectType(path string) error {
	var err error
	buf := make([]byte, 12)
	if err = headerGiver(&buf, path); err == nil {
		if !gif(buf) && !h264mpeg4AVC(buf) && !h264(buf) {
			err = code12()
		}
	}
	return err
}

func (vc *voice) isCorrectType(path string) error {
	var err error
	buf := make([]byte, 12)
	if err = headerGiver(&buf, path); err == nil {
		if !oggWithOpus(buf) && !mp3(buf) && !m4a(buf) {
			err = code12()
		}
	}
	return err
}

func (vdn *videonote) isCorrectType(path string) error {
	var err error
	buf := make([]byte, 12)
	if err = headerGiver(&buf, path); err == nil {
		if !mpeg4(buf) {
			err = code12()
		}
	}
	return err
}

func requiredPhotoData(ph *photo, num int) error {
	var err error
	if ph.Photo == "" {
		err = fmerrors.MissedRequiredField(interfacePhoto, "Write{Storage,Telgram,Internet}Photo{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredVideoData(vd *video, num int) error {
	var err error
	if vd.Video == "" {
		err = fmerrors.MissedRequiredField(interfaceVideo, "Write{Storage,Telgram,Internet}Video{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredAudioData(ad *audio, num int) error {
	var err error
	if ad.Audio == "" {
		err = fmerrors.MissedRequiredField(interfaceAudio, "Write{Storage,Telgram,Internet}Audio{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredDocumentData(dc *document, num int) error {
	var err error
	if dc.Document == "" {
		err = fmerrors.MissedRequiredField(interfaceDocument, "Write{Storage,Telgram,Internet}Document{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredAnimationData(an *animation, num int) error {
	var err error
	if an.Animation == "" {
		err = fmerrors.MissedRequiredField(interfaceAnimation, "Write{Storage,Telgram,Internet}Animation{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredVoiceData(vc *voice, num int) error {
	var err error
	if vc.Voice == "" {
		err = fmerrors.MissedRequiredField(interfaceVoice, "Write{Storage,Telgram,Internet}Voice{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}

func requiredVideoNoteData(vdn *videonote, num int) error {
	var err error
	if vdn.VideoNote == "" {
		err = fmerrors.MissedRequiredField(interfaceVideoNote, "Write{Storage,Telgram,Internet}VideoNote{FilePath/ID/URL}", num, 0, true, false)
	}
	return err
}
