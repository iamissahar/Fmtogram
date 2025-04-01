package fmtogram

type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`
	Chat                    *Chat
	BoostCount              int
	BusinessBot             *User
}

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	IsForum   bool   `json:"is_forum,omitempty"`
}

type SharedUsers struct {
	UserID    int          `json:"user_id"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Username  string       `json:"username,omitempty"`
	Photo     []*PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestID int            `json:"request_id"`
	Users     []*SharedUsers `json:"users"`
}

type ChatShared struct {
	RequestID int          `json:"request_id"`
	ChatID    int          `json:"chat_id"`
	Title     string       `json:"title,omitempty"`
	Username  string       `json:"username,omitempty"`
	Photo     []*PhotoSize `json:"photo,omitempty"`
}

type SharedUser struct {
	UserID    int          `json:"user_id"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Username  string       `json:"username,omitempty"`
	Photo     []*PhotoSize `json:"photo,omitempty"`
}

type BackgroundTypeFill struct {
	Type             string          `json:"type"`
	Fill             *BackgroundFill `json:"fill"`
	DarkThemeDimming int             `json:"dark_theme_dimming,omitempty"`
}

type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"`
	Document         *Document `json:"document"`
	DarkThemeDimming int       `json:"dark_theme_dimming,omitempty"`
	IsBlurred        bool      `json:"is_blurred,omitempty"`
	IsMoving         bool      `json:"is_moving,omitempty"`
}

type BackgroundTypePattern struct {
	Type       string          `json:"type"`
	Document   *Document       `json:"document"`
	Fill       *BackgroundFill `json:"fill"`
	Intensity  int             `json:"intensity"`
	IsInverted bool            `json:"is_inverted,omitempty"`
	IsMoving   bool            `json:"is_moving,omitempty"`
}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

type BackgroundType struct {
	BackgroundTypeFill      *BackgroundTypeFill
	BackgroundTypeWallpaper *BackgroundTypeWallpaper
	BackgroundTypePattern   *BackgroundTypePattern
	BackgroundTypeChatTheme *BackgroundTypeChatTheme
}

type ChatBackground struct {
	Type *BackgroundType `json:"type"`
}

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type Birthdate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     *bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 *User  `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int    `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
	SubscriptionPeriod      int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       int    `json:"subscription_price,omitempty"`
}

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   *User  `json:"user"`
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   *User  `json:"user"`
}

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageID int    `json:"giveaway_message_id"`
	User              *User  `json:"user,omitempty"`
	IsUnclaimed       *bool  `json:"is_unclaimed,omitempty"`
}

type ChatBoostSource struct {
	ChatBoostSourcePremium  *ChatBoostSourcePremium
	ChatBoostSourceGiftCode *ChatBoostSourceGiftCode
	ChatBoostSourceGiveaway *ChatBoostSourceGiveaway
}

type ChatBoost struct {
	BoostID        string           `json:"boost_id"`
	AddDate        int              `json:"add_date"`
	CxpirationDate int              `json:"expiration_date"`
	Source         *ChatBoostSource `json:"source"`
}

type ChatBoostUpdated struct {
	Chat  *Chat      `json:"chat"`
	Boost *ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       *Chat            `json:"chat"`
	BoostID    string           `json:"boost_id"`
	RemoveDate int              `json:"remove_date"`
	Source     *ChatBoostSource `json:"source"`
}

