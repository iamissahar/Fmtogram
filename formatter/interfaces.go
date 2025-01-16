package formatter

import (
	"mime/multipart"
	"time"

	"github.com/l1qwie/Fmtogram/types"
)

type handlerMedia interface {
	multipartFields(writer *multipart.Writer, group *[]interface{}, id int, input bool) error
	jsonFileds() (jsbody []byte, err error)
	uniqueConst() (constID int)
}

type kb interface {
	get() ([]byte, error)
	isOK() error
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
	WriteSupportStreaming() error
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

	// Receives a string that will be a caption of the audio
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

	WriteExplanationEntities(entities []*types.MessageEntity) error

	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with (formatter.IPoll).WriteCloseDate
	WriteOpenPeriod(period int) error

	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with (formatter.IPoll).WriteOpenPeriod
	WriteCloseDate(time int) error

	// Call it only if the poll needs to be immediately closed. This can be useful for poll preview.
	WriteClosed() error

	GetResponse() types.Poll
}

type ISticker interface {
	// Name of the sticker set to be set as the group sticker set
	WriteSetName(name string) error

	WriteStickerStorage(path string) error

	WriteStickerTelegram(stickerID string) error

	WriteStickerInternet(url string) error

	WriteAssociatedEmoji(emoji string) error

	WriteAssociatedEmojies(emojies []string) error

	WriteEmojiID(emojiID string) error

	WriteEmojiIDs(emojiIDs []string) error

	WriteFormat(format string) error

	WriteTitle(title string) error

	WriteStickerType(stickertype string) error

	WriteNeedsRepainting() error

	// New sticker position in the set, zero-based
	WritePosition(pos string) error

	WriteOldSticker(stickerID string) error

	// list of 0-20 search keywords for the sticker with total length of up to 64 characters
	WriteKeywords(words []string) error

	WriteMaskPosition(maskpos *types.MaskPosition) error

	WriteThumbnailStorage(path string) error

	WriteThumbnailTelegram(thumbnailID string) error

	WriteThumbnailInternet(url string) error

	WriteThumbnailFormat(format string) error

	WriteGiftID(giftID string) error

	WriteTextParseMode(parsemode string) error

	WriteTextEntities(entities []*types.MessageEntity) error
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
	WriteCommands(commands []*types.BotCommand) error

	WriteScope(scope *types.BotCommandScope) error

	// Recieves a two-letter ISO 639-1 language code.
	WriteLanguage(lang string) error

	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	WriteName(name string) error

	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	WriteDescription(description string) error

	// New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	WriteShortDescription(description string) error

	WriteMenuButton(button *types.MenuButton) error

	WriteRights(rights *types.ChatAdministratorRights) error

	WriteForChannels() error
}

type IResult interface {
	WriteCachedAudio(cachedAudio *types.InlineQueryResultCachedAudio) error

	WriteCachedDocument(cachedDocument *types.InlineQueryResultCachedDocument) error

	WriteCachedGif(cachedGif *types.InlineQueryResultCachedGif) error

	WriteCachedMpeg4Gif(cachedMpeg4Gif *types.InlineQueryResultCachedMpeg4Gif) error

	WriteCachedPhoto(cachedPhoto *types.InlineQueryResultCachedPhoto) error

	WriteCachedSticker(cachedSticker *types.InlineQueryResultCachedSticker) error

	WriteCachedVideo(cachedVideo *types.InlineQueryResultCachedVideo) error

	WriteCachedVoice(cachedVoice *types.InlineQueryResultCachedVoice) error

	WriteArticle(art *types.InlineQueryResultArticle) error

	WriteAudio(ad *types.InlineQueryResultAudio) error

	WriteContact(cont *types.InlineQueryResultContact) error

	WriteGame(game *types.InlineQueryResultGame) error

	WriteDocument(doc *types.InlineQueryResultDocument) error

	WriteGif(gif *types.InlineQueryResultGif) error

	WriteLocation(loc *types.InlineQueryResultLocation) error

	WriteMpeg4Gif(mpeg4gif *types.InlineQueryResultMpeg4Gif) error

	WritePhoto(ph *types.InlineQueryResultPhoto) error

	WriteVenue(ven *types.InlineQueryResultVenue) error

	WriteVideo(vd *types.InlineQueryResultVideo) error

	WriteVoice(vc *types.InlineQueryResultVoice) error
}

type IInlineMode interface {
	WriteQueryID(queryID string) error

	WriteWebAppQueryID(queryID string) error

	WriteResult() (res IResult, err error)

	WriteResults(length int) (res []IResult, err error)

	WriteCacheTime(time int) error

	WriteIsPersonal() error

	WriteNextOffset(offset string) error

	// Recieves an object that represents a button to be shown above inline query results. You must use exactly one of the optional fields.
	WriteButton(but *types.InlineQueryResultsButton) error

	WriteAllowUserChats() error

	WriteAllowBotChats() error

	WriteAllowGroupChats() error

	WriteAllowChannelChats() error

	GetSentWebAppMessage() types.SentWebAppMessage

	GetPreparedInlineMessage() types.PreparedInlineMessage
}

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

	// Use this only if you're going to use method methods.Dice. In anyother cases you'd prefer put emoji just in text message.
	// All emojis that available for this function are in the types.Emojis array
	WriteEmoji(emoji string) error

	// Type of action to broadcast. Choose one of types.Actions, depending on what the user is about to receive:
	// types.Actions[0] for text messages, types.Actions[1] for photos, types.Actions[2] or types.Actions[3] for videos,
	// types.Actions[4] or types.Actions[5] for voice notes, types.Actions[6] for general files, types.Actions[7]
	// for stickers, types.Actions[8] for location data, types.Actions[9] or types.Actions[10] for video notes.
	WriteAction(action string) error

	// Recieves a slice of reaction types to set on the message. Currently, as non-premium users,
	// bots can set up to one reaction per message. A custom emoji reaction can be used if it is either
	// already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	WriteReaction(reaction []*types.ReactionType) error

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

	WritePermissions(permissions *types.ChatPermissions) error

	WriteIndependentChatPermissions() error

	WriteAdministratorRights(chAdminRights *types.ChatAdministratorRights) error

	// New custom title for the administrator; 0-16 characters, emoji are not allowed
	WriteCustomTitle(title string) error

	WriteUserID(userID int) error

	// 	Unique identifier for the query to be answered
	WriteCallBackQueryID(queryID string) error

	// An alert will be shown by the client instead of a notification at the top of the chat screen
	WriteShowAlert() error

	WriteURL(url string) error

	// The maximum amount of time in seconds that the result of the callback query may be cached
	// client-side. Telegram apps will support caching starting in version 3.14.
	WriteCacheTime(time int) error

	WriteInlineMessageID(messageID string) error

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

	WriteSenderChatID(chatID int) error

	// New chat title, 1-128 characters
	WriteTitle(title string) error

	// New chat description, 0-255 characters
	WriteDescription(description string) error

	// You can call this function after calling Send(). It returns you a structure with some data about the chat you just sent a message to
	GetResponse() types.Chat
}

type ILink interface {
	// Invite link name; 0-32 characters
	WriteName(name string) error

	// Point in time (Unix timestamp) when the link will expire
	WriteExpireDate(date time.Duration) error

	// The maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	WriteMemberLimit(limit int) error

	// If users joining the chat via the link need to be approved by chat administrators.
	// If called, member_limit can't be specified
	WriteJoinRequest() error

	// The invite link to edit
	WriteIniveLink(link string) error

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
