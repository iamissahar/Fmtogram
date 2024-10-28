package formatter

import (
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/types"
)

type handlerMedia interface {
	multipartFields(*multipart.Writer, *[]interface{}, int, bool) error
	jsonFileds() ([]byte, error)
}

type handlerKB interface {
	MultipartFields(*multipart.Writer) error
	JsonFields() ([]byte, error)
}

type IPhoto interface {
	// Receives a path of a photo
	WritePhotoStorage(string)

	// Recieves a Telegram ID of a photo
	WritePhotoTelegram(string)

	// Recieves an URL-link of a photo
	WritePhotoInternet(string)

	// Receives a string that will be a caption of the photo. If you're going to send
	// one picture and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(string)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	WriteCaptionEntities([]*types.MessageEntity)

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia()

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler()

	// You can call this function after calling Send(). It returns you an array with some data about the photo you just sent. 4 is the biggest amount of data. Sometimes there might be nil
	GetResponse() [4]types.PhotoSize
}

type IVideo interface {
	// Receives a path of a video
	WriteVideoStorage(string)

	// Recieves a Telegram ID of a video
	WriteVideoTelegram(string)

	// Recieves an URL-link of a video
	WriteVideoInternet(string)

	// Receives a string that will be a caption of the video. If you're going to send
	// one video and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(string)

	WriteCaptionEntities([]*types.MessageEntity)
	WriteDuration(int)

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler()
	WriteHeight(int)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia()
	WriteSupportsStreaming()
	WriteThumbnailStorage(string)
	WriteThumbnailTelegram(string)
	WriteThumbnailInternet(string)
	WriteWidth(int)

	// You can call this function after calling Send(). It returns you a structure with some data about the video you just sent
	GetResponse() types.Video
}

type IAudio interface {
	// Receives a path of an audio
	WriteAudioStorage(string)

	// Receives a Telegram ID of an audio
	WriteAudioTelegram(string)

	// Recueves an URL-link of an audio
	WriteAudioInternet(string)

	// Receives a string that will be a caption of the audio. If you're going to send
	// one audio and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(string)

	WriteCaptionEntities([]*types.MessageEntity)

	// Duration of the audio in seconds
	WriteDuration(int)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	WritePerformer(string)
	WriteThumbnailStorage(string)
	WriteThumbnailTelegram(string)
	WriteThumbnailInternet(string)
	WriteTitle(string)

	// You can call this function after calling Send(). It returns you a structure with some data about the audio you just sent
	GetResponse() types.Audio
}

type IDocument interface {
	// Receives a string that will be a caption of the document. If you're going to send
	// one document and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(string)

	WriteCaptionEntities([]*types.MessageEntity)
	WriteDisableContentTypeDetection()

	// Receives a path of an audio
	WriteDocumentStorage(string)

	// Receives a Telegram ID of an audio
	WriteDocumentTelegram(string)

	// Recueves an URL-link of an audio
	WriteDocumentInternet(string)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	WriteThumbnailStorage(string)
	WriteThumbnailTelegram(string)
	WriteThumbnailInternet(string)

	// You can call this function after calling Send(). It returns you a structure with some data about the document you just sent
	GetResponse() types.Document
}

type IAnimation interface {
	// Receives a path of an animation
	WriteAnimationStorage(string)

	// Receives a Telegram ID of an animation
	WriteAnimationTelegram(string)

	// Receives an URL-link of an animation
	WriteAnimationInternet(string)

	// Duration of sent animation in seconds
	WriteDuration(int)

	// Animation width
	WriteWidth(int)

	// Animation height
	WriteHeight(int)

	// Receives a path of an animation
	WriteThumbnailStorage(string)

	// Receives a Telegram ID of an animation
	WriteThumbnailTelegram(string)

	// Receives an URL-link of an animation
	WriteThumbnailInternet(string)

	// Receives a string of animation caption, 0-1024 characters after entities parsing
	WriteCaption(string)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	// A slice of special entities that appear in the caption, which can be specified instead of {IAnimation}.WriteParseMode()
	WriteCaptionEntities([]*types.MessageEntity)

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia()

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler()
}

type IMSGInformation interface {
	// Doesn't revieve anything. If the function was called, the message you send to a client will be gotten without a notification
	WriteDisableNotification()

	// Recieves a slice of special entities that appear in message text, which can be specified instead of WriteParseMode()
	WriteEntities([]*types.MessageEntity)

	// Recieves a link preview generation options for the message
	WriteLinkPreviewOptions(*types.LinkPreviewOptions)

	// Recieves an unique identifier of the message effect to be added to the message; for private chats only
	WriteMessageEffectID(string)

	// Recieves an unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	WriteMessageThreadID(int)

	// Recieves an unique identifier for the message you want to do somthing
	WriteMessageID(int)

	// Recieves a slice of unique identifiers for the messages you want to do somthing. Only ints (number of message id). The identifiers must be specified in a strictly increasing order.
	WriteMessageIDs([]int) error

	// New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	WriteCaption(string)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	// Doesn't recieve anything. If the function was called, the text-message will be protected from forwarding and saving
	WriteProtectContent()

	// The text content of the message
	WriteString(string)

	// Call it, if the caption must be shown above the message media (only if you copy a message that has Media Type). Ignored if a new caption isn't specified.
	WriteShowCaptionAboveMedia()

	// Recieves a structure that has description of the message to reply to
	WriteReplyParameters(*types.ReplyParameters)

	// You can call this function after calling Send(). It returns you a structure with some data about the user you just sent a message to
	GetResponse() types.User

	// Call it after methods.ForwardMessages, methods.CopyMessage and methods.CopyMessages. In any other cases it returns nil
	GetMessageIDs() []int
}

