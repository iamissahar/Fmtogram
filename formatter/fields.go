package formatter

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/fmerrors"
)

func writeFileToMultipart(writer *multipart.Writer, fieldname, filename string) error {
	var part io.Writer
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()

		part, err = writer.CreateFormFile(fieldname, filename)

		if err == nil {
			_, err = io.Copy(part, file)
			if err != nil {
				fmerrors.CantCopyFile(err)
			}
		} else {
			fmerrors.CantCreateFormFile(err)
		}
	} else {
		fmerrors.CantOpenFile(err)
	}
	return err
}

func addThumbnail(writer *multipart.Writer, thumbnail string, gottenfrom int) error {
	var err error
	if gottenfrom == Storage {
		err = writeFileToMultipart(writer, "thumbnail", thumbnail)
	} else {
		err = writer.WriteField("thumbnail", thumbnail)

		if err != nil {
			fmerrors.CantWriteField(err)
		}
	}
	return err
}

func (ph *photo) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var body []byte
	var err error
	if input {
		if ph.gottenFrom != Storage {
			ph.Media = ph.Photo
		} else {
			err = writeFileToMultipart(writer, ph.Photo, ph.Photo)
			ph.Media = fmt.Sprintf("attach://%s", ph.Photo)
		}
	} else {
		if ph.gottenFrom != Storage {
			err = writer.WriteField("photo", ph.Photo)
		} else {
			err = writeFileToMultipart(writer, "photo", ph.Photo)
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "photo")
	}
	if err == nil && ph.Caption != "" {
		err = field(writer, "caption", ph.Caption)
	}
	if err == nil && ph.ParseMode != "" {
		err = field(writer, "parse_mode", ph.ParseMode)
	}
	if err == nil && ph.CaptionEntities != nil {
		if body, err = json.Marshal(ph.CaptionEntities); err == nil {
			err = field(writer, "caption_entities", string(body))
		}
	}
	if err == nil && ph.ShowCaptionAboveMedia {
		err = field(writer, "show_caption_above_media", "True")
	}
	if err == nil && ph.HasSpoiler {
		err = field(writer, "has_spoiler", "True")
	}
	if group != nil {
		(*group)[i] = ph
	}
	return err
}

func (ph *photo) jsonFileds() ([]byte, error) {
	return json.Marshal(ph)
}

func (ph *photo) uniqueConst() int {
	return constPhoto
}

func (vd *video) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var body []byte
	var err error
	if input {
		if vd.videoGottenFrom != Storage {
			vd.Media = vd.Video
		} else {
			err = writeFileToMultipart(writer, vd.Video, vd.Video)
			vd.Media = fmt.Sprintf("attach://%s", vd.Video)
		}
	} else {
		if vd.videoGottenFrom != Storage {
			err = writer.WriteField("video", vd.Video)
		} else {
			err = writeFileToMultipart(writer, "video", vd.Video)
		}
	}
	if err == nil && input {
		err = field(writer, "type", "video")
	}
	if err == nil && vd.Thumbnail != "" {
		err = addThumbnail(writer, vd.Thumbnail, vd.thumbnailGottenFrom)
	}
	if err == nil && vd.Caption != "" {
		err = field(writer, "caption", vd.Caption)
	}
	if err == nil && vd.ParseMode != "" {
		err = field(writer, "parse_mode", vd.ParseMode)
	}
	if err == nil && vd.CaptionEntities != nil {
		if body, err = json.Marshal(vd.CaptionEntities); err == nil {
			err = field(writer, "caption_entities", string(body))
		}
	}
	if err == nil && vd.ShowCaptionAboveMedia {
		err = field(writer, "show_caption_above_media", "True")
	}
	if err == nil && vd.Width != 0 {
		err = field(writer, "width", fmt.Sprintf("%d", vd.Width))
	}
	if err == nil && vd.Height != 0 {
		err = field(writer, "height", fmt.Sprintf("%d", vd.Height))
	}
	if err == nil && vd.Duration != 0 {
		err = field(writer, "duration", fmt.Sprintf("%d", vd.Duration))
	}
	if err == nil && vd.SupportsStreaming {
		err = field(writer, "supports_streaming", "True")
	}
	if err == nil && vd.HasSpoiler {
		err = field(writer, "has_spoiler", "True")
	}
	if group != nil {
		(*group)[i] = vd
	}
	return err
}

