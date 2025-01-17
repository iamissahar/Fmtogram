package unit

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type paramsT struct {
	name          string
	integer       int
	str           string
	date          time.Duration
	array         []*types.MessageEntity
	arrayInt      []int
	link          *types.LinkPreviewOptions
	replyP        *types.ReplyParameters
	reaction      []*types.ReactionType
	permis        *types.ChatPermissions
	adminrights   *types.ChatAdministratorRights
	errors        []*types.PassportElementError
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type prmTestContainer struct {
	name          string
	inputInt      []int
	inputStr      []string
	inputDates    []time.Duration
	inputArrInt   [][]int
	inputArr      [][]*types.MessageEntity
	inputLink     []*types.LinkPreviewOptions
	inputReplyP   []*types.ReplyParameters
	inputReaction [][]*types.ReactionType
	inputPermis   []*types.ChatPermissions
	inputAdmin    []*types.ChatAdministratorRights
	inputErr      [][]*types.PassportElementError
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int)
}

func (prmtc prmTestContainer) defaultParamT(i int) *paramsT {
	return &paramsT{name: prmtc.name, isExpectedErr: prmtc.isExpectedErr[i], codeErr: prmtc.codeErr[i]}
}

func putParamWriteDisableNotification(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteDisableNotification
}

func putParamWriteEntities(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteEntities
	p.array = prmtc.inputArr[i]
}

func putParamWriteLinkPreviewOptions(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteLinkPreviewOptions
	p.link = prmtc.inputLink[i]
}

func putParamWriteMessageEffectID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteMessageEffectID
	p.str = prmtc.inputStr[i]
}

func putParamWriteMessageThreadID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteMessageThreadID
	p.integer = prmtc.inputInt[i]
}

func putParamWriteMessageID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteMessageID
	p.integer = prmtc.inputInt[i]
}

func putParamWriteMessageIDs(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteMessageIDs
	p.arrayInt = prmtc.inputArrInt[i]
}

func putParamWriteCaption(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteCaption
	p.str = prmtc.inputStr[i]
}

func putParamWriteParseMode(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteParseMode
	p.str = prmtc.inputStr[i]
}

func putParamWriteProtectContent(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteProtectContent
}

func putParamWriteString(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteString
	p.str = prmtc.inputStr[i]
}

func putParamWriteShowCaptionAboveMedia(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteShowCaptionAboveMedia
}

func putParamWriteReplyParameters(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteReplyParameters
	p.replyP = prmtc.inputReplyP[i]
}

func putParamWriteAllowPaidBroadcast(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteAllowPaidBroadcast
}

func putParamWriteStarCount(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteStarCount
	p.integer = prmtc.inputInt[i]
}

func putParamWritePayload(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WritePayload
	p.str = prmtc.inputStr[i]
}

func putParamWriteEmoji(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteEmoji
	p.str = prmtc.inputStr[i]
}

func putParamWriteAction(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteAction
	p.str = prmtc.inputStr[i]
}

func putParamWriteReaction(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteReaction
	p.reaction = prmtc.inputReaction[i]
}

func putParamWriteReactionIsBig(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteReactionIsBig
}

func putParamWriteOffset(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteOffset
	p.integer = prmtc.inputInt[i]
}

func putParamWriteLimit(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteLimit
	p.integer = prmtc.inputInt[i]
}

func putParamWriteEmojiStatusCustomEmojiID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteEmojiStatusCustomEmojiID
	p.str = prmtc.inputStr[i]
}

func putParamWriteEmojiStatusExpirationDate(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteEmojiStatusExpirationDate
	p.integer = prmtc.inputInt[i]
}

func putParamWriteFileID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteFileID
	p.str = prmtc.inputStr[i]
}

func putParamWriteUntilDate(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteUntilDate
	p.date = prmtc.inputDates[i]
}

func putParamWriteRevokeMessages(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteRevokeMessages
}

func putParamWriteOnlyIfBanned(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteOnlyIfBanned
}

func putParamWritePermissions(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WritePermissions
	p.permis = prmtc.inputPermis[i]
}

