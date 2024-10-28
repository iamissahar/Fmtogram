package formatter

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	interfacePhoto    string = "IPhoto"
	interfaceVideo    string = "IVideo"
	interfaceAudio    string = "IAudio"
	interfaceDocument string = "IDocument"
	interfaceInf      string = "MSG Information"
	interfaceChat     string = "IChat"
	inKB              string = "Inline Keyboard"
	replyKB           string = "Reply Keyboard"
	button            string = "Button"
	inbtn             string = "Inline Button"
	rpbtn             string = "Reply Button"
	checkString       int    = -1
	checkArray        int    = -2
	checkBool         int    = -3
	checkInt          int    = -4
)

func (ph *photo) writePhoto(photo, object string) {
	if !isItEmply(ph, checkString, ph.Photo) {
		logs.DataIsntEmply(interfacePhoto, "Photo", ph.Photo)
	}
	ph.Photo, ph.Media = photo, photo
	logs.DataWrittenSuccessfully(interfacePhoto, object)
}

func (ph *photo) WritePhotoStorage(photo string) {
	ph.GottenFrom = Storage
	ph.writePhoto(photo, "Photo From Storage")
}

func (ph *photo) WritePhotoTelegram(photo string) {
	ph.GottenFrom = Telegram
	ph.writePhoto(photo, "Photo From Telegram")
}

func (ph *photo) WritePhotoInternet(photo string) {
	ph.GottenFrom = Internet
	ph.writePhoto(photo, "Photo From The Internet")
}

func (ph *photo) WriteCaption(caption string) {
	if !isItEmply(ph, checkString, ph.Caption) {
		logs.DataIsntEmply(interfacePhoto, "Caption", ph.Caption)
	}
	ph.Caption = caption
	logs.DataWrittenSuccessfully(interfacePhoto, "Caption")
}

func (ph *photo) WriteParseMode(parsemode string) {
	if !isItEmply(ph, checkString, ph.ParseMode) {
		logs.DataIsntEmply(interfacePhoto, "Parse Mode", ph.ParseMode)
	}
	ph.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfacePhoto, "Parse Mode")
}

func (ph *photo) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(ph, checkArray, ph.CaptionEntities) {
		logs.DataIsntEmply(interfacePhoto, "Caption Entities", ph.CaptionEntities)
	}
	ph.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfacePhoto, "Caption Entities")
}

func (ph *photo) WriteShowCaptionAboveMedia() {
	if !isItEmply(ph, checkBool, ph.ShowCaptionAboveMedia) {
		logs.DataIsntEmply(interfacePhoto, "Show Caption Above Media", ph.ShowCaptionAboveMedia)
	}
	ph.ShowCaptionAboveMedia = true
	logs.SettedParam("Show Caption Above Media", interfacePhoto, ph.ShowCaptionAboveMedia)
}

func (ph *photo) WriteHasSpoiler() {
	if !isItEmply(ph, checkBool, ph.HasSpoiler) {
		logs.DataIsntEmply(interfacePhoto, "Has Spoiler", ph.HasSpoiler)
	}
	ph.HasSpoiler = true
	logs.SettedParam("Has Spoiler", interfacePhoto, ph.HasSpoiler)
}

func (ph *photo) GetResponse() [4]types.PhotoSize {
	return ph.response
}

func (vd *video) writeVideo(video, object string) {
	if !isItEmply(vd, checkString, vd.Video) {
		logs.DataIsntEmply(interfaceVideo, "Video", vd.Video)
	}
	vd.Video, vd.Media = video, video
	logs.DataWrittenSuccessfully(interfaceVideo, object)
}

func (vd *video) WriteVideoStorage(video string) {
	vd.VideoGottenFrom = Storage
	vd.writeVideo(video, "Video From Storage")
}

func (vd *video) WriteVideoTelegram(video string) {
	vd.VideoGottenFrom = Storage
	vd.writeVideo(video, "Video From Telegram")
}

func (vd *video) WriteVideoInternet(video string) {
	vd.VideoGottenFrom = Storage
	vd.writeVideo(video, "Video From the Internet")
}

func (vd *video) writeThumbnail(thumbnail, object string) {
	if !isItEmply(vd, checkString, vd.Thumbnail) {
		logs.DataIsntEmply("Video", "Thumbnail", vd.Thumbnail)
	}
	vd.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(interfaceVideo, object)
}

