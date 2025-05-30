package fmtogram

import (
	"bytes"
	"mime/multipart"
	"time"
)

type handlerMedia interface {
	proccessFile(wr *multipart.Writer, cnttype *string, isgroup bool) error
	uniqueConst() (constID int)
}

type handleFiles interface {
	single(wr *multipart.Writer, buf *bytes.Buffer, contenttype *string) (queryStr string, err error)
	multiple(wr *multipart.Writer, buf *bytes.Buffer, contenttype *string) (queryStr string, err error)
}

type kb interface {
	get() ([]byte, error)
	isOK() error
}

// Use this interface to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for
// the bot, Telegram will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update [https://core.telegram.org/bots/api#update].
// In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give
// up after a reasonable amount of attempts. If you'd like to make sure that the webhook
// was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a
// header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
type IWebHook interface {
	// HTTPS URL to send updates to. Use an empty string to remove webhook integration
	WriteUrl(url string) error

	// Upload your public key certificate so that the root certificate in use can be checked. See
	// Telegram team's self-signed guide for details:
	// https://core.telegram.org/bots/self-signed
	WriteCertificate(path string) error

	// The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	WriteIPAddress(ip string) error

	// The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	// Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	WriteMaxConnections(max int) error

	// A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message",
	// "edited_channel_post", "callback_query"] to only receive updates of these  See Update[https://core.telegram.org/bots/api#update] for a complete list of
	// available update  Specify an empty list to receive all update types except chat_member, message_reaction,
	// and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this
	// parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received
	// for a short period of time.
	WriteAllowedUpdates(upds []string) error

	// Call it to drop all pending updates
	WriteDropPendingUpdates() error

	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters.
	// Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a
	// webhook set by you.
	WriteSecretToken(token string) error

	// Call  this function when you're ready to set a webhook. Requires only WriteUrl(), but it's possible to use
	// all "Write" functions.
	Set() error

	// Call this function to delete a webhook. Doesn't require anything, but could be given WriteDropPendingUpdates()
	Delete() error

	// Call this function to get all information about a webhook. Doesn't require any data. If the bot is using usual update system,
	// it will return an object with the url field empty.
	Info() (inf *WebhookInfo, err error)

	// Get status of the request after Set(), Delete() or Info()
	GetStatus() bool

	// Get error data of the request after Set(), Delete() or Info()
	GetError() (code int, msg string)
}

type IMessage interface {
	// Unique message identifier inside this chat. In specific instances
	// (e.g., message containing a video sent to a big chat), the server might
	// automatically schedule a message instead of sending it immediately. In such cases,
	// this field will be 0 and the relevant message will be unusable until it is actually sent
	ID() int

	// Sender of the message; may be empty for messages sent to channels. For backward
	// compatibility, if the message was sent on behalf of a chat, the field contains a
	// fake sender user in non-channel chats
	Sender() *User

	// Chat the message belongs to
	Chat() *Chat

	// Date the message was sent in Unix time. It is always a positive number, representing
	// a valid date and a special framework of golang for working with time
	Date() (unix int, t time.Time)

	// For replies in the same chat and message thread, the original message.
	// Note that the IMessage interface from this function will not contain further
	// Reply() fields even if it itself is a reply.
	Reply() IMessage

	// For text messages, the actual UTF-8 text of the message
	Text() string

	// Message is an audio file, information about the file
	Audio() *Audio

	// Message is a general file, information about the file
	Document() *Document

	// Message contains paid media; information about the paid media (photo or video)
	PaidMedia() *PaidMediaInfo

	// Message is a photo, available sizes of the photo
	Photo() []*PhotoSize

	// Message is a sticker, information about the sticker
	Sticker() *Sticker

	// Message is a forwarded story
	Story() *Story

	// Message is a video, information about the video
	Video() *Video

	// Message is a video note, information about the video message
	VideoNote() *VideoNote

	// Message is a voice message, information about the file
	Voice() *Voice

	// The whole other data that isn't specified in this interface
	Message() *Message
}

type IBusiness interface {
	// Unique identifier of the business connection
	ID() string

	// Business account user that created the business connection
	Sender() *User

	// Identifier of a private chat with the user who created the business connection.
	// This number may have more than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits
	SenderChatID() int64

	// Date the connection was established in Unix time and a special framework of golang for working with time
	Date() (unix int, t time.Time)

	// True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	CanReply() bool

	// True, if the connection is active
	IsEnabled() bool
}

