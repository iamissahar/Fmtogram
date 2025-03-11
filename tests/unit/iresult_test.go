package unit

// type resultT struct {
// 	name          string
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

// type restTestContainer struct {
// 	name                string
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
// 	buildF              func(res formatter.IResult, r *resultT, i int)
// }

// func (rtc *restTestContainer) byDefault(i int) *resultT {
// 	return &resultT{name: rtc.name, isExpectedErr: rtc.isExpectedErr[i], codeErr: rtc.codeErr[i]}
// }

// func (rtc *restTestContainer) cachedAudio(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedAudio
// 	r.cachAudio = rtc.InputCachedAudio[i]
// }

// func (rtc *restTestContainer) cachedDocument(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedDocument
// 	r.cachDocument = rtc.InputCachedDocument[i]
// }

// func (rtc *restTestContainer) cachedGif(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedGif
// 	r.cachGif = rtc.InputCachedGif[i]
// }

// func (rtc *restTestContainer) cachedMpeg4Gif(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedMpeg4Gif
// 	r.cachMpeg4Gif = rtc.InputCachedMpeg4Gif[i]
// }

// func (rtc *restTestContainer) cachedPhoto(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedPhoto
// 	r.cachPhoto = rtc.InputCachedPhoto[i]
// }

// func (rtc *restTestContainer) cachedSticker(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedSticker
// 	r.cachSticker = rtc.InputCachedSticker[i]
// }

// func (rtc *restTestContainer) cachedVideo(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedVideo
// 	r.cachVideo = rtc.InputCachedVideo[i]
// }

// func (rtc *restTestContainer) cachedVoice(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteCachedVoice
// 	r.cachVoice = rtc.InputCachedVoice[i]
// }

// func (rtc *restTestContainer) article(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteArticle
// 	r.art = rtc.InputArticle[i]
// }

// func (rtc *restTestContainer) audio(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteAudio
// 	r.ad = rtc.InputAudio[i]
// }

// func (rtc *restTestContainer) contact(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteContact
// 	r.con = rtc.InputContact[i]
// }

// func (rtc *restTestContainer) game(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteGame
// 	r.g = rtc.InputGame[i]
// }

// func (rtc *restTestContainer) document(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteDocument
// 	r.doc = rtc.InputDocument[i]
// }

// func (rtc *restTestContainer) gif(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteGif
// 	r.gf = rtc.InputGif[i]
// }

// func (rtc *restTestContainer) location(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteLocation
// 	r.loc = rtc.InputLocation[i]
// }

// func (rtc *restTestContainer) mpeg4Gif(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteMpeg4Gif
// 	r.mpg4gf = rtc.InputMpeg4Gif[i]
// }

// func (rtc *restTestContainer) photo(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WritePhoto
// 	r.ph = rtc.InputPhoto[i]
// }

// func (rtc *restTestContainer) venue(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteVenue
// 	r.ven = rtc.InputVenue[i]
// }

// func (rtc *restTestContainer) video(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteVideo
// 	r.vd = rtc.InputVideo[i]
// }

// func (rtc *restTestContainer) voice(res formatter.IResult, r *resultT, i int) {
// 	r.testedFunc = res.WriteVoice
// 	r.vc = rtc.InputVoice[i]
// }

// func (rtc *restTestContainer) writeCachedAudio() {
// 	rtc.name = "(IResult).WriteCachedAudio()"
// 	rtc.InputCachedAudio = []*types.InlineQueryResultCachedAudio{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedAudio
// }

// func (rtc *restTestContainer) writeCachedDocument() {
// 	rtc.name = "(IResult).WriteCachedDocument()"
// 	rtc.InputCachedDocument = []*types.InlineQueryResultCachedDocument{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedDocument
// }

// func (rtc *restTestContainer) writeCachedGif() {
// 	rtc.name = "(IResult).WriteCachedGif()"
// 	rtc.InputCachedGif = []*types.InlineQueryResultCachedGif{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedGif
// }

// func (rtc *restTestContainer) writeCachedMpeg4Gif() {
// 	rtc.name = "(IResult).WriteCachedMpeg4Gif()"
// 	rtc.InputCachedMpeg4Gif = []*types.InlineQueryResultCachedMpeg4Gif{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedMpeg4Gif
// }

// func (rtc *restTestContainer) writeCachedPhoto() {
// 	rtc.name = "(IResult).WriteCachedPhoto()"
// 	rtc.InputCachedPhoto = []*types.InlineQueryResultCachedPhoto{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedPhoto
// }