func (vd *video) jsonFileds() ([]byte, error) {
	return json.Marshal(vd)
}

func (vd *video) uniqueConst() int {
	return constVideo
}

func (ad *audio) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var body []byte
	var err error
	if input {
		if ad.audioGottenFrom != Storage {
			ad.Media = ad.Audio
		} else {
			err = writeFileToMultipart(writer, ad.Audio, ad.Audio)
			ad.Media = fmt.Sprintf("attach://%s", ad.Audio)
		}
	} else {
		if ad.audioGottenFrom != Storage {
			err = field(writer, "audio", ad.Audio)
		} else {
			err = writeFileToMultipart(writer, "audio", ad.Audio)
		}
	}
	if err == nil && input {
		err = field(writer, "type", "audio")
	}
	if err == nil && ad.Thumbnail != "" {
		err = addThumbnail(writer, ad.Thumbnail, ad.thumbnailGottenFrom)
	}
	if err == nil && ad.ParseMode != "" {
		err = field(writer, "parse_mode", ad.ParseMode)
	}
	if err == nil && ad.CaptionEntities != nil {
		if body, err = json.Marshal(ad.CaptionEntities); err == nil {
			err = field(writer, "caption_entities", string(body))
		}
	}
	if err == nil && ad.Duration != 0 {
		err = field(writer, "duration", fmt.Sprintf("%d", ad.Duration))
	}
	if err == nil && ad.Performer != "" {
		err = field(writer, "performer", ad.Performer)
	}
	if err == nil && ad.Title != "" {
		err = field(writer, "title", ad.Title)
	}
	if group != nil {
		(*group)[i] = ad
	}
	return err
}

func (ad *audio) jsonFileds() ([]byte, error) {
	return json.Marshal(ad)
}

func (ad *audio) uniqueConst() int {
	return constAudio
}

func (dc *document) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var body []byte
	var err error
	if input {
		if dc.documentGottenFrom != Storage {
			dc.Media = dc.Document
		} else {
			err = writeFileToMultipart(writer, dc.Document, dc.Document)
			dc.Media = fmt.Sprintf("attach://%s", dc.Document)
		}
	} else {
		if dc.documentGottenFrom != Storage {
			err = field(writer, "document", dc.Document)
		} else {
			err = writeFileToMultipart(writer, "document", dc.Document)
		}
	}
	if err == nil && input {
		err = field(writer, "type", "document")
	}
	if err == nil && dc.Thumbnail != "" {
		err = addThumbnail(writer, dc.Thumbnail, dc.thumbnailGottenFrom)
	}
	if err == nil && dc.Caption != "" {
		err = field(writer, "caption", dc.Caption)
	}
	if err == nil && dc.ParseMode != "" {
		err = field(writer, "parse_mode", dc.ParseMode)
	}
	if err == nil && dc.CaptionEntities != nil {
		if body, err = json.Marshal(dc.CaptionEntities); err == nil {
			err = field(writer, "caption_entities", string(body))
		}
	}
	if err == nil && dc.DisableContentTypeDetection {
		err = field(writer, "disable_content_type_detection", "True")
	}
	if group != nil {
		(*group)[i] = dc
	}
	return err
}

func (dc *document) jsonFileds() ([]byte, error) {
	return json.Marshal(dc)
}

func (dc *document) uniqueConst() int {
	return constDoc
}

func (an *animation) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	if input {
		if an.animationGottenFrom != Storage {
			an.Media = an.Animation
		} else {
			err = writeFileToMultipart(writer, an.Animation, an.Animation)
			an.Media = fmt.Sprintf("attach://%s", an.Animation)
		}
	} else {
		if an.animationGottenFrom != Storage {
			err = field(writer, "animation", an.Animation)
		} else {
			err = writeFileToMultipart(writer, "animation", an.Animation)
		}
	}
	if err == nil && input {
		err = field(writer, "type", "animation")
	}
	if err == nil && an.Thumbnail != "" {
		err = addThumbnail(writer, an.Thumbnail, an.thumbnailGottenFrom)
	}
	if err == nil && an.Width != 0 {
		err = field(writer, "width", fmt.Sprintf("%d", an.Width))
	}
	if err == nil && an.Height != 0 {
		err = field(writer, "height", fmt.Sprintf("%d", an.Height))
	}
	if err == nil && an.Duration != 0 {
		err = field(writer, "duration", fmt.Sprintf("%d", an.Duration))
	}
	if err == nil && an.HasSpoiler {
		err = field(writer, "has_spoiler", "True")
	}
	if group != nil {
		(*group)[i] = an
	}
	return err
}

