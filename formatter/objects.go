package formatter

import (
	"time"

	"github.com/l1qwie/Fmtogram/types"
)

type photo struct {
	Type                  string                 `json:"type,omitempty"`
	Media                 string                 `json:"media,omitempty"`
	Photo                 string                 `json:"photo,omitempty"`
	Caption               string                 `json:"caption,omitempty"`
	ParseMode             string                 `json:"parse_mode,omitempty"`
	CaptionEntities       []*types.MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                   `json:"has_spoiler,omitempty"`
	gottenFrom            int
	response              [4]types.PhotoSize
}

type video struct {
	Type                  string                 `json:"type,omitempty"`
	Media                 string                 `json:"media,omitempty"`
	Video                 string                 `json:"video,omitempty"`
	Thumbnail             string                 `json:"thumbnail,omitempty"`
	Caption               string                 `json:"caption,omitempty"`
	ParseMode             string                 `json:"parse_mode,omitempty"`
	CaptionEntities       []*types.MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty"`
	Width                 int                    `json:"width,omitempty"`
	Height                int                    `json:"height,omitempty"`
	Duration              int                    `json:"duration,omitempty"`
	SupportsStreaming     bool                   `json:"supports_streaming,omitempty"`
	HasSpoiler            bool                   `json:"has_spoiler,omitempty"`
	videoGottenFrom       int
	thumbnailGottenFrom   int
	response              types.Video
}

type audio struct {
	Type                string                 `json:"type,omitempty"`
	Media               string                 `json:"media,omitempty"`
	Audio               string                 `json:"audio,omitempty"`
	Thumbnail           string                 `json:"thumbnail,omitempty"`
	Caption             string                 `json:"caption,omitempty"`
	ParseMode           string                 `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.MessageEntity `json:"caption_entities,omitempty"`
	Duration            int                    `json:"duration,omitempty"`
	Performer           string                 `json:"performer,omitempty"`
	Title               string                 `json:"title,omitempty"`
	audioGottenFrom     int
	thumbnailGottenFrom int
	response            types.Audio
}

type document struct {
	Type                        string                 `json:"type,omitempty"`
	Media                       string                 `json:"media,omitempty"`
	Document                    string                 `json:"document,omitempty"`
	Thumbnail                   string                 `json:"thumbnail,omitempty"`
	Caption                     string                 `json:"caption,omitempty"`
	ParseMode                   string                 `json:"parse_mode,omitempty"`
	CaptionEntities             []*types.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
	documentGottenFrom          int
	thumbnailGottenFrom         int
	response                    types.Document
}

type animation struct {
	Type                string `json:"type,omitempty"`
	Media               string `json:"media,omitempty"`
	Animation           string `json:"animation,omitempty"`
	Thumbnail           string `json:"thumbnail,omitempty"`
	Width               int    `json:"width,omitempty"`
	Height              int    `json:"height,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	HasSpoiler          bool   `json:"has_spoiler,omitempty"`
	animationGottenFrom int
	thumbnailGottenFrom int
	response            types.Animation
}

type voice struct {
	Type       string `json:"type,omitempty"`
	Media      string `json:"media,omitempty"`
	Voice      string `json:"voice,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	gottenFrom int
	response   types.Voice
}

type videonote struct {
	Type                string `json:"type,omitempty"`
	Media               string `json:"media,omitempty"`
	VideoNote           string `json:"video_note,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	Length              int    `json:"length,omitempty"`
	Thumbnail           string `json:"thumbnail,omitempty"`
	videoGottenFrom     int
	thumbnailGottenFrom int
	response            types.VideoNote
}

