package formatter

import (
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/types"
)

type handlerMedia interface {
	multipartFields(writer *multipart.Writer, group *[]interface{}, id int, input bool) error
	jsonFileds() (jsbody []byte, err error)
	uniqueConst() (constID int)
}

type handlerKB interface {
	multipartFields(writer *multipart.Writer) error
	jsonFields() (jsbody []byte, err error)
}

type IPhoto interface {
	// Receives a path of a photo. The path can't be an empty string.
	//
	// path - string. Has to be the path to a photo in your pc
	//
	// if the function has been already mentioned, it returns an error with code 20
	WritePhotoStorage(path string) error

	// Recieves a Telegram ID of a photo. The ID can't be an empty string.
	//
	// photoID - string. Has to be given by telegram, and be data from IPhoto.GetResponse()
	//
	// if the function has been already mentioned, it returns an error with code 20
	WritePhotoTelegram(photoID string) error

	// Recieves an URL-link of a photo. The URL can't be an empty string.
	WritePhotoInternet(URL string) error

	// Receives a string that will be a caption of the photo. If you're going to send
	// not only one picture and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface. Caption cannot be an empty string.
	WriteCaption(caption string) error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message.
	// Parsemode can be only types.HTML, types.Markdown and types.MarkdownV2. Cannot be an empty string.
	WriteParseMode(parsemode string) error

	// entities can't be an empty slice.
	WriteCaptionEntities(entities []*types.MessageEntity) error

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia() error

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler() error

	// You can call this function after calling Send(). It returns you an array with some data about the photo you just sent. 4 is the biggest amount of data. Sometimes there might be nil
	GetResponse() [4]types.PhotoSize
}

type IVideo interface {
	// Receives a path of a video
	WriteVideoStorage(path string) error

	// Recieves a Telegram ID of a video
	WriteVideoTelegram(videoID string) error

	// Recieves an URL-link of a video
	WriteVideoInternet(URL string) error

	// Receives a string that will be a caption of the video. If you're going to send
	// one video and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(caption string) error

	WriteCaptionEntities(entities []*types.MessageEntity) error
	WriteDuration(duration int) error

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler() error
	WriteHeight(height int) error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string) error

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia() error
	WriteSupportsStreaming() error
	WriteThumbnailStorage(path string) error
	WriteThumbnailTelegram(thumbnailID string) error
	WriteThumbnailInternet(URL string) error
	WriteWidth(width int) error

	// You can call this function after calling Send(). It returns you a structure with some data about the video you just sent
	GetResponse() types.Video
}

type IAudio interface {
	// Receives a path of an audio
	WriteAudioStorage(path string) error

	// Receives a Telegram ID of an audio
	WriteAudioTelegram(AudioID string) error

	// Recueves an URL-link of an audio
	WriteAudioInternet(URL string) error

	// Receives a string that will be a caption of the audio. If you're going to send
	// one audio and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(caption string) error

	WriteCaptionEntities(entities []*types.MessageEntity) error

	// Duration of the audio in seconds
	WriteDuration(duration int) error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string) error

	WritePerformer(performer string) error
	WriteThumbnailStorage(path string) error
	WriteThumbnailTelegram(thumbnailID string) error
	WriteThumbnailInternet(URL string) error
	WriteTitle(title string) error

	// You can call this function after calling Send(). It returns you a structure with some data about the audio you just sent
	GetResponse() types.Audio
}

type IDocument interface {
	// Receives a path of an audio
	WriteDocumentStorage(path string) error

	// Receives a Telegram ID of an audio
	WriteDocumentTelegram(documentID string) error

	// Recueves an URL-link of an audio
	WriteDocumentInternet(URL string) error

	// Receives a string that will be a caption of the document. If you're going to send
	// one document and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(caption string) error

	WriteCaptionEntities(entities []*types.MessageEntity) error
	WriteDisableContentTypeDetection() error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string) error

	WriteThumbnailStorage(path string) error
	WriteThumbnailTelegram(thumbnailID string) error
	WriteThumbnailInternet(URL string) error

	// You can call this function after calling Send(). It returns you a structure with some data about the document you just sent
	GetResponse() types.Document
}