type IBusinessMessage interface {
	// Unique identifier of the business connection
	ConnectionID() string

	// Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	Chat() *Chat

	// The list of identifiers of deleted messages in the chat of the business account
	MessageIDs() []int
}

type IMessageReaction interface {
	// The chat containing the message the user reacted to
	Chat() *Chat

	// Unique identifier of the message inside the chat
	MessageID() int

	// The user that changed the reaction, if the user isn't anonymous
	Sender() *User

	// The chat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat() *Chat

	// Date of the change in Unix time and a special framework of golang for working with time
	Date() (unix int, t time.Time)

	// Previous list of reaction types that were set by the sender
	OldReaction() []*ReactionType

	// New list of reaction types that have been set by the sender
	NewReaction() []*ReactionType
}

type IMessageReactionCount interface {
	// The chat containing the message
	Chat() *Chat

	// Unique message identifier inside the chat
	MessageID() int

	// Date of the change in Unix time and a special framework of golang for working with time
	Date() (unix int, t time.Time)

	// List of reactions that are present on the message
	Reactions() []*ReactionCount
}

type IInlineQuery interface {
	// Unique identifier for this query
	ID() string

	// Sender
	Sender() *User

	// 	Text of the query (up to 256 characters)
	Query() string

	// Offset of the results to be returned, can be controlled by the bot
	Offset() string
	// Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the
	// inline query sender, “private”, “group”, “supergroup”, or “channel”. The chat type should be always known
	// for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	ChatType() string

	// Sender location, only for bots that request user location
	Location() *Location
}

type IInlineResult interface {
	// The unique identifier for the result that was chosen
	ID() string

	// The sender of the request that chose the result
	Sender() *User

	// Sender location, only for bots that require user location
	Location() *Location

	// Identifier of the sent inline message. Available only if there is (formatter.IKeyboard).Inline() attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	InlineMessageID() string

	// The query that was used to obtain the result
	Query() string
}

// After the user presses a callback button, Telegram clients will display a progress bar until
// you call methods.AnswerCallbackQuery. It is, therefore, necessary to react by calling methods.AnswerCallbackQuery
// even if no notification to the user is needed (e.g., without specifying any of the optional parameters).
type ICallbackQuery interface {
	// Unique identifier for this query
	ID() string

	// Sender
	Sender() *User

	// Message sent by the bot with the callback button that originated the query
	Message() *MaybeInaccessibleMessage

	// Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID() string

	// Global identifier, uniquely corresponding to the chat to which the message with the callback
	// button was sent. Useful for high scores in games.
	ChatInstance() string

	// Data associated with the callback button. Be aware that the message originated the query can
	// contain no callback buttons with this data.
	Data() string

	// Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName() string
}

type IShippingQuery interface {
	// Unique query identifier
	ID() string

	// User who sent the query
	Sender() *User

	// 	Bot-specified invoice payload
	InvoicePayload() string

	// User specified shipping address
	ShippingAddress() *ShippingAddress
}

type IPreCheckoutQuery interface {
	// Unique query identifier
	ID() string

	// User who sent the query
	Sender() *User

	// Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	Currency() string

	// Total price in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows
	// the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount() int

	// Bot-specified invoice payload
	InvoicePayload() string

	// Identifier of the shipping option chosen by the user
	ShippingOptionID() string

	// Order information provided by the user
	OrderInfo() *OrderInfo
}

type IPaidMedia interface {
	// User who purchased the media
	Sender() *User

	// Bot-specified paid media payload
	Payload() string
}

type IPollUpdate interface {
	// Unique poll identifier
	ID() string

	// Poll question, 1-300 characters
	Question() string

	// Poll type, currently can be “regular” or “quiz”
	Type() string

	// Slice of poll options
	Options() []*PollOption

	// True, if the poll is closed
	IsClosed() bool

	// 0-based identifier of the correct answer option. Available only for polls in the quiz mode,
	// which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionID() string

	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation() string

	// All other possible data
	Poll() *Poll
}

type IPollAnswer interface {
	// Unique poll identifier
	ID() string

	// The chat that changed the answer to the poll, if the voter is anonymous
	VoterChat() *Chat

	// The user that changed the answer to the poll, if the voter isn't anonymous
	Sender() *User

	// 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIDs() []string
}

