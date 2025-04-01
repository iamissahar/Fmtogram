package unit

import (
	"bytes"
	"strings"
	"testing"

	fmtogram "github.com/iamissahar/Fmtogram"
)

type paymentT struct {
	name          string
	str           string
	integer       int
	prices        []*fmtogram.LabeledPrice
	arrInt        []int
	shOpt         []*fmtogram.ShippingOption
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type payTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	inputPrice    [][]*fmtogram.LabeledPrice
	inputArrInt   [][]int
	InputShOpt    [][]*fmtogram.ShippingOption
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(pay fmtogram.IPayment, p *paymentT, i int)
}

func (paytc *payTestContainer) title(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteTitle
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) description(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteDescription
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) payload(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePayload
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) providerToken(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteProviderToken
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) currency(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteCurrency
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) pricesData(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePrices
	p.prices = paytc.inputPrice[i]
}

func (paytc *payTestContainer) subscriptionPeriod(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSubscriptionPeriod
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) maxTipAmount(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteMaxTipAmount
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) suggestedTipAmounts(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSuggestedTipAmounts
	p.arrInt = paytc.inputArrInt[i]
}

func (paytc *payTestContainer) startParameter(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteStartParameter
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) providerData(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteProviderData
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) photoUrl(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoUrl
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) photoSize(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoSize
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) photoWidth(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoWidth
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) photoHeight(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoHeight
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) needName(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedName
}

func (paytc *payTestContainer) needPhoneNumber(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedPhoneNumber
}

func (paytc *payTestContainer) needEmail(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedEmail
}

func (paytc *payTestContainer) needShippingAddress(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedShippingAddress
}

func (paytc *payTestContainer) sendPhoneNumberToProvider(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSendPhoneNumberToProvider
}

func (paytc *payTestContainer) sendEmailToProvider(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSendEmailToProvider
}

func (paytc *payTestContainer) isFlexiable(pay fmtogram.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteIsFlexible
}

