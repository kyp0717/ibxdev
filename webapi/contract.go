/******

*****/
package webapi

import (
	"bytes"
	"encoding/json"
	//"fmt"
	//"github.com/fatih/color"
	"log"
)

// Account API - 8 endpoints total
// Get Brokerage Acct
const searchUrl = "/iserver/secdef/search"

func SearchSymbol(symbol string) ([]byte, error) {
	reqJson, err := json.Marshal(map[string]string{
		"symbol": symbol,
	})
	if err != nil {
		log.Fatalln(err)
	}
	url := BaseURL + searchUrl
	data, err := IbPost(url, bytes.NewBuffer(reqJson))
	return data, err

}
