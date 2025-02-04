package integrated

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

var parsemode = []string{types.HTML, types.Markdown, types.MarkdownV2}
var photodata = []string{"../media/tel-aviv.jpg", "AgACAgIAAxkDAAIQOmeVLmLOfKYAAacYIhF6AAERhIoVlSYAAhjzMRvJOKlIR8eAgDzDgJIBAAMCAANtAAM2BA",
	"https://www.aljazeera.com/wp-content/uploads/2025/01/AFP__20250120__36UX28A__v2__HighRes__TopshotUsPoliticsTrumpInauguration-1737420954.jpg?w=770&resize=770,513"}
var audiodata = []string{"../media/sound.mp3", "CQACAgIAAxUHZ6IojQe818oNFe4bl9lH23Up7X0AAtVlAALmFxlJU-d8VdX3MFM2BA",
	"https://pixabay.com/music/download/vlog-music-beat-trailer-showreel-promo-background-intro-theme-274290.mp3"}
var docdata = []string{"../media/Resume.pdf", "BQACAgIAAxUHZ6Imx1aO5DIjSTDxSNIM5Sx8F5AAArhlAALmFxlJ4tfoQKuqKkE2BA",
	"https://www.aljazeera.com/wp-content/uploads/2025/01/AFP__20250120__36UX28A__v2__HighRes__TopshotUsPoliticsTrumpInauguration-1737420954.jpg?w=770&resize=770,513"}
var videodata = []string{"../media/black.mp4", "BAACAgIAAxkDAAIT-GeWRcV86dcv9CKSUWigMceO6OnTAALVYwACyTixSIm8GQnP7rn3NgQ",
	"https://www.pexels.com/download/video/6646588/"}
var animdata = []string{"../media/prichinatryski.mp4", "CgACAgIAAxkDAAIXLWehBtw_GBKJjOWUGUMuseLL5ilpAAL9YQAC5hcJSZl2JdZMJNA8NgQ",
	"https://www.icegif.com/wp-content/uploads/2025/02/patrick-mahomes-icegif-1.gif"}
var voicedata = []string{"../media/dimaJOSKAproNATO.ogg", "AwACAgIAAxkDAAIXd2ehChKgOVjVLvEfNW18seQfOzCAAAJDYgAC5hcJSdrW7HHHzSUDNgQ",
	"https://s33.aconvert.com/convert/p3r68-cdx67/0ye4j-z7u9r.ogg"}
var videoNdata = []string{"../media/black.mp4", "BAACAgIAAxkDAAIX02ehEJZfVRFXkXTl8wLAQJ5AXH41AALWYgAC5hcJSX6PBZ_I_kmPNgQ",
	"https://www.pexels.com/download/video/6646588/"}

var thumbaudio = []string{"../media/tel-aviv.jpg", "AAMCAgADFQdnoiiNB7zXyg0V7huX2UfbdSntfQAC1WUAAuYXGUlT53xV1fcwUwEAB20AAzYE"}
var thumbdoc = []string{"../media/tel-aviv.jpg", "AAMCAgADFQdnoibHVo7kMiNJMPFI0gzlLHwXkAACuGUAAuYXGUni1-hAq6oqQQEAB20AAzYE"}
var thumbvideo = []string{"../media/tel-aviv.jpg", "AAMCAgADGQMAAhbXZ5qcsV0fZsV-_N1uLTQVUcgp1lYAArFhAAKXWdhIeYErRMz_ObEBAAdtAAM2BA"}
var thumbanim = []string{"../media/tel-aviv.jpg", "AAMCAgADGQMAAhctZ6EG3D8YEomM5ZQZQy6x4svmKWkAAv1hAALmFwlJmXYl1kwk0DwBAAdtAAM2BA"}
var thumbvideoN = []string{"../media/tel-aviv.jpg", "AAMCAgADGQMAAhfTZ6EQll9VEVeRdOXzAsBAnkBcfjUAAtZiAALmFwlJfo8Fn8j-SY8BAAdtAAM2BA"}

type mediagroup struct {
	amout     int
	photos    []formatter.IPhoto
	videos    []formatter.IVideo
	audios    []formatter.IAudio
	documents []formatter.IDocument
	phFunc    func(formatter.IPhoto) []func(string) error
	vdFunc    func(formatter.IVideo) []func(string) error
	adFunc    func(formatter.IAudio) []func(string) error
	docFunc   func(formatter.IDocument) []func(string) error
	all       bool
}

