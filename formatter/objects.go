package formatter

import "github.com/iamissahar/Fmtogram/types"

type photo struct {
	Type                  string                 `json:"type,omitempty" field:"type" ifempty:"false"`
	Media                 string                 `json:"media,omitempty" field:"media" ifempty:"false"`
	Photo                 string                 `json:"photo,omitempty" field:"photo" ifempty:"false"`
	Caption               string                 `json:"caption,omitempty" field:"caption" ifempty:"false"`
	ParseMode             string                 `json:"parse_mode,omitempty" field:"parse_mode" ifempty:"false"`
	CaptionEntities       []*types.MessageEntity `json:"caption_entities,omitempty" field:"caption_entities" ifempty:"false"`
	ShowCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty" field:"show_caption_above_media" ifempty:"false"`
	HasSpoiler            bool                   `json:"has_spoiler,omitempty" field:"has_spoiler" ifempty:"false"`
	gottenFrom            int
}

type video struct {
	Type                  string                 `json:"type,omitempty" field:"type" ifempty:"false"`
	Media                 string                 `json:"media,omitempty" field:"media" ifempty:"false"`
	Video                 string                 `json:"video,omitempty" field:"video" ifempty:"false"`
	Thumbnail             string                 `json:"thumbnail,omitempty" field:"thumbnail" ifempty:"false"`
	Caption               string                 `json:"caption,omitempty" field:"caption" ifempty:"false"`
	ParseMode             string                 `json:"parse_mode,omitempty" field:"parse_mode" ifempty:"false"`
	CaptionEntities       []*types.MessageEntity `json:"caption_entities,omitempty" field:"caption_entities" ifempty:"false"`
	ShowCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty" field:"show_caption_above_media" ifempty:"false"`
	Width                 int                    `json:"width,omitempty" field:"width" ifempty:"false"`
	Height                int                    `json:"height,omitempty" field:"height" ifempty:"false"`
	Duration              int                    `json:"duration,omitempty" field:"duration" ifempty:"false"`
	SupportsStreaming     bool                   `json:"supports_streaming,omitempty" field:"supports_streaming" ifempty:"false"`
	HasSpoiler            bool                   `json:"has_spoiler,omitempty" field:"has_spoiler" ifempty:"false"`
	Cover                 string                 `json:"cover,omitempty" field:"cover" ifempty:"false"`
	StartTimestamp        *int                   `json:"start_timestamp,omitempty" field:"start_timestamp" ifempty:"false"`
	videoGottenFrom       int
	thumbnailGottenFrom   int
	coverGottenFrom       int
}

type audio struct {
	Type                string                 `json:"type,omitempty" field:"type" ifempty:"false"`
	Media               string                 `json:"media,omitempty" field:"media" ifempty:"false"`
	Audio               string                 `json:"audio,omitempty" field:"audio" ifempty:"false"`
	Thumbnail           string                 `json:"thumbnail,omitempty" field:"thumbnail" ifempty:"false"`
	Caption             string                 `json:"caption,omitempty" field:"caption" ifempty:"false"`
	ParseMode           string                 `json:"parse_mode,omitempty" field:"parse_mode" ifempty:"false"`
	CaptionEntities     []*types.MessageEntity `json:"caption_entities,omitempty" field:"caption_entities" ifempty:"false"`
	Duration            int                    `json:"duration,omitempty" field:"duration" ifempty:"false"`
	Performer           string                 `json:"performer,omitempty" field:"performer" ifempty:"false"`
	Title               string                 `json:"title,omitempty" field:"title" ifempty:"false"`
	audioGottenFrom     int
	thumbnailGottenFrom int
}

