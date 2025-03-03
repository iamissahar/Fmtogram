package integrated

// import (
// 	"math/rand"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/l1qwie/Fmtogram/executer"
// 	"github.com/l1qwie/Fmtogram/formatter/methods"
// 	"github.com/l1qwie/Fmtogram/testbotdata"
// 	"github.com/l1qwie/Fmtogram/types"
// )

// func msgRReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.MessageReaction); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func msgRAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteReaction([]*types.ReactionType{{
// 		ReactionTypeEmoji: &types.ReactionTypeEmoji{
// 			Type:  "emoji",
// 			Emoji: "💩",
// 		},
// 	}}); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.MessageReaction); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestMessageReaction(t *testing.T) {
// 	t.Run("Req", msgRReq)
// 	t.Run("All", msgRAll)
// }

// func uppReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UserProfilePhotos); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func uppAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteLimit(8); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteOffset(7); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UserProfilePhotos); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestUserProfilePhotos(t *testing.T) {
// 	t.Run("Req", uppReq)
// 	t.Run("All", uppAll)
// }

// func uesReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UserEmojiStatus); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func uesAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteEmojiStatusCustomEmojiID(""); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteEmojiStatusExpirationDate(66); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UserEmojiStatus); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestUserEmojiStatus(t *testing.T) {
// 	t.Run("Req", uesReq)
// 	t.Run("All", uesAll)
// }

// func TestFile(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	// if err = tc.prm.WriteFileID(videodata[1]); err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.File); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if tc.get.File() == nil {
// 		t.Fatal("tc.get.File() should not be nil")
// 	}
// 	filepath := "./example.mp4"
// 	if err = executer.DownloadFile(tc.get.File().FilePath, "./example.mp4"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = os.Remove(filepath); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func banmReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(unknownid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.BanMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func banmAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(unknownid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUntilDate(time.Minute); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteRevokeMessages(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.BanMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestBanMember(t *testing.T) {
// 	t.Run("Req", banmReq)
// 	t.Run("All", banmAll)
// }

// func unbanmReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(unknownid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UnbanMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: method is available for supergroup and channel chats only" {
// 		t.Fatal(msg)
// 	}
// }

// func unbanmAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(unknownid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteOnlyIfBanned(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UnbanMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: method is available for supergroup and channel chats only" {
// 		t.Fatal(msg)
// 	}
// }

// func TestUnbanMember(t *testing.T) {
// 	t.Run("Req", unbanmReq)
// 	t.Run("All", unbanmAll)
// }

// func restReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(unknownid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WritePermissions(chatperm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.RestrictMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func restAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(unknownid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WritePermissions(chatperm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteIndependentChatPermissions(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUntilDate(time.Minute); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.RestrictMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestRestrict(t *testing.T) {
// 	t.Run("Req", restReq)
// 	t.Run("All", restAll)
// }

