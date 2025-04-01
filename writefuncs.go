package fmtogram

import (
	"fmt"
	"net/http"

	"github.com/iamissahar/Fmtogram/methods"
)

type emojich struct {
	permissibleQuantity int
	amount              int
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
	if (parsemode != HTML) && (parsemode != Markdown) && (parsemode != MarkdownV2) {
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

func checkEntities(entities, cell []*MessageEntity) error {
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
			ph.Photo, ph.Media = photo, fmt.Sprintf("attach://%s", photo)
			ph.Type = "photo"
		}
	}
	return err
}

func (ph *photo) WritePhotoTelegram(photo string) error {
	var err error
	if err = checkStringValue(photo, ph.Photo); err == nil {
		ph.gottenFrom = Tg
		ph.Photo, ph.Media = photo, photo
		ph.Type = "photo"
	}
	return err
}

func (ph *photo) WritePhotoInternet(photo string) error {
	var err error
	if err = checkStringValue(photo, ph.Photo); err == nil {
		ph.gottenFrom = Internet
		ph.Photo, ph.Media = photo, photo
		ph.Type = "photo"
	}
	return err
}

func (ph *photo) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, ph.Caption); err == nil {
		ph.Caption = caption
	}
	return err
}

func (ph *photo) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, ph.ParseMode); err == nil {
		ph.ParseMode = parsemode
	}
	return err
}

func (ph *photo) WriteCaptionEntities(captionEntities []*MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, ph.CaptionEntities); err == nil {
		ph.CaptionEntities = captionEntities
	}
	return err
}

func (ph *photo) WriteShowCaptionAboveMedia() error {
	var err error
	if ph.ShowCaptionAboveMedia {
		err = code10()
	} else {
		ph.ShowCaptionAboveMedia = true
	}
	return err
}

func (ph *photo) WriteHasSpoiler() error {
	var err error
	if ph.HasSpoiler {
		err = code10()
	} else {
		ph.HasSpoiler = true
	}
	return err
}

func (vd *video) WriteVideoStorage(video string) error {
	var err error
	if err = checkStringValue(video, vd.Video); err == nil {
		if err = vd.isCorrectType(video); err == nil {
			vd.videoGottenFrom = Storage
			vd.Video, vd.Media = video, fmt.Sprintf("attach://%s", video)
			vd.Type = "video"
		}
	}
	return err
}

func (vd *video) WriteVideoTelegram(video string) error {
	var err error
	if err = checkStringValue(video, vd.Video); err == nil {
		vd.videoGottenFrom = Tg
		vd.Video, vd.Media = video, video
		vd.Type = "video"
	}
	return err
}

func (vd *video) WriteVideoInternet(video string) error {
	var err error
	if err = checkStringValue(video, vd.Video); err == nil {
		vd.videoGottenFrom = Internet
		vd.Video, vd.Media = video, video
		vd.Type = "video"
	}
	return err
}

func (vd *video) WriteThumbnail(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, vd.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			vd.thumbnailGottenFrom = Storage
			vd.Thumbnail = thumbnail
		}
	}
	return err
}

func (vd *video) WriteThumbnailID(thumbnailID string) error {
	var err error
	if err = checkStringValue(thumbnailID, vd.Thumbnail); err == nil {
		vd.thumbnailGottenFrom = Tg
		vd.Thumbnail = thumbnailID
	}
	return err
}

func (vd *video) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, vd.Caption); err == nil {
		vd.Caption = caption
	}
	return err
}

func (vd *video) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, vd.ParseMode); err == nil {
		vd.ParseMode = parsemode
	}
	return err
}

func (vd *video) WriteCaptionEntities(captionEntities []*MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, vd.CaptionEntities); err == nil {
		vd.CaptionEntities = captionEntities
	}
	return err
}

func (vd *video) WriteShowCaptionAboveMedia() error {
	var err error
	if !vd.ShowCaptionAboveMedia {
		vd.ShowCaptionAboveMedia = true
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteWidth(width int) error {
	var err error
	if err = checkIntegerValue(width, vd.Width); err == nil {
		vd.Width = width
	}
	return err
}

func (vd *video) WriteHeight(height int) error {
	var err error
	if err = checkIntegerValue(height, vd.Height); err == nil {
		vd.Height = height
	}
	return err
}

func (vd *video) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, vd.Duration); err == nil {
		vd.Duration = duration
	}
	return err
}

func (vd *video) WriteSupportsStreaming() error {
	var err error
	if !vd.SupportsStreaming {
		vd.SupportsStreaming = true
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteHasSpoiler() error {
	var err error
	if !vd.HasSpoiler {
		vd.HasSpoiler = true
	} else {
		err = code10()
	}
	return err
}

func (vd *video) WriteCoverStorage(path string) error {
	var err error
	if err = checkStringValue(path, vd.Cover); err == nil {
		vd.Cover = fmt.Sprintf("attach://%s", path)
		vd.coverGottenFrom = Storage
	}
	return err
}

func (vd *video) WriteCoverTelegram(coverID string) error {
	var err error
	if err = checkStringValue(coverID, vd.Cover); err == nil {
		vd.Cover = coverID
		vd.coverGottenFrom = Tg
	}
	return err
}

func (vd *video) WriteCoverInternet(url string) error {
	var err error
	if err = checkStringValue(url, vd.Cover); err == nil {
		vd.Cover = url
		vd.coverGottenFrom = Internet
	}
	return err
}

func (ad *audio) WriteAudioStorage(audio string) error {
	var err error
	if err = checkStringValue(audio, ad.Audio); err == nil {
		if err = ad.isCorrectType(audio); err == nil {
			ad.audioGottenFrom = Storage
			ad.Audio, ad.Media = audio, fmt.Sprintf("attach://%s", audio)
			ad.Type = "audio"
		}
	}
	return err
}

func (ad *audio) WriteAudioTelegram(audio string) error {
	var err error
	if err = checkStringValue(audio, ad.Audio); err == nil {
		ad.audioGottenFrom = Tg
		ad.Audio, ad.Media = audio, audio
		ad.Type = "audio"
	}
	return err
}

func (ad *audio) WriteAudioInternet(audio string) error {
	var err error
	if err = checkStringValue(audio, ad.Audio); err == nil {
		ad.audioGottenFrom = Internet
		ad.Audio, ad.Media = audio, audio
		ad.Type = "audio"
	}
	return err
}

func (ad *audio) WriteThumbnail(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, ad.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			ad.thumbnailGottenFrom = Storage
			ad.Thumbnail = thumbnail
		}
	}
	return err
}