type document struct {
	Type                        string                 `json:"type,omitempty" field:"type" ifempty:"false"`
	Media                       string                 `json:"media,omitempty" field:"media" ifempty:"false"`
	Document                    string                 `json:"document,omitempty" field:"document" ifempty:"false"`
	Thumbnail                   string                 `json:"thumbnail,omitempty" field:"thumbnail" ifempty:"false"`
	Caption                     string                 `json:"caption,omitempty" field:"caption" ifempty:"false"`
	ParseMode                   string                 `json:"parse_mode,omitempty" field:"parse_mode" ifempty:"false"`
	CaptionEntities             []*types.MessageEntity `json:"caption_entities,omitempty" field:"caption_entities" ifempty:"false"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty" field:"disable_content_type_detection" ifempty:"false"`
	documentGottenFrom          int
	thumbnailGottenFrom         int
}

type animation struct {
	Type                string `json:"type,omitempty" field:"type" ifempty:"false"`
	Media               string `json:"media,omitempty" field:"media" ifempty:"false"`
	Animation           string `json:"animation,omitempty" field:"animation" ifempty:"false"`
	Thumbnail           string `json:"thumbnail,omitempty" field:"thumbnail" ifempty:"false"`
	Width               int    `json:"width,omitempty" field:"width" ifempty:"false"`
	Height              int    `json:"height,omitempty" field:"height" ifempty:"false"`
	Duration            int    `json:"duration,omitempty" field:"duration" ifempty:"false"`
	HasSpoiler          bool   `json:"has_spoiler,omitempty" field:"has_spoiler" ifempty:"false"`
	animationGottenFrom int
	thumbnailGottenFrom int
}

type voice struct {
	Type       string `json:"type,omitempty" field:"type" ifempty:"false"`
	Media      string `json:"media,omitempty" field:"media" ifempty:"false"`
	Voice      string `json:"voice,omitempty" field:"voice" ifempty:"false"`
	Duration   int    `json:"duration,omitempty" field:"duration" ifempty:"false"`
	gottenFrom int
}

type videonote struct {
	VideoNote           string `json:"video_note,omitempty" field:"vidoe_note" ifempty:"false"`
	Duration            int    `json:"duration,omitempty" field:"duration" ifempty:"false"`
	Length              int    `json:"length,omitempty" field:"length" ifempty:"false"`
	Thumbnail           string `json:"thumbnail,omitempty" field:"thumbnail" ifempty:"false"`
	videoGottenFrom     int
	thumbnailGottenFrom int
}

type location struct {
	Latitude             float64 `json:"latitude,omitempty" field:"latitude" ifempty:"false"`
	Longitude            float64 `json:"longitude,omitempty" field:"longitude" ifempty:"false"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty" field:"horizontal_accuracy" ifempty:"false"`
	LivePeriod           int     `json:"live_period,omitempty" field:"live_period" ifempty:"false"`
	Heading              int     `json:"heading,omitempty" field:"heading" ifempty:"false"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty" field:"proximity_alert_radius" ifempty:"false"`
	Title                string  `json:"title,omitempty" field:"title" ifempty:"false"`
	Address              string  `json:"address,omitempty" field:"address" ifempty:"false"`
	FoursquareID         string  `json:"foursquare_id,omitempty" field:"foursquare_id" ifempty:"false"`
	FoursquareType       string  `json:"foursquare_type,omitempty" field:"foursquare_type" ifempty:"false"`
	GooglePlaceID        string  `json:"google_place_id,omitempty" field:"google_place_id" ifempty:"false"`
	GooglePlaceType      string  `json:"google_place_type,omitempty" field:"google_place_type" ifempty:"false"`
}

type contact struct {
	PhoneNumber string `json:"phone_number,omitempty" field:"phone_number" ifempty:"false"`
	FirstName   string `json:"first_name,omitempty" field:"first_name" ifempty:"false"`
	LastName    string `json:"last_name,omitempty" field:"last_name" ifempty:"false"`
	Vcard       string `json:"vcard,omitempty" field:"vcard" ifempty:"false"`
}