type testcase struct {
	msg           *formatter.Message
	prm           formatter.IParameters
	ch            formatter.IChat
	ph            formatter.IPhoto
	vd            formatter.IVideo
	an            formatter.IAnimation
	get           formatter.IGet
	ad            formatter.IAudio
	dc            formatter.IDocument
	vc            formatter.IVoice
	vdn           formatter.IVideoNote
	loc           formatter.ILocation
	mg            *mediagroup
	addMedia      func(*testcase, *testing.T)
	mediaF        func(*testcase) []func(string) error
	mediadata     []string
	thumbdata     []string
	common        func(*testcase, *testing.T)
	kbF           func(*testcase, *testing.T)
	thumb         bool
	thumbnailF    func(*testcase) []func(string) error
	code          [3]int
	errmsg        [3]string
	withoutString bool
	paid          bool
}

func photoArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.ph.WritePhotoStorage, tc.ph.WritePhotoTelegram, tc.ph.WritePhotoInternet}
}

func videoArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.vd.WriteVideoStorage, tc.vd.WriteVideoTelegram, tc.vd.WriteVideoInternet}
}

func audioArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.ad.WriteAudioStorage, tc.ad.WriteAudioTelegram, tc.ad.WriteAudioInternet}
}

func docArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.dc.WriteDocumentStorage, tc.dc.WriteDocumentTelegram, tc.dc.WriteDocumentInternet}
}

func animArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.an.WriteAnimationStorage, tc.an.WriteAnimationTelegram, tc.an.WriteAnimationInternet}
}

func voiceArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.vc.WriteVoiceStorage, tc.vc.WriteVoiceTelegram, tc.vc.WriteVoiceInternet}
}

func videoNArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.vdn.WriteVideoNoteStorage, tc.vdn.WriteVideoNoteTelegram, tc.vdn.WriteVideoNoteInternet}
}

func addPhoto(tc *testcase, t *testing.T) {
	if err := tc.msg.AddPhoto(tc.ph); err != nil {
		t.Fatal(err)
	}
}

func addVideo(tc *testcase, t *testing.T) {
	if err := tc.msg.AddVideo(tc.vd); err != nil {
		t.Fatal(err)
	}
}

func addAudio(tc *testcase, t *testing.T) {
	if err := tc.msg.AddAudio(tc.ad); err != nil {
		t.Fatal(err)
	}
}

func addDoc(tc *testcase, t *testing.T) {
	if err := tc.msg.AddDocument(tc.dc); err != nil {
		t.Fatal(err)
	}
}

func addAnim(tc *testcase, t *testing.T) {
	if err := tc.msg.AddAnimation(tc.an); err != nil {
		t.Fatal(err)
	}
}

func addVoice(tc *testcase, t *testing.T) {
	if err := tc.msg.AddVoice(tc.vc); err != nil {
		t.Fatal(err)
	}
}

func addVideoNote(tc *testcase, t *testing.T) {
	if err := tc.msg.AddVideoNote(tc.vdn); err != nil {
		t.Fatal(err)
	}
}

func audioThumb(tc *testcase) []func(string) error {
	return []func(string) error{tc.ad.WriteThumbnail, tc.ad.WriteThumbnailID}
}

func docThumb(tc *testcase) []func(string) error {
	return []func(string) error{tc.dc.WriteThumbnail, tc.dc.WriteThumbnailID}
}

func videoThumb(tc *testcase) []func(string) error {
	return []func(string) error{tc.vd.WriteThumbnail, tc.vd.WriteThumbnailID}
}

func animThumb(tc *testcase) []func(string) error {
	return []func(string) error{tc.an.WriteThumbnail, tc.an.WriteThumbnailID}
}

func videoNThumb(tc *testcase) []func(string) error {
	return []func(string) error{tc.vdn.WriteThumbnail, tc.vdn.WriteThumbnailID}
}