type IAnimation interface {
	// Receives a path of an animation
	WriteAnimationStorage(path string) error

	// Receives a Telegram ID of an animation
	WriteAnimationTelegram(animationID string) error

	// Receives an URL-link of an animation
	WriteAnimationInternet(URL string) error

	// Duration of sent animation in seconds
	WriteDuration(duration int) error

	// Animation width
	WriteWidth(width int) error

	// Animation height
	WriteHeight(height int) error

	// Receives a path of an animation
	WriteThumbnailStorage(path string) error

	// Receives a Telegram ID of an animation
	WriteThumbnailTelegram(thumbnailID string) error

	// Receives an URL-link of an animation
	WriteThumbnailInternet(URL string) error

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler() error

	// You can call this function after calling Send(). It returns you a structure with some data about the animation you just sent
	GetResponse() types.Animation
}

type IVoice interface {
	// Receives a path of a voice file in formats: .OGG encoded with OPUS, or in .MP3 format, or in .M4A format
	WriteVoiceStorage(path string) error

	// Receives a Telegram ID of a voice
	WriteVoiceTelegram(voiceID string) error

	// Receives an URL-link of a voice
	WriteVoiceInternet(URL string) error

	// Duration of sent voice in seconds
	WriteDuration(duration int) error

	// You can call this function after calling Send(). It returns you a structure with some data about the voice you just sent
	GetResponse() types.Voice
}

type IVideoNote interface {
	// Receives a path of a voice-note file in formats: MPEG4 videos of up to 1 minute long
	WriteVideoNoteStorage(path string) error

	// Receives a Telegram ID of a voice-note
	WriteVideoNoteTelegram(videonoteID string) error

	// Receives an URL-link of a voice-note
	WriteVideoNoteInternet(URL string) error

	// Duration of sent voice-note in seconds
	WriteDuration(duration int) error

	// Video width and height, i.e. diameter of the voice-note
	WriteLength(length int) error

	// Receives a path of a thumbnail of voice-note
	WriteThumbnailStorage(path string) error

	// Receives a Telegram ID of a thumbnail
	WriteThumbnailTelegram(thumbnailID string) error

	// Receives an URL-link of a thumbnail
	WriteThumbnailInternet(URL string) error

	// You can call this function after calling Send(). It returns you a structure with some data about the video-note you just sent
	GetResponse() types.VideoNote
}

type ILocation interface {
	// Latitude of the location
	WriteLatitude(lat float64) error

	// Longitude of the location
	WriteLongitude(long float64) error

	// The radius of uncertainty for the location, measured in meters; 0-1500
	WriteHorizontalAccuracy(horacc float64) error

	// Period in seconds during which the location will be updated (see https://telegram.org/blog/live-locations,
	// should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	WriteLivePeriod(period int) error

	// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	WriteHeading(heading int) error

	// For live locations, a maximum distance for proximity alerts about approaching another chat member,
	// in meters. Must be between 1 and 100000 if specified.
	WriteProximityAlertRadius(proxalrad int) error

	// Name of the venue
	WriteTitle(title string) error

	// Address of the venue
	WriteAddress(address string) error

	// Foursquare identifier of the venue
	WriteFoursquareID(foursquareID string) error

	// Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”)
	WriteFoursquareType(foursquareType string) error

	// Google Places identifier of the venue
	WriteGooglePlaceID(googlePlaceID string) error

	// Google Places type of the venue. (See https://developers.google.com/maps/documentation/places/web-service/supported_types)
	WriteGooglePlaceType(googlePlaceType string) error

	GetResponse() types.Venue
}

type IContact interface {
	// Receives the contact's phone number
	WritePhoneNumber(phone string) error

	// Receives the contact's first name
	WriteFirstName(fname string) error

	// Receives the contact's last name
	WriteLastName(lname string) error

	// Recieves additional data about the contact in the form of a vCard, 0-2048 bytes. What vCard is: https://en.wikipedia.org/wiki/VCard
	WriteVCard(vcard string) error

	GetResponse() types.Contact
}