type poll struct {
	Question             string                 `json:"question,omitempty" field:"question" ifempty:"false"`
	QuestionParsemode    string                 `json:"question_parse_mode,omitempty" field:"question_parse_mode" ifempty:"false"`
	QuestionEntities     []*types.MessageEntity `json:"question_entities,omitempty" field:"question_entities" ifempty:"false"`
	Options              []*types.PollOption    `json:"options,omitempty" field:"options" ifempty:"false"`
	IsAnonymous          *bool                  `json:"is_anonymous,omitempty" field:"is_anonymous" ifempty:"true"`
	Type                 string                 `json:"type,omitempty" field:"type" ifempty:"false"`
	AllowMultipleAnswers bool                   `json:"allows_multiple_answers,omitempty" field:"allows_multiple_answers" ifempty:"false"`
	CorrectOptionID      string                 `json:"correct_option_id,omitempty" field:"correct_option_id" ifempty:"false"`
	Explanation          string                 `json:"explanation,omitempty" field:"explanation" ifempty:"false"`
	ExplanationParsemode string                 `json:"explanation_parse_mode,omitempty" field:"explanation_parse_mode" ifempty:"false"`
	ExplanationEntities  []*types.MessageEntity `json:"explanation_entities,omitempty" field:"explanation_entities" ifempty:"false"`
	OpenPeriod           int                    `json:"open_period,omitempty" field:"open_period" ifempty:"false"`
	CloseDate            int                    `json:"close_date,omitempty" field:"close_date" ifempty:"false"`
	IsClosed             bool                   `json:"is_closed,omitempty" field:"is_closed" ifempty:"false"`
}

type sticker struct {
	SetName             string              `json:"sticker_set_name,omitempty" field:"sticker_set_name" ifempty:"false"`
	Sticker             string              `json:"sticker,omitempty" field:"sticker" ifempty:"false"`
	Emoji               string              `json:"emoji,omitempty" field:"emoji" ifempty:"false"`
	Emojies             []string            `json:"emoji_list,omitempty" field:"emoji_list" ifempty:"false"`
	EmojiID             string              `json:"custom_emoji_id,omitempty" field:"custom_emoji_id" ifempty:"false"`
	EmojiIDs            []string            `json:"custom_emoji_ids,omitempty" field:"custom_emoji_ids" ifempty:"false"`
	StickerFormat       string              `json:"sticker_format,omitempty" field:"sticker_format" ifempty:"false"`
	Position            string              `json:"position,omitempty" field:"position" ifempty:"false"`
	OldSticker          string              `json:"old_sticker,omitempty" field:"old_sticker" ifempty:"false"`
	Keywords            []string            `json:"keywords,omitempty" field:"keywords" ifempty:"false"`
	MaskPosition        *types.MaskPosition `json:"mask_position,omitempty" field:"mask_position" ifempty:"false"`
	Thumbnail           string              `json:"thumbnail,omitempty" field:"thumbnail" ifempty:"false"`
	Format              string              `json:"format,omitempty" field:"format" ifempty:"false"`
	GiftID              string              `json:"gift_id,omitempty" field:"gift_id" ifempty:"false"`
	PayForUpgrade       bool                `json:"pay_for_upgrade,omitempty" field:"pay_for_upgrade" ifempty:"false"`
	stickerGottenFrom   int
	thumbnailGottenFrom int
	thumbnailType       string
}

type forum struct {
	Name        string  `json:"name,omitempty" field:"name" ifempty:"false"`
	IconColor   int     `json:"icon_color,omitempty" field:"icon_color" ifempty:"false"`
	IconEmojiID *string `json:"icon_custom_emoji_id,omitempty" field:"icon_custom_emoji_id" ifempty:"true"`
}

