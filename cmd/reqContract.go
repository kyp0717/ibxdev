/*
 */
package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"io/ioutil"
	// "log"
	"crypto/tls"
	"encoding/json"
	"net/http"
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

//const baseURL = "https://localhost:5000/v1/portal/"
//const endpt = "iserver/accounts"

// getContractInfoCmd represents the getContractInfo command
var reqContractCmd = &cobra.Command{
	Use:   "reqContract",
	Short: "A brief description of your command",
	Long:  `A longer description to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conid, _ := cmd.Flags().GetString("conid")
		if conid == "" {
			reqContract("242506861")

		} else {
			reqContract(conid)
		}
	},
}

func init() {
	rootCmd.AddCommand(reqContractCmd)

	reqContractCmd.Flags().String("conid", "242506861", "specify Contract ID for instrument")
}

func reqContract(conid string) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url := UrlBase + "iserver/contract/" + conid + "/info"
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	var ctx ContractInfo

	json.Unmarshal([]byte(data), &ctx)

	fmt.Printf("%+v\n", ctx)
	spew.Dump(ctx)

}