type IChat interface {
	// Reseives an unique identifier of the business connection on behalf of which the message will be sent
	WriteBusinessConnectionID(string)

	// These the functions might be used only if you need to send a message (!NOT RESPONSE!) to a client.
	// There are some information about the chat the request was gotten from, so you do not need it at all.
	//
	// Receives a number (int). Must be the chatID you want to send a message to
	WriteChatID(int)

	// Receives the name of the channel. Must be the name of the channel you want to send a message to. Works only with public channales
	WriteChatName(string)

	// // Receives the URL of the chat. Must be the the URL of you want to send a message to
	// WriteChatURL(chatURL string)

	// Receives a number (int). Must be the chatID of the original message that was sent
	WriteFromChatID(int)

	// Receives the name of the chat. Must be the name of the chat of the original message that was sent
	WriteFromChatName(string)

	// You can call this function after calling Send(). It returns you a structure with some data about the chat you just sent a message to
	GetResponse() types.Chat
}

// This interface represents a custom keyboard with reply options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
type IReply interface {
	// Receives 2 coordinates. The first one is a line you expect to see the button and the second one is a position on the line you just mentioned.
	// The coordinates mustn't be a random numbers. You have to call function Set() first. Returns an interface of the new created button
	// and might return and error, if the coordinates are wrong
	NewButton(int, int) (IReplyButton, error)

	// Receives a slice of integers and sets the plan of future reply-keyboard. Be careful with data, because one array cell is considered
	// as row, so, for example: a slice []int{2, 1, 4} means that in the reply-keyboard will be 3 rows and in the first one there'll be 2 buttons, in the next one: just one
	// and in the last one: 4 buttons
	Set([]int)
}

// This interface represents one button of IReply interface. WriteString() is requied, and at most one of the rest functions must be used to specify type of the button. For simple text buttons, String can be used instead of this object to specify the button text.
type IReplyButton interface {
	// If specified, pressing the button will open a list of suitable users. Available in private chats only.
	WriteRequestChat(*types.KeyboardButtonRequestChat)

	// Doesn't receives anything. If it is called, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	WriteRequestContact()

	// Doesn't receives anything. If it is called, the user's current location will be sent when the button is pressed. Available in private chats only.
	WriteRequestLocation()

	// Receives a structure with some data. The user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WriteRequestPoll(*types.KeyboardButtonPollType)

	// If specified, pressing the button will open a list of suitable users. Available in private chats only.
	WriteRequestUsers(*types.KeyboardButtonRequestUsers)

	// The text content that will be on the button
	WriteString(string)

	// Receives a structure that describes Web App. Will be launched when the button is pressed. Available in private chats only.
	WriteWebApp(*types.WebAppInfo)
}

// This interfaceinterface represents an inline keyboard that appears right next to the message it belongs to.
type IInline interface {
	// Receives 2 coordinates. The first one is a line you expect to see the button and the second one is a position on the line you just mentioned.
	// The coordinates mustn't be a random numbers. You have to call function Set() first. Returns an interface of the new created button
	// and might return and error, if the coordinates are wrong
	NewButton(int, int) (IInlineButton, error)

	// Receives a slice of integers and sets the plan of future inline-keyboard. Be careful with data, because one array cell is one array cell is considered
	// as row, so, for example: a slice []int{2, 1, 4} means that in the inline-keyboard will be 3 rows and in the first one there'll be 2 buttons, in the next one: just one
	// and in the last one: 4 buttons
	Set([]int)
}

// This interface represents one button of an IInline interface. WriteString is requied. The rest aren't, but exactly one function must be used to specify type of the button.
type IInlineButton interface {
	// Receives some text string that will be received by the bot after a client pressed on the button
	WriteCallbackData(string)

	// Receives the description of the game that will be launched when the user presses the button.
	//
	// This type of button must always be the first button in the first row
	WriteCallbackGame(*types.CallbackGame)

	// Recieves an HTTPS URL structure used to automatically authorize the user.
	WriteLoginUrl(*types.LoginUrl)

	// Specify True, to send a Pay button. Substrings “⭐” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.
	//
	// This type of button must always be the first button in the first row and can only be used in invoice messages.
	WritePay()

	// Receives some text string that will be seen by client on the button
	WriteString(string)

	// If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	//
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages
	// sent on behalf of a Telegram Business account.
	WriteSwitchInlineQuery(string)

	// If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field.
	// Not supported for messages sent on behalf of a Telegram Business account.
	WriteSwitchInlineQueryChosenChat(*types.SwitchInlineQueryChosenChat)

	// If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages
	// sent on behalf of a Telegram Business account.
	WriteSwitchInlineQueryCurrentChat(string)

	// Receives a URL string that will be received by the bot after a client pressed on the button
	WriteURL(string)

	// Receives the description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	WriteWebApp(*types.WebAppInfo)
}