func (vd *video) WriteThumbnailStorage(thumbnail string) {
	vd.ThumbnailGottenFrom = Storage
	vd.writeThumbnail(thumbnail, "Thumbnail From Storage")
}

func (vd *video) WriteThumbnailTelegram(thumbnail string) {
	vd.ThumbnailGottenFrom = Telegram
	vd.writeThumbnail(thumbnail, "Thumbnail From Telegram")
}

func (vd *video) WriteThumbnailInternet(thumbnail string) {
	vd.ThumbnailGottenFrom = Internet
	vd.writeThumbnail(thumbnail, "Thumbnail From the Internet")
}

func (vd *video) WriteCaption(caption string) {
	if !isItEmply(vd, checkString, vd.Caption) {
		logs.DataIsntEmply("Video", "Caption", vd.Caption)
	}
	vd.Caption = caption
	logs.DataWrittenSuccessfully(interfaceVideo, "Caption")
}

func (vd *video) WriteParseMode(parsemode string) {
	if !isItEmply(vd, checkString, vd.ParseMode) {
		logs.DataIsntEmply("Video", "Parse Mode", vd.ParseMode)
	}
	vd.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceVideo, "Parse Mode")
}

func (vd *video) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(vd, checkArray, vd.CaptionEntities) {
		logs.DataIsntEmply("Video", "Caption Entities", vd.CaptionEntities)
	}
	vd.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfaceVideo, "Caption Entities")
}

func (vd *video) WriteShowCaptionAboveMedia() {
	if !isItEmply(vd, checkBool, vd.ShowCaptionAboveMedia) {
		logs.DataIsntEmply("Video", "Show Caption Above Media", vd.ShowCaptionAboveMedia)
	}
	vd.ShowCaptionAboveMedia = true
	logs.SettedParam("Show Caption Above Media", "Video", vd.ShowCaptionAboveMedia)
}

func (vd *video) WriteWidth(width int) {
	if !isItEmply(vd, checkInt, vd.Width) {
		logs.DataIsntEmply("Video", "Width", vd.Width)
	}
	vd.Width = width
	logs.DataWrittenSuccessfully(interfaceVideo, "Width")
}

func (vd *video) WriteHeight(height int) {
	if !isItEmply(vd, checkInt, vd.Height) {
		logs.DataIsntEmply("Video", "Height", vd.Height)
	}
	vd.Height = height
	logs.DataWrittenSuccessfully(interfaceVideo, "Height")
}

func (vd *video) WriteDuration(duration int) {
	if !isItEmply(vd, checkInt, vd.Duration) {
		logs.DataIsntEmply("Video", "Duration", vd.Duration)
	}
	vd.Duration = duration
	logs.DataWrittenSuccessfully(interfaceVideo, "Duration")
}

func (vd *video) WriteSupportsStreaming() {
	if !isItEmply(vd, checkBool, vd.SupportsStreaming) {
		logs.DataIsntEmply("Video", "Supports Streaming", vd.SupportsStreaming)
	}
	vd.SupportsStreaming = true
	logs.SettedParam("Supports Streaming", "Video", vd.SupportsStreaming)
}

func (vd *video) WriteHasSpoiler() {
	if !isItEmply(vd, checkBool, vd.HasSpoiler) {
		logs.DataIsntEmply("Video", "Has Spoiler", vd.HasSpoiler)
	}
	vd.HasSpoiler = true
	logs.SettedParam("Has Spoiler", "Video", vd.SupportsStreaming)
}

func (vd *video) GetResponse() types.Video {
	return vd.response
}

func (ad *audio) writeAudio(audio, object string) {
	if !isItEmply(ad, checkString, ad.Audio) {
		logs.DataIsntEmply(interfaceAudio, "Audio", ad.Audio)
	}
	ad.Audio, ad.Media = audio, audio
	logs.DataWrittenSuccessfully(interfaceAudio, object)
}

func (ad *audio) WriteAudioStorage(audio string) {
	ad.AudioGottenFrom = Storage
	ad.writeAudio(audio, "Audio From Storage")
}

func (ad *audio) WriteAudioTelegram(audio string) {
	ad.AudioGottenFrom = Telegram
	ad.writeAudio(audio, "Audio From Telegram")
}