func (ad *audio) WriteThumbnailID(thumbnailID string) error {
	var err error
	if err = checkStringValue(thumbnailID, ad.Thumbnail); err == nil {
		ad.thumbnailGottenFrom = Tg
		ad.Thumbnail = thumbnailID
	}
	return err
}

func (ad *audio) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, ad.Caption); err == nil {
		ad.Caption = caption
	}
	return err
}

func (ad *audio) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, ad.ParseMode); err == nil {
		ad.ParseMode = parsemode
	}
	return err
}

func (ad *audio) WriteCaptionEntities(captionEntities []*MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, ad.CaptionEntities); err == nil {
		ad.CaptionEntities = captionEntities
	}
	return err
}

func (ad *audio) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, ad.Duration); err == nil {
		ad.Duration = duration
	}
	return err
}

func (ad *audio) WritePerformer(performer string) error {
	var err error
	if err = checkStringValue(performer, ad.Performer); err == nil {
		ad.Performer = performer
	}
	return err
}

func (ad *audio) WriteTitle(title string) error {
	var err error
	if err = checkStringValue(title, ad.Title); err == nil {
		ad.Title = title
	}
	return err
}

func (dc *document) WriteDocumentStorage(document string) error {
	var err error
	if err = checkStringValue(document, dc.Document); err == nil {
		dc.documentGottenFrom = Storage
		dc.Document, dc.Media = document, fmt.Sprintf("attach://%s", document)
		dc.Type = "document"
	}
	return err
}

func (dc *document) WriteDocumentTelegram(document string) error {
	var err error
	if err = checkStringValue(document, dc.Document); err == nil {
		dc.documentGottenFrom = Tg
		dc.Document, dc.Media = document, document
		dc.Type = "document"
	}
	return err
}

func (dc *document) WriteDocumentInternet(document string) error {
	var err error
	if err = checkStringValue(document, dc.Document); err == nil {
		dc.documentGottenFrom = Internet
		dc.Document, dc.Media = document, document
		dc.Type = "document"
	}
	return err
}

func (dc *document) WriteThumbnail(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, dc.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			dc.thumbnailGottenFrom = Storage
			dc.Thumbnail = thumbnail
		}
	}
	return err
}

func (dc *document) WriteThumbnailID(thumbnailID string) error {
	var err error
	if err = checkStringValue(thumbnailID, dc.Thumbnail); err == nil {
		dc.thumbnailGottenFrom = Tg
		dc.Thumbnail = thumbnailID
	}
	return err
}

func (dc *document) WriteCaption(caption string) error {
	var err error
	if err = checkCaption(caption, dc.Caption); err == nil {
		dc.Caption = caption
	}
	return err
}

func (dc *document) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, dc.ParseMode); err == nil {
		dc.ParseMode = parsemode
	}
	return err
}

func (dc *document) WriteCaptionEntities(captionEntities []*MessageEntity) error {
	var err error
	if err = checkEntities(captionEntities, dc.CaptionEntities); err == nil {
		dc.CaptionEntities = captionEntities
	}
	return err
}

func (dc *document) WriteDisableContentTypeDetection() error {
	var err error
	if !dc.DisableContentTypeDetection {
		dc.DisableContentTypeDetection = true
	} else {
		err = code10()
	}
	return err
}

func (an *animation) WriteAnimationStorage(animation string) error {
	var err error
	if err = checkStringValue(animation, an.Animation); err == nil {
		if err = an.isCorrectType(animation); err == nil {
			an.animationGottenFrom = Storage
			an.Animation, an.Media = animation, fmt.Sprintf("attach://%s", animation)
		}
	}
	return err
}

func (an *animation) WriteAnimationTelegram(animation string) error {
	var err error
	if err = checkStringValue(animation, an.Animation); err == nil {
		an.animationGottenFrom = Tg
		an.Animation = animation
	}
	return err
}

func (an *animation) WriteAnimationInternet(animation string) error {
	var err error
	if err = checkStringValue(animation, an.Animation); err == nil {
		an.animationGottenFrom = Internet
		an.Animation = animation
	}
	return err
}

func (an *animation) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, an.Duration); err == nil {
		an.Duration = duration
	}
	return err
}

func (an *animation) WriteWidth(width int) error {
	var err error
	if err = checkIntegerValue(width, an.Width); err == nil {
		an.Width = width
	}
	return err
}

