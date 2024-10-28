package types

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
	MessageOriginUser       *MessageOriginUser
	MessageOriginHiddenUser *MessageOriginHiddenUser
	MessageOriginChat       *MessageOriginChat
	MessageOriginChannel    *MessageOriginChannel
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

type Message struct {
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
	ReplyToMessage                *Message                       `json:"reply_to_message"`
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
