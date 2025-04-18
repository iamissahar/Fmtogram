package integrated

import (
	"fmt"
	"math/rand"
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
	"github.com/iamissahar/Fmtogram/testbotdata"
)

func (tc *testcase) defaultSet() {
	tc.whattocheck[status] = struct{}{}
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[chat] = struct{}{}
	tc.whattocheck[sender] = struct{}{}
	tc.whattocheck[date] = struct{}{}
	tc.whattocheck[msgid] = struct{}{}
	tc.whattocheck[replyed] = struct{}{}
	tc.whattocheck[msg] = struct{}{}
}

func sendMessage(tc *testcase, t *testing.T, title string, kbF func(*testcase, *testing.T)) {
	var err error
	for i := 0; i < 4; i++ {
		tt := new(testcase)
		tt.init()
		if i == 3 {
			if err = tt.msg.WriteEntities(entities); err != nil {
				t.Fatal(err)
			}
		} else {
			if err = tt.msg.WriteParseMode(parsemode[i]); err != nil {
				t.Fatal(err)
			}
		}
		if err = tt.msg.WriteString(textformsg); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.WriteLinkPreviewOptions(linkpopt); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.WriteDisableNotification(); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.WriteProtectContent(); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.WriteMessageEffectID(msgEffect); err != nil {
			t.Fatal(err)
		}
		if err = tt.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
			t.Fatal(err)
		}
		kbF(tt, t)
		if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
			t.Fatal(err)
		}
		t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
		t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
		tt.checkResponse(i)
		tc.timetochange <- struct{}{}
	}
}

func msgReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
	if err = tc.msg.WriteString(textformsg); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "message", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "message", tc.get.Response())
	tc.checkResponse(0)
}

func TestSendMessage(t *testing.T) {
	t.Run("Req", msgReq)
	t.Run("All", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			t.Run(kbnames[i], func(t *testing.T) {
				t.Parallel()
				tc := new(testcase)
				defer func() { tc.workdone <- struct{}{} }()
				tc.init()
				tc.defaultSet()
				go tc.changeToken(nil, nil, t)
				sendMessage(tc, t, kbnames[i], kb[i])
			})
		}
	})
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
	if err := tc.msg.WritePhoto(tc.ph); err != nil {
		t.Fatal(err)
	}
}

func addVideo(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteVideo(tc.vd); err != nil {
		t.Fatal(err)
	}
}

func addAudio(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteAudio(tc.ad); err != nil {
		t.Fatal(err)
	}
}

func addDoc(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteDocument(tc.dc); err != nil {
		t.Fatal(err)
	}
}

func addAnim(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteAnimation(tc.an); err != nil {
		t.Fatal(err)
	}
}

func addVoice(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteVoice(tc.vc); err != nil {
		t.Fatal(err)
	}
}

func addVideoNote(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteVideoNote(tc.vdn); err != nil {
		t.Fatal(err)
	}
}

func addSticker(tc *testcase, t *testing.T) {
	if err := tc.msg.WriteSticker(tc.st); err != nil {
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

func sendMediaReq(t *testing.T, tc *testcase, addMorePrm []func(*testcase), title string, i int) {
	var err error
	tt := new(testcase)
	tt.init()
	tt.defaultSet()
	addMorePrm[i](tt)
	if tc.paid {
		if err = tt.msg.WriteStarCount(stars); err != nil {
			t.Fatal(err)
		}
	}
	if err = tc.mediaF(tt)[i](tc.mediadata[i]); err != nil {
		t.Fatal(err)
	}
	tc.addMedia(tt, t)
	if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
	tt.checkResponse(i)
	tc.timetochange <- struct{}{}
}

func sendMediaAll(t *testing.T, tc *testcase, addMorePrm []func(*testcase), kbF func(*testcase, *testing.T), title string) {
	var err error
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			<-tc.allowtocontinue
			tt := new(testcase)
			tt.init()
			tt.defaultSet()
			addMorePrm[i](tt)
			if tc.paid {
				if err = tt.msg.WriteStarCount(stars); err != nil {
					t.Fatal(err)
				}
			}
			if !tc.withoutString {
				if err = tt.msg.WriteString(textformsg); err != nil {
					t.Fatal(err)
				}
				if j == 3 {
					if err = tt.msg.WriteEntities(entities); err != nil {
						t.Fatal(err)
					}
				} else if j >= 0 {
					if err = tt.msg.WriteParseMode(parsemode[j]); err != nil {
						t.Fatal(err)
					}
				}
			}
			if err = tc.mediaF(tt)[i](tc.mediadata[i]); err != nil {
				t.Fatal(err)
			}
			if tc.thumb {
				if i == 0 || i == 1 {
					if err = tc.thumbnailF(tt)[i](tc.thumbdata[i]); err != nil {
						t.Fatal(err)
					}
				}
			}
			tc.common(tc, tt, t)
			tc.addMedia(tt, t)
			kbF(tt, t)
			if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
				t.Fatal(err)
			}
			t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
			t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
			tt.checkResponse(i)
			tc.timetochange <- struct{}{}
		}
	}
}