type UserChatBoosts struct {
	Boosts []*ChatBoost `json:"boosts"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

type ChatAdministratorRights struct {
	IsAnonymous         *bool `json:"is_anonymous,omitempty"`
	CanManageChat       *bool `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   *bool `json:"can_delete_messages,omitempty"`
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  *bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   *bool `json:"can_promote_members,omitempty"`
	CanChangeInfo       *bool `json:"can_change_info,omitempty"`
	CanInviteUsers      *bool `json:"can_invite_users,omitempty"`
	CanPostStories      *bool `json:"can_post_stories,omitempty"`
	CanEditStories      *bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    *bool `json:"can_delete_stories,omitempty"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
}

type ChatMember struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsAnonymous           *bool  `json:"is_anonymous,omitempty"`
	CustomTitle           string `json:"custom_title,omitempty"`
	CanBeEdited           *bool  `json:"can_be_edited,omitempty"`
	CanManageChat         *bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages     *bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats   *bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers    *bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     *bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo         *bool  `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool  `json:"can_invite_users,omitempty"`
	CanPostStories        *bool  `json:"can_post_stories,omitempty"`
	CanEditStories        *bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories      *bool  `json:"can_delete_stories,omitempty"`
	CanPostMessages       *bool  `json:"can_post_messages,omitempty"`
	CanEditMessages       *bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages        *bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool  `json:"can_manage_topics,omitempty"`
	IsMember              *bool  `json:"is_member,omitempty"`
	CanSendMessages       *bool  `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool  `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool  `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool  `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool  `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool  `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     *bool  `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool  `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool  `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews,omitempty"`
	UntilDate             *int   `json:"until_date,omitempty"`
}

type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`
	From       *User           `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int             `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`
	From                    *User           `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           *ChatMember     `json:"old_chat_member"`
	NewChatMember           *ChatMember     `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"`
}

