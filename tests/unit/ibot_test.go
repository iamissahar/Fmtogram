package unit

import (
	"strings"
	"testing"

	"github.com/iamissahar/Fmtogram/formatter"
	"github.com/iamissahar/Fmtogram/types"
)

type botT struct {
	name          string
	str           string
	commands      []*types.BotCommand
	scope         *types.BotCommandScope
	mbutt         *types.MenuButton
	rights        *types.ChatAdministratorRights
	testedFunc    interface{}
	isExpectedErr bool
	codeErr       string
}

type botTestContainer struct {
	name          string
	inputStr      []string
	inputCom      [][]*types.BotCommand
	inputScope    []*types.BotCommandScope
	inputMButton  []*types.MenuButton
	inputRights   []*types.ChatAdministratorRights
	isExpectedErr []bool
	codeErr       []string
	amount, until int
	buildF        func(bot formatter.IBot, b *botT, i int)
}

func (bottc *botTestContainer) defaultBotT(i int) *botT {
	return &botT{name: bottc.name, isExpectedErr: bottc.isExpectedErr[i], codeErr: bottc.codeErr[i]}
}

func (bottc *botTestContainer) commands(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteCommands
	b.commands = bottc.inputCom[i]
}

func (bottc *botTestContainer) scope(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteScope
	b.scope = bottc.inputScope[i]
}

func (bottc *botTestContainer) language(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteLanguage
	b.str = bottc.inputStr[i]
}

func (bottc *botTestContainer) Wname(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteName
	b.str = bottc.inputStr[i]
}

func (bottc *botTestContainer) description(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteDescription
	b.str = bottc.inputStr[i]
}

func (bottc *botTestContainer) shortDescription(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteShortDescription
	b.str = bottc.inputStr[i]
}

func (bottc *botTestContainer) menuButton(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteMenuButton
	b.mbutt = bottc.inputMButton[i]
}

func (bottc *botTestContainer) rights(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteRights
	b.rights = bottc.inputRights[i]
}

func (bottc *botTestContainer) forchannels(bot formatter.IBot, b *botT, i int) {
	b.testedFunc = bot.WriteForChannels
}

func (bottc *botTestContainer) writeCommands() {
	bottc.name = "(IBot).WriteCommands"
	bottc.inputCom = [][]*types.BotCommand{{{Command: "Something", Description: "Does something"}},
		nil,
		{{Command: "Something", Description: "Does something"}, nil, nil},
		{{Command: "Something", Description: "Does something"}}, {{Command: "Something", Description: "Does something"}}}
	bottc.isExpectedErr = []bool{false, true, true, false, true}
	bottc.codeErr = []string{"", "20", "5", "", "10"}
	bottc.amount, bottc.until = 5, 3
	bottc.buildF = bottc.commands
}

func (bottc *botTestContainer) writeScope() {
	bottc.name = "(IBot).WriteScope"
	bottc.inputScope = []*types.BotCommandScope{{}, nil, {}, {}}
	bottc.isExpectedErr = []bool{false, true, false, true}
	bottc.codeErr = []string{"", "20", "", "10"}
	bottc.amount, bottc.until = 4, 2
	bottc.buildF = bottc.scope
}

func (bottc *botTestContainer) writeLanguage() {
	bottc.name = "(IBot).WriteLanguage"
	bottc.inputStr = []string{"en", "", "asdasda", "english", "ru", "hb"}
	bottc.isExpectedErr = []bool{false, false, true, true, false, true}
	bottc.codeErr = []string{"", "", "20", "20", "", "10"}
	bottc.amount, bottc.until = 6, 4
	bottc.buildF = bottc.language
}

func (bottc *botTestContainer) writeName() {
	bottc.name = "(IBot).WriteName"
	bottc.inputStr = []string{"Name", "", strings.Repeat("h", 64), strings.Repeat("k", 65), "is a name", "noname"}
	bottc.isExpectedErr = []bool{false, false, false, true, false, true}
	bottc.codeErr = []string{"", "", "", "20", "", "10"}
	bottc.amount, bottc.until = 6, 4
	bottc.buildF = bottc.Wname
}

func (bottc *botTestContainer) writeDescription() {
	bottc.name = "(IBot).WriteDescription"
	bottc.inputStr = []string{"Description", "", strings.Repeat("h", 512), strings.Repeat("k", 513), "is a description", "no descriptions"}
	bottc.isExpectedErr = []bool{false, false, false, true, false, true}
	bottc.codeErr = []string{"", "", "", "20", "", "10"}
	bottc.amount, bottc.until = 6, 4
	bottc.buildF = bottc.description
}

