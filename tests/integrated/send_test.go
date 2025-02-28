package integrated

import (
	"math/rand"
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

func (tc *testcase) sendMessage(t *testing.T) {
	var err error
	for i := 0; i < 4; i++ {
		tt := new(testcase)
		tt.init()
		if i == 3 {
			if err = tt.prm.WriteEntities(entities); err != nil {
				t.Fatal(err)
			}
		} else {
			if err = tt.prm.WriteParseMode(parsemode[i]); err != nil {
				t.Fatal(err)
			}
		}
		if err = tt.prm.WriteString(textformsg); err != nil {
			t.Fatal(err)
		}
		if err = tt.ch.WriteChatID(chatid); err != nil {
			t.Fatal(err)
		}
		if err = tt.prm.WriteLinkPreviewOptions(linkpopt); err != nil {
			t.Fatal(err)
		}
		if err = tt.prm.WriteDisableNotification(); err != nil {
			t.Fatal(err)
		}
		if err = tt.prm.WriteProtectContent(); err != nil {
			t.Fatal(err)
		}
		if err = tt.prm.WriteMessageEffectID(msgEffect); err != nil {
			t.Fatal(err)
		}
		if err = tt.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.AddParameters(tt.prm); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.AddChat(tt.ch); err != nil {
			t.Fatal(err)
		}
		tc.kbF(tt, t)
		tc.send(tt.msg, t)
		tc.checkResponse(t, i)
		timetochange <- struct{}{}
	}
}

func msgReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	if err = tc.prm.WriteString(textformsg); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	tc.checkResponse(t, 0)
}

func msgAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(nil, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	for i := 0; i < 3; i++ {
		if i == 0 {
			t.Log("Test Message With Inline Keyboard")
			tc.kbF = tc.inlineKb
		} else if i == 1 {
			t.Log("Test Message With ReplyMarkup Keyboard")
			tc.kbF = tc.replyKb
		} else {
			t.Log("Test Message With ForceReply Keyboard")
			tc.kbF = tc.forceKb
		}
		tc.sendMessage(t)
	}
}