type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              string                `json:"title,omitempty"`
	Username                           string                `json:"username,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	AccentColorID                      int64                 `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	AvailableReactions                 []*ReactionType       `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               int64                 `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	Description                        string                `json:"description,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
	CanSendGift                        bool                  `json:"can_send_gift,omitempty"`
}

type PaidMediaPurchased struct {
	From *User  `json:"from"`
	Pl   string `json:"paid_media_payload"`
}

type Update struct {
	UpdateID               int                          `json:"update_id"`
	Message                *TelegramMessage             `json:"message"`
	EditedMsg              *TelegramMessage             `json:"edited_message"`
	ChanPost               *TelegramMessage             `json:"channel_post"`
	EditedChanPost         *TelegramMessage             `json:"edited_channel_post"`
	BusinessConn           *BusinessConnection          `json:"business_connection"`
	BusinessMsg            *TelegramMessage             `json:"business_message"`
	EditedBusinessMsg      *TelegramMessage             `json:"edited_business_message"`
	DeletedBusinessMessage *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	MessageR               *MessageReactionUpdated      `json:"message_reaction"`
	MessageRcount          *MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQ                *InlineQuery                 `json:"inline_query"`
	ChosenInlineR          *ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQ              *CallbackQuery               `json:"callback_query"`
	ShippingQ              *ShippingQuery               `json:"shipping_query"`
	PreCheckoutQ           *PreCheckoutQuery            `json:"pre_checkout_query"`
	PaidMedia              *PaidMediaPurchased          `json:"purchased_paid_media"`
	PollUpd                *Poll                        `json:"poll"`
	PollAnswr              *PollAnswer                  `json:"poll_answer"`
	MyChatMem              *ChatMemberUpdated           `json:"my_chat_member"`
	ChatMem                *ChatMemberUpdated           `json:"chat_member"`
	ChatJoinReq            *ChatJoinRequest             `json:"chat_join_request"`
	ChatBoostUpd           *ChatBoostUpdated            `json:"chat_boost"`
	RemovedChatBoostUpd    *ChatBoostRemoved            `json:"removed_chat_boost"`
}

type Telegram struct {
	Ok     bool      `json:"ok"`
	Result []*Update `json:"result"`
}

type ChatFullInfoResponse struct {
	Ok     bool          `json:"ok"`
	Result *ChatFullInfo `json:"result"`
}

type ChatMemberResponse struct {
	Ok     bool        `json:"ok"`
	Result *ChatMember `json:"result"`
}

type ChatMembersResponse struct {
	Ok     bool          `json:"ok"`
	Result []*ChatMember `json:"result"`
}

type StickersResponse struct {
	Ok     bool       `json:"ok"`
	Result []*Sticker `json:"result"`
}

type ForumResponse struct {
	Ok     bool        `json:"ok"`
	Result *ForumTopic `json:"result"`
}

type UserBoostsResponse struct {
	Ok     bool            `json:"ok"`
	Result *UserChatBoosts `json:"result"`
}

type BusinessConResponse struct {
	Ok     bool                `json:"ok"`
	Result *BusinessConnection `json:"result"`
}

type BotCommandResponse struct {
	Ok     bool          `json:"ok"`
	Result []*BotCommand `json:"result"`
}

type BotNameResponse struct {
	Ok     bool     `json:"ok"`
	Result *BotName `json:"result"`
}

type BotDescriptionResponse struct {
	Ok     bool            `json:"ok"`
	Result *BotDescription `json:"result"`
}

type BotShorDescriptionResponse struct {
	Ok     bool                 `json:"ok"`
	Result *BotShortDescription `json:"result"`
}

type MenuButtonResponse struct {
	Ok     bool        `json:"ok"`
	Result *MenuButton `json:"result"`
}

type AdminRightResponse struct {
	Ok     bool                     `json:"ok"`
	Result *ChatAdministratorRights `json:"result"`
}

type PollResponse struct {
	Ok     bool  `json:"ok"`
	Result *Poll `json:"result"`
}

type StickerSetResponse struct {
	Ok     bool        `json:"ok"`
	Result *StickerSet `json:"result"`
}

type GiftsResponse struct {
	Ok     bool   `json:"ok"`
	Result *Gifts `json:"result"`
}

type WepAppMsgResponse struct {
	Ok     bool               `json:"ok"`
	Result *SentWebAppMessage `json:"result"`
}

type PreparedInlineMessageResponse struct {
	Ok     bool                   `json:"ok"`
	Result *PreparedInlineMessage `json:"result"`
}

type StringResponse struct {
	Ok     bool   `json:"ok"`
	Result string `json:"result"`
}

type StarTransactionResponse struct {
	Ok     bool             `json:"ok"`
	Result *StarTransaction `json:"result"`
}

type GameHighScoresResponse struct {
	Ok     bool             `json:"ok"`
	Result []*GameHighScore `json:"result"`
}

type Error struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type GetMee struct {
	Ok     bool  `json:"ok"`
	Result *User `json:"result,omitempty"`
}

type MessageResponse struct {
	Ok     bool             `json:"ok"`
	Result *TelegramMessage `json:"result,omitempty"`
}

type MediaGroupResponse struct {
	Ok     bool               `json:"ok"`
	Result []*TelegramMessage `json:"result,omitempty"`
}

type MessageIDResponse struct {
	Ok     bool       `json:"ok"`
	Result *MessageID `json:"result,omitempty"`
}

type MessageIDsResponse struct {
	Ok     bool         `json:"ok"`
	Result []*MessageID `json:"result,omitempty"`
}

type SimpleResponse struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type UserProfilePhotosResponse struct {
	Ok     bool               `json:"ok"`
	Result *UserProfilePhotos `json:"result"`
}

type FileResponse struct {
	Ok     bool  `json:"ok"`
	Result *File `json:"result"`
}

type InviteLinkResponse struct {
	Ok     bool            `json:"ok"`
	Result *ChatInviteLink `json:"result"`
}

type IntResponse struct {
	Ok     bool `json:"ok"`
	Result int  `json:"result"`
}

type WebhookResponse struct {
	Ok     bool         `json:"ok"`
	Result *WebhookInfo `json:"result"`
}

type KeyboardButtonRequestUsers struct {
	RequestID       int  `json:"request_id"`
	UserIsBot       bool `json:"user_is_bot,omitempty"`
	UserIsPremium   bool `json:"user_is_premium,omitempty"`
	MaxQuantity     int  `json:"max_quantity,omitempty"`
	RequestName     bool `json:"request_name,omitempty"`
	RequestUsername bool `json:"request_username,omitempty"`
	RequestPhoto    bool `json:"request_photo,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequestID               int                      `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`
	RequestTitle            bool                     `json:"request_title,omitempty"`
	RequestUsername         bool                     `json:"request_username,omitempty"`
	RequestPhoto            bool                     `json:"request_photo,omitempty"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`
	IsPersistent          bool                `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool                `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string              `json:"input_field_placeholder,omitempty"`
	Selective             bool                `json:"selective,omitempty"`
}

