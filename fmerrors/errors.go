package fmerrors

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

const requiredColumnWidth = 60

type FME struct {
	Code   int
	String string
	Stack  string
}

// func (f *FME) Error() string {
// 	f.Stack = string(debug.Stack())
// 	return fmt.Sprintf("[ERROR]\tcode: %d\tmessage: %s\nStack: %s", f.Code, f.String, f.Stack)
// }

func (f *FME) Error() string {
	return fmt.Sprintf("%d", f.Code)
}

func TelegramError(code int, dis string) error {
	return fmt.Errorf("[ERROR:%d] %s", code, dis)
}

func AssertTest(compared, comparedfunc, comparator, comparatorfunc string) (err error) {
	body := "[TEST ERROR] The data doesn't match. The original passed value in %s (%s) is different from the final %s (%s)"
	err = fmt.Errorf(fmt.Sprintf(body, comparedfunc, compared, comparatorfunc, comparator))
	return err
}

func JustError(message string) (err error) {
	err = fmt.Errorf("[TEST ERROR] The length of the passed arrays does not match. %s", message)
	return err
}

func UpdatesMisstakes(part string) (err error) {
	body := fmt.Sprintf("[ERROR] Couldn't get any updates: %s", part)
	err = fmt.Errorf(body)

	return err
}

func MadeMisstake(err error) {
	log.Println("You have a misstake:", err)
	debug.PrintStack()
}

func MakeIntestines() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from log.Fatal:", r)
		debug.PrintStack()
		os.Exit(1)
	}
}

func ErrorInUserData(err error) {
	log.Printf("[ERROR] You have an error in user's data from the response of Telegram: %s", err)
}

func ChatIDIsNil() {
	log.Print("[ERROR] You're trying to delete a message, but you haven't indicated the ChatID where the bot should delete the message")
}

func MessageIDIsNil() {
	log.Print("[ERROR] You're trying to delete a message, but you haven't indicated the Deleted Message ID which bot should delete")
}

func RequestError(err error) {
	log.Printf("[ERROR] You have an error while trying to send a request to Telegram: %s", err)
}

func CantOpenFile(err error) {
	log.Printf("[ERROR] Fmtogram couldn't open the file to send after: %s", err)
}

func CantCopyFile(err error) {
	log.Printf("[ERROR] Fmtgoram couldn't copy a part to the file: %s", err)
}

func CantCreateFormFile(err error) {
	log.Printf("[ERROR] Fmtogram couldn't create a Form File: %s", err)
}

func CantWriteField(err error) {
	log.Printf("[ERROR] Fmtogram couldn't write a multipart field: %s", err)
}

func CantMarshalJSON(err error) {
	log.Printf("[ERROR] Fmtogram couldn't marshal a json object: %s", err)
}

func ChatIDIsMissed() error {
	return fmt.Errorf("[ERROR]\tYou are required to provide the data in WriteChat{ID,Name}")
}

func MissedGottenFrom(object, kindofgotten string, num int) error {
	return fmt.Errorf("[ERROR] in object {%s} (number in queue is %d) missed %s", object, num, kindofgotten)
}

func ButtosDoesntFit(val string, data int) error {
	return fmt.Errorf("[ERROR] data (%d) in value '%s' is out of range. 'line' is a position in the first array in [][]. 'pos' it is a position in the second array in [][]", data, val)
}

func MustBe(objects, fields, funcnames []string) error {
	var mesErr string
	for i := range objects {
		mesErr += fmt.Sprintf("There must be some data in %s field in object %s. Use %s to add your data.\n", fields[i], objects[i], funcnames[i])
	}
	return fmt.Errorf("[ERROR] Some data is missed! There's more info:\n %s", mesErr)
}

