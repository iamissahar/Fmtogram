package unit

// import (
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/iamissahar/Fmtogram/formatter"
// )

// type linkT struct {
// 	name          string
// 	integer       int
// 	str           string
// 	date          time.Duration
// 	testedFunc    interface{}
// 	isExpectedErr bool
// 	codeErr       string
// }

// type linkTestContainer struct {
// 	name          string
// 	inputInt      []int
// 	inputStr      []string
// 	inputDates    []time.Duration
// 	isExpectedErr []bool
// 	codeErr       []string
// 	amount, until int
// 	buildF        func(prm formatter.ILink, l *linkT, i int)
// }

// func (linktc linkTestContainer) defaultLinkT(i int) *linkT {
// 	return &linkT{name: linktc.name, isExpectedErr: linktc.isExpectedErr[i], codeErr: linktc.codeErr[i]}
// }

// func (linktc *linkTestContainer) putLinkWriteName(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteName
// 	l.str = linktc.inputStr[i]
// }

// func (linktc *linkTestContainer) putLinkWriteExpireDate(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteExpireDate
// 	l.date = linktc.inputDates[i]
// }

// func (linktc *linkTestContainer) putLinkWriteMemberLimit(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteMemberLimit
// 	l.integer = linktc.inputInt[i]
// }

// func (linktc *linkTestContainer) putLinkWriteJoinRequest(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteJoinRequest
// }

// func (linktc *linkTestContainer) putLinkWriteIniveLink(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteInviteLink
// 	l.str = linktc.inputStr[i]
// }

// func (linktc *linkTestContainer) putLinkWriteSubscriptionPeriod(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteSubscriptionPeriod
// 	l.integer = linktc.inputInt[i]
// }

// func (linktc *linkTestContainer) putLinkWriteSubscriptionPrice(link formatter.ILink, l *linkT, i int) {
// 	l.testedFunc = link.WriteSubscriptionPrice
// 	l.integer = linktc.inputInt[i]
// }

// func (linktc *linkTestContainer) writeName() {
// 	linktc.name = "(ILink).WriteName()"
// 	linktc.inputStr = []string{"ASKDAK;LSD", "", "g", strings.Repeat("s", 32), strings.Repeat("s", 33), "dsa;ldl;a", "something"}
// 	linktc.isExpectedErr = []bool{false, true, false, false, true, false, true}
// 	linktc.codeErr = []string{"", "20", "", "", "20", "", "10"}
// 	linktc.amount, linktc.until = 7, 5
// 	linktc.buildF = linktc.putLinkWriteName
// }

// func (linktc *linkTestContainer) writeExpireDate() {
// 	linktc.name = "(ILink).WriteExpireDate()"
// 	linktc.inputDates = []time.Duration{time.Hour, 0, time.Microsecond, time.Minute, time.Nanosecond, time.Second}
// 	linktc.isExpectedErr = []bool{false, true, false, false, false, true}
// 	linktc.codeErr = []string{"", "20", "", "", "", "10"}
// 	linktc.amount, linktc.until = 6, 4
// 	linktc.buildF = linktc.putLinkWriteExpireDate
// }

// func (linktc *linkTestContainer) writeMemberLimit() {
// 	linktc.name = "(ILink).WriteMemberLimit()"
// 	linktc.inputInt = []int{123, 0, -1, 99999, 100000, 534, 88}
// 	linktc.isExpectedErr = []bool{false, true, true, false, true, false, true}
// 	linktc.codeErr = []string{"", "20", "20", "", "20", "", "10"}
// 	linktc.amount, linktc.until = 7, 5
// 	linktc.buildF = linktc.putLinkWriteMemberLimit
// }

// func (linktc *linkTestContainer) writeJoinRequest() {
// 	linktc.name = "(ILink).WriteJoinRequest()"
// 	linktc.isExpectedErr = []bool{false, false, true}
// 	linktc.codeErr = []string{"", "", "10"}
// 	linktc.amount, linktc.until = 3, 1
// 	linktc.buildF = linktc.putLinkWriteJoinRequest
// }

