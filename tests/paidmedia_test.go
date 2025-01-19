package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func printPaidVideo(vds []formatter.IVideo, t *testing.T) {
	for _, vd := range vds {
		t.Log(vd.GetResponse())
	}
}

func paidVideoMax(vd formatter.IVideo, t *testing.T) {
	if err := vd.WriteThumbnailTelegram("AgACAgIAAxkDAAICJGcf5rBCBtGmJm-IRRgrYK2XeIjqAAJn7jEbe80AAUmH5LwgebRpFwEAAwIAA3MAAzYE"); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteWidth(66); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteHeight(11); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteDuration(55); err != nil {
		t.Fatal(err)
	}
	if err := vd.WriteSupportsStreaming(); err != nil {
		t.Fatal(err)
	}
}

func paidVideoStMax(msg *formatter.Message, vds []formatter.IVideo, t *testing.T) {
	for _, vd := range vds {
		if err := vd.WriteVideoStorage("media_test/musk.mp4"); err != nil {
			t.Fatal(err)
		}
		paidVideoMax(vd, t)
		if err := msg.AddVideo(vd); err != nil {
			t.Fatal(err)
		}
	}
}

func paidVideoTgMax(msg *formatter.Message, vds []formatter.IVideo, t *testing.T) {
	for _, vd := range vds {
		if err := vd.WriteVideoTelegram("BAACAgIAAxkDAAICW2cmSEhOkyaHOAABq1xHkIJ4eRwwxwAC1FsAAlq0MUlyV6gDqCe2NjYE"); err != nil {
			t.Fatal(err)
		}
		paidVideoMax(vd, t)
		if err := msg.AddVideo(vd); err != nil {
			t.Fatal(err)
		}
	}
}

func paidMediaReq(msg *formatter.Message, pr formatter.IParameters, t *testing.T) {
	if err := pr.WriteStarCount(4); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddParameters(pr); err != nil {
		t.Fatal(err)
	}
}

func paidMediaUnreq(msg *formatter.Message, pr formatter.IParameters, t *testing.T) {
	if err := pr.WriteStarCount(4); err != nil {
		t.Fatal(err)
	}
	if err := pr.WritePayload("payload"); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteCaption("SHALOM! Love from <b>Israel</b>!"); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteParseMode(types.HTML); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteShowCaptionAboveMedia(); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := pr.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddParameters(pr); err != nil {
		t.Fatal(err)
	}
}

func paidPhotoReqFieldsWithName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoReqFieldsWithName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoReqFieldsWithName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoReqFieldsWithName(t *testing.T) {
	t.Run("1x", paidPhotoReqFieldsWithName1x)
	t.Run("2x", paidPhotoReqFieldsWithName2x)
	t.Run("10x", paidPhotoReqFieldsWithName10x)
}

func paidPhotoReqFieldsWithoutName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoReqFieldsWithoutName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoReqFieldsWithoutName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoReqFieldsWithoutName(t *testing.T) {
	t.Run("1x", paidPhotoReqFieldsWithoutName1x)
	t.Run("2x", paidPhotoReqFieldsWithoutName2x)
	t.Run("10x", paidPhotoReqFieldsWithoutName10x)
}

func paidPhotoUnReqFieldsWithName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoUnReqFieldsWithName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createReplyKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoUnReqFieldsWithName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoUnReqFieldsWithName(t *testing.T) {
	t.Run("1x", paidPhotoUnReqFieldsWithName1x)
	t.Run("2x", paidPhotoUnReqFieldsWithName2x)
	t.Run("10x", paidPhotoUnReqFieldsWithName10x)
}