func (bottc *botTestContainer) writeShortDescription() {
	bottc.name = "(IBot).WriteShortDescription"
	bottc.inputStr = []string{"Short Description", "", strings.Repeat("h", 120), strings.Repeat("k", 121), "is a short description", "no short descriptions"}
	bottc.isExpectedErr = []bool{false, false, false, true, false, true}
	bottc.codeErr = []string{"", "", "", "20", "", "10"}
	bottc.amount, bottc.until = 6, 4
	bottc.buildF = bottc.shortDescription
}

func (bottc *botTestContainer) writeMenuButton() {
	bottc.name = "(IBot).WriteMenuButton"
	bottc.inputMButton = []*types.MenuButton{{}, nil, {}, {}}
	bottc.isExpectedErr = []bool{false, true, false, true}
	bottc.codeErr = []string{"", "20", "", "10"}
	bottc.amount, bottc.until = 4, 2
	bottc.buildF = bottc.menuButton
}

func (bottc *botTestContainer) writeRights() {
	bottc.name = "(IBot).WriteRights"
	bottc.inputRights = []*types.ChatAdministratorRights{{}, nil, {}, {}}
	bottc.isExpectedErr = []bool{false, true, false, true}
	bottc.codeErr = []string{"", "20", "", "10"}
	bottc.amount, bottc.until = 4, 2
	bottc.buildF = bottc.rights
}

func (bottc *botTestContainer) writeForChannels() {
	bottc.name = "(IBot).WriteForChannels"
	bottc.isExpectedErr = []bool{false, false, true}
	bottc.codeErr = []string{"", "", "10"}
	bottc.amount, bottc.until = 3, 1
	bottc.buildF = bottc.forchannels
}

func (bot *botT) startTest(part string, i int, t *testing.T) {
	switch f := bot.testedFunc.(type) {
	case func(string) error:
		printTestLog(part, bot.name, bot.codeErr, bot.str, bot.isExpectedErr, i, t)
		checkError(f(bot.str), bot.isExpectedErr, bot.codeErr, t)
	case func([]*types.BotCommand) error:
		printTestLog(part, bot.name, bot.codeErr, bot.commands, bot.isExpectedErr, i, t)
		checkError(f(bot.commands), bot.isExpectedErr, bot.codeErr, t)
	case func(*types.BotCommandScope) error:
		printTestLog(part, bot.name, bot.codeErr, bot.scope, bot.isExpectedErr, i, t)
		checkError(f(bot.scope), bot.isExpectedErr, bot.codeErr, t)
	case func(*types.MenuButton) error:
		printTestLog(part, bot.name, bot.codeErr, bot.mbutt, bot.isExpectedErr, i, t)
		checkError(f(bot.mbutt), bot.isExpectedErr, bot.codeErr, t)
	case func(*types.ChatAdministratorRights) error:
		printTestLog(part, bot.name, bot.codeErr, bot.rights, bot.isExpectedErr, i, t)
		checkError(f(bot.rights), bot.isExpectedErr, bot.codeErr, t)
	case func() error:
		printTestLog(part, bot.name, bot.codeErr, true, bot.isExpectedErr, i, t)
		checkError(f(), bot.isExpectedErr, bot.codeErr, t)
	default:
		t.Fatal("unexpected type of tested function")
	}
}

func (bottc *botTestContainer) createTestArrays(msg *formatter.Message) ([]UnitTest, []UnitTest) {
	var bot formatter.IBot
	a, b := make([]UnitTest, bottc.until), make([]UnitTest, bottc.amount-bottc.until)
	for i, j := 0, 0; i < bottc.amount; i++ {
		bb := bottc.defaultBotT(i)
		if i < bottc.until {
			bottc.buildF(msg.NewBot(), bb, i)
			a[i] = bb
		} else {
			if j%2 == 0 {
				bot = msg.NewBot()
			}
			bottc.buildF(bot, bb, i)
			b[j] = bb
			j++
		}
	}
	return a, b
}

func (bottc *botTestContainer) mainLogic(msg *formatter.Message, t *testing.T) {
	botcontainer, doublecontainer := bottc.createTestArrays(msg)
	check(botcontainer, t)
	doubleCheck(doublecontainer, t)
}

func TestBotWriteCommands(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeCommands()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteScope(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeScope()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteLanguage(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeLanguage()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteName(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeName()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteDescription(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeDescription()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteShortDescription(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeShortDescription()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteMenuButton(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeMenuButton()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteRights(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeRights()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}

func TestBotWriteForChannels(t *testing.T) {
	t.Parallel()
	bottc := new(botTestContainer)
	bottc.writeForChannels()
	msg := formatter.CreateEmpltyMessage()
	bottc.mainLogic(msg, t)
}