func (ad *audio) WriteAudioInternet(audio string) {
	ad.AudioGottenFrom = Internet
	ad.writeAudio(audio, "Audio From the Internet")
}

func (ad *audio) writeThumbnail(thumbnail, object string) {
	if !isItEmply(ad, checkString, ad.Thumbnail) {
		logs.DataIsntEmply("Audio", "Thumbnail", ad.Thumbnail)
	}
	ad.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(interfaceAudio, object)
}

func (ad *audio) WriteThumbnailStorage(thumbnail string) {
	ad.ThumbnailGottenFrom = Storage
	ad.writeThumbnail(thumbnail, "Thumbnail From Storage")
}

func (ad *audio) WriteThumbnailTelegram(thumbnail string) {
	ad.ThumbnailGottenFrom = Telegram
	ad.writeThumbnail(thumbnail, "Thumbnail From Telegram")
}

func (ad *audio) WriteThumbnailInternet(thumbnail string) {
	ad.ThumbnailGottenFrom = Internet
	ad.writeThumbnail(thumbnail, "Thumbnail From the Internet")
}

func (ad *audio) WriteCaption(caption string) {
	if !isItEmply(ad, checkString, ad.Caption) {
		logs.DataIsntEmply("Audio", "Caption", ad.Caption)
	}
	ad.Caption = caption
	logs.DataWrittenSuccessfully(interfaceAudio, "Caption")
}

func (ad *audio) WriteParseMode(parsemode string) {
	if !isItEmply(ad, checkString, ad.ParseMode) {
		logs.DataIsntEmply("Audio", "Parse Mode", ad.ParseMode)
	}
	ad.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceAudio, "Parse Mode")
}

func (ad *audio) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(ad, checkArray, ad.CaptionEntities) {
		logs.DataIsntEmply("Audio", "Caption Entities", ad.CaptionEntities)
	}
	ad.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfaceAudio, "Caption Entities")
}

func (ad *audio) WriteDuration(duration int) {
	if !isItEmply(ad, checkInt, ad.Duration) {
		logs.DataIsntEmply("Audio", "Duration", ad.Duration)
	}
	ad.Duration = duration
	logs.DataWrittenSuccessfully(interfaceAudio, "Duration")
}

func (ad *audio) WritePerformer(performer string) {
	if !isItEmply(ad, checkString, ad.Performer) {
		logs.DataIsntEmply("Audio", "Performer", ad.Performer)
	}
	ad.Performer = performer
	logs.DataWrittenSuccessfully(interfaceAudio, "Performer")
}

func (ad *audio) WriteTitle(title string) {
	if !isItEmply(ad, checkString, ad.Title) {
		logs.DataIsntEmply("Audio", "Title", ad.Title)
	}
	ad.Title = title
	logs.DataWrittenSuccessfully(interfaceAudio, "Title")
}

func (ad *audio) GetResponse() types.Audio {
	return ad.response
}

func (dc *document) writeDocument(document, object string) {
	if !isItEmply(dc, checkString, dc.Document) {
		logs.DataIsntEmply("Document", "Document", dc.Document)
	}
	dc.Document = document
	logs.DataWrittenSuccessfully(interfaceDocument, object)
}

func (dc *document) WriteDocumentStorage(document string) {
	dc.DocumentGottenFrom = Storage
	dc.writeDocument(document, "Document From Storage")
}

func (dc *document) WriteDocumentTelegram(document string) {
	dc.DocumentGottenFrom = Telegram
	dc.writeDocument(document, "Document From Telegram")
}

func (dc *document) WriteDocumentInternet(document string) {
	dc.DocumentGottenFrom = Internet
	dc.writeDocument(document, "Document From the Internet")
}

func (dc *document) writeThumbnail(thumbnail, object string) {
	if !isItEmply(dc, checkString, dc.Thumbnail) {
		logs.DataIsntEmply("Document", "Thumbnail", dc.Thumbnail)
	}
	dc.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(interfaceDocument, object)
}

func (dc *document) WriteThumbnailStorage(thumbnail string) {
	dc.ThumbnailGottenFrom = Storage
	dc.writeThumbnail(thumbnail, "Thumbnail From Storage")
}

