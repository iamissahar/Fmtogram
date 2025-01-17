package formatter

import (
	"fmt"
	"time"

	"github.com/l1qwie/Fmtogram/fmerrors"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

type emojich struct {
	permissibleQuantity int
	amount              int
}

func code01() error {
	err := new(fmerrors.FME)
	err.Code = 1
	err.String = "there isnt a place to put a data in"
	return err
}

func code3() error {
	err := new(fmerrors.FME)
	err.Code = 3
	err.String = "media holder is full"
	return err
}

func code5() error {
	err := new(fmerrors.FME)
	err.Code = 5
	err.String = "incorrect data in an input slice"
	return err
}

func code10() error {
	err := new(fmerrors.FME)
	err.Code = 10
	err.String = "the data is already present"
	return err
}

func code12() error {
	err := new(fmerrors.FME)
	err.Code = 12
	err.String = "incorrect type of file"
	return err
}

func code14() error {
	err := new(fmerrors.FME)
	err.Code = 14
	err.String = "media method is specified, media data isn't"
	return err
}

func code20() error {
	err := new(fmerrors.FME)
	err.Code = 20
	err.String = "incorrect input data"
	return err
}

func code21() error {
	err := new(fmerrors.FME)
	err.Code = 21
	err.String = "missed required data"
	return err
}

func code25() error {
	err := new(fmerrors.FME)
	err.Code = 25
	err.String = "incompatible data"
	return err
}

func code54() error {
	err := new(fmerrors.FME)
	err.Code = 54
	err.String = "isn't allowed to be united with others"
	return err
}

func checkCaption(caption, cell string) error {
	var err error
	cap := len([]rune(caption))
	if (cap > 0) && (cap <= 1024) {
		if cell != "" {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func checkParseMode(parsemode, cell string) error {
	var err error
	if (parsemode != types.HTML) && (parsemode != types.Markdown) && (parsemode != types.MarkdownV2) {
		err = code20()
	} else {
		if cell != "" {
			err = code10()
		}
	}
	return err
}

func checkInputPlaceHolder(placeholder, cell string) error {
	var err error
	p := []rune(placeholder)
	if len(p) > 0 && len(p) <= 64 {
		if cell != "" {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func checkEntities(entities, cell []*types.MessageEntity) error {
	var err error
	if len(entities) != 0 {
		for i := 0; (i < len(entities)) && (err == nil); i++ {
			if entities[i] == nil {
				err = code5()
			}
		}
		if err == nil {
			if len(cell) != 0 {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func checkStringValue(val, cell string) error {
	var err error
	if val != "" {
		if cell != "" {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func checkIntegerValue(val, cell int) error {
	var err error
	if val > 0 {
		if cell != 0 {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ph *photo) WritePhotoStorage(photo string) error {
	var err error
	if err = checkStringValue(photo, ph.Photo); err == nil {
		if err = ph.isCorrectType(photo); err == nil {
			ph.gottenFrom = Storage
			ph.Photo, ph.Media = photo, photo
			logs.DataWrittenSuccessfully(interfacePhoto, "Photo From Storage")
		}
	}
	return err
}

func (ph *photo) WritePhotoTelegram(photo string) error {
	var err error
	if err = checkStringValue(photo, ph.Photo); err == nil {
		ph.gottenFrom = Telegram
		ph.Photo, ph.Media = photo, photo
		logs.DataWrittenSuccessfully(interfacePhoto, "Photo From Telegram")

	}
	return err
}

func (ph *photo) WritePhotoInternet(photo string) error {
	var err error
	if err = checkStringValue(photo, ph.Photo); err == nil {
		ph.gottenFrom = Internet
		ph.Photo, ph.Media = photo, photo
		logs.DataWrittenSuccessfully(interfacePhoto, "Photo From Internet")

	}
	return err
}

func (ph *photo) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, ph.Caption); err == nil {
		ph.Caption = caption
		logs.DataWrittenSuccessfully(interfacePhoto, "Caption")
	}
	return err
}

func (ph *photo) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, ph.ParseMode); err == nil {
		ph.ParseMode = parsemode
		logs.DataWrittenSuccessfully(interfacePhoto, "Parse Mode")
	}
	return err
}

func (ph *photo) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, ph.CaptionEntities); err == nil {
		ph.CaptionEntities = captionEntities
		logs.DataWrittenSuccessfully(interfacePhoto, "Caption Entities")
	}
	return err
}

func (ph *photo) WriteShowCaptionAboveMedia() error {
	var err error
	if ph.ShowCaptionAboveMedia {
		err = code10()
	} else {
		ph.ShowCaptionAboveMedia = true
		logs.SettedParam("Show Caption Above Media", interfacePhoto, ph.ShowCaptionAboveMedia)
	}
	return err
}

func (ph *photo) WriteHasSpoiler() error {
	var err error
	if ph.HasSpoiler {
		err = code10()
	} else {
		ph.HasSpoiler = true
		logs.SettedParam("Has Spoiler", interfacePhoto, ph.HasSpoiler)
	}
	return err
}

func (ph *photo) GetResponse() [4]types.PhotoSize {
	return ph.response
}

func (vd *video) WriteVideoStorage(video string) error {
	var err error
	if err = checkStringValue(video, vd.Video); err == nil {
		if err = vd.isCorrectType(video); err == nil {
			vd.videoGottenFrom = Storage
			vd.Video, vd.Media = video, video
			logs.DataWrittenSuccessfully(interfaceVideo, "Photo From Storage")
		}
	}
	return err
}

func (vd *video) WriteVideoTelegram(video string) error {
	var err error
	if err = checkStringValue(video, vd.Video); err == nil {
		vd.videoGottenFrom = Telegram
		vd.Video, vd.Media = video, video
		logs.DataWrittenSuccessfully(interfaceVideo, "Photo From Telegram")
	}
	return err
}

func (vd *video) WriteVideoInternet(video string) error {
	var err error
	if err = checkStringValue(video, vd.Video); err == nil {
		vd.videoGottenFrom = Internet
		vd.Video, vd.Media = video, video
		logs.DataWrittenSuccessfully(interfaceVideo, "Photo From Internet")
	}
	return err
}

func (vd *video) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, vd.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			vd.thumbnailGottenFrom = Storage
			vd.Thumbnail = thumbnail
			logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail From Storage")
		}
	}
	return err
}

func (vd *video) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, vd.Thumbnail); err == nil {
		vd.thumbnailGottenFrom = Telegram
		vd.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail From Telegram")

	}
	return err
}

func (vd *video) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, vd.Thumbnail); err == nil {
		vd.thumbnailGottenFrom = Internet
		vd.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail From Internet")

	}
	return err
}

func (vd *video) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, vd.Caption); err == nil {
		vd.Caption = caption
		logs.DataWrittenSuccessfully(interfaceVideo, "Caption")
	}
	return err
}

func (vd *video) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, vd.ParseMode); err == nil {
		vd.ParseMode = parsemode
		logs.DataWrittenSuccessfully(interfaceVideo, "Parse Mode")
	}
	return err
}

func (vd *video) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, vd.CaptionEntities); err == nil {
		vd.CaptionEntities = captionEntities
		logs.DataWrittenSuccessfully(interfaceVideo, "Caption Entities")
	}
	return err
}

func (vd *video) WriteShowCaptionAboveMedia() error {
	var err error
	if !vd.ShowCaptionAboveMedia {
		vd.ShowCaptionAboveMedia = true
		logs.SettedParam("Show Caption Above Media", "Video", vd.ShowCaptionAboveMedia)
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteWidth(width int) error {
	var err error
	if err = checkIntegerValue(width, vd.Width); err == nil {
		vd.Width = width
		logs.DataWrittenSuccessfully(interfaceVideo, "Width")
	}
	return err
}

func (vd *video) WriteHeight(height int) error {
	var err error
	if err = checkIntegerValue(height, vd.Height); err == nil {
		vd.Height = height
		logs.DataWrittenSuccessfully(interfaceVideo, "Height")
	}
	return err
}

func (vd *video) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, vd.Duration); err == nil {
		vd.Duration = duration
		logs.DataWrittenSuccessfully(interfaceVideo, "Duration")
	}
	return err
}

func (vd *video) WriteSupportStreaming() error {
	var err error
	if !vd.SupportsStreaming {
		vd.SupportsStreaming = true
		logs.SettedParam("Supports Streaming", "Video", vd.SupportsStreaming)
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteHasSpoiler() error {
	var err error
	if !vd.HasSpoiler {
		vd.HasSpoiler = true
		logs.SettedParam("Has Spoiler", "Video", vd.SupportsStreaming)
	} else {
		err = code10()
	}
	return err
}

func (vd *video) GetResponse() types.Video {
	return vd.response
}

func (ad *audio) WriteAudioStorage(audio string) error {
	var err error
	if err = checkStringValue(audio, ad.Audio); err == nil {
		if err = ad.isCorrectType(audio); err == nil {
			ad.audioGottenFrom = Storage
			ad.Audio, ad.Media = audio, audio
			logs.DataWrittenSuccessfully(interfaceAudio, "Audio From Storage")
		}
	}
	return err
}

func (ad *audio) WriteAudioTelegram(audio string) error {
	var err error
	if err = checkStringValue(audio, ad.Audio); err == nil {
		ad.audioGottenFrom = Telegram
		ad.Audio, ad.Media = audio, audio
		logs.DataWrittenSuccessfully(interfaceAudio, "Audio From Telegram")
	}
	return err
}

func (ad *audio) WriteAudioInternet(audio string) error {
	var err error
	if err = checkStringValue(audio, ad.Audio); err == nil {
		ad.audioGottenFrom = Internet
		ad.Audio, ad.Media = audio, audio
		logs.DataWrittenSuccessfully(interfaceAudio, "Audio From Internet")
	}
	return err
}

func (ad *audio) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, ad.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			ad.thumbnailGottenFrom = Storage
			ad.Thumbnail = thumbnail
			logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail From Storage")
		}
	}
	return err
}

func (ad *audio) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, ad.Thumbnail); err == nil {
		ad.thumbnailGottenFrom = Telegram
		ad.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail From Telegram")
	}
	return err
}

func (ad *audio) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, ad.Thumbnail); err == nil {
		ad.thumbnailGottenFrom = Internet
		ad.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail From Internet")
	}
	return err
}

func (ad *audio) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, ad.Caption); err == nil {
		ad.Caption = caption
		logs.DataWrittenSuccessfully(interfaceAudio, "Caption")
	}
	return err
}