type ShippingOption struct {
	ID     string          `json:"id"`
	Title  string          `json:"title"`
	Prices []*LabeledPrice `json:"prices"`
}

type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	Date int    `json:"date"`
	Url  string `json:"url"`
}

type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

type RevenueWithdrawalState struct {
	*RevenueWithdrawalStatePending
	*RevenueWithdrawalStateSucceeded
	*RevenueWithdrawalStateFailed
}

type AffiliateInfo struct {
	User               User `json:"affiliate_user"`
	Chat               Chat `json:"affiliate_chat"`
	CommissionPerMille int  `json:"commission_per_mille"`
	Amount             int  `json:"amount"`
	NanostarAmount     int  `json:"nanostar_amount"`
}

type TransactionPartnerUser struct {
	Type               string         `json:"type"`
	User               User           `json:"user"`
	Affiliate          *AffiliateInfo `json:"affiliate,omitempty"`
	InvoicePayload     string         `json:"invoice_payload,omitempty"`
	SubscriptionPeriod int            `json:"subscription_period,omitempty"`
	PaidMedia          []PaidMedia    `json:"paid_media,omitempty"`
	PaidMediaPayload   string         `json:"paid_media_payload,omitempty"`
	Gift               *Gift          `json:"gift,omitempty"`
}

type TransactionPartnerChat struct {
	Type string `json:"type"`
	Chat *Chat  `json:"chat"`
	Gift *Gift  `json:"gift"`
}

type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`
	SponsorUser        *User  `json:"sponsor_user,omitempty"`
	CommissionPerMille int    `json:"commission_per_mille"`
}

type TransactionPartnerFragment struct {
	Type            string                  `json:"type"`
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}

type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`
	RequestCount int    `json:"request_count"`
}

type TransactionPartnerOther struct {
	Type string `json:"type"`
}

type TransactionPartner struct {
	*TransactionPartnerUser
	*TransactionPartnerChat
	*TransactionPartnerAffiliateProgram
	*TransactionPartnerFragment
	*TransactionPartnerTelegramAds
	*TransactionPartnerTelegramApi
	*TransactionPartnerOther
}

type StarTransaction struct {
	ID             string              `json:"id"`
	Amount         int                 `json:"amount"`
	NanostarAmount *int                `json:"nanostar_amount,omitempty"`
	Date           int                 `json:"date"`
	Source         *TransactionPartner `json:"source,omitempty"`
	Receiver       *TransactionPartner `json:"receiver,omitempty"`
}

type StarTransactions struct {
	ST []*StarTransaction `json:"transactions"`
}

type PassportElementError struct {
	Source      string   `json:"source,omitempty"`
	Type        string   `json:"type,omitempty"`
	FieldName   string   `json:"field_name,omitempty"`
	DataHash    *string  `json:"data_hash,omitempty"`
	ElementHash *string  `json:"element_hash,omitempty"`
	FileHashes  []string `json:"file_hashes,omitempty"`
	Message     string   `json:"message,omitempty"`
}

type InaccessibleMessage struct {
	Chat      *Chat `json:"chat"`
	MessageID int   `json:"message_id"`
	Date      int   `json:"date"`
}

type MessageID struct {
	MessageID int `json:"message_id"`
}