type location struct {
	Latitude             float64 `json:"latitude,omitempty"`
	Longitude            float64 `json:"longitude,omitempty"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
	Title                string  `json:"title,omitempty"`
	Address              string  `json:"address,omitempty"`
	FoursquareID         string  `json:"foursquare_id,omitempty"`
	FoursquareType       string  `json:"foursquare_type,omitempty"`
	GooglePlaceID        string  `json:"google_place_id,omitempty"`
	GooglePlaceType      string  `json:"google_place_type,omitempty"`
	response             types.Venue
}

type contact struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
	response    types.Contact
}

type poll struct {
	Question             string                 `json:"question,omitempty"`
	QuestionParsemode    string                 `json:"question_parse_mode,omitempty"`
	QuestionEntities     []*types.MessageEntity `json:"question_entities,omitempty"`
	Options              []*types.PollOption    `json:"options,omitempty"`
	IsAnonymous          bool                   `json:"is_anonymous,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	AllowMultipleAnswers bool                   `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID      string                 `json:"correct_option_id,omitempty"`
	Explanation          string                 `json:"explanation,omitempty"`
	ExplanationParsemode string                 `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities  []*types.MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod           int                    `json:"open_period,omitempty"`
	CloseDate            int                    `json:"close_date,omitempty"`
	IsClosed             bool                   `json:"is_closed,omitempty"`
	response             types.Poll
}

type sticker struct {
	SetName             string              `json:"sticker_set_name,omitempty"`
	Sticker             string              `jsno:"sticker"`
	Stickers            [50]string          `json:"stickers"`
	Name                string              `json:"name"`
	Emoji               string              `json:"emoji"`
	Emojies             []string            `json:"emoji_list"`
	EmojiID             string              `json:"custom_emoji_id"`
	EmojiIDs            []string            `json:"custom_emoji_ids"`
	Format              string              `json:"sticker_format"`
	Title               string              `json:"title"`
	StickerType         string              `json:"sticker_type"`
	NeedsRepainting     bool                `json:"needs_repainting"`
	Position            string              `json:"position"`
	OldSticker          string              `json:"old_sticker"`
	Keywords            []string            `json:"keywords"`
	MaskPosition        *types.MaskPosition `json:"mask_position"`
	Thumbnail           string              `json:"thumbnail"`
	ThumbnailFormat     string              `json:"format"`
	GiftID              string              `json:"gift_id"`
	PayForUpgrade       bool                `json:"pay_for_upgrade"`
	stickerGottenFrom   int
	thumbnailGottenFrom int
	thumbnailType       string
}

type forum struct {
	Name        string `json:"name,omitempty"`
	IconColor   int    `json:"icon_color,omitempty"`
	IconEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type information struct {
	MessageID                  int                            `json:"message_id,omitempty"`
	MessageIDs                 []int                          `json:"message_ids,omitempty"`
	Text                       string                         `json:"text,omitempty"`
	Caption                    string                         `json:"caption,omitempty"`
	ParseMode                  string                         `json:"parse_mode,omitempty"`
	MessageThreadID            int                            `json:"message_thread_id,omitempty"`
	Entities                   []*types.MessageEntity         `json:"entities,omitempty"`
	CaptionEntities            []*types.MessageEntity         `json:"caption_entities,omitempty"`
	LinkPreviewOptions         *types.LinkPreviewOptions      `json:"link_preview_options,omitempty"`
	DisableNotification        bool                           `json:"disable_notification,omitempty"`
	ProtectContent             bool                           `json:"protect_content,omitempty"`
	MessageEffectID            string                         `json:"message_effect_id,omitempty"`
	ShowCaptionAboveMedia      bool                           `json:"show_caption_above_media,omitempty"`
	ReplyParameters            *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	AllowPaidBroadcast         bool                           `json:"allow_paid_broadcast,omitempty"`
	StarCount                  int                            `json:"star_count,omitempty"`
	Payload                    string                         `json:"payload,omitempty"`
	Emoji                      string                         `json:"emoji,omitempty"`
	Action                     string                         `json:"action,omitempty"`
	Reaction                   []*types.ReactionType          `json:"reaction,omitempty"`
	ReactionIsBig              bool                           `json:"is_big,omitempty"`
	Offset                     int                            `json:"offset,omitempty"`
	Limit                      int                            `json:"limit,omitempty"`
	EmojiStatusCustomEmojiID   string                         `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate  int                            `json:"emoji_status_expiration_date,omitempty"`
	FileID                     string                         `json:"file_id,omitempty"`
	UntilDate                  time.Duration                  `json:"until_date,omitempty"`
	RevokeMessages             bool                           `json:"revoke_messages,omitempty"`
	OnlyIfBanned               bool                           `json:"only_if_banned,omitempty"`
	Permissions                *types.ChatPermissions         `json:"permissions,omitempty"`
	IndependentChatPermissions bool                           `json:"use_independent_chat_permissions,omitempty"`
	AdminRights                *types.ChatAdministratorRights `json:",omitempty"`
	CustomTitle                string                         `json:"custom_title,omitempty"`
	UserID                     int                            `json:"user_id,omitempty"`
	CallBackQueryID            string                         `json:"callback_query_id,omitempty"`
	ShowAlert                  bool                           `json:"show_alert,omitempty"`
	Url                        string                         `json:"url,omitempty"`
	CacheTime                  int                            `json:"cache_time,omitempty"`
	InlineMessageID            string                         `json:"inline_message_id,omitempty"`
	Errors                     []*types.PassportElementError  `json:"errors,omitempty"`
	GiftParseMode              string                         `json:"text_parse_mode,omitempty"`
	GiftEntities               []*types.MessageEntity         `json:"text_entities,omitempty"`
	Description                string                         `json:"custom_description,omitempty"`
	RemoveCaption              bool                           `json:"remove_caption,omitempty"`
	response                   types.User
	responseMessageIDs         []int
}

