package fmtogram

import "github.com/iamissahar/Fmtogram/methods"

const (
	Storage               int    = 1
	Tg                    int    = 2
	Internet              int    = 3
	constPhoto            int    = 0
	constAudio            int    = 1
	constVideo            int    = 2
	constDoc              int    = 3
	constAnim             int    = 4
	constVoice            int    = 5
	constVideoNote        int    = 6
	allowed               int    = 1
	interfaceKeyboard     string = "IKeyboard"
	interfaceParam        string = "IParameters"
	interfacePhoto        string = "IPhoto"
	interfaceVideo        string = "IVideo"
	interfaceAudio        string = "IAudio"
	interfaceDocument     string = "IDocument"
	interfaceChat         string = "IChat"
	interfaceAnimation    string = "IAnimation"
	interfaceVoice        string = "IVoice"
	interfaceInKB         string = "Inline Keyboard"
	interfaceReplyKB      string = "Reply Keyboard"
	interfaceForceReplyKB string = "ForceReply Keyobard"
	interfaceVideoNote    string = "IVideoNote"
	interfaceLocation     string = "ILocation"
	interfaceContact      string = "IConatact"
	interfacePoll         string = "IPoll"
	interfaceLink         string = "ILink"
	interfaceSticker      string = "ISticker"
	interfaceForum        string = "IForum"
	interfaceBot          string = "IBot"
	button                string = "Button"
	inbtn                 string = "Inline Button"
	rpbtn                 string = "Reply Button"
	interfaceInlineMode   string = "IInlineMode"
	interfaceResult       string = "IResult"
	interfacePayment      string = "IPayment"
	interfaceGame         string = "IGame"
	interfaceGet          string = "IGet"
	checkString           int    = -1
	checkArray            int    = -2
	checkBool             int    = -3
	checkInt              int    = -4
	wURL                  int    = 0
	wCallback             int    = 1
	wWebApp               int    = 2
	wLoginUrl             int    = 3
	wSwitchIn             int    = 4
	wSwitchInQuery        int    = 5
	wSwitchInQueryCh      int    = 6
	wGame                 int    = 7
	wPay                  int    = 8
	wReqUsers             int    = 0
	wReqChat              int    = 1
	wReqContact           int    = 2
	wReqLocation          int    = 3
	wReqPoll              int    = 4
	wReqWebApp            int    = 5
	added                 int    = 1
	Month                 int    = 2592000
	telegramStars         string = "XTR"
)

type Currency struct {
	Min float64
	Max float64
}