func putParamWriteIndependentChatPermissions(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteIndependentChatPermissions
}

func putParamWriteAdministratorRights(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteAdministratorRights
	p.adminrights = prmtc.inputAdmin[i]
}

func putParamWriteCustomTitle(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteCustomTitle
	p.str = prmtc.inputStr[i]
}

func putParamWriteUserID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteUserID
	p.integer = prmtc.inputInt[i]
}

func putParamWriteCallBackQueryID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteCallBackQueryID
	p.str = prmtc.inputStr[i]
}

func putParamWriteShowAlert(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteShowAlert
}

func putParamWriteURL(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteURL
	p.str = prmtc.inputStr[i]
}

func putParamWriteCacheTime(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteCacheTime
	p.integer = prmtc.inputInt[i]
}

func putParamWriteInlineMessageID(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteInlineMessageID
	p.str = prmtc.inputStr[i]
}

func putParamWriteErrors(prmtc prmTestContainer, prm formatter.IParameters, p *paramsT, i int) {
	p.testedFunc = prm.WriteErrors
	p.errors = prmtc.inputErr[i]
}

func (prmtc *prmTestContainer) writeDisableNotification() {
	prmtc.name = "(IParameters).WriteDisableNotification"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteDisableNotification
}

func (prmtc *prmTestContainer) writeEntities() {
	prmtc.name = "(IParameters).WriteEntities"
	prmtc.inputArr = [][]*types.MessageEntity{
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}},
		{},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}, nil, nil},
		{{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}, {{Type: "text_link", Offset: 0, Length: 7, Url: "https://youtube.com"}}}
	prmtc.isExpectedErr = []bool{false, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "5", "", "10"}
	prmtc.amount, prmtc.until = 5, 3
	prmtc.buildF = putParamWriteEntities
}

func (prmtc *prmTestContainer) writeLinkPreviewOptions() {
	prmtc.name = "(IParameters).WriteLinkPreviewOptions"
	prmtc.inputLink = []*types.LinkPreviewOptions{{IsDisabled: true, URL: "https://youtube.com", PreferLargeMedia: true, ShowAboveText: true},
		nil,
		{IsDisabled: true, URL: "https://youtube.com", PreferLargeMedia: true, ShowAboveText: true}, {IsDisabled: true, URL: "https://youtube.com", PreferLargeMedia: true, ShowAboveText: true}}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteLinkPreviewOptions
}

func (prmtc *prmTestContainer) writeMessageEffectID() {
	prmtc.name = "(IParameters).WriteMessageEffectID"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteMessageEffectID
}

func (prmtc *prmTestContainer) writeMessageThreadID() {
	prmtc.name = "(IParameters).WriteMessageThreadID"
	prmtc.inputInt = []int{12331, 0, 222, 12345}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteMessageThreadID
}

func (prmtc *prmTestContainer) writeMessageID() {
	prmtc.name = "(IParameters).WriteMessageID"
	prmtc.inputInt = []int{55, 0, 12356, 12345}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteMessageID
}

func (prmtc *prmTestContainer) writeMessageIDs() {
	prmtc.name = "(IParameters).WriteMessageIDs"
	prmtc.inputArrInt = [][]int{{55, 777, 12356, 12345},
		{},
		{55, 777, 0, 12345},
		{55, 777, 12356, 12345}, {55, 777, 12356, 12345}}
	prmtc.isExpectedErr = []bool{false, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "5", "", "10"}
	prmtc.amount, prmtc.until = 5, 3
	prmtc.buildF = putParamWriteMessageIDs
}

func (prmtc *prmTestContainer) writeCaption() {
	prmtc.name = "(IParameters).WriteCaption"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteCaption
}

func (prmtc *prmTestContainer) writeParseMode() {
	prmtc.name = "(IParameters).WriteParseMode"
	prmtc.inputStr = []string{types.HTML, types.Markdown, types.MarkdownV2, "", "something else", types.HTML, types.Markdown}
	prmtc.isExpectedErr = []bool{false, false, false, true, true, false, true}
	prmtc.codeErr = []string{"", "", "", "20", "20", "", "10"}
	prmtc.amount, prmtc.until = 7, 5
	prmtc.buildF = putParamWriteParseMode
}