// func (linktc *linkTestContainer) writeIniveLink() {
// 	linktc.name = "(ILink).WriteIniveLink()"
// 	linktc.inputStr = []string{"ASKDAK;LSD", "", "dsa;ldl;a", "something"}
// 	linktc.isExpectedErr = []bool{false, true, false, true}
// 	linktc.codeErr = []string{"", "20", "", "10"}
// 	linktc.amount, linktc.until = 4, 2
// 	linktc.buildF = linktc.putLinkWriteIniveLink
// }

// func (linktc *linkTestContainer) writeSubscriptionPeriod() {
// 	linktc.name = "(ILink).WriteSubscriptionPeriod()"
// 	linktc.inputInt = []int{2592000, 0, 123151, -3439842012341, 2592000, 2592000}
// 	linktc.isExpectedErr = []bool{false, true, true, true, false, true}
// 	linktc.codeErr = []string{"", "20", "20", "20", "", "10"}
// 	linktc.amount, linktc.until = 6, 4
// 	linktc.buildF = linktc.putLinkWriteSubscriptionPeriod
// }

// func (linktc *linkTestContainer) writeSubscriptionPrice() {
// 	linktc.name = "(ILink).WriteSubscriptionPrice()"
// 	linktc.inputInt = []int{123, 0, -1, 1, 2500, 2501, 123, 744}
// 	linktc.isExpectedErr = []bool{false, true, true, false, false, true, false, true}
// 	linktc.codeErr = []string{"", "20", "20", "", "", "20", "", "10"}
// 	linktc.amount, linktc.until = 8, 6
// 	linktc.buildF = linktc.putLinkWriteSubscriptionPrice
// }

// func (link *linkT) startTest(part string, i int, t *testing.T) {
// 	switch f := link.testedFunc.(type) {
// 	case func(string) error:
// 		printTestLog(part, link.name, link.codeErr, link.str, link.isExpectedErr, i, t)
// 		checkError(f(link.str), link.isExpectedErr, link.codeErr, t)
// 	case func(int) error:
// 		printTestLog(part, link.name, link.codeErr, link.integer, link.isExpectedErr, i, t)
// 		checkError(f(link.integer), link.isExpectedErr, link.codeErr, t)
// 	case func(time.Duration) error:
// 		printTestLog(part, link.name, link.codeErr, link.date, link.isExpectedErr, i, t)
// 		checkError(f(link.date), link.isExpectedErr, link.codeErr, t)
// 	case func() error:
// 		printTestLog(part, link.name, link.codeErr, true, link.isExpectedErr, i, t)
// 		checkError(f(), link.isExpectedErr, link.codeErr, t)
// 	}
// }

// func (linktc *linkTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
// 	var l formatter.ILink
// 	a, b := make([]UnitTest, linktc.until), make([]UnitTest, linktc.amount-linktc.until)
// 	for i, j := 0, 0; i < linktc.amount; i++ {
// 		p := linktc.defaultLinkT(i)
// 		if i < linktc.until {
// 			linktc.buildF(msg.NewLink(), p, i)
// 			a[i] = p
// 		} else {
// 			if j%2 == 0 {
// 				l = msg.NewLink()
// 			}
// 			linktc.buildF(l, p, i)
// 			b[j] = p
// 			j++
// 		}
// 	}
// 	return a, b
// }

// func (linktc *linkTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
// 	linkcontainer, doublecontainer := linktc.createTestArrays(msg)
// 	t.Log(linkcontainer, doublecontainer)
// 	check(linkcontainer, t)
// 	doubleCheck(doublecontainer, t)
// }

// func TestLinkWriteName(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeName()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }

// func TestLinWriteExpireDate(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeExpireDate()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }

// func TestLinkWriteMemberLimit(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeMemberLimit()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }

// func TestLinkWriteJoinRequest(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeJoinRequest()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }

// func TestLinkWriteIniveLink(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeIniveLink()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }

// func TestLinkWriteSubscriptionPeriod(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeSubscriptionPeriod()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }

// func TestLinkWriteSubscriptionPrice(t *testing.T) {
// 	t.Parallel()
// 	linktc := new(linkTestContainer)
// 	linktc.writeSubscriptionPrice()
// 	msg := formatter.CreateEmpltyMessage()
// 	linktc.mainLogic(msg, t)
// }