// func promReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.PromoteMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func promAll(t *testing.T) {
// 	var err error
// 	var b bool
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteAdministratorRights(&types.ChatAdministratorRights{CanChangeInfo: &b}); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.PromoteMember); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestPromote(t *testing.T) {
// 	t.Run("Req", promReq)
// 	t.Run("All", promAll)
// }

// func TestChAdminiTitle(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteCustomTitle(titles[rand.Intn(3)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ChatAdministratorTitle); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: not enough rights to change custom title of the user" {
// 		t.Fatal(msg)
// 	}
// }

// func TestBanSender(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteSenderChatID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.BanSenderChat); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: not enough rights to restrict/unrestrict chat member" {
// 		t.Fatal(msg)
// 	}
// }

// func chpReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WritePermissions(chatperm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ChatPermissions); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: not enough rights to change chat permissions" {
// 		t.Fatal(msg)
// 	}
// }

// func chpAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WritePermissions(chatperm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteIndependentChatPermissions(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ChatPermissions); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: not enough rights to change chat permissions" {
// 		t.Fatal(msg)
// 	}
// }

// func TestChatPermissions(t *testing.T) {
// 	t.Run("Req", chpReq)
// 	t.Run("All", chpAll)
// }

// func TestExportLink(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ExportInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if str := tc.get.String(); str == "" {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func clReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	} else {
// 		t.Log(l)
// 	}
// }

// func clAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteName(groupnames[rand.Intn(3)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteExpireDate(time.Hour); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteMemberLimit(5); err != nil {
// 		t.Fatal(err)
// 	}
// 	// if err = tc.link.WriteJoinRequest(); err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	if err = tc.msg.AddMethod(methods.CreateInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddLink(tc.link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func TestCreateLink(t *testing.T) {
// 	t.Run("Req", clReq)
// 	t.Run("All", clAll)
// }

// func edlReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteInviteLink(link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.EditInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddLink(tc.link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func edlAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteInviteLink(link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteName(groupnames[rand.Intn(3)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteExpireDate(time.Hour); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteMemberLimit(5); err != nil {
// 		t.Fatal(err)
// 	}
// 	// if err = tc.link.WriteJoinRequest(); err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	if err = tc.msg.AddMethod(methods.EditInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddLink(tc.link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func TestEditLink(t *testing.T) {
// 	t.Run("Req", edlReq)
// 	t.Run("All", edlAll)
// }

// func crSLReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteSubscriptionPeriod(month); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteSubscriptionPrice(1); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateSubInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddLink(tc.link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: PRICING_CHAT_INVALID" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func crSLAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteName(groupnames[rand.Intn(3)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteSubscriptionPeriod(month); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteSubscriptionPrice(1); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateSubInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddLink(tc.link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: PRICING_CHAT_INVALID" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func TestCreateSubLink(t *testing.T) {
// 	t.Run("Req", crSLReq)
// 	t.Run("All", crSLAll)
// }

// func TestRevokeLink(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.link.WriteInviteLink(link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.RevokeInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddLink(tc.link); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if l := tc.get.InviteLink(); l == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func TestSetChatPhoto(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	// if err = tc.ph.WritePhotoStorage(photodata[0]); err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	if err = tc.msg.AddMethod(methods.SetChatPhoto); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddPhoto(tc.ph); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestDelChatPhoto(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.DeleteChatPhoto); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestChatTitle(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteTitle(titles[rand.Intn(3)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ChatTitle); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestChatDescription(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteDescription(titles[rand.Intn(3)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ChatDescription); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestPinMessage(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.PinMessage); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestUnpinMessage(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UnpinMessage); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestLeave(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.LeaveChat); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestGetChat(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetChat); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if ch := tc.get.ChatInfo(); ch == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	}
// }

// func TestGetAdmins(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetAdmins); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if m := tc.get.Members(); m == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	} else {
// 		for _, mm := range m {
// 			t.Log(mm)
// 		}
// 	}
// }

// func TestGetMemberCount(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetMemberCount); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if m := tc.get.Integer(); m == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	} else {
// 		t.Log(m)
// 	}
// }

// // func TestChatStickerSet(t *testing.T) {
// // 	var err error
// // 	tc := new(testcase)
// // 	tc.init()
// // 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	// USAPresident
// // 	if err = tc.st.WriteSetName("Mr. Trump"); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if err = tc.msg.AddMethod(methods.ChatStickerSet); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if err = tc.msg.AddSticker(tc.st); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if err = tc.msg.AddChat(tc.ch); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	tc.send(tc.msg, t)
// // 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// // 		t.Fatal(msg)
// // 	}
// // }

// func TestGetTopicStickers(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.msg.AddMethod(methods.GetForumIconStickers); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if m := tc.get.Stickers(); m == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	} else {
// 		t.Log(m)
// 	}
// }

// func TestCreateTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.fr.WriteName(topicnames[rand.Intn(4)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.fr.WriteIconColor(iconcolor); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.fr.WriteIconEmojiID(iconemojiid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddForum(tc.fr); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateForumTopic); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if fr := tc.get.Forum(); fr == nil {
// 		t.Fatal("the response shouldn't be empty")
// 	} else {
// 		t.Log(fr)
// 	}
// }

// func TestEditTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageThreadID(10); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.fr.WriteName(topicnames[rand.Intn(4)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.fr.WriteIconEmojiID(""); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddForum(tc.fr); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.EditForumTopic); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestCloseTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageThreadID(10); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CloseForumTopic); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestReopenTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageThreadID(10); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ReopenForumTopic); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestAnswerCallBack(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.prm.WriteCallBackQueryID("asldal:sdla"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteString("SHALOM, PRAVOSLAVNIE!"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteShowAlert(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteURL("tttesty_bot_bot", "https://youtube.com"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteCacheTime(2); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.AnswerCallbackQuery); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 400 && msg != "Bad Request: query is too old and response timeout expired or query ID is invalid" {
// 		t.Fatal(msg)
// 	}
// }

// func TestGetBoosts(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetUsersBoosts); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestSetCommands(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.bot.WriteCommands(botcomm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.bot.WriteScope(botcommscope); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.bot.WriteLanguage(english); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.SetMyCommands); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddBot(tc.bot); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestGetCommands(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.bot.WriteLanguage(english); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddBot(tc.bot); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetMyCommands); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if com := tc.get.Commands(); com == nil {
// 		t.Fatal("the response shouldn't be nil")
// 	} else {
// 		t.Log(com)
// 	}
// }

// func TestSetName(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.bot.WriteName(topicnames[rand.Intn(4)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddBot(tc.bot); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.SetMyName); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func TestGetName(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.msg.AddMethod(methods.GetMyName); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// 	if name := tc.get.String(); name == "" {
// 		t.Fatal("response souldn't be empty")
// 	} else {
// 		t.Log(name)
// 	}
// }

// func TestCeateSet(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteSetName("random_name_112323_by_secondtestbotforb_bot"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteSetTitle(topicnames[rand.Intn(4)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	for i := 0; i < 2; i++ {
// 		st := tc.msg.NewSticker()
// 		// if i/2 == 0 {
// 		// if err = st.WriteStickerStorage(stickerdata[0]); err != nil {
// 		// 	t.Fatal(err)
// 		// }
// 		// } else {
// 		// if err = st.WriteStickerTelegram(stickerdata[1]); err != nil {
// 		// 	t.Fatal(err)
// 		// }
// 		// }
// 		if err = st.WriteFormat("static"); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = st.WriteAssociatedEmojies([]string{emojies[rand.Intn(len(emojies)-1)]}); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = st.WriteKeywords([]string{keywords[rand.Intn(len(keywords)-1)]}); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tc.msg.AddSticker(st); err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// 	if err = tc.prm.WriteStickerType("regular"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateNewStickerSet); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	if code, msg := tc.get.Error(); code != 0 && msg != "" {
// 		t.Fatal(msg)
// 	}
// }

// func fmsgReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[chat] = struct{}{}
// 	tc.whattocheck[sender] = struct{}{}
// 	tc.whattocheck[date] = struct{}{}
// 	tc.whattocheck[msgid] = struct{}{}
// 	tc.whattocheck[replyed] = struct{}{}
// 	tc.whattocheck[msg] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ForwardMessage); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func fmsgAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[chat] = struct{}{}
// 	tc.whattocheck[sender] = struct{}{}
// 	tc.whattocheck[date] = struct{}{}
// 	tc.whattocheck[msgid] = struct{}{}
// 	tc.whattocheck[replyed] = struct{}{}
// 	tc.whattocheck[msg] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteDisableNotification(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteProtectContent(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ForwardMessage); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func oneFMsg(t *testing.T) {
// 	t.Run("Req", fmsgReq)
// 	t.Run("All", fmsgAll)
// }

// func fmsgsReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[msgids] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ForwardMessages); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func fmsgsAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[msgids] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteDisableNotification(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteProtectContent(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ForwardMessages); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func mtoneFMsg(t *testing.T) {
// 	t.Run("Req", fmsgsReq)
// 	t.Run("All", fmsgsAll)
// }

// func TestFMessage(t *testing.T) {
// 	t.Run("OneMessage", oneFMsg)
// 	t.Run("MoreThan1Message", mtoneFMsg)
// }

// func copyMsgReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[msgid] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CopyMessage); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func (tc *testcase) copyMessage(t *testing.T) {
// 	var err error
// 	for i := 0; i < 4; i++ {
// 		tt := new(testcase)
// 		tt.init()
// 		if i == 3 {
// 			if err = tt.prm.WriteEntities(entities); err != nil {
// 				t.Fatal(err)
// 			}
// 		} else {
// 			if err = tt.prm.WriteParseMode(parsemode[i]); err != nil {
// 				t.Fatal(err)
// 			}
// 		}
// 		if err = tt.ch.WriteChatID(chatid); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.ch.WriteFromChatID(chatid); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteVideoStartTimestamp(6); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteString(textformsg); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteShowCaptionAboveMedia(); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteDisableNotification(); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteProtectContent(); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.prm.WriteReplyParameters(tc.getReplyPrm()); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.msg.AddMethod(methods.CopyMessage); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.msg.AddChat(tt.ch); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = tt.msg.AddParameters(tt.prm); err != nil {
// 			t.Fatal(err)
// 		}
// 		tc.kbF(tt, t)
// 		tc.send(tt.msg, t)
// 		tc.checkResponse(t, 0)
// 	}
// }

// func copyMsgAll(t *testing.T) {
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[msgid] = struct{}{}
// 	for i := 0; i < 3; i++ {
// 		if i == 0 {
// 			t.Log("Test Copy Message With Inline Keyboard")
// 			tc.kbF = tc.inlineKb
// 		} else if i == 1 {
// 			t.Log("Test Copy Message With ReplyMarkup Keyboard")
// 			tc.kbF = tc.replyKb
// 		} else {
// 			t.Log("Test Copy Message With ForceReply Keyboard")
// 			tc.kbF = tc.forceKb
// 		}
// 		tc.copyMessage(t)
// 	}
// }

// func copyMessage(t *testing.T) {
// 	t.Run("Req", copyMsgReq)
// 	t.Run("All", copyMsgAll)
// }

// func copyMsgsReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[msgids] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CopyMessages); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func copyMsgsAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[msgids] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteFromChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[types.BotID]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteDisableNotification(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteProtectContent(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteRemoveCaption(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CopyMessages); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	tc.send(tc.msg, t)
// 	tc.checkResponse(t, 0)
// }

// func copyMessages(t *testing.T) {
// 	t.Run("Req", copyMsgsReq)
// 	t.Run("All", copyMsgsAll)
// }

// func TestCopy(t *testing.T) {
// 	t.Run("OneMessage", copyMessage)
// 	t.Run("MoreThan1Message", copyMessages)
// }

// func TestGetUpdates(t *testing.T) {
// 	var offset int
// 	var err error
// 	types.BotID = testbotdata.TestsInGroup2_bot
// 	if err = executer.GetOffset(&offset, types.BotID); err != nil {
// 		t.Fatal(err)
// 	}
// 	for err != nil {
// 		err = executer.GetOffset(&offset, types.BotID)
// 		time.Sleep(time.Second / 10)
// 	}
// 	tg := new(types.Telegram)
// 	if err = executer.GetUpdates(tg, &offset, types.BotID); err != nil {
// 		t.Fatal(err)
// 	}
// }