func TestSMessage(t *testing.T) {
	t.Run("Req", msgReq)
	t.Run("All", msgAll)
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

func stickerArr(tc *testcase) []func(string) error {
	return []func(string) error{tc.st.WriteStickerStorage, tc.st.WriteStickerTelegram, tc.st.WriteStickerInternet}
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

func addSticker(tc *testcase, t *testing.T) {
	if err := tc.msg.AddSticker(tc.st); err != nil {
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

func (tc *testcase) sendMediaReq(t *testing.T, isvoice bool) {
	var err error
	for i := 0; i < 3; i++ {
		test := new(testcase)
		test.init()
		if tc.paid {
			if err = test.prm.WriteStarCount(stars); err != nil {
				t.Fatal(err)
			}
		}
		if err = tc.mediaF(test)[i](tc.mediadata[i]); err != nil {
			t.Fatal(err)
		}
		if err = test.ch.WriteChatID(chatid); err != nil {
			t.Fatal(err)
		}
		if err = test.msg.AddChat(test.ch); err != nil {
			t.Fatal(err)
		}
		if err = test.msg.AddParameters(test.prm); err != nil {
			t.Fatal(err)
		}
		tc.addMedia(test, t)
		tc.send(test.msg, t)
		if isvoice && i == 2 {
			delete(tc.whattocheck, vc)
			tc.whattocheck[doc] = struct{}{}
		}
		tc.checkResponse(t, i)
		timetochange <- struct{}{}
	}
}

func (tc *testcase) sendMediaAll(t *testing.T, isvoice bool) {
	var err error
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			tcc := new(testcase)
			tcc.init()
			if tc.paid {
				if err = tcc.prm.WriteStarCount(stars); err != nil {
					t.Fatal(err)
				}
			}
			if !tc.withoutString {
				if err = tcc.prm.WriteString(textformsg); err != nil {
					t.Fatal(err)
				}
				if j == 3 {
					if err = tcc.prm.WriteEntities(entities); err != nil {
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
			tc.send(tcc.msg, t)
			if isvoice && i == 2 {
				delete(tc.whattocheck, vc)
				tc.whattocheck[doc] = struct{}{}
			} else {
				delete(tc.whattocheck, doc)
			}
			tc.checkResponse(t, i)
			timetochange <- struct{}{}
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
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	if err = tc.ch.WriteChatID(chatid); err != nil {
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
	tc.init()
	go tc.changeToken(photodata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[ph] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata[tc.token]
	tc.sendMediaReq(t, false)
}

func photoAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(photodata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[ph] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata[tc.token]
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
		tc.sendMediaAll(t, false)
	}
}

func TestPhoto(t *testing.T) {
	t.Run("Req", photoReq)
	t.Run("All", photoAll)
}

func audioCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WritePerformer(emojies[rand.Intn(8)]); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WriteTitle(topicnames[rand.Intn(4)]); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
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
	tc.init()
	go tc.changeToken(audiodata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[audio] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addAudio
	tc.mediaF = audioArr
	tc.mediadata = audiodata[tc.token]
	tc.code = [3]int{0, 0, 400}
	tc.errmsg = [3]string{"", "", "[ERROR:400] Bad Request: wrong type of the web page content"}
	tc.sendMediaReq(t, false)
}

func audioAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(audiodata, thumbaudio)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[audio] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addAudio
	tc.mediaF = audioArr
	tc.mediadata = audiodata[tc.token]
	tc.common = audioCommon
	tc.thumb = true
	tc.thumbnailF = audioThumb
	tc.thumbdata = thumbaudio[tc.token]
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
		tc.sendMediaAll(t, false)
	}
}

func TestAudio(t *testing.T) {
	t.Run("Req", audioReq)
	t.Run("All", audioAll)
}

func documentCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
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
	tc.init()
	go tc.changeToken(docdata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[doc] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addDoc
	tc.mediaF = docArr
	tc.mediadata = docdata[tc.token]
	tc.sendMediaReq(t, false)
}

func docAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(docdata, thumbdoc)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[doc] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addDoc
	tc.mediaF = docArr
	tc.mediadata = docdata[tc.token]
	tc.common = documentCommon
	tc.thumb = true
	tc.thumbnailF = docThumb
	tc.thumbdata = thumbdoc[tc.token]
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
		tc.sendMediaAll(t, false)
	}
}

func TestDocument(t *testing.T) {
	t.Run("Req", docReq)
	t.Run("All", docAll)
}

func videoCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
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
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
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
	tc.init()
	go tc.changeToken(videodata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[video] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata[tc.token]
	tc.sendMediaReq(t, false)
}

func videoAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(videodata, thumbvideo)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[video] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata[tc.token]
	tc.common = videoCommon
	tc.thumb = true
	tc.thumbnailF = videoThumb
	tc.thumbdata = thumbvideo[tc.token]
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
		tc.sendMediaAll(t, false)
	}
}

func TestVideo(t *testing.T) {
	t.Run("Req", videoReq)
	t.Run("All", videoAll)
}

func animCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
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
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
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
	tc.init()
	go tc.changeToken(animdata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[anim] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addAnim
	tc.mediaF = animArr
	tc.mediadata = animdata[tc.token]
	tc.sendMediaReq(t, false)
}

func animAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(animdata, thumbanim)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[anim] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addAnim
	tc.mediaF = animArr
	tc.mediadata = animdata[tc.token]
	tc.common = animCommon
	tc.thumb = true
	tc.thumbnailF = animThumb
	tc.thumbdata = thumbanim[tc.token]
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
		tc.sendMediaAll(t, false)
	}
}

func TestAnimation(t *testing.T) {
	t.Run("Req", animReq)
	t.Run("All", animAll)
}

func voiceCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
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
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
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
	tc.init()
	go tc.changeToken(voicedata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[vc] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVoice
	tc.mediaF = voiceArr
	tc.mediadata = voicedata[tc.token]
	tc.sendMediaReq(t, true)
}

func voiceAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(voicedata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[vc] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVoice
	tc.mediaF = voiceArr
	tc.mediadata = voicedata[tc.token]
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
		tc.sendMediaAll(t, true)
	}
}

func TestVoice(t *testing.T) {
	t.Run("Req", voiceReq)
	t.Run("All", voiceAll)
}

func videoNCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
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
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
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
	tc.init()
	go tc.changeToken(videoNdata, nil)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[vdn] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVideoNote
	tc.mediaF = videoNArr
	tc.mediadata = videoNdata[tc.token]
	tc.sendMediaReq(t, false)
}

func videoNAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	go tc.changeToken(videoNdata, thumbvideoN)
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	// tc.whattocheck[vdn] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVideoNote
	tc.mediaF = videoNArr
	tc.mediadata = videoNdata[tc.token]
	tc.common = videoNCommon
	tc.thumb = true
	tc.withoutString = true
	tc.thumbnailF = videoNThumb
	tc.thumbdata = thumbvideoN[tc.token]
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
		tc.sendMediaAll(t, false)
	}
}

func TestVideoNote(t *testing.T) {
	t.Run("Req", videoNReq)
	t.Run("All", videoNAll)
}

func paidPhReq(t *testing.T) {
	tc := new(testcase)
	tc.init()
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[ph] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata[types.BotID]
	tc.paid = true
	tc.sendMediaReq(t, false)
}

func paidVdReq(t *testing.T) {
	tc := new(testcase)
	tc.init()
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[vd] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata[types.BotID]
	tc.paid = true
	tc.sendMediaReq(t, false)
}

func paidReq(t *testing.T) {
	t.Run("Photo", paidPhReq)
	t.Run("Video", paidVdReq)
}

func paidPhAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[ph] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addPhoto
	tc.mediaF = photoArr
	tc.mediadata = photodata[types.BotID]
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
		tc.sendMediaAll(t, false)
	}
}

