package formatter

// func (fm *Formatter) WriteMessageID()
// ReplyParameters
// ReplyMarkup

// // Save a text of the message for sending
// func (fm *Formatter) WriteString(text string) {
// 	fm.Message.Text = text
// 	logs.DataWrittenSuccessfully("Text Of The Message")
// }

// // Save a Parse Mode of the message for sending
// func (fm *Formatter) WriteParseMode(mode string) {
// 	fm.Message.ParseMode = mode
// 	logs.DataWrittenSuccessfully("Parse Mode")
// }

// func (fm *Formatter) ProtectAllContent() {
// 	fm.Message.ProtectContent = true
// 	logs.DataWrittenSuccessfully("Protected Content")
// }

// func (fm *Formatter) ProtectMedia(numberOfMedia int) {

// }

// func (fm *Formatter) AddCaptionEntities(entities []types.MessageEntity) {
// 	fm.Message.CaptionEntities = entities
// }

// func (fm *Formatter) SetIkbdDim(dim []int) {
// 	logs.DataWrittenSuccessfully("The Structure Of Inline Keyboard")
// 	fm.Keyboard.Keyboard = make([][]btn, len(dim))
// 	for i := 0; i < len(dim); i++ {
// 		fm.Keyboard.Keyboard[i] = make([]btn, dim[i])
// 	}
// }

// func (fm *Formatter) doRutine() {
// 	if fm.Keyboard.x == len(fm.Keyboard.Keyboard[fm.Keyboard.y]) {
// 		fm.Keyboard.x = 0
// 		fm.Keyboard.y = fm.Keyboard.y + 1
// 	}
// }

// func (fm *Formatter) WriteInlineButtonCmd(label, cmd string) {
// 	logs.DataWrittenSuccessfully("A CMD-Button Of Inline Keyboard")
// 	fm.doRutine()
// 	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Label = label
// 	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].what = bCmd
// 	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Cmd = cmd

// 	fm.Keyboard.x = fm.Keyboard.x + 1

// }

// func (fm *Formatter) WriteInlineButtonUrl(label, url string) {
// 	logs.DataWrittenSuccessfully("A URL-Button Of Inline Keyboard")
// 	fm.doRutine()
// 	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Label = label
// 	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].what = bUrl
// 	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Url = url

// 	fm.Keyboard.x = fm.Keyboard.x + 1
// }
