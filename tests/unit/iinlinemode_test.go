package unit

// import (
// 	"testing"

// 	"github.com/iamissahar/Fmtogram/formatter"
// 	"github.com/iamissahar/Fmtogram/types"
// )

// type inlineModeT struct {
// 	name          string
// 	str           string
// 	integer       int
// 	but           *types.InlineQueryResultsButton
// 	cachAudio     *types.InlineQueryResultCachedAudio
// 	cachDocument  *types.InlineQueryResultCachedDocument
// 	cachGif       *types.InlineQueryResultCachedGif
// 	cachMpeg4Gif  *types.InlineQueryResultCachedMpeg4Gif
// 	cachPhoto     *types.InlineQueryResultCachedPhoto
// 	cachSticker   *types.InlineQueryResultCachedSticker
// 	cachVideo     *types.InlineQueryResultCachedVideo
// 	cachVoice     *types.InlineQueryResultCachedVoice
// 	art           *types.InlineQueryResultArticle
// 	ad            *types.InlineQueryResultAudio
// 	con           *types.InlineQueryResultContact
// 	g             *types.InlineQueryResultGame
// 	doc           *types.InlineQueryResultDocument
// 	gf            *types.InlineQueryResultGif
// 	loc           *types.InlineQueryResultLocation
// 	mpg4gf        *types.InlineQueryResultMpeg4Gif
// 	ph            *types.InlineQueryResultPhoto
// 	ven           *types.InlineQueryResultVenue
// 	vd            *types.InlineQueryResultVideo
// 	vc            *types.InlineQueryResultVoice
// 	testedFunc    interface{}
// 	isExpectedErr bool
// 	codeErr       string
// }

// type inmTestContainer struct {
// 	name                string
// 	inputStr            []string
// 	inputInt            []int
// 	inputButton         []*types.InlineQueryResultsButton
// 	InputCachedAudio    []*types.InlineQueryResultCachedAudio
// 	InputCachedDocument []*types.InlineQueryResultCachedDocument
// 	InputCachedGif      []*types.InlineQueryResultCachedGif
// 	InputCachedMpeg4Gif []*types.InlineQueryResultCachedMpeg4Gif
// 	InputCachedPhoto    []*types.InlineQueryResultCachedPhoto
// 	InputCachedSticker  []*types.InlineQueryResultCachedSticker
// 	InputCachedVideo    []*types.InlineQueryResultCachedVideo
// 	InputCachedVoice    []*types.InlineQueryResultCachedVoice
// 	InputArticle        []*types.InlineQueryResultArticle
// 	InputAudio          []*types.InlineQueryResultAudio
// 	InputContact        []*types.InlineQueryResultContact
// 	InputGame           []*types.InlineQueryResultGame
// 	InputDocument       []*types.InlineQueryResultDocument
// 	InputGif            []*types.InlineQueryResultGif
// 	InputLocation       []*types.InlineQueryResultLocation
// 	InputMpeg4Gif       []*types.InlineQueryResultMpeg4Gif
// 	InputPhoto          []*types.InlineQueryResultPhoto
// 	InputVenue          []*types.InlineQueryResultVenue
// 	InputVideo          []*types.InlineQueryResultVideo
// 	InputVoice          []*types.InlineQueryResultVoice
// 	isExpectedErr       []bool
// 	codeErr             []string
// 	amount, until       int
// 	buildF              func(inmode formatter.IInlineMode, in *inlineModeT, i int)
// 	nodouble            bool
// }

// func (inmtc *inmTestContainer) byDefault(i int) *inlineModeT {
// 	return &inlineModeT{name: inmtc.name, isExpectedErr: inmtc.isExpectedErr[i], codeErr: inmtc.codeErr[i]}
// }

// func (inmtc *inmTestContainer) queryID(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteQueryID
// 	in.str = inmtc.inputStr[i]
// }

// func (inmtc *inmTestContainer) webAppQueryID(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteWebAppQueryID
// 	in.str = inmtc.inputStr[i]
// }

// func (inmtc *inmTestContainer) cachedAudio(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedAudio
// 	r.cachAudio = inmtc.InputCachedAudio[i]
// }

// func (inmtc *inmTestContainer) cachedDocument(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedDocument
// 	r.cachDocument = inmtc.InputCachedDocument[i]
// }

// func (inmtc *inmTestContainer) cachedGif(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedGif
// 	r.cachGif = inmtc.InputCachedGif[i]
// }