func paidPhotoUnReqFieldsWithoutName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoUnReqFieldsWithoutName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoUnReqFieldsWithoutName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto(),
		msg.NewPhoto(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidPhotoUnReqFieldsWithoutName(t *testing.T) {
	t.Run("1x", paidPhotoUnReqFieldsWithoutName1x)
	t.Run("2x", paidPhotoUnReqFieldsWithoutName2x)
	t.Run("10x", paidPhotoUnReqFieldsWithoutName10x)
}

func paidPhoto(t *testing.T) {
	t.Run("ReqFieldsWithName", paidPhotoReqFieldsWithName)
	t.Run("ReqFieldsWithoutName", paidPhotoReqFieldsWithoutName)
	t.Run("UnReqFieldsWithName", paidPhotoUnReqFieldsWithName)
	t.Run("UnReqFieldsWithoutName", paidPhotoUnReqFieldsWithoutName)
}

// Paid Video Tests
func paidVideoReqFieldsWithName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoReqFieldsWithName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoReqFieldsWithName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoReqFieldsWithName(t *testing.T) {
	t.Run("1x", paidVideoReqFieldsWithName1x)
	t.Run("2x", paidVideoReqFieldsWithName2x)
	t.Run("10x", paidVideoReqFieldsWithName10x)
}

func paidVideoReqFieldsWithoutName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoReqFieldsWithoutName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoReqFieldsWithoutName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoReqFieldsWithoutName(t *testing.T) {
	t.Run("1x", paidVideoReqFieldsWithoutName1x)
	t.Run("2x", paidVideoReqFieldsWithoutName2x)
	t.Run("10x", paidVideoReqFieldsWithoutName10x)
}

func paidVideoUnReqFieldsWithName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []formatter.IVideo{msg.NewVideo()}
	paidVideoStMax(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printPaidVideo(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoUnReqFieldsWithName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []formatter.IVideo{msg.NewVideo(), msg.NewVideo()}
	paidVideoStMax(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createReplyKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printPaidVideo(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoUnReqFieldsWithName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []formatter.IVideo{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	paidVideoStMax(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printPaidVideo(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoUnReqFieldsWithName(t *testing.T) {
	t.Run("1x", paidVideoUnReqFieldsWithName1x)
	t.Run("2x", paidVideoUnReqFieldsWithName2x)
	t.Run("10x", paidVideoUnReqFieldsWithName10x)
}

func paidVideoUnReqFieldsWithoutName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []formatter.IVideo{msg.NewVideo()}
	paidVideoStMax(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printPaidVideo(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoUnReqFieldsWithoutName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []formatter.IVideo{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	paidVideoStMax(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printPaidVideo(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoUnReqFieldsWithoutName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []formatter.IVideo{msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo(),
		msg.NewVideo(), msg.NewVideo()}
	paidVideoStMax(msg, holder, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printPaidVideo(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoUnReqFieldsWithoutName(t *testing.T) {
	t.Run("1x", paidVideoUnReqFieldsWithoutName1x)
	t.Run("2x", paidVideoUnReqFieldsWithoutName2x)
	t.Run("10x", paidVideoUnReqFieldsWithoutName10x)
}

func paidVideo(t *testing.T) {
	t.Run("ReqFieldsWithName", paidVideoReqFieldsWithName)
	t.Run("ReqFieldsWithoutName", paidVideoReqFieldsWithoutName)
	t.Run("UnReqFieldsWithName", paidVideoUnReqFieldsWithName)
	t.Run("UnReqFieldsWithoutName", paidVideoUnReqFieldsWithoutName)
}

// Photo and Video Tests
func paidVideoAndPhotoReqFieldsWithName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoReqFieldsWithName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto(), msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoReqFieldsWithName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoReqFieldsWithName(t *testing.T) {
	t.Run("1x", paidVideoAndPhotoReqFieldsWithName1x)
	t.Run("2x", paidVideoAndPhotoReqFieldsWithName2x)
	t.Run("10x", paidVideoAndPhotoReqFieldsWithName10x)
}

func paidVideoAndPhotoReqFieldsWithoutName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoReqFieldsWithoutName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto(), msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinSt(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoReqFieldsWithoutName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder := []interface{}{msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto(),
		msg.NewVideo(), msg.NewPhoto()}
	writeAndAddMediaMinTg(msg, holder, t)
	paidMediaReq(msg, pr, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoReqFieldsWithoutName(t *testing.T) {
	t.Run("1x", paidVideoAndPhotoReqFieldsWithoutName1x)
	t.Run("2x", paidVideoAndPhotoReqFieldsWithoutName2x)
	t.Run("10x", paidVideoAndPhotoReqFieldsWithoutName10x)
}

func paidVideoAndPhotoUnReqFieldsWithName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto()}
	holder2 := []formatter.IVideo{msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder1, t)
	paidVideoStMax(msg, holder2, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printPaidVideo(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoUnReqFieldsWithName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []formatter.IVideo{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder1, t)
	paidVideoTgMax(msg, holder2, t)
	paidMediaUnreq(msg, pr, t)
	createReplyKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printPaidVideo(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoUnReqFieldsWithName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []formatter.IVideo{msg.NewVideo(), msg.NewVideo(), msg.NewVideo(), msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinTg(msg, holder1, t)
	paidVideoTgMax(msg, holder2, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.PaidMedia); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printPaidVideo(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoUnReqFieldsWithName(t *testing.T) {
	t.Run("1x", paidVideoAndPhotoUnReqFieldsWithName1x)
	t.Run("2x", paidVideoAndPhotoUnReqFieldsWithName2x)
	t.Run("10x", paidVideoAndPhotoUnReqFieldsWithName10x)
}

func paidVideoAndPhotoUnReqFieldsWithoutName1x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto()}
	holder2 := []formatter.IVideo{msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder1, t)
	paidVideoStMax(msg, holder2, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printPaidVideo(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoUnReqFieldsWithoutName2x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []formatter.IVideo{msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinSt(msg, holder1, t)
	paidVideoTgMax(msg, holder2, t)
	paidMediaUnreq(msg, pr, t)
	createReplyKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printPaidVideo(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoUnReqFieldsWithoutName10x(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	pr := msg.NewParameters()
	ch := createChat(msg, t)
	holder1 := []interface{}{msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto(), msg.NewPhoto()}
	holder2 := []formatter.IVideo{msg.NewVideo(), msg.NewVideo(), msg.NewVideo(), msg.NewVideo(), msg.NewVideo()}
	writeAndAddMediaMinTg(msg, holder1, t)
	paidVideoTgMax(msg, holder2, t)
	paidMediaUnreq(msg, pr, t)
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	printResponse(holder1, t)
	printPaidVideo(holder2, t)
	t.Log(pr.GetResponse())
	t.Log(ch.GetResponse())
}

func paidVideoAndPhotoUnReqFieldsWithoutName(t *testing.T) {
	t.Run("1x", paidVideoAndPhotoUnReqFieldsWithoutName1x)
	t.Run("2x", paidVideoAndPhotoUnReqFieldsWithoutName2x)
	t.Run("10x", paidVideoAndPhotoUnReqFieldsWithoutName10x)
}

func paidPhotoAndVideo(t *testing.T) {
	t.Run("ReqFieldsWithName", paidVideoAndPhotoReqFieldsWithName)
	t.Run("ReqFieldsWithoutName", paidVideoAndPhotoReqFieldsWithoutName)
	t.Run("UnReqFieldsWithName", paidVideoAndPhotoUnReqFieldsWithName)
	t.Run("UnReqFieldsWithoutName", paidVideoAndPhotoUnReqFieldsWithoutName)
}

func paidmediaFunctional(t *testing.T) {
	t.Run("Photo", paidPhoto)
	t.Run("Video", paidVideo)
	t.Run("UnitedPxV", paidPhotoAndVideo)
}

func TestSendPaidMedia(t *testing.T) {
	t.Run("Functional", paidmediaFunctional)
}