func (an *animation) jsonFileds() ([]byte, error) {
	return json.Marshal(an)
}

func (an *animation) uniqueConst() int {
	return constAnim
}

func (vc *voice) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	if input {
		if vc.gottenFrom != Storage {
			vc.Media = vc.Voice
		} else {
			err = writeFileToMultipart(writer, vc.Voice, vc.Voice)
			vc.Media = fmt.Sprintf("attach://%s", vc.Voice)
		}
	} else {
		if vc.gottenFrom != Storage {
			err = field(writer, "voice", vc.Voice)
		} else {
			err = writeFileToMultipart(writer, "voice", vc.Voice)
		}
	}
	if err == nil && input {
		err = field(writer, "type", "voice")
	}
	if err == nil && vc.Duration != 0 {
		err = field(writer, "duration", fmt.Sprintf("%d", vc.Duration))
	}
	if group != nil {
		(*group)[i] = vc
	}
	return err
}

func (vc *voice) jsonFileds() ([]byte, error) {
	return json.Marshal(vc)
}

func (*voice) uniqueConst() int {
	return constVoice
}

func (vdn *videonote) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	if input {
		if vdn.videoGottenFrom != Storage {
			vdn.Media = vdn.VideoNote
		} else {
			err = writeFileToMultipart(writer, vdn.VideoNote, vdn.VideoNote)
			vdn.Media = fmt.Sprintf("attach://%s", vdn.VideoNote)
		}
	} else {
		if vdn.videoGottenFrom != Storage {
			err = field(writer, "video_note", vdn.VideoNote)
		} else {
			err = writeFileToMultipart(writer, "video_note", vdn.VideoNote)
		}
	}
	if err == nil && input {
		err = field(writer, "type", "video_note")
	}
	if err == nil && vdn.Duration != 0 {
		err = field(writer, "duration", fmt.Sprintf("%d", vdn.Duration))

	}
	if err == nil && vdn.Length != 0 {
		err = field(writer, "length", fmt.Sprintf("%d", vdn.Length))
	}
	if err == nil && vdn.Thumbnail != "" {
		err = addThumbnail(writer, vdn.Thumbnail, vdn.thumbnailGottenFrom)
	}
	if group != nil {
		(*group)[i] = vdn
	}
	return err
}

func (vc *videonote) jsonFileds() ([]byte, error) {
	return json.Marshal(vc)
}

func (*videonote) uniqueConst() int {
	return constVideoNote
}

func (st *sticker) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	var body []byte
	if input {
		if st.stickerGottenFrom == Storage {
			err = writeFileToMultipart(writer, st.Sticker, st.Sticker)
			st.Sticker = fmt.Sprintf("attach://%s", st.Sticker)
		}
	} else {
		if st.stickerGottenFrom != Storage {
			err = field(writer, "sticker", st.Sticker)
		} else {
			err = writeFileToMultipart(writer, "sticker", st.Sticker)
		}
	}
	if err == nil && st.Emoji != "" {
		err = field(writer, "emoji", st.Emoji)
	}
	if err == nil && st.Format != "" {
		err = field(writer, "sticker_format", st.Format)
		err = field(writer, "format", st.Format)
	}
	if err == nil && st.StickerType != "" {
		err = field(writer, "sticker_type", st.Format)
	}
	if err == nil && st.NeedsRepainting {
		err = field(writer, "needs_repainting", "True")
	}
	if err == nil && st.Name != "" {
		err = field(writer, "name", st.Name)
	}
	if err == nil && st.OldSticker != "" {
		err = field(writer, "old_sticker", st.OldSticker)
	}
	if err == nil && st.Emojies != nil {
		if body, err = json.Marshal(st.Emojies); err == nil {
			err = field(writer, "emoji_list", string(body))
		}
	}
	if err == nil && st.MaskPosition != nil {
		if body, err = json.Marshal(st.MaskPosition); err == nil {
			err = field(writer, "mask_position", string(body))
		}
	}
	if err == nil && st.Keywords != nil {
		if body, err = json.Marshal(st.Keywords); err == nil {
			err = field(writer, "keywords", string(body))
		}
	}
	if err == nil && st.Thumbnail != "" {
		err = addThumbnail(writer, st.Thumbnail, st.thumbnailGottenFrom)
	}
	if err == nil && st.ThumbnailFormat != "" {
		err = field(writer, "format", st.ThumbnailFormat)
	}
	if group != nil {
		(*group)[i] = st
	}
	return err
}

