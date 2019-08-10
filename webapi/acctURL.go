package accounts

const BaseURL = "https://localhost:5000/v1/portal"

// Account API - 8 endpoints total
// Get Brokerage Acct
const AcctURL = "/iserver/accounts"

type IbAccunts struct {
	Accounts []string `json:"accounts"`
	Aliases  struct {
	} `json:"aliases"`
	SelectedAccount string `json:"selectedAccount"`
}

// Post - Select Acct
