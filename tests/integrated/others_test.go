package integrated

import (
	"testing"
	"unicode/utf16"
)

// func msgRReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "message-reaction", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "message-reaction", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func msgRAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteReaction([]*fmtogram.ReactionType{{Type: "emoji", Emoji: "💩"}}); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "message-reaction", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "message-reaction", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestMessageReaction(t *testing.T) {
// 	t.Run("Req", msgRReq)
// 	t.Run("All", msgRAll)
// }

// func uppReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[profph] = struct{}{}
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UserProfilePhotos); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-user-profle-photos", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-user-profle-photos", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func uppAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[profph] = struct{}{}
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-user-profle-photos", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-user-profle-photos", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestUserProfilePhotos(t *testing.T) {
// 	t.Run("Req", uppReq)
// 	t.Run("All", uppAll)
// }

// func uesReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.UserEmojiStatus); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-user-emoji", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-user-emoji", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func uesAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-user-emoji", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-user-emoji", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestUserEmojiStatus(t *testing.T) {
// 	t.Run("Req", uesReq)
// 	t.Run("All", uesAll)
// }

// func TestFile(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[file] = struct{}{}
// 	if err = tc.prm.WriteFileID(videodata[tc.token][1]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.File); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-file", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-file", tc.get.Response())
// 	tc.checkResponse(0)
// 	filepath := "./example.mp4"
// 	if err = fmtogram.DownloadFile(tc.get.File().FilePath, "./example.mp4"); err != nil {
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
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "ban-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "ban-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func banmAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "ban-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "ban-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestBanMember(t *testing.T) {
// 	t.Run("Req", banmReq)
// 	t.Run("All", banmAll)
// }

// func unbanmReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "unban-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "unban-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func unbanmAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "unban-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "unban-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestUnbanMember(t *testing.T) {
// 	t.Run("Req", unbanmReq)
// 	t.Run("All", unbanmAll)
// }

// func restReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "restrict-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "restrict-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func restAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "restrict-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "restrict-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestRestrict(t *testing.T) {
// 	t.Run("Req", restReq)
// 	t.Run("All", restAll)
// }

// func promReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "promote-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "promote-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func promAll(t *testing.T) {
// 	var err error
// 	var b bool
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(testybotid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteAdministratorRights(&fmtogram.ChatAdministratorRights{CanChangeInfo: &b}); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "promote-member", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "promote-member", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestPromote(t *testing.T) {
// 	t.Run("Req", promReq)
// 	t.Run("All", promAll)
// }

// func TestChAdminiTitle(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-admin-title", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-admin-title", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestBanSender(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "ban-sender", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "ban-sender", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func chpReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func chpAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestChatPermissions(t *testing.T) {
// 	t.Run("Req", chpReq)
// 	t.Run("All", chpAll)
// }

// func TestExportLink(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[str] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.ExportInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func clReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateInviteLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func clAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestCreateLink(t *testing.T) {
// 	t.Run("Req", clReq)
// 	t.Run("All", clAll)
// }

// func edlReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func edlAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestEditLink(t *testing.T) {
// 	t.Run("Req", edlReq)
// 	t.Run("All", edlAll)
// }

// func crSLReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func crSLAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestCreateSubLink(t *testing.T) {
// 	t.Run("Req", crSLReq)
// 	t.Run("All", crSLAll)
// }

// func TestRevokeLink(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[invl] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestSetChatPhoto(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ph.WritePhotoStorage(photodata[tc.token][0]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.SetChatPhoto); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddPhoto(tc.ph); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-permissions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-permissions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestDelChatPhoto(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.DeleteChatPhoto); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "delete-chat-photo", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "delete-chat-photo", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestChatTitle(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-title", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-title", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestChatDescription(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-chat-description", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-chat-description", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestPinMessage(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "pin-chat-message", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "pin-chat-message", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestUnpinMessage(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "unpin-chat-message", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "unpin-chat-message", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestLeave(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.LeaveChat); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "leave-chat", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "leave-chat", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetChat(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[chinf] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetChat); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-chat", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-chat", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetAdmins(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[memb] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetAdmins); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-admins", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-admins", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetMemberCount(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[integer] = struct{}{}
// 	if err = tc.ch.WriteChatID(supergroupid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetMemberCount); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-member-count", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-member-count", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[stkrs] = struct{}{}
// 	if err = tc.msg.AddMethod(methods.GetForumIconStickers); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-topic-stickers", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-topic-stickers", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestCreateTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[frm] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "create-topic", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "create-topic", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestEditTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "edit-topic", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "edit-topic", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestCloseTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "close-topic", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "close-topic", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestReopenTopic(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "reopen-topic", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "reopen-topic", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestAnswerCallBack(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "answer-callback", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "answer-callback", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetBoosts(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[boosts] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-boosts", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-boosts", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestSetCommands(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-commands", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-commands", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetCommands(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[comm] = struct{}{}
// 	if err = tc.bot.WriteLanguage(english); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddBot(tc.bot); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetMyCommands); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-commands", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-commands", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestSetName(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.bot.WriteName(topicnames[rand.Intn(4)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddBot(tc.bot); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.SetMyName); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-name", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-name", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetName(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[str] = struct{}{}
// 	if err = tc.msg.AddMethod(methods.GetMyName); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-name", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-name", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestCeateSet(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteSetName(fmt.Sprintf("random_name_4423563_by_%s", botnames[tc.token])); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteSetTitle(topicnames[rand.Intn(4)]); err != nil {
// 		t.Fatal(err)
// 	}
// 	for i := 0; i < 2; i++ {
// 		st := tc.msg.NewSticker()
// 		// if i/2 == 0 {
// 		// if err = st.WriteStickerStorage(stickerdata[tc.token][0]); err != nil {
// 		// t.Fatal(err)
// 		// }
// 		// } else {
// 		if err = st.WriteStickerTelegram(stickerdata[tc.token][1]); err != nil {
// 			t.Fatal(err)
// 		}
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "create-sticker-set", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "create-sticker-set", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "forward-message", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "forward-message", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "forward-message", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "forward-message", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "forward-messages", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "forward-messages", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "forward-messages", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "forward-messages", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "copy-message", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "copy-message", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func cpyMsg(tc *testcase, t *testing.T, kbF func(*testcase), title string) {
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
// 		if err = tt.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
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
// 		kbF(tt)
// 		if err = tc.msg.AddToken(tc.token); err != nil {
// 			t.Fatal(err)
// 		}
// 		send(tc.msg, t)
// 		t.Logf("[TEST:%s] Request:\n%s", title, tc.get.Request())
// 		t.Logf("[TEST:%s] Response:\n%s", title, tc.get.Response())
// 		tc.checkResponse(i)
// 		tc.timetochange <- struct{}{}
// 	}
// }

// func copyMsgAll(t *testing.T) {
// 	for i := 0; i < 3; i++ {
// 		t.Run(kbnames[i], func(t *testing.T) {
// 			t.Parallel()
// 			tc := new(testcase)
// 			tc.init()
// 			defer func() { tc.workdone <- struct{}{} }()
// 			go tc.changeToken(nil, nil)
// 			tc.whattocheck[status] = struct{}{}
// 			tc.whattocheck[errr] = struct{}{}
// 			tc.whattocheck[msgid] = struct{}{}
// 			cpyMsg(tc, t, kb[i], kbnames[i])
// 		})
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
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "copy-messages", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "copy-messages", tc.get.Response())
// 	tc.checkResponse(0)
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
// 	if err = tc.prm.WriteMessageIDs(testbotdata.MessageIDs[tc.token]); err != nil {
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
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "copy-messages", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "copy-messages", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func copyMessages(t *testing.T) {
// 	t.Run("Req", copyMsgsReq)
// 	t.Run("All", copyMsgsAll)
// }

// func TestCopy(t *testing.T) {
// 	t.Run("OneMessage", copyMessage)
// 	t.Run("MoreThan1Message", copyMessages)
// }

// func TestAnswerInlineQuery(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.errmsg[0] = "Bad Request: query is too old and response timeout expired or query ID is invalid"
// 	tc.code[0] = 400
// 	tc.inline.WriteIntoArray()
// 	if err = tc.inline.WritePhoto(&fmtogram.InlineQueryResultPhoto{Type: "photo", PhotoURL: photodata[tc.token][2]}); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.inline.WriteQueryID(queryid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.inline.WriteCacheTime(22); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.inline.WriteIsPersonal(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.inline.WriteNextOffset("l"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.inline.WriteButton(&fmtogram.InlineQueryResultsButton{Text: "YOUTUBE", WebApp: &fmtogram.WebAppInfo{Url: "https://youtube.com"}}); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddInlineMode(tc.inline); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.AnswerInlineQuery); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "answer-inline", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "answer-inline", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func crinvlinkReq(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[str] = struct{}{}
// 	if err = tc.pay.WriteTitle(invtitle); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteDescription(invdescrip); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePayload(invpayload); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteCurrency(starsName); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePrices(prices); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddPayment(tc.pay); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateInvoiceLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "create-invoice-link", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "create-invoice-link", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func crinvlinkAll(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[status] = struct{}{}
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[str] = struct{}{}
// 	if err = tc.pay.WriteTitle(invtitle); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteDescription(invdescrip); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePayload(invpayload); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteCurrency(starsName); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePrices(prices); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteSubscriptionPeriod(2592000); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteProviderData("{}"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePhotoUrl(photodata[tc.token][2]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePhotoSize(12); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePhotoWidth(3); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WritePhotoHeight(6); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddPayment(tc.pay); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.CreateInvoiceLink); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "create-invoice-link", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "create-invoice-link", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestCreateInvoiceLink(t *testing.T) {
// 	t.Run("Req", crinvlinkReq)
// 	t.Run("All", crinvlinkAll)
// }

// func TestAnswerShippingQuery(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.code[0] = 400
// 	tc.errmsg[0] = "Bad Request: query is too old and response timeout expired or query ID is invalid"
// 	if err = tc.pay.WriteShippingID("VALID_SHIPPING_ID"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.pay.WriteOK(false); err != nil {
// 		t.Fatal(err)
// 	}
// 	// if err = tc.pay.WriteShippingOptions([]*fmtogram.ShippingOption{{ID: "VALID_ID", Title: invtitle, Prices: prices}}); err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	if err = tc.pay.WriteErrorMessage("ERROR_MESSAGE"); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddPayment(tc.pay); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.AnswerShippingQuery); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "answer-shipping-query", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "answer-shipping-query", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetStarTransactions(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.whattocheck[status] = struct{}{}
// 	// tc.whattocheck[stars] = struct{}{}
// 	if err = tc.prm.WriteLimit(11); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteOffset(2); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetStarTransactions); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-star-transactions", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-star-transactions", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestSetPassport(t *testing.T) {
// 	var err error
// 	var s string
// 	tc := new(testcase)
// 	tc.init()
// 	tc.code[0] = 400
// 	tc.errmsg[0] = "Bad Request: DATA_HASH_SIZE_INVALID"
// 	tc.whattocheck[errr] = struct{}{}
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteErrors([]*fmtogram.PassportElementError{{
// 		Source:    "data",
// 		Type:      "personal_details",
// 		FieldName: "field_name",
// 		DataHash:  &s,
// 		Message:   "Hi there!",
// 	},
// 	}); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.SetPassportDataErrors); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-passport", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-passport", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestSetGameScore(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.code[0] = 400
// 	tc.errmsg[0] = "Bad Request: game score can't be set"
// 	tc.whattocheck[errr] = struct{}{}
// 	tc.token = testbotdata.TestsInGroup2_bot
// 	if err = tc.game.WriteScore(66); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddGame(tc.game); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.SetGameScore); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "set-game-score", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "set-game-score", tc.get.Response())
// 	tc.checkResponse(0)
// }

// func TestGetScore(t *testing.T) {
// 	var err error
// 	tc := new(testcase)
// 	tc.init()
// 	tc.code[0] = 400
// 	tc.errmsg[0] = "Bad Request: MESSAGE_ID_INVALID"
// 	tc.whattocheck[errr] = struct{}{}
// 	// tc.whattocheck[status] = struct{}{}
// 	// tc.whattocheck[score] = struct{}{}
// 	tc.token = testbotdata.TestsInGroup2_bot
// 	if err = tc.prm.WriteUserID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.ch.WriteChatID(chatid); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.prm.WriteMessageID(testbotdata.MessageID[tc.token]); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddChat(tc.ch); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddParameters(tc.prm); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddToken(tc.token); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err = tc.msg.AddMethod(methods.GetGameHighScores); err != nil {
// 		t.Fatal(err)
// 	}
// 	send(tc.msg, t)
// 	t.Logf("[TEST:%s] Request:\n%s", "get-game-score", tc.get.Request())
// 	t.Logf("[TEST:%s] Response:\n%s", "get-game-score", tc.get.Response())
// 	tc.checkResponse(0)
// }

// // func TestGetUpdates(t *testing.T) {
// // 	var offset int
// // 	var err error
// // 	tc := new(testcase)
// // 	tc.init()
// // 	if err = executer.GetOffset(&offset, tc.token); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	for err != nil {
// // 		err = executer.GetOffset(&offset, tc.token)
// // 		time.Sleep(time.Second / 10)
// // 	}
// // 	tg := new(fmtogram.Telegram)
// // 	if err = executer.GetUpdates(tg, &offset, tc.token); err != nil {
// // 		t.Fatal(err)
// // 	}
// // }

func TestUTF16(t *testing.T) {
	text := "There is some kind of text"
	utf16Runes := utf16.Encode([]rune(text))

	t.Log("UTF-16 length:", len(utf16Runes))

	offset := 1
	length := 4
	if offset+length > len(utf16Runes) {
		t.Log("Ошибка: offset выходит за пределы текста в UTF-16")
	} else {
		t.Log("Все ок")
	}
}
