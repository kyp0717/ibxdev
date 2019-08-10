/******

*****/
package webapi

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
)

// Account API - 8 endpoints total
// Get Brokerage Acct
const AcctURL = "/iserver/accounts"
const BaseURL = "https://localhost:5000/v1/portal"

type IbAccunts struct {
	Accounts []string `json:"accounts"`
	Aliases  struct {
	} `json:"aliases"`
	SelectedAccount string `json:"selectedAccount"`
}

func GetAccts2() {
	url := BaseURL + AcctURL
	data, _ := IbGet(url)
	var accts IbAccunts
	json.Unmarshal([]byte(data), &accts)
	accts.Print()
}

func GetAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbGet(url)
	return data, err
}

func (a IbAccunts) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}
