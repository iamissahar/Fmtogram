package formatter

import (
	"bytes"
	"os"

	"github.com/l1qwie/Fmtogram/formatter/methods"
)

var matrix [7][7]int

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
	return len(buf) >= 36 && bytes.HasPrefix(buf, []byte("OggS")) && bytes.HasPrefix(buf[28:], []byte("OpusHead"))
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
	buf := make([]byte, 64)
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

func init() {
	matrix[constPhoto][constPhoto] = allowed
	matrix[constPhoto][constVideo] = allowed
	matrix[constAudio][constAudio] = allowed
	matrix[constVideo][constPhoto] = allowed
	matrix[constVideo][constVideo] = allowed
	matrix[constDoc][constDoc] = allowed
}

func compatibilityCheck(msg *Message, mediaID int) bool {
	res := true
	for j := 0; j < msg.fm.mh.i; j++ {
		othersID := msg.fm.mh.storage[j].uniqueConst()
		if (matrix[mediaID][othersID] != allowed) || (msg.fm.method == methods.PaidMedia) {
			res = false
		}
	}
	return res
}

func isDefaultChat(ch *chat) bool {
	res := true
	if ch.FromChatID != nil {
		res = false
	}
	if ch.ID != nil {
		res = false
	}
	if ch.BusinessConnectionID != "" {
		res = false
	}
	return res
}

func isDefaultParams(pr *information) bool {
	res := true
	if pr.MessageID != 0 {
		res = false
	}
	if pr.MessageIDs != nil {
		res = false
	}
	if pr.Text != "" {
		res = false
	}
	if pr.Caption != "" {
		res = false
	}
	if pr.ParseMode != "" {
		res = false
	}
	if pr.MessageThreadID != 0 {
		res = false
	}
	if pr.Entities != nil {
		res = false
	}
	if pr.CaptionEntities != nil {
		res = false
	}
	if pr.LinkPreviewOptions != nil {
		res = false
	}
	if pr.DisableNotification {
		res = false
	}
	if pr.ProtectContent {
		res = false
	}
	if pr.MessageEffectID != "" {
		res = false
	}
	if pr.ShowCaptionAboveMedia {
		res = false
	}
	if pr.ReplyParameters != nil {
		res = false
	}
	if pr.AllowPaidBroadcast {
		res = false
	}
	if pr.StarCount != 0 {
		res = false
	}
	if pr.Payload != "" {
		res = false
	}
	return res
}