var Currencies = map[string]Currency{
	"AED": {3.67, 36729.90},
	"AFN": {73.33, 733288.06},
	"ALL": {95.25, 952542.67},
	"AMD": {400.37, 4003726.20},
	"ARS": {1041.26, 10412550.54},
	"AUD": {1.61, 16099.43},
	"AZN": {1.70, 17002.06},
	"BAM": {1.90, 18978.40},
	"BDT": {121.32, 1213247.42},
	"BGN": {1.90, 19012.20},
	"BHD": {0.377, 3768.830},
	"BND": {1.36, 13648.14},
	"BOB": {6.90, 69002.66},
	"BRL": {6.06, 60560.96},
	"BYN": {3.27, 32675.79},
	"CAD": {1.44, 14404.20},
	"CHF": {0.91, 9115.95},
	"CLP": {1014, 10135600},
	"CNY": {7.33, 73283.00},
	"COP": {4341.75, 43417500.00},
	"CRC": {501.10, 5011011.61},
	"CZK": {24.59, 245873.00},
	"DKK": {7.25, 72519.65},
	"DOP": {61.20, 611993.44},
	"DZD": {135.57, 1355741.00},
	"EGP": {50.39, 503927.01},
	"ETB": {127.66, 1276572.46},
	"EUR": {0.97, 9720.30},
	"GBP": {0.82, 8207.89},
	"GEL": {2.84, 28404.74},
	"GHS": {14.88, 148776.37},
	"GTQ": {7.71, 77085.80},
	"HKD": {7.79, 77855.95},
	"HNL": {25.40, 253970.56},
	"HRK": {7.38, 73795.48},
	"HUF": {402.04, 4020350.22},
	"IDR": {16391.20, 163912000.00},
	"ILS": {3.61, 36115.30},
	"INR": {86.61, 866114.98},
	"IQD": {1310000, 13100000000},
	"IRR": {42087.51, 420875081.19},
	"ISK": {141, 1410900},
	"JMD": {157.05, 1570537.39},
	"JOD": {0.709, 7091.000},
	"JPY": {156, 1555105},
	"KES": {129.50, 1294975.94},
	"KGS": {87.45, 874494.99},
	"KRW": {1459, 14586950},
	"KZT": {529.30, 5292993.20},
	"LBP": {89061.59, 890615892.75},
	"LKR": {295.86, 2958556.92},
	"MAD": {10.04, 100385.23},
	"MDL": {18.73, 187323.52},
	"MMK": {3247.96, 32479609.92},
	"MNT": {3398.00, 33980001.07},
	"MOP": {8.01, 80096.45},
	"MUR": {46.95, 469503.41},
	"MVR": {15.41, 154050.09},
	"MXN": {20.89, 208855.85},
	"MYR": {4.50, 45030.55},
	"MZN": {63.89, 638914.30},
	"NGN": {1556.42, 15564202.47},
	"NIO": {36.74, 367420.20},
	"NOK": {11.38, 113776.50},
	"NPR": {138.24, 1382398.88},
	"NZD": {1.79, 17866.71},
	"PAB": {1.00, 9985.06},
	"PEN": {3.76, 37577.99},
	"PHP": {58.60, 585990.21},
	"PKR": {278.37, 2783712.56},
	"PLN": {4.15, 41486.00},
	"PYG": {7869, 78691261},
	"QAR": {3.64, 36352.79},
	"RON": {4.84, 48386.02},
	"RSD": {113.89, 1138860.04},
	"RUB": {107.57, 1075719.60},
	"SAR": {3.75, 37520.06},
	"SEK": {11.17, 111688.70},
	"SGD": {1.37, 13668.85},
	"SYP": {13002.00, 130019999.85},
	"THB": {34.45, 344464.97},
	"TJS": {10.88, 108840.03},
	"TRY": {35.56, 355566.90},
	"TTD": {6.78, 67816.84},
	"TWD": {32.93, 329258.04},
	"TZS": {2525.00, 25249998.35},
	"UAH": {42.09, 420889.08},
	"UGX": {3682, 36815650},
	"USD": {1.00, 10000.00},
	"UYU": {43.91, 439090.19},
	"UZS": {12935.81, 129358109.75},
	"VND": {25330, 253300000},
	"YER": {249.17, 2491749.63},
	"ZAR": {18.75, 187492.35},
	"XTR": {},
}

