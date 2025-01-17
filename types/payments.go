package types

type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

type PaidMediaPurchased struct {
	From    User   `json:"from"`
	Payload string `json:"paid_media_payload"`
}

type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	Date int    `json:"date"`
	Url  string `json:"url"`
}

type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

type RevenueWithdrawalState struct {
	*RevenueWithdrawalStatePending
	*RevenueWithdrawalStateSucceeded
	*RevenueWithdrawalStateFailed
}

type AffiliateInfo struct {
	User               User `json:"affiliate_user"`
	Chat               Chat `json:"affiliate_chat"`
	CommissionPerMille int  `json:"commission_per_mille"`
	Amount             int  `json:"amount"`
	NanostarAmount     int  `json:"nanostar_amount"`
}

type TransactionPartnerUser struct {
	Type               string         `json:"type"`
	User               User           `json:"user"`
	Affiliate          *AffiliateInfo `json:"affiliate,omitempty"`
	InvoicePayload     string         `json:"invoice_payload,omitempty"`
	SubscriptionPeriod int            `json:"subscription_period,omitempty"`
	PaidMedia          []PaidMedia    `json:"paid_media,omitempty"`
	PaidMediaPayload   string         `json:"paid_media_payload,omitempty"`
	Gift               *Gift          `json:"gift,omitempty"`
}

type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`
	SponsorUser        *User  `json:"sponsor_user,omitempty"`
	CommissionPerMille int    `json:"commission_per_mille"`
}

type TransactionPartnerFragment struct {
	Type            string                  `json:"type"`
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}

type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`
	RequestCount int    `json:"request_count"`
}

type TransactionPartnerOther struct {
	Type string `json:"type"`
}

type TransactionPartner struct {
	*TransactionPartnerUser
	*TransactionPartnerAffiliateProgram
	*TransactionPartnerFragment
	*TransactionPartnerTelegramAds
	*TransactionPartnerTelegramApi
	*TransactionPartnerOther
}

type StarTransaction struct {
	ID             string              `json:"id"`
	Amount         int                 `json:"amount"`
	NanostarAmount *int                `json:"nanostar_amount,omitempty"`
	Date           int                 `json:"date"`
	Source         *TransactionPartner `json:"source,omitempty"`
	Receiver       *TransactionPartner `json:"receiver,omitempty"`
}

type StarTransactions struct {
	ST []*StarTransaction `json:"transactions"`
}