func photoCommon(oldtc, tc *testcase, t *testing.T) {
	var err error
	if err = tc.ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func photoIntoMap(tc *testcase) {
	tc.whattocheck[ph] = struct{}{}
}

func photoReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(photodata, nil, t)
			tc.addMedia = addPhoto
			tc.mediaF = photoArr
			tc.mediadata = photodata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapPhoto, fmt.Sprintf("%s-photo", filenames[i]), i)
		})
	}
}

func photoAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(photodata, nil, t)
			tc.addMedia = addPhoto
			tc.mediaF = photoArr
			tc.mediadata = photodata[tc.bs.Token]
			tc.common = photoCommon
			sendMediaAll(t, tc, addToMapPhoto, kb[i], kbnames[i])
		})
	}
}

func TestSendPhoto(t *testing.T) {
	t.Run("Req", photoReq)
	t.Run("All", photoAll)
}

func audioCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.ad.WriteDuration(3); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WritePerformer(emojies[rand.Intn(8)]); err != nil {
		t.Fatal(err)
	}
	if err = tc.ad.WriteTitle(topicnames[rand.Intn(4)]); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func audioIntoMap(tc *testcase) {
	tc.whattocheck[ad] = struct{}{}
}

func audioReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(audiodata, nil, t)
			tc.addMedia = addAudio
			tc.mediaF = audioArr
			tc.mediadata = audiodata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapAudio, fmt.Sprintf("%s-audio", filenames[i]), i)
		})
	}
}

func audioAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(audiodata, thumbaudio, t)
			tc.addMedia = addAudio
			tc.mediaF = audioArr
			tc.mediadata = audiodata[tc.bs.Token]
			tc.common = audioCommon
			tc.thumb = true
			tc.thumbnailF = audioThumb
			tc.thumbdata = thumbaudio[tc.bs.Token]
			sendMediaAll(t, tc, addToMapAudio, kb[i], kbnames[i])
		})
	}
}

func TestSendAudio(t *testing.T) {
	t.Run("Req", audioReq)
	t.Run("All", audioAll)
}

func documentCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.dc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func docIntoMap(tc *testcase) {
	tc.whattocheck[doc] = struct{}{}
}

func docReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(docdata, nil, t)
			tc.addMedia = addDoc
			tc.mediaF = docArr
			tc.mediadata = docdata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapDoc, fmt.Sprintf("%s-document", filenames[i]), i)
		})
	}
}

func docAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(docdata, thumbdoc, t)
			tc.addMedia = addDoc
			tc.mediaF = docArr
			tc.mediadata = docdata[tc.bs.Token]
			tc.common = documentCommon
			tc.thumb = true
			tc.thumbnailF = docThumb
			tc.thumbdata = thumbdoc[tc.bs.Token]
			sendMediaAll(t, tc, addToMapDoc, kb[i], kbnames[i])
		})
	}
}

func TestSendDocument(t *testing.T) {
	t.Run("Req", docReq)
	t.Run("All", docAll)
}

func videoCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.vd.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteWidth(55); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err = tc.vd.WriteSupportsStreaming(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func vidIntoMap(tc *testcase) {
	tc.whattocheck[vd] = struct{}{}
}

func videoReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(videodata, nil, t)
			tc.addMedia = addVideo
			tc.mediaF = videoArr
			tc.mediadata = videodata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapVideo, fmt.Sprintf("%s-video", filenames[i]), i)
		})
	}
}

func videoAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(videodata, thumbvideo, t)
			tc.addMedia = addVideo
			tc.mediaF = videoArr
			tc.mediadata = videodata[tc.bs.Token]
			tc.common = videoCommon
			tc.thumb = true
			tc.thumbnailF = videoThumb
			tc.thumbdata = thumbvideo[tc.bs.Token]
			sendMediaAll(t, tc, addToMapVideo, kb[i], kbnames[i])
		})
	}
}