func (dc *document) WriteThumbnailTelegram(thumbnail string) {
	dc.ThumbnailGottenFrom = Telegram
	dc.writeThumbnail(thumbnail, "Thumbnail From Telegram")
}

func (dc *document) WriteThumbnailInternet(thumbnail string) {
	dc.ThumbnailGottenFrom = Internet
	dc.writeThumbnail(thumbnail, "Thumbnail From the Internet")
}

func (dc *document) WriteCaption(caption string) {
	if !isItEmply(dc, checkString, dc.Caption) {
		logs.DataIsntEmply("Document", "Caption", dc.Caption)
	}
	dc.Caption = caption
	logs.DataWrittenSuccessfully(interfaceDocument, "Caption")
}

func (dc *document) WriteParseMode(parsemode string) {
	if !isItEmply(dc, checkString, dc.ParseMode) {
		logs.DataIsntEmply("Document", "Parse Mode", dc.ParseMode)
	}
	dc.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceDocument, "Parse Mode")
}

func (dc *document) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(dc, checkArray, dc.CaptionEntities) {
		logs.DataIsntEmply("Document", "Caption Entities", dc.CaptionEntities)
	}
	dc.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfaceDocument, "Caption Entities")
}

func (dc *document) WriteDisableContentTypeDetection() {
	if !isItEmply(dc, checkArray, dc.DisableContentTypeDetection) {
		logs.DataIsntEmply("Document", "Disable Content Type Detection", dc.DisableContentTypeDetection)
	}
	dc.DisableContentTypeDetection = true
	logs.SettedParam("Disable Content Type Detection", "Document", dc.DisableContentTypeDetection)
}

func (dc *document) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(dc, checkInt, dc.DocumentGottenFrom) {
		logs.DataIsntEmply(interfaceDocument, "Gotten From", dc.DocumentGottenFrom)
	}
	dc.DocumentGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceDocument, "Gotten From")
}

func (dc *document) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(dc, checkInt, dc.ThumbnailGottenFrom) {
		logs.DataIsntEmply(interfaceDocument, "Thumbnail Gotten From", dc.ThumbnailGottenFrom)
	}
	dc.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceDocument, "Thumbnail Gotten From")
}

func (dc *document) GetResponse() types.Document {
	return dc.response
}

func (an *animation) writeAnimation(animation, object string) {
	if !isItEmply(an, checkString, an.Animation) {
		logs.DataIsntEmply("WriteAnimation{Storage/Telegram/URL}()", "IAnimation", an.Animation)
	}
	an.Animation = animation
	logs.DataWrittenSuccessfully(interfaceAnimation, object)
}

func (an *animation) WriteAnimationStorage(animation string) {
	if !isItEmply(an, checkInt, an.Animation) {
		logs.DataIsntEmply(interfaceDocument, "Animation", an.Animation)
	}
	an.AnimationGottenFrom = Storage
	an.writeAnimation(animation, "Thumbnail From Storage")
}

func (inf *information) WriteString(text string) {
	if inf.Text != "" {
		logs.DataIsntEmply(interfaceInf, "Text", inf.Text)
	}
	inf.Text = text
	logs.DataWrittenSuccessfully(interfaceInf, "Text")
}

func (inf *information) WriteParseMode(parsemode string) {
	if inf.ParseMode != "" {
		logs.DataIsntEmply(interfaceInf, "Parse Mode", inf.ParseMode)
	}
	inf.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceInf, "Parse Mode")
}

func (inf *information) WriteMessageThreadID(messageID int) {
	if inf.MessageThreadID != 0 {
		logs.DataIsntEmply(interfaceInf, "Message Thread ID", inf.MessageThreadID)
	}
	inf.MessageThreadID = messageID
	logs.DataWrittenSuccessfully(interfaceInf, "Message Thread ID")
}

func (inf *information) WriteDisableNotification() {
	if inf.DisableNotification {
		logs.DataIsntEmply(interfaceInf, "Disable Notification", inf.DisableNotification)
	}
	inf.DisableNotification = true
	logs.SettedParam("Disable Notification", interfaceInf, inf.DisableNotification)
}

func (inf *information) WriteProtectContent() {
	if inf.ProtectContent {
		logs.DataIsntEmply(interfaceInf, "Protect Content", inf.ProtectContent)
	}
	inf.ProtectContent = true
	logs.SettedParam("Protect Content", interfaceInf, inf.ProtectContent)
}