func (an *animation) WriteHeight(height int) error {
	var err error
	if err = checkIntegerValue(height, an.Height); err == nil {
		an.Height = height
	}
	return err
}

func (an *animation) WriteThumbnail(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, an.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			an.thumbnailGottenFrom = Storage
			an.Thumbnail = thumbnail
		}
	}
	return err
}

func (an *animation) WriteThumbnailID(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, an.Thumbnail); err == nil {
		an.thumbnailGottenFrom = Tg
		an.Thumbnail = thumbnail
	}
	return err
}

func (an *animation) WriteHasSpoiler() error {
	var err error
	if !an.HasSpoiler {
		an.HasSpoiler = true
	} else {
		err = code10()
	}
	return err
}

func (vc *voice) WriteVoiceStorage(voice string) error {
	var err error
	if err = checkStringValue(voice, vc.Voice); err == nil {
		if err = vc.isCorrectType(voice); err == nil {
			vc.gottenFrom = Storage
			vc.Voice = voice
		}
	}
	return err
}

func (vc *voice) WriteVoiceTelegram(voiceID string) error {
	var err error
	if err = checkStringValue(voiceID, vc.Voice); err == nil {
		vc.gottenFrom = Tg
		vc.Voice = voiceID
	}
	return err
}

func (vc *voice) WriteVoiceInternet(URL string) error {
	var err error
	if err = checkStringValue(URL, vc.Voice); err == nil {
		vc.gottenFrom = Internet
		vc.Voice = URL
	}
	return err
}

func (vc *voice) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, vc.Duration); err == nil {
		vc.Duration = duration
	}
	return err
}

func (vdn *videonote) WriteVideoNoteStorage(videonote string) error {
	var err error
	if err = checkStringValue(videonote, vdn.VideoNote); err == nil {
		if err = vdn.isCorrectType(videonote); err == nil {
			vdn.videoGottenFrom = Storage
			vdn.VideoNote = videonote
		}
	}
	return err
}

func (vdn *videonote) WriteVideoNoteTelegram(vdnID string) error {
	var err error
	if err = checkStringValue(vdnID, vdn.VideoNote); err == nil {
		vdn.videoGottenFrom = Tg
		vdn.VideoNote = vdnID
	}
	return err
}

func (vdn *videonote) WriteVideoNoteInternet(URL string) error {
	var err error
	if err = checkStringValue(URL, vdn.VideoNote); err == nil {
		vdn.videoGottenFrom = Internet
		vdn.VideoNote = URL
	}
	return err
}

func (vdn *videonote) WriteDuration(duration int) error {
	var err error
	if err = checkIntegerValue(duration, vdn.Duration); err == nil {
		vdn.Duration = duration
	}
	return err
}

func (vdn *videonote) WriteLength(length int) error {
	var err error
	if err = checkIntegerValue(length, vdn.Length); err == nil {
		vdn.Length = length
	}
	return err
}

func (vdn *videonote) WriteThumbnail(thumbnail string) error {
	var err error
	if err = checkStringValue(thumbnail, vdn.Thumbnail); err == nil {
		if err = isThumbnailCorrectType(thumbnail); err == nil {
			vdn.thumbnailGottenFrom = Storage
			vdn.Thumbnail = thumbnail
		}
	}
	return err
}

func (vdn *videonote) WriteThumbnailID(thumbnailID string) error {
	var err error
	if err = checkStringValue(thumbnailID, vdn.Thumbnail); err == nil {
		vdn.thumbnailGottenFrom = Tg
		vdn.Thumbnail = thumbnailID
	}
	return err
}

func (loc *location) WriteLatitude(lat float64) error {
	var err error
	if (lat >= -90) && (lat <= 90) {
		if loc.Latitude == 0 {
			loc.Latitude = lat
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
	}
	return err
}

func (loc *location) WriteAddress(address string) error {
	var err error
	if err = checkStringValue(address, loc.Address); err == nil {
		loc.Address = address
	}
	return err
}

func (loc *location) WriteFoursquareID(foursquareID string) error {
	var err error
	if err = checkStringValue(foursquareID, loc.FoursquareID); err == nil {
		loc.FoursquareID = foursquareID
	}
	return err
}

func (loc *location) WriteFoursquareType(foursquareType string) error {
	var err error
	if err = checkStringValue(foursquareType, loc.FoursquareType); err == nil {
		loc.FoursquareType = foursquareType
	}
	return err
}

func (loc *location) WriteGooglePlaceID(googlePlaceID string) error {
	var err error
	if err = checkStringValue(googlePlaceID, loc.GooglePlaceID); err == nil {
		loc.GooglePlaceID = googlePlaceID
	}
	return err
}

func (loc *location) WriteGooglePlaceType(googlePlaceType string) error {
	var err error
	if err = checkStringValue(googlePlaceType, loc.GooglePlaceType); err == nil {
		loc.GooglePlaceType = googlePlaceType
	}
	return err
}

func (c *contact) WritePhoneNumber(phone string) error {
	var err error
	if err = checkStringValue(phone, c.PhoneNumber); err == nil {
		c.PhoneNumber = phone
	}
	return err
}

func (c *contact) WriteFirstName(fname string) error {
	var err error
	if err = checkStringValue(fname, c.FirstName); err == nil {
		c.FirstName = fname
	}
	return err
}

func (c *contact) WriteLastName(lname string) error {
	var err error
	if err = checkStringValue(lname, c.LastName); err == nil {
		c.LastName = lname
	}
	return err
}

func (c *contact) WriteVCard(vcard string) error {
	var err error
	if err = checkStringValue(vcard, c.Vcard); err == nil {
		c.Vcard = vcard
	}
	return err
}

func (p *poll) WriteQuestion(question string) error {
	var err error
	q := len([]rune(question))
	if (q >= 1) && (q <= 300) {
		if p.Question == "" {
			p.Question = question
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
	}
	return err
}

func (p *poll) WriteQuestionEntities(entities []*MessageEntity) error {
	var err error
	if err = checkEntities(entities, p.QuestionEntities); err == nil {
		p.QuestionEntities = entities
	}
	return err
}

func (p *poll) WriteOptions(options []*PollOption) error {
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
			} else {
				err = code10()
			}
		}
	} else {
		err = code20()
	}
	return err
}