type MessageReactionCountUpdated struct {
	Chat      *Chat            `json:"chat"`
	MessageID int              `json:"message_id"`
	Date      int              `json:"date"`
	Reactions []*ReactionCount `json:"reactions"`
}

type MessageReactionUpdated struct {
	Chat        *Chat           `json:"chat"`
	MessageID   int             `json:"message_id"`
	User        *User           `json:"user,omitempty"`
	ActorChat   *Chat           `json:"actor_chat,omitempty"`
	Date        int             `json:"date"`
	OldReaction []*ReactionType `json:"old_reaction"`
	NewReaction []*ReactionType `json:"new_reaction"`
}

type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       int    `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           int    `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

type MessageOriginChat struct {
	Type            string `json:"type"`
	Date            int    `json:"date"`
	SenderChat      *Chat  `json:"sender_chat"`
	AuthorSignature string `json:"author_signature"`
}

type MessageOriginChannel struct {
	Type            string `json:"type"`
	Date            int    `json:"date"`
	Chat            *Chat  `json:"chat"`
	MessageID       int    `json:"message_id"`
	AuthorSignature string `json:"author_signature"`
}

type MessageOrigin struct {
	*MessageOriginUser
	*MessageOriginHiddenUser
	*MessageOriginChat
	*MessageOriginChannel
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url"`
	User          *User  `json:"user"`
	Language      string `json:"language"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type TelegramMessage struct {
	MessageID                     int                            `json:"message_id"`
	MessageThreadID               int                            `json:"message_thread_id"`
	From                          *User                          `json:"from"`
	SenderChat                    *Chat                          `json:"sender_chat"`
	SenderBoostCount              int                            `json:"sender_boost_count"`
	SenderBusinessBot             *User                          `json:"sender_business_bot"`
	Date                          int                            `json:"date"`
	BusinessConnectionID          string                         `json:"business_connection_id"`
	Chat                          *Chat                          `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin"`
	IsTopicMessage                bool                           `json:"is_topic_message"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward"`
	ReplyToMessage                *TelegramMessage               `json:"reply_to_message"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply"`
	Quote                         *TextQuote                     `json:"quote"`
	ReplyToStory                  *Story                         `json:"reply_to_story"`
	ViaBot                        *User                          `json:"via_bot"`
	EditDate                      int                            `json:"edit_date"`
	HasProtectedContent           bool                           `json:"has_protected_content"`
	IsFrommOffline                bool                           `json:"is_from_offline"`
	MediaGroupID                  string                         `json:"media_group_id"`
	AuthorSignature               string                         `json:"author_signature"`
	Text                          string                         `json:"text"`
	Entities                      []*MessageEntity               `json:"entities"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options"`
	EffectID                      string                         `json:"effect_id"`
	Animation                     *Animation                     `json:"animation"`
	Audio                         *Audio                         `json:"audio"`
	Document                      *Document                      `json:"document"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media"`
	Photo                         []*PhotoSize                   `json:"photo"`
	Sticker                       *Sticker                       `json:"sticker"`
	Story                         *Story                         `json:"story"`
	Video                         *Video                         `json:"video"`
	VideoNote                     *VideoNote                     `json:"video_note"`
	Voice                         *Voice                         `json:"voice"`
	Caption                       string                         `json:"caption"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler"`
	Contact                       *Contact                       `json:"contact"`
	Dice                          *Dice                          `json:"dice"`
	Game                          *Game                          `json:"game"`
	Poll                          *Poll                          `json:"poll"`
	Venue                         *Venue                         `json:"venue"`
	Location                      *Location                      `json:"location"`
	NewChatMembers                []*User                        `json:"new_chat_members"`
	LeftChatMember                *User                          `json:"left_chat_member"`
	NewChatTitle                  string                         `json:"new_chat_title"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`
	GroupChatCreated              bool                           `json:"group_chat_created"`
	SuperGroupChatCreated         bool                           `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                           `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               int                            `json:"migrate_from_chat_id"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message"`
	Invoice                       *Invoice                       `json:"invoice"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment"`
	UsersShared                   *UsersShared                   `json:"users_shared"`
	ChatShared                    *ChatShared                    `json:"chat_shared"`
	ConnectedWebsite              string                         `json:"connected_website"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  *PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created"`
	Giveaway                      *Giveaway                      `json:"giveaway"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    *WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
}

type MaybeInaccessibleMessage struct {
	Message             *Message
	InaccessibleMessage *InaccessibleMessage
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Video struct {
	FileID         string       `json:"file_id"`
	FileUniqueID   string       `json:"file_unique_id"`
	Width          int          `json:"width"`
	Height         int          `json:"height"`
	Duration       int          `json:"duration"`
	Thumbnail      *PhotoSize   `json:"thumbnail,omitempty"`
	FileName       string       `json:"file_name,omitempty"`
	MimeType       string       `json:"mime_type,omitempty"`
	FileSize       int          `json:"file_size,omitempty"`
	Cover          []*PhotoSize `json:"cover,omitempty"`
	StartTimestamp int          `json:"start_timestamp,omitempty"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type PaidMedia struct {
	Type     string       `json:"type"`
	Width    int          `json:"width,omitempty"`
	Height   int          `json:"height,omitempty"`
	Duration int          `json:"duration,omitempty"`
	Photo    []*PhotoSize `json:"photo"`
	Video    *Video       `json:"video"`
}

type PaidMediaInfo struct {
	StarCount int          `json:"star_count"`
	PaidMedia []*PaidMedia `json:"paid_media"`
}

type Story struct {
	Chat *Chat `json:"chat"`
	ID   int   `json:"id"`
}

type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id,omitempty"`
	SmallFileUniqueID string `json:"small_file_unique_id,omitempty"`
	BigFileID         string `json:"big_file_id,omitempty"`
	BigFileUniqueID   string `json:"big_file_unique_id,omitempty"`
}

