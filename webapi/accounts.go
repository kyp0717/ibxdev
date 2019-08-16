/******

*****/
package webapi

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
	"os"
)

// Account API - 8 endpoints total
// Get Brokerage Acct

var AcctNum string

type IbAccounts struct {
	Accounts []string `json:"accounts"`
	Aliases  struct {
	} `json:"aliases"`
	SelectedAccount string `json:"selectedAccount"`
}

func GetAccts() IbAccounts {
	url := UrlBase + UrlAccounts
	data, _ := IbGet(url)
	fmt.Println(string(data))
	var accts IbAccounts
	json.Unmarshal([]byte(data), &accts)
	os.Setenv("IBACCOUNT", accts.Accounts[0])
	AcctNum = accts.Accounts[0]
	return accts
}

func GetAccts_old() ([]byte, error) {
	url := UrlBase + UrlAccounts
	fmt.Println(url)
	data, err := IbGet(url)
	return data, err
}

func (a IbAccounts) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println("--- Account Information --- ")
	fmt.Printf("--- Accounts: %s \n", cyan(a.Accounts[0]))
}