func (p *poll) WriteAnonymous(yesno bool) error {
	var err error
	if p.IsAnonymous == nil {
		p.IsAnonymous = &yesno
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
	}
	return err
}

func (p *poll) WriteExplanationEntities(entities []*MessageEntity) error {
	var err error
	if err = checkEntities(entities, p.ExplanationEntities); err == nil {
		p.ExplanationEntities = entities
	}
	return err
}

func (p *poll) WriteOpenPeriod(period int) error {
	var err error
	if p.CloseDate == 0 {
		if (period >= 5) && (period <= 600) {
			if p.OpenPeriod == 0 {
				p.OpenPeriod = period
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
	} else {
		err = code10()
	}
	return err
}

// func (l *link) WriteName(name string) error {
// 	var err error
// 	n := len([]rune(name))
// 	if n > 0 && n <= 32 {
// 		if l.Name == "" {
// 			l.Name = name
// 		} else {
// 			err = code10()
// 		}
// 	} else {
// 		err = code20()

// 	}
// 	return err
// }

// func (l *link) WriteExpireDate(date time.Duration) error {
// 	var err error
// 	if date > 0 {
// 		if l.ExpireDate == 0 {
// 			l.ExpireDate = time.Now().Unix() + int64(date.Seconds())
// 		} else {
// 			err = code10()
// 		}
// 	} else {
// 		err = code20()
// 	}
// 	return err
// }

// func (l *link) WriteMemberLimit(limit int) error {
// 	var err error
// 	if !l.JoinRequest {
// 		if limit > 0 && limit <= 99999 {
// 			if l.MemberLimit == 0 {
// 				l.MemberLimit = limit
// 			} else {
// 				err = code10()
// 			}
// 		} else {
// 			err = code20()
// 		}
// 	} else {
// 		err = code25()
// 	}
// 	return err
// }

// func (l *link) WriteJoinRequest() error {
// 	var err error
// 	if l.MemberLimit == 0 {
// 		if !l.JoinRequest {
// 			l.JoinRequest = true
// 		} else {
// 			err = code10()
// 		}
// 	} else {
// 		err = code25()
// 	}
// 	return err
// }

// func (l *link) WriteInviteLink(link string) error {
// 	var err error
// 	if err = checkStringValue(link, l.InviteLink); err == nil {
// 		l.InviteLink = link
// 	}
// 	return err
// }

// func (l *link) WriteSubscriptionPeriod(period int) error {
// 	var err error
// 	if period == Month {
// 		if l.SubscriptionPeriod == 0 {
// 			l.SubscriptionPeriod = period
// 		} else {
// 			err = code10()
// 		}
// 	} else {
// 		err = code20()
// 	}
// 	return err
// }

// func (l *link) WriteSubscriptionPrice(price int) error {
// 	var err error
// 	if price > 0 && price <= 2500 {
// 		if l.SubscriptionPrice == 0 {
// 			l.SubscriptionPrice = price
// 		} else {
// 			err = code10()
// 		}
// 	} else {
// 		err = code20()
// 	}
// 	return err
// }

func (st *sticker) WriteStickerStorage(path string) error {
	var err error
	if err = checkStringValue(path, st.Sticker); err == nil {
		if err = st.isCorrectType(path); err == nil {
			st.Sticker, st.stickerGottenFrom = path, Storage
		}
	}
	return err
}

func (st *sticker) WriteStickerTelegram(stickerID string) error {
	var err error
	if err = checkStringValue(stickerID, st.Sticker); err == nil {
		st.Sticker, st.stickerGottenFrom = stickerID, Tg
	}
	return err
}

func (st *sticker) WriteStickerInternet(url string) error {
	var err error
	if err = checkStringValue(url, st.Sticker); err == nil {
		st.Sticker, st.stickerGottenFrom = url, Tg
	}
	return err
}

func (st *sticker) WriteAssociatedEmoji(emoji string) error {
	var err error
	if err = checkStringValue(emoji, st.Emoji); err == nil {
		st.Emoji = emoji
	}
	return err
}

func (msg *Message) WriteParseMode(parsemode string) error {
	var err error
	if err = checkParseMode(parsemode, msg.fm.prm.ParseMode); err == nil {
		msg.fm.prm.ParseMode, msg.fm.prm.GiftParseMode = parsemode, parsemode
	}
	return err
}

func (msg *Message) WriteMessageThreadID(messageID int) error {
	var err error
	if err = checkIntegerValue(messageID, msg.fm.prm.MessageThreadID); err == nil {
		msg.fm.prm.MessageThreadID = messageID
	}
	return err
}

func (msg *Message) WriteDisableNotification() error {
	var err error
	if !msg.fm.prm.DisableNotification {
		msg.fm.prm.DisableNotification = true
	} else {
		err = code10()
	}
	return err
}

func (msg *Message) WriteProtectContent() error {
	var err error
	if !msg.fm.prm.ProtectContent {
		msg.fm.prm.ProtectContent = true
	} else {
		err = code10()
	}
	return err
}

func (msg *Message) WriteMessageEffectID(messageID string) error {
	var err error
	if err = checkStringValue(messageID, msg.fm.prm.MessageEffectID); err == nil {
		msg.fm.prm.MessageEffectID = messageID
	}
	return err
}

func (msg *Message) WriteEntities(entities []*MessageEntity) error {
	var err error
	if err = checkEntities(entities, msg.fm.prm.Entities); err == nil {
		msg.fm.prm.Entities = entities
		msg.fm.prm.CaptionEntities = entities
		msg.fm.prm.GiftEntities = entities
	}
	return err
}

func (msg *Message) WriteLinkPreviewOptions(lpo *LinkPreviewOptions) error {
	var err error
	if lpo != nil {
		if msg.fm.prm.LinkPreviewOptions == nil {
			msg.fm.prm.LinkPreviewOptions = lpo
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteShowCaptionAboveMedia() error {
	var err error
	if !msg.fm.prm.ShowCaptionAboveMedia {
		msg.fm.prm.ShowCaptionAboveMedia = true
	} else {
		err = code10()
	}
	return err
}

func (msg *Message) WriteReplyParameters(reply *ReplyParameters) error {
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
			if msg.fm.prm.ReplyParameters == nil {
				msg.fm.prm.ReplyParameters = reply
			} else {
				err = code10()
			}
		}

	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteAllowPaidBroadcast() error {
	var err error
	if !msg.fm.prm.AllowPaidBroadcast {
		msg.fm.prm.AllowPaidBroadcast = true
	} else {
		err = code10()
	}
	return err
}

func (msg *Message) WriteStarCount(amount int) error {
	var err error
	if err = checkIntegerValue(amount, msg.fm.prm.StarCount); err == nil {
		msg.fm.prm.StarCount = amount
	}
	return err
}

func (msg *Message) WritePayload(payload string) error {
	var err error
	if err = checkStringValue(payload, msg.fm.prm.Payload); err == nil {
		msg.fm.prm.Payload = payload
	}
	return err
}

func makeMap() map[rune]*emojich {
	emojiMap := make(map[rune]*emojich, len(Emojis))
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

func (msg *Message) WriteEmoji(emoji string) error {
	var err error
	emojiMap := makeMap()
	em := []rune(emoji)
	if len(em) > 0 {
		if _, ok := emojiMap[em[0]]; !ok {
			err = code20()
		} else {
			if err = isArrEmojiOK(em, emojiMap); err == nil {
				if emojiMap[em[0]].amount < emojiMap[em[0]].permissibleQuantity {
					if msg.fm.prm.Emoji == "" {
						msg.fm.prm.Emoji = emoji
						if !msg.fm.notchange {
							msg.fm.method = methods.Dice
							msg.fm.tgr = new(MessageResponse)
							msg.fm.httpMethod = http.MethodPost
						}
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

func (msg *Message) WriteBusinessConnectionID(connectionID string) error {
	var err error
	if err = checkStringValue(connectionID, msg.fm.prm.BusinessConnectionID); err == nil {
		msg.fm.prm.BusinessConnectionID = connectionID
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
		}
	}
	return err
}

func (inb *inlineKeyboardButton) WriteWebApp(wbapp *WebAppInfo) error {
	var err error
	if (wbapp != nil) && (wbapp.Url != "") {
		if inb.WebApp == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wWebApp] = added
				inb.WebApp = wbapp
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteLoginUrl(logurl *LoginUrl) error {
	var err error
	if logurl != nil {
		if inb.LoginUrl == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wLoginUrl] = added
				inb.LoginUrl = logurl
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
	if inb.SwitchInlineQuery == nil {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wSwitchIn] = added
			inb.SwitchInlineQuery = &sw
		}
	} else {
		err = code10()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryCurrentChat(swcch string) error {
	var err error
	if inb.SwitchInlineQueryCurrentChat == nil {
		if err = inb.checkOthers(); err == nil {
			inb.storage[wSwitchInQuery] = added
			inb.SwitchInlineQueryCurrentChat = &swcch
		}
	} else {
		err = code10()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryChosenChat(sw *SwitchInlineQueryChosenChat) error {
	var err error
	if sw != nil {
		if inb.SwitchInlineQueryChosenChat == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wSwitchInQueryCh] = added
				inb.SwitchInlineQueryChosenChat = sw
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (inb *inlineKeyboardButton) WriteCallbackGame(game *CallbackGame) error {
	var err error
	if game != nil {
		if inb.CallbackGame == nil {
			if err = inb.checkOthers(); err == nil {
				inb.storage[wGame] = added
				inb.CallbackGame = game
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
	} else {
		err = code10()
	}

	return err
}

func (rp *reply) WriteResizeKeyboard() error {
	var err error
	if !rp.ResizeKeyboard {
		rp.ResizeKeyboard = true
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteOneTimeKeyboard() error {
	var err error
	if !rp.OneTimeKeyboard {
		rp.OneTimeKeyboard = true
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteInputFieldPlaceholder(placeholder string) error {
	var err error
	if err = checkInputPlaceHolder(placeholder, rp.InputFieldPlaceholder); err == nil {
		rp.InputFieldPlaceholder = placeholder
	}
	return err
}

func (rp *reply) WriteSelective() error {
	var err error
	if !rp.Selective {
		rp.Selective = true
	} else {
		err = code10()
	}
	return err
}

func (rp *reply) WriteRemove() error {
	var err error
	if !rp.Remove {
		rp.Remove = true
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
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestUsers(requs *KeyboardButtonRequestUsers) error {
	var err error
	if requs != nil {
		if rpb.RequestUsers == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqUsers] = added
				rpb.RequestUsers = requs
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestChat(reqch *KeyboardButtonRequestChat) error {
	var err error
	if reqch != nil {
		if rpb.RequestChat == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqChat] = added
				rpb.RequestChat = reqch
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
		}
	} else {
		err = code10()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteRequestPoll(poll *KeyboardButtonPollType) error {
	var err error
	if poll != nil {
		if rpb.RequestPoll == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqPoll] = added
				rpb.RequestPoll = poll
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (rpb *replyKeyboardButton) WriteWebApp(webapp *WebAppInfo) error {
	var err error
	if (webapp != nil) && (webapp.Url != "") {
		if rpb.WebApp == nil {
			if err = rpb.checkOthers(); err == nil {
				rpb.storage[wReqWebApp] = added
				rpb.WebApp = webapp
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
	} else {
		err = code10()
	}
	return err
}

func (frp *forcereply) WriteInputFieldPlaceholder(placeholder string) error {
	var err error
	if err = checkInputPlaceHolder(placeholder, frp.InputFieldPlaceholder); err == nil {
		frp.InputFieldPlaceholder = placeholder
	}
	return err
}

func (frp *forcereply) WriteSelective() error {
	var err error
	if !frp.Selective {
		frp.Selective = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteTitle(title string) error {
	var err error
	t := len([]rune(title))
	if t > 0 && t <= 32 {
		if p.Title == "" {
			p.Title = title
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
	if p.ProviderToken == nil {
		p.ProviderToken = &token
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteCurrency(currency string) error {
	var err error
	if _, ok := Currencies[currency]; ok {
		if (currency == telegramStars && (p.Prices == nil || p.Prices[0].Amount <= 2500)) || currency != telegramStars {
			if p.Currency == "" {
				p.Currency = currency
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

func (p *payment) WritePrices(prices []*LabeledPrice) error {
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
	if p.StartParameter == nil {
		p.StartParameter = &prm
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteProviderData(data string) error {
	var err error
	if err = checkStringValue(data, p.ProviderData); err == nil {
		p.ProviderData = data
	}
	return err
}

func (p *payment) WriteSubscriptionPeriod(period int) error {
	var err error
	if (p.Currency == "" || p.Currency == telegramStars) && (p.Prices == nil || p.Prices[0].Amount <= 2500) && period == Month {
		if p.SubscriptionPeriod == 0 {
			p.SubscriptionPeriod = period
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
	if p.Currency != telegramStars && p.SubscriptionPeriod == 0 {
		if len(amounts) > 0 && len(amounts) <= 4 {
			for i := 0; i < len(amounts) && err == nil; i++ {
				if amounts[i] < 1 {
					err = code5()
				}
			}
			if err == nil {
				if p.SuggestedTipAmounts == nil {
					p.SuggestedTipAmounts = amounts
				} else {
					err = code10()
				}
			}
		} else {
			err = code20()
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
	}
	return err
}

func (p *payment) WritePhotoSize(size int) error {
	var err error
	if err = checkIntegerValue(size, p.PhotoSize); err == nil {
		p.PhotoSize = size
	}
	return err
}

func (p *payment) WritePhotoWidth(width int) error {
	var err error
	if err = checkIntegerValue(width, p.PhotoWidth); err == nil {
		p.PhotoWidth = width
	}
	return err
}

func (p *payment) WritePhotoHeight(height int) error {
	var err error
	if err = checkIntegerValue(height, p.PhotoHeight); err == nil {
		p.PhotoHeight = height
	}
	return err
}

func (p *payment) WriteNeedName() error {
	var err error
	if !p.NeedName {
		p.NeedName = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteNeedPhoneNumber() error {
	var err error
	if !p.NeedPhoneNumber {
		p.NeedPhoneNumber = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteNeedEmail() error {
	var err error
	if !p.NeedEmail {
		p.NeedEmail = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteNeedShippingAddress() error {
	var err error
	if !p.NeedShippingAddress {
		p.NeedShippingAddress = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteSendPhoneNumberToProvider() error {
	var err error
	if !p.SendPhoneNumberToProvider {
		p.SendPhoneNumberToProvider = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteSendEmailToProvider() error {
	var err error
	if !p.SendEmailToProvider {
		p.SendEmailToProvider = true
	} else {
		err = code10()
	}
	return err
}

func (p *payment) WriteIsFlexible() error {
	var err error
	if !p.IsFlexible {
		p.IsFlexible = true
	} else {
		err = code10()
	}
	return err
}

func (g *game) WriteShortName(name string) error {
	var err error
	if err = checkStringValue(name, g.ShortName); err == nil {
		g.ShortName = name
	}
	return err
}

func (wh *webhook) WriteUrl(url string) error {
	var err error
	if wh.Url == nil {
		wh.Url = &url
	} else {
		err = code10()
	}
	return err
}

func (wh *webhook) WriteCertificate(path string) error {
	var err error
	if path != "" {
		if wh.Certificate == "" {
			wh.Certificate = path
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (wh *webhook) WriteIPAddress(ip string) error {
	var err error
	if ip != "" {
		if wh.IpAddress == "" {
			wh.IpAddress = ip
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (wh *webhook) WriteMaxConnections(max int) error {
	var err error
	if max > 0 && max <= 100 {
		if wh.MaxConnections == 0 {
			wh.MaxConnections = max
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (wh *webhook) WriteAllowedUpdates(upds []string) error {
	var err error
	if upds != nil {
		if wh.AllowedUpdates == nil {
			wh.AllowedUpdates = upds
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (wh *webhook) WriteDropPendingUpdates() error {
	var err error
	if !wh.DropPendingUpdates {
		wh.DropPendingUpdates = true
	} else {
		err = code10()
	}
	return err
}

func (wh *webhook) WriteSecretToken(token string) error {
	var err error
	t := len([]rune(token))
	if t > 0 && t <= 256 {
		if wh.SecretToken == "" {
			wh.SecretToken = token
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (g *gift) WriteID(id string) error {
	var err error
	if err = checkStringValue(id, g.ID); err == nil {
		g.ID = id
	}
	return err
}

func (g *gift) WritePayForUpgrade() error {
	var err error
	if !g.PayForUpgrade {
		g.PayForUpgrade = true
	} else {
		err = code10()
	}
	return err
}

func (msg *Message) WriteString(str string) error {
	var err error
	if err = checkStringValue(str, msg.fm.prm.Text); err == nil {
		msg.fm.prm.Text = str
		msg.fm.prm.Caption = str
	}
	return err
}

func (msg *Message) WriteAction(action string) error {
	var err error
	var found bool
	for i := 0; (i < len(Actions)) && (!found); i++ {
		if Actions[i] == action {
			found = true
		} else if i+1 == len(Actions) {
			err = code20()
		}
	}
	if err == nil {
		if msg.fm.prm.Action == "" {
			msg.fm.prm.Action = action
			if !msg.fm.notchange {
				msg.fm.method = methods.ChatAction
				msg.fm.tgr = new(SimpleResponse)
				msg.fm.httpMethod = http.MethodGet
			}
		} else {
			err = code10()
		}
	}
	return err
}

func (msg *Message) WriteUserID(userID int) error {
	var err error
	if err = checkIntegerValue(userID, msg.fm.prm.UserID); err == nil {
		msg.fm.prm.UserID = userID
	}
	return err
}

func (msg *Message) WritePhoto(ph IPhoto) error {
	var err error
	if p, ok := ph.(*photo); ok {
		if p.Photo != "" {
			if compatibilityCheck(msg, constPhoto) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if p.gottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange && msg.fm.method != methods.PaidMedia {
						if msg.fm.mh.i == 0 {
							msg.fm.method = methods.Photo
							msg.fm.tgr = new(MessageResponse)
							msg.fm.httpMethod = http.MethodPost
						} else {
							msg.fm.method = methods.MediaGroup
							msg.fm.tgr = new(MediaGroupResponse)
							msg.fm.httpMethod = http.MethodPost
						}
					}
					msg.fm.mh.storage[msg.fm.mh.i] = p
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteAudio(ad IAudio) error {
	var err error
	if a, ok := ad.(*audio); ok {
		if a.Audio != "" {
			if compatibilityCheck(msg, constAudio) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if a.audioGottenFrom == Storage || a.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange {
						if msg.fm.mh.i == 0 {
							msg.fm.method = methods.Audio
							msg.fm.tgr = new(MessageResponse)
							msg.fm.httpMethod = http.MethodPost
						} else {
							msg.fm.method = methods.MediaGroup
							msg.fm.tgr = new(MediaGroupResponse)
							msg.fm.httpMethod = http.MethodPost
						}
					}
					msg.fm.mh.storage[msg.fm.mh.i] = a
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteVideo(vd IVideo) error {
	var err error
	if v, ok := vd.(*video); ok {
		if v.Video != "" {
			if compatibilityCheck(msg, constVideo) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if v.videoGottenFrom == Storage || v.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange && msg.fm.method != methods.PaidMedia {
						if msg.fm.mh.i == 0 {
							msg.fm.method = methods.Video
							msg.fm.tgr = new(MessageResponse)
							v.StartTimestamp = &msg.fm.prm.VideoStartTimestamp
							msg.fm.httpMethod = http.MethodPost
						} else {
							msg.fm.method = methods.MediaGroup
							msg.fm.tgr = new(MediaGroupResponse)
							msg.fm.httpMethod = http.MethodPost
						}
					}
					msg.fm.mh.storage[msg.fm.mh.i] = v
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteDocument(doc IDocument) error {
	var err error
	if d, ok := doc.(*document); ok {
		if d.Document != "" {
			if compatibilityCheck(msg, constDoc) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if d.documentGottenFrom == Storage || d.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange {
						if msg.fm.mh.i == 0 {
							msg.fm.method = methods.Document
							msg.fm.tgr = new(MessageResponse)
							msg.fm.httpMethod = http.MethodPost
						} else {
							msg.fm.method = methods.MediaGroup
							msg.fm.tgr = new(MediaGroupResponse)
							msg.fm.httpMethod = http.MethodPost
						}
					}
					msg.fm.mh.storage[msg.fm.mh.i] = d
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteVideoNote(vdn IVideoNote) error {
	var err error
	if v, ok := vdn.(*videonote); ok {
		if v.VideoNote != "" {
			if compatibilityCheck(msg, constVideoNote) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if v.videoGottenFrom == Storage || v.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange {
						msg.fm.method = methods.VideoNote
						msg.fm.tgr = new(MessageResponse)
						msg.fm.httpMethod = http.MethodPost
					}
					msg.fm.mh.storage[msg.fm.mh.i] = v
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteVoice(vc IVoice) error {
	var err error
	if v, ok := vc.(*voice); ok {
		if v.Voice != "" {
			if compatibilityCheck(msg, constVoice) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if v.gottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange {
						msg.fm.method = methods.Voice
						msg.fm.tgr = new(MessageResponse)
						msg.fm.httpMethod = http.MethodPost
					}
					msg.fm.mh.storage[msg.fm.mh.i] = v
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteAnimation(anim IAnimation) error {
	var err error
	if a, ok := anim.(*animation); ok {
		if a.Animation != "" {
			if compatibilityCheck(msg, constAnim) {
				if msg.fm.mh.i+1 <= len(msg.fm.mh.storage) {
					if a.animationGottenFrom == Storage || a.thumbnailGottenFrom == Storage {
						msg.fm.mh.atLeastOnce = true
					}
					if !msg.fm.notchange {
						msg.fm.method = methods.Animation
						msg.fm.tgr = new(MessageResponse)
						msg.fm.httpMethod = http.MethodPost
					}
					msg.fm.mh.storage[msg.fm.mh.i] = a
					msg.fm.mh.i++
					msg.fm.mh.amount++
				} else {
					err = code3()
				}
			} else {
				err = code54()
			}
		} else {
			err = code21()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteSticker(st ISticker) error {
	var err error
	if s, ok := st.(*sticker); ok {
		if msg.fm.sticker.storage[msg.fm.sticker.i] == nil {
			if ((s.Sticker != "") || (s.SetName != "")) || (s.GiftID != "") {
				if s.stickerGottenFrom == Storage {
					msg.fm.sticker.atLeastOnce = true
				}
				msg.fm.sticker.storage[msg.fm.sticker.i] = s
				msg.fm.sticker.i++
				msg.fm.sticker.amount++
				if !msg.fm.notchange {
					if s.GiftID != "" {
						msg.fm.method = methods.Gift
					} else {
						msg.fm.method = methods.Sticker
					}
				}
				msg.fm.tgr = new(MessageResponse)
				msg.fm.httpMethod = http.MethodPost
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteLocation(loc ILocation) error {
	var err error
	if l, ok := loc.(*location); ok {
		if msg.fm.loc == nil {
			if (l.Latitude != 0) && (l.Longitude) != 0 {
				msg.fm.loc = l
				if !msg.fm.notchange {
					if msg.fm.loc.Title == "" {
						msg.fm.method = methods.Location
					} else {
						msg.fm.method = methods.Venue
					}
				}
				msg.fm.httpMethod = http.MethodPost
				msg.fm.tgr = new(MessageResponse)
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteContact(con IContact) error {
	var err error
	if c, ok := con.(*contact); ok {
		if msg.fm.con == nil {
			if (c.PhoneNumber != "") && (c.FirstName != "") {
				msg.fm.con = c
				if !msg.fm.notchange {
					msg.fm.method = methods.Contact
					msg.fm.tgr = new(MessageResponse)
					msg.fm.httpMethod = http.MethodPost
				}
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteGift(gg IGift) error {
	var err error
	if g, ok := gg.(*gift); ok {
		if msg.fm.sticker.storage[msg.fm.sticker.i] == nil {
			if g.ID != "" {
				msg.fm.gift = g
				if !msg.fm.notchange {
					msg.fm.method = methods.Gift
					msg.fm.tgr = new(SimpleResponse)
				}
				msg.fm.httpMethod = http.MethodPost
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WritePoll(p IPoll) error {
	var err error
	if pp, ok := p.(*poll); ok {
		if msg.fm.poll == nil {
			if (pp.Question != "") && (pp.Options != nil) {
				msg.fm.poll = pp
				if !msg.fm.notchange {
					msg.fm.method = methods.Poll
					msg.fm.tgr = new(MessageResponse)
					msg.fm.httpMethod = http.MethodPost
				}
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteGame(gamename string) error {
	var err error
	msg.fm.game = new(game)
	if err = checkStringValue(gamename, msg.fm.game.ShortName); err == nil {
		msg.fm.game.ShortName = gamename
		if !msg.fm.notchange {
			msg.fm.method = methods.Game
			msg.fm.tgr = new(MessageResponse)
			msg.fm.httpMethod = http.MethodPost
		}
	}
	return err
}

func (msg *Message) WritePayment(pay IPayment) error {
	var err error
	if p, ok := pay.(*payment); ok {
		if msg.fm.payment == nil {
			if ((p.Title != "") && (p.Description != "") && (p.Payload != "") && (p.Currency != "") && (p.Prices != nil)) ||
				(p.ShippingID != "" && p.OK != nil) && (*p.OK && p.ShippingOptions != nil || !*p.OK && p.ErrorMessage != "") {
				msg.fm.payment = p
				if !msg.fm.notchange {
					msg.fm.method = methods.Invoice
					msg.fm.tgr = new(MessageResponse)
					msg.fm.httpMethod = http.MethodPost
				}
			} else {
				err = code21()
			}
		} else {
			err = code10()
		}
	} else {
		err = code20()
	}
	return err
}

func (msg *Message) WriteKeyboard(kb IKeyboard) error {
	var err error
	if k, ok := kb.(*keyboard); ok {
		if err = isNill(k.Keyboard); err != nil {
			if err = k.Keyboard.isOK(); err == nil {
				if msg.fm.kb == nil {
					msg.fm.kb = k
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