type IPoll interface {
	// Receives a string that will be the question of a poll, 1-300 characters
	WriteQuestion(question string) error

	WriteQuestionParseMode(parseode string) error

	WriteQuestionEntities(entities []*types.MessageEntity) error

	// Receives a slice of 2-10 answer options
	WriteOptions(options []*types.PollOption) error

	// Call it only if the poll needs to be anonymous
	WriteAnonymous() error

	// Receives a poll type, “quiz” or “regular”
	WriteType(polltype string) error

	//  Call it only if the poll needs to allow multiple answers, ignored for polls in quiz mode
	WriteAllowMultipleAnswers() error

	// 0-based identifier of the correct answer option, required for polls in quiz mode
	WriteCorrectOptionID(optID int) error

	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll,
	// 1-200 characters with at most 2 line feeds after entities parsing
	WriteExplanation(explanation string) error

	WriteExplanationParseMode(parsemode string) error

	WriteExplanationEntities(entities []*types.MessageEntity) error

	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with (formatter.IPoll).WriteCloseDate
	WriteOpenPeriod(period int) error

	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with (formatter.IPoll).WriteOpenPeriod
	WriteCloseDate(time int) error

	// Call it only if the poll needs to be immediately closed. This can be useful for poll preview.
	WriteClosed() error

	GetResponse() types.Poll
}

// type ISticker interface {
// 	// Receives a path of a sticker file in formats: static .WEBP, animated .TGS or video .WEBM
// 	WriteStickerStorage(path string) error

// 	// Receives a Telegram ID of a sticker
// 	WriteStickerTelegram(stickerID string) error

// 	// Receives an URL-link of a sticker
// 	WriteStickerInternet(URL string) error

// 	// 	Emoji associated with the sticker; only for just uploaded stickers
// 	WriteEmoji(emoji string) error
// }

type IParameters interface {
	// Doesn't revieve anything. If the function was called, the message you send to a client will be gotten without a notification
	WriteDisableNotification() error

	// Recieves a slice of special entities that appear in message text, which can be specified instead of WriteParseMode()
	WriteEntities(entities []*types.MessageEntity) error

	// Recieves a link preview generation options for the message
	WriteLinkPreviewOptions(link *types.LinkPreviewOptions) error

	// Recieves an unique identifier of the message effect to be added to the message; for private chats only
	WriteMessageEffectID(effectID string) error

	// Recieves an unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	WriteMessageThreadID(threadID int) error

	// Recieves an unique identifier for the message you want to do somthing
	WriteMessageID(ID int) error

	// Recieves a slice of unique identifiers for the messages you want to do somthing. Only ints (number of message id). The identifiers must be specified in a strictly increasing order.
	WriteMessageIDs(IDs []int) error

	// New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	WriteCaption(caption string) error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string) error

	// Doesn't recieve anything. If the function was called, the text-message will be protected from forwarding and saving
	WriteProtectContent() error

	// The text content of the message
	WriteString(text string) error

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia() error

	// Recieves a structure that has description of the message to reply to
	WriteReplyParameters(replyPar *types.ReplyParameters) error

	// Call it to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	WriteAllowPaidBroadcast() error

	// The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	WriteStarCount(amount int) error

	// Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	WritePayload(payload string) error

	// You can call this function after calling Send(). It returns you a structure with some data about the user you just sent a message to
	GetResponse() types.User

	// Call it after methods.ForwardMessages, methods.CopyMessage and methods.CopyMessages. In any other cases it returns nil
	GetMessageIDs() []int
}

type IChat interface {
	// Reseives an unique identifier of the business connection on behalf of which the message will be sent
	WriteBusinessConnectionID(connID string) error

	// These the functions might be used only if you need to send a message (!NOT RESPONSE!) to a client.
	// There are some information about the chat the request was gotten from, so you do not need it at all.
	//
	// Receives a number (int). Must be the chatID you want to send a message to
	WriteChatID(ID int) error

	// Receives the name of the channel. Must be the name of the channel you want to send a message to. Works only with public channales
	WriteChatName(name string) error

	// // Receives the URL of the chat. Must be the the URL of you want to send a message to
	// WriteChatURL(chatURL string)

	// Receives a number (int). Must be the chatID of the original message that was sent
	WriteFromChatID(ID int) error

	// Receives the name of the chat. Must be the name of the chat of the original message that was sent
	WriteFromChatName(name string) error

	// You can call this function after calling Send(). It returns you a structure with some data about the chat you just sent a message to
	GetResponse() types.Chat
}

