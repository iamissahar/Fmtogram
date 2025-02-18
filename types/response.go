package types

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

type Telegram struct {
	Ok     bool      `json:"ok"`
	Result []*Update `json:"result,omitempty"`
}

type GetMe struct {
	Ok     bool  `json:"ok"`
	Result *User `json:"result,omitempty"`
}

type MessageResponse struct {
	Ok     bool     `json:"ok"`
	Result *Message `json:"result,omitempty"`
}

type MediaGroupResponse struct {
	Ok     bool       `json:"ok"`
	Result []*Message `json:"result,omitempty"`
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