func TestSendVideo(t *testing.T) {
	t.Run("Req", videoReq)
	t.Run("All", videoAll)
}

func animCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.an.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteWidth(55); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func animIntoMap(tc *testcase) {
	tc.whattocheck[anim] = struct{}{}
}

func animReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(animdata, nil, t)
			tc.addMedia = addAnim
			tc.mediaF = animArr
			tc.mediadata = animdata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapAnim, fmt.Sprintf("%s-audio", filenames[i]), i)
		})
	}
}

func animAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(animdata, thumbanim, t)
			tc.addMedia = addAnim
			tc.mediaF = animArr
			tc.mediadata = animdata[tc.bs.Token]
			tc.common = animCommon
			tc.thumb = true
			tc.thumbnailF = animThumb
			tc.thumbdata = thumbanim[tc.bs.Token]
			sendMediaAll(t, tc, addToMapAnim, kb[i], kbnames[i])
		})
	}
}

func TestSendAnimation(t *testing.T) {
	t.Run("Req", animReq)
	t.Run("All", animAll)
}

func voiceCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.vc.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.an.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func vcIntoMap(tc *testcase) {
	tc.whattocheck[vc] = struct{}{}
}

func vcIntoMap1(tc *testcase) {
	tc.whattocheck[doc] = struct{}{}
}

func voiceReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(voicedata, nil, t)
			tc.addMedia = addVoice
			tc.mediaF = voiceArr
			tc.mediadata = voicedata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapVoice, fmt.Sprintf("%s-voice", filenames[i]), i)
		})
	}
}

func voiceAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(voicedata, nil, t)
			tc.addMedia = addVoice
			tc.mediaF = voiceArr
			tc.mediadata = voicedata[tc.bs.Token]
			tc.common = voiceCommon
			sendMediaAll(t, tc, addToMapVoice, kb[i], kbnames[i])
		})
	}
}

func TestSendVoice(t *testing.T) {
	t.Run("Req", voiceReq)
	t.Run("All", voiceAll)
}

func videoNCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.vdn.WriteDuration(111); err != nil {
		t.Fatal(err)
	}
	if err = tc.vdn.WriteLength(55); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func vdnIntoMap(tc *testcase) {
	tc.whattocheck[vdn] = struct{}{}
}

func vdnIntoMap1(tc *testcase) {
	tc.whattocheck[vd] = struct{}{}
}

func videoNReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(videoNdata, nil, t)
			tc.addMedia = addVideoNote
			tc.mediaF = videoNArr
			tc.mediadata = videoNdata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapVdn, fmt.Sprintf("%s-videonote", filenames[i]), i)
		})
	}
}

func videoNAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(videoNdata, thumbvideoN, t)
			tc.addMedia = addVideoNote
			tc.mediaF = videoNArr
			tc.mediadata = videoNdata[tc.bs.Token]
			tc.common = videoNCommon
			tc.thumb = true
			tc.withoutString = true
			tc.thumbnailF = videoNThumb
			tc.thumbdata = thumbvideoN[tc.bs.Token]
			sendMediaAll(t, tc, addToMapVdn1, kb[i], kbnames[i])
		})
	}
}

func TestSendNoteVideo(t *testing.T) {
	t.Run("Req", videoNReq)
	t.Run("All", videoNAll)
}

func paidPhReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(photodata, nil, t)
			tc.addMedia = addPhoto
			tc.mediaF = photoArr
			tc.mediadata = photodata[tc.bs.Token]
			tc.paid = true
			sendMediaReq(t, tc, addToMapPhoto, fmt.Sprintf("%s-paid-photo", filenames[i]), i)
		})
	}
}

func paidVdReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(videodata, nil, t)
			tc.addMedia = addVideo
			tc.mediaF = videoArr
			tc.mediadata = videodata[tc.bs.Token]
			tc.paid = true
			sendMediaReq(t, tc, addToMapVideo, fmt.Sprintf("%s-paid-video", filenames[i]), i)
		})
	}
}

func paidReq(t *testing.T) {
	t.Run("Photo", paidPhReq)
	t.Run("Video", paidVdReq)
}

func paidPhAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(photodata, nil, t)
			tc.addMedia = addPhoto
			tc.mediaF = photoArr
			tc.mediadata = photodata[tc.bs.Token]
			tc.paid = true
			tc.common = photoCommon
			sendMediaAll(t, tc, addToMapPhoto, kb[i], kbnames[i])
		})
	}
}

func paidVdAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(videodata, thumbvideo, t)
			tc.addMedia = addVideo
			tc.mediaF = videoArr
			tc.mediadata = videodata[tc.bs.Token]
			tc.common = videoCommon
			tc.thumb = true
			tc.paid = true
			tc.thumbnailF = videoThumb
			tc.thumbdata = thumbvideo[tc.bs.Token]
			sendMediaAll(t, tc, addToMapVideo, kb[i], kbnames[i])
		})
	}
}

func paidAll(t *testing.T) {
	t.Run("Photo", paidPhAll)
	t.Run("Video", paidVdAll)
}

func TestSendPaidMedia(t *testing.T) {
	t.Run("Req", paidReq)
	t.Run("All", paidAll)
}

func mgPhotoAll(ph fmtogram.IPhoto, q *int, t *testing.T) {
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

func mgVideoAll(vd fmtogram.IVideo, q *int, majorq int, token string, t *testing.T) {
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
		if err = vd.WriteThumbnail(thumbvideo[token][0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = vd.WriteThumbnailID(thumbvideo[token][1]); err != nil {
			t.Fatal(err)
		}
	}
	*q++
	if *q > 3 {
		*q = 0
	}
}

func mgDocumentAll(dc fmtogram.IDocument, q *int, majorq int, token string, t *testing.T) {
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
		if err = dc.WriteThumbnail(thumbdoc[token][0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = dc.WriteThumbnailID(thumbdoc[token][1]); err != nil {
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

func mgAudioAll(ad fmtogram.IAudio, q *int, majorq int, token string, t *testing.T) {
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
		if err = ad.WriteThumbnail(thumbaudio[token][0]); err != nil {
			t.Fatal(err)
		}
	} else if majorq == 1 {
		if err = ad.WriteThumbnailID(thumbaudio[token][1]); err != nil {
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
			tc.mg.photos[p] = fmtogram.NewPhoto()
			if err = tc.mg.phFunc(tc.mg.photos[p])[q](photodata[tc.bs.Token][q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgPhotoAll(tc.mg.photos[p], &justcounter, t)
			}
			if err = tc.msg.WritePhoto(tc.mg.photos[p]); err != nil {
				t.Fatal(err)
			}
			p++
		} else if obj[i] == video {
			tc.mg.videos[v] = fmtogram.NewVideo()
			if err = tc.mg.vdFunc(tc.mg.videos[v])[q](videodata[tc.bs.Token][q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgVideoAll(tc.mg.videos[v], &justcounter, q, tc.bs.Token, t)
			}
			if err = tc.msg.WriteVideo(tc.mg.videos[v]); err != nil {
				t.Fatal(err)
			}
			v++
		} else if obj[i] == audio {
			tc.mg.audios[a] = fmtogram.NewAudio()
			if err = tc.mg.adFunc(tc.mg.audios[a])[q](audiodata[tc.bs.Token][q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgAudioAll(tc.mg.audios[a], &justcounter, q, tc.bs.Token, t)
			}
			if err = tc.msg.WriteAudio(tc.mg.audios[a]); err != nil {
				t.Fatal(err)
			}
			a++
		} else {
			tc.mg.documents[d] = fmtogram.NewDocument()
			if err = tc.mg.docFunc(tc.mg.documents[d])[q](docdata[tc.bs.Token][q]); err != nil {
				t.Fatal(err)
			}
			if tc.mg.all {
				mgDocumentAll(tc.mg.documents[d], &justcounter, q, tc.bs.Token, t)
			}
			if err = tc.msg.WriteDocument(tc.mg.documents[d]); err != nil {
				t.Fatal(err)
			}
			d++
		}
	}
}

func photoF(ph fmtogram.IPhoto) []func(string) error {
	return []func(string) error{ph.WritePhotoStorage, ph.WritePhotoTelegram, ph.WritePhotoInternet}
}

func videoF(vd fmtogram.IVideo) []func(string) error {
	return []func(string) error{vd.WriteVideoStorage, vd.WriteVideoTelegram, vd.WriteVideoInternet}
}

func documentF(doc fmtogram.IDocument) []func(string) error {
	return []func(string) error{doc.WriteDocumentStorage, doc.WriteDocumentTelegram, doc.WriteDocumentInternet}
}

func audioF(ad fmtogram.IAudio) []func(string) error {
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
		tc.mg.photos = make([]fmtogram.IPhoto, tc.mg.amout)
		tc.whattocheck[phs] = struct{}{}
	} else if objtype == video {
		tp = video
		tc.mg.vdFunc = videoF
		tc.mg.videos = make([]fmtogram.IVideo, tc.mg.amout)
		delete(tc.whattocheck, phs)
		tc.whattocheck[vds] = struct{}{}
	} else if objtype == audio {
		tp = audio
		tc.mg.adFunc = audioF
		tc.mg.audios = make([]fmtogram.IAudio, tc.mg.amout)
		delete(tc.whattocheck, vds)
		tc.whattocheck[ads] = struct{}{}
	} else if objtype == together {
		tc.mg.vdFunc = videoF
		tc.mg.phFunc = photoF
		tc.mg.videos = make([]fmtogram.IVideo, tc.mg.amout/2)
		tc.mg.photos = make([]fmtogram.IPhoto, tc.mg.amout/2)
		delete(tc.whattocheck, ads)
		tc.whattocheck[phs] = struct{}{}
		tc.whattocheck[vds] = struct{}{}
	} else {
		tp = document
		tc.mg.docFunc = documentF
		tc.mg.documents = make([]fmtogram.IDocument, tc.mg.amout)
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

func mediaGroup(all bool, objtype int, title string, t *testing.T) {
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
			}
			if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
				t.Fatal(err)
			}
			t.Logf("[TEST:%s] Request:\n%s", title, tc.get.Request())
			t.Logf("[TEST:%s] Response:\n%s", title, tc.get.Response())
			tc.checkResponse(i)
		}
	}
}

func mediaReq(t *testing.T) {
	t.Run("Photo", func(t *testing.T) { t.Parallel(); mediaGroup(false, photo, "photo", t) })
	t.Run("Video", func(t *testing.T) { t.Parallel(); mediaGroup(false, video, "video", t) })
	t.Run("TogetherPhVd", func(t *testing.T) { t.Parallel(); mediaGroup(false, together, "photo+video", t) })
	t.Run("Document", func(t *testing.T) { t.Parallel(); mediaGroup(false, document, "document", t) })
	t.Run("Audio", func(t *testing.T) { t.Parallel(); mediaGroup(false, audio, "audio", t) })
}

func mediaGroupCommon(tc *testcase, t *testing.T) {
	var err error
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func mediaAll(t *testing.T) {
	t.Run("Photo", func(t *testing.T) { t.Parallel(); mediaGroup(true, photo, "photo", t) })
	t.Run("Video", func(t *testing.T) { t.Parallel(); mediaGroup(true, video, "video", t) })
	t.Run("TogetherPhVd", func(t *testing.T) { t.Parallel(); mediaGroup(true, together, "photo+video", t) })
	t.Run("Document", func(t *testing.T) { t.Parallel(); mediaGroup(true, document, "document", t) })
	t.Run("Audio", func(t *testing.T) { t.Parallel(); mediaGroup(true, audio, "audio", t) })
}

func TestSendMediaGroup(t *testing.T) {
	t.Run("Req", mediaReq)
	t.Run("All", mediaAll)
}

func locReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
	if err = tc.loc.WriteLatitude(latitude); err != nil {
		t.Fatal(err)
	}
	if err = tc.loc.WriteLongitude(longitude); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteLocation(tc.loc); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "location", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "location", tc.get.Response())
	tc.checkResponse(0)
}

func locAll(tc *testcase, t *testing.T, kb func(*testcase, *testing.T), title string, i int) {
	var err error
	tt := new(testcase)
	tt.init()
	tc.defaultSet()
	if err = tt.loc.WriteLatitude(latitude); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteLongitude(longitude); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteHorizontalAccuracy(22); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteLivePeriod(0x7FFFFFFF); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteHeading(360); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteProximityAlertRadius(1); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteLocation(tt.loc); err != nil {
		t.Fatal(err)
	}
	kb(tt, t)
	if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
	tt.checkResponse(i)
	tc.timetochange <- struct{}{}
}

func TestSendLocation(t *testing.T) {
	t.Run("Req", func(t *testing.T) {
		t.Parallel()
		locReq(t)
	})
	t.Run("All", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			t.Run(kbnames[i], func(t *testing.T) {
				t.Parallel()
				tc := new(testcase)
				tc.init()
				defer func() { tc.workdone <- struct{}{} }()
				go tc.changeToken(nil, nil, t)
				locAll(tc, t, kb[i], kbnames[i], i)
			})
		}
	})
}

func venReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
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
	if err = tc.msg.WriteLocation(tc.loc); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "venue", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "venue", tc.get.Response())
	tc.checkResponse(0)
}