func (inf *information) WriteMessageEffectID(messageID string) {
	if inf.MessageEffectID != "" {
		logs.DataIsntEmply(interfaceInf, "Message Effect ID", inf.MessageEffectID)
	}
	inf.MessageEffectID = messageID
	logs.DataWrittenSuccessfully(interfaceInf, "Message Effect ID")
}

func (inf *information) WriteEntities(entities []*types.MessageEntity) {
	if len(inf.Entities) != 0 {
		logs.DataIsntEmply(interfaceInf, "Entities", inf.Entities)
	}
	inf.Entities = entities
	logs.DataWrittenSuccessfully(interfaceInf, "Entities")
}

func (inf *information) WriteLinkPreviewOptions(lpo *types.LinkPreviewOptions) {
	if inf.LinkPreviewOptions != nil {
		logs.DataIsntEmply(interfaceInf, "Link Preview Options", inf.LinkPreviewOptions)
	}
	inf.LinkPreviewOptions = lpo
	logs.DataWrittenSuccessfully(interfaceInf, "Link Preview Options")
}

func (inf *information) WriteMessageID(messageID int) {
	if inf.MessageID != 0 {
		logs.DataIsntEmply(interfaceInf, "Message ID", inf.MessageID)
	}
	inf.MessageID = messageID
	logs.DataWrittenSuccessfully(interfaceInf, "Message ID")
}

func (inf *information) WriteMessageIDs(messageIDs []int) error {
	var err error
	if inf.MessageIDs != nil {
		logs.DataIsntEmply(interfaceInf, "Message IDs", inf.MessageIDs)
	}
	inf.MessageIDs = messageIDs
	logs.DataWrittenSuccessfully(interfaceInf, "Message IDs")
	return err
}

func (inf *information) WriteCaption(caption string) {
	if inf.Caption != "" {
		logs.DataIsntEmply(interfaceInf, "Caption", inf.Caption)
	}
	inf.Caption = caption
	logs.DataWrittenSuccessfully(interfaceInf, "Caption")
}

func (inf *information) WriteShowCaptionAboveMedia() {
	if inf.ShowCaptionAboveMedia {
		logs.DataIsntEmply(interfaceInf, "Show Caption Above Media", inf.ShowCaptionAboveMedia)
	}
	inf.ShowCaptionAboveMedia = true
	logs.DataWrittenSuccessfully(interfaceInf, "Show Caption Above Media")
}

func (inf *information) WriteReplyParameters(reply *types.ReplyParameters) {
	if inf.ReplyParameters != nil {
		logs.DataIsntEmply(interfaceInf, "Reply Parameters", inf.ReplyParameters)
	}
	inf.ReplyParameters = reply
	logs.DataWrittenSuccessfully(interfaceInf, "Reply Parameters")
}

func (inf *information) GetResponse() types.User {
	return inf.response
}

func (inf *information) GetMessageIDs() []int {
	return inf.responseMessageIDs
}

func (ch *chat) WriteChatID(chatID int) {
	if ch.ID != nil {
		logs.DataIsntEmply(interfaceChat, "Chat ID", ch.ID)
	}
	ch.ID = chatID
	logs.DataWrittenSuccessfully(interfaceChat, "Chat ID")
}

func (ch *chat) WriteChatName(chatname string) {
	if ch.ID != nil {
		logs.DataIsntEmply(interfaceChat, "Chat Name", ch.ID)
	}
	ch.ID = fmt.Sprint("@", chatname)
	logs.DataWrittenSuccessfully(interfaceChat, "Chat Name")
}

func (ch *chat) WriteBusinessConnectionID(connectionID string) {
	if ch.BusinessConnectionID != "" {
		logs.DataIsntEmply(interfaceChat, "Business Connection ID", ch.BusinessConnectionID)
	}
	ch.BusinessConnectionID = connectionID
	logs.DataWrittenSuccessfully(interfaceChat, "Business Connection ID")
}

func (ch *chat) WriteFromChatID(chatID int) {
	if ch.FromChatID != nil {
		logs.DataIsntEmply(interfaceChat, "From Chat ID", ch.FromChatID)
	}
	ch.FromChatID = chatID
	logs.DataWrittenSuccessfully(interfaceChat, "From Chat ID")
}

