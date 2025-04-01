package fmtogram

func (g *get) Status() bool {
	return g.status
}

func (g *get) Error() (int, string) {
	return g.errorCode, g.errorMsg
}

func (g *get) Chat() *Chat {
	return g.chat
}

func (g *get) Date() int {
	return g.date
}

func (g *get) MessageID() int {
	return g.msgID
}

func (g *get) MessageIDs() []int {
	return g.msgIDs
}

func (g *get) Sender() *User {
	return g.sender
}

func (g *get) Replyed() IGet {
	return g.replyed
}

func (g *get) ForwardOrigin() *MessageOrigin {
	return g.msgOrigin
}

func (g *get) Photo() []*PhotoSize {
	return g.photo
}

func (g *get) Audio() *Audio {
	return g.audio
}

func (g *get) Document() *Document {
	return g.document
}

func (g *get) Video() *Video {
	return g.video
}

func (g *get) Animation() *Animation {
	return g.anim
}

func (g *get) Voice() *Voice {
	return g.voice
}

func (g *get) VideoNote() *VideoNote {
	return g.vdn
}

func (g *get) PaidMedia() *PaidMedia {
	return g.paid
}

func (g *get) MediaGroupID() string {
	return g.mg.id
}

func (g *get) Photos() [][]*PhotoSize {
	return g.mg.photos
}

func (g *get) Videos() []*Video {
	return g.mg.videos
}

func (g *get) Audios() []*Audio {
	return g.mg.audios
}

func (g *get) Documents() []*Document {
	return g.mg.docs
}

func (g *get) Poll() *Poll {
	return g.poll
}

func (g *get) Dice() *Dice {
	return g.dice
}

func (g *get) ProfilePhotos() *UserProfilePhotos {
	return g.uprph
}

func (g *get) File() *File {
	return g.file
}

func (g *get) Stickers() []*Sticker {
	return g.stickers
}

func (g *get) Gifts() []*Gift {
	return g.gifts
}

func (g *get) Message() *TelegramMessage {
	return g.msg
}

func (g *get) Messages() []*TelegramMessage {
	return g.msgs
}

func (g *get) String() string {
	return g.str
}

func (g *get) InviteLink() *ChatInviteLink {
	return g.invlink
}

func (g *get) ChatInfo() *ChatFullInfo {
	return g.chatinfo
}

func (g *get) Members() []*ChatMember {
	return g.members
}

func (g *get) Integer() *int {
	return g.integer
}

func (g *get) Forum() *ForumTopic {
	return g.forum
}

func (g *get) Boosts() []*ChatBoost {
	return g.boosts
}

func (g *get) BusinessConnection() *BusinessConnection {
	return g.bconn
}

func (g *get) Commands() []*BotCommand {
	return g.botcomm
}

func (g *get) MenuButton() *MenuButton {
	return g.menuButton
}

func (g *get) AdminRights() *ChatAdministratorRights {
	return g.admin
}

func (g *get) PreparedInlineMessage() *PreparedInlineMessage {
	return g.prepinlmsg
}

func (g *get) StarTransaction() *StarTransaction {
	return g.startrans
}

func (g *get) Score() []*GameHighScore {
	return g.score
}

func (g *get) Request() string {
	return g.request
}

func (g *get) Response() string {
	return g.response
}

func (wh *webhook) GetStatus() bool {
	return wh.status
}

func (wh *webhook) GetError() (int, string) {
	return wh.errorcode, wh.errormsg
}
