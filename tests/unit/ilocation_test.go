package unit

import (
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type locationT struct {
	name          string
	str           string
	integer       int
	float         float64
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type locTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	inputFloat    []float64
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT
}

func putWriteLocationLatitude(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, "", 0, loctc.inputFloat[i], loc.WriteLatitude, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationLongitude(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, "", 0, loctc.inputFloat[i], loc.WriteLongitude, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationHorizontalAccuracy(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, "", 0, loctc.inputFloat[i], loc.WriteHorizontalAccuracy, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationLivePeriod(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, "", loctc.inputInt[i], 0, loc.WriteLivePeriod, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationHeading(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, "", loctc.inputInt[i], 0, loc.WriteHeading, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationProximityAlertRadius(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, "", loctc.inputInt[i], 0, loc.WriteProximityAlertRadius, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationTitle(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, loctc.inputStr[i], 0, 0, loc.WriteTitle, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationAddress(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, loctc.inputStr[i], 0, 0, loc.WriteAddress, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationFoursquareID(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, loctc.inputStr[i], 0, 0, loc.WriteFoursquareID, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationFoursquareType(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, loctc.inputStr[i], 0, 0, loc.WriteFoursquareType, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationGooglePlaceID(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, loctc.inputStr[i], 0, 0, loc.WriteGooglePlaceID, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func putWriteLocationGooglePlaceType(loctc locTestContainer, loc fmtogram.ILocation, i int) *locationT {
	return &locationT{loctc.name, loctc.inputStr[i], 0, 0, loc.WriteGooglePlaceType, loctc.isExpectedErr[i], loctc.codeErr[i]}
}

func (loctc *locTestContainer) writeLatitude() {
	loctc.name = "(ILocation).WriteLatitude"
	loctc.inputFloat = []float64{90, -110, -91, 96, -77, 44}
	loctc.isExpectedErr = []bool{false, true, true, true, false, true}
	loctc.codeErr = []string{"", "20", "20", "20", "", "10"}
	loctc.amount, loctc.until = 6, 4
	loctc.buildF = putWriteLocationLatitude
}

func (loctc *locTestContainer) writeLongitude() {
	loctc.name = "(ILocation).WriteLongitude"
	loctc.inputFloat = []float64{90, -181, -200, 181, 180, -180}
	loctc.isExpectedErr = []bool{false, true, true, true, false, true}
	loctc.codeErr = []string{"", "20", "20", "20", "", "10"}
	loctc.amount, loctc.until = 6, 4
	loctc.buildF = putWriteLocationLongitude
}

func (loctc *locTestContainer) writeHorizontalAccuracy() {
	loctc.name = "(ILocation).WriteHorizontalAccuracy"
	loctc.inputFloat = []float64{90, -181, -1, 150000, 1500, 999}
	loctc.isExpectedErr = []bool{false, true, true, true, false, true}
	loctc.codeErr = []string{"", "20", "20", "20", "", "10"}
	loctc.amount, loctc.until = 6, 4
	loctc.buildF = putWriteLocationHorizontalAccuracy
}

func (loctc *locTestContainer) writeLivePeriod() {
	loctc.name = "(ILocation).WriteLivePeriod"
	loctc.inputInt = []int{60, 40, -1, 86401, 0x7FFFFFFF, 86400, 999}
	loctc.isExpectedErr = []bool{false, true, true, true, false, false, true}
	loctc.codeErr = []string{"", "20", "20", "20", "", "", "10"}
	loctc.amount, loctc.until = 7, 5
	loctc.buildF = putWriteLocationLivePeriod
}

func (loctc *locTestContainer) writeHeading() {
	loctc.name = "(ILocation).WriteHeading"
	loctc.inputInt = []int{1, 0, -1, 86401, 360, 20, 55}
	loctc.isExpectedErr = []bool{false, true, true, true, false, false, true}
	loctc.codeErr = []string{"", "20", "20", "20", "", "", "10"}
	loctc.amount, loctc.until = 7, 5
	loctc.buildF = putWriteLocationHeading
}

func (loctc *locTestContainer) writeProximityAlertRadius() {
	loctc.name = "(ILocation).WriteProximityAlertRadius"
	loctc.inputInt = []int{1, 0, -1, 100001, 360, 20, 55}
	loctc.isExpectedErr = []bool{false, true, true, true, false, false, true}
	loctc.codeErr = []string{"", "20", "20", "20", "", "", "10"}
	loctc.amount, loctc.until = 7, 5
	loctc.buildF = putWriteLocationProximityAlertRadius
}

func (loctc *locTestContainer) writeTitle() {
	loctc.name = "(ILocation).WriteTitle"
	loctc.inputStr = []string{"title", "", "TITLE IN CAPS", "Nothing special"}
	loctc.isExpectedErr = []bool{false, true, false, true}
	loctc.codeErr = []string{"", "20", "", "10"}
	loctc.amount, loctc.until = 4, 2
	loctc.buildF = putWriteLocationTitle
}

func (loctc *locTestContainer) writeAddress() {
	loctc.name = "(ILocation).WriteAddress"
	loctc.inputStr = []string{"address", "", "ulica Pushkina dom Kalatushkina", "~"}
	loctc.isExpectedErr = []bool{false, true, false, true}
	loctc.codeErr = []string{"", "20", "", "10"}
	loctc.amount, loctc.until = 4, 2
	loctc.buildF = putWriteLocationAddress
}

func (loctc *locTestContainer) writeFoursquareID() {
	loctc.name = "(ILocation).WriteFoursquareID"
	loctc.inputStr = []string{"something", "", "ulica Pushkina dom Kalatushkina", "~"}
	loctc.isExpectedErr = []bool{false, true, false, true}
	loctc.codeErr = []string{"", "20", "", "10"}
	loctc.amount, loctc.until = 4, 2
	loctc.buildF = putWriteLocationFoursquareID
}

func (loctc *locTestContainer) writeFoursquareType() {
	loctc.name = "(ILocation).WriteFoursquareType"
	loctc.inputStr = []string{"arts_entertainment/default", "", "kinda type", "not a type"}
	loctc.isExpectedErr = []bool{false, true, false, true}
	loctc.codeErr = []string{"", "20", "", "10"}
	loctc.amount, loctc.until = 4, 2
	loctc.buildF = putWriteLocationFoursquareType
}

func (loctc *locTestContainer) writeGooglePlaceID() {
	loctc.name = "(ILocation).WriteGooglePlaceID"
	loctc.inputStr = []string{"google id", "", "GOOGLE/AMAZON", "Not an ID"}
	loctc.isExpectedErr = []bool{false, true, false, true}
	loctc.codeErr = []string{"", "20", "", "10"}
	loctc.amount, loctc.until = 4, 2
	loctc.buildF = putWriteLocationGooglePlaceID
}

func (loctc *locTestContainer) writeGooglePlaceType() {
	loctc.name = "(ILocation).WriteGooglePlaceType"
	loctc.inputStr = []string{"googletype", "", "GOOGLE/AMAZON type", "Not a google type at all"}
	loctc.isExpectedErr = []bool{false, true, false, true}
	loctc.codeErr = []string{"", "20", "", "10"}
	loctc.amount, loctc.until = 4, 2
	loctc.buildF = putWriteLocationGooglePlaceType
}

func (loc *locationT) callStrF(testedF func(string) error, t *testing.T) {
	if !loc.isExpectedErr {
		if err := testedF(loc.str); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(loc.str); err.Error() != loc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (loc *locationT) callIntF(testedF func(int) error, t *testing.T) {
	if !loc.isExpectedErr {
		if err := testedF(loc.integer); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(loc.integer); err.Error() != loc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (loc *locationT) callFloat64F(testedF func(float64) error, t *testing.T) {
	if !loc.isExpectedErr {
		if err := testedF(loc.float); err != nil {
			t.Fatalf(errMsg, err)
		}
	} else {
		if err := testedF(loc.float); err.Error() != loc.codeErr {
			t.Fatalf(errMsg, err)
		}
	}
}

func (loc *locationT) startTest(part string, i int, t *testing.T) {
	switch f := loc.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, loc.name, loc.codeErr, loc.str, loc.isExpectedErr, i, t)
		loc.callStrF(f, t)
	case func(int) error:
		printTestLog(part, loc.name, loc.codeErr, loc.integer, loc.isExpectedErr, i, t)
		loc.callIntF(f, t)
	case func(float64) error:
		printTestLog(part, loc.name, loc.codeErr, loc.float, loc.isExpectedErr, i, t)
		loc.callFloat64F(f, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (loctc *locTestContainer) createTestArrays(msg *fmtogram.Message) ([]UnitTest, []UnitTest) {
	var loc fmtogram.ILocation
	a, b := make([]UnitTest, loctc.until), make([]UnitTest, loctc.amount-loctc.until)
	for i, j := 0, 0; i < loctc.amount; i++ {
		if i < loctc.until {
			loc = fmtogram.NewLocation()
			a[i] = loctc.buildF(*loctc, loc, i)
		} else {
			if j%2 == 0 {
				loc = fmtogram.NewLocation()
			}
			b[j] = loctc.buildF(*loctc, loc, i)
			j++
		}
	}
	return a, b
}

func mainLocationLogic(msg *fmtogram.Message, loctc locTestContainer, t *testing.T) {
	locationcontainer, doublecontainer := loctc.createTestArrays(msg)
	check(locationcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteLocationLatitude(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeLatitude()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationLongitude(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeLongitude()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationHorizontalAccuracy(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeHorizontalAccuracy()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationLivePeriod(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeLivePeriod()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationHeading(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeHeading()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationProximityAlertRadius(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeProximityAlertRadius()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationTitle(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeTitle()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationAddress(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeAddress()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationFoursquareID(t *testing.T) {
	locct := new(locTestContainer)
	locct.writeFoursquareID()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationFoursquareType(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeFoursquareType()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationGooglePlaceID(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeGooglePlaceID()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}

func TestWriteLocationGooglePlaceType(t *testing.T) {
	t.Parallel()
	locct := new(locTestContainer)
	locct.writeGooglePlaceType()
	msg := fmtogram.CreateEmpltyMessage()
	mainLocationLogic(msg, *locct, t)
}
