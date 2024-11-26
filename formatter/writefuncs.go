package formatter

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/fmerrors"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

func code12() error {
	err := new(fmerrors.FME)
	err.Code = 12
	err.String = "incorrect type of file"
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

func code999() error {
	err := new(fmerrors.FME)
	err.Code = 999
	err.String = "incorrect input data type"
	return err
}

func (ph *photo) writePhoto(photo, object string) error {
	var err error
	if ph.Photo == "" {
		ph.Photo, ph.Media = photo, photo
		logs.DataWrittenSuccessfully(interfacePhoto, object)
	} else {
		err = code10()
	}
	return err
}

func (ph *photo) WritePhotoStorage(photo string) error {
	var err error
	if photo != "" {
		if err = ph.isCorrectType(photo); err == nil {
			ph.gottenFrom = Storage
			err = ph.writePhoto(photo, "Photo From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (ph *photo) WritePhotoTelegram(photo string) error {
	var err error
	if photo != "" {
		ph.gottenFrom = Telegram
		err = ph.writePhoto(photo, "Photo From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (ph *photo) WritePhotoInternet(photo string) error {
	var err error
	if photo != "" {
		ph.gottenFrom = Internet
		err = ph.writePhoto(photo, "Photo From The Internet")
	} else {
		err = code20()
	}
	return err
}

func (ph *photo) WriteCaption(caption string) error {
	var err error
	if caption != "" {
		if ph.Caption != "" {
			err = code10()
		} else {
			ph.Caption = caption
			logs.DataWrittenSuccessfully(interfacePhoto, "Caption")
		}
	} else {
		err = code20()
	}
	return err
}

func (ph *photo) WriteParseMode(parsemode string) error {
	var err error
	if parsemode != types.HTML && parsemode != types.Markdown && parsemode != types.MarkdownV2 {
		err = code20()
	} else {
		if ph.ParseMode != "" {
			err = code10()
		} else {
			ph.ParseMode = parsemode
			logs.DataWrittenSuccessfully(interfacePhoto, "Parse Mode")
		}
	}
	return err
}

func (ph *photo) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if len(captionEntities) != 0 {

		for i := 0; (i < len(captionEntities)) && (err == nil); i++ {
			if captionEntities[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if len(ph.CaptionEntities) == 0 {
				ph.CaptionEntities = captionEntities
				logs.DataWrittenSuccessfully(interfacePhoto, "Caption Entities")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
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

func (vd *video) writeVideo(video, object string) error {
	var err error
	if vd.Video == "" {
		vd.Video, vd.Media = video, video
		logs.DataWrittenSuccessfully(interfaceVideo, object)
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteVideoStorage(video string) error {
	var err error
	if video != "" {
		if err = vd.isCorrectType(video); err == nil {
			vd.videoGottenFrom = Storage
			err = vd.writeVideo(video, "Video From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteVideoTelegram(video string) error {
	var err error
	if video != "" {
		vd.videoGottenFrom = Telegram
		err = vd.writeVideo(video, "Video From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteVideoInternet(video string) error {
	var err error
	if video != "" {
		vd.videoGottenFrom = Internet
		err = vd.writeVideo(video, "Video From the Internet")
	} else {
		err = code20()
	}
	return err
}

func (vd *video) writeThumbnail(thumbnail, object string) error {
	var err error
	if vd.Thumbnail == "" {
		vd.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceVideo, object)
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if thumbnail != "" {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			vd.thumbnailGottenFrom = Storage
			err = vd.writeThumbnail(thumbnail, "Thumbnail From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if thumbnail != "" {
		vd.thumbnailGottenFrom = Telegram
		err = vd.writeThumbnail(thumbnail, "Thumbnail From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if thumbnail != "" {
		vd.thumbnailGottenFrom = Internet
		err = vd.writeThumbnail(thumbnail, "Thumbnail From the Internet")
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteCaption(caption string) error {
	var err error
	if caption != "" {
		if vd.Caption == "" {
			vd.Caption = caption
			logs.DataWrittenSuccessfully(interfaceVideo, "Caption")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteParseMode(parsemode string) error {
	var err error
	if parsemode != types.HTML && parsemode != types.Markdown && parsemode != types.MarkdownV2 {
		err = code20()
	} else {
		if vd.ParseMode == "" {
			vd.ParseMode = parsemode
			logs.DataWrittenSuccessfully(interfaceVideo, "Parse Mode")
		} else {
			err = code10()
		}
	}
	return err
}

func (vd *video) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if len(captionEntities) != 0 {

		for i := 0; (i < len(captionEntities)) && (err == nil); i++ {
			if captionEntities[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if len(vd.CaptionEntities) == 0 {
				vd.CaptionEntities = captionEntities
				logs.DataWrittenSuccessfully(interfaceVideo, "Caption Entities")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
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
	if width > 0 {
		if vd.Width == 0 {
			vd.Width = width
			logs.DataWrittenSuccessfully(interfaceVideo, "Width")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteHeight(height int) error {
	var err error
	if height > 0 {
		if vd.Height == 0 {
			vd.Height = height
			logs.DataWrittenSuccessfully(interfaceVideo, "Height")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteDuration(duration int) error {
	var err error
	if duration > 0 {
		if vd.Duration == 0 {
			vd.Duration = duration
			logs.DataWrittenSuccessfully(interfaceVideo, "Duration")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vd *video) WriteSupportsStreaming() error {
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

func (ad *audio) writeAudio(audio, object string) error {
	var err error
	if ad.Audio == "" {
		ad.Audio, ad.Media = audio, audio
		logs.DataWrittenSuccessfully(interfaceAudio, object)
	} else {
		err = code10()
	}
	return err
}

func (ad *audio) WriteAudioStorage(audio string) error {
	var err error
	if audio != "" {
		if err = ad.isCorrectType(audio); err == nil {
			ad.audioGottenFrom = Storage
			err = ad.writeAudio(audio, "Audio From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteAudioTelegram(audio string) error {
	var err error
	if audio != "" {
		ad.audioGottenFrom = Telegram
		err = ad.writeAudio(audio, "Audio From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteAudioInternet(audio string) error {
	var err error
	if audio != "" {
		ad.audioGottenFrom = Internet
		err = ad.writeAudio(audio, "Audio From the Internet")
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) writeThumbnail(thumbnail, object string) error {
	var err error
	if ad.Thumbnail == "" {
		ad.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceAudio, object)
	} else {
		err = code10()
	}
	return err
}

func (ad *audio) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if thumbnail != "" {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			ad.thumbnailGottenFrom = Storage
			err = ad.writeThumbnail(thumbnail, "Thumbnail From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if thumbnail != "" {
		ad.thumbnailGottenFrom = Telegram
		err = ad.writeThumbnail(thumbnail, "Thumbnail From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if thumbnail != "" {
		ad.thumbnailGottenFrom = Internet
		err = ad.writeThumbnail(thumbnail, "Thumbnail From the Internet")
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteCaption(caption string) error {
	var err error
	if caption != "" {
		if ad.Caption == "" {
			ad.Caption = caption
			logs.DataWrittenSuccessfully(interfaceAudio, "Caption")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteParseMode(parsemode string) error {
	var err error
	if (parsemode != types.HTML) && (parsemode != types.Markdown) && (parsemode != types.MarkdownV2) {
		err = code20()
	} else {
		if ad.ParseMode == "" {
			ad.ParseMode = parsemode
			logs.DataWrittenSuccessfully(interfaceAudio, "Parse Mode")
		} else {
			err = code10()
		}
	}
	return err
}

func (ad *audio) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if len(captionEntities) != 0 {

		for i := 0; (i < len(captionEntities)) && (err == nil); i++ {
			if captionEntities[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if len(ad.CaptionEntities) == 0 {
				ad.CaptionEntities = captionEntities
				logs.DataWrittenSuccessfully(interfaceAudio, "Caption Entities")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteDuration(duration int) error {
	var err error
	if duration != 0 {
		if ad.Duration == 0 {
			ad.Duration = duration
			logs.DataWrittenSuccessfully(interfaceAudio, "Duration")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WritePerformer(performer string) error {
	var err error
	if performer != "" {
		if ad.Performer == "" {
			ad.Performer = performer
			logs.DataWrittenSuccessfully(interfaceAudio, "Performer")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) WriteTitle(title string) error {
	var err error
	if title != "" {
		if ad.Title == "" {
			ad.Title = title
			logs.DataWrittenSuccessfully(interfaceAudio, "Title")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (ad *audio) GetResponse() types.Audio {
	return ad.response
}

func (dc *document) writeDocument(document, object string) error {
	var err error
	if dc.Document == "" {
		dc.Document = document
		logs.DataWrittenSuccessfully(interfaceDocument, object)
	} else {
		err = code10()
	}
	return err
}

func (dc *document) WriteDocumentStorage(document string) error {
	var err error
	if document != "" {
		dc.documentGottenFrom = Storage
		err = dc.writeDocument(document, "Document From Storage")
	} else {
		err = code20()
	}
	return err
}

func (dc *document) WriteDocumentTelegram(document string) error {
	var err error
	if document != "" {
		dc.documentGottenFrom = Telegram
		err = dc.writeDocument(document, "Document From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (dc *document) WriteDocumentInternet(document string) error {
	var err error
	if document != "" {
		dc.documentGottenFrom = Internet
		err = dc.writeDocument(document, "Document From the Internet")
	} else {
		err = code20()
	}
	return err
}

func (dc *document) writeThumbnail(thumbnail, object string) error {
	var err error
	if dc.Thumbnail == "" {
		dc.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceDocument, object)
	} else {
		err = code10()
	}
	return err
}

func (dc *document) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if thumbnail != "" {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			dc.thumbnailGottenFrom = Storage
			err = dc.writeThumbnail(thumbnail, "Thumbnail From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (dc *document) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if thumbnail != "" {
		dc.thumbnailGottenFrom = Telegram
		err = dc.writeThumbnail(thumbnail, "Thumbnail From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (dc *document) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if thumbnail != "" {
		dc.thumbnailGottenFrom = Internet
		err = dc.writeThumbnail(thumbnail, "Thumbnail From the Internet")
	} else {
		err = code20()
	}
	return err
}

func (dc *document) WriteCaption(caption string) error {
	var err error
	if caption != "" {
		if dc.Caption == "" {
			dc.Caption = caption
			logs.DataWrittenSuccessfully(interfaceDocument, "Caption")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (dc *document) WriteParseMode(parsemode string) error {
	var err error
	if (parsemode != types.HTML) && (parsemode != types.Markdown) && (parsemode != types.MarkdownV2) {
		err = code20()
	} else {
		if dc.ParseMode == "" {
			dc.ParseMode = parsemode
			logs.DataWrittenSuccessfully(interfaceDocument, "Parse Mode")
		} else {
			err = code10()
		}
	}
	return err
}

func (dc *document) WriteCaptionEntities(captionEntities []*types.MessageEntity) error {
	var err error
	if len(captionEntities) != 0 {

		for i := 0; (i < len(captionEntities)) && (err == nil); i++ {
			if captionEntities[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if len(dc.CaptionEntities) == 0 {
				dc.CaptionEntities = captionEntities
				logs.DataWrittenSuccessfully(interfaceDocument, "Caption Entities")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
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

func (an *animation) writeAnimation(animation, object string) error {
	var err error
	if an.Animation == "" {
		an.Animation = animation
		logs.DataWrittenSuccessfully(interfaceAnimation, object)
	} else {
		err = code10()
	}
	return err
}

func (an *animation) WriteAnimationStorage(animation string) error {
	var err error
	if animation != "" {
		if err = an.isCorrectType(animation); err == nil {
			an.animationGottenFrom = Storage
			err = an.writeAnimation(animation, "Animation From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteAnimationTelegram(animation string) error {
	var err error
	if animation != "" {
		an.animationGottenFrom = Telegram
		err = an.writeAnimation(animation, "Animation From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteAnimationInternet(animation string) error {
	var err error
	if animation != "" {
		an.animationGottenFrom = Internet
		err = an.writeAnimation(animation, "Animation From Internet")
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteDuration(duration int) error {
	var err error
	if duration > 0 {
		if an.Duration == 0 {
			an.Duration = duration
			logs.DataWrittenSuccessfully(interfaceAnimation, "Duration")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteWidth(width int) error {
	var err error
	if width > 0 {
		if an.Width == 0 {
			an.Width = width
			logs.DataWrittenSuccessfully(interfaceAnimation, "Width")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteHeight(height int) error {
	var err error
	if height > 0 {
		if an.Height == 0 {
			an.Height = height
			logs.DataWrittenSuccessfully(interfaceAnimation, "Height")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (an *animation) writeThumbnail(thumbnail, object string) error {
	var err error
	if an.Thumbnail == "" {
		an.Thumbnail = thumbnail
		logs.DataWrittenSuccessfully(interfaceAnimation, object)
	} else {
		err = code10()
	}
	return err
}

func (an *animation) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if thumbnail != "" {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			an.thumbnailGottenFrom = Storage
			err = an.writeThumbnail(thumbnail, "Thumbnail From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteThumbnailTelegram(thumbnail string) error {
	var err error
	if thumbnail != "" {
		an.thumbnailGottenFrom = Telegram
		err = an.writeThumbnail(thumbnail, "Thumbnail From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (an *animation) WriteThumbnailInternet(thumbnail string) error {
	var err error
	if thumbnail != "" {
		an.thumbnailGottenFrom = Internet
		err = an.writeThumbnail(thumbnail, "Thumbnail From Internet")
	} else {
		err = code20()
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

func (vc *voice) writeVoice(voice, object string) error {
	var err error
	if vc.Voice == "" {
		vc.Voice = voice
		logs.DataWrittenSuccessfully(interfaceVoice, object)
	} else {
		err = code10()
	}
	return err
}

func (vc *voice) WriteVoiceStorage(voice string) error {
	var err error
	if voice != "" {
		if err = vc.isCorrectType(voice); err == nil {
			vc.gottenFrom = Storage
			err = vc.writeVoice(voice, "Animation From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (vc *voice) WriteVoiceTelegram(voiceID string) error {
	var err error
	if voiceID != "" {
		vc.gottenFrom = Telegram
		err = vc.writeVoice(voiceID, "Animation From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (vc *voice) WriteVoiceInternet(URL string) error {
	var err error
	if URL != "" {
		vc.gottenFrom = Internet
		err = vc.writeVoice(URL, "Animation From Internet")
	} else {
		err = code20()
	}
	return err
}

func (vc *voice) WriteDuration(duration int) error {
	var err error
	if duration > 0 {
		if vc.Duration == 0 {
			vc.Duration = duration
			logs.DataWrittenSuccessfully(interfaceVoice, "Duration")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vc *voice) GetResponse() types.Voice {
	return vc.response
}

func (vdn *videonote) writeVideoNote(videonote, object string) error {
	var err error
	if vdn.VideoNote == "" {
		vdn.VideoNote = videonote
		logs.DataWrittenSuccessfully(interfaceVoice, object)
	} else {
		err = code10()
	}
	return err
}

func (vdn *videonote) WriteVideoNoteStorage(videonote string) error {
	var err error
	if videonote != "" {
		if err = vdn.isCorrectType(videonote); err == nil {
			vdn.videoGottenFrom = Storage
			err = vdn.writeVideoNote(videonote, "Video-Note From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) WriteVideoNoteTelegram(vdnID string) error {
	var err error
	if vdnID != "" {
		vdn.videoGottenFrom = Telegram
		err = vdn.writeVideoNote(vdnID, "Video-Note From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) WriteVideoNoteInternet(URL string) error {
	var err error
	if URL != "" {
		vdn.videoGottenFrom = Internet
		err = vdn.writeVideoNote(URL, "Video-Note From Internet")
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) WriteDuration(duration int) error {
	var err error
	if duration > 0 {
		if vdn.Duration == 0 {
			vdn.Duration = duration
			logs.DataWrittenSuccessfully(interfaceVideoNote, "Duration")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) WriteLength(length int) error {
	var err error
	if length > 0 {
		if vdn.Length == 0 {
			vdn.Length = length
			logs.DataWrittenSuccessfully(interfaceVideoNote, "Length")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) writeThumbnail(thumb, object string) error {
	var err error
	if vdn.Thumbnail == "" {
		vdn.Thumbnail = thumb
		logs.DataWrittenSuccessfully(interfaceVideoNote, object)
	} else {
		err = code10()
	}
	return err
}

func (vdn *videonote) WriteThumbnailStorage(thumbnail string) error {
	var err error
	if thumbnail != "" {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			vdn.thumbnailGottenFrom = Storage
			err = vdn.writeThumbnail(thumbnail, "Thumbnail From Storage")
		}
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) WriteThumbnailTelegram(thumbnailID string) error {
	var err error
	if thumbnailID != "" {
		vdn.thumbnailGottenFrom = Telegram
		err = vdn.writeThumbnail(thumbnailID, "Thumbnail From Telegram")
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) WriteThumbnailInternet(URL string) error {
	var err error
	if URL != "" {
		vdn.thumbnailGottenFrom = Internet
		err = vdn.writeThumbnail(URL, "Thumbnail From Internet")
	} else {
		err = code20()
	}
	return err
}

func (vdn *videonote) GetResponse() types.VideoNote {
	return vdn.response
}

func (inf *information) WriteString(text string) error {
	var err error
	if text != "" {
		if inf.Text == "" {
			inf.Text = text
			logs.DataWrittenSuccessfully(interfaceInf, "Text")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteParseMode(parsemode string) error {
	var err error
	if (parsemode != types.HTML) && (parsemode != types.Markdown) && (parsemode != types.MarkdownV2) {
		err = code20()
	} else {
		if inf.ParseMode == "" {
			inf.ParseMode = parsemode
			logs.DataWrittenSuccessfully(interfaceInf, "Parse Mode")
		} else {
			err = code10()
		}
	}
	return err
}

func (inf *information) WriteMessageThreadID(messageID int) error {
	var err error
	if messageID > 0 {
		if inf.MessageID == 0 {
			inf.MessageThreadID = messageID
			logs.DataWrittenSuccessfully(interfaceInf, "Message Thread ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
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
	if messageID != "" {
		if inf.MessageEffectID == "" {
			inf.MessageEffectID = messageID
			logs.DataWrittenSuccessfully(interfaceInf, "Message Effect ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WriteEntities(entities []*types.MessageEntity) error {
	var err error
	if len(entities) != 0 {

		for i := 0; (i < len(entities)) && (err == nil); i++ {
			if entities[i] == nil {
				err = code5()
			}
		}

		if err == nil {
			if len(inf.Entities) == 0 {
				inf.Entities = entities
				logs.DataWrittenSuccessfully(interfaceInf, "Entities")
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
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
	if messageID > 0 {
		if inf.MessageID == 0 {
			inf.MessageID = messageID
			logs.DataWrittenSuccessfully(interfaceInf, "Message ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
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
	if caption != "" {
		if inf.Caption == "" {
			inf.Caption = caption
			logs.DataWrittenSuccessfully(interfaceInf, "Caption")
		} else {
			err = code10()
		}
	} else {
		err = code20()
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
			switch ch := reply.ChatID.(type) {
			case int:
				if ch <= 0 {
					err = code20()
				}
			case string:
				if ch == "" {
					err = code20()
				}
			default:
				err = code999()
			}

			if err == nil {
				if inf.ReplyParameters == nil {
					inf.ReplyParameters = reply
					logs.DataWrittenSuccessfully(interfaceInf, "Reply Parameters")
				} else {
					err = code10()
				}
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
	if amount > 0 {
		if inf.StarCount == 0 {
			inf.StarCount = amount
			logs.DataWrittenSuccessfully(interfaceInf, "Star Count")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inf *information) WritePayload(payload string) error {
	var err error
	if payload != "" {
		if inf.Payload == "" {
			inf.Payload = payload
			logs.DataWrittenSuccessfully(interfaceInf, "Payload")
		} else {
			err = code10()
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
	if connectionID != "" {
		if ch.BusinessConnectionID == "" {
			ch.BusinessConnectionID = connectionID
			logs.DataWrittenSuccessfully(interfaceChat, "Business Connection ID")
		} else {
			err = code10()
		}
	} else {
		err = code20()
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

func (in *inline) Set(plan []int) error {
	var err error
	if len(plan) > 0 {
		in.Keyboard = new(inlineKeyboard)
		in.Keyboard.InlineKeyboard = make([][]*inlineKeyboardButton, len(plan))
		for i := range in.Keyboard.InlineKeyboard {
			in.Keyboard.InlineKeyboard[i] = make([]*inlineKeyboardButton, plan[i])
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

	if (line >= 0) && (pos >= 0) && len(in.Keyboard.InlineKeyboard) > line && len(in.Keyboard.InlineKeyboard[line]) > pos {

		if in.Keyboard.InlineKeyboard[line][pos] == nil {
			in.Keyboard.InlineKeyboard[line][pos] = new(inlineKeyboardButton)
			but = in.Keyboard.InlineKeyboard[line][pos]
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return but, err
}

func (inb *inlineKeyboardButton) WriteString(text string) error {
	var err error
	if text != "" {
		if inb.Text == "" {
			inb.Text = text
			logs.DataWrittenSuccessfully(inbtn, "Text")
		} else {
			err = code10()
		}
	} else {
		err = code20()
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
	if url != "" {
		if inb.Url == "" {
			err = inb.checkOthers()
			if err == nil {
				inb.storage[wURL] = added
				inb.Url = url
				logs.DataWrittenSuccessfully(inbtn, "URL")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteCallbackData(text string) error {
	var err error
	if text != "" {
		if inb.CallbackData == "" {
			err = inb.checkOthers()
			if err == nil {
				inb.storage[wCallback] = added
				inb.CallbackData = text
				logs.DataWrittenSuccessfully(inbtn, "Callback Data")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteWebApp(wbapp *types.WebAppInfo) error {
	var err error
	if wbapp != nil && wbapp.Url != "" {
		if inb.WebApp == nil {
			err = inb.checkOthers()
			if err == nil {
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
			err = inb.checkOthers()
			if err == nil {
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
	if sw != "" {
		if inb.SwitchInlineQuery == "" {
			err = inb.checkOthers()
			if err == nil {
				inb.storage[wSwitchIn] = added
				inb.SwitchInlineQuery = sw
				logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryCurrentChat(swcch string) error {
	var err error
	if swcch != "" {
		if inb.SwitchInlineQueryCurrentChat == "" {
			err = inb.checkOthers()
			if err == nil {
				inb.storage[wSwitchInQuery] = added
				inb.SwitchInlineQueryCurrentChat = swcch
				logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Current Chat")
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat) error {
	var err error
	if sw != nil {
		if inb.SwitchInlineQueryChosenChat == nil {
			err = inb.checkOthers()
			if err == nil {
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
	if inb.CallbackGame == nil {
		err = inb.checkOthers()
		if err == nil {
			inb.storage[wGame] = added
			inb.CallbackGame = game
			logs.DataWrittenSuccessfully(inbtn, "Callback Game")
		}
	} else {
		err = code10()
	}
	return err
}

func (inb *inlineKeyboardButton) WritePay() error {
	var err error
	if !inb.Pay {
		err = inb.checkOthers()
		if err == nil {
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
		rp.Keyboard = new(replyKeyboard)
		rp.Keyboard.Keyboard = make([][]*replyKeyboardButton, len(plan))
		for i := range rp.Keyboard.Keyboard {
			rp.Keyboard.Keyboard[i] = make([]*replyKeyboardButton, plan[i])
		}
	} else {
		err = code20()
	}
	return err
}

func (rp *reply) WriteIsPersistent() error {
	var err error
	if !rp.Keyboard.IsPersistent {
		rp.Keyboard.IsPersistent = true
		logs.SettedParam("Is Persistent", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteResizeKeyboard() error {
	var err error
	if !rp.Keyboard.ResizeKeyboard {
		rp.Keyboard.ResizeKeyboard = true
		logs.SettedParam("Resize Keyboard", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteOneTimeKeyboard() error {
	var err error
	if !rp.Keyboard.OneTimeKeyboard {
		rp.Keyboard.OneTimeKeyboard = true
		logs.SettedParam("One Time Keyboard", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteInputFieldPlaceholder(placeholder string) error {
	var err error
	if placeholder != "" {
		if rp.Keyboard.InputFieldPlaceholder == "" {
			rp.Keyboard.InputFieldPlaceholder = placeholder
			logs.DataWrittenSuccessfully(interfaceReplyKB, "Input Field Placeholder")
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (rp *reply) WriteSelective() error {
	var err error
	if !rp.Keyboard.Selective {
		rp.Keyboard.Selective = true
		logs.SettedParam("Selective", interfaceReplyKB, true)
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) NewButton(line, pos int) (IReplyButton, error) {
	var err error
	but := new(replyKeyboardButton)

	if (line >= 0) && (pos >= 0) && (len(rp.Keyboard.Keyboard) > line) && (len(rp.Keyboard.Keyboard[line]) > pos) {

		if rp.Keyboard.Keyboard[line][pos] == nil {
			rp.Keyboard.Keyboard[line][pos] = new(replyKeyboardButton)
			but = rp.Keyboard.Keyboard[line][pos]
		} else {
			err = code10()
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
	if text != "" {
		if rpb.Text == "" {
			rpb.Text = text
			logs.DataWrittenSuccessfully(rpbtn, "Text")

		} else {
			err = code10()
		}
	} else {
		err = code20()
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
	if (poll != nil) && (poll.Type != "") {
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
