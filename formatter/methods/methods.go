package methods

const (
	// Use this method to send text messages. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IParameters).WriteString() [Required]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteLinkPreviewOptions() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Error()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Message string = "sendMessage"

	// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded.
	// Use usual function GetResponse() in interfaces you have mentioned to have response data from Telegram. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IChat).WriteFromChatID() [Required]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Required]
	// 	(formatter.IParameters).WriteVideoStartTimestamp() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Error()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	ForwardMessage string = "forwardMessage"

	// Use this method to forward multiple messages of any kind. If some of the specified
	// messages can't be found or forwarded, they are skipped. Service messages and messages
	// with protected content can't be forwarded. Album grouping is kept for forwarded
	// messagesUse this method to forward multiple messages of any kind. If some of the specified
	// messages can't be found or forwarded, they are skipped. Service messages and messages with
	// protected content can't be forwarded. Album grouping is kept for forwarded messages.
	// On success call (formatter.IParameters).GetMessageIDs() to have the response data from Telegram. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IChat).WriteFromChatID() [Required]
	// 	(formatter.IParameters).WriteMessageIDs() [Required]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Error()
	// 	(formatter.IGet).MessageIDs()
	ForwardMessages string = "forwardMessages"

	// Use this method to copy messages of any kind. Service messages, paid media messages,
	// giveaway messages, giveaway winners messages, and invoice messages can't be copied. Copying means that the bot copys
	// the message and then sends it to the chat with the unique identifier you put in (formatter.IChat).WriteChat{ID/Name}()
	// On success call (formatter.IParameters).GetMessageIDs() to have the response data from Telegram. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IChat).WriteFromChatID() [Required]
	// 	(formatter.IParameters).WriteMessageID() [Required]
	// 	(formatter.IParameters).WriteCaption() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteShowCaptionAboveMedia() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IParameters).WriteVideoStartTimestamp() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Error()
	// 	(formatter.IGet).MessageID()
	CopyMessage string = "copyMessage"

	// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied,
	// they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages,
	// and invoice messages can't be copied. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IChat).WriteFromChatID() [Required]
	// 	(formatter.IParameters).WriteMessageIDs() [Required]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteRemoveCaption() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Error()
	// 	(formatter.IGet).MessageIDs()
	CopyMessages string = "copyMessages"

	// Use this method to send photos. Might be used by default. Just mention a photo in Message structure.
	// Call GetResponse() in each interface you have created and gave to Message structure. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IPhoto).WritePhotoStorage(), (formatter.IPhoto).WritePhotoTelegram()  or (formatter.IPhoto).WritePhotoInternet() [Required]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteShowCaptionAboveMedia() [Optional]
	// 	(formatter.IPhoto).WriteHasSpoiler() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Photo()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Photo string = "sendPhoto"

	// Use this method to send audio files, if you want Telegram clients to display them in the music player.
	// Your audio must be in the .MP3 or .M4A format. On success, call GetResponse() in each interface you have
	// created and gave to Message structure. Bots can currently send audio files of up to 50 MB in size,
	// this limit may be changed in the future. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IAudio).WriteAudioStorage(), (formatter.IAudio).WriteAudioTelegram() or (formatter.IAudio).WriteAudioInternet() [Required]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IAudio).WriteDuration() [Optional]
	// 	(formatter.IAudio).WritePerformer() [Optional]
	// 	(formatter.IAudio).WriteTitle() [Optional]
	// 	(formatter.IAudio).WriteThumbnail() or (formatter.IAudio).WriteThumbnailID() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Audio()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Audio string = "sendAudio"

	// Use this method to send general files.
	// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IDocument).WriteDocumentStorage(), (formatter.IDocument).WriteDocumentTelegram() or (formatter.IDocument).WriteDocumentInternet() [Required]
	// 	(formatter.IDocument).WriteThumbnail() or (formatter.IDocument).WriteThumbnailID() [Optional]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IDocument).WriteDisableContentTypeDetection() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Document()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Document string = "sendDocument"

	// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document).
	// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IVideo).WriteVideoStorage(), (formatter.IVideo).WriteVideoTelegram() or (formatter.IVideo).WriteVideoInternet() [Required]
	// 	(formatter.IVideo).WriteDuration() [Optional]
	// 	(formatter.IVideo).WriteWidth() [Optional]
	// 	(formatter.IVideo).WriteHeight() [Optional]
	// 	(formatter.IVideo).WriteThumbnail() or (formatter.IVideo).WriteThumbnailID()  [Optional]
	// 	(formatter.IVideo).WriteCoverStorage(), (formatter.IVideo).WriteCoverTelegram() or (formatter.IVideo).WriteCoverInternet() [Optional]
	// 	(formatter.IParameters).WriteVideoStartTimestamp() [Optional]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteShowCaptionAboveMedia() [Optional]
	//	(formatter.IVideo).WriteHasSpoiler() [Optional]
	// 	(formatter.IVideo).WriteSupportsStreaming() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Video()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Video string = "sendVideo"

	// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
	// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IAnimation).WriteAnimationStorage(), (formatter.IAnimation).WriteAnimationTelegram() or (formatter.IAnimation).WriteAnimationInternet() [Required]
	// 	(formatter.IAnimation).WriteDuration() [Optional]
	// 	(formatter.IAnimation).WriteWidth() [Optional]
	// 	(formatter.IAnimation).WriteHeight() [Optional]
	// 	(formatter.IAnimation).WriteThumbnail() or (formatter.IAnimation).WriteThumbnailID()  [Optional]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteShowCaptionAboveMedia() [Optional]
	//	(formatter.IAnimation).WriteHasSpoiler() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Animation()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Animation string = "sendAnimation"

	// Use this method to send audio files, if you want Telegram clients to display the file
	// as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS,
	// or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document).
	// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IVoice).WriteVoiceStorage(), (formatter.IVoice).WriteVoiceTelegram() or (formatter.IVoice).WriteVoiceInternet() [Required]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IVoice).WriteDuration() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Voice()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	Voice string = "sendVoice"

	// Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IVideoNote).WriteVideoNoteStorage(), (formatter.IVideoNote).WriteVideoNoteTelegram() or (formatter.IVideoNote).WriteVideoNoteInternet() [Required]
	// 	(formatter.IVideoNote).WriteDuration() [Optional]
	// 	(formatter.IVideoNote).WriteLength() [Optional]
	// 	(formatter.IVideoNote).WriteThumbnail() or (formatter.IVideoNote).WriteThumbnailID() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).VideoNote()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	VideoNote string = "sendVideoNote"

	// Use this method to send paid media. Paid media can be only (formatter.IPhoto) or (formatter.IVideo). Also, the fields of methods.Photo
	// and methds.PaidMedia are diffirent. So, please make sure because all data that you might mention may be ignored
	// because it isn't allowed to send this with this method. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteStarCount() [Required]
	// 	(formatter.IParameters).WritePayload() [Optional]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteParseMode() [Optional]
	// 	(formatter.IParameters).WriteEntities() [Optional]
	// 	(formatter.IParameters).WriteShowCaptionAboveMedia() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Next a few parameters depend on what you're sending. If you're sending a photo:
	// 	(formatter.IPhoto).WritePhotoStorage(), (formatter.IPhoto).WritePhotoTelegram() or (formatter.IPhoto).WritePhotoInternet() [Required]
	// If you're sending a video:
	// 	(formatter.IVideo).WriteVideoStorage(), (formatter.IVideo).WriteVideoTelegram() or (formatter.IVideo).WriteVideoInternet() [Required]
	// 	(formatter.IVideo).WriteThumbnail() or (formatter.IVideo).WriteThumbnailID() [Optional]
	// 	(formatter.IVideo).WriteWidth() [Optional]
	// 	(formatter.IVideo).WriteHeight() [Optional]
	// 	(formatter.IVideo).WriteDuration() [Optional]
	// 	(formatter.IVideo).WriteSupportsStreaming() [Optional]
	// 	(formatter.IVideo).WriteCoverStorage(), (formatter.IVideo).WriteCoverTelegram() or (formatter.IVideo).WriteCoverInternet() [Optional]
	// 	(formatter.IParameters).WriteVideoStartTimestamp() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).PaidMedia()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Message() (all possible data included)
	PaidMedia string = "sendPaidMedia"

	// Use this method to send a group of photos, videos, documents or audios as an album.
	// Documents and audio files can be only grouped in an album with messages of the same type.
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// Next a few parameters depend on what you're going to send. If you're going to send photos:
	// 	(formatter.IPhoto).WritePhotoStorage(), (formatter.IPhoto).WritePhotoTelegram() or (formatter.IPhoto).WritePhotoInternet() [Required]
	// 	(formatter.IPhoto).WriteCaption() [Optional]
	// 	(formatter.IPhoto).WriteParseMode() [Optional]
	// 	(formatter.IPhoto).WriteCaptionEntities() [Optional]
	// 	(formatter.IPhoto).WriteShowCaptionAboveMedia() [Optional]
	// 	(formatter.IPhoto).WriteHasSpoiler() [Optional]
	// If you're going to send videos:
	//	(formatter.IVideo).WriteVideoStorage(), (formatter.IVideo).WriteVideoTelegram() or (formatter.IVideo).WriteVideoInternet() [Required]
	// 	(formatter.IVideo).WriteThumbnail() or (formatter.IVideo).WriteThumbnailID() [Optional]
	// 	(formatter.IVideo).WriteCaption() [Optional]
	// 	(formatter.IVideo).WriteParseMode() [Optional]
	// 	(formatter.IVideo).WriteCaptionEntities() [Optional]
	// 	(formatter.IVideo).WriteShowCaptionAboveMedia() [Optional]
	// 	(formatter.IVideo).WriteWidth() [Optional]
	// 	(formatter.IVideo).WriteHeight() [Optional]
	// 	(formatter.IVideo).WriteDuration() [Optional]
	// 	(formatter.IVideo).WriteSupportsStreaming() [Optional]
	// 	(formatter.IVideo).WriteHasSpoiler() [Optional]
	// If you're going to send documents:
	// 	(formatter.IDocument).WriteDocumentStorage(), (formatter.IDocument).WriteDocumentTelegram() or (formatter.IDocument).WriteDocumentInternet() [Required]
	// 	(formatter.IDocument).WriteThumbnail() or (formatter.IDocument).WriteThumbnailID() [Optional]
	// 	(formatter.IDocument).WriteCaption() [Optional]
	// 	(formatter.IDocument).WriteParseMode() [Optional]
	// 	(formatter.IDocument).WriteCaptionEntities() [Optional]
	// 	(formatter.IDocument).WriteDisableContentTypeDetection() [Optional]
	// If you're going to send audios:
	// 	(formatter.IAudio).WriteAudioStorage(), (formatter.IAudio).WriteAudioTelegram() or (formatter.IAudio).WriteAudioInternet() [Required]
	// 	(formatter.IAudio).WriteCaption() [Optional]
	// 	(formatter.IAudio).WriteParseMode() [Optional]
	// 	(formatter.IAudio).WriteCaptionEntities() [Optional]
	// 	(formatter.IAudio).WriteDuration() [Optional]
	// 	(formatter.IAudio).WritePerformer() [Optional]
	// 	(formatter.IAudio).WriteTitle() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageIDs()
	// 	(formatter.IGet).MediaGroupID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Photos()
	// 	(formatter.IGet).Videos()
	// 	(formatter.IGet).Audios()
	// 	(formatter.IGet).Documents()
	// 	(formatter.IGet).Replyed()
	// or
	// 	(formatter.IGet).Messages() (all possible data included)
	MediaGroup string = "sendMediaGroup"

	// Use this method to send point on the map. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.ILocation).WriteLatitude() [Required]
	// 	(formatter.ILocation).WriteLongitude() [Required]
	// 	(formatter.ILocation).WriteHorizontalAccuracy() [Optional]
	// 	(formatter.ILocation).WriteLivePeriod() [Optional]
	// 	(formatter.ILocation).WriteHeading() [Optional]
	// 	(formatter.ILocation).WriteProximityAlertRadius() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Replyed()
	Location string = "sendLocation"

	// Use this method to send information about a venue. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.ILocation).WriteLatitude() [Required]
	// 	(formatter.ILocation).WriteLongitude() [Required]
	// 	(formatter.ILocation).WriteTitle() [Required]
	// 	(formatter.ILocation).WriteAddress() [Required]
	// 	(formatter.ILocation).WriteFoursquareID() [Optional]
	// 	(formatter.ILocation).WriteFoursquareType() [Optional]
	// 	(formatter.ILocation).WriteGooglePlaceID() [Optional]
	// 	(formatter.ILocation).WriteGooglePlaceType() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Replyed()
	Venue string = "sendVenue"

	// Use this method to send phone contacts. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IContact).WritePhoneNumber() [Required]
	// 	(formatter.IContact).WriteFirstName() [Required]
	// 	(formatter.IContact).WriteLastName() [Optional]
	// 	(formatter.IContact).WriteVCard() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Replyed()
	Contact string = "sendContact"

	// Use this method to send a native poll. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IPoll).WriteQuestion() [Required]
	// 	(formatter.IPoll).WriteOptions() [Required]
	// 	(formatter.IPoll).WriteQuestionParseMode() [Optional]
	// 	(formatter.IPoll).WriteQuestionEntities() [Optional]
	//	(formatter.IPoll).WriteAnonymous() [Optional]
	// 	(formatter.IPoll).WriteType() [Optional]
	// 	(formatter.IPoll).WriteAllowMultipleAnswers() [Optional]
	// 	(formatter.IPoll).WriteCorrectOptionID() [Optional]
	// 	(formatter.IPoll).WriteExplanation() [Optional]
	// 	(formatter.IPoll).WriteExplanationParseMode() [Optional]
	// 	(formatter.IPoll).WriteExplanationEntities() [Optional]
	// 	(formatter.IPoll).WriteOpenPeriod() [Optional]
	// 	(formatter.IPoll).WriteCloseDate() [Optional]
	// 	(formatter.IPoll).WriteClosed() [Optional]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Poll()
	// 	(formatter.IGet).Replyed()
	Poll string = "sendPoll"

	// Use this method to send an animated emoji that will display a random value. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IParameters).WriteEmoji() [Required]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteReply() or (formatter.IKeyboard).WriteInline() or (formatter.IKeyboard).WriteForceReply() [Optional]
	// Get the results here:
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Dice()
	// 	(formatter.IGet).Replyed()
	Dice string = "sendDice"

	// Use this method when you need to tell the user that something is happening on
	// the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot,
	// Telegram clients clear its typing status)
	// 	Example: The an ImageBot needs some time to process a request and upload the image.
	// 	Instead of sending a text message along the lines of “Retrieving image, please wait…”,
	// 	the bot may use ChatAction method with (formatter.IParameters).WriteAction(types.Action[1]). The user will see a
	// 	“sending photo” status for the bot. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IParameters).WriteAction() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ChatAction string = "sendChatAction"

	// Use this method to change the chosen reactions on a message.
	// Service messages can't be reacted to. Automatically forwarded messages from a channel
	// to its discussion group have the same available reactions as messages in the channel.
	// Bots can't use paid reactions. Parameters:
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageID() [Required]
	// 	(formatter.IParameters).WriteReaction() [Optional]
	// 	(formatter.IParameters).WriteReactionIsBig() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	MessageReaction string = "setMessageReaction"

	// Use this method to get a list of profile pictures for a user. The file
	// can then be downloaded via the link "https://api.telegram.org/file/bot<token>/<file_path>",
	// where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteOffset() [Optional]
	// 	(formatter.IParameters).WriteLimit() [Optional]
	// Get the results here:
	// 	(formatter.IGet).ProfilePhotos()
	UserProfilePhotos string = "getUserProfilePhotos"

	// Changes the emoji status for a given user that previously allowed the bot to manage
	// their emoji status via the Mini App method https://core.telegram.org/bots/webapps#initializing-mini-apps. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteEmojiStatusCustomEmojiID() [Optional]
	// 	(formatter.IParameters).WriteEmojiStatusExpirationDate() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UserEmojiStatus string = "setUserEmojiStatus"

	// Use this method to get basic information about a file and prepare it for downloading.
	// For the moment, bots can download files of up to 20MB in size. Parameters:
	//	(formatter.IParameters).WriteFileID() [Required]
	// Get the results here:
	// 	(formatter.IGet).File()
	File string = "getFile"

	// Use this method to ban a user in a group, a supergroup or a channel.
	// In the case of supergroups and channels, the user will not be able to return
	// to the chat on their own using invite links, etc., unless unbanned first. The bot must be an
	// administrator in the chat for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteUntilDate() [Optional]
	// 	(formatter.IParameters).WriteRevokeMessages() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	BanMember string = "banChatMember"

	// Use this method to unban a previously banned user in a supergroup or channel. The user
	// will not return to the group or channel automatically, but will be able to join via link, etc.
	// The bot must be an administrator for this to work. By default, this method guarantees that after
	// the call the user is not a member of the chat, but will be able to join it. So if the user is a member
	// of the chat they will also be removed from the chat. If you don't want this, use
	// 	(formatter.IParameters).WriteOnlyIfBanned()
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteOnlyIfBanned() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnbanMember string = "unbanChatMember"

	// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for
	// this to work and must have the appropriate administrator rights. Pass true for all fields in *types.ChatPermissions
	// that you put in (formatter.IParameters).WritePermissions() restrictions from a user. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WritePermissions() [Required]
	// 	(formatter.IParameters).WriteIndependentChatPermissions() [Optional]
	// 	(formatter.IParameters).WriteUntilDate() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	RestrictMember string = "restrictChatMember"

	// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator
	// in the chat for this to work and must have the appropriate administrator rights. Pass False for all fields in *types.ChatAdministratorRights
	// that you put in (formatter.IParameters).WriteAdministratorRights() parameters to demote a user. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteAdministratorRights() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	PromoteMember string = "promoteChatMember"

	// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteCustomTitle() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ChatAdministratorTitle string = "setChatAdministratorCustomTitle"

	// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned,
	// the owner of the banned chat won't be able to send messages on behalf of any of their channels.
	// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IChat).WriteSenderChatID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	BanSenderChat string = "banChatSenderChat"

	// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator
	// for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IChat).WriteSenderChatID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnbanSenderChat string = "unbanChatSenderChat"

	// Use this method to set default chat permissions for all members.
	// The bot must be an administrator in the group or a supergroup for
	// this to work and must have the types.ChatAdministratorRights.CanRestrictMembers administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WritePermissions() [Required]
	// 	(formatter.IParameters).WriteIndependentChatPermissions() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ChatPermissions string = "setChatPermissions"

	// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked.
	// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IParameters).GetInviteLink()
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formamtter.IGet).String()
	ExportInviteLink string = "exportChatInviteLink"

	// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must
	// have the appropriate administrator rights. The link can be revoked using the method RevokeInviteLink. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.Ilink).WriteName() [Optional]
	// 	(formatter.Ilink).WriteExpireDate() [Optional]
	// 	(formatter.Ilink).WriteMemberLimit() [Optional]
	// 	(formatter.Ilink).WriteJoinRequest() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).InviteLink()
	CreateInviteLink string = "createChatInviteLink"

	// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work
	// and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.Ilink).WriteInviteLink() [Required]
	// 	(formatter.Ilink).WriteName() [Optional]
	// 	(formatter.Ilink).WriteExpireDate() [Optional]
	// 	(formatter.Ilink).WriteMemberLimit() [Optional]
	// 	(formatter.Ilink).WriteJoinRequest() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).InviteLink()
	EditInviteLink string = "editChatInviteLink"

	// Use this method to create a subscription invite link for a channel chat. The bot must have the types.ChatAdministratorRights.CanInviteUsers administrator rights.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.Ilink).WriteName() [Optional]
	// 	(formatter.ILink).WriteSubscriptionPeriod() [Required]
	// 	(formatter.ILink).WriteSubscriptionPrice() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).InviteLink()
	CreateSubInviteLink string = "createChatSubscriptionInviteLink"

	// Use this method to edit a subscription invite link created by the bot. The bot must have the types.ChatAdministratorRights.CanInviteUsers administrator rights.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.Ilink).WriteInviteLink() [Required]
	// 	(formatter.Ilink).WriteName() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).InviteLink()
	EditSubInviteLink string = "editChatSubscriptionInviteLink"

	// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated.
	// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.Ilink).WriteInviteLink() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).InviteLink()
	RevokeInviteLink string = "revokeChatInviteLink"

	// Use this method to approve a chat join request. The bot must be an administrator in the chat for this
	// to work and must have the types.ChatAdministratorRights.CanInviteUsers administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formmatter.IParameters).WriteUserID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ApproveJoinReq string = "approveChatJoinRequest"

	// Use this method to decline a chat join request. The bot must be an administrator in the
	// chat for this to work and must have the types.ChatAdministratorRights.CanInviteUsers administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formmatter.IParameters).WriteUserID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	DeclineJoinReq string = "declineChatJoinRequest"

	// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the
	// chat for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IPhoto).WritePhotoStorage() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetChatPhoto string = "setChatPhoto"

	// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this
	// to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	DeleteChatPhoto string = "deleteChatPhoto"

	// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat
	// for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// (formatter.IChat).WriteTitle() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ChatTitle string = "setChatTitle"

	// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for
	// this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IChat).WriteDescription() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ChatDescription string = "setChatDescription"

	// Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat,
	// the bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanPinMessages
	// administrator right in a supergroup or types.ChatAdministratorRights.CanEditMessages administrator right in a channel. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formmatter.IParameters).WriteMessageID() [Required]
	// 	(formatter.IParameters).WriteDisableNotification [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	PinMessage string = "pinChatMessage"

	// Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat,
	// the bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanPinMessages
	// administrator right in a supergroup or types.ChatAdministratorRights.CanEditMessages administrator right in a channel. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formmatter.IParameters).WriteMessageID() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnpinMessage string = "unpinChatMessage"

	// Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat,
	// the bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanPinMessages
	// administrator right in a supergroup or types.ChatAdministratorRights.CanEditMessages administrator right in a channel. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnpinAll string = "unpinAllChatMessages"

	// Use this method for your bot to leave a group, supergroup or channel. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	LeaveChat string = "leaveChat"

	// Use this method to get up-to-date information about the chat. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	//	(formatter.IGet).ChatInfo()
	GetChat string = "getChat"

	// Use this method to get a list of administrators in a chat, which aren't bots. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Members()
	GetAdmins string = "getChatAdministrators"

	// Use this method to get the number of members in a chat. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Integer()
	GetMemberCount string = "getChatMemberCount"

	// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the
	// bot is an administrator in the chat. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Members() [always zero element]
	GetMember string = "getChatMember"

	//Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat
	// for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.ISticker).WriteSetName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ChatStickerSet string = "setChatStickerSet"

	// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat
	// for this to work and must have the appropriate administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	DeleteChatStickerSet string = "deleteChatStickerSet"

	// No parameters. The results are here:
	// 	(formatter.IGet).Stickers()
	GetForumIconStickers string = "getForumTopicIconStickers"

	// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IForum).WriteName()  [Required]
	// 	(formatter.IForum).WriteIconColor() [Optional]
	// 	(formatter.IForum).WriteIconEmojiID() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet)..Forum()
	CreateForumTopic string = "createForumTopic"

	// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	// 	(formatter.IForum).WriteName()  [Optional]
	// 	(formatter.IForum).WriteIconEmojiID() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	EditForumTopic string = "editForumTopic"

	// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	CloseForumTopic string = "closeForumTopic"

	// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat
	// for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ReopenForumTopic string = "reopenForumTopic"

	// Use this method to delete a forum topic along with all its messages in a forum supergroup chat.
	// The bot must be an administrator in the chat for this to work and must have the types.ChatAdministratorRights.CanDeleteMessages administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	DeleteForumTopic string = "deleteForumTopic"

	// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this
	// to work and must have the types.ChatAdministratorRights.CanPinMessages administrator right in the supergroup. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnpinAllForumMessages string = "unpinAllForumTopicMessages"

	// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IForum).WriteName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	EditGeneralForum string = "editGeneralForumTopic"

	// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	CloseGeneralForum string = "closeGeneralForumTopic"

	// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. The topic will be automatically unhidden if it was hidden.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	ReopenGeneralForum string = "reopenGeneralForumTopic"

	// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. The topic will be automatically closed if it was open.
	// Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	HideGeneralForum string = "hideGeneralForumTopic"

	// Use this method to unhide the 'General' topic in a forum supergroup chat.The bot must be an administrator in
	// the chat for this to work and must have the types.ChatAdministratorRights.CanManageTopics administrator rights. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnhideGeneralForum string = "unhideGeneralForumTopic"

	// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this
	// to work and must have the types.ChatAdministratorRights.CanPinMessages administrator right in the supergroup. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	UnpinAllGeneralForumMessages string = "unpinAllGeneralForumTopicMessages"

	// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. Parameters:
	// 	(formatter.IParameters).WriteCallbackQueryID() [Required]
	// 	(formatter.IParameters).WriteString() [Optional]
	// 	(formatter.IParameters).WriteShowAlert() [Optional]
	// 	(formatter.IParameters).WriteURL() [Optional]
	// 	(formatter.IParameters).CacheTime() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	AnswerCallbackQuery string = "answerCallbackQuery"

	// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteUserID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Boosts()
	GetUsersBoosts string = "getUserChatBoosts"

	// Use this method to get information about the connection of the bot with a business account. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).BusinessConnection()
	GetBusinessConnection string = "getBusinessConnection"

	// Use this method to change the list of the bot's commands. Parameters:
	// 	(formatter.IBot).WriteCommands() [Required]
	// 	(formatter.IBot).WriteScope() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetMyCommands string = "setMyCommands"

	// Use this method to delete the list of the bot's commands for the given scope and user language.
	// After deletion, higher level commands will be shown to affected users. Parameters:
	// 	(formatter.IBot).WriteScope() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	DeleteMyCommands string = "deleteMyCommands"

	// Use this method to get the current list of the bot's commands for the given scope and user language. Parameters:
	// 	(formatter.IBot).WriteScope() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(fotmatter.IGet).Commands()
	GetMyCommands string = "getMyCommands"

	// Use this method to change the bot's name. Parameters:
	// 	(formatter.IBot).WriteName() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetMyName string = "setMyName"

	// Use this method to get the current bot name for the given user language. Parameters:
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).String()
	GetMyName string = "getMyName"

	// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Parameters:
	// 	(formatter.IBot).WriteDescription() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetMyDescription string = "setMyDescription"

	// Use this method to get the current bot description for the given user language. Parameters:
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).String()
	GetMyDescription string = "getMyDescription"

	// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot.
	// Parameters:
	// 	(formatter.IBot).WriteShortDescription() [Optional]
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetMyShortDescription string = "setMyShortDescription"

	// Use this method to get the current bot short description for the given user language. Parameters:
	// 	(formatter.IBot).WriteLanguage() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).String()
	GetMyShortDescription string = "getMyShortDescription"

	// Use this method to change the bot's menu button in a private chat, or the default menu button. Parameters:
	// 	(formatter.IChat).WriteChatID() [Optional]
	// 	(formatter.IBot).WriteMenuButton() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetChatMenuButton string = "setChatMenuButton"

	// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Parameters:
	// 	(formatter.IChat).WriteChatID() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).MenuButton()
	GetChatMenuButton string = "getChatMenuButton"

	// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels.
	// These rights will be suggested to users, but they are free to modify the list before adding the bot. Parameters:
	// 	(formatter.IBot).WriteRights() [Optional]
	// 	(formatter.IBot).WriteForChannels() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	SetMyDefaultAdminRights string = "setMyDefaultAdministratorRights"

	// Use this method to get the current default administrator rights of the bot. Parameters:
	// 	(formatter.IBot).WriteForChannels() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).AdminRights()
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
	// If the edited message is not an inline message call these to get the resutls:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Message()
	// otherwise:
	// 	(formatter.IGet).Status()
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
	// If the edited message is not an inline message call these to get the resutls:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Message()
	// otherwise:
	// 	(formatter.IGet).Status()
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
	// If the edited message is not an inline message call these to get the resutls:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Message()
	// otherwise:
	// 	(formatter.IGet).Status()
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
	// If the edited message is not an inline message call these to get the resutls:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Message()
	// otherwise:
	// 	(formatter.IGet).Status()
	EditLiveLocation string = "editMessageLiveLocation"

	// Use this method to stop updating a live location message before data you put in (formatter.ILocation).WriteLivePeriod() expires. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// If the edited message is not an inline message call these to get the resutls:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Message()
	// otherwise:
	// 	(formatter.IGet).Status()
	StopLiveLocation string = "stopMessageLiveLocation"

	// Use this method to edit only the reply markup of messages. Note that business messages that were not sent by the bot and do
	// not contain an inline keyboard can only be edited within 48 hours from the time they were sent. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// If the edited message is not an inline message call these to get the resutls:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Message()
	// otherwise:
	// 	(formatter.IGet).Status()
	EditReplyMarkup string = "editMessageReplyMarkup"

	// Use this method to stop a poll which was sent by the bot. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageID() [Required]
	// 	and a keyboard: (formatter.IKeyboard).WriteInline() [Optional]
	// Get the results here:
	// 	(formatter.IGet).Status()
	// 	(formatter.IGet).Poll()
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
	// Get the results here:
	// 	(formatter.IGet).Status()
	DeleteMessage string = "deleteMessage"

	// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Parameters:
	// 	(formatter.IChat).WriteChatID() or (formatter.IChat).WriteChatName() [Required]
	// 	(formatter.IParameters).WriteMessageIDs() [Required]
	// Get the results here:
	// 	(formatter.IGet).Status()
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
	// 	(formatter.IGet).MessageID()
	// 	(formatter.IGet).Sender()
	// 	(formatter.IGet).Chat()
	// 	(formatter.IGet).Date()
	// 	(formatter.IGet).Sticker()
	// 	(formatter.IGet).Replyed()
	// or get the all results:
	// 	(formatter.IGet).Message()
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
	// 	(formatter.ISticker).WriteThumbnail() or (formatter.ISticker).WriteThumbnailID() [Optional]
	//	(formatter.ISticker).WriteThumbnailFormat() [Optional]
	SetStickerSetThumbnail string = "setStickerSetThumbnail"

	// Use this method to set the thumbnail of a custom emoji sticker set. Parameters:
	// 	(formatter.ISticker).WriteSetName() [Required]
	// 	(formatter.ISticker).WriteEmojiID() [Optional]
	SetEmojiStickerSetThumbnail string = "setCustomEmojiStickerSetThumbnail"

	// Use this method to delete a sticker set that was created by the bot. Parameters:
	// 	(formatter.ISticker).WriteSetName()
	DeleteStickerSet string = "deleteStickerSet"

	// Has no parameters. Get the results here:
	// 	(formatter.IGet).Gifts()
	GetAvailableGifts string = "getAvailableGifts"

	// Sends a gift to the given user. The gift can't be converted to Telegram Stars by the user. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IChat).WriteChatID or (formatter.IChat).WriteChatName() [Optional]
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

	// If you sent an invoice requesting a shipping address and the parameter formatter.IPayment).WriteIsFlexible() was specified,
	// the Bot API will send an (formatter.IPayment).GetUpdate with a shipping_query field to the bot. Use this method to reply to shipping queries.
	// Parameters:
	// 	(formatter.IPayment).WriteShippingID() [Required]
	// 	(formatter.IPayment).WriteOK() [Required]
	// 	(formatter.IPayment).WriteShippingOptions() [Optional]
	// 	(formatter.IPayment).WriteErrorMessage() [Optional]
	AnswerShippingQuery string = "answerShippingQuery"

	// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of
	// a (formatter.IPayment).GetUpdate with the field PreCheckoutQuery. Use this method to respond to such pre-checkout queries. Note: The Bot API must
	// receive an answer within 10 seconds after the pre-checkout query was sent. Parameters:
	// 	(formatter.IPayment).WritePreCheckoutID() [Required]
	// 	(formatter.IPayment).WriteOK() [Required]
	// 	(formatter.IIPayment).WriteErrorMessage() [Optional]
	AnswerPreCheckoutQuery string = "answerPreCheckoutQuery"

	// Returns the bot's Telegram Star transactions in chronological order. Parameters:
	// 	(formatter.IPayment).WriteLimit() [Optional]
	// 	(formatter.IPaymentt).WriteOffset() [Optional]
	// Get the results here:
	// 	(formatter.IPayment).GetStarTransactions()
	GetStarTransactions string = "getStarTransactions"

	// Refunds a successful payment in Telegram Stars. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IPayment).WriteTelegramPaymentChargeID() [Required]
	RefundStarPayment string = "refundStarPayment"

	// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IPayment).WriteTelegramPaymentChargeID() [Required]
	// 	(formatter.IPayment).WriteIsCanceled() [Required]
	EditUserStarSubscription string = "editUserStarSubscription"

	// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit
	// their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Use this if the
	// data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if
	// a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some
	// details in the error message to make sure the user knows how to correct the issues. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IParameters).WriteErrors() [Required]
	SetPassportDataErrors string = "setPassportDataErrors"

	// Use this method to send a game. Parameters:
	// 	(formatter.IChat).WriteBusinessConnectionID() [Optional]
	// 	(formatter.IChat).WriteChatID [Required]
	// 	(formatter.IParameters).WriteMessageThreadID() [Optional]
	// 	(formatter.IGame).WriteShortName() [Required]
	// 	(formatter.IParameters).WriteDisableNotification() [Optional]
	// 	(formatter.IParameters).WriteProtectContent() [Optional]
	// 	(formatter.IParameters).WriteAllowPaidBroadcast() [Optional]
	// 	(formatter.IParameters).WriteMessageEffectID() [Optional]
	// 	(formatter.IParameters).WriteReplyParameters() [Optional]
	// 	(formatter.IKeyboard).WriteInline() [Optional]
	Game string = "sendGame"

	// Use this method to set the score of the specified user in a game message. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IGame).WriteScore() [Required]
	// 	(formatter.IGame).WriteForce() [Optional]
	// 	(formatter.IGame).WriteDisableEditMessage() [Optional]
	// 	(formatter.IChat).WriteChatID() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	//	(formatter.IParameters).WriteInlineMessageID() [Optional]
	SetGameScore string = "setGameScore"

	// Use this method to get data for high score tables. Will return the score of the
	// specified user and several of their neighbors in a game. This method will currently
	// return scores for the target user, plus two of their closest neighbors on each side. Will also
	// return the top three users if the user and their neighbors are not among them. Please note that this
	// behavior is subject to change. Parameters:
	// 	(formatter.IParameters).WriteUserID() [Required]
	// 	(formatter.IChat).WriteChatID() [Optional]
	// 	(formatter.IParameters).WriteMessageID() [Optional]
	// 	(formatter.IParameters).WriteInlineMessageID() [Optional]
	GetGameHighScores string = "getGameHighScores"
)

var Lvl1 = []string{Message, Photo, Audio,
	Document, Video, Animation,
	Voice, VideoNote, PaidMedia,
	MediaGroup, Location, Venue,
	Contact, Poll, Dice,
	ChatAction, Sticker, Gift,
	Invoice, Game}
