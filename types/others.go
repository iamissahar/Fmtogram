package types

var (
	BotID             string
	DEBUG             bool
	INFO              bool
	NoWarningMessages bool
	// StartFunc func(*Telegram, *Returned) *formatter.Formatter
)

const (
	Markdown    string = "Markdown"
	HTML        string = "HTML"
	TelegramAPI string = "https://api.telegram.org/"
)

type ForumTopicClosed struct {
	// Currently holds no information
}

type ForumTopicReopened struct {
	// Currently holds no information
}

type GeneralForumTopicHidden struct {
	// Currently holds no information
}

type GeneralForumTopicUnhidden struct {
	// Currently holds no information
}

type GiveawayCreated struct {
	// Currently holds no information
}

type VideoChatStarted struct {
	// This object represents a service message about a video chat started in the chat. Currently holds no information.
}

type CallbackGame struct {
	// A placeholder, currently holds no information
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	URL              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type WebAppInfo struct {
	Url string `json:"url"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type TextQuote struct {
	Text     string           `json:"text"`
	Entities []*MessageEntity `json:"entities"`
	Position int              `json:"position"`
	IsManual bool             `json:"is_manual"`
}

type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePlaceID   string    `json:"google_place_id,omitempty"`
	GooglePlaceType string    `json:"google_place_type,omitempty"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`
	Colors []int  `json:"colors"`
}

type BackgroundFill struct {
	BackgroundFillSolid            *BackgroundFillSolid
	BackgroundFillGradient         *BackgroundFillGradient
	BackgroundFillFreeformGradient *BackgroundFillFreeformGradient
}

type Giveaway struct {
	Chats                         []*Chat  `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
}

type GiveawayWinners struct {
	Chat                          *Chat   `json:"chat"`
	GiveawayMessageID             int     `json:"giveaway_message_id"`
	WinnersSelectionDate          int     `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []*User `json:"winners"`
	AdditionalChatCount           int     `json:"additional_chat_count,omitempty"`
	PremiumSubscriptionMonthCount int     `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int     `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool    `json:"only_new_members,omitempty"`
	WasRefunded                   bool    `json:"was_refunded,omitempty"`
	PrizeDescription              string  `json:"prize_description,omitempty"`
}

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Files       []*PassportFile `json:"files"`
	FrontSide   *PassportFile   `json:"front_side"`
	ReverseSide *PassportFile   `json:"reverse_side"`
	Selfie      *PassportFile   `json:"selfie"`
	Translation []*PassportFile `json:"translation"`
	Hash        string          `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YSHift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities"`
	Animation    *Animation       `json:"animation"`
}

type PollOption struct {
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int              `json:"voter_count"`
}

type InputPollOption struct {
	Text          string           `json:"text"`
	TextParseMode string           `json:"text_parse_mode,omitempty"`
	TextEntities  []*MessageEntity `json:"text_entities,omitempty"`
}

type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	QuestionEnities       []*MessageEntity `json:"question_entities,omitempty"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`
	Chat               *Chat               `json:"chat"`
	MessageID          int                 `json:"message_id"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`
	Animation          *Animation          `json:"animation"`
	Audio              *Audio              `json:"audio"`
	Document           *Document           `json:"document"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media"`
	Photo              []*PhotoSize        `json:"photo"`
	Sticker            *Sticker            `json:"sticker"`
	Story              *Story              `json:"story"`
	Video              []*Video            `json:"video"`
	VideoNote          *VideoNote          `json:"video_note"`
	Voice              *Voice              `json:"voice"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler"`
	Contact            *Contact            `json:"contact"`
	Dice               *Dice               `json:"dice"`
	Game               *Game               `json:"game"`
	Giveaway           *Giveaway           `json:"giveaway"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners"`
	Invoice            *Invoice            `json:"invoice"`
	Location           *Location           `json:"location"`
	Poll               *Poll               `json:"poll"`
	Venue              *Venue              `json:"venue"`
}

type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type ReactionTypePaid struct {
	Type string `json:"type"`
}

type ReactionType struct {
	ReactionTypeEmoji       *ReactionTypeEmoji
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji
	ReactionTypePaid        *ReactionTypePaid
}

type ReactionCount struct {
	Type       *ReactionType `json:"type"`
	TotalCount int           `json:"total_count"`
}

type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`
}

type ReplyParameters struct {
	MessageID                int              `json:"message_id"`
	ChatID                   interface{}      `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"`
	Quote                    string           `json:"quote,omitempty"`
	QuoteParseMode           string           `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []*MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int              `json:"quote_position,omitempty"`
}

type MenuButtonCommands struct {
	Type string `json:"type"`
}

type MenuButtonWebApp struct {
	Type   string      `json:"type"`
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

type MenuButton struct {
	MenuButtonCommands *MenuButtonCommands
	MenuButtonWebApp   *MenuButtonWebApp
	MenuButtonDefault  *MenuButtonDefault
}

type Update struct {
	UpdateID               int                          `json:"update_id"`
	Message                *Message                     `json:"message"`
	EditedMessage          *Message                     `json:"edited_message"`
	ChanPost               *Message                     `json:"channel_post"`
	EditedChanPost         *Message                     `json:"edited_channel_post"`
	BusinessConn           *BusinessConnection          `json:"business_connection"`
	BusinessMessage        *Message                     `json:"business_message"`
	EditedBusinessMessage  *Message                     `json:"edited_business_message"`
	DeletedBusinessMessage *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	MessageReaction        *MessageReactionUpdated      `json:"message_reaction"`
	MessageReactionCount   *MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQuery            *InlineQuery                 `json:"inline_query"`
	ChosenInlineResult     *ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQuery          *CallbackQuery               `json:"callback_query"`
	ShippingQuery          *ShippingQuery               `json:"shipping_query"`
	PreCheckoutQuery       *PreCheckoutQuery            `json:"pre_checkout_query"`
	Poll                   *Poll                        `json:"poll"`
	PollAnswer             *PollAnswer                  `json:"poll_answer"`
	MyChatMember           *ChatMemberUpdated           `json:"my_chat_member"`
	ChatMember             *ChatMemberUpdated           `json:"chat_member"`
	ChatJoinRequest        *ChatJoinRequest             `json:"chat_join_request"`
	ChatBoost              *ChatBoostUpdated            `json:"chat_boost"`
	RemovedChatBoost       *ChatBoostRemoved            `json:"removed_chat_boost"`
}

type TelegramError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Telegram struct {
	Ok          bool           `json:"ok"`
	Result      []*Update      `json:"result,omitempty"`
	Error       *TelegramError `json:"error,omitempty"`
	ErrorCode   int            `json:"error_code,omitempty"`
	Description string         `json:"description,omitempty"`
}

type GetMe struct {
	Ok     bool           `json:"ok"`
	Result *User          `json:"result,omitempty"`
	Error  *TelegramError `json:"error,omitempty"`
}

type BadResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type MessageResponse struct {
	Ok          bool           `json:"ok"`
	Result      *Message       `json:"result,omitempty"`
	Error       *TelegramError `json:"error,omitempty"`
	ErrorCode   int            `json:"error_code,omitempty"`
	Description string         `json:"description,omitempty"`
}
