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
	Ok     bool        `json:"ok"`
	Result *BotCommand `json:"result"`
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
