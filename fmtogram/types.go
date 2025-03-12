package fmtogram

import "github.com/iamissahar/Fmtogram/types"

type Message struct {
	MessageID                     int                                  `json:"message_id"`
	MessageThreadID               int                                  `json:"message_thread_id"`
	From                          *types.User                          `json:"from"`
	SenderChat                    *types.Chat                          `json:"sender_chat"`
	SenderBoostCount              int                                  `json:"sender_boost_count"`
	SenderBusinessBot             *types.User                          `json:"sender_business_bot"`
	DateUnix                      int                                  `json:"date"`
	BusinessConnectionID          string                               `json:"business_connection_id"`
	Ch                            *types.Chat                          `json:"chat"`
	ForwardOrigin                 *types.MessageOrigin                 `json:"forward_origin"`
	IsTopicMessage                bool                                 `json:"is_topic_message"`
	IsAutomaticForward            bool                                 `json:"is_automatic_forward"`
	ReplyToMessage                *Message                             `json:"reply_to_message"`
	ExternalReply                 *types.ExternalReplyInfo             `json:"external_reply"`
	Quote                         *types.TextQuote                     `json:"quote"`
	ReplyToStory                  *types.Story                         `json:"reply_to_story"`
	ViaBot                        *types.User                          `json:"via_bot"`
	EditDate                      int                                  `json:"edit_date"`
	HasProtectedContent           bool                                 `json:"has_protected_content"`
	IsFrommOffline                bool                                 `json:"is_from_offline"`
	MediaGroupID                  string                               `json:"media_group_id"`
	AuthorSignature               string                               `json:"author_signature"`
	Txt                           string                               `json:"text"`
	Entities                      []*types.MessageEntity               `json:"entities"`
	LinkPreviewOptions            *types.LinkPreviewOptions            `json:"link_preview_options"`
	EffectID                      string                               `json:"effect_id"`
	Animation                     *types.Animation                     `json:"animation"`
	Ad                            *types.Audio                         `json:"audio"`
	Doc                           *types.Document                      `json:"document"`
	PaidM                         *types.PaidMediaInfo                 `json:"paid_media"`
	Ph                            []*types.PhotoSize                   `json:"photo"`
	Stckr                         *types.Sticker                       `json:"sticker"`
	Stry                          *types.Story                         `json:"story"`
	Vd                            *types.Video                         `json:"video"`
	Vdn                           *types.VideoNote                     `json:"video_note"`
	Vc                            *types.Voice                         `json:"voice"`
	Caption                       string                               `json:"caption"`
	CaptionEntities               []*types.MessageEntity               `json:"caption_entities"`
	ShowCaptionAboveMedia         bool                                 `json:"show_caption_above_media"`
	HasMediaSpoiler               bool                                 `json:"has_media_spoiler"`
	Contact                       *types.Contact                       `json:"contact"`
	Dice                          *types.Dice                          `json:"dice"`
	Game                          *types.Game                          `json:"game"`
	Poll                          *types.Poll                          `json:"poll"`
	Venue                         *types.Venue                         `json:"venue"`
	Location                      *types.Location                      `json:"location"`
	NewChatMembers                []*types.User                        `json:"new_chat_members"`
	LeftChatMember                *types.User                          `json:"left_chat_member"`
	NewChatTitle                  string                               `json:"new_chat_title"`
	NewChatPhoto                  []*types.PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                                 `json:"delete_chat_photo"`
	GroupChatCreated              bool                                 `json:"group_chat_created"`
	SuperGroupChatCreated         bool                                 `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                                 `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *types.MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               int                                  `json:"migrate_from_chat_id"`
	PinnedMessage                 *types.MaybeInaccessibleMessage      `json:"pinned_message"`
	Invoice                       *types.Invoice                       `json:"invoice"`
	SuccessfulPayment             *types.SuccessfulPayment             `json:"successful_payment"`
	RefundedPayment               *types.RefundedPayment               `json:"refunded_payment"`
	UsersShared                   *types.UsersShared                   `json:"users_shared"`
	ChatShared                    *types.ChatShared                    `json:"chat_shared"`
	ConnectedWebsite              string                               `json:"connected_website"`
	WriteAccessAllowed            *types.WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  *types.PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       *types.ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	BoostAdded                    *types.ChatBoostAdded                `json:"boost_added"`
	ChatBackgroundSet             *types.ChatBackground                `json:"chat_background_set"`
	ForumTopicCreated             *types.ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              *types.ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              *types.ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            *types.ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       *types.GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     *types.GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	GiveawayCreated               *types.GiveawayCreated               `json:"giveaway_created"`
	Giveaway                      *types.Giveaway                      `json:"giveaway"`
	GiveawayWinners               *types.GiveawayWinners               `json:"giveaway_winners"`
	GiveawayCompleted             *types.GiveawayCompleted             `json:"giveaway_completed"`
	VideoChatScheduled            *types.VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              *types.VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                *types.VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  *types.VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    *types.WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   *types.InlineKeyboardMarkup          `json:"reply_markup"`
}

