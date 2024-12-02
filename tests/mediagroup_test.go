package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

func writePhotoFields(ph formatter.IPhoto, t *testing.T) {
	if err := ph.WriteCaption("<b>SHalom!</b> YA VAS LUBLU"); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := ph.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func writeVideoFields(vd formatter.IVideo, t *testing.T) {
	if err := vd.WriteCaption("YA LUCHSHE <b>VSEH!</b>"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(11); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHasSpoiler(); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteParseMode(types.Markdown); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(55); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteThumbnailTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(100); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
}

func writeAudioFields(ad formatter.IAudio, t *testing.T) {
	if err := ad.WriteCaption("It's supposed to be an audio"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteParseMode(types.MarkdownV2); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteDuration(4); err != nil {
		t.Fatal(err)
	}
	if err := ad.WritePerformer("performer"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteTitle("Title is here! It's a sound!"); err != nil {
		t.Fatal(err)
	}
	if err := ad.WriteThumbnailTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
}

func writeDocFields(doc formatter.IDocument, t *testing.T) {
	if err := doc.WriteCaption("It's <b>just</b> a resume"); err != nil {
		t.Fatal(err)
	}
	if err := doc.WriteDisableContentTypeDetection(); err != nil {
		t.Fatal(err)
	}
	if err := doc.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := doc.WriteThumbnailTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
}

func writePhotoSt(ph formatter.IPhoto, t *testing.T) {
	if err := ph.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	writePhotoFields(ph, t)
}

func writePhotoTg(ph formatter.IPhoto, t *testing.T) {
	if err := ph.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
	writePhotoFields(ph, t)
}

func writeVideoSt(vd formatter.IVideo, t *testing.T) {
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	writeVideoFields(vd, t)
}

func writeVideoTg(vd formatter.IVideo, t *testing.T) {
	if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
		t.Fatal(err)
	}
	writeVideoFields(vd, t)
}

func writeAudioSt(ad formatter.IAudio, t *testing.T) {
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	writeAudioFields(ad, t)
}

func writeAudioTg(ad formatter.IAudio, t *testing.T) {
	if err := ad.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
		t.Fatal(err)
	}
	writeAudioFields(ad, t)
}

func writeDocumentSt(doc formatter.IDocument, t *testing.T) {
	if err := doc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	writeDocFields(doc, t)
}

func writeDocumenTg(doc formatter.IDocument, t *testing.T) {
	if err := doc.WriteDocumentTelegram("BQACAgIAAxUHZ04Lv9KGTAKVx03M3Meuj1-rrcMAAqBlAAKHvHFKKa-WTPa24802BA"); err != nil {
		t.Fatal(err)
	}
	writeDocFields(doc, t)
}

func writeAndAddMediaMinSt(msg *formatter.Message, storage []interface{}, t *testing.T) {
	for i := 0; i < 10 && i < len(storage); i++ {
		switch media := storage[i].(type) {
		case formatter.IPhoto:
			if err := media.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddPhoto(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IVideo:
			if err := media.WriteVideoStorage("media_test/musk.mp4"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddVideo(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IAudio:
			if err := media.WriteAudioStorage("media_test/sound.mp3"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddAudio(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IDocument:
			if err := media.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddDocument(media); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func writeAndAddMediaMinTg(msg *formatter.Message, storage []interface{}, t *testing.T) {
	for i := 0; i < 10 && i < len(storage); i++ {
		switch media := storage[i].(type) {
		case formatter.IPhoto:
			if err := media.WritePhotoTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddPhoto(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IVideo:
			if err := media.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddVideo(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IAudio:
			if err := media.WriteAudioTelegram("CQACAgIAAxkDAAICOGcf7hnLD4tEsb5uMKPBcxZywiG3AAImeQACe80AAUnK0NGMf0Nv2zYE"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddAudio(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IDocument:
			if err := media.WriteDocumentTelegram("BQACAgIAAxUHZ04Lv9KGTAKVx03M3Meuj1-rrcMAAqBlAAKHvHFKKa-WTPa24802BA"); err != nil {
				t.Fatal(err)
			}
			if err := msg.AddDocument(media); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func writeAndAddMediaMaxSt(msg *formatter.Message, storage []interface{}, t *testing.T) {
	for i := 0; i < 10 && i < len(storage); i++ {
		switch media := storage[i].(type) {
		case formatter.IPhoto:
			writePhotoSt(media, t)
			if err := msg.AddPhoto(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IVideo:
			writeVideoSt(media, t)
			if err := msg.AddVideo(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IAudio:
			writeAudioSt(media, t)
			if err := msg.AddAudio(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IDocument:
			writeDocumentSt(media, t)
			if err := msg.AddDocument(media); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func writeAndAddMediaMaxTg(msg *formatter.Message, storage []interface{}, t *testing.T) {
	for i := 0; i < 10 && i < len(storage); i++ {
		switch media := storage[i].(type) {
		case formatter.IPhoto:
			writePhotoTg(media, t)
			if err := msg.AddPhoto(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IVideo:
			writeVideoTg(media, t)
			if err := msg.AddVideo(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IAudio:
			writeAudioTg(media, t)
			if err := msg.AddAudio(media); err != nil {
				t.Fatal(err)
			}
		case formatter.IDocument:
			writeDocumenTg(media, t)
			if err := msg.AddDocument(media); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func writeAllMsgParams(pr formatter.IParameters, t *testing.T) {
	if err := pr.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
}

func printResponse(mediaHolder []interface{}, t *testing.T) {
	for _, val := range mediaHolder {
		switch media := val.(type) {
		case formatter.IPhoto:
			t.Log(media.GetResponse())
		case formatter.IVideo:
			t.Log(media.GetResponse())
		case formatter.IAudio:
			t.Log(media.GetResponse())
		case formatter.IDocument:
			t.Log(media.GetResponse())
		default:
			t.Fatal("unknown data type")
		}
	}
}

// Photo Tests
func photoSt2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func photoSt2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func photoSt2x(t *testing.T) {
	t.Run("Req", photoSt2xReq)
	t.Run("Unreq", photoSt2xUnreq)
}

func photoSt10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func photoSt10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func photoSt10x(t *testing.T) {
	t.Run("Req", photoSt10xReq)
	t.Run("Unreq", photoSt10xUnreq)
}

func photoTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func photoTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func photoTg2x(t *testing.T) {
	t.Run("Req", photoTg2xReq)
	t.Run("Unreq", photoTg2xUnreq)
}

func photoTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func photoTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func photoTg10x(t *testing.T) {
	t.Run("Req", photoTg10xReq)
	t.Run("Unreq", photoTg10xUnreq)
}

func photoStAndTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto()}
	holder2 := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder1, t)
	writeAndAddMediaMinSt(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func photoStAndTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder1 := []interface{}{msg.NewPhoto()}
	holder2 := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMaxTg(msg, holder1, t)
	writeAndAddMediaMaxSt(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func photoStAndTg2x(t *testing.T) {
	t.Run("Req", photoStAndTg2xReq)
	t.Run("Unreq", photoStAndTg2xUnreq)
}

func photoStAndTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []interface{}{msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder1, t)
	writeAndAddMediaMinSt(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func photoStAndTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder1 := []interface{}{msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []interface{}{msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxTg(msg, holder1, t)
	writeAndAddMediaMaxSt(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func photoStAndTg10x(t *testing.T) {
	t.Run("Req", photoStAndTg10xReq)
	t.Run("Unreq", photoStAndTg10xUnreq)
}

// Video Tests
func videoSt2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoSt2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoSt2x(t *testing.T) {
	t.Run("Req", videoSt2xReq)
	t.Run("Unreq", videoSt2xUnreq)
}

func videoSt10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoSt10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoSt10x(t *testing.T) {
	t.Run("Req", videoSt10xReq)
	t.Run("Unreq", videoSt10xUnreq)
}

func videoTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoTg2x(t *testing.T) {
	t.Run("Req", videoTg2xReq)
	t.Run("Unreq", videoTg2xUnreq)
}

func videoTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoTg10x(t *testing.T) {
	t.Run("Req", videoTg10xReq)
	t.Run("Unreq", videoTg10xUnreq)
}

func videoStAndTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	holder2 := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo()}
	writeAndAddMediaMinTg(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func videoStAndTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder1 := []interface{}{msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	holder2 := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo()}
	writeAndAddMediaMaxTg(msg, holder1, t)
	writeAndAddMediaMaxTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoStAndTg2x(t *testing.T) {
	t.Run("Req", videoStAndTg2xReq)
	t.Run("Unreq", videoStAndTg2xUnreq)
}

func videoStAndTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	holder2 := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func videoStAndTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder1 := []interface{}{msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	holder2 := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo()}
	writeAndAddMediaMaxSt(msg, holder1, t)
	writeAndAddMediaMaxTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoStAndTg10x(t *testing.T) {
	t.Run("Req", videoStAndTg10xReq)
	t.Run("Unreq", videoStAndTg10xUnreq)
}

func videoAndPhotoSt2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoAndPhotoSt2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoAndPhotoSt2x(t *testing.T) {
	t.Run("Req", videoAndPhotoSt2xReq)
	t.Run("Unreq", videoAndPhotoSt2xUnreq)
}

func videoAndPhotoSt10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(), msg.NewVideo(),
		msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoAndPhotoSt10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(), msg.NewVideo(),
		msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoAndPhotoSt10x(t *testing.T) {
	t.Run("Req", videoAndPhotoSt10xReq)
	t.Run("Unreq", videoAndPhotoSt10xUnreq)
}

func videoAndPhotoTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoAndPhotoTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoAndPhotoTg2x(t *testing.T) {
	t.Run("Req", videoAndPhotoTg2xReq)
	t.Run("Unreq", videoAndPhotoTg2xUnreq)
}

func videoAndPhotoTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(), msg.NewVideo(),
		msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func videoAndPhotoTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(), msg.NewVideo(),
		msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoAndPhotoTg10x(t *testing.T) {
	t.Run("Req", videoAndPhotoTg10xReq)
	t.Run("Unreq", videoAndPhotoTg10xUnreq)
}

func videoAndPhotoStAndTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewVideo()}
	holder2 := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func videoAndPhotoStAndTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder1 := []interface{}{msg.NewVideo()}
	holder2 := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMaxSt(msg, holder1, t)
	writeAndAddMediaMaxTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoAndPhotoStAndTg2x(t *testing.T) {
	t.Run("Req", videoAndPhotoStAndTg2xReq)
	t.Run("Unreq", videoAndPhotoStAndTg2xUnreq)
}

func videoAndPhotoStAndTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewVideo(), msg.NewPhoto(), msg.NewVideo(), msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []interface{}{msg.NewVideo(), msg.NewPhoto(), msg.NewVideo(), msg.NewPhoto(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func videoAndPhotoStAndTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	pr := msg.NewMessage()
	holder1 := []interface{}{msg.NewVideo(), msg.NewPhoto(), msg.NewVideo(), msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []interface{}{msg.NewVideo(), msg.NewPhoto(), msg.NewVideo(), msg.NewPhoto(), msg.NewVideo()}
	writeAndAddMediaMaxSt(msg, holder1, t)
	writeAndAddMediaMaxTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func videoAndPhotoStAndTg10x(t *testing.T) {
	t.Run("Req", videoAndPhotoStAndTg10xReq)
	t.Run("Unreq", videoAndPhotoStAndTg10xUnreq)
}

// Audio Tests
func audioSt2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func audioSt2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func audioSt2x(t *testing.T) {
	t.Run("Req", audioSt2xReq)
	t.Run("Unreq", audioSt2xUnreq)
}

func audioSt10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func audioSt10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func audioSt10x(t *testing.T) {
	t.Run("Req", audioSt10xReq)
	t.Run("Unreq", audioSt10xUnreq)
}

func audioTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func audioTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func audioTg2x(t *testing.T) {
	t.Run("Req", audioTg2xReq)
	t.Run("Unreq", audioTg2xUnreq)
}

func audioTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func audioTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func audioTg10x(t *testing.T) {
	t.Run("Req", audioTg10xReq)
	t.Run("Unreq", audioTg10xUnreq)
}

func audioStAndTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewAudio()}
	holder2 := []interface{}{msg.NewAudio()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func audioStAndTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewAudio()}
	holder2 := []interface{}{msg.NewAudio()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func audioStAndTg2x(t *testing.T) {
	t.Run("Req", audioStAndTg2xReq)
	t.Run("Unreq", audioStAndTg2xUnreq)
}

func audioStAndTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(), msg.NewAudio()}
	holder2 := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func audioStAndTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(), msg.NewAudio()}
	holder2 := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMaxSt(msg, holder1, t)
	writeAndAddMediaMaxTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func audioStAndTg10x(t *testing.T) {
	t.Run("Req", audioStAndTg10xReq)
	t.Run("Unreq", audioStAndTg10xUnreq)
}

// Document Tests
func documentSt2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func documentSt2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func documentSt2x(t *testing.T) {
	t.Run("Req", documentSt2xReq)
	t.Run("Unreq", documentSt2xUnreq)
}

func documentSt10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinSt(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func documentSt10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMaxSt(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func documentSt10x(t *testing.T) {
	t.Run("Req", documentSt10xReq)
	t.Run("Unreq", documentSt10xUnreq)
}

func documentTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func documentTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func documentTg2x(t *testing.T) {
	t.Run("Req", documentTg2xReq)
	t.Run("Unreq", documentTg2xUnreq)
}

func documentTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinTg(msg, holder, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(ch.GetResponse())
}

func documentTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMaxTg(msg, holder, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func documentTg10x(t *testing.T) {
	t.Run("Req", documentTg10xReq)
	t.Run("Unreq", documentTg10xUnreq)
}

func documentStAndTg2xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewDocument()}
	holder2 := []interface{}{msg.NewDocument()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func documentStAndTg2xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewDocument()}
	holder2 := []interface{}{msg.NewDocument()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func documentStAndTg2x(t *testing.T) {
	t.Run("Req", documentStAndTg2xReq)
	t.Run("Unreq", documentStAndTg2xUnreq)
}

func documentStAndTg10xReq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument()}
	holder2 := []interface{}{msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(ch.GetResponse())
}

func documentStAndTg10xUnreq(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewMessage()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument()}
	holder2 := []interface{}{msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinSt(msg, holder1, t)
	writeAndAddMediaMinTg(msg, holder2, t)
	writeAllMsgParams(pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printResponse(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func documentStAndTg10x(t *testing.T) {
	t.Run("Req", documentStAndTg10xReq)
	t.Run("Unreq", documentStAndTg10xUnreq)
}

func mediagroupFunctional(t *testing.T) {
	// t.Parallel()
	t.Run("2xPhotoStorage", photoSt2x)
	t.Run("10xPhotoStorage", photoSt10x)
	t.Run("2xPhotoTelegram", photoTg2x)
	t.Run("10xPhotoTelegram", photoTg10x)
	t.Run("2xPhotoTelegramAndStorage", photoStAndTg2x)
	t.Run("10xPhotoTelegramAndStorage", photoStAndTg10x)
	t.Run("2xVideoStorage", videoSt2x)
	t.Run("10xVideoStorage", videoSt10x)
	t.Run("2xVideoTelegram", videoTg2x)
	t.Run("10xVideoTelegram", videoTg10x)
	t.Run("2xVideoTelegramAndStorage", videoStAndTg2x)
	t.Run("10xVideoTelegramAndStorage", videoStAndTg10x)
	t.Run("2xVideoAndPhotoStorage", videoAndPhotoSt2x)
	t.Run("10xVideoAndPhotoStorage", videoAndPhotoSt10x)
	t.Run("2xVideoAndPhotoTelegram", videoAndPhotoTg2x)
	t.Run("V10xideoAndPhotoTelegram", videoAndPhotoTg10x)
	t.Run("2xVideoAndPhotoTelegramAndStorage", videoAndPhotoStAndTg2x)
	t.Run("10xVideoAndPhotoTelegramAndStorage", videoAndPhotoStAndTg10x)
	t.Run("2xAudioStorage", audioSt2x)
	t.Run("10xAudioStorage", audioSt10x)
	t.Run("2xAudioTelegram", audioTg2x)
	t.Run("10xAudioTelegram", audioTg10x)
	t.Run("2xAudioTelegramAndStorage", audioStAndTg2x)
	t.Run("10xAudioTelegramAndStorage", audioStAndTg10x)
	t.Run("2xDocumentxStorage", documentSt2x)
	t.Run("10xDocumentStorage", documentSt10x)
	t.Run("2xDocumentxTelegram", documentTg2x)
	t.Run("10xDocumentTelegram", documentTg10x)
	t.Run("2xDocumentTelegramAndStorage", documentStAndTg2x)
	t.Run("10xDocumentTelegramAndStorage", documentStAndTg10x)
}

func photo2xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewPhoto(), msg.NewPhoto()}, t)
}

func photo10xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}

	writeAndAddMediaMinSt(msg, storage, t)
}

func photoVideoX1OK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewPhoto(), msg.NewVideo()}, t)
}

func photoVideoX10OK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, storage, t)
}

func photoCode3(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, storage, t)
	switch photo := storage[10].(type) {
	case formatter.IPhoto:
		if err := photo.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
			t.Fatal(err)
		}
		if err := msg.AddPhoto(photo); err.Error() != code3 {
			t.Fatal(err)
		}
	}
}

func photoCode54(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	photo := msg.NewPhoto()
	if err := photo.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(photo); err != nil {
		t.Fatal(err)
	}
	ad := msg.NewAudio()
	if err := ad.WriteAudioInternet("AS:KLDAKLSDLK"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err.Error() != code54 {
		t.Fatal(err)
	}
}

func photoMedia(t *testing.T) {
	t.Parallel()
	t.Run("2xPhotoOK", photo2xOK)
	t.Run("10xPhotoOK", photo10xOK)
	t.Run("1xVideo&1xPhotoOK", photoVideoX1OK)
	t.Run("5xVideo&5xPhotoOK", photoVideoX10OK)
	t.Run("Code3", photoCode3)
	t.Run("Code54", photoCode54)
}

func video2xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewVideo(), msg.NewVideo()}, t)
}

func video10xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, storage, t)
}

func videoCode3(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo()}
	writeAndAddMediaMinSt(msg, storage, t)
	switch video := storage[10].(type) {
	case formatter.IVideo:
		if err := video.WriteVideoStorage("media_test/musk.mp4"); err != nil {
			t.Fatal(err)
		}
		if err := msg.AddVideo(video); err.Error() != code3 {
			t.Fatal(err)
		}
	}
}

func videoCode54(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	photo := msg.NewPhoto()
	if err := photo.WritePhotoStorage("media_test/photo1.jpg"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(photo); err != nil {
		t.Fatal(err)
	}
	dc := msg.NewDocument()
	if err := dc.WriteDocumentInternet("AS:KLDAKLSDLK"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(dc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func videoMeida(t *testing.T) {
	t.Parallel()
	t.Run("2xVideoOK", video2xOK)
	t.Run("10xVideoOK", video10xOK)
	t.Run("Code3", videoCode3)
	t.Run("Code54", videoCode54)
}

func audio2xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewAudio(), msg.NewAudio()}, t)
}

func audio10xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio()}
	writeAndAddMediaMinSt(msg, storage, t)
}

func audioCode3(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio(), msg.NewAudio(),
		msg.NewAudio()}
	writeAndAddMediaMinSt(msg, storage, t)
	switch ad := storage[10].(type) {
	case formatter.IAudio:
		if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
			t.Fatal(err)
		}
		if err := msg.AddAudio(ad); err.Error() != code3 {
			t.Fatal(err)
		}
	}
}

func audioCode54(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioStorage("media_test/sound.mp3"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	doc := msg.NewDocument()
	if err := doc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioMedia(t *testing.T) {
	t.Parallel()
	t.Run("2xAudioOK", audio2xOK)
	t.Run("10xAudioOK", audio10xOK)
	t.Run("Code3", audioCode3)
	t.Run("Code54", audioCode54)
}

func document2xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewDocument(), msg.NewDocument()}, t)
}

func document10xOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument()}
	writeAndAddMediaMinSt(msg, storage, t)
}

func documentCode3(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	storage := []interface{}{msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument(), msg.NewDocument(),
		msg.NewDocument()}
	writeAndAddMediaMinSt(msg, storage, t)
	switch doc := storage[10].(type) {
	case formatter.IDocument:
		if err := doc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
			t.Fatal(err)
		}
		if err := msg.AddDocument(doc); err.Error() != code3 {
			t.Fatal(err)
		}
	}
}

func documentCode54(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentStorage("media_test/Resume.pdf"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("ASLK:JDKLASDASSd"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentMedia(t *testing.T) {
	t.Parallel()
	t.Run("2xDocumentOK", document2xOK)
	t.Run("10xDocumentOK", document10xOK)
	t.Run("Code3", documentCode3)
	t.Run("Code54", documentCode54)
}

func photoXphoto(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewPhoto(), msg.NewPhoto()}, t)
}

func photoXvideo(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewPhoto(), msg.NewVideo()}, t)
}

func photoXaudio(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram(":AL:SD:LASDL:ASDL:A"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err.Error() != code54 {
		t.Fatal(err)
	}
}

func photoXdocument(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram(":AL:SD:LASDL:ASDL:A"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func photoXanimation(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram(":AL:SD:LASDL:ASDL:A"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAnimation(an); err.Error() != code54 {
		t.Fatal(err)
	}
}

func photoXvoice(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram(":AL:SD:LASDL:ASDL:A"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVoice(vc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func photoXvideonote(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram(":AL:SD:LASDL:ASDL:A"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err != nil {
		t.Fatal(err)
	}
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err.Error() != code54 {
		t.Fatal(err)
	}
}

func photoCompatibility(t *testing.T) {
	t.Parallel()
	t.Run("xPhoto", photoXphoto)
	t.Run("xVideo", photoXvideo)
	t.Run("xAudio", photoXaudio)
	t.Run("xDocument", photoXdocument)
	t.Run("xAnimation", photoXanimation)
	t.Run("xVoice", photoXvoice)
	t.Run("xVideoNote", photoXvideonote)
	// t.Run("xSticker")
}

func videoXvideo(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewVideo(), msg.NewVideo()}, t)
}

func videoXphoto(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewVideo(), msg.NewPhoto()}, t)
}

func videoXaudio(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("ASL:DLKSGHL:SDF:L12412"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err.Error() != code54 {
		t.Fatal(err)
	}
}
func videoXdocument(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("ASL:DLKSGHL:SDF:L12412"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func videoXanimation(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("ASL:DLKSGHL:SDF:L12412"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAnimation(an); err.Error() != code54 {
		t.Fatal(err)
	}
}

func videoXvoice(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("ASL:DLKSGHL:SDF:L12412"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVoice(vc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func videoXvideonote(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	vd := msg.NewVideo()
	if err := vd.WriteVideoTelegram("ASL:DLKSGHL:SDF:L12412"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err != nil {
		t.Fatal(err)
	}
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err.Error() != code54 {
		t.Fatal(err)
	}
}

func videoCompatibility(t *testing.T) {
	t.Parallel()
	t.Run("xVideo", videoXvideo)
	t.Run("xPhoto", videoXphoto)
	t.Run("xAudio", videoXaudio)
	t.Run("xDocument", videoXdocument)
	t.Run("xAnimation", videoXanimation)
	t.Run("xVoice", videoXvoice)
	t.Run("xVideoNote", videoXvideonote)
	// t.Run("xSticker")
}

func audioXaudio(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewAudio(), msg.NewAudio()}, t)
}

func audioXvideo(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("I!@#$JIKL!#@KLJ$!@#KLJ"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioXphoto(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("I!@#$JIKL!#@KLJ$!@#KLJ"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("KL:%@KL:$%KL:@#KL$"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioXdocument(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("I!@#$JIKL!#@KLJ$!@#KLJ"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioXanimation(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("I!@#$JIKL!#@KLJ$!@#KLJ"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAnimation(an); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioXvoice(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("I!@#$JIKL!#@KLJ$!@#KLJ"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVoice(vc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioXvideonote(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("I!@#$JIKL!#@KLJ$!@#KLJ"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err != nil {
		t.Fatal(err)
	}
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err.Error() != code54 {
		t.Fatal(err)
	}
}

func audioCompatibility(t *testing.T) {
	t.Parallel()
	t.Run("xAudio", audioXaudio)
	t.Run("xVideo", audioXvideo)
	t.Run("xPhoto", audioXphoto)
	t.Run("xDocument", audioXdocument)
	t.Run("xAnimation", audioXanimation)
	t.Run("xVoice", audioXvoice)
	t.Run("xVideoNote", audioXvideonote)
	// t.Run("xSticker")
}

func documentXdocument(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	writeAndAddMediaMinSt(msg, []interface{}{msg.NewDocument(), msg.NewDocument()}, t)
}

func documentXaudio(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	ad := msg.NewAudio()
	if err := ad.WriteAudioTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAudio(ad); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentXvideo(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	vd := msg.NewVideo()
	if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideo(vd); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentXphoto(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	ph := msg.NewPhoto()
	if err := ph.WritePhotoTelegram("KL:%@KL:$%KL:@#KL$"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddPhoto(ph); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentXanimation(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	an := msg.NewAnimation()
	if err := an.WriteAnimationTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddAnimation(an); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentXvoice(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	vc := msg.NewVoice()
	if err := vc.WriteVoiceTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVoice(vc); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentXvideonote(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	doc := msg.NewDocument()
	if err := doc.WriteDocumentTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddDocument(doc); err != nil {
		t.Fatal(err)
	}
	vdn := msg.NewVideoNote()
	if err := vdn.WriteVideoNoteTelegram("AS:LDA:LSD:LA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddVideoNote(vdn); err.Error() != code54 {
		t.Fatal(err)
	}
}

func documentCompatibility(t *testing.T) {
	t.Parallel()
	t.Run("xDocument", documentXdocument)
	t.Run("xAudio", documentXaudio)
	t.Run("xVideo", documentXvideo)
	t.Run("xPhoto", documentXphoto)
	t.Run("xAnimation", documentXanimation)
	t.Run("xVoice", documentXvoice)
	t.Run("xVideoNote", documentXvideonote)
	// t.Run("xSticker")
}

func mediagroupCompatibility(t *testing.T) {
	t.Parallel()
	t.Run("Photo", photoCompatibility)
	t.Run("Video", videoCompatibility)
	t.Run("Auido", audioCompatibility)
	t.Run("Document", documentCompatibility)
}

func mediagroupUnit(t *testing.T) {
	t.Parallel()
	t.Run("PhotoMedia", photoMedia)
	t.Run("VideoMedia", videoMeida)
	t.Run("AudioMedia", audioMedia)
	t.Run("DocumentMedia", documentMedia)
	t.Run("Compatibility", mediagroupCompatibility)
}

func TestSendMediaGroup(t *testing.T) {
	t.Parallel()
	t.Run("Functional", mediagroupFunctional)
	t.Run("Unit", mediagroupUnit)
}
