package unit

import (
	"strings"
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
)

type forumT struct {
	name          string
	str           string
	integer       int
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type frmTestContainer struct {
	name          string
	inputStr      []string
	inputInt      []int
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(frm formatter.IForum, f *forumT, i int)
}

func (frmtc *frmTestContainer) defaultForumT(i int) *forumT {
	return &forumT{name: frmtc.name, isExpectedErr: frmtc.isExpectedErr[i], codeErr: frmtc.codeErr[i]}
}

func (frmtc *frmTestContainer) Wname(frm formatter.IForum, f *forumT, i int) {
	f.testedFunc = frm.WriteName
	f.str = frmtc.inputStr[i]
}

func (frmtc *frmTestContainer) iconColor(frm formatter.IForum, f *forumT, i int) {
	f.testedFunc = frm.WriteIconColor
	f.integer = frmtc.inputInt[i]
}

func (frmtc *frmTestContainer) iconEmojiID(frm formatter.IForum, f *forumT, i int) {
	f.testedFunc = frm.WriteIconEmojiID
	f.str = frmtc.inputStr[i]
}

func (frmtc *frmTestContainer) writeName() {
	frmtc.name = "(IForum).WriteName()"
	frmtc.inputStr = []string{"a;l,ksdsal;'d", "", "y", strings.Repeat("h", 128), strings.Repeat("i", 129), "name", "noname"}
	frmtc.isExpectedErr = []bool{false, true, false, false, true, false, true}
	frmtc.codeErr = []string{"", "20", "", "", "20", "", "10"}
	frmtc.amount, frmtc.until = 7, 5
	frmtc.buildF = frmtc.Wname
}

func (frmtc *frmTestContainer) writeIconColor() {
	frmtc.name = "(IForum).WriteIconColor()"
	frmtc.inputInt = []int{0xFFD67E, 0, 230123901, 0xFFD67E, 0xCB86DB, 0x8EEE98, 0xFF93B2, 0xFB6F5F,
		0xFB6F5F, 0xFB6F5F}
	frmtc.isExpectedErr = []bool{false, true, true, false, false, false, false, false,
		false, true}
	frmtc.codeErr = []string{"", "20", "20", "", "", "", "", "",
		"", "10"}
	frmtc.amount, frmtc.until = 10, 8
	frmtc.buildF = frmtc.iconColor
}

func (frmtc *frmTestContainer) writeIconEmojiID() {
	frmtc.name = "(IForum).WriteIconEmojiID()"
	frmtc.inputStr = []string{"a;l,ksdsal;'d", "", "y", "name", "noname"}
	frmtc.isExpectedErr = []bool{false, false, false, false, true}
	frmtc.codeErr = []string{"", "", "", "", "10"}
	frmtc.amount, frmtc.until = 5, 3
	frmtc.buildF = frmtc.iconEmojiID
}

func (frm *forumT) startTest(part string, i int, t *testing.T) {
	switch f := frm.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, frm.name, frm.codeErr, frm.str, frm.isExpectedErr, i, t)
		checkError(f(frm.str), frm.isExpectedErr, frm.codeErr, t)
	case func(int) error:
		printTestLog(part, frm.name, frm.codeErr, frm.integer, frm.isExpectedErr, i, t)
		checkError(f(frm.integer), frm.isExpectedErr, frm.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (frmtc *frmTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var st formatter.IForum
	a, b := make([]UnitTest, frmtc.until), make([]UnitTest, frmtc.amount-frmtc.until)
	for i, j := 0, 0; i < frmtc.amount; i++ {
		s := frmtc.defaultForumT(i)
		if i < frmtc.until {
			frmtc.buildF(msg.NewForum(), s, i)
			a[i] = s
		} else {
			if j%2 == 0 {
				st = msg.NewForum()
			}
			frmtc.buildF(st, s, i)
			b[j] = s
			j++
		}
	}
	return a, b
}

func (frmtc *frmTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	frmcontainer, doublecontainer := frmtc.createTestArrays(msg)
	check(frmcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestForumWriteName(t *testing.T) {
	t.Parallel()
	frmtc := new(frmTestContainer)
	frmtc.writeName()
	msg := formatter.CreateEmpltyMessage()
	frmtc.mainLogic(msg, t)
}

func TestForumWriteIconColor(t *testing.T) {
	t.Parallel()
	frmtc := new(frmTestContainer)
	frmtc.writeIconColor()
	msg := formatter.CreateEmpltyMessage()
	frmtc.mainLogic(msg, t)
}
func TestForumWriteIconEmojiID(t *testing.T) {
	t.Parallel()
	frmtc := new(frmTestContainer)
	frmtc.writeIconEmojiID()
	msg := formatter.CreateEmpltyMessage()
	frmtc.mainLogic(msg, t)
}