// func (rtc *restTestContainer) writeCachedSticker() {
// 	rtc.name = "(IResult).WriteCachedSticker()"
// 	rtc.InputCachedSticker = []*types.InlineQueryResultCachedSticker{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedSticker
// }

// func (rtc *restTestContainer) writeCachedVideo() {
// 	rtc.name = "(IResult).WriteCachedVideo()"
// 	rtc.InputCachedVideo = []*types.InlineQueryResultCachedVideo{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedVideo
// }

// func (rtc *restTestContainer) writeCachedVoice() {
// 	rtc.name = "(IResult).WriteCachedVoice()"
// 	rtc.InputCachedVoice = []*types.InlineQueryResultCachedVoice{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.cachedVoice
// }

// func (rtc *restTestContainer) writeArticle() {
// 	rtc.name = "(IResult).WriteArticle()"
// 	rtc.InputArticle = []*types.InlineQueryResultArticle{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.article
// }

// func (rtc *restTestContainer) writeAudio() {
// 	rtc.name = "(IResult).WriteAudio()"
// 	rtc.InputAudio = []*types.InlineQueryResultAudio{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.audio
// }

// func (rtc *restTestContainer) writeContact() {
// 	rtc.name = "(IResult).WriteContact()"
// 	rtc.InputContact = []*types.InlineQueryResultContact{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.contact
// }

// func (rtc *restTestContainer) writeGame() {
// 	rtc.name = "(IResult).WriteGame()"
// 	rtc.InputGame = []*types.InlineQueryResultGame{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.game
// }

// func (rtc *restTestContainer) writeDocument() {
// 	rtc.name = "(IResult).WriteDocument()"
// 	rtc.InputDocument = []*types.InlineQueryResultDocument{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.document
// }

// func (rtc *restTestContainer) writeGif() {
// 	rtc.name = "(IResult).WriteGif()"
// 	rtc.InputGif = []*types.InlineQueryResultGif{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.gif
// }

// func (rtc *restTestContainer) writeLocation() {
// 	rtc.name = "(IResult).WriteLocation()"
// 	rtc.InputLocation = []*types.InlineQueryResultLocation{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.location
// }

// func (rtc *restTestContainer) writeMpeg4Gif() {
// 	rtc.name = "(IResult).WriteMpeg4Gif()"
// 	rtc.InputMpeg4Gif = []*types.InlineQueryResultMpeg4Gif{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.mpeg4Gif
// }

// func (rtc *restTestContainer) writePhoto() {
// 	rtc.name = "(IResult).WritePhoto()"
// 	rtc.InputPhoto = []*types.InlineQueryResultPhoto{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.photo
// }

// func (rtc *restTestContainer) writeVenue() {
// 	rtc.name = "(IResult).WriteVenue()"
// 	rtc.InputVenue = []*types.InlineQueryResultVenue{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.venue
// }

// func (rtc *restTestContainer) writeVideo() {
// 	rtc.name = "(IResult).WriteVideo()"
// 	rtc.InputVideo = []*types.InlineQueryResultVideo{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.video
// }

// func (rtc *restTestContainer) writeVoice() {
// 	rtc.name = "(IResult).WriteVoice()"
// 	rtc.InputVoice = []*types.InlineQueryResultVoice{{}, nil, {}, {}}
// 	rtc.isExpectedErr = []bool{false, true, false, true}
// 	rtc.codeErr = []string{"", "20", "", "10"}
// 	rtc.amount, rtc.until = 4, 2
// 	rtc.buildF = rtc.voice
// }

