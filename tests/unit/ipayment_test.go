package unit

import (
	"bytes"
	"strings"
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

type paymentT struct {
	name          string
	str           string
	integer       int
	prices        []*types.LabeledPrice
	arrInt        []int
	shOpt         []*types.ShippingOption
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type payTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	inputPrice    [][]*types.LabeledPrice
	inputArrInt   [][]int
	InputShOpt    [][]*types.ShippingOption
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(pay formatter.IPayment, p *paymentT, i int)
}

func (paytc *payTestContainer) title(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteTitle
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) description(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteDescription
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) payload(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePayload
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) providerToken(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteProviderToken
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) currency(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteCurrency
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) pricesData(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePrices
	p.prices = paytc.inputPrice[i]
}

func (paytc *payTestContainer) subscriptionPeriod(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSubscriptionPeriod
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) maxTipAmount(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteMaxTipAmount
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) suggestedTipAmounts(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSuggestedTipAmounts
	p.arrInt = paytc.inputArrInt[i]
}

func (paytc *payTestContainer) startParameter(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteStartParameter
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) providerData(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteProviderData
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) photoUrl(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoUrl
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) photoSize(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoSize
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) photoWidth(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoWidth
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) photoHeight(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePhotoHeight
	p.integer = paytc.inputInt[i]
}

func (paytc *payTestContainer) needName(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedName
}

func (paytc *payTestContainer) needPhoneNumber(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedPhoneNumber
}

func (paytc *payTestContainer) needEmail(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedEmail
}

func (paytc *payTestContainer) needShippingAddress(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteNeedShippingAddress
}

func (paytc *payTestContainer) sendPhoneNumberToProvider(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSendPhoneNumberToProvider
}

func (paytc *payTestContainer) sendEmailToProvider(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteSendEmailToProvider
}

func (paytc *payTestContainer) isFlexiable(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteIsFlexible
}

func (paytc *payTestContainer) shippingID(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteShippingID
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) ok(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteOK
}

func (paytc *payTestContainer) shippingOptions(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteShippingOptions
	p.shOpt = paytc.InputShOpt[i]
}

func (paytc *payTestContainer) errorMessage(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteErrorMessage
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) preCheckoutID(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WritePreCheckoutID
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) telegramPaymentChargeID(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteTelegramPaymentChargeID
	p.str = paytc.inputStr[i]
}

func (paytc *payTestContainer) isCanceled(pay formatter.IPayment, p *paymentT, i int) {
	p.testedFunc = pay.WriteIsCanceled
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
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
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
	paytc.inputPrice = [][]*types.LabeledPrice{{{}}, nil, {{}, nil, {}}, {{}}, {{}}}
	paytc.isExpectedErr = []bool{false, true, true, false, true}
	paytc.codeErr = []string{"", "20", "5", "", "10"}
	paytc.amount, paytc.until = 5, 3
	paytc.buildF = paytc.pricesData
}

func (paytc *payTestContainer) writeSubscriptionPeriod() {
	paytc.name = "(IPayment).WriteSubscriptionPeriod()"
	paytc.inputInt = []int{123123, -231, 0, formatter.Month, formatter.Month, formatter.Month}
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
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
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

func (paytc *payTestContainer) writeShippingID() {
	paytc.name = "(IPayment).WriteShippingID()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.shippingID
}

func (paytc *payTestContainer) writeOK() {
	paytc.name = "(IPayment).WriteOK()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.ok
}

func (paytc *payTestContainer) writeShippingOptions() {
	paytc.name = "(IPayment).WriteShippingOptions()"
	paytc.InputShOpt = [][]*types.ShippingOption{{}, nil, {{}, {}, nil}, {}, {}}
	paytc.isExpectedErr = []bool{false, true, true, false, true}
	paytc.codeErr = []string{"", "20", "5", "", "10"}
	paytc.amount, paytc.until = 5, 3
	paytc.buildF = paytc.shippingOptions
}

func (paytc *payTestContainer) writeErrorMessage() {
	paytc.name = "(IPayment).WriteErrorMessage()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.errorMessage
}

func (paytc *payTestContainer) writePreCheckoutID() {
	paytc.name = "(IPayment).WritePreCheckoutID()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.preCheckoutID
}

func (paytc *payTestContainer) writeTelegramPaymentChargeID() {
	paytc.name = "(IPayment).WriteTelegramPaymentChargeID()"
	paytc.inputStr = []string{"asdasdasdad", "", "ASKLJDAKSJLFJKLASJKLD", "AS:>DLKAWIO$*(AS)(ED()ASD)*()"}
	paytc.isExpectedErr = []bool{false, true, false, true}
	paytc.codeErr = []string{"", "20", "", "10"}
	paytc.amount, paytc.until = 4, 2
	paytc.buildF = paytc.telegramPaymentChargeID
}

func (paytc *payTestContainer) writeIsCanceled() {
	paytc.name = "(IPayment).writeIsCanceled()"
	paytc.isExpectedErr = []bool{false, false, true}
	paytc.codeErr = []string{"", "", "10"}
	paytc.amount, paytc.until = 3, 1
	paytc.buildF = paytc.isCanceled
}

func (pay *paymentT) startTest(part string, i int, t *testing.T) {
	switch f := pay.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, pay.name, pay.codeErr, pay.str, pay.isExpectedErr, i)
		checkError(f(pay.str), pay.isExpectedErr, pay.codeErr, t)
	case func(int) error:
		printTestLog(part, pay.name, pay.codeErr, pay.integer, pay.isExpectedErr, i)
		checkError(f(pay.integer), pay.isExpectedErr, pay.codeErr, t)
	case func([]*types.LabeledPrice) error:
		printTestLog(part, pay.name, pay.codeErr, pay.prices, pay.isExpectedErr, i)
		checkError(f(pay.prices), pay.isExpectedErr, pay.codeErr, t)
	case func([]*types.ShippingOption) error:
		printTestLog(part, pay.name, pay.codeErr, pay.shOpt, pay.isExpectedErr, i)
		checkError(f(pay.shOpt), pay.isExpectedErr, pay.codeErr, t)
	case func([]int) error:
		printTestLog(part, pay.name, pay.codeErr, pay.arrInt, pay.isExpectedErr, i)
		checkError(f(pay.arrInt), pay.isExpectedErr, pay.codeErr, t)
	case func() error:
		printTestLog(part, pay.name, pay.codeErr, true, pay.isExpectedErr, i)
		checkError(f(), pay.isExpectedErr, pay.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (paytc *payTestContainer) byDefault(i int) *paymentT {
	return &paymentT{name: paytc.name, isExpectedErr: paytc.isExpectedErr[i], codeErr: paytc.codeErr[i]}
}

func (paytc *payTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var pay formatter.IPayment
	a, b := make([]UnitTest, paytc.until), make([]UnitTest, paytc.amount-paytc.until)
	for i, j := 0, 0; i < paytc.amount; i++ {
		s := paytc.byDefault(i)
		if i < paytc.until {
			paytc.buildF(msg.NewPayment(), s, i)
			a[i] = s
		} else {
			if j%2 == 0 {
				pay = msg.NewPayment()
			}
			paytc.buildF(pay, s, i)
			b[j] = s
			j++
		}
	}
	return a, b
}

func (paytc *payTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	paycontainer, doublecontainer := paytc.createTestArrays(msg)
	check(paycontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestPaymentWriteTitle(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeTitle()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteDescription(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeDescription()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePayload(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePayload()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteProviderToken(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeProviderToken()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteCurrency(t *testing.T) {
	for key := range formatter.Currencies {
		paytc := new(payTestContainer)
		paytc.writeCurrency(key)
		msg := formatter.CreateEmpltyMessage()
		paytc.mainLogic(msg, t)
	}
}

func TestPaymentWritePrices(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePrices()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSubscriptionPeriod(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeSubscriptionPeriod()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteMaxTipAmount(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeMaxTipAmount()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSuggestedTipAmounts(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeSuggestedTipAmounts()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteStartParameter(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeStartParameter()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteProviderData(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeProviderData()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoUrl(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePhotoUrl()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoSize(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePhotoSize()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoWidth(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePhotoWidth()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePhotoHeight(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePhotoHeight()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedName(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeNeedName()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedPhoneNumber(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeNeedPhoneNumber()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedEmail(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeNeedEmail()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteNeedShippingAddress(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeNeedShippingAddress()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSendPhoneNumberToProvider(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeSendPhoneNumberToProvider()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteSendEmailToProvider(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeSendEmailToProvider()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteIsFlexible(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeIsFlexiable()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteShippingID(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeShippingID()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteOK(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeOK()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteShippingOptions(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeShippingOptions()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteErrorMessage(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeErrorMessage()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWritePreCheckoutID(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writePreCheckoutID()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteTelegramPaymentChargeID(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeTelegramPaymentChargeID()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}

func TestPaymentWriteIsCaneled(t *testing.T) {
	paytc := new(payTestContainer)
	paytc.writeIsCanceled()
	msg := formatter.CreateEmpltyMessage()
	paytc.mainLogic(msg, t)
}
