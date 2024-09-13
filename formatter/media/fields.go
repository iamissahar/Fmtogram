package media

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	Storage  int = 1
	Telegram int = 2
	Internet int = 3
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

func (ph *Photo) MultipartFields(writer *multipart.Writer, group []interface{}, i int, input bool) (string, error) {
	err := commonFields(writer, ph.CaptionEntities)
	if err == nil {
		if input {
			err = writeFileToMultipart(writer, ph.Photo, ph.Photo)
			ph.Media = fmt.Sprintf("attach://%s", ph.Photo)
		} else {
			err = writeFileToMultipart(writer, "photo", ph.Photo)
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
		group[i] = ph
	}
	return methods.Photo, err
}

func (ph *Photo) JsonFileds() (string, []byte, error) {
	phjson, err := json.Marshal(ph)
	return methods.Photo, phjson, err
}

func (ph *Photo) CheckGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if ph.gottenFrom == 0 {
		err = errors.MissedGottenFrom(objectPhoto, "Gotten From", numberInQueue)
	}
	if ph.gottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func (vd *Photo) CheckThumbnailGottenFrom(numberInQueue int) (bool, error) {
	return false, nil
}

func (vd *Video) MultipartFields(writer *multipart.Writer, group []interface{}, i int, input bool) (string, error) {
	err := commonFields(writer, vd.CaptionEntities)
	if err == nil {
		if input {
			err = writeFileToMultipart(writer, vd.Video, vd.Video)
			vd.Media = fmt.Sprintf("attach://%s", vd.Video)
		} else {
			err = writeFileToMultipart(writer, "video", vd.Video)
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
		group[i] = vd
	}
	return methods.Video, err
}

func (vd *Video) JsonFileds() (string, []byte, error) {
	vidjson, err := json.Marshal(vd)
	return methods.Video, vidjson, err
}

func (vd *Video) CheckGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if vd.VideoGottenFrom == 0 {
		err = errors.MissedGottenFrom(objectVideo, "Gotten From", numberInQueue)
	}
	if vd.VideoGottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func (vd *Video) CheckThumbnailGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if vd.Thumbnail != "" {
		if vd.ThumbnailGottenFrom == 0 {
			err = errors.MissedGottenFrom(objectVideo, "Thumbnail Gotten From", numberInQueue)
		}
	}
	if vd.ThumbnailGottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func (ad *Audio) MultipartFields(writer *multipart.Writer, group []interface{}, i int, input bool) (string, error) {
	err := commonFields(writer, ad.CaptionEntities)
	if err == nil {
		if input {
			err = writeFileToMultipart(writer, ad.Audio, ad.Audio)
			ad.Media = fmt.Sprintf("attach://%s", ad.Audio)
		} else {
			err = writeFileToMultipart(writer, "audio", ad.Audio)
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
		group[i] = ad
	}
	return methods.Audio, err
}

func (ad *Audio) JsonFileds() (string, []byte, error) {
	adjson, err := json.Marshal(ad)
	return methods.Video, adjson, err
}

func (ad *Audio) CheckGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if ad.AudioGottenFrom == 0 {
		err = errors.MissedGottenFrom(objectAudio, "Gotten From", numberInQueue)
	}
	if ad.AudioGottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func (ad *Audio) CheckThumbnailGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if ad.Thumbnail != "" {
		if ad.ThumbnailGottenFrom == 0 {
			err = errors.MissedGottenFrom(objectAudio, "Thumbnail Gotten From", numberInQueue)
		}
	}
	if ad.ThumbnailGottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func (dc *Document) MultipartFields(writer *multipart.Writer, group []interface{}, i int, input bool) (string, error) {
	err := commonFields(writer, dc.CaptionEntities)
	if err == nil {
		if input {
			err = writeFileToMultipart(writer, dc.Document, dc.Document)
			dc.Media = fmt.Sprintf("attach://%s", dc.Document)
		} else {
			err = writeFileToMultipart(writer, "document", dc.Document)
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
		group[i] = dc
	}
	return methods.Document, err
}

func (dc *Document) JsonFileds() (string, []byte, error) {
	dcjson, err := json.Marshal(dc)
	return methods.Video, dcjson, err
}

func (dc *Document) CheckGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if dc.DocumentGottenFrom == 0 {
		err = errors.MissedGottenFrom(objectDocument, "Gotten From", numberInQueue)
	}
	if dc.DocumentGottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func (dc *Document) CheckThumbnailGottenFrom(numberInQueue int) (bool, error) {
	var (
		fromStorage bool
		err         error
	)
	if dc.Thumbnail != "" {
		if dc.ThumbnailGottenFrom == 0 {
			err = errors.MissedGottenFrom(objectDocument, "Thumbnail Gotten From", numberInQueue)
		}
	}
	if dc.ThumbnailGottenFrom == Storage {
		fromStorage = true
	}
	return fromStorage, err
}

func Group(writer *multipart.Writer, group []interface{}) error {
	log.Print(group)
	mediaJSON, err := json.Marshal(group)
	if err == nil && len(mediaJSON) != 0 {
		err = writer.WriteField("media", string(mediaJSON))
	}
	log.Print(string(mediaJSON))
	return err
}