func venAll(tc *testcase, t *testing.T, kb func(*testcase, *testing.T), title string, i int) {
	var err error
	tt := new(testcase)
	tt.init()
	tt.defaultSet()
	if err = tt.loc.WriteLatitude(latitude); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteLongitude(longitude); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteTitle(placetitle); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteAddress(city); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteFoursquareID("44"); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteFoursquareType("arts_entertainment/default"); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteGooglePlaceID("AS:KLDKAL:SDK:LAK:LSL:"); err != nil {
		t.Fatal(err)
	}
	if err = tt.loc.WriteGooglePlaceType("bank"); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteLocation(tt.loc); err != nil {
		t.Fatal(err)
	}
	kb(tt, t)
	if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
	tt.checkResponse(i)
	tc.timetochange <- struct{}{}
}

func TestSendVenue(t *testing.T) {
	t.Run("Req", func(t *testing.T) {
		t.Parallel()
		venReq(t)
	})
	t.Run("All", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			t.Run(kbnames[i], func(t *testing.T) {
				t.Parallel()
				tc := new(testcase)
				tc.init()
				defer func() { tc.workdone <- struct{}{} }()
				go tc.changeToken(nil, nil, t)
				venAll(tc, t, kb[i], kbnames[i], i)
			})
		}
	})
}

func conReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
	if err = tc.con.WritePhoneNumber(phonenum); err != nil {
		t.Fatal(err)
	}
	if err = tc.con.WriteFirstName(name); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteContact(tc.con); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "contact", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "contact", tc.get.Response())
	tc.checkResponse(0)
}

func conAll(tc *testcase, t *testing.T, kb func(*testcase, *testing.T), title string) {
	var err error
	tt := new(testcase)
	tt.init()
	tt.defaultSet()
	if err = tt.con.WritePhoneNumber(phonenum); err != nil {
		t.Fatal(err)
	}
	if err = tt.con.WriteFirstName(name); err != nil {
		t.Fatal(err)
	}
	if err = tt.con.WriteLastName(lastname); err != nil {
		t.Fatal(err)
	}
	if err = tt.con.WriteVCard(vcard); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteContact(tt.con); err != nil {
		t.Fatal(err)
	}
	kb(tt, t)
	if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
	tt.checkResponse(0)
	tc.timetochange <- struct{}{}
}

func TestSendContact(t *testing.T) {
	t.Run("Req", conReq)
	t.Run("All", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			t.Run(kbnames[i], func(t *testing.T) {
				t.Parallel()
				tc := new(testcase)
				tc.init()
				defer func() { tc.workdone <- struct{}{} }()
				go tc.changeToken(nil, nil, t)
				conAll(tc, t, kb[i], kbnames[i])
			})
		}
	})
}

func pollReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
	if err = tc.poll.WriteQuestion(question); err != nil {
		t.Fatal(err)
	}
	if err = tc.poll.WriteOptions(pollOpt); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WritePoll(tc.poll); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "poll", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "poll", tc.get.Response())
	tc.checkResponse(0)
}

func pollAll(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	defer func() { tc.workdone <- struct{}{} }()
	go tc.changeToken(nil, nil, t)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				<-tc.allowtocontinue
				tt := new(testcase)
				tt.init()
				if j == 3 {
					if err = tt.poll.WriteQuestionEntities(entities); err != nil {
						t.Fatal(err)
					}
				} else {
					if err = tt.poll.WriteQuestionParseMode(parsemode[j]); err != nil {
						t.Fatal(err)
					}
				}
				if k == 3 {
					if err = tt.poll.WriteExplanationEntities(entities); err != nil {
						t.Fatal(err)
					}
				} else {
					if err = tt.poll.WriteExplanationParseMode(parsemode[k]); err != nil {
						t.Fatal(err)
					}
				}
				if err = tt.poll.WriteQuestion(question); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteOptions(pollOpt); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteAnonymous(false); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteType("quiz"); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteAllowMultipleAnswers(); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteCorrectOptionID("00"); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteOpenPeriod(300); err != nil {
					t.Fatal(err)
				}
				if err = tt.msg.WriteDisableNotification(); err != nil {
					t.Fatal(err)
				}
				if err = tt.msg.WriteProtectContent(); err != nil {
					t.Fatal(err)
				}
				if err = tt.msg.WriteMessageEffectID(msgEffect); err != nil {
					t.Fatal(err)
				}
				if err = tt.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
					t.Fatal(err)
				}
				if err = tt.poll.WriteExplanation(explanation); err != nil {
					t.Fatal(err)
				}
				if err = tt.msg.WritePoll(tt.poll); err != nil {
					t.Fatal(err)
				}
				kb[i](tt, t)
				if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
					t.Fatal(err)
				}
				t.Logf("[TEST:%s] Request:\n%s", "poll", tt.get.Request())
				t.Logf("[TEST:%s] Response:\n%s", "poll", tt.get.Response())
				tt.checkResponse(i)
				tc.timetochange <- struct{}{}
			}
		}
	}
}