type information struct {
	MessageID                  int                            `json:"message_id,omitempty" field:"message_id" ifempty:"false"`
	MessageIDs                 []int                          `json:"message_ids,omitempty" field:"message_ids" ifempty:"false"`
	Text                       string                         `json:"text,omitempty" field:"text" ifempty:"false"`
	Caption                    string                         `json:"caption,omitempty" field:"caption" ifempty:"false"`
	ParseMode                  string                         `json:"parse_mode,omitempty" field:"parse_mode" ifempty:"false"`
	MessageThreadID            int                            `json:"message_thread_id,omitempty" field:"message_thread_id" ifempty:"false"`
	Entities                   []*types.MessageEntity         `json:"entities,omitempty" field:"entities" ifempty:"false"`
	CaptionEntities            []*types.MessageEntity         `json:"caption_entities,omitempty" field:"caption_entities" ifempty:"false"`
	LinkPreviewOptions         *types.LinkPreviewOptions      `json:"link_preview_options,omitempty" field:"link_preview_options" ifempty:"false"`
	DisableNotification        bool                           `json:"disable_notification,omitempty" field:"disable_notification" ifempty:"false"`
	ProtectContent             bool                           `json:"protect_content,omitempty" field:"protect_content" ifempty:"false"`
	MessageEffectID            string                         `json:"message_effect_id,omitempty" field:"message_effect_id" ifempty:"false"`
	ShowCaptionAboveMedia      bool                           `json:"show_caption_above_media,omitempty" field:"show_caption_above_media" ifempty:"false"`
	ReplyParameters            *types.ReplyParameters         `json:"reply_parameters,omitempty" field:"reply_parameters" ifempty:"false"`
	AllowPaidBroadcast         bool                           `json:"allow_paid_broadcast,omitempty" field:"allow_paid_broadcast" ifempty:"false"`
	StarCount                  int                            `json:"star_count,omitempty" field:"star_count" ifempty:"false"`
	Payload                    string                         `json:"payload,omitempty" field:"payload" ifempty:"false"`
	Emoji                      string                         `json:"emoji,omitempty" field:"emoji" ifempty:"false"`
	Action                     string                         `json:"action,omitempty" field:"action" ifempty:"false"`
	Reaction                   []*types.ReactionType          `json:"reaction,omitempty" field:"reaction" ifempty:"false"`
	ReactionIsBig              bool                           `json:"is_big,omitempty" field:"is_big" ifempty:"false"`
	Offset                     int                            `json:"offset,omitempty" field:"offset" ifempty:"false"`
	Limit                      int                            `json:"limit,omitempty" field:"limit" ifempty:"false"`
	EmojiStatusCustomEmojiID   *string                        `json:"emoji_status_custom_emoji_id,omitempty" field:"emoji_status_custom_emoji_id" ifempty:"false"`
	EmojiStatusExpirationDate  int                            `json:"emoji_status_expiration_date,omitempty" field:"emoji_status_expiration_date" ifempty:"false"`
	FileID                     string                         `json:"file_id,omitempty" field:"file_id" ifempty:"false"`
	UntilDate                  int64                          `json:"until_date,omitempty" field:"until_date" ifempty:"false"`
	RevokeMessages             bool                           `json:"revoke_messages,omitempty" field:"revoke_messages" ifempty:"false"`
	OnlyIfBanned               bool                           `json:"only_if_banned,omitempty" field:"only_if_banned" ifempty:"false"`
	Permissions                *types.ChatPermissions         `json:"permissions,omitempty" field:"permissions" ifempty:"false"`
	IndependentChatPermissions bool                           `json:"use_independent_chat_permissions,omitempty" field:"use_independent_chat_permissions" ifempty:"false"`
	AdminRights                *types.ChatAdministratorRights `json:"rights,omitempty" field:"rights" ifempty:"false"`
	CustomTitle                string                         `json:"custom_title,omitempty" field:"custom_title" ifempty:"false"`
	UserID                     int                            `json:"user_id,omitempty" field:"user_id" ifempty:"false"`
	CallBackQueryID            string                         `json:"callback_query_id,omitempty" field:"callback_query_id" ifempty:"false"`
	ShowAlert                  bool                           `json:"show_alert,omitempty" field:"show_alert" ifempty:"false"`
	Url                        string                         `json:"url,omitempty" field:"url" ifempty:"false"`
	CacheTime                  int                            `json:"cache_time,omitempty" field:"cache_time" ifempty:"false"`
	InlineMessageID            string                         `json:"inline_message_id,omitempty" field:"inline_message_id" ifempty:"false"`
	Errors                     []*types.PassportElementError  `json:"errors,omitempty" field:"errors" ifempty:"false"`
	GiftParseMode              string                         `json:"text_parse_mode,omitempty" field:"text_parse_mode" ifempty:"false"`
	GiftEntities               []*types.MessageEntity         `json:"text_entities,omitempty" field:"text_entities" ifempty:"false"`
	Description                string                         `json:"custom_description,omitempty" field:"custom_description" ifempty:"false"`
	RemoveCaption              bool                           `json:"remove_caption,omitempty" field:"remove_caption" ifempty:"false"`
	IsAnonymous                *bool                          `json:"is_anonymous,omitempty" field:"is_anonymous" ifempty:"true"`
	CanManageChat              *bool                          `json:"can_manage_chat,omitempty" field:"can_manage_chat" ifempty:"true"`
	CanDeleteMessages          *bool                          `json:"can_delete_messages,omitempty" field:"can_delete_messages" ifempty:"true"`
	CanManageVideoChats        *bool                          `json:"can_manage_video_chats,omitempty" field:"can_manage_video_chats" ifempty:"true"`
	CanRestrictMembers         *bool                          `json:"can_restrict_members,omitempty" field:"can_restrict_members" ifempty:"true"`
	CanPromoteMembers          *bool                          `json:"can_promote_members,omitempty" field:"can_promote_members" ifempty:"true"`
	CanChangeInfo              *bool                          `json:"can_change_info,omitempty" field:"can_change_info" ifempty:"true"`
	CanInviteUsers             *bool                          `json:"can_invite_users,omitempty" field:"can_invite_users" ifempty:"true"`
	CanPostStories             *bool                          `json:"can_post_stories,omitempty" field:"can_post_stories" ifempty:"true"`
	CanEditStories             *bool                          `json:"can_edit_stories,omitempty" field:"can_edit_stories" ifempty:"true"`
	CanDeleteStories           *bool                          `json:"can_delete_stories,omitempty" field:"can_delete_stories" ifempty:"true"`
	CanPostMessages            *bool                          `json:"can_post_messages,omitempty" field:"can_post_messages" ifempty:"true"`
	CanEditMessages            *bool                          `json:"can_edit_messages,omitempty" field:"can_edit_messages" ifempty:"true"`
	CanPinMessages             *bool                          `json:"can_pin_messages,omitempty" field:"can_pin_messages" ifempty:"true"`
	CanManageTopics            *bool                          `json:"can_manage_topics,omitempty" field:"can_manage_topics" ifempty:"true"`
	VideoStartTimestamp        int                            `json:"video_start_timestamp,omitempty" field:"video_start_timestamp" ifempty:"false"`
	SetName                    string                         `json:"name,omitempty" field:"name" ifempty:"false"`
	SetTitle                   string                         `json:"title,omitempty" field:"title" ifempty:"false"`
	StickerType                string                         `json:"sticker_type,omitempty" field:"sticker_type" ifempty:"false"`
	NeedsRepainting            bool                           `json:"needs_repainting,omitempty" field:"needs_repainting" ifempty:"false"`
}

