/******

*****/
package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ibxdev/internal"
	//"github.com/fatih/color"
	"log"
)

type ContractInfo struct {
	RTH            bool   `json:"r_t_h"`
	ConID          string `json:"con_id"`
	Company_Name   string `json:"company_name"`
	Exchange       string `json:"exchange"`
	LocalSymbol    string `json:"local_symbol"`
	InstrumentType string `json:"instrument_type"`
	Currency       string `json:"currency"`
	CompanyName    string `json:"companyname"`
	Category       string `json:"category"`
	Industry       string `json:"industry"`
	Rules          struct {
		OrderTypes        []string `json:"orderTypes"`
		OrderTypesOutside []string `json:"orderTypesOutside"`
		DefaultSize       int      `json:"defaultSize"`
		SizeIncrement     int      `json:"sizeIncrement"`
		TifTypes          []string `json:"tifTypes"`
		LimitPrice        int      `json:"limitPrice"`
		Stopprice         int      `json:"stopprice"`
		Preview           bool     `json:"preview"`
		DisplaySize       string   `json:"displaySize"`
		Increment         string   `json:"increment"`
	} `json:"rules"`
}

func SearchSymbol(symbol string) ([]byte, error) {
	reqJson, err := json.Marshal(map[string]string{
		"symbol": symbol,
	})
	if err != nil {
		log.Fatalln(err)
	}
	url := UrlBase + UrlCtxSearch
	data, err := IbPost(url, bytes.NewBuffer(reqJson))
	return data, err

}

func GetByConID(conid string) ContractInfo {
	url := UrlBase + "iserver/contract/" + conid + "/info"
	data, err := IbGet(url)
	internal.CheckErr(err)

	var ctx ContractInfo
	json.Unmarshal([]byte(data), &ctx)
	fmt.Printf("%+v\n", ctx)
	//spew.Dump(ctx)
	return ctx
}
