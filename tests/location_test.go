package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func writeVenueReqFields(loc formatter.ILocation, msg *formatter.Message, t *testing.T) {
	if err := loc.WriteLatitude(56.5132); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2350); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteTitle("HATA S KRAIU"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteAddress("улица Пушкина, дом Калатушкина"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err != nil {
		t.Fatal(err)
	}
}

func writeVenueUnreqFields(loc formatter.ILocation, msg *formatter.Message, t *testing.T) {
	if err := loc.WriteLatitude(56.5132); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2350); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteTitle("HATA S KRAIU"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteAddress("улица Пушкина, дом Калатушкина"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteGooglePlaceID("12334"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteGooglePlaceType("cafe"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteFoursquareID("0x151d4c90dbf22a73:0x59df6666a223fbd2"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteFoursquareType("AKSDALK:SDLKA"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err != nil {
		t.Fatal(err)
	}
}

func writeLocReqFields(loc formatter.ILocation, msg *formatter.Message, t *testing.T) {
	if err := loc.WriteLatitude(56.5132); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2350); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err != nil {
		t.Fatal(err)
	}
}

func writeLocUnreqFields(loc formatter.ILocation, msg *formatter.Message, t *testing.T) {
	if err := loc.WriteLatitude(56.5132); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2350); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteHeading(54); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteHorizontalAccuracy(90); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLivePeriod(60); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteProximityAlertRadius(656); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err != nil {
		t.Fatal(err)
	}
}

func locationReqFieldsWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	writeLocReqFields(loc, msg, t)
	if err := msg.AddMethod(methods.Location); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
}

func locationReqFieldsWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	writeLocReqFields(loc, msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
}

func locationReqFields(t *testing.T) {
	t.Run("WithName", locationReqFieldsWithName)
	t.Run("WithoutName", locationReqFieldsWithoutName)
}

func locationUnreqFieldsWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	param := msg.NewMessage()
	writeLocUnreqFields(loc, msg, t)
	if err := param.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteMessageEffectID("507"); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	createReplyKeyboard(msg, t)
	if err := msg.AddMethod(methods.Location); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
	t.Log(param.GetResponse())
}

func locationUnreqFieldsWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	param := msg.NewMessage()
	writeLocUnreqFields(loc, msg, t)
	if err := param.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteMessageEffectID("507"); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
	t.Log(param.GetResponse())
}

func locationUnreqFields(t *testing.T) {
	t.Run("WithName", locationUnreqFieldsWithName)
	t.Run("WithoutName", locationUnreqFieldsWithoutName)
}

func locationFunctionalSendLocation(t *testing.T) {
	t.Run("ReqFileds", locationReqFields)
	t.Run("UnreqFields", locationUnreqFields)
}

func venueReqFieldsWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	writeVenueReqFields(loc, msg, t)
	if err := msg.AddMethod(methods.Venue); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
}

func venueReqFieldsWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	writeVenueReqFields(loc, msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
}

func venueReqFields(t *testing.T) {
	t.Run("WithName", venueReqFieldsWithName)
	t.Run("WithoutName", venueReqFieldsWithoutName)
}

func venueUnreqFieldsWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	param := msg.NewMessage()
	writeVenueUnreqFields(loc, msg, t)
	if err := param.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteMessageEffectID("507"); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.Venue); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
	t.Log(param.GetResponse())
}

func venueUnreqFieldsWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	ch := createChat(msg, t)
	loc := msg.NewLocation()
	param := msg.NewMessage()
	writeVenueUnreqFields(loc, msg, t)
	if err := param.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteMessageEffectID("507"); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteReplyParameters(&types.ReplyParameters{MessageID: 532, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	createInlineKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(ch.GetResponse())
	t.Log(loc.GetResponse())
	t.Log(param.GetResponse())
}

func venueUnreqFields(t *testing.T) {
	t.Run("WithName", venueUnreqFieldsWithName)
	t.Run("WithoutName", venueUnreqFieldsWithoutName)
}

func locationFunctionalSendVenue(t *testing.T) {
	t.Run("ReqFileds", venueReqFields)
	t.Run("UnreqFields", venueUnreqFields)
}

func locationFunctional(t *testing.T) {
	t.Run("sendLocation", locationFunctionalSendLocation)
	t.Run("sendVenue", locationFunctionalSendVenue)
}

func locationLatOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLatitude(77.2315); err != nil {
		t.Fatal(err)
	}
}