// func (inmtc *inmTestContainer) cachedMpeg4Gif(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedMpeg4Gif
// 	r.cachMpeg4Gif = inmtc.InputCachedMpeg4Gif[i]
// }

// func (inmtc *inmTestContainer) cachedPhoto(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedPhoto
// 	r.cachPhoto = inmtc.InputCachedPhoto[i]
// }

// func (inmtc *inmTestContainer) cachedSticker(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedSticker
// 	r.cachSticker = inmtc.InputCachedSticker[i]
// }

// func (inmtc *inmTestContainer) cachedVideo(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedVideo
// 	r.cachVideo = inmtc.InputCachedVideo[i]
// }

// func (inmtc *inmTestContainer) cachedVoice(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteCachedVoice
// 	r.cachVoice = inmtc.InputCachedVoice[i]
// }

// func (inmtc *inmTestContainer) article(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteArticle
// 	r.art = inmtc.InputArticle[i]
// }

// func (inmtc *inmTestContainer) audio(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteAudio
// 	r.ad = inmtc.InputAudio[i]
// }

// func (inmtc *inmTestContainer) contact(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteContact
// 	r.con = inmtc.InputContact[i]
// }

// func (inmtc *inmTestContainer) game(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteGame
// 	r.g = inmtc.InputGame[i]
// }

// func (inmtc *inmTestContainer) document(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteDocument
// 	r.doc = inmtc.InputDocument[i]
// }

// func (inmtc *inmTestContainer) gif(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteGif
// 	r.gf = inmtc.InputGif[i]
// }

// func (inmtc *inmTestContainer) location(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteLocation
// 	r.loc = inmtc.InputLocation[i]
// }

// func (inmtc *inmTestContainer) mpeg4Gif(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteMpeg4Gif
// 	r.mpg4gf = inmtc.InputMpeg4Gif[i]
// }

// func (inmtc *inmTestContainer) photo(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WritePhoto
// 	r.ph = inmtc.InputPhoto[i]
// }

// func (inmtc *inmTestContainer) venue(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteVenue
// 	r.ven = inmtc.InputVenue[i]
// }

// func (inmtc *inmTestContainer) video(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteVideo
// 	r.vd = inmtc.InputVideo[i]
// }

// func (inmtc *inmTestContainer) voice(inmode formatter.IInlineMode, r *inlineModeT, i int) {
// 	r.testedFunc = inmode.WriteVoice
// 	r.vc = inmtc.InputVoice[i]
// }

// func (inmtc *inmTestContainer) cacheTime(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteCacheTime
// 	in.integer = inmtc.inputInt[i]
// }

// func (inmtc *inmTestContainer) isPersonal(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteIsPersonal
// }

// func (inmtc *inmTestContainer) nextOffset(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteNextOffset
// 	in.str = inmtc.inputStr[i]
// }

// func (inmtc *inmTestContainer) button(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteButton
// 	in.but = inmtc.inputButton[i]
// }

// func (inmtc *inmTestContainer) allowUsersChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteAllowUserChats
// }

// func (inmtc *inmTestContainer) allowBotChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteAllowUserChats
// }

// func (inmtc *inmTestContainer) allowGroupChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteAllowUserChats
// }

// func (inmtc *inmTestContainer) allowChannelChats(inmode formatter.IInlineMode, in *inlineModeT, i int) {
// 	in.testedFunc = inmode.WriteAllowUserChats
// }

// func (inmtc *inmTestContainer) writeQueryID() {
// 	inmtc.name = "(IInlineMode).WriteQueryID()"
// 	inmtc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
// 	inmtc.isExpectedErr = []bool{false, true, false, true}
// 	inmtc.codeErr = []string{"", "20", "", "10"}
// 	inmtc.amount, inmtc.until = 4, 2
// 	inmtc.buildF = inmtc.queryID
// }

// func (inmtc *inmTestContainer) writeWebAppQueryID() {
// 	inmtc.name = "(IInlineMode).WriteWebAppQueryID()"
// 	inmtc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
// 	inmtc.isExpectedErr = []bool{false, true, false, true}
// 	inmtc.codeErr = []string{"", "20", "", "10"}
// 	inmtc.amount, inmtc.until = 4, 2
// 	inmtc.buildF = inmtc.webAppQueryID
// }