type StickerSet struct {
	Name        string     `json:"name,omitempty"`
	Title       string     `json:"title,omitempty"`
	StickerType string     `json:"sticker_type,omitempty"`
	Stickers    []Sticker  `json:"stickers,omitempty"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
}

type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    string        `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`
	FileSize         int           `json:"file_size,omitempty"`
}

type InputMediaPhoto struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             string           `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
}

type InputMediaVideo struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Thumbnail             interface{}      `json:"thumbnail,omitempty"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             string           `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
	Width                 int              `json:"width,omitempty"`
	Height                int              `json:"height,omitempty"`
	Duration              int              `json:"duration,omitempty"`
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
}

type InputMediaAnimation struct {
	Type                  string           `json:"type"`
	Media                 string           `json:"media"`
	Thumbnail             interface{}      `json:"thumbnail,omitempty"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             string           `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
	Width                 int              `json:"width,omitempty"`
	Height                int              `json:"height,omitempty"`
	Duration              int              `json:"duration,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
}

type InputMediaAudio struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumbnail       interface{}      `json:"thumbnail,omitempty"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       string           `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Duration        int              `json:"duration,omitempty"`
	Performer       string           `json:"performer,omitempty"`
	Title           string           `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type                        string           `json:"type"`
	Media                       string           `json:"media"`
	Thumbnail                   interface{}      `json:"thumbnail,omitempty"`
	Caption                     string           `json:"caption,omitempty"`
	ParseMode                   string           `json:"parse_mode,omitempty"`
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
}

type InputMedia struct {
	InputMediaAnimation *InputMediaAnimation
	InputMediaDocument  *InputMediaDocument
	InputMediaAudio     *InputMediaAudio
	InputMediaPhoto     *InputMediaPhoto
	InputMediaVideo     *InputMediaVideo
}

type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
}