func (tc *testcase) sendMediaReq(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		test := new(testcase)
		test.init()
		if tc.paid {
			if err = test.prm.WriteStarCount(31); err != nil {
				t.Fatal(err)
			}
		}
		if err = tc.mediaF(test)[i](tc.mediadata[i]); err != nil {
			t.Fatal(err)
		}
		if err = test.ch.WriteChatID(738070596); err != nil {
			t.Fatal(err)
		}
		if err = test.msg.AddChat(test.ch); err != nil {
			t.Fatal(err)
		}
		if err = test.msg.AddParameters(test.prm); err != nil {
			t.Fatal(err)
		}
		tc.addMedia(test, t)
		send(test.msg, t)
		if code, msg := test.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
			t.Fatal(msg)
		}
	}
}

func (tc *testcase) sendMediaAll(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			tcc := new(testcase)
			tcc.init()
			if tc.paid {
				if err = tcc.prm.WriteStarCount(31); err != nil {
					t.Fatal(err)
				}
			}
			if !tc.withoutString {
				if err = tcc.prm.WriteString("there's a string"); err != nil {
					t.Fatal(err)
				}
				if j == 3 {
					if err = tcc.prm.WriteEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
						t.Fatal(err)
					}
				} else if j >= 0 {
					if err = tcc.prm.WriteParseMode(parsemode[j]); err != nil {
						t.Fatal(err)
					}
				}
			}
			if err = tc.mediaF(tcc)[i](tc.mediadata[i]); err != nil {
				t.Fatal(err)
			}
			if tc.thumb {
				if i == 0 || i == 1 {
					if err = tc.thumbnailF(tcc)[i](tc.thumbdata[i]); err != nil {
						t.Fatal(err)
					}
				}
			}
			tc.common(tcc, t)
			tc.addMedia(tcc, t)
			tc.kbF(tcc, t)
			send(tcc.msg, t)
			if code, msg := tcc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func photoCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func photoReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata
	tc.sendMediaReq(t)
}

func photoAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata
	tc.common = photoCommon
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Photo With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Photo With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Photo With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestPhoto(t *testing.T) {
	t.Run("Required", photoReq)
	t.Run("All", photoAll)
}

func audioCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WritePerformer("performer"); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WriteTitle("title"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func audioReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addAudio
	tc.mediaF = audioArr
	tc.mediadata = audiodata
	tc.code = [3]int{0, 0, 400}
	tc.errmsg = [3]string{"", "", "[ERROR:400] Bad Request: wrong type of the web page content"}
	tc.sendMediaReq(t)
}

func audioAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addAudio
	tc.mediaF = audioArr
	tc.mediadata = audiodata
	tc.common = audioCommon
	tc.thumb = true
	tc.thumbnailF = audioThumb
	tc.thumbdata = thumbaudio
	tc.code = [3]int{0, 0, 400}
	tc.errmsg = [3]string{"", "", "[ERROR:400] Bad Request: wrong type of the web page content"}
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Audio With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Audio With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Audio With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestAudio(t *testing.T) {
	t.Run("Required", audioReq)
	t.Run("All", audioAll)
}

func documentCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func docReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addDoc
	tc.mediaF = docArr
	tc.mediadata = docdata
	tc.sendMediaReq(t)
}

func docAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addDoc
	tc.mediaF = docArr
	tc.mediadata = docdata
	tc.common = documentCommon
	tc.thumb = true
	tc.thumbnailF = docThumb
	tc.thumbdata = thumbdoc
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Document With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Document With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Document With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestDocument(t *testing.T) {
	t.Run("Required", docReq)
	t.Run("All", docAll)
}

func videoCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteWidth(55); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteSupportsStreaming(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func videoReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata
	tc.sendMediaReq(t)
}

func videoAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata
	tc.common = videoCommon
	tc.thumb = true
	tc.thumbnailF = videoThumb
	tc.thumbdata = thumbvideo
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Video With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Video With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Video With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestVideo(t *testing.T) {
	t.Run("Req", videoReq)
	t.Run("All", videoAll)
}

func animCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteWidth(55); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func animReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addAnim
	tc.mediaF = animArr
	tc.mediadata = animdata
	tc.sendMediaReq(t)
}

func animAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addAnim
	tc.mediaF = animArr
	tc.mediadata = animdata
	tc.common = animCommon
	tc.thumb = true
	tc.thumbnailF = animThumb
	tc.thumbdata = thumbanim
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Animation With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Animation With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Animation With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestAnimation(t *testing.T) {
	t.Run("Req", animReq)
	t.Run("All", animAll)
}

func voiceCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.vc.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func voiceReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVoice
	tc.mediaF = voiceArr
	tc.mediadata = voicedata
	tc.sendMediaReq(t)
}

func voiceAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVoice
	tc.mediaF = voiceArr
	tc.mediadata = voicedata
	tc.common = voiceCommon
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Voice With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Voice With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Voice With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestVoice(t *testing.T) {
	t.Run("Req", voiceReq)
	t.Run("All", voiceAll)
}

func videoNCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.vdn.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.vdn.WriteLength(55); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
}

func videoNReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVideoNote
	tc.mediaF = videoNArr
	tc.mediadata = videoNdata
	tc.sendMediaReq(t)
}

func videoNAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVideoNote
	tc.mediaF = videoNArr
	tc.mediadata = videoNdata
	tc.common = videoNCommon
	tc.thumb = true
	tc.withoutString = true
	tc.thumbnailF = videoNThumb
	tc.thumbdata = thumbvideoN
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Video Note With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Video Note With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Video Note With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func TestVideoNote(t *testing.T) {
	t.Run("Req", videoNReq)
	t.Run("All", videoNAll)
}

func paidPhReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata
	tc.paid = true
	tc.sendMediaReq(t)
}

func paidVdReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata
	tc.paid = true
	tc.sendMediaReq(t)
}

func paidReq(t *testing.T) {
	t.Run("Photo", paidPhReq)
	t.Run("Video", paidVdReq)
}

func paidPhAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata
	tc.paid = true
	tc.common = photoCommon
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Photo With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Photo With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Photo With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func paidVdAll(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata
	tc.common = videoCommon
	tc.thumb = true
	tc.paid = true
	tc.thumbnailF = videoThumb
	tc.thumbdata = thumbvideo
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Video With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Video With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Video With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMediaAll(t)
	}
}

func paidAll(t *testing.T) {
	t.Run("Photo", paidPhAll)
	t.Run("Video", paidVdAll)
}

func TestPaidMedia(t *testing.T) {
	t.Run("Req", paidReq)
	t.Run("All", paidAll)
}

