package unit

import (
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
)

type contactT struct {
	name          string
	str           string
	testedFunc    func(string) error
	isExpectedErr bool
	codeErr       string
}

type conTestContainer struct {
	name          string
	inputStr      []string
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(contc conTestContainer, con formatter.IContact, i int) *contactT
}

func putWriteContactPhoneNumber(contc conTestContainer, con formatter.IContact, i int) *contactT {
	return &contactT{contc.name, contc.inputStr[i], con.WritePhoneNumber, contc.isExpectedErr[i], contc.codeErr[i]}
}

func putWriteContactFirstName(contc conTestContainer, con formatter.IContact, i int) *contactT {
	return &contactT{contc.name, contc.inputStr[i], con.WriteFirstName, contc.isExpectedErr[i], contc.codeErr[i]}
}

func putWriteContactLastName(contc conTestContainer, con formatter.IContact, i int) *contactT {
	return &contactT{contc.name, contc.inputStr[i], con.WriteLastName, contc.isExpectedErr[i], contc.codeErr[i]}
}

func putWriteContactVCard(contc conTestContainer, con formatter.IContact, i int) *contactT {
	return &contactT{contc.name, contc.inputStr[i], con.WriteVCard, contc.isExpectedErr[i], contc.codeErr[i]}
}

func (conct *conTestContainer) writePhoneNumber() {
	conct.name = "(IContact).WritePhoneNumber"
	conct.inputStr = []string{"phone", "", "0912315", "20874172783"}
	conct.isExpectedErr = []bool{false, true, false, true}
	conct.codeErr = []string{"", "20", "", "10"}
	conct.amount, conct.until = 4, 2
	conct.buildF = putWriteContactPhoneNumber
}

func (conct *conTestContainer) writeFirstName() {
	conct.name = "(IContact).WriteFirstName"
	conct.inputStr = []string{"phone", "", "0912315", "20874172783"}
	conct.isExpectedErr = []bool{false, true, false, true}
	conct.codeErr = []string{"", "20", "", "10"}
	conct.amount, conct.until = 4, 2
	conct.buildF = putWriteContactFirstName
}

func (conct *conTestContainer) writeLastName() {
	conct.name = "(IContact).WriteLastName"
	conct.inputStr = []string{"phone", "", "0912315", "20874172783"}
	conct.isExpectedErr = []bool{false, true, false, true}
	conct.codeErr = []string{"", "20", "", "10"}
	conct.amount, conct.until = 4, 2
	conct.buildF = putWriteContactLastName
}

func (conct *conTestContainer) writeVCard() {
	conct.name = "(IContact).WriteVCard"
	conct.inputStr = []string{"phone", "", "0912315", "20874172783"}
	conct.isExpectedErr = []bool{false, true, false, true}
	conct.codeErr = []string{"", "20", "", "10"}
	conct.amount, conct.until = 4, 2
	conct.buildF = putWriteContactVCard
}
func (con *contactT) startTest(part string, i int, t *testing.T) {
	printTestLog(part, con.name, con.codeErr, con.str, con.isExpectedErr, i, t)
	if !con.isExpectedErr {
		if err := con.testedFunc(con.str); err != nil {
			t.Fatal(err)
		}
	} else {
		if err := con.testedFunc(con.str); err.Error() != con.codeErr {
			t.Fatal(err)
		}
	}
}

func (contc *conTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var con formatter.IContact
	a, b := make([]UnitTest, contc.until), make([]UnitTest, contc.amount-contc.until)
	for i, j := 0, 0; i < contc.amount; i++ {
		if i < contc.until {
			con = msg.NewContact()
			a[i] = contc.buildF(*contc, con, i)
		} else {
			if j%2 == 0 {
				con = msg.NewContact()
			}
			b[j] = contc.buildF(*contc, con, i)
			j++
		}
	}
	return a, b
}

func mainContactLogic(msg *formatter.Message, contc conTestContainer, t *testing.T) {
	contactcontainer, doublecontainer := contc.createTestArrays(msg)
	check(contactcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestWriteContactPhoneNumber(t *testing.T) {
	t.Parallel()
	contc := new(conTestContainer)
	contc.writePhoneNumber()
	msg := formatter.CreateEmpltyMessage()
	mainContactLogic(msg, *contc, t)
}

func TestWriteContactFirstName(t *testing.T) {
	t.Parallel()
	contc := new(conTestContainer)
	contc.writeFirstName()
	msg := formatter.CreateEmpltyMessage()
	mainContactLogic(msg, *contc, t)
}

func TestWriteContactLastName(t *testing.T) {
	t.Parallel()
	contc := new(conTestContainer)
	contc.writeLastName()
	msg := formatter.CreateEmpltyMessage()
	mainContactLogic(msg, *contc, t)
}

func TestWriteContactVCard(t *testing.T) {
	t.Parallel()
	contc := new(conTestContainer)
	contc.writeVCard()
	msg := formatter.CreateEmpltyMessage()
	mainContactLogic(msg, *contc, t)
}