type chat struct {
	FromChatID           interface{} `json:"from_chat_id,omitempty"`
	ID                   interface{} `json:"chat_id,omitempty"`
	BusinessConnectionID string      `json:"business_connection_id,omitempty"`
	SenderChatID         int         `json:"sender_chat_id,omitempty"`
	Title                string      `json:"title,omitempty"`
	Description          string      `json:"description,omitempty"`
	response             types.Chat
}

type link struct {
	Name               string        `json:"name,omitempty"`
	ExpireDate         time.Duration `json:"expire_date,omitempty"`
	MemberLimit        int           `json:"member_limit,omitempty"`
	JoinRequest        bool          `json:"creates_join_request,omitempty"`
	InviteLink         string        `json:"invite_link,omitempty"`
	SubscriptionPeriod int           `json:"subscription_period,omitempty"`
	SubscriptionPrice  int           `json:"subscription_price,omitempty"`
}

type bot struct {
	Commands         []*types.BotCommand            `json:"commands,omitempty"`
	Scope            *types.BotCommandScope         `json:"scope,omitempty"`
	Language         string                         `json:"language_code,omitempty"`
	Name             string                         `json:"name,omitempty"`
	Description      string                         `json:"description,omitempty"`
	ShortDescription string                         `json:"short_description,omitempty"`
	MenuButtom       *types.MenuButton              `json:"menu_button,omitempty"`
	Rights           *types.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels      bool                           `json:"for_channels,omitempty"`
}

type inlineKeyboardButton struct {
	Text                         string                             `json:"text"`
	Url                          string                             `json:"url,omitempty"`
	CallbackData                 string                             `json:"callback_data,omitempty"`
	WebApp                       *types.WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *types.LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            *string                            `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                            `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *types.SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 *types.CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                               `json:"pay,omitempty"`
	storage                      [9]int
}

type keyboard struct {
	Keyboard kb `json:"reply_markup,omitempty"`
}

type forcereply struct {
	Forcereply            bool   `json:"force_reply,omitempty"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

type inline struct {
	Keyboard [][]*inlineKeyboardButton `json:"inline_keyboard"`
}

type replyKeyboardButton struct {
	Text            string                            `json:"text"`
	RequestUsers    *types.KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *types.KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                              `json:"request_contact,omitempty"`
	RequestLocation bool                              `json:"request_location,omitempty"`
	RequestPoll     *types.KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *types.WebAppInfo                 `json:"web_app,omitempty"`
	storage         [6]int
}

type reply struct {
	Keyboard              [][]*replyKeyboardButton `json:"keyboard"`
	IsPersistent          bool                     `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool                     `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool                     `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string                   `json:"input_field_placeholder,omitempty"`
	Selective             bool                     `json:"selective,omitempty"`
	Remove                bool                     `json:"remove_keyboard,omitempty"`
}

type inlinemode struct {
	QueryID           string                          `json:"inline_query_id"`
	WebAppQueryID     string                          `json:"web_app_query_id"`
	Result            IResult                         `json:"result"`
	Results           []IResult                       `json:"results"`
	CacheTime         int                             `json:"cache_time"`
	IsPersonal        bool                            `json:"is_personal"`
	NextOffset        string                          `json:"next_offset"`
	Button            *types.InlineQueryResultsButton `json:"button"`
	AllowUserChats    bool                            `json:"allow_user_chats"`
	AllowBotChats     bool                            `json:"allow_bot_chats"`
	AllowGroupChats   bool                            `json:"allow_group_chats"`
	AllowChannelChats bool                            `json:"allow_channel_chats"`
	webAppResponse    *types.SentWebAppMessage
	inMessageResponse *types.PreparedInlineMessage
}

