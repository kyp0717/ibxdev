/******

*****/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	// "log"
	"crypto/tls"
	"encoding/json"
	"github.com/fatih/color"
	"net/http"
)

type IbStatus struct {
	Authenticated bool     `json:"authenticated"`
	Connected     bool     `json:"connected"`
	Competing     bool     `json:"competing"`
	Fail          string   `json:"fail"`
	Message       string   `json:"message"`
	Prompts       []string `json:"prompts"`
}

// getstatusCmd represents the getstatus command
var chkStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "check status of connection",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		chkStatus()
	},
}

func init() {
	rootCmd.AddCommand(chkStatusCmd)
}

func chkStatus() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	ibURL := BaseURL + "iserver/auth/status"

	resp, err := http.Get(ibURL)

	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	var status IbStatus
	json.Unmarshal([]byte(data), &status)
	status.Print()

}

func (s IbStatus) Print() {
	// use like func
	// red := color.FgRed.Render
	// kgreen := color.FgGreen.Render
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("Connection: %s \n", red(s.Connected))
	fmt.Printf("This is a %s and this is %s.\n", red("warning"), red("error"))

}