type IChatMember interface {
	// Chat the user belongs to
	Chat() *Chat

	// Performer of the action, which resulted in the change
	Sender() *User

	// Date of the change in Unix time and a special framework of golang for working with time
	Date() (unix int, t time.Time)

	// Previous information about the chat member
	OldMember() *ChatMember

	// New information about the chat member
	NewMemmber() *ChatMember

	// Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	InviteLink() *ChatInviteLink

	// True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaJoinRequest() bool

	// True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink() bool
}

type IJoinRequest interface {
	// Chat to which the request was sent
	Chat() *Chat

	// User that sent the join request
	Sender() *User

	// Identifier of a private chat with the user who sent the join request. This number may have more
	// than 32 significant bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision
	// float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to
	// send messages until the join request is processed, assuming no other administrator contacted the user.
	UserChatID() int64

	// Date the request was sent in Unix time and a special framework of golang for working with time
	Date() (unix int, t time.Time)

	// Bio of the user.
	Bio() string

	// Chat invite link that was used by the user to send the join request
	InviteLink() *ChatInviteLink
}

type IBoostUpdate interface {
	// Chat which was boosted
	Chat() *Chat

	// Information about the chat boost
	Boost() *ChatBoost
}

type IBoostRemove interface {
	// Chat which was boosted
	Chat() *Chat

	// Unique identifier of the boost
	BoostID() string

	// Point in time (Unix timestamp) when the boost was removed and a special framework of golang for working with time
	RemoveDate() (unix int, t time.Time)

	// Source of the removed boost
	Source() *ChatBoostSource
}

type ITelegram interface {
	// The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially.
	// This identifier becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates
	// or to restore the correct update sequence, should they get out of order. If there are no new updates for at least
	// a week, then identifier of the next update will be chosen randomly instead of sequentially.
	ID() int

	// New incoming message of any kind - text, photo, sticker, etc.
	Message() IMessage

	// New version of a message that is known to the bot and was edited. This update may at times be triggered by
	// changes to message fields that are either unavailable or not actively used by your bot.
	EditedMessage() IMessage

	// New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost() IMessage

	// New version of a channel post that is known to the bot and was edited. This update may at times be
	// triggered by changes to message fields that are either unavailable or not actively used by your bot.
	EditedChannelPost() IMessage

	// The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	BusinessConnection() IBusiness

	// New message from a connected business account
	BusinessMessage() IMessage

	// New version of a message from a connected business account
	EditedBusinessMessage() IMessage

	// Messages were deleted from a connected business account
	DeletedBusinessMessages() IBusinessMessage

	// A reaction to a message was changed by a user. The bot must be an administrator in the chat.
	// The update isn't received for reactions set by bots.
	MessageReaction() IMessageReaction

	// Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the
	// chat. The data is grouped and can be sent with delay up to a few minutes.
	MessageReactionCount() IMessageReactionCount

	// New incoming inline query
	InlineQuery() IInlineQuery

	// The result of an inline query that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	ChosenInlineResult() IInlineResult

	// New incoming callback query
	CallbackQuery() ICallbackQuery

	// New incoming shipping query. Only for invoices with flexible price
	ShippingQuery() IShippingQuery

	// New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery() IPreCheckoutQuery

	// A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	PurchasedPaidMedia() IPaidMedia

	// New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	Poll() IPollUpdate

	// A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer() IPollAnswer

	// The bot's chat member status was updated in a chat. For private chats, this data is received
	// only when the bot is blocked or unblocked by the user.
	MyChatMember() IChatMember

	// A chat member's status was updated in a chat. The bot must be an administrator in the chat.
	ChatMember() IChatMember

	// A request to join the chat has been sent. The bot must have the
	// ChatAdministratorRights.CanInviteUsers administrator right in the chat to receive this data.
	ChatJoinRequest() IJoinRequest

	// A chat boost was added or changed. The bot must be an administrator in the chat to receive this data.
	ChatBoost() IBoostUpdate

	// A boost was removed from a chat. The bot must be an administrator in the chat to receive this data.
	RemovedChatBoost() IBoostRemove
}