func (ad *audio) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, ad.ParseMode); err == nil {
		ad.ParseMode = parsemode
		logs.DataWrittenSuccessfully(interfaceAudio, "Parse Mode")
	}
	return err
}

func (ad *audio) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, ad.CaptionEntities); err == nil {
		ad.CaptionEntities = captionEntities
		logs.DataWrittenSuccessfully(interfaceAudio, "Caption Entities")
	}
	return err
}

func (ad *audio) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, ad.Duration); err == nil {
		ad.Duration = duration
		logs.DataWrittenSuccessfully(interfaceAudio, "Duration")
	}
	return err
}

func (ad *audio) WritePerformer(performer string) error {
	var err error
	if err = checkStringValue(performer, ad.Performer); err == nil {
		ad.Performer = performer
		logs.DataWrittenSuccessfully(interfaceAudio, "Performer")
	}
	return err
}

func (ad *audio) WriteTitle(title string) error {
	var err error
	if err = checkStringValue(title, ad.Title); err == nil {
		ad.Title = title
		logs.DataWrittenSuccessfully(interfaceAudio, "Title")
	}
	return err
}

func (ad *audio) GetResponse() types.Audio {
	return ad.response
}

func (dc *document) WriteDocumentStorage(document string) error {
	var err error
	if err = checkStringValue(document, dc.Document); err == nil {
		dc.documentGottenFrom = Storage
		dc.Document, dc.Media = document, document
		logs.DataWrittenSuccessfully(interfaceDocument, "Document From Storage")
	}
	return err
}

func (dc *document) WriteDocumentTelegram(document string) error {
	var err error
	if err = checkStringValue(document, dc.Document); err == nil {
		dc.documentGottenFrom = Telegram
		dc.Document, dc.Media = document, document
		logs.DataWrittenSuccessfully(interfaceDocument, "Document From Telegram")
	}
	return err
}

func (dc *document) WriteDocumentInternet(document string) error {
	var err error
	if err = checkStringValue(document, dc.Document); err == nil {
		dc.documentGottenFrom = Internet
		dc.Document, dc.Media = document, document
		logs.DataWrittenSuccessfully(interfaceDocument, "Document From Internet")
	}
	return err
}

func (dc *document) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, dc.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			dc.thumbnailGottenFrom = Storage
			dc.Thumbnail = thumbnail
			logs.DataWrittenSuccessfully(interfaceDocument, "Thumbnail From Storage")
		}
	}
	return err
}

func (dc *document) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, dc.Thumbnail); err == nil {
		dc.thumbnailGottenFrom = Telegram
		dc.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceDocument, "Thumbnail From Telegram")
	}
	return err
}

func (dc *document) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, dc.Thumbnail); err == nil {
		dc.thumbnailGottenFrom = Internet
		dc.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceDocument, "Thumbnail From Internet")
	}
	return err
}

func (dc *document) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, dc.Caption); err == nil {
		dc.Caption = caption
		logs.DataWrittenSuccessfully(interfaceDocument, "Caption")
	}
	return err
}

func (dc *document) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, dc.ParseMode); err == nil {
		dc.ParseMode = parsemode
		logs.DataWrittenSuccessfully(interfaceAudio, "Parse Mode")
	}
	return err
}

func (dc *document) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, dc.CaptionEntities); err == nil {
		dc.CaptionEntities = captionEntities
		logs.DataWrittenSuccessfully(interfaceDocument, "Caption Entities")
	}
	return err
}

func (dc *document) WriteDisableContentTypeDetection() error {
	var err error
	if !dc.DisableContentTypeDetection {
		dc.DisableContentTypeDetection = true
		logs.SettedParam("Disable Content Type Detection", interfaceDocument, dc.DisableContentTypeDetection)
	} else {
		err = code10()
	}
	return err
}

func (dc *document) GetResponse() types.Document {
	return dc.response
}

func (an *animation) WriteAnimationStorage(animation string) error {
	var err error
	if err = checkStringValue(animation, an.Animation); err == nil {
		if err = an.isCorrectType(animation); err == nil {
			an.animationGottenFrom = Storage
			an.Animation = animation
			logs.DataWrittenSuccessfully(interfaceAnimation, "Animation From Storage")
		}
	}
	return err
}

func (an *animation) WriteAnimationTelegram(animation string) error {
	var err error
	if err = checkStringValue(animation, an.Animation); err == nil {
		an.animationGottenFrom = Telegram
		an.Animation = animation
		logs.DataWrittenSuccessfully(interfaceAnimation, "Animation From Telegram")
	}
	return err
}

func (an *animation) WriteAnimationInternet(animation string) error {
	var err error
	if err = checkStringValue(animation, an.Animation); err == nil {
		an.animationGottenFrom = Internet
		an.Animation = animation
		logs.DataWrittenSuccessfully(interfaceAnimation, "Animation From Internet")
	}
	return err
}

func (an *animation) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, an.Duration); err == nil {
		an.Duration = duration
		logs.DataWrittenSuccessfully(interfaceAnimation, "Duration")
	}
	return err
}

func (an *animation) WriteWidth(width int) error {
	var err error
	if err = checkIntegerValue(width, an.Width); err == nil {
		an.Width = width
		logs.DataWrittenSuccessfully(interfaceAnimation, "Width")
	}
	return err
}

func (an *animation) WriteHeight(height int) error {
	var err error
	if err = checkIntegerValue(height, an.Height); err == nil {
		an.Height = height
		logs.DataWrittenSuccessfully(interfaceAnimation, "Height")
	}
	return err
}

func (an *animation) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, an.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			an.thumbnailGottenFrom = Storage
			an.Thumbnail = thumbnail
			logs.DataWrittenSuccessfully(interfaceAnimation, "Thumbnail From Storage")
		}
	}
	return err
}

func (an *animation) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, an.Thumbnail); err == nil {
		an.thumbnailGottenFrom = Telegram
		an.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceAnimation, "Thumbnail From Telegram")
	}
	return err
}

func (an *animation) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, an.Thumbnail); err == nil {
		an.thumbnailGottenFrom = Internet
		an.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceAnimation, "Thumbnail From Internet")
	}
	return err
}

func (an *animation) WriteHasSpoiler() error {
	var err error
	if !an.HasSpoiler {
		an.HasSpoiler = true
		logs.DataWrittenSuccessfully(interfaceAnimation, "Has Spoiler")
	} else {
		err = code10()
	}
	return err
}

func (an *animation) GetResponse() types.Animation {
	return an.response
}

func (vc *voice) WriteVoiceStorage(voice string) error {
	var err error
	if err = checkStringValue(voice, vc.Voice); err == nil {
		if err = vc.isCorrectType(voice); err == nil {
			vc.gottenFrom = Storage
			vc.Voice = voice
			logs.DataWrittenSuccessfully(interfaceVoice, "Voice From Storage")
		}
	}
	return err
}

func (vc *voice) WriteVoiceTelegram(voiceID string) error {
	var err error
	if err = checkStringValue(voiceID, vc.Voice); err == nil {
		vc.gottenFrom = Telegram
		vc.Voice = voiceID
		logs.DataWrittenSuccessfully(interfaceVoice, "Voice From Telegram")
	}
	return err
}

func (vc *voice) WriteVoiceInternet(URL string) error {
	var err error
	if err = checkStringValue(URL, vc.Voice); err == nil {
		vc.gottenFrom = Internet
		vc.Voice = URL
		logs.DataWrittenSuccessfully(interfaceVoice, "Voice From Internet")
	}
	return err
}

func (vc *voice) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, vc.Duration); err == nil {
		vc.Duration = duration
		logs.DataWrittenSuccessfully(interfaceVoice, "Duration")
	}
	return err
}

func (vc *voice) GetResponse() types.Voice {
	return vc.response
}

func (vdn *videonote) WriteVideoNoteStorage(videonote string) error {
	var err error
	if err = checkStringValue(videonote, vdn.VideoNote); err == nil {
		if err = vdn.isCorrectType(videonote); err == nil {
			vdn.videoGottenFrom = Storage
			vdn.VideoNote, vdn.Media = videonote, videonote
			logs.DataWrittenSuccessfully(interfaceVoice, "Video-Note From Storage")
		}
	}
	return err
}

func (vdn *videonote) WriteVideoNoteTelegram(vdnID string) error {
	var err error
	if err = checkStringValue(vdnID, vdn.VideoNote); err == nil {
		vdn.videoGottenFrom = Telegram
		vdn.VideoNote, vdn.Media = vdnID, vdnID
		logs.DataWrittenSuccessfully(interfaceVoice, "Video-Note From Telegram")
	}
	return err
}

func (vdn *videonote) WriteVideoNoteInternet(URL string) error {
	var err error
	if err = checkStringValue(URL, vdn.VideoNote); err == nil {
		vdn.videoGottenFrom = Internet
		vdn.VideoNote, vdn.Media = URL, URL
		logs.DataWrittenSuccessfully(interfaceVoice, "Video-Note From Internet")
	}
	return err
}

func (vdn *videonote) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, vdn.Duration); err == nil {
		vdn.Duration = duration
		logs.DataWrittenSuccessfully(interfaceVideoNote, "Duration")
	}
	return err
}

func (vdn *videonote) WriteLength(length int) error {
	var err error
	if err = checkIntegerValue(length, vdn.Length); err == nil {
		vdn.Length = length
		logs.DataWrittenSuccessfully(interfaceVideoNote, "Length")
	}
	return err
}

func (vdn *videonote) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, vdn.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			vdn.thumbnailGottenFrom = Storage
			vdn.Thumbnail = thumbnail
			logs.DataWrittenSuccessfully(interfaceVideoNote, "Thumbnail From Storage")
		}
	}
	return err
}

func (vdn *videonote) WriteThumbnailTelegram(thumbnailID string) error {
	var err error
	if err = checkStringValue(thumbnailID, vdn.Thumbnail); err == nil {
		vdn.thumbnailGottenFrom = Telegram
		vdn.Thumbnail = thumbnailID
		logs.DataWrittenSuccessfully(interfaceVideoNote, "Thumbnail From Telegram")
	}
	return err
}

func (vdn *videonote) WriteThumbnailInternet(URL string) error {
	var err error
	if err = checkStringValue(URL, vdn.Thumbnail); err == nil {
		vdn.thumbnailGottenFrom = Internet
		vdn.Thumbnail = URL
		logs.DataWrittenSuccessfully(interfaceVideoNote, "Thumbnail From Internet")
	}
	return err
}

func (vdn *videonote) GetResponse() types.VideoNote {
	return vdn.response
}