// This interface represents a custom keyboard with reply options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
type IReply interface {
	// Receives 2 coordinates. The first one is a line you expect to see the button and the second one is a position on the line you just mentioned.
	// The coordinates mustn't be a random numbers. You have to call function Set() first. Returns an interface of the new created button
	// and might return and error, if the coordinates are wrong
	NewButton(line int, position int) (button IReplyButton, err error)

	// Receives a slice of integers and sets the plan of future reply-keyboard. Be careful with data, because one array cell is considered
	// as row, so, for example: a slice []int{2, 1, 4} means that in the reply-keyboard will be 3 rows and in the first one there'll be 2 buttons, in the next one: just one
	// and in the last one: 4 buttons
	Set(plan []int) error

	// Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	WriteIsPersistent() error

	// Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	WriteResizeKeyboard() error

	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard
	// in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	WriteOneTimeKeyboard() error

	// The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	WriteInputFieldPlaceholder(placeholder string) error

	// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message
	// is a reply to a message in the same chat and forum topic, sender of the original message.
	//
	// Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
	WriteSelective() error
}

// This interface represents one button of IReply interface. WriteString() is requied, and at most one of the rest functions must be used to specify type of the button. For simple text buttons, String can be used instead of this object to specify the button text.
type IReplyButton interface {
	// If specified, pressing the button will open a list of suitable users. Available in private chats only.
	WriteRequestChat(rchat *types.KeyboardButtonRequestChat) error

	// Doesn't receives anything. If it is called, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	WriteRequestContact() error

	// Doesn't receives anything. If it is called, the user's current location will be sent when the button is pressed. Available in private chats only.
	WriteRequestLocation() error

	// Receives a structure with some data. The user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WriteRequestPoll(rpoll *types.KeyboardButtonPollType) error

	// If specified, pressing the button will open a list of suitable users. Available in private chats only.
	WriteRequestUsers(rusers *types.KeyboardButtonRequestUsers) error

	// The text content that will be on the button
	WriteString(text string) error

	// Receives a structure that describes Web App. Will be launched when the button is pressed. Available in private chats only.
	WriteWebApp(webapp *types.WebAppInfo) error
}

// This interfaceinterface represents an inline keyboard that appears right next to the message it belongs to.
type IInline interface {
	// Receives 2 coordinates. The first one is a line you expect to see the button and the second one is a position on the line you just mentioned.
	// The coordinates mustn't be a random numbers. You have to call function Set() first. Returns an interface of the new created button
	// and might return and error, if the coordinates are wrong
	NewButton(line int, position int) (button IInlineButton, err error)

	// Receives a slice of integers and sets the plan of future inline-keyboard. Be careful with data, because one array cell is one array cell is considered
	// as row, so, for example: a slice []int{2, 1, 4} means that in the inline-keyboard will be 3 rows and in the first one there'll be 2 buttons, in the next one: just one
	// and in the last one: 4 buttons
	Set(plan []int) error
}

// This interface represents one button of an IInline interface. WriteString is requied. The rest aren't, but exactly one function must be used to specify type of the button.
type IInlineButton interface {
	// Receives some text string that will be received by the bot after a client pressed on the button
	WriteCallbackData(text string) error

	// Receives the description of the game that will be launched when the user presses the button.
	//
	// This type of button must always be the first button in the first row
	WriteCallbackGame(game *types.CallbackGame) error

	// Recieves an HTTPS URL structure used to automatically authorize the user.
	WriteLoginUrl(logUrl *types.LoginUrl) error

	// Specify True, to send a Pay button. Substrings “⭐” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.
	//
	// This type of button must always be the first button in the first row and can only be used in invoice messages.
	WritePay() error

	// Receives some text string that will be seen by client on the button
	WriteString(text string) error

	// If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	//
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages
	// sent on behalf of a Telegram Business account.
	WriteSwitchInlineQuery(sw string) error

	// If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field.
	// Not supported for messages sent on behalf of a Telegram Business account.
	WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat) error

	// If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages
	// sent on behalf of a Telegram Business account.
	WriteSwitchInlineQueryCurrentChat(sw string) error

	// Receives a URL string that will be received by the bot after a client pressed on the button
	WriteURL(URL string) error

	// Receives the description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	WriteWebApp(webapp *types.WebAppInfo) error
}
