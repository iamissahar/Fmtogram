package tests

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/types"
)

func writeReqContactFields(msg *formatter.Message, con formatter.IContact, t *testing.T) {
	if err := con.WritePhoneNumber("+79041293212245"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteFirstName("John"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddContact(con); err != nil {
		t.Fatal(err)
	}
}

func writeUnreqContactFields(msg *formatter.Message, con formatter.IContact, t *testing.T) {
	if err := con.WritePhoneNumber("+79041293212245"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteFirstName("John"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteLastName("Doe"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteVCard("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nTEL:+1234567890\nEND:VCARD"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddContact(con); err != nil {
		t.Fatal(err)
	}
}

func scontactReqFieldsWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	ch := createChat(msg, t)
	writeReqContactFields(msg, con, t)
	if err := msg.AddMethod(methods.Contact); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(con.GetResponse())
	t.Log(ch.GetResponse())
}

func scontactReqFieldsWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	ch := createChat(msg, t)
	writeReqContactFields(msg, con, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(con.GetResponse())
	t.Log(ch.GetResponse())
}

func scontactReqFields(t *testing.T) {
	t.Run("WithName", scontactReqFieldsWithName)
	t.Run("WithoutName", scontactReqFieldsWithoutName)
}

func scontactUnreqFieldsWithName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	param := msg.NewMessage()
	ch := createChat(msg, t)
	writeUnreqContactFields(msg, con, t)
	if err := param.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteMessageEffectID("507"); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	createInlineKeyboard(msg, t)
	if err := msg.AddMethod(methods.Contact); err != nil {
		t.Fatal(err)
	}
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(con.GetResponse())
	t.Log(ch.GetResponse())
}

func scontactUnreqFieldsWithoutName(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	param := msg.NewMessage()
	ch := createChat(msg, t)
	writeUnreqContactFields(msg, con, t)
	if err := param.WriteDisableNotification(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteProtectContent(); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteMessageEffectID("507"); err != nil {
		t.Fatal(err)
	}
	if err := param.WriteReplyParameters(&types.ReplyParameters{MessageID: 996, ChatID: 738070596}); err != nil {
		t.Fatal(err)
	}
	createReplyKeyboard(msg, t)
	if _, err := msg.Send(); err != nil {
		t.Fatal(err)
	}
	t.Log(con.GetResponse())
	t.Log(ch.GetResponse())
}

func scontactUnreqFields(t *testing.T) {
	t.Run("WithName", scontactUnreqFieldsWithName)
	t.Run("WithoutName", scontactUnreqFieldsWithoutName)
}

func sendContactFunctional(t *testing.T) {
	t.Run("ReqFields", scontactReqFields)
	t.Run("UnreqFields", scontactUnreqFields)
}

func contactFunctional(t *testing.T) {
	t.Run("sendContact", sendContactFunctional)
}

func contactPhoneNumOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WritePhoneNumber("+79041293212"); err != nil {
		t.Fatal(err)
	}
}

func contactPhoneNum10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WritePhoneNumber("+79041293212556"); err != nil {
		t.Fatal(err)
	}
	if err := con.WritePhoneNumber("+79041293212556"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func contactPhoneNum20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WritePhoneNumber(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func contactPhoneNum(t *testing.T) {
	t.Run("OK", contactPhoneNumOK)
	t.Run("Code10", contactPhoneNum10)
	t.Run("Code20", contactPhoneNum20)
}

func contactFirstNOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteFirstName("John"); err != nil {
		t.Fatal(err)
	}
}

func contactFirstN10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteFirstName("John"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteFirstName("John"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func contactFirstN20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteFirstName(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func contactFirstN(t *testing.T) {
	t.Run("OK", contactFirstNOK)
	t.Run("Code10", contactFirstN10)
	t.Run("Code20", contactFirstN20)
}

func contactLastNOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteLastName("Doe"); err != nil {
		t.Fatal(err)
	}
}

func contactLastN10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteLastName("Doe"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteLastName("Doe"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func contactLastN20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteLastName(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func contactLastN(t *testing.T) {
	t.Run("OK", contactLastNOK)
	t.Run("Code10", contactLastN10)
	t.Run("Code20", contactLastN20)
}

func contactVCardOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteVCard("vCard"); err != nil {
		t.Fatal(err)
	}
}

func contactVCard10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteVCard("vCard"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteVCard("vCard"); err.Error() != code10 {
		t.Fatal(err)
	}
}

func contactVCard20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WriteVCard(""); err.Error() != code20 {
		t.Fatal(err)
	}
}

func contactVCard(t *testing.T) {
	t.Run("OK", contactVCardOK)
	t.Run("Code10", contactVCard10)
	t.Run("Code20", contactVCard20)
}

func contactAddOK(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WritePhoneNumber("+79041293212556"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteFirstName("John"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddContact(con); err != nil {
		t.Fatal(err)
	}
}

func contactAdd10(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := con.WritePhoneNumber("+79041293212556"); err != nil {
		t.Fatal(err)
	}
	if err := con.WriteFirstName("John"); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddContact(con); err != nil {
		t.Fatal(err)
	}
	if err := msg.AddContact(con); err.Error() != code10 {
		t.Fatal(err)
	}
}

func contactAdd20(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	if err := msg.AddContact(nil); err.Error() != code20 {
		t.Fatal(err)
	}
}

func contactAdd21(t *testing.T) {
	msg := formatter.CreateEmpltyMessage()
	con := msg.NewContact()
	if err := msg.AddContact(con); err.Error() != code21 {
		t.Fatal(err)
	}
}

func contactAdd(t *testing.T) {
	t.Run("OK", contactAddOK)
	t.Run("Code10", contactAdd10)
	t.Run("Code20", contactAdd20)
	t.Run("Code21", contactAdd21)
}

func contactUnit(t *testing.T) {
	t.Parallel()
	t.Run("PhoneNumber", contactPhoneNum)
	t.Run("FirstName", contactFirstN)
	t.Run("LastName", contactLastN)
	t.Run("vCard", contactVCard)
	t.Run("Add", contactAdd)
}

func TestSendContact(t *testing.T) {
	t.Parallel()
	t.Run("Functional", contactFunctional)
	t.Run("Unit", contactUnit)
}