func (prmtc *prmTestContainer) writeProtectContent() {
	prmtc.name = "(IParameters).WriteProtectContent"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteProtectContent
}

func (prmtc *prmTestContainer) writeString() {
	prmtc.name = "(IParameters).WriteString"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteString
}

func (prmtc *prmTestContainer) writeShowCaptionAboveMedia() {
	prmtc.name = "(IParameters).WriteShowCaptionAboveMedia"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteShowCaptionAboveMedia
}

func (prmtc *prmTestContainer) writeReplyParameters() {
	prmtc.name = "(IParameters).WriteReplyParameters"
	prmtc.inputReplyP = []*types.ReplyParameters{{MessageID: 999, Quote: "I am from the government and I am here to help"},
		nil,
		{MessageID: 999, Quote: "I am from the government and I am here to help"}, {MessageID: 999, Quote: "I am from the government and I am here to help"}}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteReplyParameters
}

func (prmtc *prmTestContainer) writeAllowPaidBroadcast() {
	prmtc.name = "(IParameters).WriteAllowPaidBroadcast"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteAllowPaidBroadcast
}

func (prmtc *prmTestContainer) writeStarCount() {
	prmtc.name = "(IParameters).WriteStarCount"
	prmtc.inputInt = []int{55, 0, 99, 45}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteStarCount
}

func (prmtc *prmTestContainer) writePayload() {
	prmtc.name = "(IParameters).WritePayload"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWritePayload
}

func (prmtc *prmTestContainer) writeEmoji() {
	prmtc.name = "(IParameters).WriteEmoji"
	prmtc.inputStr = []string{types.Emojis[1],
		"",
		strings.Repeat(types.Emojis[0], 6), strings.Repeat(types.Emojis[0], 7),
		strings.Repeat(types.Emojis[1], 6), strings.Repeat(types.Emojis[1], 7),
		strings.Repeat(types.Emojis[2], 5), strings.Repeat(types.Emojis[2], 6),
		strings.Repeat(types.Emojis[3], 5), strings.Repeat(types.Emojis[3], 6),
		strings.Repeat(types.Emojis[4], 6), strings.Repeat(types.Emojis[4], 7),
		strings.Repeat(types.Emojis[5], 64), strings.Repeat(types.Emojis[5], 65),
		fmt.Sprint(strings.Repeat(types.Emojis[0], 6), strings.Repeat(types.Emojis[1], 6), strings.Repeat(types.Emojis[2], 5),
			strings.Repeat(types.Emojis[3], 5), strings.Repeat(types.Emojis[4], 6), strings.Repeat(types.Emojis[5], 64)),
		"something",
		types.Emojis[1], types.Emojis[1]} //18
	prmtc.isExpectedErr = []bool{false,
		true,
		false, true,
		false, true,
		false, true,
		false, true,
		false, true,
		false, true,
		true,
		true,
		false, true} //18
	prmtc.codeErr = []string{"",
		"20",
		"", "20",
		"", "20",
		"", "20",
		"", "20",
		"", "20",
		"", "20",
		"20",
		"20",
		"", "10"} //18
	prmtc.amount, prmtc.until = 18, 16
	prmtc.buildF = putParamWriteEmoji
}

func (prmtc *prmTestContainer) writeAction() {
	prmtc.name = "(IParameters).WriteAction"
	prmtc.inputStr = []string{types.Actions[0], "", "a;lsd;lad;l",
		types.Actions[1], types.Actions[2], types.Actions[3],
		types.Actions[4], types.Actions[5], types.Actions[6],
		types.Actions[7], types.Actions[8], types.Actions[9], types.Actions[10],
		types.Actions[1], types.Actions[1]}
	prmtc.isExpectedErr = []bool{false, true, true,
		false, false, false,
		false, false, false,
		false, false, false,
		false, true}
	prmtc.codeErr = []string{"", "20", "20",
		"", "", "",
		"", "", "",
		"", "", "",
		"", "10"}
	prmtc.amount, prmtc.until = 14, 12
	prmtc.buildF = putParamWriteAction
}