type IGet interface {
	Status() bool
	Error() (code int, msg string)
	Chat() *Chat
	Sender() *User
	Date() int
	MessageID() int
	MessageIDs() []int
	Replyed() IGet
	ForwardOrigin() *MessageOrigin
	Photo() []*PhotoSize
	Audio() *Audio
	Document() *Document
	Video() *Video
	Animation() *Animation
	Voice() *Voice
	VideoNote() *VideoNote
	PaidMedia() *PaidMedia
	MediaGroupID() string
	Photos() [][]*PhotoSize
	Videos() []*Video
	Audios() []*Audio
	Documents() []*Document
	Poll() *Poll
	Dice() *Dice
	ProfilePhotos() *UserProfilePhotos
	File() *File
	Stickers() []*Sticker
	Gifts() []*Gift
	Message() *TelegramMessage
	Messages() []*TelegramMessage
	String() string
	InviteLink() *ChatInviteLink
	ChatInfo() *ChatFullInfo
	Members() []*ChatMember
	Integer() *int
	Forum() *ForumTopic
	Boosts() []*ChatBoost
	BusinessConnection() *BusinessConnection
	Commands() []*BotCommand
	MenuButton() *MenuButton
	AdminRights() *ChatAdministratorRights
	PreparedInlineMessage() *PreparedInlineMessage
	StarTransaction() *StarTransaction
	Score() []*GameHighScore
	Request() string
	Response() string
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
	// Parsemode can be only HTML, Markdown and MarkdownV2. Cannot be an empty string.
	WriteParseMode(parsemode string) error

	// entities can't be an empty slice.
	WriteCaptionEntities(entities []*MessageEntity) error

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia() error

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler() error
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

	WriteCaptionEntities(entities []*MessageEntity) error
	WriteDuration(duration int) error

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler() error
	WriteHeight(height int) error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: HTML, Markdown and...
	WriteParseMode(parsemode string) error

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia() error

	WriteSupportsStreaming() error

	WriteThumbnail(path string) error

	WriteThumbnailID(thimbnailID string) error

	WriteWidth(width int) error

	WriteCoverStorage(path string) error

	WriteCoverTelegram(coverID string) error

	WriteCoverInternet(url string) error
}

type IAudio interface {
	// Receives a path of an audio
	WriteAudioStorage(path string) error

	// Receives a Telegram ID of an audio
	WriteAudioTelegram(AudioID string) error

	// Recueves an URL-link of an audio
	WriteAudioInternet(URL string) error

	// Receives a string that will be a caption of the audio
	WriteCaption(caption string) error

	WriteCaptionEntities(entities []*MessageEntity) error

	// Duration of the audio in seconds
	WriteDuration(duration int) error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: HTML, Markdown and...
	WriteParseMode(parsemode string) error

	WritePerformer(performer string) error

	WriteThumbnail(path string) error

	WriteThumbnailID(thimbnailID string) error

	WriteTitle(title string) error
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

	WriteCaptionEntities(entities []*MessageEntity) error
	WriteDisableContentTypeDetection() error

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: HTML, Markdown and...
	WriteParseMode(parsemode string) error

	WriteThumbnail(path string) error

	WriteThumbnailID(thimbnailID string) error
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
	WriteThumbnail(path string) error

	// Receives a Telegram ID of an animation
	WriteThumbnailID(thumbnailID string) error

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler() error
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
	WriteThumbnail(path string) error

	// Receives a Telegram ID of a thumbnail
	WriteThumbnailID(thumbnailID string) error
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

	// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360
	WriteHeading(heading int) error

	// For live locations, a maximum distance for proximity alerts about approaching another chat member,
	// in meters. Must be between 1 and 100000
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
}

type IPoll interface {
	// Receives a string that will be the question of a poll, 1-300 characters
	WriteQuestion(question string) error

	WriteQuestionParseMode(parseode string) error

	WriteQuestionEntities(entities []*MessageEntity) error

	// Receives a slice of 2-10 answer options
	WriteOptions(options []*PollOption) error

	// Call it only if the poll needs to be anonymous
	WriteAnonymous(yesno bool) error

	// Receives a poll type, "quiz" or "regular"
	WriteType(polltype string) error

	//  Call it only if the poll needs to allow multiple answers, ignored for polls in quiz mode
	WriteAllowMultipleAnswers() error

	// 0-based identifier of the correct answer option, required for polls in quiz mode
	WriteCorrectOptionID(optID string) error

	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll,
	// 1-200 characters with at most 2 line feeds after entities parsing
	WriteExplanation(explanation string) error

	WriteExplanationParseMode(parsemode string) error

	WriteExplanationEntities(entities []*MessageEntity) error

	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with (formatter.IPoll).WriteCloseDate
	WriteOpenPeriod(period int) error

	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with (formatter.IPoll).WriteOpenPeriod
	WriteCloseDate(time int) error

	// Call it only if the poll needs to be immediately closed. This can be useful for poll preview.
	WriteClosed() error
}