func (ch *chat) WriteFromChatName(chatname string) {
	if ch.FromChatID != nil {
		logs.DataIsntEmply(interfaceChat, "From Chat ID", ch.FromChatID)
	}
	ch.FromChatID = fmt.Sprint("@", chatname)
	logs.DataWrittenSuccessfully(interfaceChat, "From Chat ID")
}

func (in *inline) Set(plan []int) {
	in.Keyboard = new(inlineKeyboard)
	in.Keyboard.InlineKeyboard = make([][]*inlineKeyboardButton, len(plan))
	for i := range in.Keyboard.InlineKeyboard {
		in.Keyboard.InlineKeyboard[i] = make([]*inlineKeyboardButton, plan[i])
	}
}

func (ch *chat) GetResponse() types.Chat {
	return ch.response
}

func (in *inline) NewButton(line, pos int) (IInlineButton, error) {
	var err error

	if (line >= 0) && (pos >= 0) && len(in.Keyboard.InlineKeyboard) > line && len(in.Keyboard.InlineKeyboard[line]) > pos {

		if in.Keyboard.InlineKeyboard[line][pos] != nil {
			logs.DataIsntEmply(inKB, fmt.Sprintf("%s line: %d, position: %d", button, line, pos), in.Keyboard.InlineKeyboard[line][pos])
		}
		in.Keyboard.InlineKeyboard[line][pos] = new(inlineKeyboardButton)

		return in.Keyboard.InlineKeyboard[line][pos], nil

	} else {

		if len(in.Keyboard.InlineKeyboard) > line {
			err = errors.ButtosDoesntFit("line", line)
		} else if len(in.Keyboard.InlineKeyboard[line]) > pos {
			err = errors.ButtosDoesntFit("pos", pos)
		}
	}
	return nil, err
}

func (inb *inlineKeyboardButton) WriteString(text string) {
	if inb.Text != "" {
		logs.DataIsntEmply(inbtn, "Text", inb.Text)
	}
	inb.Text = text
	logs.DataWrittenSuccessfully(inbtn, "Text")
}

func (inb *inlineKeyboardButton) WriteURL(url string) {
	if inb.Url != "" {
		logs.DataIsntEmply(inbtn, "URL", inb.Url)
	}
	if inb.One < 1 {
		inb.Url = url
		inb.One++
		logs.DataWrittenSuccessfully(inbtn, "URL")
	}
}

func (inb *inlineKeyboardButton) WriteCallbackData(text string) {
	if inb.CallbackData != "" {
		logs.DataIsntEmply(inbtn, "Callback Data", inb.CallbackData)
	}
	inb.CallbackData = text
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Callback Data")
}

func (inb *inlineKeyboardButton) WriteWebApp(wbapp *types.WebAppInfo) {
	if inb.WebApp != nil {
		logs.DataIsntEmply(inbtn, "Web App", *inb.WebApp)
	}
	inb.WebApp = wbapp
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Web App")
}