type chat struct {
	FromChatID           interface{} `json:"from_chat_id,omitempty" field:"from_chat_id" ifempty:"false"`
	ID                   interface{} `json:"chat_id,omitempty" field:"chat_id" ifempty:"false"`
	BusinessConnectionID string      `json:"business_connection_id,omitempty" field:"business_connection_id" ifempty:"false"`
	SenderChatID         int         `json:"sender_chat_id,omitempty" field:"sender_chat_id" ifempty:"false"`
	Title                string      `json:"title,omitempty" field:"title" ifempty:"false"`
	Description          string      `json:"description,omitempty" field:"description" ifempty:"false"`
}

type link struct {
	Name               string `json:"name,omitempty" field:"name" ifempty:"false"`
	ExpireDate         int64  `json:"expire_date,omitempty" field:"expire_date" ifempty:"false"`
	MemberLimit        int    `json:"member_limit,omitempty" field:"member_limit" ifempty:"false"`
	JoinRequest        bool   `json:"creates_join_request,omitempty" field:"creates_join_request" ifempty:"false"`
	InviteLink         string `json:"invite_link,omitempty" field:"invite_link" ifempty:"false"`
	SubscriptionPeriod int    `json:"subscription_period,omitempty" field:"subscription_period" ifempty:"false"`
	SubscriptionPrice  int    `json:"subscription_price,omitempty" field:"subscription_price" ifempty:"false"`
}