type ISticker interface {
	WriteStickerStorage(path string) error

	WriteStickerTelegram(stickerID string) error

	WriteStickerInternet(url string) error

	WriteAssociatedEmoji(emoji string) error

	// WriteAssociatedEmojies(emojies []string) error

	// WriteEmojiID(emojiID string) error

	// WriteEmojiIDs(emojiIDs []string) error

	// WriteFormat(format string) error

	// // New sticker position in the set, zero-based
	// WritePosition(pos string) error

	// WriteOldSticker(stickerID string) error

	// // list of 0-20 search keywords for the sticker with total length of up to 64 characters
	// WriteKeywords(words []string) error

	// WriteMaskPosition(maskpos *MaskPosition) error

	// WriteThumbnailStorage(path string) error

	// WriteThumbnailTelegram(thumbnailID string) error

	// WriteThumbnailInternet(url string) error

	// WriteThumbnailFormat(format string) error

	// WriteGiftID(giftID string) error

	// WritePayForUpgrade() error
}

type IForum interface {
	// Topic name, 1-128 characters
	WriteName(name string) error

	// Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590
	// (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	WriteIconColor(color int) error

	WriteIconEmojiID(emojiID string) error
}

type IBot interface {
	// Recieves a slice of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	WriteCommands(commands []*BotCommand) error

	WriteScope(scope *BotCommandScope) error

	// Recieves a two-letter ISO 639-1 language code.
	WriteLanguage(lang string) error

	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	WriteName(name string) error

	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	WriteDescription(description string) error

	// New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	WriteShortDescription(description string) error

	WriteMenuButton(button *MenuButton) error

	WriteRights(rights *ChatAdministratorRights) error

	WriteForChannels() error
}

type IInlineMode interface {
	WriteQueryID(queryID string) error

	WriteWebAppQueryID(queryID string) error

	WriteCachedAudio(cachedAudio *InlineQueryResultCachedAudio) error

	WriteCachedDocument(cachedDocument *InlineQueryResultCachedDocument) error

	WriteCachedGif(cachedGif *InlineQueryResultCachedGif) error

	WriteCachedMpeg4Gif(cachedMpeg4Gif *InlineQueryResultCachedMpeg4Gif) error

	WriteCachedPhoto(cachedPhoto *InlineQueryResultCachedPhoto) error

	WriteCachedSticker(cachedSticker *InlineQueryResultCachedSticker) error

	WriteCachedVideo(cachedVideo *InlineQueryResultCachedVideo) error

	WriteCachedVoice(cachedVoice *InlineQueryResultCachedVoice) error

	WriteArticle(art *InlineQueryResultArticle) error

	WriteAudio(ad *InlineQueryResultAudio) error

	WriteContact(cont *InlineQueryResultContact) error

	WriteGame(game *InlineQueryResultGame) error

	WriteDocument(doc *InlineQueryResultDocument) error

	WriteGif(gif *InlineQueryResultGif) error

	WriteLocation(loc *InlineQueryResultLocation) error

	WriteMpeg4Gif(mpeg4gif *InlineQueryResultMpeg4Gif) error

	WritePhoto(ph *InlineQueryResultPhoto) error

	WriteVenue(ven *InlineQueryResultVenue) error

	WriteVideo(vd *InlineQueryResultVideo) error

	WriteVoice(vc *InlineQueryResultVoice) error

	WriteCacheTime(time int) error

	WriteIsPersonal() error

	WriteNextOffset(offset string) error

	// Recieves an object that represents a button to be shown above inline query results. You must use exactly one of the optional fields.
	WriteButton(but *InlineQueryResultsButton) error

	WriteAllowUserChats() error

	WriteAllowBotChats() error

	WriteAllowGroupChats() error

	WriteAllowChannelChats() error

	WriteIntoArray()
}

