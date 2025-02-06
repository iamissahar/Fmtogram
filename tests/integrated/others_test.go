package integrated

import (
	"testing"

	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func msgRReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.MessageReaction); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func msgRAll(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.ch.WriteChatID(738070596); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteMessageID(testbotdata.MessageID[types.BotID]); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteReaction([]*types.ReactionType{{
		ReactionTypeEmoji: &types.ReactionTypeEmoji{
			Type:  "emoji",
			Emoji: "💩",
		},
	}}); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.MessageReaction); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddChat(tc.ch); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func TestMessageReaction(t *testing.T) {
	t.Run("Req", msgRReq)
	t.Run("All", msgRAll)
}

func uppReq(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.prm.WriteUserID(8033103339); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.UserProfilePhotos); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func uppAll(t *testing.T) {
	var err error
	tc := new(testcase)
	tc.init()
	if err = tc.prm.WriteUserID(8033103339); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteLimit(8); err != nil {
		t.Fatal(err)
	}
	if err = tc.prm.WriteOffset(7); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddMethod(methods.UserProfilePhotos); err != nil {
		t.Fatal(err)
	}
	if err = tc.msg.AddParameters(tc.prm); err != nil {
		t.Fatal(err)
	}
	send(tc.msg, t)
	if code, msg := tc.get.Error(); code != 0 && msg != "" {
		t.Fatal(msg)
	}
}

func TestUserProfilePhotos(t *testing.T) {
	t.Run("Req", uppReq)
	t.Run("All", uppAll)
}

func uesReq(t *testing.T) {

}

func uesAll(t *testing.T) {

}

func TestUserEmojiStatus(t *testing.T) {
	t.Run("Req", uesReq)
	t.Run("All", uesAll)
}
