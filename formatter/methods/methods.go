package methods

const (
	// Use this method to send text messages
	Message string = "sendMessage"

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
	Voice string = "sendVoice"

	// Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages.
	VideoNote string = "sendVideoNote"

	// Use this method to send paid media. Paid media can be only IPhoto or IVideo. Also, the fields of methods.Photo
	// and methds.PaidMedia are diffirent. So, please make sure because all data that you might mention may be ignored
	// because it isn't allowed to send this with this method
	PaidMedia string = "sendPaidMedia"

	// Use this method to send a group of photos, videos, documents or audios as an album.
	// Documents and audio files can be only grouped in an album with messages of the same type.
	MediaGroup string = "sendMediaGroup"

	// Use this method to send point on the map
	Location string = "sendLocation"

	// Use this method to send information about a venue.
	Venue string = "sendVenue"

	// Use this method to send phone contacts.
	Contact string = "sendContact"

	// Use this method to send a native poll.
	Poll string = "sendPoll"

	// Use this method to send an animated emoji that will display a random value
	Dice string = "sendDice"

	// Use this method when you need to tell the user that something is happening on
	// the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot,
	// Telegram clients clear its typing status)
	// 	Example: The an ImageBot needs some time to process a request and upload the image.
	// 	Instead of sending a text message along the lines of “Retrieving image, please wait…”,
	// 	the bot may use ChatAction method with IParameters.WriteAction(types.Action[1]). The user will see a
	// 	“sending photo” status for the bot.
	ChatAction string = "sendChatAction"

	// Use this method to change the chosen reactions on a message.
	// Service messages can't be reacted to. Automatically forwarded messages from a channel
	// to its discussion group have the same available reactions as messages in the channel.
	// Bots can't use paid reactions.
	MessageReaction string = "setMessageReaction"

	// Use this method to get a list of profile pictures for a user. The file
	// can then be downloaded via the link "https://api.telegram.org/file/bot<token>/<file_path>",
	// where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour.
	// Get the results here:
	// 	(IParameters).GetProfilePhotos()
	UserProfilePhotos string = "getUserProfilePhotos"

	// Changes the emoji status for a given user that previously allowed the bot to manage
	// their emoji status via the Mini App method https://core.telegram.org/bots/webapps#initializing-mini-apps
	UserEmojiStatus string = "setUserEmojiStatus"

	// Use this method to get basic information about a file and prepare it for downloading.
	// For the moment, bots can download files of up to 20MB in size. Get the results here:
	// 	(IParameters).GetFile()
	File string = "getFile"

	// Use this method to ban a user in a group, a supergroup or a channel.
	// In the case of supergroups and channels, the user will not be able to return
	// to the chat on their own using invite links, etc., unless unbanned first. The bot must be an
	// administrator in the chat for this to work and must have the appropriate administrator rights.
	BanMember string = "banChatMember"

	// Use this method to unban a previously banned user in a supergroup or channel. The user
	// will not return to the group or channel automatically, but will be able to join via link, etc.
	// The bot must be an administrator for this to work. By default, this method guarantees that after
	// the call the user is not a member of the chat, but will be able to join it. So if the user is a member
	// of the chat they will also be removed from the chat. If you don't want this, use
	// 	(IParameters).WriteOnlyIfBanned()
	UnbanMember string = "unbanChatMember"

	// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for
	// this to work and must have the appropriate administrator rights. Pass true for all fields in *types.ChatPermissions
	// that you put in (IParameters).WritePermissions() restrictions from a user.
	RestrictMember string = "restrictChatMember"

	// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator
	// in the chat for this to work and must have the appropriate administrator rights. Pass False for all fields in *types.ChatAdministratorRights
	// that you put in (IParameters).WriteAdministratorRights() parameters to demote a user.
	PromoteMember string = "promoteChatMember"

	// Use this method to set a custom title for an administrator in a supergroup promoted by the bot.
	ChatAdministratorTitle string = "setChatAdministratorCustomTitle"

	// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned,
	// the owner of the banned chat won't be able to send messages on behalf of any of their channels.
	// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights.
	BanSenderChat string = "banChatSenderChat"

	// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator
	// for this to work and must have the appropriate administrator rights.
	UnbanSenderChat string = "unbanChatSenderChat"

	// Use this method to set default chat permissions for all members.
	// The bot must be an administrator in the group or a supergroup for
	// this to work and must have the types.ChatAdministratorRights.CanRestrictMembers administrator rights.
	ChatPermissions string = "setChatPermissions"

	// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked.
	// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Get the results here:
	// 	(IParameters).GetInviteLink()
	ExportInviteLink string = "exportChatInviteLink"

	// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must
	// have the appropriate administrator rights. The link can be revoked using the method RevokeInviteLink. Get the results here:
	// 	(IParameters).GetNewInviteLink()
	CreateInviteLink string = "createChatInviteLink"

	// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work
	// and must have the appropriate administrator rights. Get the results here:
	// 	(IParameters).GetNewInviteLink()
	EditInviteLink string = "editChatInviteLink"

	// Use this method to create a subscription invite link for a channel chat. The bot must have the types.ChatAdministratorRights.CanInviteUsers administrator rights.
	// Get the results here:
	// 	(IParameters).GetNewInviteLink()
	CreateSubInviteLink string = "createChatSubscriptionInviteLink"

	// Use this method to edit a subscription invite link created by the bot. The bot must have the types.ChatAdministratorRights.CanInviteUsers administrator rights.
	// Get the results here:
	// 	(IParameters).GetNewInviteLink()
	EditSubInviteLink string = "editChatSubscriptionInviteLink"

	// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated.
	// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Get the results here:
	// 	(IParameters).GetNewInviteLink()
	RevokeInviteLink string = "revokeChatInviteLink"

	// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanInviteUsers administrator rights.
	ApproveJoinReq string = "approveChatJoinRequest"

	// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanInviteUsers administrator rights.
	DeclineJoinReq string = "declineChatJoinRequest"

	// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the
	// chat for this to work and must have the appropriate administrator rights.
	SetChatPhoto string = "setChatPhoto"

	// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this
	// to work and must have the appropriate administrator rights.
	DeleteChatPhoto string = "deleteChatPhoto"

	// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat
	// for this to work and must have the appropriate administrator rights.
	ChatTitle string = "setChatTitle"

	// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for
	// this to work and must have the appropriate administrator rights.
	ChatDescription string = "setChatDescription"

	// Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat,
	// the bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanPinMessages
	// administrator right in a supergroup or types.ChatAdministratorRights.CanEditMessages administrator right in a channel.
	PinMessage string = "pinChatMessage"

	// Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat,
	// the bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanPinMessages
	// administrator right in a supergroup or types.ChatAdministratorRights.CanEditMessages administrator right in a channel.
	UnpinMessage string = "unpinChatMessage"

	// Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat,
	// the bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanPinMessages
	// administrator right in a supergroup or types.ChatAdministratorRights.CanEditMessages administrator right in a channel.
	UnpinAll string = "unpinAllChatMessages"

	// Use this method for your bot to leave a group, supergroup or channel.
	LeaveChat string = "leaveChat"

	// Use this method to get up-to-date information about the chat. Get the results here:
	// 	(IChat).GetChat()
	GetChat string = "getChat"

	// Use this method to get a list of administrators in a chat, which aren't bots. Get the results here:
	// 	(IChat).GetMembers()
	GetAdmins string = "getChatAdministrators"

	// Use this method to get the number of members in a chat. Get the results here:
	// 	(IChat).GetMembersAmount()
	GetMemberCount string = "getChatMemberCount"

	// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the
	// bot is an administrator in the chat. Get the results here:
	// 	(IChat).GetMembers() [always zero element]
	GetMember string = "getChatMember"

	//Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat
	// for this to work and must have the appropriate administrator rights. Parameters:
	// 	(IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(Isticker).WriteSetName() [Required]
	ChatStickerSet string = "getChatStickerSet"

	// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat
	// for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	DeleteChatStickerSet string = "deleteChatStickerSet"

	// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IForum).WriteName()  [Required]
	// 	(formatter.IForum).WriteIconColor() [Optional]
	// 	(formatter.IForum).WriteIconEmojiID() [Optional]
	// Get the results here:
	// 	(formatter.IForum).GetForum()
	CreateForumTopic string = "createForumTopic"

	// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	// 	(formatter.IForum).WriteName()  [Optional]
	// 	(formatter.IForum).WriteIconEmojiID() [Optional]
	EditForumTopic string = "editForumTopic"

	// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	CloseForumTopic string = "closeForumTopic"

	// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	ReopenForumTopic string = "reopenForumTopic"

	// Use this method to delete a forum topic along with all its messages in a forum supergroup chat.
	// The bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanDeleteMessages administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	DeleteForumTopic string = "deleteForumTopic"

	// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this
	// to work and must have the types.ChatAdministratorRights.CanPinMessages administrator right in the supergroup. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	UnpinAllForumMessages string = "unpinAllForumTopicMessages"

	// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IForum).WriteName() [Required]
	EditGeneralForum string = "editGeneralForumTopic"

	// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	CloseGeneralForum string = "closeGeneralForumTopic"

	// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. The topic will be automatically unhidden if it was hidden.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	ReopenGGeneralForum string = "reopenGeneralForumTopic"

	// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. The topic will be automatically closed if it was open.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	HideGeneralForum string = "hideGeneralForumTopic"

	// Use this method to unhide the 'General' topic in a forum supergroup chat.The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	UnhideGeneralForum string = "unhideGeneralForumTopic"

	// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this
	// to work and must have the types.ChatAdministratorRights.CanPinMessages administrator right in the supergroup. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	UnpinAllGeneralForumMessages string = "unpinAllGeneralForumTopicMessages"

	// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. Parameters:
	// 	(formatter.IParameters).WriteCallbackQueryID() [Required]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteShowAlert() [Optional]
	// 	(formatter.IParameters).WriteURL() [Optional]
	// 	(formatter.IParameters).CacheTime() [Optional]
	AnswerCallbackQuery string = "answerCallbackQuery"

	// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// Get the results here:
	// 	(IChat).GetUsersBoosts()
	GetUsersBoosts string = "getUserChatBoosts"

	// Use this method to get information about the connection of the bot with a business account. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Required]
	// Get the results here:
	// 	(formatter.IChat).GetBusinessConnection()
	GetBusinessConnection string = "getBusinessConnection"

	// Use this method to change the list of the bot's commands. Parameters:
	// 	(formatter.IBot).WriteCommands() [Required]
	// 	(formatter.IBot).WriteScope() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	SetMyCommands string = "setMyCommands"

	// Use this method to delete the list of the bot's commands for the given scope and user language.
	// After deletion, higher level commands will be shown to affected users. Parameters:
	// 	(formatter.IBot).WriteScope() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	DeleteMyCommands string = "deleteMyCommands"

	// Use this method to get the current list of the bot's commands for the given scope and user language. Parameters:
	// 	(formatter.IBot).WriteScope() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IBot).GetCommands()
	GetMyCommands string = "getMyCommands"

	// Use this method to change the bot's name. Parameters:
	// 	(formatter.IBot).WriteName() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	SetMyName string = "setMyName"

	// Use this method to get the current bot name for the given user language. Parameters:
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IBot).GetName()
	GetMyName string = "getMyName"

	// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Parameters:
	// 	(formatter.IBot).WriteDescription() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	SetMyDescription string = "setMyDescription"

	// Use this method to get the current bot description for the given user language. Parameters:
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IBot).GetDescription()
	GetMyDescription string = "getMyDescription"

	// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot.
	// Parameters:
	// 	(formatter.IBot).WriteShortDescription() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	SetMyShortDescription string = "setMyShortDescription"

	// Use this method to get the current bot short description for the given user language. Parameters:
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IBot).GetShortDescription()
	GetMyShortDescription string = "getMyShortDescription"

	// Use this method to change the bot's menu button in a private chat, or the default menu button. Parameters:
	// 	(formatter.IChat).WriteChatID() [Optional]
	// 	(formatter.IBot).WriteMenuButton() [Optional]
	SetChatMenuButton string = "setChatMenuButton"

	// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Parameters:
	// 	(formatter.IChat).WriteChatID() [Optional]
	// Get the resutls here:
	// 	(formatter.IBot).GetMenuButton()
	GetChatMenuButton string = "getChatMenuButton"

	// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels.
	// These rights will be suggested to users, but they are free to modify the list before adding the bot. Parameters:
	// 	(formatter.IBot).WriteRights() [Optional]
	// 	(formatter.IBot).WriteForChannels() [Optional]
	SetMyDefaultAdminRights string = "setMyDefaultAdministratorRights"

	// Use this method to get the current default administrator rights of the bot. Parameters:
	// 	(formatter.IBot).WriteForChannels() [Optional]
	// Get the resutls here:
	// 	(formatter.IBot).GetAdminRights()
	GetMyDefaultAdministratorRights string = "getMyDefaultAdministratorRights"

	// Use this method to edit text and game messages.Note that business messages that were not sent by the bot and
	// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	(formatter.IParameters).WriteString() [Required]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteLinkPreviewOptions() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the resutls here:
	// 	(formatter.IParameters).GetMessage()
	EditText string = "editMessageText"

	// Use this method to edit captions of messages. Note that business messages that were not sent by the bot and do not
	// contain an inline keyboard can only be edited within 48 hours from the time they were sent. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	(formatter.IParameters).WriteCaption() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteShowCaptionAboveMedia() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the resutls here:
	// 	(formatter.IParameters).GetMessage()
	EditCaption string = "editMessageCaption"

	// Use this method to edit animation, audio, document, photo, or video messages, or to add media to
	// text messages. If a message is part of a message album, then it can be edited only to an audio for
	// audio albums, only to a document for document albums and to a photo or a video otherwise. When an
	// inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL.
	// Note that business messages that were not sent by the bot and do not contain an inline keyboard
	// can only be edited within 48 hours from the time they were sent. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	any media: (formatter.IPhoto), (formatter.IVideo), (formatter.IAudio), (formatter.IDocument) or (formatter.IAnimation) [Required]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the resutls here:
	// 	(formatter.IParameters).GetMessage()
	EditMedia string = "editMessageMedia"

	// Use this method to edit live location messages. A location can be edited until its (formatter.ILocation).WriteLivePeriod() expires or
	// editing is explicitly disabled by a call to method StopLiveLocation. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	(formatter.ILocation).WriteLatitude() [Required]
	// 	(formatter.ILocation).WriteLongitude() [Required]
	// 	(formatter.ILocation).WriteLivePeriod() [Optional]
	// 	(formatter.ILocation).WriteHorizontalAccuracy() [Optional]
	// 	(formatter.ILocation).WriteHeading() [Optional]
	// 	(formatter.ILocation).WriteProximityAlertRadius() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the resutls here:
	// 	(formatter.IParameters).GetMessage()
	EditLiveLocation string = "editMessageLiveLocation"

	// Use this method to stop updating a live location message before data you put in (formatter.ILocation).WriteLivePeriod() expires. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the resutls here:
	// 	(formatter.IParameters).GetMessage()
	StopLiveLocation string = "stopMessageLiveLocation"

	// Use this method to edit only the reply markup of messages. Note that business messages that were not sent by the bot and do
	// not contain an inline keyboard can only be edited within 48 hours from the time they were sent. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the resutls here:
	// 	(formatter.IParameters).GetMessage()
	EditReplyMarkup string = "editMessageReplyMarkup"

	// Use this method to stop a poll which was sent by the bot. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageID() [Required]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	StopPoll string = "stopPoll"

	// Use this method to delete a message, including service messages, with the following limitations:
	// 	- A message can only be deleted if it was sent less than 48 hours ago.
	// 	- Service messages about a supergroup, channel, or forum topic creation can't be deleted.
	// 	- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
	// 	- Bots can delete outgoing messages in private chats, groups, and supergroups.
	// 	- Bots can delete incoming messages in private chats.
	// 	- Bots granted types.ChatAdministratorRights.CanPostMessages permissions can delete outgoing messages in channels.
	// 	- If the bot is an administrator of a group, it can delete any message there.
	// 	- If the bot has types.ChatAdministratorRights.CanDeleteMessages permission in a supergroup or a channel, it can delete any message there.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageID() [Required]
	DeleteMessage string = "deleteMessage"

	// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageIDs() [Required]
	DeleteMessages string = "deleteMessages"

	// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.ISticker).WriteStickerStorage(), (formatter.ISticker).WriteStickerTelegram() or (formatter.ISticker).WriteStickerInternet() [Required]
	//	(formatter.ISticker).WriteAssociatedEmoji() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteReply(), (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IParameters).GetMessage()
	Sticker string = "sendSticker"

	// Use this method to get a sticker set. Parameters:
	// 	(formatter.ISticker).WriteSetName() [Required]
	// Get the result here:
	// 	(formatter.ISticker).GetSet()
	GetStickerSet string = "getStickerSet"

	// Use this method to get information about custom emoji stickers by their identifiers. Parameters:
	// 	(formatter.ISticker).WriteEmojiIDs() [Required]
	// Get the results here:
	// 	(formatter.ISticker).GetStickers()
	GetEmojiStickers string = "getCustomEmojiStickers"

	// Use this method to upload a file with a sticker for later use in the CreateNewStickerSet,
	// AddStickerToSet, or ReplaceStickerInSet methods (the file can be used multiple times). Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.ISticker).WriteStickerStorage() [Required]
	// 	(formatter.ISticker).WriteFormat() [Required]
	// Get the results here:
	// 	(formatter.IParameters).GetFile()
	UploadSticker string = "uploadStickerFile"

	// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Parameters:
	//	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.ISticker).WriteTitle() [Required]
	// 	(formatter.ISticker).WriteStickerType() [Optional]
	// 	(formatter.ISticker).WriteNeedsRepainting() [Optional]
	// It's possible to add 1-50 sticker a time. Please, every time create a new Sticker (formatter.Message).NewSticker() and specify these functions:
	// 	(formatter.ISticker).WriteStickerStorage(), (formatter.ISticker).WriteStickerTelegram() or (formatter.ISticker).WriteStickerInternet() [Required]
	// 	(formatter.ISticker).WriteFormat() [Optional]
	// 	(formatter.ISticker).WriteAssociatedEmojies() [Optional]
	// 	(formatter.ISticker).WriteMaskPosition() [Optional]
	// 	(formatter.ISticker).WriteKeywords() [Optional]
	CreateNewStickerSet string = "createNewStickerSet"

	// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200
	// stickers. Other sticker sets can have up to 120 stickers. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.ISticker).WriteStickerStorage(), (formatter.ISticker).WriteStickerTelegram() or (formatter.ISticker).WriteStickerInternet() [Required]
	// 	(formatter.ISticker).WriteFormat() [Optional]
	// 	(formatter.ISticker).WriteAssociatedEmojies() [Optional]
	// 	(formatter.ISticker).WriteMaskPosition() [Optional]
	// 	(formatter.ISticker).WriteKeywords() [Optional]
	AddStickerToSet string = "addStickerToSet"

	// Use this method to move a sticker in a set created by the bot to a specific position. Parameters:
	// 	(formatter.ISticker).WriteStickerTelegram() [Required]
	// 	(formatter.ISticker).WritePosition() [Required]
	SetStickerPositionInSet string = "setStickerPositionInSet"

	// Use this method to delete a sticker from a set created by the bot. Parameters:
	// 	(formatter.ISticker).WriteStickerTelegram() [Required]
	DeleteStickerFromSet string = "deleteStickerFromSet"

	// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent
	// to calling DeleteStickerFromSet, then AddStickerToSet, then SetStickerPositionInSet. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.ISticker).WriteStickerStorage(), (formatter.ISticker).WriteStickerTelegram() or (formatter.ISticker).WriteStickerInternet() [Required]
	// 	(formatter.ISticker).WriteFormat() [Optional]
	// 	(formatter.ISticker).WriteAssociatedEmojies() [Optional]
	// 	(formatter.ISticker).WriteMaskPosition() [Optional]
	// 	(formatter.ISticker).WriteKeywords() [Optional]
	ReplaceStickerInSet string = "replaceStickerInSet"

	// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker
	// must belong to a sticker set created by the bot. Parameters:
	// 	(formatter.ISticker).WriteStickerTelegram() [Required]
	// 	(formatter.ISticker).WriteAssociatedEmojies() [Required]
	SetStickerEmojiList string = "setStickerEmojiList"

	// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot.
	// Parameters:
	// 	(formatter.ISticker).WriteStickerTelegram() [Required]
	// 	(formatter.ISticker).WriteKeywords() [Optional]
	SetStickerKeywords string = "setStickerKeywords"

	// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Parameters:
	// 	(formatter.ISticker).WriteStickerTelegram() [Required]
	// 	(formatter.ISticker).WriteMaskPosition() [Optional]
	SetStickerMaskPosition string = "setStickerMaskPosition"

	// Use this method to set the title of a created sticker set. Parameters:
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.ISticker).WriteTitle() [Required]
	SetStickerSetTitle string = "setStickerSetTitle"

	// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must m
	// atch the format of the stickers in the set. Parameters:
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.ISticker).WriteThumbnailStorage(), 	(formatter.ISticker).WriteThumbnailTelegram() or (formatter.ISticker).WriteThumbnailInternet() [Optional]
	//	(formatter.ISticker).WriteThumbnailFormat() [Optional]
	SetStickerSetThumbnail string = "setStickerSetThumbnail"

	// Use this method to set the thumbnail of a custom emoji sticker set. Parameters:
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.ISticker).WriteEmojiID() [Optional]
	SetEmojiStickerSetThumbnail string = "setCustomEmojiStickerSetThumbnail"

	// Use this method to delete a sticker set that was created by the bot. Parameters:
	// 	(formatter.ISticker).WriteSetName()
	DeleteStickerSet string = "deleteStickerSet"

	// Sends a gift to the given user. The gift can't be converted to Telegram Stars by the user. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.ISticker).WriteGiftID() [Required]
	// 	(formatter.ISticker).WritePayForUpgrade() [Optional]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	Gift string = "sendGift"

	// Verifies a user on behalf of the organization which is represented by the bot. More: https://telegram.org/verify#third-party-verification. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteDescription() [Optional]
	VerifyUser string = "verifyUser"

	// Verifies a chat on behalf of the organization which is represented by the bot. More: https://telegram.org/verify#third-party-verification. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteDescription() [Optional]
	VerifyChat string = "verifyChat"

	// Removes verification from a user who is currently verified on behalf of the organization represented by the bot.
	// More: https://telegram.org/verify#third-party-verification. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	RemoveUserVerification string = "removeUserVerification"

	// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot.
	// More: https://telegram.org/verify#third-party-verification. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	RemoveChatVerification string = "removeChatVerification"

	// Use this method to send answers to an inline query. No more than 50 results per query are allowed. Parameters:
	// 	(formatter.IInlineMode).WriteQueryID() [Required]
	// 	(formatter.IInlineMode).WriteResults() [Required]
	// 	(formatter.IInlineMode).WriteCacheTime() [Optional]
	// 	(formatter.IInlineMode).WriteIsPersonal() [Optional]
	// 	(formatter.IInlineMode).WriteNextOffset() [Optional]
	// 	(formatter.IInlineMode).WriteButton() [Optional]
	AnswerInlineQuery string = "answerInlineQuery"

	// Use this method to set the result of an interaction with a Web App and send a corresponding message
	// on behalf of the user to the chat from which the query originated. More: https://core.telegram.org/bots/webapps. Parameters:
	// 	(formatter.IIlineMode).WriteWebAppQueryID() [Required]
	// 	(formatter.IInlineMode).WriteResult() [Required]
	// Get the results here:
	// 	(formatter.IInlineMode).GetSentWebAppMessage()
	AnswerWebAppQuery string = "answerWebAppQuery"

	// Stores a message that can be sent by a user of a Mini App. Parameters:
	//	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IInlineMode).WriteResult() [Required]
	// 	(formatter.IInlineMode).WriteAllowUserChats() [Optional]
	// 	(formatter.IInlineMode).WriteAllowBotChats() [Optional]
	// 	(formatter.IInlineMode).WriteAllowGroupChats() [Optional]
	// 	(formatter.IInlineMode).WriteAllowChannelChats() [Optional]
	// Get the results here:
	// 	(formatter.IInlineMode).GetPreparedInlineMessage()
	SavePreparedInlineMessage string = "savePreparedInlineMessage"

	// Use this method to send invoices. Parameters:
	// 	(IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IPayment).WriteTitle() [Required]
	// 	(formatter.IPayment).WriteDescription() [Required]
	// 	(formatter.IPayment).WritePayload() [Required]
	// 	(formatter.IPayment).WriteProviderToken() [Optional]
	// 	(formatter.IPayment).WriteCurrency() [Required]
	// 	(formatter.IPayment).WritePrices() [Required]
	// 	(formatter.IPayment).WriteMaxTipAmount() [Optional]
	// 	(formatter.IPayment).WriteSuggestedTipAmounts() [Optional]
	// 	(formatter.IPayment).WriteStartParameter() [Optional]
	// 	(formatter.IPayment).WritePhotoUrl() [Optional]
	// 	(formatter.IPayment).WritePhotoSize() [Optional]
	// 	(formatter.IPayment).WritePhotoWidth() [Optional]
	// 	(formatter.IPayment).WritePhotoHeight() [Optional]
	// 	(formatter.IPayment).WriteNeedPhoneNumber() [Optional]
	// 	(formatter.IPayment).WriteNeedEmail() [Optional]
	// 	(formatter.IPayment).WriteNeedShippingAddress() [Optional]
	// 	(formatter.IPayment).WriteSendPhoneNumberToProvider() [Optional]
	// 	(formatter.IPayment).WriteSendEmailToProvider() [Optional]
	// 	(formatter.IPayment).WriteIsFlexible() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteInline() [Optional]
	// Get the results here:
	// 	(formatter.IParameters).GetMessage()
	Invoice string = "sendInvoice"

	// Use this method to create a link for an invoice. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IPayment).WriteTitle() [Required]
	// 	(formatter.IPayment).WriteDescription() [Required]
	// 	(formatter.IPayment).WritePayload() [Required]
	// 	(formatter.IPayment).WriteProviderToken() [Optional]
	// 	(formatter.IPayment).WriteCurrency() [Required]
	// 	(formatter.IPayment).WritePrices() [Required]
	// 	(formatter.IPayment).WriteSubscriptionPeriod() [Optional]
	// 	(formatter.IPayment).WriteMaxTipAmount() [Optional]
	// 	(formatter.IPayment).WriteSuggestedTipAmounts() [Optional]
	// 	(formatter.IPayment).WriteProviderData() [Optional]
	// 	(formatter.IPayment).WritePhotoUrl() [Optional]
	// 	(formatter.IPayment).WritePhotoSize() [Optional]
	// 	(formatter.IPayment).WritePhotoWidth() [Optional]
	// 	(formatter.IPayment).WritePhotoHeight() [Optional]
	// 	(formatter.IPayment).WriteNeedPhoneNumber() [Optional]
	// 	(formatter.IPayment).WriteNeedEmail() [Optional]
	// 	(formatter.IPayment).WriteNeedShippingAddress() [Optional]
	// 	(formatter.IPayment).WriteSendPhoneNumberToProvider() [Optional]
	// 	(formatter.IPayment).WriteSendEmailToProvider() [Optional]
	// 	(formatter.IPayment).WriteIsFlexible() [Optional]
	// Get the results here:
	// 	(formatter.IPayment).GetInvoiceLink()
	CreateInvoiceLink string = "createInvoiceLink"

	//
	AnswerShippingQuery string = "answerShippingQuery"
)

var Media = map[string]struct{}{Photo: {}, Audio: {}, Document: {}, Video: {}, MediaGroup: {}, Animation: {}, Voice: {}}