type IParameters interface {
	// Doesn't revieve anything. If the function was called, the message you send to a client will be gotten without a notification
	WriteDisableNotification() error

	// Recieves a slice of special entities that appear in message text, which can be specified instead of WriteParseMode()
	WriteEntities(entities []*MessageEntity) error

	// Recieves a link preview generation options for the message
	WriteLinkPreviewOptions(link *LinkPreviewOptions) error

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

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: HTML, Markdown and...
	WriteParseMode(parsemode string) error

	// Doesn't recieve anything. If the function was called, the text-message will be protected from forwarding and saving
	WriteProtectContent() error

	// The text content of the message
	WriteString(text string) error

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia() error

	// Recieves a structure that has description of the message to reply to
	WriteReplyParameters(replyPar *ReplyParameters) error

	// Call it to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	WriteAllowPaidBroadcast() error

	// The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	WriteStarCount(amount int) error

	// Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	WritePayload(payload string) error

	// Use this only if you're going to use method methods.Dice. In anyother cases you'd prefer put emoji just in text message.
	// All emojis that available for this function are in the Emojis array
	WriteEmoji(emoji string) error

	// Type of action to broadcast. Choose one of Actions, depending on what the user is about to receive:
	// Actions[0] for text messages, Actions[1] for photos, Actions[2] or Actions[3] for videos,
	// Actions[4] or Actions[5] for voice notes, Actions[6] for general files, Actions[7]
	// for stickers, Actions[8] for location data, Actions[9] or Actions[10] for video notes.
	WriteAction(action string) error

	// Recieves a slice of reaction types to set on the message. Currently, as non-premium users,
	// bots can set up to one reaction per message. A custom emoji reaction can be used if it is either
	// already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	WriteReaction(reaction []*ReactionType) error

	// Call it to set the reaction with a big animation
	WriteReactionIsBig() error

	// Sequential number of the first photo to be returned. By default, all photos are returned.
	WriteOffset(offset int) error

	WriteLimit(limit int) error

	WriteEmojiStatusCustomEmojiID(emojiID string) error

	WriteEmojiStatusExpirationDate(date int) error

	WriteFileID(fileID string) error

	// Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than
	// 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	WriteUntilDate(date time.Duration) error

	// Call it to delete all messages from the chat for the user that is being removed. If you don't call it when you just banned a user, the
	// user will be able to see messages in the group that were sent before the user was removed.
	// Always setted for supergroups and channels
	WriteRevokeMessages() error

	// Do nothing if the user is not banned
	WriteOnlyIfBanned() error

	WritePermissions(permissions *ChatPermissions) error

	WriteIndependentChatPermissions() error

	WriteAdministratorRights(chAdminRights *ChatAdministratorRights) error

	// New custom title for the administrator; 0-16 characters, emoji are not allowed
	WriteCustomTitle(title string) error

	WriteUserID(userID int) error

	// 	Unique identifier for the query to be answered
	WriteCallBackQueryID(queryID string) error

	// An alert will be shown by the client instead of a notification at the top of the chat screen
	WriteShowAlert() error

	WriteURL(botnickname, url string) error

	// The maximum amount of time in seconds that the result of the callback query may be cached
	// client-side. Telegram apps will support caching starting in version 3.14.
	WriteCacheTime(time int) error

	WriteInlineMessageID(messageID string) error

	WriteErrors(errors []*PassportElementError) error

	WriteDescription(desc string) error

	WriteRemoveCaption() error

	WriteVideoStartTimestamp(seconds int) error

	WriteSetName(name string) error

	WriteSetTitle(title string) error

	WriteStickerType(stickertype string) error

	WriteNeedsRepainting() error
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

	WriteSenderChatID(chatID int) error

	// New chat title, 1-128 characters
	WriteTitle(title string) error

	// New chat description, 0-255 characters
	WriteDescription(description string) error
}

type ILink interface {
	// Invite link name; 0-32 characters
	WriteName(name string) error

	// Point in time (Unix timestamp) when the link will expire
	WriteExpireDate(date time.Duration) error

	// The maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999. Incompatible with (formatter.ILink).WriteJoinRequest()
	WriteMemberLimit(limit int) error

	// If users joining the chat via the link need to be approved by chat administrators.
	// Incompatible with (formatter.ILink).WriteMemberLimit()
	WriteJoinRequest() error

	// The invite link to edit
	WriteInviteLink(link string) error

	// The number of seconds the subscription will be active for before the next payment.
	// Currently, it must always be 2592000 (30 days).
	WriteSubscriptionPeriod(period int) error

	// The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500
	WriteSubscriptionPrice(price int) error
}