// func (inmtc *inmTestContainer) writeCachedAudio() {
// 	inmtc.name = "(IInlineMode).WriteCachedAudio()"
// 	inmtc.InputCachedAudio = []*types.InlineQueryResultCachedAudio{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedAudio
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedDocument() {
// 	inmtc.name = "(IInlineMode).WriteCachedDocument()"
// 	inmtc.InputCachedDocument = []*types.InlineQueryResultCachedDocument{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedDocument
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedGif() {
// 	inmtc.name = "(IInlineMode).WriteCachedGif()"
// 	inmtc.InputCachedGif = []*types.InlineQueryResultCachedGif{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedGif
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedMpeg4Gif() {
// 	inmtc.name = "(IInlineMode).WriteCachedMpeg4Gif()"
// 	inmtc.InputCachedMpeg4Gif = []*types.InlineQueryResultCachedMpeg4Gif{{}, nil, {}, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false, true}
// 	inmtc.codeErr = []string{"", "20", "", "10"}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedMpeg4Gif
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedPhoto() {
// 	inmtc.name = "(IInlineMode).WriteCachedPhoto()"
// 	inmtc.InputCachedPhoto = []*types.InlineQueryResultCachedPhoto{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedPhoto
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedSticker() {
// 	inmtc.name = "(IInlineMode).WriteCachedSticker()"
// 	inmtc.InputCachedSticker = []*types.InlineQueryResultCachedSticker{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedSticker
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedVideo() {
// 	inmtc.name = "(IInlineMode).WriteCachedVideo()"
// 	inmtc.InputCachedVideo = []*types.InlineQueryResultCachedVideo{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedVideo
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCachedVoice() {
// 	inmtc.name = "(IInlineMode).WriteCachedVoice()"
// 	inmtc.InputCachedVoice = []*types.InlineQueryResultCachedVoice{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.cachedVoice
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeArticle() {
// 	inmtc.name = "(IInlineMode).WriteArticle()"
// 	inmtc.InputArticle = []*types.InlineQueryResultArticle{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.article
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeAudio() {
// 	inmtc.name = "(IInlineMode).WriteAudio()"
// 	inmtc.InputAudio = []*types.InlineQueryResultAudio{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.audio
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeContact() {
// 	inmtc.name = "(IInlineMode).WriteContact()"
// 	inmtc.InputContact = []*types.InlineQueryResultContact{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.contact
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeGame() {
// 	inmtc.name = "(IInlineMode).WriteGame()"
// 	inmtc.InputGame = []*types.InlineQueryResultGame{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.game
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeDocument() {
// 	inmtc.name = "(IInlineMode).WriteDocument()"
// 	inmtc.InputDocument = []*types.InlineQueryResultDocument{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.document
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeGif() {
// 	inmtc.name = "(IInlineMode).WriteGif()"
// 	inmtc.InputGif = []*types.InlineQueryResultGif{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false, true}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.gif
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeLocation() {
// 	inmtc.name = "(IInlineMode).WriteLocation()"
// 	inmtc.InputLocation = []*types.InlineQueryResultLocation{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.location
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeMpeg4Gif() {
// 	inmtc.name = "(IInlineMode).WriteMpeg4Gif()"
// 	inmtc.InputMpeg4Gif = []*types.InlineQueryResultMpeg4Gif{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.mpeg4Gif
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writePhoto() {
// 	inmtc.name = "(IInlineMode).WritePhoto()"
// 	inmtc.InputPhoto = []*types.InlineQueryResultPhoto{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.photo
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeVenue() {
// 	inmtc.name = "(IInlineMode).WriteVenue()"
// 	inmtc.InputVenue = []*types.InlineQueryResultVenue{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.venue
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeVideo() {
// 	inmtc.name = "(IInlineMode).WriteVideo()"
// 	inmtc.InputVideo = []*types.InlineQueryResultVideo{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.video
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeVoice() {
// 	inmtc.name = "(IInlineMode).WriteVoice()"
// 	inmtc.InputVoice = []*types.InlineQueryResultVoice{{}, nil, {}}
// 	inmtc.isExpectedErr = []bool{false, true, false}
// 	inmtc.codeErr = []string{"", "20", ""}
// 	inmtc.amount, inmtc.until = 3, 3
// 	inmtc.buildF = inmtc.voice
// 	inmtc.nodouble = true
// }

// func (inmtc *inmTestContainer) writeCacheTime() {
// 	inmtc.name = "(IInlineMode).WriteCacheTime()"
// 	inmtc.inputInt = []int{0, -12124, -1, 1231, 34341234, 645, 1}
// 	inmtc.isExpectedErr = []bool{true, true, true, false, false, false, true}
// 	inmtc.codeErr = []string{"20", "20", "20", "", "", "", "10"}
// 	inmtc.amount, inmtc.until = 7, 5
// 	inmtc.buildF = inmtc.cacheTime
// }

// func (inmtc *inmTestContainer) writeIsPersonal() {
// 	inmtc.name = "(IInlineMode).WriteIsPersonal()"
// 	inmtc.isExpectedErr = []bool{false, false, true}
// 	inmtc.codeErr = []string{"", "", "10"}
// 	inmtc.amount, inmtc.until = 3, 1
// 	inmtc.buildF = inmtc.isPersonal
// }

// func (inmtc *inmTestContainer) writeNextOffset() {
// 	inmtc.name = "(IInlineMode).WriteNextOffset()"
// 	inmtc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
// 	inmtc.isExpectedErr = []bool{false, false, false, true}
// 	inmtc.codeErr = []string{"", "", "", "10"}
// 	inmtc.amount, inmtc.until = 4, 2
// 	inmtc.buildF = inmtc.nextOffset
// }

// func (inmtc *inmTestContainer) writeButton() {
// 	inmtc.name = "(IInlineMode).WriteButton()"
// 	inmtc.inputButton = []*types.InlineQueryResultsButton{{Text: "nothing", WebApp: nil, StartParameter: "?"},
// 		{Text: "something"}, nil,
// 		{Text: "nothing", WebApp: nil}, {Text: "absolutely nothing", StartParameter: ""},
// 		{Text: "nothing", WebApp: nil, StartParameter: "?"}, {Text: "nothing", WebApp: nil, StartParameter: "?"}}
// 	inmtc.isExpectedErr = []bool{false, true, true, true, true, false, true}
// 	inmtc.codeErr = []string{"", "20", "20", "20", "20", "", "10"}
// 	inmtc.amount, inmtc.until = 7, 5
// 	inmtc.buildF = inmtc.button
// }

// func (inmtc *inmTestContainer) writeAllowUserChats() {
// 	inmtc.name = "(IInlineMode).WriteAllowUserChats()"
// 	inmtc.isExpectedErr = []bool{false, false, true}
// 	inmtc.codeErr = []string{"", "", "10"}
// 	inmtc.amount, inmtc.until = 3, 1
// 	inmtc.buildF = inmtc.allowUsersChats
// }

// func (inmtc *inmTestContainer) writeAllowBotChats() {
// 	inmtc.name = "(IInlineMode).WriteAllowBotChats()"
// 	inmtc.isExpectedErr = []bool{false, false, true}
// 	inmtc.codeErr = []string{"", "", "10"}
// 	inmtc.amount, inmtc.until = 3, 1
// 	inmtc.buildF = inmtc.allowBotChats
// }

// func (inmtc *inmTestContainer) writeAllowGroupChats() {
// 	inmtc.name = "(IInlineMode).WriteAllowGroupChats()"
// 	inmtc.isExpectedErr = []bool{false, false, true}
// 	inmtc.codeErr = []string{"", "", "10"}
// 	inmtc.amount, inmtc.until = 3, 1
// 	inmtc.buildF = inmtc.allowGroupChats
// }

// func (inmtc *inmTestContainer) writeAllowChannelChats() {
// 	inmtc.name = "(IInlineMode).WriteAllowChannelChats()"
// 	inmtc.isExpectedErr = []bool{false, false, true}
// 	inmtc.codeErr = []string{"", "", "10"}
// 	inmtc.amount, inmtc.until = 3, 1
// 	inmtc.buildF = inmtc.allowChannelChats
// }

// func (inm *inlineModeT) startTest(part string, i int, t *testing.T) {
// 	switch f := inm.testedFunc.(type) {
// 	case func(string) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.str, inm.isExpectedErr, i, t)
// 		checkError(f(inm.str), inm.isExpectedErr, inm.codeErr, t)
// 	case func(int) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.integer, inm.isExpectedErr, i, t)
// 		checkError(f(inm.integer), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultsButton) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.but, inm.isExpectedErr, i, t)
// 		checkError(f(inm.but), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedAudio) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachAudio, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachAudio), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedDocument) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachDocument, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachDocument), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedGif) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachGif, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachGif), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedMpeg4Gif) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachMpeg4Gif, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachMpeg4Gif), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedPhoto) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachPhoto, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachPhoto), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedSticker) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachSticker, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachSticker), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedVideo) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachVideo, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachVideo), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultCachedVoice) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.cachVoice, inm.isExpectedErr, i, t)
// 		checkError(f(inm.cachVoice), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultArticle) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.art, inm.isExpectedErr, i, t)
// 		checkError(f(inm.art), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultAudio) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.ad, inm.isExpectedErr, i, t)
// 		checkError(f(inm.ad), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultContact) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.con, inm.isExpectedErr, i, t)
// 		checkError(f(inm.con), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultGame) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.g, inm.isExpectedErr, i, t)
// 		checkError(f(inm.g), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultDocument) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.doc, inm.isExpectedErr, i, t)
// 		checkError(f(inm.doc), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultGif) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.gf, inm.isExpectedErr, i, t)
// 		checkError(f(inm.gf), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultLocation) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.loc, inm.isExpectedErr, i, t)
// 		checkError(f(inm.loc), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultMpeg4Gif) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.mpg4gf, inm.isExpectedErr, i, t)
// 		checkError(f(inm.mpg4gf), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultPhoto) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.ph, inm.isExpectedErr, i, t)
// 		checkError(f(inm.ph), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultVenue) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.ven, inm.isExpectedErr, i, t)
// 		checkError(f(inm.ven), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultVideo) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.vd, inm.isExpectedErr, i, t)
// 		checkError(f(inm.vd), inm.isExpectedErr, inm.codeErr, t)
// 	case func(*types.InlineQueryResultVoice) error:
// 		printTestLog(part, inm.name, inm.codeErr, inm.vc, inm.isExpectedErr, i, t)
// 		checkError(f(inm.vc), inm.isExpectedErr, inm.codeErr, t)
// 	case func() error:
// 		printTestLog(part, inm.name, inm.codeErr, true, inm.isExpectedErr, i, t)
// 		checkError(f(), inm.isExpectedErr, inm.codeErr, t)
// 	default:
// 		t.Fatal("unexpected type of tested function")
// 	}
// }

// func (inmtc *inmTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
// 	var inm formatter.IInlineMode
// 	a, b := make([]UnitTest, inmtc.until), make([]UnitTest, inmtc.amount-inmtc.until)
// 	for i, j := 0, 0; i < inmtc.amount; i++ {
// 		s := inmtc.byDefault(i)
// 		if i < inmtc.until {
// 			inmtc.buildF(msg.NewInlineMode(), s, i)
// 			a[i] = s
// 		} else {
// 			if j%2 == 0 {
// 				inm = msg.NewInlineMode()
// 			}
// 			inmtc.buildF(inm, s, i)
// 			b[j] = s
// 			j++
// 		}
// 	}
// 	return a, b
// }

// func (inmtc *inmTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
// 	inmcontainer, doublecontainer := inmtc.createTestArrays(msg)
// 	check(inmcontainer, t)
// 	if !inmtc.nodouble {
// 		doubleCheck(doublecontainer, t)
// 	}
// }

// func TestInlineModeWriteQueryID(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeQueryID()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteWebAppQueryID(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeWebAppQueryID()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedAudio(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedAudio()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedDocument(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedDocument()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedGif(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedGif()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedMpeg4Gif(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedMpeg4Gif()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedPhoto(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedPhoto()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedSticker(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedSticker()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedVideo(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedVideo()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCachedVoice(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCachedVoice()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteArticle(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeArticle()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteAudio(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeAudio()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteContact(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeContact()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteGame(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeGame()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteDocument(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeDocument()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteGif(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeGif()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteLocation(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeLocation()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteMpeg4Gif(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeMpeg4Gif()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWritePhoto(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writePhoto()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteVenue(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeVenue()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteVideo(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeVideo()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteVoice(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeVoice()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteCacheTime(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeCacheTime()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteIsPersonal(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeIsPersonal()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteNextOffset(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeNextOffset()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteButton(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeButton()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteAllowUserChats(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeAllowUserChats()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteAllowBotChats(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeAllowBotChats()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteAllowGroupChats(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeAllowGroupChats()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }

// func TestInlineModeWriteAllowChannelChats(t *testing.T) {
// 	t.Parallel()
// 	inmtc := new(inmTestContainer)
// 	inmtc.writeAllowChannelChats()
// 	msg := formatter.CreateEmpltyMessage()
// 	inmtc.mainLogic(msg, t)
// }