func (paytc *payTestContainer) writeTitle() {
	paytc.name = "(IPayment).WriteTitle()"
	paytc.inputStr = []string{"asdasdasdad", "", "1", strings.Repeat("s", 32), strings.Repeat("s", 33), "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "", "20", "", "10"}
	paytc.amount, paytc.until = 7, 5
	paytc.buildF = paytc.title
}

func (paytc *payTestContainer) writeDescription() {
	paytc.name = "(IPayment).WriteDescription()"
	paytc.inputStr = []string{"asdasdasdad", "", "1", strings.Repeat("s", 255), strings.Repeat("s", 256), "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "", "20", "", "10"}
	paytc.amount, paytc.until = 7, 5
	paytc.buildF = paytc.description
}

func (paytc *payTestContainer) writePayload() {
	paytc.name = "(IPayment).WritePayload()"
	paytc.inputStr = []string{"asdasdasdad", "", string(bytes.Repeat([]byte{1}, 1)), string(bytes.Repeat([]byte{1}, 128)), string(bytes.Repeat([]byte{1}, 129)), "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "", "20", "", "10"}
	paytc.amount, paytc.until = 7, 5
	paytc.buildF = paytc.payload
}

func (paytc *payTestContainer) writeProviderToken() {
	paytc.name = "(IPayment).WriteProviderToken()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, false, false, true}
	paytc.codeErr = []string{"", "", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.providerToken
}

func (paytc *payTestContainer) writeCurrency(currency string) {
	paytc.name = "(IPayment).WriteCurrency()"
	paytc.inputStr = []string{"asdasdasdad", currency, "", currency, currency}
	paytc.isExpectedErr = []bool{true, false, true, false, true}
	paytc.codeErr = []string{"20", "", "20", "", "10"}
	paytc.amount, paytc.until = 5, 3
	paytc.buildF = paytc.currency
}

func (paytc *payTestContainer) writePrices() {
	paytc.name = "(IPayment).WritePrices()"
	paytc.inputPrice = [][]*fmtogram.LabeledPrice{{{}}, nil, {{}, nil, {}}, {{}}, {{}}}
	paytc.isExpectedErr = []bool{false, true, true, false, true}
	paytc.codeErr = []string{"", "20", "5", "", "10"}
	paytc.amount, paytc.until = 5, 3
	paytc.buildF = paytc.pricesData
}

func (paytc *payTestContainer) writeSubscriptionPeriod() {
	paytc.name = "(IPayment).WriteSubscriptionPeriod()"
	paytc.inputInt = []int{123123, -231, 0, fmtogram.Month, fmtogram.Month, fmtogram.Month}
	paytc.isExpectedErr = []bool{true, true, true, false, false, true}
	paytc.codeErr = []string{"20", "20", "20", "", "", "10"}
	paytc.amount, paytc.until = 6, 4
	paytc.buildF = paytc.subscriptionPeriod
}

func (paytc *payTestContainer) writeMaxTipAmount() {
	paytc.name = "(IPayment).WriteMaxTipAmount()"
	paytc.inputInt = []int{123123, -231, 0, 1235, 345}
	paytc.isExpectedErr = []bool{false, true, true, false, true}
	paytc.codeErr = []string{"", "20", "20", "", "10"}
	paytc.amount, paytc.until = 5, 3
	paytc.buildF = paytc.maxTipAmount
}

func (paytc *payTestContainer) writeSuggestedTipAmounts() {
	paytc.name = "(IPayment).WriteSuggestedTipAmounts()"
	paytc.inputArrInt = [][]int{{1235, 52, 46}, nil, {523, -123}, {11, 0}, {1235, 52, 46}, {1235, 52, 46}}
	paytc.isExpectedErr = []bool{false, true, true, true, false, true}
	paytc.codeErr = []string{"", "20", "5", "5", "", "10"}
	paytc.amount, paytc.until = 6, 4
	paytc.buildF = paytc.suggestedTipAmounts
}

func (paytc *payTestContainer) writeStartParameter() {
	paytc.name = "(IPayment).WriteStartParameter()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, false, false, true}
	paytc.codeErr = []string{"", "", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.startParameter
}

func (paytc *payTestContainer) writeProviderData() {
	paytc.name = "(IPayment).WriteProviderData()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.providerData
}

func (paytc *payTestContainer) writePhotoUrl() {
	paytc.name = "(IPayment).WritePhotoUrl()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.photoUrl
}

func (paytc *payTestContainer) writePhotoSize() {
	paytc.name = "(IPayment).WritePhotoSize()"
	paytc.inputInt = []int{131, 0, -21, 1, 123, 563}
	paytc.isExpectedErr = []bool{false, true, true, false, false, true}
	paytc.codeErr = []string{"", "20", "20", "", "", "10"}
	paytc.amount, paytc.until = 6, 4
	paytc.buildF = paytc.photoSize
}

func (paytc *payTestContainer) writePhotoWidth() {
	paytc.name = "(IPayment).WritePhotoWidth()"
	paytc.inputInt = []int{131, 0, -21, 1, 123, 563}
	paytc.isExpectedErr = []bool{false, true, true, false, false, true}
	paytc.codeErr = []string{"", "20", "20", "", "", "10"}
	paytc.amount, paytc.until = 6, 4
	paytc.buildF = paytc.photoWidth
}

func (paytc *payTestContainer) writePhotoHeight() {
	paytc.name = "(IPayment).WritePhotoHeight()"
	paytc.inputInt = []int{131, 0, -21, 1, 123, 563}
	paytc.isExpectedErr = []bool{false, true, true, false, false, true}
	paytc.codeErr = []string{"", "20", "20", "", "", "10"}
	paytc.amount, paytc.until = 6, 4
	paytc.buildF = paytc.photoHeight
}

func (paytc *payTestContainer) writeNeedName() {
	paytc.name = "(IPayment).WriteNeedName()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.needName
}

func (paytc *payTestContainer) writeNeedPhoneNumber() {
	paytc.name = "(IPayment).WriteNeedPhoneNumber()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.needPhoneNumber
}

func (paytc *payTestContainer) writeNeedEmail() {
	paytc.name = "(IPayment).WriteNeedEmail()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.needEmail
}

func (paytc *payTestContainer) writeNeedShippingAddress() {
	paytc.name = "(IPayment).WriteNeedShippingAddress()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.needShippingAddress
}

func (paytc *payTestContainer) writeSendPhoneNumberToProvider() {
	paytc.name = "(IPayment).WriteSendPhoneNumberToProvider()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.sendPhoneNumberToProvider
}

func (paytc *payTestContainer) writeSendEmailToProvider() {
	paytc.name = "(IPayment).WriteSendEmailToProvider()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.sendEmailToProvider
}

func (paytc *payTestContainer) writeIsFlexiable() {
	paytc.name = "(IPayment).WriteIsFlexiable()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.isFlexiable
}

func (pay *paymentT) startTest(part string, i int, t *testing.T) {
	switch f := pay.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, pay.name, pay.codeErr, pay.str, pay.isExpectedErr, i, t)
		checkError(f(pay.str), pay.isExpectedErr, pay.codeErr, t)
	case func(int) error:
		printTestLog(part, pay.name, pay.codeErr, pay.integer, pay.isExpectedErr, i, t)
		checkError(f(pay.integer), pay.isExpectedErr, pay.codeErr, t)
	case func([]*fmtogram.LabeledPrice) error:
		printTestLog(part, pay.name, pay.codeErr, pay.prices, pay.isExpectedErr, i, t)
		checkError(f(pay.prices), pay.isExpectedErr, pay.codeErr, t)
	case func([]*fmtogram.ShippingOption) error:
		printTestLog(part, pay.name, pay.codeErr, pay.shOpt, pay.isExpectedErr, i, t)
		checkError(f(pay.shOpt), pay.isExpectedErr, pay.codeErr, t)
	case func([]int) error:
		printTestLog(part, pay.name, pay.codeErr, pay.arrInt, pay.isExpectedErr, i, t)
		checkError(f(pay.arrInt), pay.isExpectedErr, pay.codeErr, t)
	case func() error:
		printTestLog(part, pay.name, pay.codeErr, true, pay.isExpectedErr, i, t)
		checkError(f(), pay.isExpectedErr, pay.codeErr, t)
	case func(bool) error:
		printTestLog(part, pay.name, pay.codeErr, false, pay.isExpectedErr, i, t)
		checkError(f(false), pay.isExpectedErr, pay.codeErr, t)
	case func():

	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (paytc *payTestContainer) byDefault(i int) *paymentT {
	return &paymentT{name: paytc.name, isExpectedErr: paytc.isExpectedErr[i], codeErr: paytc.codeErr[i]}
}

func (paytc *payTestContainer) createTestArrays(msg *fmtogram.Message) ([]UnitTest, []UnitTest) {
	var pay fmtogram.IPayment
	a, b := make([]UnitTest, paytc.until), make([]UnitTest, paytc.amount-paytc.until)
	for i, j := 0, 0; i < paytc.amount; i++ {
		s := paytc.byDefault(i)
		if i < paytc.until {
			paytc.buildF(fmtogram.NewPayment(), s, i)
			a[i] = s
		} else {
			if j%2 == 0 {
				pay = fmtogram.NewPayment()
			}
			paytc.buildF(pay, s, i)
			b[j] = s
			j++
		}
	}
	return a, b
}

func (paytc *payTestContainer) mainLogic(msg *fmtogram.Message, t *testing.T) {
	paycontainer, doublecontainer := paytc.createTestArrays(msg)
	check(paycontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestPaymentWriteTitle(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeTitle()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteDescription(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeDescription()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePayload(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writePayload()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteProviderToken(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeProviderToken()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteCurrency(t *testing.T) {
	t.Parallel()
	for key := range fmtogram.Currencies {
		paytc := new(payTestContainer)
		paytc.writeCurrency(key)
		msg := fmtogram.CreateEmpltyMessage()
		paytc.mainLogic(msg, t)
	}
}

func TestPaymentWritePrices(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writePrices()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSubscriptionPeriod(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeSubscriptionPeriod()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteMaxTipAmount(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeMaxTipAmount()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSuggestedTipAmounts(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeSuggestedTipAmounts()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteStartParameter(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeStartParameter()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteProviderData(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeProviderData()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoUrl(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writePhotoUrl()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoSize(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writePhotoSize()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoWidth(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writePhotoWidth()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoHeight(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writePhotoHeight()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedName(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeNeedName()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedPhoneNumber(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeNeedPhoneNumber()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedEmail(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeNeedEmail()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedShippingAddress(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeNeedShippingAddress()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSendPhoneNumberToProvider(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeSendPhoneNumberToProvider()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSendEmailToProvider(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeSendEmailToProvider()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteIsFlexible(t *testing.T) {
	t.Parallel()
	paytc := new(payTestContainer)
	paytc.writeIsFlexiable()
	msg := fmtogram.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}