func (inb *inlineKeyboardButton) WriteLoginUrl(logurl *types.LoginUrl) {
	if inb.LoginUrl != nil {
		logs.DataIsntEmply(inbtn, "Login URL", *inb.LoginUrl)
	}
	inb.LoginUrl = logurl
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Login URL")
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQuery(sw string) {
	if inb.SwitchInlineQuery != "" {
		logs.DataIsntEmply(inbtn, "Switch Inline Query", inb.SwitchInlineQuery)
	}
	inb.SwitchInlineQuery = sw
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query")
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryCurrentChat(swcch string) {
	if inb.SwitchInlineQueryCurrentChat != "" {
		logs.DataIsntEmply(inbtn, "Switch Inline Query Current Chat", inb.SwitchInlineQueryCurrentChat)
	}
	inb.SwitchInlineQueryCurrentChat = swcch
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Current Chat")
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat) {
	if inb.SwitchInlineQueryChosenChat != nil {
		logs.DataIsntEmply(inbtn, "Switch Inline Query Chosen Chat", inb.SwitchInlineQueryChosenChat)
	}
	inb.SwitchInlineQueryChosenChat = sw
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Chosen Chat")
}

func (inb *inlineKeyboardButton) WriteCallbackGame(game *types.CallbackGame) {
	if inb.CallbackGame != nil {
		logs.DataIsntEmply(inbtn, "Callback Game", inb.CallbackGame)
	}
	inb.CallbackGame = game
	inb.One++
	logs.DataWrittenSuccessfully(inbtn, "Callback Game")
}

func (inb *inlineKeyboardButton) WritePay() {
	if inb.Pay {
		logs.DataIsntEmply(inbtn, "Pay", inb.Pay)
	}
	inb.Pay = true
	inb.One++
	logs.SettedParam("Pay", inbtn, true)
}

func (rp *reply) Set(plan []int) {
	rp.Keyboard = new(replyKeyboard)
	rp.Keyboard.Keyboard = make([][]*replyKeyboardButton, len(plan))
	for i := range rp.Keyboard.Keyboard {
		rp.Keyboard.Keyboard[i] = make([]*replyKeyboardButton, plan[i])
	}
}

func (rp *reply) NewButton(line, pos int) (IReplyButton, error) {
	var err error

	if (line >= 0) && (pos >= 0) && (len(rp.Keyboard.Keyboard) > line) && (len(rp.Keyboard.Keyboard[line]) > pos) {

		if rp.Keyboard.Keyboard[line][pos] != nil {
			logs.DataIsntEmply(replyKB, fmt.Sprintf("%s line: %d, position: %d", button, line, pos), rp.Keyboard.Keyboard[line][pos])
		}
		rp.Keyboard.Keyboard[line][pos] = new(replyKeyboardButton)

		return rp.Keyboard.Keyboard[line][pos], nil

	} else {
		if len(rp.Keyboard.Keyboard) > line {
			err = errors.ButtosDoesntFit("line", line)
		} else if len(rp.Keyboard.Keyboard[line]) > pos {
			err = errors.ButtosDoesntFit("pos", pos)
		}
	}
	return nil, err
}

func (rpb *replyKeyboardButton) WriteString(text string) {
	if rpb.Text != "" {
		logs.DataIsntEmply(rpbtn, "Text", rpb.Text)
	}
	rpb.Text = text
	logs.DataWrittenSuccessfully(rpbtn, "Text")
}

func (rpb *replyKeyboardButton) WriteRequestUsers(requs *types.KeyboardButtonRequestUsers) {
	if rpb.RequestUsers != nil {
		logs.DataIsntEmply(rpbtn, "Request Users", rpb.RequestUsers)
	}
	rpb.RequestUsers = requs
	logs.DataWrittenSuccessfully(rpbtn, "Request Users")
}

func (rpb *replyKeyboardButton) WriteRequestChat(reqch *types.KeyboardButtonRequestChat) {
	if rpb.RequestChat != nil {
		logs.DataIsntEmply(rpbtn, "Request Chat", rpb.RequestChat)
	}
	rpb.RequestChat = reqch
	logs.DataWrittenSuccessfully(rpbtn, "Request Chat")
}

func (rpb *replyKeyboardButton) WriteRequestContact() {
	if rpb.RequestContact {
		logs.DataIsntEmply(rpbtn, "Request Contact", rpb.RequestContact)
	}
	rpb.RequestContact = true
	logs.SettedParam("Request Contact", rpbtn, true)
}

func (rpb *replyKeyboardButton) WriteRequestLocation() {
	if rpb.RequestLocation {
		logs.DataIsntEmply(rpbtn, "Request Location", rpb.RequestLocation)
	}
	rpb.RequestLocation = true
	logs.SettedParam("Request Location", rpbtn, true)
}

func (rpb *replyKeyboardButton) WriteRequestPoll(poll *types.KeyboardButtonPollType) {
	if rpb.RequestPoll != nil {
		logs.DataIsntEmply(rpbtn, "Request Poll", rpb.RequestPoll)
	}
	rpb.RequestPoll = poll
	logs.DataWrittenSuccessfully(rpbtn, "Request Poll")
}

func (rpb *replyKeyboardButton) WriteWebApp(webapp *types.WebAppInfo) {
	if rpb.WebApp != nil {
		logs.DataIsntEmply(rpbtn, "Web App", rpb.WebApp)
	}
	rpb.WebApp = webapp
	logs.DataWrittenSuccessfully(rpbtn, "Web App")
}