func locationLat10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLatitude(77.2315); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLatitude(77.2315); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationLat20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLatitude(91.35446); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationLat(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationLatOk)
	t.Run("Code10", locationLat10)
	t.Run("Code20", locationLat20)
}

func locationLongOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLongitude(101.2315); err != nil {
		t.Fatal(err)
	}
}

func locationLong10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLongitude(101.2315); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2315); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationLong20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLongitude(-181.2315); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationLong(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationLongOk)
	t.Run("Code10", locationLong10)
	t.Run("Code20", locationLong20)
}

func locationHorAccOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteHorizontalAccuracy(556); err != nil {
		t.Fatal(err)
	}
}

func locationHorAcc10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteHorizontalAccuracy(556); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteHorizontalAccuracy(556); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationHorAcc20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteHorizontalAccuracy(1700); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationHorAcc(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationHorAccOk)
	t.Run("Code10", locationHorAcc10)
	t.Run("Code20", locationHorAcc20)
}

func locationLivePeriodOkNum(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLivePeriod(4563); err != nil {
		t.Fatal(err)
	}
}

func locationLivePeriodOkFlag(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLivePeriod(0x7FFFFFFF); err != nil {
		t.Fatal(err)
	}
}

func locationLivePeriodOk(t *testing.T) {
	t.Run("60-86400", locationLivePeriodOkNum)
	t.Run("0x7FFFFFFF", locationLivePeriodOkFlag)
}

func locationLivePeriod10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLivePeriod(4563); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLivePeriod(0x7FFFFFFF); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationLivePeriod20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLivePeriod(10); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationLivePeriod(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationLivePeriodOk)
	t.Run("Code10", locationLivePeriod10)
	t.Run("Code20", locationLivePeriod20)
}

func locationHeadingOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteHeading(5); err != nil {
		t.Fatal(err)
	}
}

func locationHeading10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteHeading(5); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteHeading(57); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationHeading20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteHeading(-21); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationHeading(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationHeadingOk)
	t.Run("Code10", locationHeading10)
	t.Run("Code20", locationHeading20)
}

func locationProxAlRadOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteProximityAlertRadius(100000); err != nil {
		t.Fatal(err)
	}
}

func locationProxAlRad10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteProximityAlertRadius(100000); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteProximityAlertRadius(10); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationProxAlRad20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteProximityAlertRadius(100001); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationProxAlRad(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationProxAlRadOk)
	t.Run("Code10", locationProxAlRad10)
	t.Run("Code20", locationProxAlRad20)
}

func locationTitleOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteTitle("HATA S KRAIU"); err != nil {
		t.Fatal(err)
	}
}