func (loc *location) WriteLatitude(lat float64) error {
	var err error
	if (lat >= -90) && (lat <= 90) {
		if loc.Latitude == 0 {
			loc.Latitude = lat
			logs.DataWrittenSuccessfully(interfaceLocation, "Latitude")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (loc *location) WriteLongitude(long float64) error {
	var err error
	if (long >= -180) && (long <= 180) {
		if loc.Longitude == 0 {
			loc.Longitude = long
			logs.DataWrittenSuccessfully(interfaceLocation, "Longitude")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (loc *location) WriteHorizontalAccuracy(horacc float64) error {
	var err error
	if (horacc >= 0) && (horacc <= 1500) {
		if loc.HorizontalAccuracy == 0 {
			loc.HorizontalAccuracy = horacc
			logs.DataWrittenSuccessfully(interfaceLocation, "Horizontal Accuracy")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (loc *location) WriteLivePeriod(period int) error {
	var err error
	if ((period >= 60) && (period <= 86400)) || (period == 0x7FFFFFFF) {
		if loc.LivePeriod == 0 {
			loc.LivePeriod = period
			logs.DataWrittenSuccessfully(interfaceLocation, "Live Period")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (loc *location) WriteHeading(heading int) error {
	var err error
	if (heading >= 1) && (heading <= 360) {
		if loc.Heading == 0 {
			loc.Heading = heading
			logs.DataWrittenSuccessfully(interfaceLocation, "Heading")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (loc *location) WriteProximityAlertRadius(proxalrad int) error {
	var err error
	if (proxalrad >= 1) && (proxalrad <= 100000) {
		if loc.ProximityAlertRadius == 0 {
			loc.ProximityAlertRadius = proxalrad
			logs.DataWrittenSuccessfully(interfaceLocation, "Proximity Alert Radius")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (loc *location) WriteTitle(title string) error {
	var err error
	if err = checkStringValue(title, loc.Title); err == nil {
		loc.Title = title
		logs.DataWrittenSuccessfully(interfaceLocation, "Title")
	}
	return err
}

func (loc *location) WriteAddress(address string) error {
	var err error
	if err = checkStringValue(address, loc.Address); err == nil {
		loc.Address = address
		logs.DataWrittenSuccessfully(interfaceLocation, "Address")
	}
	return err
}

func (loc *location) WriteFoursquareID(foursquareID string) error {
	var err error
	if err = checkStringValue(foursquareID, loc.FoursquareID); err == nil {
		loc.FoursquareID = foursquareID
		logs.DataWrittenSuccessfully(interfaceLocation, "Foursquare ID")
	}
	return err
}

func (loc *location) WriteFoursquareType(foursquareType string) error {
	var err error
	if err = checkStringValue(foursquareType, loc.FoursquareType); err == nil {
		loc.FoursquareType = foursquareType
		logs.DataWrittenSuccessfully(interfaceLocation, "Foursquare Type")
	}
	return err
}

func (loc *location) WriteGooglePlaceID(googlePlaceID string) error {
	var err error
	if err = checkStringValue(googlePlaceID, loc.GooglePlaceID); err == nil {
		loc.GooglePlaceID = googlePlaceID
		logs.DataWrittenSuccessfully(interfaceLocation, "Google Place ID")
	}
	return err
}

func (loc *location) WriteGooglePlaceType(googlePlaceType string) error {
	var err error
	if err = checkStringValue(googlePlaceType, loc.GooglePlaceType); err == nil {
		loc.GooglePlaceType = googlePlaceType
		logs.DataWrittenSuccessfully(interfaceLocation, "Google Place Type")
	}
	return err
}

func (loc *location) GetResponse() types.Venue {
	return loc.response
}

func (c *contact) WritePhoneNumber(phone string) error {
	var err error
	if err = checkStringValue(phone, c.PhoneNumber); err == nil {
		c.PhoneNumber = phone
		logs.DataWrittenSuccessfully(interfaceContact, "Phone Number")
	}
	return err
}

func (c *contact) WriteFirstName(fname string) error {
	var err error
	if err = checkStringValue(fname, c.FirstName); err == nil {
		c.FirstName = fname
		logs.DataWrittenSuccessfully(interfaceContact, "First Name")
	}
	return err
}

func (c *contact) WriteLastName(lname string) error {
	var err error
	if err = checkStringValue(lname, c.LastName); err == nil {
		c.LastName = lname
		logs.DataWrittenSuccessfully(interfaceContact, "Last Name")
	}
	return err
}

func (c *contact) WriteVCard(vcard string) error {
	var err error
	if err = checkStringValue(vcard, c.Vcard); err == nil {
		c.Vcard = vcard
		logs.DataWrittenSuccessfully(interfaceContact, "vCard")
	}
	return err
}

func (c *contact) GetResponse() types.Contact {
	return c.response
}

func (p *poll) WriteQuestion(question string) error {
	var err error
	q := len([]rune(question))
	if (q >= 1) && (q <= 300) {
		if p.Question == "" {
			p.Question = question
			logs.DataWrittenSuccessfully(interfacePoll, "Question")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *poll) WriteQuestionParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, p.QuestionParsemode); err == nil {
		p.QuestionParsemode = parsemode
		logs.DataWrittenSuccessfully(interfacePoll, "Question Parse Mode")
	}
	return err
}

func (p *poll) WriteQuestionEntities(entities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(entities, p.QuestionEntities); err == nil {
		p.QuestionEntities = entities
		logs.DataWrittenSuccessfully(interfacePoll, "Question Entities")
	}
	return err
}

func (p *poll) WriteOptions(options []*types.PollOption) error {
	var err error
	if len(options) != 0 {

		for i := 0; (i < len(options)) && (err == nil); i++ {
			if options[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if len(p.Options) == 0 {
				p.Options = options
				logs.DataWrittenSuccessfully(interfacePoll, "Options")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (p *poll) WriteAnonymous() error {
	var err error
	if !p.IsAnonymous {
		p.IsAnonymous = true
		logs.SettedParam("Anonymous", interfacePoll, true)
	} else {
		err = code10()
	}
	return err
}

func (p *poll) WriteType(polltype string) error {
	var err error
	if (polltype == "quiz") || (polltype == "regular") {
		if p.Type == "" {
			p.Type = polltype
			logs.DataWrittenSuccessfully(interfacePoll, "Type")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *poll) WriteAllowMultipleAnswers() error {
	var err error
	if !p.AllowMultipleAnswers {
		p.AllowMultipleAnswers = true
		logs.SettedParam("Allow Multiple Answers", interfacePoll, true)
	} else {
		err = code10()
	}
	return err
}

func (p *poll) WriteCorrectOptionID(optID string) error {
	var err error
	if err = checkStringValue(optID, p.CorrectOptionID); err == nil {
		r := []rune(optID)
		if r[0] == '0' {
			p.CorrectOptionID = optID
			logs.DataWrittenSuccessfully(interfacePoll, "Correct Option ID")
		} else {
			err = code20()
		}
	}
	return err
}

func (p *poll) WriteExplanation(explanation string) error {
	var err error
	if err = checkStringValue(explanation, p.Explanation); err == nil {
		exp := []rune(explanation)
		if len(exp) >= 1 && len(exp) <= 200 {
			p.Explanation = explanation
			logs.DataWrittenSuccessfully(interfacePoll, "Explanation")
		} else {
			err = code20()
		}
	}
	return err
}

func (p *poll) WriteExplanationParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, p.ExplanationParsemode); err == nil {
		p.ExplanationParsemode = parsemode
		logs.DataWrittenSuccessfully(interfacePoll, "Explanation Parse Mode")
	}
	return err
}

func (p *poll) WriteExplanationEntities(entities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(entities, p.ExplanationEntities); err == nil {
		p.ExplanationEntities = entities
		logs.DataWrittenSuccessfully(interfacePoll, "Explanation Entities")
	}
	return err
}

func (p *poll) WriteOpenPeriod(period int) error {
	var err error
	if p.CloseDate == 0 {
		if (period >= 5) && (period <= 600) {
			if p.OpenPeriod == 0 {
				p.OpenPeriod = period
				logs.DataWrittenSuccessfully(interfacePoll, "Open Period")
			} else {
				err = code10()
			}
		} else {
			err = code20()
		}
	} else {
		err = code25()
	}
	return err
}

func (p *poll) WriteCloseDate(time int) error {
	var err error
	if p.OpenPeriod == 0 {
		if (time >= 5) && (time <= 600) {
			if p.CloseDate == 0 {
				p.CloseDate = time
				logs.DataWrittenSuccessfully(interfacePoll, "Close Date")
			} else {
				err = code10()
			}
		} else {
			err = code20()
		}
	} else {
		err = code25()
	}
	return err
}

func (p *poll) WriteClosed() error {
	var err error
	if !p.IsClosed {
		p.IsClosed = true
		logs.SettedParam("Closed", interfacePoll, true)
	} else {
		err = code10()
	}
	return err
}

func (p *poll) GetResponse() types.Poll {
	return p.response
}

func (l *link) WriteName(name string) error {
	var err error
	n := len([]rune(name))
	if n > 0 && n <= 32 {
		if l.Name == "" {
			l.Name = name
			logs.DataWrittenSuccessfully(interfaceLink, "Name")
		} else {
			err = code10()
		}
	} else {
		err = code20()

	}
	return err
}

func (l *link) WriteExpireDate(date time.Duration) error {
	var err error
	if date > 0 {
		if l.ExpireDate == 0 {
			l.ExpireDate = date
			logs.DataWrittenSuccessfully(interfaceLink, "Expire Date")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (l *link) WriteMemberLimit(limit int) error {
	var err error
	if limit > 0 && limit <= 99999 {
		if l.MemberLimit == 0 {
			l.MemberLimit = limit
			logs.DataWrittenSuccessfully(interfaceLink, "Member Limit")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (l *link) WriteJoinRequest() error {
	var err error
	if !l.JoinRequest {
		l.JoinRequest = true
		logs.SettedParam("Join Request", interfaceLink, true)
	} else {
		err = code10()
	}
	return err
}

func (l *link) WriteIniveLink(link string) error {
	var err error
	if err = checkStringValue(link, l.InviteLink); err == nil {
		l.InviteLink = link
		logs.DataWrittenSuccessfully(interfaceLink, "Invite Link")
	}
	return err
}

func (l *link) WriteSubscriptionPeriod(period int) error {
	var err error
	if period == Month {
		if l.SubscriptionPeriod == 0 {
			l.SubscriptionPeriod = period
			logs.DataWrittenSuccessfully(interfaceLink, "Subscription Period")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (l *link) WriteSubscriptionPrice(price int) error {
	var err error
	if price > 0 && price <= 2500 {
		if l.SubscriptionPrice == 0 {
			l.SubscriptionPrice = price
			logs.DataWrittenSuccessfully(interfaceLink, "Subscription Price")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteSetName(name string) error {
	var err error
	if err = checkStringValue(name, st.SetName); err == nil {
		st.SetName = name
		logs.DataWrittenSuccessfully(interfaceSticker, "Set Name")
	}
	return err
}

func (st *sticker) WriteStickerStorage(path string) error {
	var err error
	if err = checkStringValue(path, st.Sticker); err == nil {
		if err = st.isCorrectType(path); err == nil {
			st.Sticker, st.stickerGottenFrom = path, Storage
			logs.DataWrittenSuccessfully(interfaceSticker, "Sticker From Storage")
		}
	}
	return err
}

func (st *sticker) WriteStickerTelegram(stickerID string) error {
	var err error
	if err = checkStringValue(stickerID, st.Sticker); err == nil {
		st.Sticker, st.stickerGottenFrom = stickerID, Telegram
		logs.DataWrittenSuccessfully(interfaceSticker, "Sticker From Telegram")
	}
	return err
}

func (st *sticker) WriteStickerInternet(url string) error {
	var err error
	if err = checkStringValue(url, st.Sticker); err == nil {
		st.Sticker, st.stickerGottenFrom = url, Telegram
		logs.DataWrittenSuccessfully(interfaceSticker, "Sticker From The Internet")
	}
	return err
}

func (st *sticker) WriteAssociatedEmoji(emoji string) error {
	var err error
	if err = checkStringValue(emoji, st.Emoji); err == nil {
		st.Emoji = emoji
		logs.DataWrittenSuccessfully(interfaceSticker, "Associated Emoji")
	}
	return err
}

func (st *sticker) WriteAssociatedEmojies(emojies []string) error {
	var err error
	if emojies != nil {
		for i := 0; (i < len(emojies)) && (err == nil); i++ {
			if emojies[i] == "" {
				err = code5()
			}
		}
		if err == nil {
			if st.Emojies == nil {
				st.Emojies = emojies
				logs.DataWrittenSuccessfully(interfaceSticker, "Associated Emojies")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteEmojiID(emojiID string) error {
	var err error
	if err = checkStringValue(emojiID, st.EmojiID); err == nil {
		st.EmojiID = emojiID
		logs.DataWrittenSuccessfully(interfaceSticker, "Emoji ID")
	}
	return err
}

func (st *sticker) WriteEmojiIDs(emojiIDs []string) error {
	var err error
	if emojiIDs != nil {
		for i := 0; (i < len(emojiIDs)) && (err == nil); i++ {
			if emojiIDs[i] == "" {
				err = code5()
			}
		}
		if err == nil {
			if st.EmojiIDs == nil {
				st.EmojiIDs = emojiIDs
				logs.DataWrittenSuccessfully(interfaceSticker, "Emoji IDs")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteFormat(format string) error {
	var err error
	if format == "static" || format == "animated" || format == "video" {
		if st.Format == "" {
			st.Format = format
			logs.DataWrittenSuccessfully(interfaceSticker, "Format")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteTitle(title string) error {
	var err error
	t := len([]rune(title))
	if t > 0 && t <= 64 {
		if st.Title == "" {
			st.Title = title
			logs.DataWrittenSuccessfully(interfaceSticker, "Title")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteStickerType(stickertype string) error {
	var err error
	if stickertype == "regular" || stickertype == "mask" || stickertype == "custom_emoji" {
		if st.StickerType == "" {
			st.StickerType = stickertype
			logs.DataWrittenSuccessfully(interfaceSticker, "Sticker Type")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteNeedsRepainting() error {
	var err error
	if !st.NeedsRepainting {
		st.NeedsRepainting = true
		logs.SettedParam("Needs Repainting", interfaceSticker, true)
	} else {
		err = code10()
	}
	return err
}

func (st *sticker) WritePosition(pos string) error {
	var err error
	p := []rune(pos)
	if (len(p) > 0) && (p[0] == '0') {
		if st.Position == "" {
			st.Position = pos
			logs.DataWrittenSuccessfully(interfaceSticker, "Position")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteOldSticker(stickerID string) error {
	var err error
	if err = checkStringValue(stickerID, st.OldSticker); err == nil {
		st.OldSticker = stickerID
		logs.DataWrittenSuccessfully(interfaceSticker, "Old Sticker")
	}
	return err
}

func (st *sticker) WriteKeywords(words []string) error {
	var err error
	if len(words) <= 20 {
		for i := 0; (i < len(words)) && (err == nil); i++ {
			if len([]rune(words[i])) > 64 {
				err = code5()
			}
		}
		if err == nil {
			if st.Keywords == nil {
				st.Keywords = words
				logs.DataWrittenSuccessfully(interfaceSticker, "Keywords")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteMaskPosition(maskpos *types.MaskPosition) error {
	var err error
	if maskpos != nil {
		if st.MaskPosition == nil {
			st.MaskPosition = maskpos
			logs.DataWrittenSuccessfully(interfaceSticker, "Mask Position")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteThumbnailStorage(path string) error {
	var err error
	if err = checkStringValue(path, st.Thumbnail); err == nil {
		if err = st.isThumbnailCorrectType(path); err == nil {
			st.Thumbnail, st.thumbnailGottenFrom = path, Storage
			logs.DataWrittenSuccessfully(interfaceSticker, "Thumbnail From Storage")
		}
	}
	return err
}

func (st *sticker) WriteThumbnailTelegram(thumbnailID string) error {
	var err error
	if err = checkStringValue(thumbnailID, st.Thumbnail); err == nil {
		st.Thumbnail, st.thumbnailGottenFrom = thumbnailID, Telegram
		logs.DataWrittenSuccessfully(interfaceSticker, "Thumbnail From Telegram")
	}
	return err
}

func (st *sticker) WriteThumbnailInternet(url string) error {
	var err error
	if err = checkStringValue(url, st.Thumbnail); err == nil {
		st.Thumbnail, st.thumbnailGottenFrom = url, Internet
		logs.DataWrittenSuccessfully(interfaceSticker, "Thumbnail From The Internet")
	}
	return err
}

func (st *sticker) WriteThumbnailFormat(format string) error {
	var err error
	if (format == "static" && (st.thumbnailType == ".WEBP/.PNG" || st.thumbnailType == "")) ||
		(format == "animated" && (st.thumbnailType == ".TGS" || st.thumbnailType == "")) ||
		(format == "video" && (st.thumbnailType == ".WEBM" || st.thumbnailType == "")) {
		if st.ThumbnailFormat == "" {
			st.ThumbnailFormat = format
			logs.DataWrittenSuccessfully(interfaceSticker, "Thumbnail Format")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (st *sticker) WriteGiftID(giftID string) error {
	var err error
	if err = checkStringValue(giftID, st.GiftID); err == nil {
		st.GiftID = giftID
		logs.DataWrittenSuccessfully(interfaceSticker, "Gift ID")
	}
	return err
}

func (st *sticker) WriteTextParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, st.ParseMode); err == nil {
		st.ParseMode = parsemode
		logs.DataWrittenSuccessfully(interfaceSticker, "Text Parse Mode")
	}
	return err
}

func (st *sticker) WriteTextEntities(entities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(entities, st.Entities); err == nil {
		st.Entities = entities
		logs.DataWrittenSuccessfully(interfaceSticker, "Text Entities")
	}
	return err
}

func (f *forum) WriteName(name string) error {
	var err error
	n := len([]rune(name))
	if n > 0 && n <= 128 {
		if f.Name == "" {
			f.Name = name
			logs.DataWrittenSuccessfully(interfaceForum, "Name")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (f *forum) WriteIconColor(color int) error {
	var err error
	if color == 0x6FB9F0 || color == 0xFFD67E || color == 0xCB86DB || color == 0x8EEE98 || color == 0xFF93B2 || color == 0xFB6F5F {
		if f.IconColor == 0 {
			f.IconColor = color
			logs.DataWrittenSuccessfully(interfaceForum, "Icon Color")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (f *forum) WriteIconEmojiID(emojiID string) error {
	var err error
	if err = checkStringValue(emojiID, f.IconEmojiID); err == nil {
		f.IconEmojiID = emojiID
		logs.DataWrittenSuccessfully(interfaceForum, "Icon Emoji ID")
	}
	return err
}

func (b *bot) WriteCommands(commands []*types.BotCommand) error {
	var err error
	if commands != nil {
		for i := 0; (i < len(commands)) && (err == nil); i++ {
			if commands[i] == nil {
				err = code5()
			}
		}
		if err == nil {
			if b.Commands == nil {
				b.Commands = commands
				logs.DataWrittenSuccessfully(interfaceBot, "Commands")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteScope(scope *types.BotCommandScope) error {
	var err error
	if scope != nil {
		if b.Scope == nil {
			b.Scope = scope
			logs.DataWrittenSuccessfully(interfaceBot, "Scope")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteLanguage(lang string) error {
	var err error
	l := len([]rune(lang))
	if l == 0 || l == 2 {
		if b.Language == "" {
			b.Language = lang
			logs.DataWrittenSuccessfully(interfaceBot, "Language")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteName(name string) error {
	var err error
	n := len([]rune(name))
	if n >= 0 && n <= 64 {
		if b.Name == "" {
			b.Name = name
			logs.DataWrittenSuccessfully(interfaceBot, "Name")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteDescription(description string) error {
	var err error
	d := len([]rune(description))
	if d >= 0 && d <= 512 {
		if b.Description == "" {
			b.Description = description
			logs.DataWrittenSuccessfully(interfaceBot, "Description")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteShortDescription(description string) error {
	var err error
	d := len([]rune(description))
	if d >= 0 && d <= 120 {
		if b.ShortDescription == "" {
			b.ShortDescription = description
			logs.DataWrittenSuccessfully(interfaceBot, "Description")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteMenuButton(button *types.MenuButton) error {
	var err error
	if button != nil {
		if b.MenuButtom == nil {
			b.MenuButtom = button
			logs.DataWrittenSuccessfully(interfaceBot, "Menu Buttom")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteRights(rights *types.ChatAdministratorRights) error {
	var err error
	if rights != nil {
		if b.Rights == nil {
			b.Rights = rights
			logs.DataWrittenSuccessfully(interfaceBot, "Rights")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (b *bot) WriteForChannels() error {
	var err error
	if !b.ForChannels {
		b.ForChannels = true
		logs.SettedParam("For Channels", interfaceBot, true)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteString(text string) error {
	var err error
	if err = checkStringValue(text, inf.Text); err == nil {
		inf.Text = text
		logs.DataWrittenSuccessfully(interfaceInf, "Text")
	}
	return err
}

func (inf *information) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, inf.ParseMode); err == nil {
		inf.ParseMode = parsemode
		logs.DataWrittenSuccessfully(interfaceAudio, "Parse Mode")
	}
	return err
}

func (inf *information) WriteMessageThreadID(messageID int) error {
	var err error
	if err = checkIntegerValue(messageID, inf.MessageThreadID); err == nil {
		inf.MessageThreadID = messageID
		logs.DataWrittenSuccessfully(interfaceInf, "Message Thread ID")
	}
	return err
}

func (inf *information) WriteDisableNotification() error {
	var err error
	if !inf.DisableNotification {
		inf.DisableNotification = true
		logs.SettedParam("Disable Notification", interfaceInf, inf.DisableNotification)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteProtectContent() error {
	var err error
	if !inf.ProtectContent {
		inf.ProtectContent = true
		logs.SettedParam("Protect Content", interfaceInf, inf.ProtectContent)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteMessageEffectID(messageID string) error {
	var err error
	if err = checkStringValue(messageID, inf.MessageEffectID); err == nil {
		inf.MessageEffectID = messageID
		logs.DataWrittenSuccessfully(interfaceInf, "Message Effect ID")
	}
	return err
}

func (inf *information) WriteEntities(entities []*types.MessageEntity) error {
	var err error
	if err = checkEntities(entities, inf.Entities); err == nil {
		inf.Entities = entities
		logs.DataWrittenSuccessfully(interfaceInf, "Entities")
	}
	return err
}

func (inf *information) WriteLinkPreviewOptions(lpo *types.LinkPreviewOptions) error {
	var err error
	if lpo != nil {
		if inf.LinkPreviewOptions == nil {
			inf.LinkPreviewOptions = lpo
			logs.DataWrittenSuccessfully(interfaceInf, "Link Preview Options")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteMessageID(messageID int) error {
	var err error
	if err = checkIntegerValue(messageID, inf.MessageID); err == nil {
		inf.MessageID = messageID
		logs.DataWrittenSuccessfully(interfaceInf, "Message ID")
	}
	return err
}

func (inf *information) WriteMessageIDs(messageIDs []int) error {
	var err error
	if len(messageIDs) > 0 {

		for i := 0; (i < len(messageIDs)) && (err == nil); i++ {
			if messageIDs[i] == 0 {
				err = code5()
			}
		}

		if err == nil {
			if len(inf.MessageIDs) == 0 {
				inf.MessageIDs = messageIDs
				logs.DataWrittenSuccessfully(interfaceInf, "Message IDs")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, inf.Caption); err == nil {
		inf.Caption = caption
		logs.DataWrittenSuccessfully(interfaceInf, "Caption")
	}
	return err
}

func (inf *information) WriteShowCaptionAboveMedia() error {
	var err error
	if !inf.ShowCaptionAboveMedia {
		inf.ShowCaptionAboveMedia = true
		logs.DataWrittenSuccessfully(interfaceInf, "Show Caption Above Media")
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteReplyParameters(reply *types.ReplyParameters) error {
	var err error
	if reply != nil {

		if len(reply.QuoteEntities) > 0 {
			for i := 0; (i < len(reply.QuoteEntities)) && (err == nil); i++ {
				if reply.QuoteEntities[i] == nil {
					err = code5()
				}
			}
		}

		if err == nil {
			if inf.ReplyParameters == nil {
				inf.ReplyParameters = reply
				logs.DataWrittenSuccessfully(interfaceInf, "Reply Parameters")
			} else {
				err = code10()
			}
		}

	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteAllowPaidBroadcast() error {
	var err error
	if !inf.AllowPaidBroadcast {
		inf.AllowPaidBroadcast = true
		logs.DataWrittenSuccessfully(interfaceInf, "Allow Paid Broadcast")
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteStarCount(amount int) error {
	var err error
	if err = checkIntegerValue(amount, inf.StarCount); err == nil {
		inf.StarCount = amount
		logs.DataWrittenSuccessfully(interfaceInf, "Star Count")
	}
	return err
}

func (inf *information) WritePayload(payload string) error {
	var err error
	if err = checkStringValue(payload, inf.Payload); err == nil {
		inf.Payload = payload
		logs.DataWrittenSuccessfully(interfaceInf, "Payload")
	}
	return err
}

func makeMap() map[rune]*emojich {
	emojiMap := make(map[rune]*emojich, len(types.Emojis))
	emojiMap[rune('🎲')] = &emojich{permissibleQuantity: 6}
	emojiMap[rune('🎯')] = &emojich{permissibleQuantity: 6}
	emojiMap[rune('🏀')] = &emojich{permissibleQuantity: 5}
	emojiMap[rune('⚽')] = &emojich{permissibleQuantity: 5}
	emojiMap[rune('🎳')] = &emojich{permissibleQuantity: 6}
	emojiMap[rune('🎰')] = &emojich{permissibleQuantity: 64}
	return emojiMap
}

func isArrEmojiOK(arr []rune, emojimap map[rune]*emojich) error {
	var err error
	for i := 1; (i < len(arr)) && (err == nil); i++ {
		if arr[i] != arr[0] {
			err = code20()
		} else {
			emojimap[arr[0]].amount++
		}
	}
	return err
}

func (inf *information) WriteEmoji(emoji string) error {
	var err error
	emojiMap := makeMap()
	em := []rune(emoji)
	if len(em) > 0 {
		if _, ok := emojiMap[em[0]]; !ok {
			err = code20()
		} else {
			if err = isArrEmojiOK(em, emojiMap); err == nil {
				if emojiMap[em[0]].amount < emojiMap[em[0]].permissibleQuantity {
					if inf.Emoji == "" {
						inf.Emoji = emoji
						logs.DataWrittenSuccessfully(interfaceInf, "Emoji")
					} else {
						err = code10()
					}
				} else {
					err = code20()

				}
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteAction(action string) error {
	var err error
	var found bool
	for i := 0; (i < len(types.Actions)) && (!found); i++ {
		if types.Actions[i] == action {
			found = true
		} else if i+1 == len(types.Actions) {
			err = code20()
		}
	}
	if err == nil {
		if inf.Action == "" {
			inf.Action = action
			logs.DataWrittenSuccessfully(interfaceInf, "Action")
		} else {
			err = code10()
		}
	}
	return err
}

func (inf *information) WriteReaction(reaction []*types.ReactionType) error {
	var err error
	if reaction != nil {
		for i := 0; (i < len(reaction)) && (err == nil); i++ {
			if reaction[i] == nil {
				err = code5()
			}
		}
		if err == nil {
			if inf.Reaction == nil {
				inf.Reaction = reaction
				logs.DataWrittenSuccessfully(interfaceInf, "Reaction")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteReactionIsBig() error {
	var err error
	if !inf.ReactionIsBig {
		inf.ReactionIsBig = true
		logs.SettedParam("Reaction Is Big", interfaceInf, true)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteOffset(offset int) error {
	var err error
	if err = checkIntegerValue(offset, inf.Offset); err == nil {
		inf.Offset = offset
		logs.DataWrittenSuccessfully(interfaceInf, "Offset")
	}
	return err
}

func (inf *information) WriteLimit(limit int) error {
	var err error
	if limit > 0 && limit <= 100 {
		if inf.Limit == 0 {
			inf.Limit = limit
			logs.DataWrittenSuccessfully(interfaceInf, "Limit")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteEmojiStatusCustomEmojiID(emojiID string) error {
	var err error
	if err = checkStringValue(emojiID, inf.EmojiStatusCustomEmojiID); err == nil {
		inf.EmojiStatusCustomEmojiID = emojiID
		logs.DataWrittenSuccessfully(interfaceInf, "Emoji Status Custom Emoji ID")
	}
	return err
}

func (inf *information) WriteEmojiStatusExpirationDate(date int) error {
	var err error
	if err = checkIntegerValue(date, inf.EmojiStatusExpirationDate); err == nil {
		inf.EmojiStatusExpirationDate = date
		logs.DataWrittenSuccessfully(interfaceInf, "Emoji Status Expiration Date")
	}
	return err
}

func (inf *information) WriteFileID(fileID string) error {
	var err error
	if err = checkStringValue(fileID, inf.FileID); err == nil {
		inf.FileID = fileID
		logs.DataWrittenSuccessfully(interfaceInf, "File ID")
	}
	return err
}

func (inf *information) WriteUntilDate(date time.Duration) error {
	var err error
	if date > 0 {
		if inf.UntilDate == 0 {
			inf.UntilDate = date
			logs.DataWrittenSuccessfully(interfaceInf, "Until Date")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteRevokeMessages() error {
	var err error
	if !inf.RevokeMessages {
		inf.RevokeMessages = true
		logs.SettedParam("Revoke Messages", interfaceInf, true)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteOnlyIfBanned() error {
	var err error
	if !inf.OnlyIfBanned {
		inf.OnlyIfBanned = true
		logs.SettedParam("Only If Banned", interfaceInf, true)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WritePermissions(permissions *types.ChatPermissions) error {
	var err error
	if permissions != nil {
		if inf.Permissions == nil {
			inf.Permissions = permissions
			logs.DataWrittenSuccessfully(interfaceInf, "Permissions")
		} else {
			err = code10()
		}
	} else {

		err = code20()
	}
	return err
}

func (inf *information) WriteIndependentChatPermissions() error {
	var err error
	if !inf.IndependentChatPermissions {
		inf.IndependentChatPermissions = true
		logs.SettedParam("Independent Chat Permissions", interfaceInf, true)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteAdministratorRights(chAdminRights *types.ChatAdministratorRights) error {
	var err error
	if chAdminRights != nil {
		if inf.AdminRights == nil {
			inf.AdminRights = chAdminRights
			logs.DataWrittenSuccessfully(interfaceInf, "Admin Rights")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteCustomTitle(title string) error {
	var err error
	t := len([]rune(title))
	if t > 0 && t <= 16 {
		if inf.CustomTitle == "" {
			inf.CustomTitle = title
			logs.DataWrittenSuccessfully(interfaceInf, "Custom Title")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteUserID(userID int) error {
	var err error
	if err = checkIntegerValue(userID, inf.UserID); err == nil {
		inf.UserID = userID
		logs.DataWrittenSuccessfully(interfaceInf, "User ID")
	}
	return err
}

func (inf *information) WriteCallBackQueryID(queryID string) error {
	var err error
	if err = checkStringValue(queryID, inf.CallBackQueryID); err == nil {
		inf.CallBackQueryID = queryID
		logs.DataWrittenSuccessfully(interfaceInf, "CallBack Query ID")
	}
	return err
}

func (inf *information) WriteShowAlert() error {
	var err error
	if !inf.ShowAlert {
		inf.ShowAlert = true
		logs.SettedParam("Show Alert", interfaceInf, true)
	} else {
		err = code10()
	}
	return err
}

func (inf *information) WriteURL(url string) error {
	var err error
	if err = checkStringValue(url, inf.Url); err == nil {
		inf.Url = url
		logs.DataWrittenSuccessfully(interfaceInf, "URL")
	}
	return err
}

func (inf *information) WriteCacheTime(time int) error {
	var err error
	if err = checkIntegerValue(time, inf.CacheTime); err == nil {
		inf.CacheTime = time
		logs.DataWrittenSuccessfully(interfaceInf, "Cache Time")
	}
	return err
}

func (inf *information) WriteInlineMessageID(messageID string) error {
	var err error
	if err = checkStringValue(messageID, inf.InlineMessageID); err == nil {
		inf.InlineMessageID = messageID
		logs.DataWrittenSuccessfully(interfaceInf, "Inline Message ID")
	}
	return err
}

func (inf *information) WriteErrors(errors []*types.PassportElementError) error {
	var err error
	if len(errors) > 0 {
		for i := 0; i < len(errors) && err == nil; i++ {
			if errors[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if inf.Errors == nil {
				inf.Errors = errors
				logs.DataWrittenSuccessfully(interfaceInf, "Errors")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) GetResponse() types.User {
	return inf.response
}

func (inf *information) GetMessageIDs() []int {
	return inf.responseMessageIDs
}

func (ch *chat) WriteChatID(chatID int) error {
	var err error
	if chatID > 0 {
		if ch.ID == nil {
			ch.ID = chatID
			logs.DataWrittenSuccessfully(interfaceChat, "Chat ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ch *chat) WriteChatName(chatname string) error {
	var err error
	if chatname != "" {
		if ch.ID == nil {
			ch.ID = fmt.Sprint("@", chatname)
			logs.DataWrittenSuccessfully(interfaceChat, "Chat Name")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ch *chat) WriteBusinessConnectionID(connectionID string) error {
	var err error
	if err = checkStringValue(connectionID, ch.BusinessConnectionID); err == nil {
		ch.BusinessConnectionID = connectionID
		logs.DataWrittenSuccessfully(interfaceChat, "Business Connection ID")
	}
	return err
}

func (ch *chat) WriteFromChatID(chatID int) error {
	var err error
	if chatID > 0 {
		if ch.FromChatID == nil {
			ch.FromChatID = chatID
			logs.DataWrittenSuccessfully(interfaceChat, "From Chat ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ch *chat) WriteFromChatName(chatname string) error {
	var err error
	if chatname != "" {
		if ch.FromChatID == nil {
			ch.FromChatID = fmt.Sprint("@", chatname)
			logs.DataWrittenSuccessfully(interfaceChat, "From Chat ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ch *chat) WriteSenderChatID(chatID int) error {
	var err error
	if err = checkIntegerValue(chatID, ch.SenderChatID); err == nil {
		ch.SenderChatID = chatID
		logs.DataWrittenSuccessfully(interfaceChat, "Sender Chat ID")
	}
	return err
}

func (ch *chat) WriteTitle(title string) error {
	var err error
	t := len([]rune(title))
	if t > 0 && t <= 128 {
		if ch.Title == "" {
			ch.Title = title
			logs.DataWrittenSuccessfully(interfaceChat, "Title")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ch *chat) WriteDescription(description string) error {
	var err error
	d := len([]rune(description))
	if d >= 0 && d <= 255 {
		if ch.Description == "" {
			ch.Description = description
			logs.DataWrittenSuccessfully(interfaceChat, "Description")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func isNill(data interface{}) error {
	var err error
	if data != nil {
		err = code20()
	}
	return err
}

func (kb *keyboard) WriteReply() (IReply, error) {
	var err error
	rp := new(reply)
	if err = isNill(kb.Keyboard); err == nil {
		kb.Keyboard = rp
	}
	return rp, err
}

func (kb *keyboard) WriteInline() (IInline, error) {
	var err error
	in := new(inline)
	if err = isNill(kb.Keyboard); err == nil {
		kb.Keyboard = in
	}
	return in, err
}

func (kb *keyboard) WriteForceReply() (IForceReply, error) {
	var err error
	frp := new(forcereply)
	if err = isNill(kb.Keyboard); err == nil {
		kb.Keyboard = frp
	}
	return frp, err
}

func (in *inline) Set(plan []int) error {
	var err error
	if len(plan) > 0 {
		for i := 0; (i < len(plan)) && (err == nil); i++ {
			if plan[i] == 0 {
				err = code20()
			}
		}
		if err == nil {
			if in.Keyboard == nil {
				in.Keyboard = make([][]*inlineKeyboardButton, len(plan))
				for i := range in.Keyboard {
					in.Keyboard[i] = make([]*inlineKeyboardButton, plan[i])
				}
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (ch *chat) GetResponse() types.Chat {
	return ch.response
}

func (in *inline) NewButton(line, pos int) (IInlineButton, error) {
	var err error
	but := new(inlineKeyboardButton)
	if (line >= 0) && (pos >= 0) {
		if (in.Keyboard != nil) && ((len(in.Keyboard) > line) && (len(in.Keyboard[line]) > pos)) {
			if in.Keyboard[line][pos] == nil {
				in.Keyboard[line][pos] = new(inlineKeyboardButton)
				but = in.Keyboard[line][pos]
			} else {
				err = code10()
			}
		} else {
			err = code01()
		}
	} else {
		err = code20()
	}
	return but, err
}

func (inb *inlineKeyboardButton) WriteString(text string) error {
	var err error
	if err = checkStringValue(text, inb.Text); err == nil {
		inb.Text = text
		logs.DataWrittenSuccessfully(inbtn, "Text")
	}
	return err
}

func (inb *inlineKeyboardButton) checkOthers() error {
	var err error
	for i := 0; (i < len(inb.storage)) && (err == nil); i++ {
		if inb.storage[i] == added {
			err = code25()
		}
	}
	return err
}

func (inb *inlineKeyboardButton) WriteURL(url string) error {
	var err error
	if err = checkStringValue(url, inb.Url); err == nil {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wURL] = added
			inb.Url = url
			logs.DataWrittenSuccessfully(inbtn, "URL")
		}
	}
	return err
}

func (inb *inlineKeyboardButton) WriteCallbackData(text string) error {
	var err error
	if err = checkStringValue(text, inb.CallbackData); err == nil {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wCallback] = added
			inb.CallbackData = text
			logs.DataWrittenSuccessfully(inbtn, "Callback Data")
		}
	}
	return err
}

func (inb *inlineKeyboardButton) WriteWebApp(wbapp *types.WebAppInfo) error {
	var err error
	if (wbapp != nil) && (wbapp.Url != "") {
		if inb.WebApp == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wWebApp] = added
				inb.WebApp = wbapp
				logs.DataWrittenSuccessfully(inbtn, "Web App")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteLoginUrl(logurl *types.LoginUrl) error {
	var err error
	if logurl != nil {
		if inb.LoginUrl == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wLoginUrl] = added
				inb.LoginUrl = logurl
				logs.DataWrittenSuccessfully(inbtn, "Login URL")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQuery(sw string) error {
	var err error
	if err = checkStringValue(sw, inb.SwitchInlineQuery); err == nil {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wSwitchIn] = added
			inb.SwitchInlineQuery = sw
			logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query")
		}
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryCurrentChat(swcch string) error {
	var err error
	if err = checkStringValue(swcch, inb.SwitchInlineQueryCurrentChat); err == nil {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wSwitchInQuery] = added
			inb.SwitchInlineQueryCurrentChat = swcch
			logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Current Chat")
		}
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat) error {
	var err error
	if sw != nil {
		if inb.SwitchInlineQueryChosenChat == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wSwitchInQueryCh] = added
				inb.SwitchInlineQueryChosenChat = sw
				logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Chosen Chat")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteCallbackGame(game *types.CallbackGame) error {
	var err error
	if game != nil {
		if inb.CallbackGame == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wGame] = added
				inb.CallbackGame = game
				logs.DataWrittenSuccessfully(inbtn, "Callback Game")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WritePay() error {
	var err error
	if !inb.Pay {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wPay] = added
			inb.Pay = true
			logs.SettedParam("Pay", inbtn, true)
		}
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) Set(plan []int) error {
	var err error
	if len(plan) > 0 {
		for i := 0; (i < len(plan)) && (err == nil); i++ {
			if plan[i] == 0 {
				err = code20()
			}
		}
		if err == nil {
			if rp.Keyboard == nil {
				rp.Keyboard = make([][]*replyKeyboardButton, len(plan))
				for i := range rp.Keyboard {
					rp.Keyboard[i] = make([]*replyKeyboardButton, plan[i])
				}
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (rp *reply) WriteIsPersistent() error {
	var err error
	if !rp.IsPersistent {
		rp.IsPersistent = true
		logs.SettedParam("Is Persistent", interfaceReplyKB, true)
	} else {
		err = code10()
	}

	return err
}

func (rp *reply) WriteResizeKeyboard() error {
	var err error
	if !rp.ResizeKeyboard {
		rp.ResizeKeyboard = true
		logs.SettedParam("Resize Keyboard", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteOneTimeKeyboard() error {
	var err error
	if !rp.OneTimeKeyboard {
		rp.OneTimeKeyboard = true
		logs.SettedParam("One Time Keyboard", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteInputFieldPlaceholder(placeholder string) error {
	var err error
	if err = checkInputPlaceHolder(placeholder, rp.InputFieldPlaceholder); err == nil {
		rp.InputFieldPlaceholder = placeholder
		logs.DataWrittenSuccessfully(interfaceReplyKB, "Input Field Placeholder")
	}
	return err
}

func (rp *reply) WriteSelective() error {
	var err error
	if !rp.Selective {
		rp.Selective = true
		logs.SettedParam("Selective", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteRemove() error {
	var err error
	if !rp.Remove {
		rp.Remove = true
		logs.SettedParam("Remove", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) NewButton(line, pos int) (IReplyButton, error) {
	var err error
	but := new(replyKeyboardButton)
	if (line >= 0) && (pos >= 0) {
		if (rp.Keyboard != nil) && ((len(rp.Keyboard) > line) && (len(rp.Keyboard[line]) > pos)) {
			if rp.Keyboard[line][pos] == nil {
				rp.Keyboard[line][pos] = new(replyKeyboardButton)
				but = rp.Keyboard[line][pos]
			} else {
				err = code10()
			}
		} else {
			err = code01()
		}
	} else {
		err = code20()
	}
	return but, err
}

func (rpb *replyKeyboardButton) checkOthers() error {
	var err error
	for i := 0; (i < len(rpb.storage)) && (err == nil); i++ {
		if rpb.storage[i] == added {
			err = code25()
		}
	}
	return err
}

func (rpb *replyKeyboardButton) WriteString(text string) error {
	var err error
	if err = checkStringValue(text, rpb.Text); err == nil {
		rpb.Text = text
		logs.DataWrittenSuccessfully(rpbtn, "Text")
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestUsers(requs *types.KeyboardButtonRequestUsers) error {
	var err error
	if requs != nil {
		if rpb.RequestUsers == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqUsers] = added
				rpb.RequestUsers = requs
				logs.DataWrittenSuccessfully(rpbtn, "Request Users")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestChat(reqch *types.KeyboardButtonRequestChat) error {
	var err error
	if reqch != nil {
		if rpb.RequestChat == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqChat] = added
				rpb.RequestChat = reqch
				logs.DataWrittenSuccessfully(rpbtn, "Request Chat")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestContact() error {
	var err error
	if !rpb.RequestContact {
		if err = rpb.checkOthers(); err == nil {
			rpb.storage[wReqContact] = added
			rpb.RequestContact = true
			logs.SettedParam("Request Contact", rpbtn, rpb.RequestContact)
		}
	} else {
		err = code10()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestLocation() error {
	var err error
	if !rpb.RequestLocation {
		if err = rpb.checkOthers(); err == nil {
			rpb.storage[wReqLocation] = added
			rpb.RequestLocation = true
			logs.SettedParam("Request Location", rpbtn, rpb.RequestLocation)
		}
	} else {
		err = code10()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestPoll(poll *types.KeyboardButtonPollType) error {
	var err error
	if poll != nil {
		if rpb.RequestPoll == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqPoll] = added
				rpb.RequestPoll = poll
				logs.DataWrittenSuccessfully(rpbtn, "Request Poll")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteWebApp(webapp *types.WebAppInfo) error {
	var err error
	if (webapp != nil) && (webapp.Url != "") {
		if rpb.WebApp == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqWebApp] = added
				rpb.WebApp = webapp
				logs.DataWrittenSuccessfully(rpbtn, "Web App")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (frp *forcereply) WriteForceReply() error {
	var err error
	if !frp.Forcereply {
		frp.Forcereply = true
		logs.SettedParam("Force Reply", interfaceForceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (frp *forcereply) WriteInputFieldPlaceholder(placeholder string) error {
	var err error
	if err = checkInputPlaceHolder(placeholder, frp.InputFieldPlaceholder); err == nil {
		frp.InputFieldPlaceholder = placeholder
		logs.DataWrittenSuccessfully(interfaceForceReplyKB, "Input Field Placeholder")
	}
	return err
}

func (frp *forcereply) WriteSelective() error {
	var err error
	if !frp.Selective {
		frp.Selective = true
		logs.SettedParam("Selective", interfaceForceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (in *inlinemode) WriteQueryID(queryID string) error {
	var err error
	if err = checkStringValue(queryID, in.QueryID); err == nil {
		in.QueryID = queryID
		logs.DataWrittenSuccessfully(interfaceInlineMode, "Query ID")
	}
	return err
}

func (in *inlinemode) WriteWebAppQueryID(queryID string) error {
	var err error
	if err = checkStringValue(queryID, in.WebAppQueryID); err == nil {
		in.WebAppQueryID = queryID
		logs.DataWrittenSuccessfully(interfaceInlineMode, "Web App Query ID")
	}
	return err
}

func (in *inlinemode) WriteResult() (IResult, error) {
	var err error
	var r IResult
	if in.Results == nil && in.Result == nil {
		r = &result{}
		in.Result = r
	} else {
		err = code10()
	}
	return r, err
}

func (in *inlinemode) WriteResults(length int) ([]IResult, error) {
	var err error
	var r []IResult
	if length > 0 {
		if in.Results == nil && in.Result == nil {
			r = make([]IResult, length)
			for i := range r {
				r[i] = &result{}
			}
			in.Results = r
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return r, err
}

func (in *inlinemode) WriteCacheTime(time int) error {
	var err error
	if err = checkIntegerValue(time, in.CacheTime); err == nil {
		in.CacheTime = time
		logs.DataWrittenSuccessfully(interfaceInlineMode, "Cache Time")
	}
	return err
}

func (in *inlinemode) WriteIsPersonal() error {
	var err error
	if !in.IsPersonal {
		in.IsPersonal = true
		logs.SettedParam("Is Personal", interfaceInlineMode, true)
	} else {
		err = code10()
	}
	return err
}

func (in *inlinemode) WriteNextOffset(offset string) error {
	var err error
	if len(offset) <= 64 {
		if in.NextOffset == "" {
			in.NextOffset = offset
			logs.DataWrittenSuccessfully(interfaceInlineMode, "Next Offset")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (in *inlinemode) WriteButton(but *types.InlineQueryResultsButton) error {
	var err error
	if but != nil {
		if but.Text != "" {
			if (but.StartParameter != "") || (but.WebApp != nil) {
				if in.Button == nil {
					in.Button = but
					logs.DataWrittenSuccessfully(interfaceInlineMode, "Button")
				} else {
					err = code10()
				}
			} else {
				err = code20()
			}
		} else {
			err = code20()
		}
	} else {
		err = code20()
	}
	return err
}

func (in *inlinemode) WriteAllowUserChats() error {
	var err error
	if !in.AllowUserChats {
		in.AllowUserChats = true
		logs.SettedParam("Allow User Chats", interfaceInlineMode, true)
	} else {
		err = code10()
	}
	return err
}

func (in *inlinemode) WriteAllowBotChats() error {
	var err error
	if !in.AllowBotChats {
		in.AllowBotChats = true
		logs.SettedParam("Allow Bot Chats", interfaceInlineMode, true)
	} else {
		err = code10()
	}
	return err
}

func (in *inlinemode) WriteAllowGroupChats() error {
	var err error
	if !in.AllowGroupChats {
		in.AllowGroupChats = true
		logs.SettedParam("Allow Group Chats", interfaceInlineMode, true)
	} else {
		err = code10()
	}
	return err
}

func (in *inlinemode) WriteAllowChannelChats() error {
	var err error
	if !in.AllowChannelChats {
		in.AllowChannelChats = true
		logs.SettedParam("Allow Channel Chats", interfaceInlineMode, true)
	} else {
		err = code10()
	}
	return err
}

func (in *inlinemode) GetSentWebAppMessage() types.SentWebAppMessage {
	return *in.webAppResponse
}

func (in *inlinemode) GetPreparedInlineMessage() types.PreparedInlineMessage {
	return *in.inMessageResponse
}

func (r *result) WriteCachedAudio(cachedAudio *types.InlineQueryResultCachedAudio) error {
	var err error
	if cachedAudio != nil {
		if r.CachedAudio == nil {
			r.CachedAudio = cachedAudio
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Audio")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedDocument(cachedDocument *types.InlineQueryResultCachedDocument) error {
	var err error
	if cachedDocument != nil {
		if r.CachedDocument == nil {
			r.CachedDocument = cachedDocument
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Document")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedGif(cachedGif *types.InlineQueryResultCachedGif) error {
	var err error
	if cachedGif != nil {
		if r.CachedGif == nil {
			r.CachedGif = cachedGif
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Gif")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedMpeg4Gif(cachedMpeg4Gif *types.InlineQueryResultCachedMpeg4Gif) error {
	var err error
	if cachedMpeg4Gif != nil {
		if r.CachedMpeg4Gif == nil {
			r.CachedMpeg4Gif = cachedMpeg4Gif
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Mpeg4 Gif")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedPhoto(cachedPhoto *types.InlineQueryResultCachedPhoto) error {
	var err error
	if cachedPhoto != nil {
		if r.CachedPhoto == nil {
			r.CachedPhoto = cachedPhoto
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Photo")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedSticker(cachedSticker *types.InlineQueryResultCachedSticker) error {
	var err error
	if cachedSticker != nil {
		if r.CachedSticker == nil {
			r.CachedSticker = cachedSticker
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Sticker")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedVideo(cachedVideo *types.InlineQueryResultCachedVideo) error {
	var err error
	if cachedVideo != nil {
		if r.CachedVideo == nil {
			r.CachedVideo = cachedVideo
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Video")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteCachedVoice(cachedVoice *types.InlineQueryResultCachedVoice) error {
	var err error
	if cachedVoice != nil {
		if r.CachedVoice == nil {
			r.CachedVoice = cachedVoice
			logs.DataWrittenSuccessfully(interfaceResult, "Cached Voice")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteArticle(art *types.InlineQueryResultArticle) error {
	var err error
	if art != nil {
		if r.Article == nil {
			r.Article = art
			logs.DataWrittenSuccessfully(interfaceResult, "Article")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteAudio(ad *types.InlineQueryResultAudio) error {
	var err error
	if ad != nil {
		if r.Audio == nil {
			r.Audio = ad
			logs.DataWrittenSuccessfully(interfaceResult, "Audio")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteContact(cont *types.InlineQueryResultContact) error {
	var err error
	if cont != nil {
		if r.Contact == nil {
			r.Contact = cont
			logs.DataWrittenSuccessfully(interfaceResult, "Contact")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteGame(game *types.InlineQueryResultGame) error {
	var err error
	if game != nil {
		if r.Game == nil {
			r.Game = game
			logs.DataWrittenSuccessfully(interfaceResult, "Game")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteDocument(doc *types.InlineQueryResultDocument) error {
	var err error
	if doc != nil {
		if r.Document == nil {
			r.Document = doc
			logs.DataWrittenSuccessfully(interfaceResult, "Document")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteGif(gif *types.InlineQueryResultGif) error {
	var err error
	if gif != nil {
		if r.Gif == nil {
			r.Gif = gif
			logs.DataWrittenSuccessfully(interfaceResult, "Gif")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteLocation(loc *types.InlineQueryResultLocation) error {
	var err error
	if loc != nil {
		if r.Location == nil {
			r.Location = loc
			logs.DataWrittenSuccessfully(interfaceResult, "Location")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteMpeg4Gif(mpeg4gif *types.InlineQueryResultMpeg4Gif) error {
	var err error
	if mpeg4gif != nil {
		if r.Mpeg4Gif == nil {
			r.Mpeg4Gif = mpeg4gif
			logs.DataWrittenSuccessfully(interfaceResult, "Mpeg4 Gif")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WritePhoto(ph *types.InlineQueryResultPhoto) error {
	var err error
	if ph != nil {
		if r.Photo == nil {
			r.Photo = ph
			logs.DataWrittenSuccessfully(interfaceResult, "Photo")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteVenue(ven *types.InlineQueryResultVenue) error {
	var err error
	if ven != nil {
		if r.Venue == nil {
			r.Venue = ven
			logs.DataWrittenSuccessfully(interfaceResult, "Venue")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteVideo(vd *types.InlineQueryResultVideo) error {
	var err error
	if vd != nil {
		if r.Video == nil {
			r.Video = vd
			logs.DataWrittenSuccessfully(interfaceResult, "Video")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (r *result) WriteVoice(vc *types.InlineQueryResultVoice) error {
	var err error
	if vc != nil {
		if r.Voice == nil {
			r.Voice = vc
			logs.DataWrittenSuccessfully(interfaceResult, "Voice")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteTitle(title string) error {
	var err error
	t := len([]rune(title))
	if t > 0 && t <= 32 {
		if p.Title == "" {
			p.Title = title
			logs.DataWrittenSuccessfully(interfacePayment, "Title")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteDescription(description string) error {
	var err error
	t := len([]rune(description))
	if t > 0 && t <= 255 {
		if p.Description == "" {
			p.Description = description
			logs.DataWrittenSuccessfully(interfacePayment, "Description")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WritePayload(payload string) error {
	var err error
	t := len([]byte(payload))
	if t > 0 && t <= 128 {
		if p.Payload == "" {
			p.Payload = payload
			logs.DataWrittenSuccessfully(interfacePayment, "Payload")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteProviderToken(token string) error {
	var err error
	if err = checkStringValue(token, p.ProviderToken); err == nil {
		p.ProviderToken = token
		logs.DataWrittenSuccessfully(interfacePayment, "Provider Token")
	}
	return err
}

func (p *payment) WriteCurrency(currency string) error {
	var err error
	if _, ok := Currencies[currency]; ok {
		if (currency == telegramStars && (p.Prices == nil || p.Prices[0].Amount <= 2500)) || currency != telegramStars {
			if p.Currency == "" {
				p.Currency = currency
				logs.DataWrittenSuccessfully(interfacePayment, "Currency")
			} else {
				err = code10()
			}
		} else {
			err = code20()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WritePrices(prices []*types.LabeledPrice) error {
	var err error
	if len(prices) != 0 {

		for i := 0; (i < len(prices)) && (err == nil); i++ {
			if prices[i] == nil {
				err = code5()
			} else {
				if prices[i].Label == telegramStars && prices[i].Amount > 2500 {
					err = code20()
				} else if prices[i].Label != telegramStars && p.SubscriptionPeriod != 0 {
					err = code20()
				}
			}
		}

		if err == nil {
			if len(p.Prices) == 0 {
				p.Prices = prices
				logs.DataWrittenSuccessfully(interfacePayment, "Prices")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteStartParameter(prm string) error {
	var err error
	if err = checkStringValue(prm, p.StartParameter); err == nil {
		p.StartParameter = prm
		logs.DataWrittenSuccessfully(interfacePayment, "Start Parameter")
	}
	return err
}

func (p *payment) WriteProviderData(data string) error {
	var err error
	if err = checkStringValue(data, p.ProviderData); err == nil {
		p.ProviderData = data
		logs.DataWrittenSuccessfully(interfacePayment, "Provider Data")
	}
	return err
}

func (p *payment) WriteSubscriptionPeriod(period int) error {
	var err error
	if (p.Currency == "" || p.Currency == telegramStars) && (p.Prices == nil || p.Prices[0].Amount <= 2500) && period == Month {
		if p.SubscriptionPeriod == 0 {
			p.SubscriptionPeriod = period
			logs.DataWrittenSuccessfully(interfacePayment, "Subscription Period")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteMaxTipAmount(amount int) error {
	var err error
	if p.Currency != telegramStars && p.SubscriptionPeriod == 0 {
		if amount > 0 {
			if p.MaxTipAmount == 0 {
				p.MaxTipAmount = amount
				logs.DataWrittenSuccessfully(interfacePayment, "Max Tip Amount")
			} else {
				err = code10()
			}
		} else {
			err = code20()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteSuggestedTipAmounts(amounts []int) error {
	var err error
	if len(amounts) > 0 && len(amounts) <= 4 {
		for i := 0; i < len(amounts) && err == nil; i++ {
			if amounts[i] < 1 {
				err = code5()
			}
		}
		if err == nil {
			if p.SuggestedTipAmounts == nil {
				p.SuggestedTipAmounts = amounts
				logs.DataWrittenSuccessfully(interfacePayment, "Suggested Tip Amounts")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WritePhotoUrl(url string) error {
	var err error
	if err = checkStringValue(url, p.PhotoUrl); err == nil {
		p.PhotoUrl = url
		logs.DataWrittenSuccessfully(interfacePayment, "Photo Url")
	}
	return err
}

func (p *payment) WritePhotoSize(size int) error {
	var err error
	if err = checkIntegerValue(size, p.PhotoSize); err == nil {
		p.PhotoSize = size
		logs.DataWrittenSuccessfully(interfacePayment, "Photo Size")
	}
	return err
}

func (p *payment) WritePhotoWidth(width int) error {
	var err error
	if err = checkIntegerValue(width, p.PhotoWidth); err == nil {
		p.PhotoWidth = width
		logs.DataWrittenSuccessfully(interfacePayment, "Photo Width")
	}
	return err
}

func (p *payment) WritePhotoHeight(height int) error {
	var err error
	if err = checkIntegerValue(height, p.PhotoHeight); err == nil {
		p.PhotoHeight = height
		logs.DataWrittenSuccessfully(interfacePayment, "Photo Height")
	}
	return err
}

func (p *payment) WriteNeedName() error {
	var err error
	if !p.NeedName {
		p.NeedName = true
		logs.SettedParam("Need Name", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteNeedPhoneNumber() error {
	var err error
	if !p.NeedPhoneNumber {
		p.NeedPhoneNumber = true
		logs.SettedParam("Need Phone Number", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteNeedEmail() error {
	var err error
	if !p.NeedEmail {
		p.NeedEmail = true
		logs.SettedParam("Need Email", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteNeedShippingAddress() error {
	var err error
	if !p.NeedShippingAddress {
		p.NeedShippingAddress = true
		logs.SettedParam("Need Shipping Address", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteSendPhoneNumberToProvider() error {
	var err error
	if !p.SendPhoneNumberToProvider {
		p.SendPhoneNumberToProvider = true
		logs.SettedParam("Send Phone Number To Provider", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteSendEmailToProvider() error {
	var err error
	if !p.SendEmailToProvider {
		p.SendEmailToProvider = true
		logs.SettedParam("Send Email To Provider", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteIsFlexible() error {
	var err error
	if !p.IsFlexible {
		p.IsFlexible = true
		logs.SettedParam("Is Flexible", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteShippingID(id string) error {
	var err error
	if err = checkStringValue(id, p.ShippingID); err == nil {
		p.ShippingID = id
		logs.DataWrittenSuccessfully(interfacePayment, "Shipping ID")
	}
	return err
}

func (p *payment) WriteOK() error {
	var err error
	if !p.OK {
		p.OK = true
		logs.SettedParam("OK", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteShippingOptions(options []*types.ShippingOption) error {
	var err error
	if options != nil {
		for i := 0; i < len(options) && err == nil; i++ {
			if options[i] == nil {
				err = code5()
			}
		}
		if err == nil {
			if p.ShippingOptions == nil {
				p.ShippingOptions = options
				logs.DataWrittenSuccessfully(interfacePayment, "Shipping Options")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteErrorMessage(msg string) error {
	var err error
	if err = checkStringValue(msg, p.ErrorMessage); err == nil {
		p.ErrorMessage = msg
		logs.DataWrittenSuccessfully(interfacePayment, "Error Message")
	}
	return err
}

func (p *payment) WritePreCheckoutID(id string) error {
	var err error
	if err = checkStringValue(id, p.PreCheckoutID); err == nil {
		p.PreCheckoutID = id
		logs.DataWrittenSuccessfully(interfacePayment, "Pre-Checkout ID")
	}
	return err
}

func (p *payment) WriteLimit(limit int) error {
	var err error
	if limit > 0 && limit <= 100 {
		if p.Limit == 0 {
			p.Limit = limit
			logs.DataWrittenSuccessfully(interfacePayment, "Limit")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (p *payment) WriteOffset(offset int) error {
	var err error
	if err = checkIntegerValue(offset, p.Offset); err == nil {
		p.Offset = offset
		logs.DataWrittenSuccessfully(interfacePayment, "Offset")
	}
	return err
}

func (p *payment) WriteTelegramPaymentChargeID(id string) error {
	var err error
	if err = checkStringValue(id, p.TelegramPaymentChargeID); err == nil {
		p.TelegramPaymentChargeID = id
		logs.DataWrittenSuccessfully(interfacePayment, "Telegram Payment Charge ID")
	}
	return err
}

func (p *payment) WriteIsCanceled() error {
	var err error
	if !p.IsCanceled {
		p.IsCanceled = true
		logs.SettedParam("Is Canceled", interfacePayment, true)
	} else {
		err = code10()
	}
	return err
}

func (g *game) WriteShortName(name string) error {
	var err error
	if err = checkStringValue(name, g.ShortName); err == nil {
		g.ShortName = name
		logs.DataWrittenSuccessfully(interfaceGame, "Short Name")
	}
	return err
}

func (g *game) WriteScore(score int) error {
	var err error
	if score >= 0 {
		if g.Score == 0 {
			g.Score = score
			logs.DataWrittenSuccessfully(interfaceGame, "Score")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (g *game) WriteForce() error {
	var err error
	if !g.Force {
		g.Force = true
		logs.SettedParam("Force", interfaceGame, true)
	} else {
		err = code10()
	}
	return err
}

func (g *game) WriteDisableEditMessage() error {
	var err error
	if !g.DisableEditMessage {
		g.DisableEditMessage = true
		logs.SettedParam("Disable Edit Message", interfaceGame, true)
	} else {
		err = code10()
	}
	return err
}
