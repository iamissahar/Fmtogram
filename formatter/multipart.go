package formatter

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
)

func writeFileToMultipart(writer *multipart.Writer, fieldname, filename string) error {
	var part io.Writer
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()

		part, err = writer.CreateFormFile(fieldname, filename)

		if err == nil {
			_, err = io.Copy(part, file)
		}
	}
	return err
}

func (ph *photo) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	var fieldname string
	if ph.gottenFrom == Storage {
		if !isgroup {
			fieldname = "photo"
			ph.Media = ""
			ph.Type = ""
		} else {
			fieldname = ph.Photo
		}
		err = writeFileToMultipart(wr, fieldname, ph.Photo)
		ph.Photo = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (ph *photo) uniqueConst() int {
	return constPhoto
}

func (vd *video) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	var fieldname string
	if vd.videoGottenFrom == Storage {
		if !isgroup {
			fieldname = "video"
			vd.Media = ""
			vd.Type = ""
		} else {
			fieldname = vd.Video
		}
		err = writeFileToMultipart(wr, fieldname, vd.Video)
		vd.Video = ""
		*cnttype = wr.FormDataContentType()
	}
	if vd.thumbnailGottenFrom == Storage {
		err = writeFileToMultipart(wr, "thumbnail", vd.Thumbnail)
		vd.Thumbnail = ""
		*cnttype = wr.FormDataContentType()
	}
	if vd.coverGottenFrom == Storage {
		err = writeFileToMultipart(wr, "cover", vd.Cover)
		vd.Cover = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (vd *video) uniqueConst() int {
	return constVideo
}

func (ad *audio) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	var fieldname string
	if ad.audioGottenFrom == Storage {
		if !isgroup {
			fieldname = "audio"
			ad.Media = ""
			ad.Type = ""
		} else {
			fieldname = ad.Audio
		}
		err = writeFileToMultipart(wr, fieldname, ad.Audio)
		ad.Audio = ""
		*cnttype = wr.FormDataContentType()
	}
	if ad.thumbnailGottenFrom == Storage {
		err = writeFileToMultipart(wr, "thumbnail", ad.Thumbnail)
		ad.Thumbnail = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (ad *audio) uniqueConst() int {
	return constAudio
}

func (dc *document) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	var fieldname string
	if dc.documentGottenFrom == Storage {
		if !isgroup {
			fieldname = "document"
			dc.Media = ""
			dc.Type = ""
		} else {
			fieldname = dc.Document
		}
		err = writeFileToMultipart(wr, fieldname, dc.Document)
		dc.Document = ""
		*cnttype = wr.FormDataContentType()
	}
	if dc.thumbnailGottenFrom == Storage {
		err = writeFileToMultipart(wr, "thumbnail", dc.Thumbnail)
		dc.Thumbnail = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (dc *document) uniqueConst() int {
	return constDoc
}

func (an *animation) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	if an.animationGottenFrom == Storage {
		err = writeFileToMultipart(wr, "animation", an.Animation)
		an.Animation = ""
		*cnttype = wr.FormDataContentType()
	}
	if an.thumbnailGottenFrom == Storage {
		err = writeFileToMultipart(wr, "thumbnail", an.Thumbnail)
		an.Thumbnail = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (an *animation) uniqueConst() int {
	return constAnim
}

func (vc *voice) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	if vc.gottenFrom == Storage {
		err = writeFileToMultipart(wr, "voice", vc.Voice)
		vc.Voice = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (*voice) uniqueConst() int {
	return constVoice
}

func (vdn *videonote) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	if vdn.videoGottenFrom == Storage {
		err = writeFileToMultipart(wr, "video_note", vdn.VideoNote)
		vdn.VideoNote = ""
		*cnttype = wr.FormDataContentType()
	}
	if vdn.thumbnailGottenFrom == Storage {
		err = writeFileToMultipart(wr, "thumbnail", vdn.Thumbnail)
		vdn.Thumbnail = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
}

func (*videonote) uniqueConst() int {
	return constVideoNote
}

func (st *sticker) proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error {
	var err error
	if st.stickerGottenFrom == Storage {
		err = writeFileToMultipart(wr, st.Sticker, st.Sticker)
		st.Sticker = ""
		*cnttype = wr.FormDataContentType()
	}
	if st.thumbnailGottenFrom == Storage {
		err = writeFileToMultipart(wr, "thumbnail", st.Thumbnail)
		st.Thumbnail = ""
		*cnttype = wr.FormDataContentType()
	}
	return err
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
