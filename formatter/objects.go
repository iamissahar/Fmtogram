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
	GottenFrom            int
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
	VideoGottenFrom       int
	ThumbnailGottenFrom   int
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
	AudioGottenFrom     int
	ThumbnailGottenFrom int
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
	DocumentGottenFrom          int
	ThumbnailGottenFrom         int
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
	AnimationGottenFrom int
	ThumbnailGottenFrom int
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
	One                          int
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