func addMediaInGroup(tc *testcase, obj []int, q int, t *testing.T) {
	const photo, video, audio, document = 0, 1, 2, 3
	var p, v, a, d, justcounter int
	var err error
	for i := 0; i < tc.mg.amout; i++ {
		if obj[i] == photo {
			tc.mg.photos[p] = tc.msg.NewPhoto()
			if err = tc.mg.phFunc(tc.mg.photos[p])[q](photodata[q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgPhotoAll(tc.mg.photos[p], &justcounter, t)
			}
			if err = tc.msg.AddPhoto(tc.mg.photos[p]); err != nil {
				t.Fatal(err)
			}
			p++
		} else if obj[i] == video {
			tc.mg.videos[v] = tc.msg.NewVideo()
			if err = tc.mg.vdFunc(tc.mg.videos[v])[q](videodata[q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgVideoAll(tc.mg.videos[v], &justcounter, q, t)
			}
			if err = tc.msg.AddVideo(tc.mg.videos[v]); err != nil {
				t.Fatal(err)
			}
			v++
		} else if obj[i] == audio {
			tc.mg.audios[a] = tc.msg.NewAudio()
			if err = tc.mg.adFunc(tc.mg.audios[a])[q](audiodata[q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgAudioAll(tc.mg.audios[a], &justcounter, q, t)
			}
			if err = tc.msg.AddAudio(tc.mg.audios[a]); err != nil {
				t.Fatal(err)
			}
			a++
		} else {
			tc.mg.documents[d] = tc.msg.NewDocument()
			if err = tc.mg.docFunc(tc.mg.documents[d])[q](docdata[q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgDocumentAll(tc.mg.documents[d], &justcounter, q, t)
			}
			if err = tc.msg.AddDocument(tc.mg.documents[d]); err != nil {
				t.Fatal(err)
			}
			d++
		}
	}
}

func photoF(ph formatter.IPhoto) []func(string) error {
	return []func(string) error{ph.WritePhotoStorage, ph.WritePhotoTelegram, ph.WritePhotoInternet}
}

func videoF(vd formatter.IVideo) []func(string) error {
	return []func(string) error{vd.WriteVideoStorage, vd.WriteVideoTelegram, vd.WriteVideoInternet}
}

func documentF(doc formatter.IDocument) []func(string) error {
	return []func(string) error{doc.WriteDocumentStorage, doc.WriteDocumentTelegram, doc.WriteDocumentInternet}
}

func audioF(ad formatter.IAudio) []func(string) error {
	return []func(string) error{ad.WriteAudioStorage, ad.WriteAudioTelegram, ad.WriteAudioInternet}
}

func mediaPhReq(t *testing.T) {
	var err error
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			tc.init()
			tc.mg.phFunc = photoF
			tc.mg.photos = make([]formatter.IPhoto, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			if err = tc.ch.WriteChatID(738070596); err != nil {
				t.Fatal(err)
			}
			if err = tc.msg.AddChat(tc.ch); err != nil {
				t.Fatal(err)
			}
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaVdReq(t *testing.T) {
	var err error
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				obj[k] = 1
			}
			tc.init()
			tc.mg.vdFunc = videoF
			tc.mg.videos = make([]formatter.IVideo, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			if err = tc.ch.WriteChatID(738070596); err != nil {
				t.Fatal(err)
			}
			if err = tc.msg.AddChat(tc.ch); err != nil {
				t.Fatal(err)
			}
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaPhVdReq(t *testing.T) {
	var err error
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 4
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				if k%2 == 0 {
					obj[k] = 1
				}
			}
			tc.init()
			tc.mg.vdFunc = videoF
			tc.mg.phFunc = photoF
			tc.mg.videos = make([]formatter.IVideo, tc.mg.amout/2)
			tc.mg.photos = make([]formatter.IPhoto, tc.mg.amout/2)
			addMediaInGroup(tc, obj, i, t)
			if err = tc.ch.WriteChatID(738070596); err != nil {
				t.Fatal(err)
			}
			if err = tc.msg.AddChat(tc.ch); err != nil {
				t.Fatal(err)
			}
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaDocReq(t *testing.T) {
	var err error
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				obj[k] = 3
			}
			tc.init()
			tc.mg.docFunc = documentF
			tc.mg.documents = make([]formatter.IDocument, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			if err = tc.ch.WriteChatID(738070596); err != nil {
				t.Fatal(err)
			}
			if err = tc.msg.AddChat(tc.ch); err != nil {
				t.Fatal(err)
			}
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaAdReq(t *testing.T) {
	var err error
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				obj[k] = 2
			}
			tc.init()
			tc.mg.adFunc = audioF
			tc.mg.audios = make([]formatter.IAudio, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			if err = tc.ch.WriteChatID(738070596); err != nil {
				t.Fatal(err)
			}
			if err = tc.msg.AddChat(tc.ch); err != nil {
				t.Fatal(err)
			}
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaReq(t *testing.T) {
	t.Run("Photo", mediaPhReq)
	t.Run("Video", mediaVdReq)
	t.Run("TogetherPhVd", mediaPhVdReq)
	t.Run("Document", mediaDocReq)
	t.Run("Audio", mediaAdReq)
}

func mediaGroupCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID("5107584321108051014"); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(&types.ReplyParameters{MessageID: 3966, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
}

func mgPhotoAll(ph formatter.IPhoto, q *int, t *testing.T) {
	var err error
	if err = ph.WriteCaption("there's a string"); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = ph.WriteCaptionEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
			t.Fatal(err)
		}
	} else if *q >= 0 {
		if err = ph.WriteParseMode(parsemode[*q]); err != nil {
			t.Fatal(err)
		}
	}
	if err = ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	*q++
	if *q > 3 {
		*q = 0
	}
}

func mgVideoAll(vd formatter.IVideo, q *int, majorq int, t *testing.T) {
	var err error
	if err = vd.WriteCaption("there's a string"); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = vd.WriteCaptionEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
			t.Fatal(err)
		}
	} else if *q >= 0 {
		if err = vd.WriteParseMode(parsemode[*q]); err != nil {
			t.Fatal(err)
		}
	}
	if err = vd.WriteWidth(22); err != nil {
		t.Fatal(err)
	}
	if err = vd.WriteHeight(22); err != nil {
		t.Fatal(err)
	}
	if err = vd.WriteDuration(4); err != nil {
		t.Fatal(err)
	}
	if err = vd.WriteSupportsStreaming(); err != nil {
		t.Fatal(err)
	}
	if err = vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if majorq == 0 {
		if err = vd.WriteThumbnail(thumbvideo[0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = vd.WriteThumbnailID(thumbvideo[1]); err != nil {
			t.Fatal(err)
		}
	}
	*q++
	if *q > 3 {
		*q = 0
	}
}

func mgDocumentAll(dc formatter.IDocument, q *int, majorq int, t *testing.T) {
	var err error
	if err = dc.WriteCaption("there's a string"); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = dc.WriteCaptionEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
			t.Fatal(err)
		}
	} else if *q >= 0 {
		if err = dc.WriteParseMode(parsemode[*q]); err != nil {
			t.Fatal(err)
		}
	}
	if majorq == 0 {
		if err = dc.WriteThumbnail(thumbdoc[0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = dc.WriteThumbnailID(thumbdoc[1]); err != nil {
			t.Fatal(err)
		}
	}
	if err = dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	*q++
	if *q > 3 {
		*q = 0
	}
}

func mgAudioAll(ad formatter.IAudio, q *int, majorq int, t *testing.T) {
	var err error
	if err = ad.WriteCaption("there's a string"); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = ad.WriteCaptionEntities([]*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}); err != nil {
			t.Fatal(err)
		}
	} else if *q >= 0 {
		if err = ad.WriteParseMode(parsemode[*q]); err != nil {
			t.Fatal(err)
		}
	}
	if majorq == 0 {
		if err = ad.WriteThumbnail(thumbaudio[0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = ad.WriteThumbnailID(thumbaudio[1]); err != nil {
			t.Fatal(err)
		}
	}
	if err = ad.WriteDuration(22); err != nil {
		t.Fatal(err)
	}
	if err = ad.WritePerformer("Me Myself & I"); err != nil {
		t.Fatal(err)
	}
	if err = ad.WriteTitle("The Cool Song"); err != nil {
		t.Fatal(err)
	}
	*q++
	if *q > 3 {
		*q = 0
	}
}

func mediaPhAll(t *testing.T) {
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			tc.init()
			tc.mg.phFunc = photoF
			tc.mg.all = true
			tc.mg.photos = make([]formatter.IPhoto, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			mediaGroupCommon(tc, t)
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaVdAll(t *testing.T) {
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				obj[k] = 1
			}
			tc.init()
			tc.mg.vdFunc = videoF
			tc.mg.all = true
			tc.mg.videos = make([]formatter.IVideo, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			mediaGroupCommon(tc, t)
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaPhVdAll(t *testing.T) {
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 4
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				if k%2 == 0 {
					obj[k] = 1
				}
			}
			tc.init()
			tc.mg.all = true
			tc.mg.vdFunc = videoF
			tc.mg.phFunc = photoF
			tc.mg.videos = make([]formatter.IVideo, tc.mg.amout/2)
			tc.mg.photos = make([]formatter.IPhoto, tc.mg.amout/2)
			addMediaInGroup(tc, obj, i, t)
			mediaGroupCommon(tc, t)
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaDocAll(t *testing.T) {
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				obj[k] = 3
			}
			tc.init()
			tc.mg.docFunc = documentF
			tc.mg.all = true
			tc.mg.documents = make([]formatter.IDocument, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			mediaGroupCommon(tc, t)
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaAdAll(t *testing.T) {
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 3
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			for k := range obj {
				obj[k] = 2
			}
			tc.init()
			tc.mg.adFunc = audioF
			tc.mg.all = true
			tc.mg.audios = make([]formatter.IAudio, tc.mg.amout)
			addMediaInGroup(tc, obj, i, t)
			mediaGroupCommon(tc, t)
			send(tc.msg, t)
			if code, msg := tc.get.Error(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				t.Fatal(msg)
			}
		}
	}
}

func mediaAll(t *testing.T) {
	t.Run("Photo", mediaPhAll)
	t.Run("Video", mediaVdAll)
	t.Run("TogetherPhVd", mediaPhVdAll)
	t.Run("Document", mediaDocAll)
	t.Run("Audio", mediaAdAll)
}

func TestMediaGroup(t *testing.T) {
	t.Run("Req", mediaReq)
	t.Run("All", mediaAll)
}

func locReq(t *testing.T) {}

func locAll(t *testing.T) {}

func TestLocation(t *testing.T) {
	t.Run("Req", locReq)
	t.Run("All", locAll)
}