func locationTitle10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteTitle("HATA S KRAIU"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteTitle("HATA S KRAIU"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationTitle20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteTitle(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationTitle(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationTitleOk)
	t.Run("Code10", locationTitle10)
	t.Run("Code20", locationTitle20)
}

func locationAddressOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteAddress("улица Пушкина, дом Калатушкина"); err != nil {
		t.Fatal(err)
	}
}

func locationAddress10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteAddress("улица Пушкина, дом Калатушкина"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteAddress("улица Пушкина, дом Калатушкина"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationAddress20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteAddress(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationAddress(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationAddressOk)
	t.Run("Code10", locationAddress10)
	t.Run("Code20", locationAddress20)
}

func locationFoursquareIDOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteFoursquareID("AK:SLDKLASKLD"); err != nil {
		t.Fatal(err)
	}
}

func locationFoursquareID10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteFoursquareID("AK:SLDKLASKLD"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteFoursquareID("AK:SLDKLASKLD"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationFoursquareID20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteFoursquareID(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationFoursquareID(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationFoursquareIDOk)
	t.Run("Code10", locationFoursquareID10)
	t.Run("Code20", locationFoursquareID20)
}

func locationFoursquareTypeOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteFoursquareType("AS:KLDKL:ASD:KOLAS"); err != nil {
		t.Fatal(err)
	}
}

func locationFoursquareType10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteFoursquareType("AS:KLDKL:ASD:KOLAS"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteFoursquareType("AS:KLDKL:ASD:KOLAS"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationFoursquareType20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteFoursquareType(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationFoursquareType(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationFoursquareTypeOk)
	t.Run("Code10", locationFoursquareType10)
	t.Run("Code20", locationFoursquareType20)
}

func locationGooglePlaceIDOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteGooglePlaceID("12541234"); err != nil {
		t.Fatal(err)
	}
}

func locationGooglePlaceID10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteGooglePlaceID("12541234"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteGooglePlaceID("12541234"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationGooglePlaceID20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteGooglePlaceID(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationGooglePlaceID(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationGooglePlaceIDOk)
	t.Run("Code10", locationGooglePlaceID10)
	t.Run("Code20", locationGooglePlaceID20)
}

func locationGooglePlaceTypeOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteGooglePlaceType("car"); err != nil {
		t.Fatal(err)
	}
}

func locationGooglePlaceType10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteGooglePlaceType("car"); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteGooglePlaceType("car"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationGooglePlaceType20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteGooglePlaceType(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationGooglePlaceType(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationGooglePlaceTypeOk)
	t.Run("Code10", locationGooglePlaceType10)
	t.Run("Code20", locationGooglePlaceType20)
}

func locationAddOk(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLatitude(56.5132); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2350); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err != nil {
		t.Fatal(err)
	}
}

func locationAdd10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := loc.WriteLatitude(56.5132); err != nil {
		t.Fatal(err)
	}
	if err := loc.WriteLongitude(101.2350); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddLocation(loc); err.Error() != code10 {
		t.Fatal(err)
	}
}

func locationAdd20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	if err := msg.AddLocation(nil); err.Error() != code20 {
		t.Fatal(err)
	}
}

func locationAdd21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	loc := msg.NewLocation()
	if err := msg.AddLocation(loc); err.Error() != code21 {
		t.Fatal(err)
	}
}

func locationAddLocation(t *testing.T) {
	t.Parallel()
	t.Run("OK", locationAddOk)
	t.Run("Code10", locationAdd10)
	t.Run("Code20", locationAdd20)
	t.Run("Code21", locationAdd21)
}

func locationAddVenue(t *testing.T) {

}

func locationAdd(t *testing.T) {
	t.Parallel()
	t.Run("sendLocation", locationAddLocation)
	t.Run("sendVenue", locationAddVenue)
}

func locationUnit(t *testing.T) {
	t.Parallel()
	t.Run("Longitude", locationLat)
	t.Run("Latitude", locationLong)
	t.Run("HorizontalAccuracy", locationHorAcc)
	t.Run("LivePeriod", locationLivePeriod)
	t.Run("Heading", locationHeading)
	t.Run("ProximityAlertRadius", locationProxAlRad)
	t.Run("Title", locationTitle)
	t.Run("Address", locationAddress)
	t.Run("FoursquareID", locationFoursquareID)
	t.Run("FoursquareType", locationFoursquareType)
	t.Run("GooglePlaceID", locationGooglePlaceID)
	t.Run("GooglePlaceType", locationGooglePlaceType)
	t.Run("Add", locationAdd)
}

func TestSendLocation(t *testing.T) {
	t.Parallel()
	t.Run("Functional", locationFunctional)
	t.Run("Unit", locationUnit)
}
