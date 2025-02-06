package types

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
