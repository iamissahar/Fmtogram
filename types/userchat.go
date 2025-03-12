package types

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
