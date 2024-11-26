package methods

const (
	// Use this method to send photos. Might be used by default. Just mention a photo in Message structure.
	// Call GetResponse() in each interface you have created and gave to Message structure.
	Photo string = "sendPhoto"

	// Use this method to send audio files, if you want Telegram clients to display them in the music player.
	// Your audio must be in the .MP3 or .M4A format. On success, call GetResponse() in each interface you have
	// created and gave to Message structure. Bots can currently send audio files of up to 50 MB in size,
	// this limit may be changed in the future.
	Audio string = "sendAudio"

	// Use this method to send general files.
	// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
	Document string = "sendDocument"

	// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document).
	// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
	Video string = "sendVideo"

	// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
	// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
	Animation string = "sendAnimation"

	// Use this method to send audio files, if you want Telegram clients to display the file
	// as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS,
	// or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document).
	// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
	Voice      string = "sendVoice"
	VideoNote  string = "sendVideoNote"
	MediaGroup string = "sendMediaGroup"
	Message    string = "sendMessage"

	// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded.
	// Use usual function GetResponse() in interfaces you have mentioned to have response data from Telegram
	ForwardMessage string = "forwardMessage"

	// Use this method to forward multiple messages of any kind. If some of the specified
	// messages can't be found or forwarded, they are skipped. Service messages and messages
	// with protected content can't be forwarded. Album grouping is kept for forwarded
	// messagesUse this method to forward multiple messages of any kind. If some of the specified
	// messages can't be found or forwarded, they are skipped. Service messages and messages with
	// protected content can't be forwarded. Album grouping is kept for forwarded messages.
	// On success call {IMSGInformation}.GetMessageIDs() to have the response data from Telegram
	ForwardMessages string = "forwardMessages"

	// Use this method to copy messages of any kind. Service messages, paid media messages,
	// giveaway messages, giveaway winners messages, and invoice messages can't be copied. Copying means that the bot copys
	// the message and then sends it to the chat with the unique identifier you put in {IChat}.WriteChat{ID/Name}()
	// On success call {IMSGInformation}.GetMessageIDs() to have the response data from Telegram
	CopyMessage string = "copyMessage"

	// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied,
	// they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages,
	// and invoice messages can't be copied.
	CopyMessages string = "copyMessages"

	// Use this method to send paid media. Paid media can be only IPhoto or IVideo. Also, the fields of methods.Photo
	// and methds.PaidMedia are diffirent. So, please make sure because all data that you might mention may be ignored
	// because it isn't allowed to send this with this method
	PaidMedia string = "sendPaidMedia"
)

var Media map[string]struct{}

func init() {
	Media = make(map[string]struct{})

	Media[Photo] = struct{}{}
	Media[Audio] = struct{}{}
	Media[Document] = struct{}{}
	Media[Video] = struct{}{}
	Media[MediaGroup] = struct{}{}
	Media[Animation] = struct{}{}
	Media[Voice] = struct{}{}
}