func TestSendPoll(t *testing.T) {
	t.Run("Req", pollReq)
	t.Run("All", pollAll)
}

func diceReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
	if err = tc.msg.WriteEmoji(fmtogram.Emojis[0]); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "poll", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "poll", tc.get.Response())
	tc.checkResponse(0)
}

func diceAll(tc *testcase, t *testing.T, kb func(*testcase, *testing.T), title string, i int) {
	var err error
	tt := new(testcase)
	tt.init()
	tt.defaultSet()
	if err = tt.msg.WriteEmoji(fmtogram.Emojis[rand.Intn(5)]); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tt.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	kb(tt, t)
	if err = tc.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", title, tt.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", title, tt.get.Response())
	tt.checkResponse(i)
}

func TestSendDice(t *testing.T) {
	t.Run("Req", diceReq)
	t.Run("All", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			t.Run(kbnames[i], func(t *testing.T) {
				t.Parallel()
				tc := new(testcase)
				tc.init()
				defer func() { tc.workdone <- struct{}{} }()
				go tc.changeToken(nil, nil, t)
				diceAll(tc, t, kb[i], kbnames[i], i)
			})
		}
	})
}

func chactReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.whattocheck[errr] = struct{}{}
	tc.whattocheck[status] = struct{}{}
	if err = tc.msg.WriteAction(fmtogram.Actions[rand.Intn(10)]); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "poll", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "poll", tc.get.Response())
	tc.checkResponse(0)
}

func TestSendChatAction(t *testing.T) {
	t.Run("Req", chactReq)
}

func stickerCommon(oldtc *testcase, tc *testcase, t *testing.T) {
	var err error
	if err = tc.st.WriteAssociatedEmoji("😁"); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(oldtc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
}

func stckrIntoMap(tc *testcase) {
	tc.whattocheck[stkrs] = struct{}{}
}

func stickerReq(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(filenames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(stickerdata, nil, t)
			tc.addMedia = addSticker
			tc.mediaF = stickerArr
			tc.mediadata = stickerdata[tc.bs.Token]
			sendMediaReq(t, tc, addToMapSticker, fmt.Sprintf("%s-sticker", filenames[i]), i)
		})
	}
}

func stickerAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(kbnames[i], func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			defer func() { tc.workdone <- struct{}{} }()
			tc.init()
			go tc.changeToken(stickerdata, nil, t)
			tc.addMedia = addSticker
			tc.mediaF = stickerArr
			tc.mediadata = stickerdata[tc.bs.Token]
			tc.common = stickerCommon
			sendMediaAll(t, tc, addToMapSticker, kb[i], kbnames[i])
		})
	}
}

func TestSendSticker(t *testing.T) {
	t.Run("Req", stickerReq)
	t.Run("All", stickerAll)
}

func giftReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.errmsg[0] = "Bad Request: STARGIFT_INVALID"
	tc.code[0] = 400
	tc.whattocheck[errr] = struct{}{}
	if err = tc.msg.WriteUserID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.gift.WriteID(giftid); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteGift(tc.gift); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "gift", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "gift", tc.get.Response())
	tc.checkResponse(0)
}