// func (res *resultT) startTest(part string, i int, t *testing.T) {
// 	switch f := res.testedFunc.(type) {
// 	case func(*types.InlineQueryResultCachedAudio) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachAudio, res.isExpectedErr, i)
// 		checkError(f(res.cachAudio), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedDocument) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachDocument, res.isExpectedErr, i)
// 		checkError(f(res.cachDocument), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedGif) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachGif, res.isExpectedErr, i)
// 		checkError(f(res.cachGif), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedMpeg4Gif) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachMpeg4Gif, res.isExpectedErr, i)
// 		checkError(f(res.cachMpeg4Gif), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedPhoto) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachPhoto, res.isExpectedErr, i)
// 		checkError(f(res.cachPhoto), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedSticker) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachSticker, res.isExpectedErr, i)
// 		checkError(f(res.cachSticker), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedVideo) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachVideo, res.isExpectedErr, i)
// 		checkError(f(res.cachVideo), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultCachedVoice) error:
// 		printTestLog(part, res.name, res.codeErr, res.cachVoice, res.isExpectedErr, i)
// 		checkError(f(res.cachVoice), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultArticle) error:
// 		printTestLog(part, res.name, res.codeErr, res.art, res.isExpectedErr, i)
// 		checkError(f(res.art), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultAudio) error:
// 		printTestLog(part, res.name, res.codeErr, res.ad, res.isExpectedErr, i)
// 		checkError(f(res.ad), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultContact) error:
// 		printTestLog(part, res.name, res.codeErr, res.con, res.isExpectedErr, i)
// 		checkError(f(res.con), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultGame) error:
// 		printTestLog(part, res.name, res.codeErr, res.g, res.isExpectedErr, i)
// 		checkError(f(res.g), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultDocument) error:
// 		printTestLog(part, res.name, res.codeErr, res.doc, res.isExpectedErr, i)
// 		checkError(f(res.doc), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultGif) error:
// 		printTestLog(part, res.name, res.codeErr, res.gf, res.isExpectedErr, i)
// 		checkError(f(res.gf), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultLocation) error:
// 		printTestLog(part, res.name, res.codeErr, res.loc, res.isExpectedErr, i)
// 		checkError(f(res.loc), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultMpeg4Gif) error:
// 		printTestLog(part, res.name, res.codeErr, res.mpg4gf, res.isExpectedErr, i)
// 		checkError(f(res.mpg4gf), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultPhoto) error:
// 		printTestLog(part, res.name, res.codeErr, res.ph, res.isExpectedErr, i)
// 		checkError(f(res.ph), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultVenue) error:
// 		printTestLog(part, res.name, res.codeErr, res.ven, res.isExpectedErr, i)
// 		checkError(f(res.ven), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultVideo) error:
// 		printTestLog(part, res.name, res.codeErr, res.vd, res.isExpectedErr, i)
// 		checkError(f(res.vd), res.isExpectedErr, res.codeErr, t)
// 	case func(*types.InlineQueryResultVoice) error:
// 		printTestLog(part, res.name, res.codeErr, res.vc, res.isExpectedErr, i)
// 		checkError(f(res.vc), res.isExpectedErr, res.codeErr, t)
// 	default:
// 		t.Fatal("unexpected type of tested function")
// 	}
// }

// func (rtc *restTestContainer) createTestArrays(msg *formatter.Message, t *testing.T) ([]UnitTest, []UnitTest) {
// 	var r formatter.IResult
// 	var err error
// 	a, b := make([]UnitTest, rtc.until), make([]UnitTest, rtc.amount-rtc.until)
// 	for i, j := 0, 0; i < rtc.amount; i++ {
// 		s := rtc.byDefault(i)
// 		if i < rtc.until {
// 			if r, err = msg.NewInlineMode().WriteResult(); err != nil {
// 				t.Fatal(err)
// 			}
// 			rtc.buildF(r, s, i)
// 			a[i] = s
// 		} else {
// 			if j%2 == 0 {
// 				if r, err = msg.NewInlineMode().WriteResult(); err != nil {
// 					t.Fatal(err)
// 				}
// 			}
// 			rtc.buildF(r, s, i)
// 			b[j] = s
// 			j++
// 		}
// 	}
// 	return a, b
// }

// func (rtc *restTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
// 	rcontainer, doublecontainer := rtc.createTestArrays(msg, t)
// 	check(rcontainer, t)
// 	doubleCheck(doublecontainer, t)
// }

// func TestResultWriteCachedAudio(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedAudio()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedDocument(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedDocument()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedGif(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedGif()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedMpeg4Gif(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedMpeg4Gif()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedPhoto(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedPhoto()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedSticker(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedSticker()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedVideo(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedVideo()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteCachedVoice(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeCachedVoice()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteArticle(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeArticle()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteAudio(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeAudio()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteContact(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeContact()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteGame(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeGame()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteDocument(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeDocument()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteGif(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeGif()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteLocation(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeLocation()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteMpeg4Gif(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeMpeg4Gif()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWritePhoto(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writePhoto()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteVenue(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeVenue()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteVideo(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeVideo()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }

// func TestResultWriteVoice(t *testing.T) {
// 	rtc := new(restTestContainer)
// 	rtc.writeVoice()
// 	msg := formatter.CreateEmpltyMessage()
// 	rtc.mainLogic(msg, t)
// }