func MissedRequiredField(intrfc, field string, interfacenum, interfaceline int, media, kb bool) error {
	var msg string
	if len(field) > requiredColumnWidth {
		field = field[:requiredColumnWidth-3] + "..."
	} else {
		field = fmt.Sprintf("%-*s", requiredColumnWidth, field)
	}

	table := fmt.Sprintf(
		"[ERROR] Required data is missing:\n\n"+
			"┌───────────────┬──────────────────────────────────────────────────────────────┐\n"+
			"│ Interface     │ Required                                                     │\n"+
			"├───────────────┼──────────────────────────────────────────────────────────────┤\n"+
			"│ %-13s │ %s │\n"+
			"└───────────────┴──────────────────────────────────────────────────────────────┘\n",
		intrfc, field,
	)
	if media {
		msg = fmt.Sprintf("%sInterface was added as %d", table, interfacenum)
	} else if kb {
		msg = fmt.Sprintf("%sInterface was added as %d on line %d", table, interfacenum, interfaceline)
	} else {
		msg = table
	}

	return fmt.Errorf(msg)
}

func CantMakeRequest(err error) error {
	return fmt.Errorf("[ERROR] Somthing went wrong! Couldn't send the request! more information: %s", err)
}

func CantGetResponse(err error) error {
	return fmt.Errorf("[ERROR] Somthing went wrong! Couldn't get the response! more information: %s", err)
}

func CantReadResponse(err error) error {
	return fmt.Errorf("[ERROR] Somthing went wrong! Couldn't read the response! more information: %s", err)
}

func CantUnmarshal(err error) error {
	return fmt.Errorf("[ERROR] Somthing went wrong! Couldn't unmarshal the response ([]byte)! more information: %s", err)
}

// func TelegramError(errCode int, description string) error {
// 	var msg string
// 	if len(description) > requiredColumnWidth+20 {
// 		description = description[:requiredColumnWidth+17] + "..."
// 	} else {
// 		description = fmt.Sprintf("%-*s", requiredColumnWidth, description)
// 	}

// 	msg = fmt.Sprintf(
// 		"[ERROR] Telegram responsed with an error:\n\n"+
// 			"┌───────────────┬──────────────────────────────────────────────────────────────────────────────────┐\n"+
// 			"│ Code          │ Description                                                                      │\n"+
// 			"├───────────────┼──────────────────────────────────────────────────────────────────────────────────┤\n"+
// 			"│ %-13d │ %s \n"+
// 			"└───────────────┴──────────────────────────────────────────────────────────────────────────────────┘\n",
// 		errCode, description,
// 	)
// 	return fmt.Errorf(msg)
// }

func ImpossibleCombination(object1, object2, method string) error {
	return fmt.Errorf("[ERROR] Impossible combination. It isn't allowed to send '%s' and '%s' in one message. The chosen method: %s", object1, object2, method)
}

func UnsupportedType(pos int) error {
	return fmt.Errorf("[ERROR] The data type in slice messageIDs in position %d isn't allowed. Allowed types are int and string", pos)
}

func TooMuchData(interface1 string, num, line int) error {
	return fmt.Errorf("[ERROR] You mentioned too much data. It is allowed to call only one function in interface {%s} except WriteString(). The interface where you gave this data is %d in the %d line", interface1, num, line)
}

func IlligalFields(interface1, f string, num int) error {
	return fmt.Errorf("[ERROR] You call '%s' function and gave it some data (if it was required) in interface {%s} with the number in queue: %d. It is not allowed if you want to send not only this interface", f, interface1, num)
}

func UnsupportedTypeOfFile(typesoffiles string) error {
	return fmt.Errorf("[ERROR] You want to send unnsported file type. The suported types: %s", typesoffiles)
}

func MisMedia() error {
	return fmt.Errorf("[ERROR] You're supposed to add a media interface if you specified a media method.\n" +
		"Add a media interface according to your method, or do not specify any method.\n" +
		"Here are all available media interfaces:\n\n" +
		"┌──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐\n" +
		"│  formatter.IPhoto    │  formatter.IAudio    │  formatter.IVideo    │  formatter.IDocument │\n" +
		"├──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤\n" +
		"│  formatter.IAnimation│  formatter.IVoice    │  formatter.IVideoNote│                      │\n" +
		"└──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘")

}