type result struct {
	CachedAudio    *types.InlineQueryResultCachedAudio
	CachedDocument *types.InlineQueryResultCachedDocument
	CachedGif      *types.InlineQueryResultCachedGif
	CachedMpeg4Gif *types.InlineQueryResultCachedMpeg4Gif
	CachedPhoto    *types.InlineQueryResultCachedPhoto
	CachedSticker  *types.InlineQueryResultCachedSticker
	CachedVideo    *types.InlineQueryResultCachedVideo
	CachedVoice    *types.InlineQueryResultCachedVoice
	Article        *types.InlineQueryResultArticle
	Audio          *types.InlineQueryResultAudio
	Contact        *types.InlineQueryResultContact
	Game           *types.InlineQueryResultGame
	Document       *types.InlineQueryResultDocument
	Gif            *types.InlineQueryResultGif
	Location       *types.InlineQueryResultLocation
	Mpeg4Gif       *types.InlineQueryResultMpeg4Gif
	Photo          *types.InlineQueryResultPhoto
	Venue          *types.InlineQueryResultVenue
	Video          *types.InlineQueryResultVideo
	Voice          *types.InlineQueryResultVoice
}

type payment struct {
	Title                     string                  `json:"title"`
	Description               string                  `json:"description"`
	Payload                   string                  `json:"payload"`
	ProviderToken             string                  `json:"provider_token"`
	Currency                  string                  `json:"currency"`
	Prices                    []*types.LabeledPrice   `json:"prices"`
	MaxTipAmount              int                     `json:"max_tip_amount"`
	SuggestedTipAmounts       []int                   `json:"suggested_tip_amounts"`
	StartParameter            string                  `json:"start_parameter"`
	ProviderData              string                  `json:"provider_data"`
	PhotoUrl                  string                  `json:"photo_url"`
	PhotoSize                 int                     `json:"photo_size"`
	PhotoWidth                int                     `json:"photo_width"`
	PhotoHeight               int                     `json:"photo_height"`
	NeedName                  bool                    `json:"need_name"`
	NeedPhoneNumber           bool                    `json:"need_phone_number"`
	NeedEmail                 bool                    `json:"need_email"`
	NeedShippingAddress       bool                    `json:"need_shipping_address"`
	SendPhoneNumberToProvider bool                    `json:"send_phone_number_to_provider"`
	SendEmailToProvider       bool                    `json:"send_email_to_provider"`
	IsFlexible                bool                    `json:"is_flexible"`
	SubscriptionPeriod        int                     `json:"subscription_period"`
	ShippingID                string                  `json:"shipping_query_id"`
	OK                        bool                    `json:"ok"`
	ShippingOptions           []*types.ShippingOption `json:"shipping_options"`
	ErrorMessage              string                  `json:"error_message"`
	PreCheckoutID             string                  `json:"pre_checkout_query_id"`
	Offset                    int                     `json:"offset"`
	Limit                     int                     `json:"limit"`
	TelegramPaymentChargeID   string                  `json:"telegram_payment_charge_id"`
	IsCanceled                bool                    `json:"is_canceled"`
}

type game struct {
	ShortName          string `json:"game_short_name,omitempty"`
	Score              int    `json:"score,omitempty"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
}

type mediagroup struct {
	id     int
	photos [][]types.PhotoSize
	videos []types.Video
	audios []types.Audio
	docs   []types.Document
}

type get struct {
	status    bool
	errorCode int
	errorMsg  string
	chat      types.Chat
	user      types.User
	bot       types.User
	date      int
	msgID     int
	msgIDs    []int
	replyed   *get
	msgOrigin types.MessageOrigin
	photo     []types.PhotoSize
	audio     types.Audio
	document  types.Document
	video     types.Video
	anim      types.Animation
	voice     types.Voice
	vdn       types.VideoNote
	paid      types.PaidMedia
	mg        *mediagroup
}
