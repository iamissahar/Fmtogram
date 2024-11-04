package formatter

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/types"
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
				errors.CantCopyFile(err)
			}
		} else {
			errors.CantCreateFormFile(err)
		}
	} else {
		errors.CantOpenFile(err)
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
			errors.CantWriteField(err)
		}
	}
	return err
}

func commonFields(writer *multipart.Writer, captionEntities []*types.MessageEntity) error {
	var (
		err  error
		body []byte
	)
	if len(captionEntities) != 0 {
		body, err = json.Marshal(captionEntities)
		if err == nil {
			err = writer.WriteField("caption_entities", string(body))

			if err != nil {
				errors.CantWriteField(err)
			}

		} else {
			errors.CantMarshalJSON(err)
		}
	}
	return err
}

func (ph *photo) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	err := commonFields(writer, ph.CaptionEntities)
	if err == nil {
		if input {
			if ph.GottenFrom != Storage {
				ph.Media = ph.Photo
			} else {
				err = writeFileToMultipart(writer, ph.Photo, ph.Photo)
				ph.Media = fmt.Sprintf("attach://%s", ph.Photo)
			}
		} else {
			if ph.GottenFrom != Storage {
				err = writer.WriteField("photo", ph.Photo)
			} else {
				err = writeFileToMultipart(writer, "photo", ph.Photo)
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "photo")
	}
	if err == nil && ph.ShowCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && ph.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if group != nil {
		(*group)[i] = ph
	}
	return err
}

func (ph *photo) jsonFileds() ([]byte, error) {
	return json.Marshal(ph)
}

func (ph *photo) nameAndConst() (string, int) {
	return interfacePhoto, constPhoto
}