var simpleResponse = map[string]*SimpleResponse{methods.MessageReaction: new(SimpleResponse),
	methods.UserEmojiStatus:              new(SimpleResponse),
	methods.BanMember:                    new(SimpleResponse),
	methods.UnbanMember:                  new(SimpleResponse),
	methods.RestrictMember:               new(SimpleResponse),
	methods.PromoteMember:                new(SimpleResponse),
	methods.ChatAdministratorTitle:       new(SimpleResponse),
	methods.BanSenderChat:                new(SimpleResponse),
	methods.UnbanSenderChat:              new(SimpleResponse),
	methods.ChatPermissions:              new(SimpleResponse),
	methods.ApproveJoinReq:               new(SimpleResponse),
	methods.DeclineJoinReq:               new(SimpleResponse),
	methods.SetChatPhoto:                 new(SimpleResponse),
	methods.DeleteChatPhoto:              new(SimpleResponse),
	methods.ChatTitle:                    new(SimpleResponse),
	methods.ChatDescription:              new(SimpleResponse),
	methods.PinMessage:                   new(SimpleResponse),
	methods.UnpinMessage:                 new(SimpleResponse),
	methods.UnpinAll:                     new(SimpleResponse),
	methods.LeaveChat:                    new(SimpleResponse),
	methods.ChatStickerSet:               new(SimpleResponse),
	methods.DeleteChatStickerSet:         new(SimpleResponse),
	methods.EditForumTopic:               new(SimpleResponse),
	methods.CloseForumTopic:              new(SimpleResponse),
	methods.ReopenForumTopic:             new(SimpleResponse),
	methods.DeleteForumTopic:             new(SimpleResponse),
	methods.UnpinAllForumMessages:        new(SimpleResponse),
	methods.EditGeneralForum:             new(SimpleResponse),
	methods.CloseGeneralForum:            new(SimpleResponse),
	methods.ReopenGeneralForum:           new(SimpleResponse),
	methods.HideGeneralForum:             new(SimpleResponse),
	methods.UnhideGeneralForum:           new(SimpleResponse),
	methods.UnpinAllGeneralForumMessages: new(SimpleResponse),
	methods.AnswerCallbackQuery:          new(SimpleResponse),
	methods.SetMyCommands:                new(SimpleResponse),
	methods.DeleteMyCommands:             new(SimpleResponse),
	methods.SetMyName:                    new(SimpleResponse),
	methods.SetMyDescription:             new(SimpleResponse),
	methods.SetChatMenuButton:            new(SimpleResponse),
	methods.SetMyDefaultAdminRights:      new(SimpleResponse),
	methods.DeleteMessage:                new(SimpleResponse),
	methods.DeleteMessages:               new(SimpleResponse),
	methods.AddStickerToSet:              new(SimpleResponse),
	methods.SetStickerPositionInSet:      new(SimpleResponse),
	methods.DeleteStickerFromSet:         new(SimpleResponse),
	methods.ReplaceStickerInSet:          new(SimpleResponse),
	methods.SetStickerEmojiList:          new(SimpleResponse),
	methods.SetStickerKeywords:           new(SimpleResponse),
	methods.SetStickerMaskPosition:       new(SimpleResponse),
	methods.SetStickerSetTitle:           new(SimpleResponse),
	methods.SetStickerSetThumbnail:       new(SimpleResponse),
	methods.SetEmojiStickerSetThumbnail:  new(SimpleResponse),
	methods.DeleteStickerSet:             new(SimpleResponse),
	methods.Gift:                         new(SimpleResponse),
	methods.VerifyUser:                   new(SimpleResponse),
	methods.VerifyChat:                   new(SimpleResponse),
	methods.RemoveUserVerification:       new(SimpleResponse),
	methods.RemoveChatVerification:       new(SimpleResponse),
	methods.RefundStarPayment:            new(SimpleResponse),
	methods.EditUserStarSubscription:     new(SimpleResponse),
	methods.SetPassportDataErrors:        new(SimpleResponse),
	methods.AnswerInlineQuery:            new(SimpleResponse),
	methods.AnswerShippingQuery:          new(SimpleResponse),
}

var messageResponse = map[string]*MessageResponse{methods.Message: new(MessageResponse),
	methods.ForwardMessage:   new(MessageResponse),
	methods.Photo:            new(MessageResponse),
	methods.Audio:            new(MessageResponse),
	methods.Document:         new(MessageResponse),
	methods.Video:            new(MessageResponse),
	methods.Animation:        new(MessageResponse),
	methods.Voice:            new(MessageResponse),
	methods.VideoNote:        new(MessageResponse),
	methods.PaidMedia:        new(MessageResponse),
	methods.Location:         new(MessageResponse),
	methods.Venue:            new(MessageResponse),
	methods.Contact:          new(MessageResponse),
	methods.Poll:             new(MessageResponse),
	methods.Dice:             new(MessageResponse),
	methods.EditText:         new(MessageResponse),
	methods.EditCaption:      new(MessageResponse),
	methods.EditMedia:        new(MessageResponse),
	methods.EditLiveLocation: new(MessageResponse),
	methods.StopLiveLocation: new(MessageResponse),
	methods.EditReplyMarkup:  new(MessageResponse),
	methods.Sticker:          new(MessageResponse),
	methods.Invoice:          new(MessageResponse),
	methods.Game:             new(MessageResponse),
	methods.SetGameScore:     new(MessageResponse),
}

var messageIDsReponse = map[string]*MessageIDsResponse{methods.ForwardMessages: new(MessageIDsResponse),
	methods.CopyMessages: new(MessageIDsResponse),
}

var inviteLinkResponse = map[string]*InviteLinkResponse{methods.CreateInviteLink: new(InviteLinkResponse),
	methods.EditInviteLink:      new(InviteLinkResponse),
	methods.CreateSubInviteLink: new(InviteLinkResponse),
	methods.EditSubInviteLink:   new(InviteLinkResponse),
	methods.RevokeInviteLink:    new(InviteLinkResponse),
}