type BusinessConnection struct {
	BID        string      `json:"id"`
	User       *types.User `json:"user"`
	UserChatID int64       `json:"user_chat_id"`
	BDate      int         `json:"date"`
	BCanReply  bool        `json:"can_reply"`
	BIsEnabled bool        `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionID string      `json:"business_connection_id"`
	Ch                   *types.Chat `json:"chat"`
	MsgIDs               []int       `json:"message_ids"`
}

type MessageReactionUpdated struct {
	Ch       *types.Chat           `json:"chat"`
	MsgID    int                   `json:"message_id"`
	User     *types.User           `json:"user,omitempty"`
	ActorCh  *types.Chat           `json:"actor_chat,omitempty"`
	DateUnix int                   `json:"date"`
	OldReact []*types.ReactionType `json:"old_reaction"`
	NewReact []*types.ReactionType `json:"new_reaction"`
}

type MessageReactionCountUpdated struct {
	Ch       *types.Chat            `json:"chat"`
	MsgID    int                    `json:"message_id"`
	DateUnix int                    `json:"date"`
	React    []*types.ReactionCount `json:"reactions"`
}

type InlineQuery struct {
	IqID       string          `json:"id"`
	From       *types.User     `json:"from"`
	Q          string          `json:"query"`
	IqOffset   string          `json:"offset"`
	IqChatType string          `json:"chat_type"`
	IqLocation *types.Location `json:"location"`
}

type ChosenInlineResult struct {
	ResultID string          `json:"result_id"`
	From     *types.User     `json:"from"`
	Loc      *types.Location `json:"location,omitempty"`
	InMsgID  string          `json:"inline_message_id,omitempty"`
	Q        string          `json:"query"`
}

type CallbackQuery struct {
	ID              string                          `json:"id"`
	From            *types.User                     `json:"from"`
	Message         *types.MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageID string                          `json:"inline_message_id,omitempty"`
	ChatInstance    string                          `json:"chat_instance"`
	Data            string                          `json:"data,omitempty"`
	GameShortName   string                          `json:"game_short_name,omitempty"`
}

type ShippingQuery struct {
	ID              string                 `json:"id"`
	From            *types.User            `json:"from"`
	InvoicePayload  string                 `json:"invoice_payload"`
	ShippingAddress *types.ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string           `json:"id"`
	From             *types.User      `json:"from"`
	Currency         string           `json:"currency"`
	TotalAmount      int              `json:"total_amount"`
	InvoicePayload   string           `json:"invoice_payload"`
	ShippingOptionID string           `json:"shipping_option_id"`
	OrderInfo        *types.OrderInfo `json:"order_info"`
}

type Poll struct {
	ID                    string                 `json:"id"`
	Question              string                 `json:"question"`
	QuestionEnities       []*types.MessageEntity `json:"question_entities,omitempty"`
	Options               []*types.PollOption    `json:"options"`
	TotalVoterCount       int                    `json:"total_voter_count"`
	IsClosed              bool                   `json:"is_closed"`
	IsAnonymous           bool                   `json:"is_anonymous"`
	Type                  string                 `json:"type"`
	AllowsMultipleAnswers bool                   `json:"allows_multiple_answers"`
	CorrectOptionID       int                    `json:"correct_option_id,omitempty"`
	Explanation           string                 `json:"explanation,omitempty"`
	ExplanationEntities   []*types.MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int                    `json:"open_period,omitempty"`
	CloseDate             int                    `json:"close_date,omitempty"`
}

type PollAnswer struct {
	PollID    string      `json:"poll_id"`
	VoterChat *types.Chat `json:"voter_chat,omitempty"`
	User      *types.User `json:"user,omitempty"`
	OptionIDs []int       `json:"option_ids"`
}

type ChatMemberUpdated struct {
	Chat                    *types.Chat           `json:"chat"`
	From                    *types.User           `json:"from"`
	Date                    int                   `json:"date"`
	OldChatMember           *types.ChatMember     `json:"old_chat_member"`
	NewChatMember           *types.ChatMember     `json:"new_chat_member"`
	InviteLink              *types.ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool                  `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool                  `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatJoinRequest struct {
	Chat       *types.Chat           `json:"chat"`
	From       *types.User           `json:"from"`
	UserChatID int64                 `json:"user_chat_id"`
	Date       int                   `json:"date"`
	Bio        string                `json:"bio,omitempty"`
	InviteLink *types.ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatBoostUpdated struct {
	Chat  *types.Chat      `json:"chat"`
	Boost *types.ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       *types.Chat            `json:"chat"`
	BoostID    string                 `json:"boost_id"`
	RemoveDate int                    `json:"remove_date"`
	Source     *types.ChatBoostSource `json:"source"`
}

type Update struct {
	UpdateID               int                          `json:"update_id"`
	Msg                    *Message                     `json:"message"`
	EditedMsg              *Message                     `json:"edited_message"`
	ChanPost               *Message                     `json:"channel_post"`
	EditedChanPost         *Message                     `json:"edited_channel_post"`
	BusinessConn           *BusinessConnection          `json:"business_connection"`
	BusinessMsg            *Message                     `json:"business_message"`
	EditedBusinessMsg      *Message                     `json:"edited_business_message"`
	DeletedBusinessMessage *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	MessageR               *MessageReactionUpdated      `json:"message_reaction"`
	MessageRcount          *MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQ                *InlineQuery                 `json:"inline_query"`
	ChosenInlineR          *ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQ              *CallbackQuery               `json:"callback_query"`
	ShippingQ              *ShippingQuery               `json:"shipping_query"`
	PreCheckoutQ           *PreCheckoutQuery            `json:"pre_checkout_query"`
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