// You can create and send only one type of keyboard. There are 3 available types
type IKeyboard interface {
	WriteReply() (IReply, error)

	WriteInline() (IInline, error)

	WriteForceReply() (IForceReply, error)
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

	// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible, use (IReply).WriteOneTimeKeyboard()
	WriteRemove() error
}

// This interface represents one button of IReply interface. WriteString() is requied, and at most one of the rest functions must be used to specify type of the button. For simple text buttons, String can be used instead of this object to specify the button text.
type IReplyButton interface {
	// If specified, pressing the button will open a list of suitable users. Available in private chats only.
	WriteRequestChat(rchat *KeyboardButtonRequestChat) error

	// Doesn't receives anything. If it is called, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	WriteRequestContact() error

	// Doesn't receives anything. If it is called, the user's current location will be sent when the button is pressed. Available in private chats only.
	WriteRequestLocation() error

	// Receives a structure with some data. The user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WriteRequestPoll(rpoll *KeyboardButtonPollType) error

	// If specified, pressing the button will open a list of suitable users. Available in private chats only.
	WriteRequestUsers(rusers *KeyboardButtonRequestUsers) error

	// The text content that will be on the button
	WriteString(text string) error

	// Receives a structure that describes Web App. Will be launched when the button is pressed. Available in private chats only.
	WriteWebApp(webapp *WebAppInfo) error
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
	WriteCallbackGame(game *CallbackGame) error

	// Recieves an HTTPS URL structure used to automatically authorize the user.
	WriteLoginUrl(logUrl *LoginUrl) error

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
	WriteSwitchInlineQueryChosenChat(sw *SwitchInlineQueryChosenChat) error

	// If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages
	// sent on behalf of a Telegram Business account.
	WriteSwitchInlineQueryCurrentChat(sw string) error

	// Receives a URL string that will be received by the bot after a client pressed on the button
	WriteURL(URL string) error

	// Receives the description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	WriteWebApp(webapp *WebAppInfo) error
}

type IForceReply interface {
	// Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	WriteForceReply() error

	// The placeholder to be shown in the input field when the reply is active; 1-64 characters
	WriteInputFieldPlaceholder(placeholder string) error

	// Use this parameter if you want to force reply from specific users only. Targets: 1) users that
	// are @mentioned in the text of the Message object; 2) if the bot's message is a reply
	// to a message in the same chat and forum topic, sender of the original message.
	WriteSelective() error
}

type IPayment interface {
	// 1-32 characters
	WriteTitle(title string) error

	// 1-255 characters
	WriteDescription(description string) error

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	WritePayload(payload string) error

	WriteProviderToken(token string) error

	WriteCurrency(currency string) error

	WritePrices(prices []*LabeledPrice) error

	WriteSubscriptionPeriod(period int) error

	WriteMaxTipAmount(amout int) error

	WriteSuggestedTipAmounts(amounts []int) error

	WriteStartParameter(prm string) error

	WriteProviderData(data string) error

	WritePhotoUrl(url string) error

	WritePhotoSize(size int) error

	WritePhotoWidth(width int) error

	WritePhotoHeight(height int) error

	WriteNeedName() error

	WriteNeedPhoneNumber() error

	WriteNeedEmail() error

	WriteNeedShippingAddress() error

	WriteSendPhoneNumberToProvider() error

	WriteSendEmailToProvider() error

	WriteIsFlexible() error

	// WriteShippingID(id string) error

	// WriteOK(b bool) error

	// WriteShippingOptions(options []*ShippingOption) error

	// WriteErrorMessage(msg string) error

	// WritePreCheckoutID(id string) error

	// WriteTelegramPaymentChargeID(id string) error

	// WriteIsCanceled() error
}

type IGame interface {
	WriteShortName(name string) error

	// WriteScore(score int) error

	// WriteForce() error

	// WriteDisableEditMessage() error
}

type IGift interface {
	WriteID(id string) error
	WritePayForUpgrade() error
}
