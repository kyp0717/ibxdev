/******

*****/
package accounts

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
)

func GetAccts_xxx() {
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