type bot struct {
	Commands         []*types.BotCommand            `json:"commands,omitempty" field:"commands" ifempty:"false"`
	Scope            *types.BotCommandScope         `json:"scope,omitempty" field:"scope" ifempty:"false"`
	Language         *string                        `json:"language_code,omitempty" field:"language_code" ifempty:"true"`
	Name             *string                        `json:"name,omitempty" field:"name" ifempty:"true"`
	Description      *string                        `json:"description,omitempty" field:"description" ifempty:"true"`
	ShortDescription *string                        `json:"short_description,omitempty" field:"short_description" ifempty:"true"`
	MenuButtom       *types.MenuButton              `json:"menu_button,omitempty" field:"menu_button" ifempty:"false"`
	Rights           *types.ChatAdministratorRights `json:"rights,omitempty" field:"rights" ifempty:"false"`
	ForChannels      bool                           `json:"for_channels,omitempty" field:"for_channels" ifempty:"false"`
}

type inlineKeyboardButton struct {
	Text                         string                             `json:"text" field:"text" ifempty:"false"`
	Url                          string                             `json:"url,omitempty" field:"url" ifempty:"false"`
	CallbackData                 string                             `json:"callback_data,omitempty" field:"callback_data" ifempty:"false"`
	WebApp                       *types.WebAppInfo                  `json:"web_app,omitempty" field:"web_app" ifempty:"false"`
	LoginUrl                     *types.LoginUrl                    `json:"login_url,omitempty" field:"login_url" ifempty:"false"`
	SwitchInlineQuery            *string                            `json:"switch_inline_query,omitempty" field:"switch_inline_query" ifempty:"true"`
	SwitchInlineQueryCurrentChat *string                            `json:"switch_inline_query_current_chat,omitempty" field:"switch_inline_query_current_chat" ifempty:"true"`
	SwitchInlineQueryChosenChat  *types.SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty" field:"switch_inline_query_chosen_chat" ifempty:"false"`
	CallbackGame                 *types.CallbackGame                `json:"callback_game,omitempty" field:"callback_game" ifempty:"false"`
	Pay                          bool                               `json:"pay,omitempty" field:"pay" ifempty:"false"`
	storage                      [9]int
}

type keyboard struct {
	Keyboard kb `json:"reply_markup,omitempty" field:"reply_markup" ifempty:"false"`
}

type forcereply struct {
	Forcereply            bool   `json:"force_reply,omitempty" field:"force_reply" ifempty:"false"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty" field:"input_field_placeholder" ifempty:"false"`
	Selective             bool   `json:"selective,omitempty" field:"selective" ifempty:"false"`
}

type inline struct {
	Keyboard [][]*inlineKeyboardButton `json:"inline_keyboard" field:"inline_keyboard" ifempty:"false"`
}

type replyKeyboardButton struct {
	Text            string                            `json:"text,omitempty" field:"text" ifempty:"false"`
	RequestUsers    *types.KeyboardButtonRequestUsers `json:"request_users,omitempty" field:"request_users" ifempty:"false"`
	RequestChat     *types.KeyboardButtonRequestChat  `json:"request_chat,omitempty" field:"request_chat" ifempty:"false"`
	RequestContact  bool                              `json:"request_contact,omitempty" field:"request_contact" ifempty:"false"`
	RequestLocation bool                              `json:"request_location,omitempty" field:"request_location" ifempty:"false"`
	RequestPoll     *types.KeyboardButtonPollType     `json:"request_poll,omitempty" field:"request_poll" ifempty:"false"`
	WebApp          *types.WebAppInfo                 `json:"web_app,omitempty" field:"web_app" ifempty:"false"`
	storage         [6]int
}