func sendgift(i int, tc *testcase, t *testing.T) {
	var err error
	if i < 3 {
		if err = tc.msg.WriteParseMode(parsemode[i]); err != nil {
			t.Fatal(err)
		}
	} else if i == 3 {
		if err = tc.msg.WriteEntities(entities); err != nil {
			t.Fatal(err)
		}
	}
	if err = tc.msg.WriteString(textformsg); err != nil {
		t.Fatal(err)
	}
	if err = tc.gift.WritePayForUpgrade(); err != nil {
		t.Fatal(err)
	}
	if err = tc.gift.WriteID(giftid); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteUserID(chatid); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteGift(tc.gift); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "gift", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "gift", tc.get.Response())
	tc.checkResponse(0)
	tc.timetochange <- struct{}{}
}

func giftAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			tc := new(testcase)
			tc.init()
			tc.errmsg[0] = "Bad Request: STARGIFT_INVALID"
			tc.code[0] = 400
			tc.whattocheck[errr] = struct{}{}
			defer func() { tc.workdone <- struct{}{} }()
			go tc.changeToken(nil, nil, t)
			sendgift(i, tc, t)
		})
	}
}

func TestSendGift(t *testing.T) {
	t.Run("Req", giftReq)
	t.Run("All", giftAll)
}

func payReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.defaultSet()
	if err = tc.pay.WriteTitle(invtitle); err != nil {
		t.Fatal(err)
	}
	if err = tc.pay.WriteDescription(invdescrip); err != nil {
		t.Fatal(err)
	}
	if err = tc.pay.WritePayload(invpayload); err != nil {
		t.Fatal(err)
	}
	if err = tc.pay.WriteCurrency(starsName); err != nil {
		t.Fatal(err)
	}
	if err = tc.pay.WritePrices(prices); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WritePayment(tc.pay); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "payment", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "payment", tc.get.Response())
	tc.checkResponse(0)
	tc.checkResponse(0)
}

func payAll(t *testing.T) {
	var err error
	tt := new(testcase)
	tt.init()
	tt.defaultSet()
	if err = tt.pay.WriteTitle(invtitle); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WriteDescription(invdescrip); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WritePayload(invpayload); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WriteCurrency(starsName); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WriteProviderToken(""); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WritePrices(prices); err != nil {
		t.Fatal(err)
	}
	// tt.pay.WriteMaxTipAmount(15)
	// tt.pay.WriteSuggestedTipAmounts([]int{5})
	if err = tt.pay.WriteStartParameter("stars_payment"); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WriteProviderData("{}"); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WritePhotoUrl(photodata[tt.bs.Token][2]); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WritePhotoSize(12); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WritePhotoWidth(3); err != nil {
		t.Fatal(err)
	}
	if err = tt.pay.WritePhotoHeight(6); err != nil {
		t.Fatal(err)
	}
	// tt.pay.WriteNeedName()
	// tt.pay.WriteNeedPhoneNumber()
	// tt.pay.WriteNeedEmail()
	// tt.pay.WriteNeedShippingAddress()
	// tt.pay.WriteSendPhoneNumberToProvider()
	// tt.pay.WriteSendEmailToProvider()
	// tt.pay.WriteIsFlexible()
	tt.msg.WriteDisableNotification()
	tt.msg.WriteProtectContent()
	tt.msg.WriteMessageEffectID(msgEffect)
	tt.msg.WriteReplyParameters(tt.getReplyPrm())
	if err = tt.msg.WritePayment(tt.pay); err != nil {
		t.Fatal(err)
	}
	// kb[0](tt)
	if err = tt.bs.Send(tt.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "payment", tt.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "payment", tt.get.Response())
	tt.checkResponse(0)
}

func TestSendInvoice(t *testing.T) {
	t.Run("Req", payReq)
	t.Run("All", payAll)
}

func gameReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.bs.Token = testbotdata.TestsInGroup2_bot
	if err = tc.msg.WriteGame(gameshortname); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "game", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "game", tc.get.Response())
	tc.checkResponse(0)
}

func gameAll(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	tc.bs.Token = testbotdata.TestsInGroup2_bot
	if err = tc.msg.WriteGame(gameshortname); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteMessageEffectID(msgEffect); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.WriteReplyParameters(tc.getReplyPrm()); err != nil {
		t.Fatal(err)
	}
	if err = tc.bs.Send(tc.msg, chatid); err != nil && err.Error() != "22" {
		t.Fatal(err)
	}
	t.Logf("[TEST:%s] Request:\n%s", "game", tc.get.Request())
	t.Logf("[TEST:%s] Response:\n%s", "game", tc.get.Response())
	tc.checkResponse(0)
}

func TestSendGame(t *testing.T) {
	t.Run("Req", gameReq)
	t.Run("All", gameAll)
}