func (prmtc *prmTestContainer) writeReaction() {
	prmtc.name = "(IParameters).WriteReaction"
	prmtc.inputReaction = [][]*types.ReactionType{{{ReactionTypeEmoji: nil}, {ReactionTypeEmoji: nil}},
		nil,
		{{ReactionTypeEmoji: nil}, nil, nil},
		{{ReactionTypeEmoji: nil}, {ReactionTypeEmoji: nil}}, {{ReactionTypeEmoji: nil}, {ReactionTypeEmoji: nil}}}
	prmtc.isExpectedErr = []bool{false, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "5", "", "10"}
	prmtc.amount, prmtc.until = 5, 3
	prmtc.buildF = putParamWriteReaction
}

func (prmtc *prmTestContainer) writeReactionIsBig() {
	prmtc.name = "(IParameters).WriteReactionIsBig"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteReactionIsBig
}

func (prmtc *prmTestContainer) writeOffset() {
	prmtc.name = "(IParameters).WriteOffset"
	prmtc.inputInt = []int{231, 0, 999, 1231}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteOffset
}

func (prmtc *prmTestContainer) writeLimit() {
	prmtc.name = "(IParameters).WriteLimit"
	prmtc.inputInt = []int{55, 0, -1, 101, 52, 58}
	prmtc.isExpectedErr = []bool{false, true, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "20", "20", "", "10"}
	prmtc.amount, prmtc.until = 6, 4
	prmtc.buildF = putParamWriteLimit
}

func (prmtc *prmTestContainer) writeEmojiStatusCustomEmojiID() {
	prmtc.name = "(IParameters).WriteEmojiStatusCustomEmojiID"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteEmojiStatusCustomEmojiID
}

func (prmtc *prmTestContainer) writeEmojiStatusExpirationDate() {
	prmtc.name = "(IParameters).WriteEmojiStatusExpirationDate"
	prmtc.inputInt = []int{231, 0, 999, 1231}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteEmojiStatusExpirationDate
}

func (prmtc *prmTestContainer) writeFileID() {
	prmtc.name = "(IParameters).WriteFileID"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteFileID
}

func (prmtc *prmTestContainer) writeUntilDate() {
	prmtc.name = "(IParameters).WriteUntilDate"
	prmtc.inputDates = []time.Duration{time.Hour, 0, time.Microsecond, time.Minute, time.Nanosecond, time.Second}
	prmtc.isExpectedErr = []bool{false, true, false, false, false, true}
	prmtc.codeErr = []string{"", "20", "", "", "", "10"}
	prmtc.amount, prmtc.until = 6, 4
	prmtc.buildF = putParamWriteUntilDate
}

func (prmtc *prmTestContainer) writeRevokeMessages() {
	prmtc.name = "(IParameters).WriteRevokeMessages"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteRevokeMessages
}

func (prmtc *prmTestContainer) writeOnlyIfBanned() {
	prmtc.name = "(IParameters).WriteOnlyIfBanned"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteOnlyIfBanned
}

func (prmtc *prmTestContainer) writePermissions() {
	prmtc.name = "(IParameters).WritePermissions"
	prmtc.inputPermis = []*types.ChatPermissions{{CanSendMessages: true},
		nil,
		{CanSendMessages: true}, {CanSendMessages: true}}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWritePermissions
}

func (prmtc *prmTestContainer) writeIndependentChatPermissions() {
	prmtc.name = "(IParameters).WriteIndependentChatPermissions"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteIndependentChatPermissions
}

func (prmtc *prmTestContainer) writeAdministratorRights() {
	prmtc.name = "(IParameters).WriteAdministratorRights"
	prmtc.inputAdmin = []*types.ChatAdministratorRights{{IsAnonymous: true},
		nil,
		{IsAnonymous: true}, {IsAnonymous: true}}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteAdministratorRights
}