type reply struct {
	Keyboard              [][]*replyKeyboardButton `json:"keyboard,omitempty" field:"keyboard" ifempty:"false"`
	IsPersistent          bool                     `json:"is_persistent,omitempty" field:"is_persistent" ifempty:"false"`
	ResizeKeyboard        bool                     `json:"resize_keyboard,omitempty" field:"resize_keyboard" ifempty:"false"`
	OneTimeKeyboard       bool                     `json:"one_time_keyboard,omitempty" field:"one_time_keyboard" ifempty:"false"`
	InputFieldPlaceholder string                   `json:"input_field_placeholder,omitempty" field:"input_field_placeholder" ifempty:"false"`
	Selective             bool                     `json:"selective,omitempty" field:"selective" ifempty:"false"`
	Remove                bool                     `json:"remove_keyboard,omitempty" field:"remove_keyboard" ifempty:"false"`
}

type inlinemode struct {
	QueryID           string                          `json:"inline_query_id,omitempty" field:"inline_query_id" ifempty:"false"`
	WebAppQueryID     string                          `json:"web_app_query_id,omitempty" field:"web_app_query_id" ifempty:"false"`
	Result            interface{}                     `json:"result,omitempty" field:"result" ifempty:"false"`
	Results           []interface{}                   `json:"results,omitempty" field:"results" ifempty:"false"`
	CacheTime         int                             `json:"cache_time,omitempty" field:"cache_time" ifempty:"false"`
	IsPersonal        bool                            `json:"is_personal,omitempty" field:"is_personal" ifempty:"false"`
	NextOffset        string                          `json:"next_offset,omitempty" field:"next_offset" ifempty:"false"`
	Button            *types.InlineQueryResultsButton `json:"button,omitempty" field:"button" ifempty:"false"`
	AllowUserChats    bool                            `json:"allow_user_chats,omitempty" field:"allow_user_chats" ifempty:"false"`
	AllowBotChats     bool                            `json:"allow_bot_chats,omitempty" field:"allow_bot_chats" ifempty:"false"`
	AllowGroupChats   bool                            `json:"allow_group_chats,omitempty" field:"allow_group_chats" ifempty:"false"`
	AllowChannelChats bool                            `json:"allow_channel_chats,omitempty" field:"allow_channel_chats" ifempty:"false"`
	webAppResponse    *types.SentWebAppMessage
	inMessageResponse *types.PreparedInlineMessage
	resultcounter     int
	onlytothearray    bool
}