func putGroup(writer *multipart.Writer, group []interface{}) error {
	mediaJSON, err := json.Marshal(group)
	if err == nil && len(mediaJSON) != 0 {
		err = writer.WriteField("media", string(mediaJSON))
	}
	return err
}

func field(w *multipart.Writer, fieldname, value string) error {
	err := w.WriteField(fieldname, value)
	if err != nil {
		fmerrors.CantWriteField(err)
	}
	return err
}

func (inf *information) multipartFields(writer *multipart.Writer) error {
	var err error
	var body []byte
	if inf.MessageThreadID != 0 {
		err = field(writer, "message_thread_id", fmt.Sprint(inf.MessageThreadID))
	}
	if err == nil && inf.StarCount != 0 {
		err = field(writer, "star_count", fmt.Sprintf("%d", inf.StarCount))
	}
	if err == nil && inf.Payload != "" {
		err = field(writer, "payload", inf.Payload)
	}
	if err == nil && inf.UserID != 0 {
		err = field(writer, "user_id", fmt.Sprintf("%d", inf.UserID))
	}
	if err == nil && inf.Text != "" {
		err = field(writer, "caption", inf.Text)
	}
	if err == nil && inf.ParseMode != "" {
		err = field(writer, "parse_mode", inf.ParseMode)
	}
	if err == nil && inf.Entities != nil {
		if body, err = json.Marshal(inf.Entities); err == nil {
			err = field(writer, "caption_entities", string(body))
		}
	}
	if err == nil && inf.DisableNotification {
		err = field(writer, "disable_notification", "True")
	}
	if err == nil && inf.ProtectContent {
		err = field(writer, "protect_content", "True")
	}
	if err == nil && inf.AllowPaidBroadcast {
		err = field(writer, "allow_paid_broadcast", "True")
	}
	if err == nil && inf.MessageEffectID != "" {
		err = field(writer, "message_effect_id", inf.MessageEffectID)
	}
	if err == nil && inf.ReplyParameters != nil {
		if body, err = json.Marshal(inf.ReplyParameters); err == nil {
			err = field(writer, "message_effect_id", string(body))
		}
	}
	return err
}

func (inf *information) createJSON() ([]byte, error) {
	body, err := json.Marshal(inf)
	if err != nil {
		fmerrors.CantMarshalJSON(err)
	}
	return body, err
}

func (ch *chat) multipartFields(writer *multipart.Writer) error {
	err := field(writer, "chat_id", fmt.Sprint(ch.ID))
	if ch.BusinessConnectionID != "" {
		err = field(writer, "business_connection_id", ch.BusinessConnectionID)
	}
	return err
}

func (ch *chat) createJson() ([]byte, error) {
	return json.Marshal(ch)
}

func (in *inline) get() ([]byte, error) {
	return json.Marshal(in)
}

func (rp *reply) get() ([]byte, error) {
	return json.Marshal(rp)
}

func (frp *forcereply) get() ([]byte, error) {
	return json.Marshal(frp)
}

func (kb *keyboard) multipartFields(writer *multipart.Writer) error {
	body, err := kb.Keyboard.get()
	if err == nil {
		err = writer.WriteField("reply_markup", string(body))
	}
	return err
}

func (kb *keyboard) jsonFields() ([]byte, error) {
	return json.Marshal(kb)
}