func (vd *video) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	err := commonFields(writer, vd.CaptionEntities)
	if err == nil {
		if input {
			if vd.VideoGottenFrom != Storage {
				vd.Media = vd.Video
			} else {
				if err = vd.checkTypeOfFile(); err == nil {
					err = writeFileToMultipart(writer, vd.Video, vd.Video)
					vd.Media = fmt.Sprintf("attach://%s", vd.Video)
				}
			}
		} else {
			if vd.VideoGottenFrom != Storage {
				err = writer.WriteField("video", vd.Video)
			} else {
				if err = vd.checkTypeOfFile(); err == nil {
					err = writeFileToMultipart(writer, "video", vd.Video)
				}
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "video")
	}
	if err == nil && vd.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", vd.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.Width != 0 {
		err = writer.WriteField("width", fmt.Sprintf("%d", vd.Width))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.Height != 0 {
		err = writer.WriteField("height", fmt.Sprintf("%d", vd.Height))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.Thumbnail != "" {
		err = addThumbnail(writer, vd.Thumbnail, vd.ThumbnailGottenFrom)
	}
	if err == nil && vd.ShowCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.SupportsStreaming {
		err = writer.WriteField("supports_streaming", "True")
	}
	if group != nil {
		(*group)[i] = vd
	}
	return err
}

func (vd *video) jsonFileds() ([]byte, error) {
	return json.Marshal(vd)
}

func (vd *video) nameAndConst() (string, int) {
	return interfaceVideo, constVideo
}

func (ad *audio) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	err := commonFields(writer, ad.CaptionEntities)
	if err == nil {
		if input {
			if ad.AudioGottenFrom != Storage {
				ad.Media = ad.Audio
			} else {
				if err = ad.checkTypeOfFile(); err == nil {
					err = writeFileToMultipart(writer, ad.Audio, ad.Audio)
					ad.Media = fmt.Sprintf("attach://%s", ad.Audio)
				}
			}
		} else {
			if ad.AudioGottenFrom != Storage {
				err = writer.WriteField("audio", ad.Audio)
			} else {
				if err = ad.checkTypeOfFile(); err == nil {
					err = writeFileToMultipart(writer, "audio", ad.Audio)
				}
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "audio")
	}
	if err == nil && ad.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", ad.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && ad.Performer != "" {
		err = writer.WriteField("performer", ad.Performer)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && ad.Thumbnail != "" {
		err = addThumbnail(writer, ad.Thumbnail, ad.ThumbnailGottenFrom)
	}
	if group != nil {
		(*group)[i] = ad
	}
	return err
}

func (ad *audio) jsonFileds() ([]byte, error) {
	return json.Marshal(ad)
}

func (ad *audio) nameAndConst() (string, int) {
	return interfaceAudio, constAudio
}

func (dc *document) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	err := commonFields(writer, dc.CaptionEntities)
	if err == nil {
		if input {
			if dc.DocumentGottenFrom != Storage {
				dc.Media = dc.Document
			} else {
				err = writeFileToMultipart(writer, dc.Document, dc.Document)
				dc.Media = fmt.Sprintf("attach://%s", dc.Document)
			}
		} else {
			if dc.DocumentGottenFrom != Storage {
				err = writer.WriteField("document", dc.Document)
			} else {
				err = writeFileToMultipart(writer, "document", dc.Document)
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "document")
	}
	if err == nil && dc.Thumbnail != "" {
		err = addThumbnail(writer, dc.Thumbnail, dc.ThumbnailGottenFrom)
	}
	if err == nil && dc.DisableContentTypeDetection {
		err = writer.WriteField("disable_content_type_detection", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if group != nil {
		(*group)[i] = dc
	}
	return err
}

func (dc *document) jsonFileds() ([]byte, error) {
	return json.Marshal(dc)
}

func (dc *document) nameAndConst() (string, int) {
	return interfaceDocument, constDoc
}

func (an *animation) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	if input {
		if an.AnimationGottenFrom != Storage {
			an.Media = an.Animation
		} else {
			if err = an.checkTypeOfFile(); err == nil {
				err = writeFileToMultipart(writer, an.Animation, an.Animation)
				an.Media = fmt.Sprintf("attach://%s", an.Animation)
			}
		}
	} else {
		if an.AnimationGottenFrom != Storage {
			err = writer.WriteField("animation", an.Animation)
		} else {
			if err = an.checkTypeOfFile(); err == nil {
				err = writeFileToMultipart(writer, "animation", an.Animation)
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "animation")
	}
	if err == nil && an.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", an.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && an.Width != 0 {
		err = writer.WriteField("width", fmt.Sprintf("%d", an.Width))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && an.Height != 0 {
		err = writer.WriteField("height", fmt.Sprintf("%d", an.Height))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && an.Thumbnail != "" {
		err = addThumbnail(writer, an.Thumbnail, an.ThumbnailGottenFrom)
	}
	if err == nil && an.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if group != nil {
		(*group)[i] = an
	}
	return err
}

func (an *animation) jsonFileds() ([]byte, error) {
	return json.Marshal(an)
}

func (an *animation) nameAndConst() (string, int) {
	return interfaceAnimation, constAnim
}

func (vc *voice) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	if input {
		if vc.gottenFrom != Storage {
			vc.Media = vc.Voice
		} else {
			if err = vc.checkTypeOfFile(); err == nil {
				err = writeFileToMultipart(writer, vc.Voice, vc.Voice)
				vc.Media = fmt.Sprintf("attach://%s", vc.Voice)
			}
		}
	} else {
		if vc.gottenFrom != Storage {
			err = writer.WriteField("voice", vc.Voice)
		} else {
			if err = vc.checkTypeOfFile(); err == nil {
				err = writeFileToMultipart(writer, "voice", vc.Voice)
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "voice")
	}
	if err == nil && vc.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", vc.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if group != nil {
		(*group)[i] = vc
	}
	return err
}

func (vc *voice) jsonFileds() ([]byte, error) {
	return json.Marshal(vc)
}

func (*voice) nameAndConst() (string, int) {
	return interfaceVoice, constVoice
}

func (vdn *videonote) multipartFields(writer *multipart.Writer, group *[]interface{}, i int, input bool) error {
	var err error
	if input {
		if vdn.videoGottenFrom != Storage {
			vdn.Media = vdn.VideoNote
		} else {
			if err = vdn.checkTypeOfFile(); err == nil {
				err = writeFileToMultipart(writer, vdn.VideoNote, vdn.VideoNote)
				vdn.Media = fmt.Sprintf("attach://%s", vdn.VideoNote)
			}
		}
	} else {
		if vdn.videoGottenFrom != Storage {
			err = writer.WriteField("video_note", vdn.VideoNote)
		} else {
			if err = vdn.checkTypeOfFile(); err == nil {
				err = writeFileToMultipart(writer, "video_note", vdn.VideoNote)
			}
		}
	}
	if err == nil && input {
		err = writer.WriteField("type", "video_note")
	}
	if err == nil && vdn.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", vdn.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vdn.Length != 0 {
		err = writer.WriteField("length", fmt.Sprintf("%d", vdn.Length))

		if err != nil {
			errors.CantWriteField(err)
		}
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

func (*videonote) nameAndConst() (string, int) {
	return interfaceVideoNote, constVideoNote
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
		errors.CantWriteField(err)
	}
	return err
}

func (inf *information) multipartFields(writer *multipart.Writer) error {
	var err error
	if inf.Text != "" {
		err = field(writer, "text", inf.Text)
	}
	if err == nil && inf.Caption != "" {
		err = field(writer, "caption", inf.Caption)
	}
	if err == nil && inf.ParseMode != "" {
		err = field(writer, "parse_mode", inf.ParseMode)
	}
	if err == nil && inf.MessageThreadID != 0 {
		err = field(writer, "message_thread_id", fmt.Sprint(inf.MessageThreadID))
	}
	if err == nil && inf.DisableNotification {
		err = field(writer, "disable_notification", "True")
	}
	if err == nil && inf.ProtectContent {
		err = field(writer, "protect_content", "True")
	}
	if err == nil && inf.MessageEffectID != "" {
		err = field(writer, "message_effect_id", inf.MessageEffectID)
	}
	return err
}

func (inf *information) createJSON() ([]byte, error) {
	body, err := json.Marshal(inf)
	if err != nil {
		errors.CantMarshalJSON(err)
	}
	return body, err
}

func (ch *chat) multipartFields(writer *multipart.Writer) error {
	var err error
	if ch.ID != nil {
		err = field(writer, "chat_id", fmt.Sprint(ch.ID))

	} else if ch.ID == nil {
		err = errors.ChatIDIsMissed()
	}
	if ch.BusinessConnectionID != "" {
		err = field(writer, "business_connection_id", ch.BusinessConnectionID)
	}
	return err
}

func (ch *chat) createJson() ([]byte, error) {
	return json.Marshal(ch)
}

func (in *inline) MultipartFields(writer *multipart.Writer) error {
	inbody, err := json.Marshal(in.Keyboard.InlineKeyboard)
	if err == nil {
		err = writer.WriteField("reply_markup", string(inbody))
	}
	return err
}

func (in *inline) JsonFields() ([]byte, error) {
	return json.Marshal(in)
}

func (rp *reply) MultipartFields(writer *multipart.Writer) error {
	rpbody, err := json.Marshal(rp.Keyboard)
	if err == nil {
		err = writer.WriteField("reply_markup", string(rpbody))
	}
	return err
}

func (rp *reply) JsonFields() ([]byte, error) {
	return json.Marshal(rp)
}
