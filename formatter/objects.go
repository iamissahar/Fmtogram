package formatter

import "github.com/l1qwie/Fmtogram/types"

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

// type sticker struct {
// 	Type       string `json:"type,omitempty"`
// 	Media      string `json:"media,omitepty"`
// 	Sticker    string `json:"sticker,omitempty"`
// 	Emoji      string `json:"emoji,omitempty"`
// 	gottenFrom int
// 	response   types.Sticker
// }

type information struct {
	MessageID             int                       `json:"message_id,omitempty"`
	MessageIDs            []int                     `json:"message_ids,omitempty"`
	Text                  string                    `json:"text,omitempty"`
	Caption               string                    `json:"caption,omitempty"`
	ParseMode             string                    `json:"parse_mode,omitempty"`
	MessageThreadID       int                       `json:"message_thread_id,omitempty"`
	Entities              []*types.MessageEntity    `json:"entities,omitempty"`
	CaptionEntities       []*types.MessageEntity    `json:"caption_entities,omitempty"`
	LinkPreviewOptions    *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification   bool                      `json:"disable_notification,omitempty"`
	ProtectContent        bool                      `json:"protect_content,omitempty"`
	MessageEffectID       string                    `json:"message_effect_id,omitempty"`
	ShowCaptionAboveMedia bool                      `json:"show_caption_above_media,omitempty"`
	ReplyParameters       *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	AllowPaidBroadcast    bool                      `json:"allow_paid_broadcast,omitempty"`
	StarCount             int                       `json:"star_count,omitempty"`
	Payload               string                    `json:"payload,omitempty"`
	response              types.User
	responseMessageIDs    []int
}

type chat struct {
	FromChatID           interface{} `json:"from_chat_id,omitempty"`
	ID                   interface{} `json:"chat_id,omitempty"`
	BusinessConnectionID string      `json:"business_connection_id,omitempty"`
	response             types.Chat
}

type inlineKeyboardButton struct {
	Text                         string                             `json:"text"`
	Url                          string                             `json:"url,omitempty"`
	CallbackData                 string                             `json:"callback_data,omitempty"`
	WebApp                       *types.WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *types.LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                             `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                             `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *types.SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 *types.CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                               `json:"pay,omitempty"`
	storage                      [9]int
}

type inline struct {
	Keyboard *inlineKeyboard `json:"reply_markup,omitempty"`
}

type inlineKeyboard struct {
	InlineKeyboard [][]*inlineKeyboardButton `json:"inline_keyboard"`
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
	Keyboard *replyKeyboard `json:"reply_markup,omitempty"`
}

type replyKeyboard struct {
	Keyboard              [][]*replyKeyboardButton `json:"keyboard"`
	IsPersistent          bool                     `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool                     `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool                     `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string                   `json:"input_field_placeholder,omitempty"`
	Selective             bool                     `json:"selective,omitempty"`
}