type InputPaidMediaVideo struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// thumbpath         string
	Thumbnail         string `json:"thumbnail,omitempty"` // InputFile or String
	Width             int    `json:"width,omitempty"`
	Height            int    `json:"height,omitempty"`
	Duration          int    `json:"duration,omitempty"`
	SupportsStreaming bool   `json:"supports_streaming,omitempty"`
}

type InputFile struct {
	File string
}

type InputPaidMedia struct {
	InputPaidMediaPhoto *InputPaidMediaPhoto
	InputPaidMediaVideo *InputPaidMediaVideo
}

type Gift struct {
	ID               string  `json:"id"`
	Sticker          Sticker `json:"sticker"`
	StarCount        int     `json:"star_count"`
	UpgradeStarCount *int    `json:"upgrade_star_count,omitempty"`
	TotalCount       *int    `json:"total_count,omitempty"`
	RemainingCount   *int    `json:"remaining_count,omitempty"`
}

type Gifts struct {
	G []*Gift `json:"gifts"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type CallbackQuery struct {
	ID              string                    `json:"id"`
	From            *User                     `json:"from"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageID string                    `json:"inline_message_id,omitempty"`
	ChatInstance    string                    `json:"chat_instance"`
	Data            string                    `json:"data,omitempty"`
	GameShortName   string                    `json:"game_short_name,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string      `json:"text"`                      // Required
	WebApp         *WebAppInfo `json:"web_app,omitempty"`         // Optional
	StartParameter string      `json:"start_parameter,omitempty"` // Optional
}

type InputMessageContent struct {
	*InputTextMessageContent     `json:",inline,omitempty"`
	*InputLocationMessageContent `json:",inline,omitempty"`
	*InputVenueMessageContent    `json:",inline,omitempty"`
	*InputContactMessageContent  `json:",inline,omitempty"`
	*InputInvoiceMessageContent  `json:",inline,omitempty"`
}

type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                `json:"url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	PhotoURL              string                `json:"photo_url"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	PhotoWidth            int                   `json:"photo_width,omitempty"`
	PhotoHeight           int                   `json:"photo_height,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultGif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	GifURL                string                `json:"gif_url"`
	GifWidth              int                   `json:"gif_width,omitempty"`
	GifHeight             int                   `json:"gif_height,omitempty"`
	GifDuration           int                   `json:"gif_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	Mpeg4URL              string                `json:"mpeg4_url"`
	Mpeg4Width            int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height           int                   `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         int                   `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	VideoURL              string                `json:"video_url"`
	MimeType              string                `json:"mime_type"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	VideoWidth            int                   `json:"video_width,omitempty"`
	VideoHeight           int                   `json:"video_height,omitempty"`
	VideoDuration         int                   `json:"video_duration,omitempty"`
	Description           string                `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                   `json:"live_period,omitempty"`
	Heading              int                   `json:"heading,omitempty"`
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL         string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int                   `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	GooglePlaceID       string                `json:"google_place_id,omitempty"`
	GooglePlaceType     string                `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	Vcard               string                `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	PhotoFileID           string                `json:"photo_file_id"`
	Title                 string                `json:"title,omitempty"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	GifFileID             string                `json:"gif_file_id"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	Mpeg4FileID           string                `json:"mpeg4_file_id"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	VideoFileID           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InputMessageContent

type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          string              `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int    `json:"expiration_date"`
}

type BusinessConnection struct {
	ID         string `json:"id"`
	User       *User  `json:"user"`
	UserChatID int    `json:"user_chat_id"`
	Date       int    `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 *Chat  `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName string                          `json:"time_zone_name"`
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours,omitempty"`
}

type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type BotCommandScope struct {
	Type   string  `json:"type"`
	ChatID *string `json:"chat_id,omitempty"`
	UserID *int    `json:"user_id,omitempty"`
}

type BotName struct {
	Name string `json:"name"`
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}
