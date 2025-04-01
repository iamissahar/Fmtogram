package unit

// import (
// 	"testing"

// 	"github.com/iamissahar/Fmtogram/formatter"
// )

// type gameT struct {
// 	name          string
// 	str           string
// 	integer       int
// 	testedFunc    interface{}
// 	isExpectedErr bool
// 	codeErr       string
// }

// type gameTestContainer struct {
// 	name          string
// 	inputStr      []string
// 	inputInt      []int
// 	isExpectedErr []bool
// 	codeErr       []string
// 	amount, until int
// 	buildF        func(game formatter.IGame, g *gameT, i int)
// }

// func (gametc *gameTestContainer) shortName(game formatter.IGame, g *gameT, i int) {
// 	g.testedFunc = game.WriteShortName
// 	g.str = gametc.inputStr[i]
// }

// func (gametc *gameTestContainer) score(game formatter.IGame, g *gameT, i int) {
// 	g.testedFunc = game.WriteScore
// 	g.integer = gametc.inputInt[i]
// }

// func (gametc *gameTestContainer) force(game formatter.IGame, g *gameT, i int) {
// 	g.testedFunc = game.WriteForce
// }

// func (gametc *gameTestContainer) disableEditMessage(game formatter.IGame, g *gameT, i int) {
// 	g.testedFunc = game.WriteDisableEditMessage
// }

// func (gametc *gameTestContainer) writeShortName() {
// 	gametc.name = "(IGame).WriteShortName()"
// 	gametc.inputStr = []string{"asl;dl;asdl;asd", "", "A:KLSDKL:ASDKL:", "L:KASDKL:AS:LG:AKLS"}
// 	gametc.isExpectedErr = []bool{false, true, false, true}
// 	gametc.codeErr = []string{"", "20", "", "10"}
// 	gametc.amount, gametc.until = 4, 2
// 	gametc.buildF = gametc.shortName
// }

// func (gametc *gameTestContainer) writeScore() {
// 	gametc.name = "(IGame).WriteScore()"
// 	gametc.inputInt = []int{1231, 0, -1, 22, 44}
// 	gametc.isExpectedErr = []bool{false, false, true, false, true}
// 	gametc.codeErr = []string{"", "", "20", "", "10"}
// 	gametc.amount, gametc.until = 5, 3
// 	gametc.buildF = gametc.score
// }

// func (gametc *gameTestContainer) writeForce() {
// 	gametc.name = "(IGame).WriteForce()"
// 	gametc.isExpectedErr = []bool{false, false, true}
// 	gametc.codeErr = []string{"", "", "10"}
// 	gametc.amount, gametc.until = 3, 1
// 	gametc.buildF = gametc.force
// }

// func (gametc *gameTestContainer) writeDisableEditMessage() {
// 	gametc.name = "(IGame).WriteDisableEditMessage()"
// 	gametc.isExpectedErr = []bool{false, false, true}
// 	gametc.codeErr = []string{"", "", "10"}
// 	gametc.amount, gametc.until = 3, 1
// 	gametc.buildF = gametc.disableEditMessage
// }

// func (pay *gameT) startTest(part string, i int, t *testing.T) {
// 	switch f := pay.testedFunc.(type) {
// 	case func(string) error:
// 		printTestLog(part, pay.name, pay.codeErr, pay.str, pay.isExpectedErr, i, t)
// 		checkError(f(pay.str), pay.isExpectedErr, pay.codeErr, t)
// 	case func(int) error:
// 		printTestLog(part, pay.name, pay.codeErr, pay.integer, pay.isExpectedErr, i, t)
// 		checkError(f(pay.integer), pay.isExpectedErr, pay.codeErr, t)
// 	case func() error:
// 		printTestLog(part, pay.name, pay.codeErr, true, pay.isExpectedErr, i, t)
// 		checkError(f(), pay.isExpectedErr, pay.codeErr, t)
// 	default:
// 		t.Fatal("unexpected type of tested function")
// 	}
// }

// func (gametc *gameTestContainer) byDefault(i int) *gameT {
// 	return &gameT{name: gametc.name, isExpectedErr: gametc.isExpectedErr[i], codeErr: gametc.codeErr[i]}
// }

// func (gametc *gameTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
// 	var game formatter.IGame
// 	a, b := make([]UnitTest, gametc.until), make([]UnitTest, gametc.amount-gametc.until)
// 	for i, j := 0, 0; i < gametc.amount; i++ {
// 		s := gametc.byDefault(i)
// 		if i < gametc.until {
// 			gametc.buildF(msg.NewGame(), s, i)
// 			a[i] = s
// 		} else {
// 			if j%2 == 0 {
// 				game = msg.NewGame()
// 			}
// 			gametc.buildF(game, s, i)
// 			b[j] = s
// 			j++
// 		}
// 	}
// 	return a, b
// }

// func (gametc *gameTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
// 	gamecontainer, doublecontainer := gametc.createTestArrays(msg)
// 	check(gamecontainer, t)
// 	doubleCheck(doublecontainer, t)
// }

// func TestGameWriteShortName(t *testing.T) {
// 	t.Parallel()
// 	gametc := new(gameTestContainer)
// 	gametc.writeShortName()
// 	msg := formatter.CreateEmpltyMessage()
// 	gametc.mainLogic(msg, t)
// }

// func TestGameWriteScore(t *testing.T) {
// 	t.Parallel()
// 	gametc := new(gameTestContainer)
// 	gametc.writeScore()
// 	msg := formatter.CreateEmpltyMessage()
// 	gametc.mainLogic(msg, t)
// }

// func TestGameWriteForce(t *testing.T) {
// 	t.Parallel()
// 	gametc := new(gameTestContainer)
// 	gametc.writeForce()
// 	msg := formatter.CreateEmpltyMessage()
// 	gametc.mainLogic(msg, t)
// }

// func TestGameWriteDisableEditMessage(t *testing.T) {
// 	t.Parallel()
// 	gametc := new(gameTestContainer)
// 	gametc.writeDisableEditMessage()
// 	msg := formatter.CreateEmpltyMessage()
// 	gametc.mainLogic(msg, t)
// }