func (prmtc *prmTestContainer) writeCustomTitle() {
	prmtc.name = "(IParameters).WriteCustomTitle"
	prmtc.inputStr = []string{"a;lsdl;asd", "", "1", strings.Repeat("c", 16), strings.Repeat("c", 17), strings.Repeat("c", 2), strings.Repeat("c", 3)}
	prmtc.isExpectedErr = []bool{false, true, false, false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "", "20", "", "10"}
	prmtc.amount, prmtc.until = 7, 5
	prmtc.buildF = putParamWriteCustomTitle
}

func (prmtc *prmTestContainer) writeUserID() {
	prmtc.name = "(IParameters).WriteUserID"
	prmtc.inputInt = []int{231, 0, 999, 1231}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteUserID
}

func (prmtc *prmTestContainer) writeCallBackQueryID() {
	prmtc.name = "(IParameters).WriteCallBackQueryID"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteCallBackQueryID
}

func (prmtc *prmTestContainer) writeShowAlert() {
	prmtc.name = "(IParameters).WriteShowAlert"
	prmtc.isExpectedErr = []bool{false, false, true}
	prmtc.codeErr = []string{"", "", "10"}
	prmtc.amount, prmtc.until = 3, 1
	prmtc.buildF = putParamWriteShowAlert
}

func (prmtc *prmTestContainer) writeURL() {
	prmtc.name = "(IParameters).WriteURL"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteURL
}

func (prmtc *prmTestContainer) writeCacheTime() {
	prmtc.name = "(IParameters).WriteCacheTime"
	prmtc.inputInt = []int{231, 0, 999, 1231}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteCacheTime
}

func (prmtc *prmTestContainer) writeInlineMessageID() {
	prmtc.name = "(IParameters).WriteInlineMessageID"
	prmtc.inputStr = []string{"kl;asdok-", "", "0-1230-1", "23828913819"}
	prmtc.isExpectedErr = []bool{false, true, false, true}
	prmtc.codeErr = []string{"", "20", "", "10"}
	prmtc.amount, prmtc.until = 4, 2
	prmtc.buildF = putParamWriteInlineMessageID
}

func (prmtc *prmTestContainer) writeErrors() {
	prmtc.name = "(IParameters).WriteErrors()"
	prmtc.inputErr = [][]*types.PassportElementError{{{}, {}}, nil, {{}, {}, nil}, {{}, {}}, {{}, {}}}
	prmtc.isExpectedErr = []bool{false, true, true, false, true}
	prmtc.codeErr = []string{"", "20", "5", "", "10"}
	prmtc.amount, prmtc.until = 5, 3
	prmtc.buildF = putParamWriteErrors
}