func paidVdAll(t *testing.T) {
	tc := new(testcase)
	tc.init()
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[vd] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
	tc.addMedia = addVideo
	tc.mediaF = videoArr
	tc.mediadata = videodata[types.BotID]
	tc.common = videoCommon
	tc.thumb = true
	tc.paid = true
	tc.thumbnailF = videoThumb
	tc.thumbdata = thumbvideo[types.BotID]
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
		tc.sendMediaAll(t, false)
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

func mgPhotoAll(ph formatter.IPhoto, q *int, t *testing.T) {
	var err error
	if err = ph.WriteCaption(textformsg); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = ph.WriteCaptionEntities(entities); err != nil {
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
	if err = vd.WriteCaption(textformsg); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = vd.WriteCaptionEntities(entities); err != nil {
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
		if err = vd.WriteThumbnail(thumbvideo[types.BotID][0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = vd.WriteThumbnailID(thumbvideo[types.BotID][1]); err != nil {
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
	if err = dc.WriteCaption(textformsg); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = dc.WriteCaptionEntities(entities); err != nil {
			t.Fatal(err)
		}
	} else if *q >= 0 {
		if err = dc.WriteParseMode(parsemode[*q]); err != nil {
			t.Fatal(err)
		}
	}
	if majorq == 0 {
		if err = dc.WriteThumbnail(thumbdoc[types.BotID][0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = dc.WriteThumbnailID(thumbdoc[types.BotID][1]); err != nil {
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
	if err = ad.WriteCaption(textformsg); err != nil {
		t.Fatal(err)
	}
	if *q == 3 {
		if err = ad.WriteCaptionEntities(entities); err != nil {
			t.Fatal(err)
		}
	} else if *q >= 0 {
		if err = ad.WriteParseMode(parsemode[*q]); err != nil {
			t.Fatal(err)
		}
	}
	if majorq == 0 {
		if err = ad.WriteThumbnail(thumbaudio[types.BotID][0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = ad.WriteThumbnailID(thumbaudio[types.BotID][1]); err != nil {
			t.Fatal(err)
		}
	}
	if err = ad.WriteDuration(22); err != nil {
		t.Fatal(err)
	}
	if err = ad.WritePerformer(emojies[rand.Intn(8)]); err != nil {
		t.Fatal(err)
	}
	if err = ad.WriteTitle(topicnames[rand.Intn(4)]); err != nil {
		t.Fatal(err)
	}
	*q++
	if *q > 3 {
		*q = 0
	}
}

func addMediaInGroup(tc *testcase, obj []int, q int, t *testing.T) {
	var p, v, a, d, justcounter int
	var err error
	for i := 0; i < tc.mg.amout; i++ {
		if obj[i] == photo {
			tc.mg.photos[p] = tc.msg.NewPhoto()
			if err = tc.mg.phFunc(tc.mg.photos[p])[q](photodata[types.BotID][q]); err != nil {
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
			if err = tc.mg.vdFunc(tc.mg.videos[v])[q](videodata[types.BotID][q]); err != nil {
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
			if err = tc.mg.adFunc(tc.mg.audios[a])[q](audiodata[types.BotID][q]); err != nil {
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
			if err = tc.mg.docFunc(tc.mg.documents[d])[q](docdata[types.BotID][q]); err != nil {
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

func mediaGroupData(tc *testcase, obj []int, objtype int) {
	var tp int
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgids] = struct{}{}
	tc.whattocheck[groupid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[msgs] = struct{}{}
	if objtype == photo {
		tc.mg.phFunc = photoF
		tc.mg.photos = make([]formatter.IPhoto, tc.mg.amout)
		tc.whattocheck[phs] = struct{}{}
	} else if objtype == video {
		tp = video
		tc.mg.vdFunc = videoF
		tc.mg.videos = make([]formatter.IVideo, tc.mg.amout)
		delete(tc.whattocheck, phs)
		tc.whattocheck[vds] = struct{}{}
	} else if objtype == audio {
		tp = audio
		tc.mg.adFunc = audioF
		tc.mg.audios = make([]formatter.IAudio, tc.mg.amout)
		delete(tc.whattocheck, vds)
		tc.whattocheck[ads] = struct{}{}
	} else if objtype == together {
		tc.mg.vdFunc = videoF
		tc.mg.phFunc = photoF
		tc.mg.videos = make([]formatter.IVideo, tc.mg.amout/2)
		tc.mg.photos = make([]formatter.IPhoto, tc.mg.amout/2)
		delete(tc.whattocheck, ads)
		tc.whattocheck[phs] = struct{}{}
		tc.whattocheck[vds] = struct{}{}
	} else {
		tp = document
		tc.mg.docFunc = documentF
		tc.mg.documents = make([]formatter.IDocument, tc.mg.amout)
		delete(tc.whattocheck, phs)
		delete(tc.whattocheck, vds)
		tc.whattocheck[docs] = struct{}{}
	}
	for k := range obj {
		if objtype == together {
			if k%2 == 0 {
				obj[k] = video
			}
		} else {
			obj[k] = tp
		}
	}
}

func mediaGroup(all bool, objtype int, t *testing.T) {
	var err error
	var obj []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tc := new(testcase)
			tc.mg = new(mediagroup)
			if j == 0 {
				tc.mg.amout = 2
				obj = make([]int, tc.mg.amout)
			} else {
				tc.mg.amout = 10
				obj = make([]int, tc.mg.amout)
			}
			tc.init()
			mediaGroupData(tc, obj, objtype)
			addMediaInGroup(tc, obj, i, t)
			if all {
				mediaGroupCommon(tc, t)
			} else {
				if err = tc.ch.WriteChatID(chatid); err != nil {
					t.Fatal(err)
				}
				if err = tc.msg.AddChat(tc.ch); err != nil {
					t.Fatal(err)
				}
			}
			tc.send(tc.msg, t)
			tc.checkResponse(t, i)
		}
	}
}

func mediaReq(t *testing.T) {
	t.Run("Photo", func(t *testing.T) { mediaGroup(false, photo, t) })
	t.Run("Video", func(t *testing.T) { mediaGroup(false, video, t) })
	t.Run("TogetherPhVd", func(t *testing.T) { mediaGroup(false, together, t) })
	t.Run("Document", func(t *testing.T) { mediaGroup(false, document, t) })
	t.Run("Audio", func(t *testing.T) { mediaGroup(false, audio, t) })
}

func mediaGroupCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
}

func mediaAll(t *testing.T) {
	t.Run("Photo", func(t *testing.T) { mediaGroup(true, photo, t) })
	t.Run("Video", func(t *testing.T) { mediaGroup(true, video, t) })
	t.Run("TogetherPhVd", func(t *testing.T) { mediaGroup(true, together, t) })
	t.Run("Document", func(t *testing.T) { mediaGroup(true, document, t) })
	t.Run("Audio", func(t *testing.T) { mediaGroup(true, audio, t) })
}

func TestMediaGroup(t *testing.T) {
	t.Run("Req", mediaReq)
	t.Run("All", mediaAll)
}

func locReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteLatitude(latitude); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteLongitude(longitude); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddLocation(tc.loc); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func locAll(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		tc := new(testcase)
		tc.init()
		var kb = [3]func(*testcase, *testing.T){tc.replyKb, tc.inlineKb, tc.forceKb}
		if err = tc.ch.WriteChatID(chatid); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteLatitude(latitude); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteLongitude(longitude); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteHorizontalAccuracy(22); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteLivePeriod(0x7FFFFFFF); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteHeading(360); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteProximityAlertRadius(1); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteDisableNotification(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteProtectContent(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddParameters(tc.prm); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddLocation(tc.loc); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddChat(tc.ch); err != nil {
			t.Fatal(err)
		}
		kb[i](tc, t)
		tc.send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 && msg != "" {
			t.Fatal(msg)
		}
	}
}

func TestLocation(t *testing.T) {
	t.Run("Req", locReq)
	t.Run("All", locAll)
}

func venReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteLatitude(latitude); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteLongitude(longitude); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteTitle(placetitle); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteAddress(city); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddLocation(tc.loc); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func venAll(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		tc := new(testcase)
		tc.init()
		var kb = [3]func(*testcase, *testing.T){tc.replyKb, tc.inlineKb, tc.forceKb}
		if err = tc.ch.WriteChatID(chatid); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteLatitude(latitude); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteLongitude(longitude); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteTitle(placetitle); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteAddress(city); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteFoursquareID("44"); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteFoursquareType("arts_entertainment/default"); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteGooglePlaceID("AS:KLDKAL:SDK:LAK:LSL:"); err != nil {
			t.Fatal(err)
		}
		if err = tc.loc.WriteGooglePlaceType("bank"); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteDisableNotification(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteProtectContent(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddParameters(tc.prm); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddLocation(tc.loc); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddChat(tc.ch); err != nil {
			t.Fatal(err)
		}
		kb[i](tc, t)
		tc.send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 && msg != "" {
			t.Fatal(msg)
		}
	}
}

func TestVenue(t *testing.T) {
	t.Run("Req", venReq)
	t.Run("All", venAll)
}

func conReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.con.WritePhoneNumber(phonenum); err != nil {
		t.Fatal(err)
	}
	if err = tc.con.WriteFirstName(name); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddContact(tc.con); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func conAll(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		tc := new(testcase)
		tc.init()
		var kb = [3]func(*testcase, *testing.T){tc.replyKb, tc.inlineKb, tc.forceKb}
		if err = tc.ch.WriteChatID(chatid); err != nil {
			t.Fatal(err)
		}
		if err = tc.con.WritePhoneNumber(phonenum); err != nil {
			t.Fatal(err)
		}
		if err = tc.con.WriteFirstName(name); err != nil {
			t.Fatal(err)
		}
		if err = tc.con.WriteLastName(lastname); err != nil {
			t.Fatal(err)
		}
		if err = tc.con.WriteVCard(vcard); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteDisableNotification(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteProtectContent(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddParameters(tc.prm); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddContact(tc.con); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddChat(tc.ch); err != nil {
			t.Fatal(err)
		}
		kb[i](tc, t)
		tc.send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 && msg != "" {
			t.Fatal(msg)
		}
	}
}

func TestContact(t *testing.T) {
	t.Run("Req", conReq)
	t.Run("All", conAll)
}

func pollReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.poll.WriteQuestion(question); err != nil {
		t.Fatal(err)
	}
	if err = tc.poll.WriteOptions(pollOpt); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddPoll(tc.poll); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func pollAll(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				tc := new(testcase)
				tc.init()
				var kb = [3]func(*testcase, *testing.T){tc.replyKb, tc.inlineKb, tc.forceKb}
				if j == 3 {
					if err = tc.poll.WriteQuestionEntities(entities); err != nil {
						t.Fatal(err)
					}
				} else {
					if err = tc.poll.WriteQuestionParseMode(parsemode[j]); err != nil {
						t.Fatal(err)
					}
				}
				if k == 3 {
					if err = tc.poll.WriteExplanationEntities(entities); err != nil {
						t.Fatal(err)
					}
				} else {
					if err = tc.poll.WriteExplanationParseMode(parsemode[k]); err != nil {
						t.Fatal(err)
					}
				}
				if err = tc.ch.WriteChatID(chatid); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteQuestion(question); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteOptions(pollOpt); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteAnonymous(false); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteType("quiz"); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteAllowMultipleAnswers(); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteCorrectOptionID("00"); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteOpenPeriod(300); err != nil {
					t.Fatal(err)
				}
				if err = tc.prm.WriteDisableNotification(); err != nil {
					t.Fatal(err)
				}
				if err = tc.prm.WriteProtectContent(); err != nil {
					t.Fatal(err)
				}
				if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
					t.Fatal(err)
				}
				if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
					t.Fatal(err)
				}
				if err = tc.poll.WriteExplanation(explanation); err != nil {
					t.Fatal(err)
				}
				if err = tc.msg.AddPoll(tc.poll); err != nil {
					t.Fatal(err)
				}
				if err = tc.msg.AddChat(tc.ch); err != nil {
					t.Fatal(err)
				}
				if err = tc.msg.AddParameters(tc.prm); err != nil {
					t.Fatal(err)
				}
				kb[i](tc, t)
				tc.send(tc.msg, t)
				if code, msg := tc.get.Error(); code != 0 && msg != "" {
					t.Fatal(msg)
				}
			}
		}
	}
}

func TestPoll(t *testing.T) {
	t.Run("Req", pollReq)
	t.Run("All", pollAll)
}

func diceReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteEmoji(types.Emojis[0]); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func diceAll(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		tc := new(testcase)
		tc.init()
		var kb = [3]func(*testcase, *testing.T){tc.replyKb, tc.inlineKb, tc.forceKb}
		if err = tc.ch.WriteChatID(chatid); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteEmoji(types.Emojis[rand.Intn(5)]); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteDisableNotification(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteProtectContent(); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
			t.Fatal(err)
		}
		if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddParameters(tc.prm); err != nil {
			t.Fatal(err)
		}
		if err = tc.msg.AddChat(tc.ch); err != nil {
			t.Fatal(err)
		}
		kb[i](tc, t)
		tc.send(tc.msg, t)
		if code, msg := tc.get.Error(); code != 0 && msg != "" {
			t.Fatal(msg)
		}
	}
}

func TestDice(t *testing.T) {
	t.Run("Req", diceReq)
	t.Run("All", diceAll)
}

func chactReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteAction(types.Actions[rand.Intn(10)]); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	tc.send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func TestChatAction(t *testing.T) {
	t.Run("Req", chactReq)
}

// func stickerCommon(tc *testcase, t *testing.T) {
// 	var err error
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.st.WriteAssociatedEmoji("😁"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteDisableNotification(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteProtectContent(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageEffectID(msgEffect); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// }

func stickerReq(t *testing.T) {
	tc := new(testcase)
	tc.addMedia = addSticker
	tc.mediaF = stickerArr
	tc.mediadata = stickerdata[types.BotID]
	tc.sendMediaReq(t, false)
}

// func stickerAll(t *testing.T) {
// 	tc := new(testcase)
// 	tc.addMedia = addSticker
// 	tc.mediaF = stickerArr
// 	tc.mediadata = stickerdata
// 	tc.common = stickerCommon
// 	for i := 0; i < 3; i++ {
// 		if i == 0 {
// 			t.Log("Test Sticker With Inline Keyboard")
// 			tc.kbF = tc.inlineKb
// 		} else if i == 1 {
// 			t.Log("Test Sticker With ReplyMarkup Keyboard")
// 			tc.kbF = tc.replyKb
// 		} else {; err != nil {
// 		t.Fatal(err)
// 	}
// 		}
// 		tc.sendMediaAll(t, false)
// 	}
// }

func TestSticker(t *testing.T) {
	t.Run("Req", stickerReq)
	// t.Run("All", stickerAll)
}