type payment struct {
	Title                     string                  `json:"title,omitempty" field:"title" ifempty:"false"`
	Description               string                  `json:"description,omitempty" field:"description" ifempty:"false"`
	Payload                   string                  `json:"payload,omitempty" field:"payload" ifempty:"false"`
	ProviderToken             *string                 `json:"provider_token,omitempty" field:"provider_token" ifempty:"false"`
	Currency                  string                  `json:"currency,omitempty" field:"currency" ifempty:"false"`
	Prices                    []*types.LabeledPrice   `json:"prices,omitempty" field:"prices" ifempty:"false"`
	MaxTipAmount              int                     `json:"max_tip_amount,omitempty" field:"max_tip_amount" ifempty:"false"`
	SuggestedTipAmounts       []int                   `json:"suggested_tip_amounts,omitempty" field:"suggested_tip_amounts" ifempty:"false"`
	StartParameter            *string                 `json:"start_parameter,omitempty" field:"start_parameter" ifempty:"false"`
	ProviderData              string                  `json:"provider_data,omitempty" field:"provider_data" ifempty:"false"`
	PhotoUrl                  string                  `json:"photo_url,omitempty" field:"photo_url" ifempty:"false"`
	PhotoSize                 int                     `json:"photo_size,omitempty" field:"photo_size" ifempty:"false"`
	PhotoWidth                int                     `json:"photo_width,omitempty" field:"photo_width" ifempty:"false"`
	PhotoHeight               int                     `json:"photo_height,omitempty" field:"photo_height" ifempty:"false"`
	NeedName                  bool                    `json:"need_name,omitempty" field:"need_name" ifempty:"false"`
	NeedPhoneNumber           bool                    `json:"need_phone_number,omitempty" field:"need_phone_number" ifempty:"false"`
	NeedEmail                 bool                    `json:"need_email,omitempty" field:"need_email" ifempty:"false"`
	NeedShippingAddress       bool                    `json:"need_shipping_address,omitempty" field:"need_shipping_address" ifempty:"false"`
	SendPhoneNumberToProvider bool                    `json:"send_phone_number_to_provider,omitempty" field:"send_phone_number_to_provider" ifempty:"false"`
	SendEmailToProvider       bool                    `json:"send_email_to_provider,omitempty" field:"send_email_to_provider" ifempty:"false"`
	IsFlexible                bool                    `json:"is_flexible,omitempty" field:"is_flexible" ifempty:"false"`
	SubscriptionPeriod        int                     `json:"subscription_period,omitempty" field:"subscription_period" ifempty:"false"`
	ShippingID                string                  `json:"shipping_query_id,omitempty" field:"shipping_query_id" ifempty:"false"`
	OK                        *bool                   `json:"ok,omitempty" field:"ok" ifempty:"false"`
	ShippingOptions           []*types.ShippingOption `json:"shipping_options,omitempty" field:"shipping_options" ifempty:"false"`
	ErrorMessage              string                  `json:"error_message,omitempty" field:"error_message" ifempty:"false"`
	PreCheckoutID             string                  `json:"pre_checkout_query_id,omitempty" field:"pre_checkout_query_id" ifempty:"false"`
	TelegramPaymentChargeID   string                  `json:"telegram_payment_charge_id,omitempty" field:"telegram_payment_charge_id" ifempty:"false"`
	IsCanceled                bool                    `json:"is_canceled,omitempty" field:"is_canceled" ifempty:"false"`
}

type game struct {
	ShortName          string `json:"game_short_name,omitempty" field:"game_short_name" ifempty:"false"`
	Score              int    `json:"score,omitempty" field:"score" ifempty:"false"`
	Force              bool   `json:"force,omitempty" field:"force" ifempty:"false"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty" field:"disable_edit_message" ifempty:"false"`
}

type mediagroup struct {
	id     string
	photos [][]*types.PhotoSize
	videos []*types.Video
	audios []*types.Audio
	docs   []*types.Document
}

type get struct {
	status     bool
	errorCode  int
	errorMsg   string
	chat       *types.Chat
	sender     *types.User
	date       int
	msgID      int
	msgIDs     []int
	replyed    *get
	msgOrigin  *types.MessageOrigin
	photo      []*types.PhotoSize
	audio      *types.Audio
	document   *types.Document
	video      *types.Video
	anim       *types.Animation
	voice      *types.Voice
	vdn        *types.VideoNote
	paid       *types.PaidMedia
	mg         *mediagroup
	poll       *types.Poll
	dice       *types.Dice
	uprph      *types.UserProfilePhotos
	file       *types.File
	stickers   []*types.Sticker
	gifts      []*types.Gift
	msg        *types.Message
	str        string
	invlink    *types.ChatInviteLink
	chatinfo   *types.ChatFullInfo
	members    []*types.ChatMember
	integer    *int
	forum      *types.ForumTopic
	boosts     []*types.ChatBoost
	bconn      *types.BusinessConnection
	botcomm    []*types.BotCommand
	menuButton *types.MenuButton
	admin      *types.ChatAdministratorRights
	stickerset *types.StickerSet
	msgs       []*types.Message
	prepinlmsg *types.PreparedInlineMessage
	startrans  *types.StarTransaction
	score      []*types.GameHighScore
	request    string
	response   string
}