func (prm *paramsT) startTest(part string, i int, t *testing.T) {
	switch f := prm.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, prm.name, prm.codeErr, prm.str, prm.isExpectedErr, i)
		checkError(f(prm.str), prm.isExpectedErr, prm.codeErr, t)
	case func(int) error:
		printTestLog(part, prm.name, prm.codeErr, prm.integer, prm.isExpectedErr, i)
		checkError(f(prm.integer), prm.isExpectedErr, prm.codeErr, t)
	case func(time.Duration) error:
		printTestLog(part, prm.name, prm.codeErr, prm.integer, prm.isExpectedErr, i)
		checkError(f(prm.date), prm.isExpectedErr, prm.codeErr, t)
	case func([]*types.MessageEntity) error:
		printTestLog(part, prm.name, prm.codeErr, prm.array, prm.isExpectedErr, i)
		checkError(f(prm.array), prm.isExpectedErr, prm.codeErr, t)
	case func([]int) error:
		printTestLog(part, prm.name, prm.codeErr, prm.arrayInt, prm.isExpectedErr, i)
		checkError(f(prm.arrayInt), prm.isExpectedErr, prm.codeErr, t)
	case func() error:
		printTestLog(part, prm.name, prm.codeErr, true, prm.isExpectedErr, i)
		checkError(f(), prm.isExpectedErr, prm.codeErr, t)
	case func(*types.LinkPreviewOptions) error:
		printTestLog(part, prm.name, prm.codeErr, prm.link, prm.isExpectedErr, i)
		checkError(f(prm.link), prm.isExpectedErr, prm.codeErr, t)
	case func(*types.ReplyParameters) error:
		printTestLog(part, prm.name, prm.codeErr, prm.replyP, prm.isExpectedErr, i)
		checkError(f(prm.replyP), prm.isExpectedErr, prm.codeErr, t)
	case func([]*types.ReactionType) error:
		printTestLog(part, prm.name, prm.codeErr, prm.reaction, prm.isExpectedErr, i)
		checkError(f(prm.reaction), prm.isExpectedErr, prm.codeErr, t)
	case func(*types.ChatPermissions) error:
		printTestLog(part, prm.name, prm.codeErr, prm.permis, prm.isExpectedErr, i)
		checkError(f(prm.permis), prm.isExpectedErr, prm.codeErr, t)
	case func(*types.ChatAdministratorRights) error:
		printTestLog(part, prm.name, prm.codeErr, prm.adminrights, prm.isExpectedErr, i)
		checkError(f(prm.adminrights), prm.isExpectedErr, prm.codeErr, t)
	case func([]*types.PassportElementError) error:
		printTestLog(part, prm.name, prm.codeErr, prm.errors, prm.isExpectedErr, i)
		checkError(f(prm.errors), prm.isExpectedErr, prm.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (prmtc *prmTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var prm formatter.IParameters
	a, b := make([]UnitTest, prmtc.until), make([]UnitTest, prmtc.amount-prmtc.until)
	for i, j := 0, 0; i < prmtc.amount; i++ {
		p := prmtc.defaultParamT(i)
		if i < prmtc.until {
			prmtc.buildF(*prmtc, msg.NewParameters(), p, i)
			a[i] = p
		} else {
			if j%2 == 0 {
				prm = msg.NewParameters()
			}
			prmtc.buildF(*prmtc, prm, p, i)
			b[j] = p
			j++
		}
	}
	return a, b
}

func mainParametersLogic(msg *formatter.Message, prmtc prmTestContainer, t *testing.T) {
	prmcontainer, doublecontainer := prmtc.createTestArrays(msg)
	check(prmcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestParamWriteDisableNotification(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeDisableNotification()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteEntities(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeEntities()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteLinkPreviewOptions(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeLinkPreviewOptions()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageEffectID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageEffectID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageThreadID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageThreadID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteMessageIDs(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeMessageIDs()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteCaption(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeCaption()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteParseMode(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeParseMode()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteProtectContent(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeProtectContent()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteString(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeString()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteShowCaptionAboveMedia(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeShowCaptionAboveMedia()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteReplyParameters(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeReplyParameters()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteAllowPaidBroadcast(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeAllowPaidBroadcast()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteStarCount(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeStarCount()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWritePayload(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writePayload()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteEmoji(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeEmoji()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteAction(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeAction()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteReaction(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeReaction()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteReactionIsBig(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeReactionIsBig()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteOffset(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeOffset()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteLimit(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeLimit()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteEmojiStatusCustomEmojiID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeEmojiStatusCustomEmojiID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteEmojiStatusExpirationDate(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeEmojiStatusExpirationDate()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteFileID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeFileID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteUntilDate(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeUntilDate()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteRevokeMessages(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeRevokeMessages()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteOnlyIfBanned(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeOnlyIfBanned()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWritePermissions(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writePermissions()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteIndependentChatPermissions(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeIndependentChatPermissions()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteAdministratorRights(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeAdministratorRights()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteCustomTitle(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeCustomTitle()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteUserID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeUserID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteCallBackQueryID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeCallBackQueryID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteShowAlert(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeShowAlert()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteURL(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeURL()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteCacheTime(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeCacheTime()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteInlineMessageID(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeInlineMessageID()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}

func TestParamWriteErrors(t *testing.T) {
	prmtc := new(prmTestContainer)
	prmtc.writeErrors()
	msg := formatter.CreateEmpltyMessage()
	mainParametersLogic(msg, *prmtc, t)
}
