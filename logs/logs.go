package logs

import (
	"log"

	"github.com/l1qwie/Fmtogram/types"
)

var Caption string

func DebugOrInfo() {
	if types.DEBUG {
		Caption = "DEBUG"
	} else {
		Caption = "INFO"
		types.INFO = true
	}
}

func TurnOn() {
	log.Printf("[%s] The bot has been turned on", Caption)
}

func TurnOff() {
	log.Printf("[%s] The bot has been turned off", Caption)
}

func GetOffset() {
	if types.DEBUG {
		log.Panicf("[%s] Started trying to get offset from Telegram", Caption)
	}
}

func GottenOffset() {
	if types.DEBUG {
		log.Printf("[%s] The offset was given from Telegram and successfuly gotten by Fmtogram", Caption)
	}
}

func FoundSomeIformation() {
	if types.DEBUG {
		log.Printf("[%s] The bots has found some information about a user without any misstakes", Caption)
	}
}

func AlreadyInQueue() {
	if types.DEBUG {
		log.Printf("[%s] The user has already been placed in their queue", Caption)
	}
}

func FirstTimeUser() {
	if types.DEBUG {
		log.Printf("[%s] The user is a new one and has been just placed in their queue", Caption)
	}
}

func CallDeveloperFunc() {
	log.Printf("[%s] The imported function has just called by Fmtogram with some data", Caption)
}

func ReturnedIsntEmply() {
	if types.DEBUG {
		log.Printf("[%s] The returned data of the request to Telegram isn't emply and has some data", Caption)
	}
}

func ReturnedIsEmply() {
	if types.DEBUG {
		log.Printf("[%s] The returned data of the request to Telegram is emply and 'returned data' has been just genereted", Caption)
	}
}

func GottenRequest() {
	if types.DEBUG {
		log.Printf("[%s] Fmtogram has just gotten the request to send in to Telegram", Caption)
	}
}

func ErrorDuringSending() {
	log.Printf("[%s] You have an error during sending the request to Telegram", Caption)
}

func DataWrittenSuccessfully(object, text string) {
	if types.DEBUG {
		log.Printf("[%s] {%s} %s saved successfuly", Caption, object, text)
	}
}

func SettedParam(param, object string, turnOn bool) {
	var verb string
	if turnOn {
		verb = "has been turned on"
	} else {
		verb = "has been turned off"
	}
	if types.DEBUG {
		log.Printf("[%s] '%s' in objcet {%s} %s", Caption, param, object, verb)
	}
}

func NewInterfaceCreated(text string) {
	if types.DEBUG {
		log.Printf("[%s] %s interface has successfuly created", Caption, text)
	}
}

func NewInterface(text string) {
	if types.DEBUG {
		log.Printf("[%s] new interface %s has just created", Caption, text)
	}
}

func DataIsntEmply(object, param string, data interface{}) {
	if !types.NoWarningMessages {
		log.Printf("[WARNING] data that was gotten from '%s' in interface {%s} wasn't added to an empty field. Make sure that you've edited the right interface. Fmtogram is anyway going to add this data to the interface, but just to be sure the data before had been: '%v'. P.S. in case everything is OK, you can turn this message off. Just make types.NoWarningMessages true", param, object, data)
	}
}

func ObjectSaved(object string) {
	if types.DEBUG {
		log.Printf("[%s] object '%s' has been saved", Caption, object)
	}
}

func InterfaceSaved(interf string) {
	if types.DEBUG {
		log.Printf("[%s] interface {%s} has saved to the message successfuly", Caption, interf)
	}
}

func DataMightBeLost(intrfc, function, data, intrfc1, function1 string) {
	if !types.NoWarningMessages {
		log.Printf("[WARNING] You might lose your data. You put '%s' in function %s in {%s} interface. Insted put it in {%s} interface, function %s", data, function, intrfc, intrfc1, function1)
	}
}

func MethodSaved(method string) {
	if types.DEBUG {
		log.Printf("[%s] method '%s' has been saved as a prefered one", Caption, method)
	}
}
