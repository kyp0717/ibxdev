/******

*****/
package webapi

import (
	"fmt"
	// "log"
	//"encoding/json"
	"github.com/fatih/color"
	/*"os"*/)

type IbStatus struct {
	Authenticated bool     `json:"authenticated"`
	Connected     bool     `json:"connected"`
	Competing     bool     `json:"competing"`
	Fail          string   `json:"fail"`
	Message       string   `json:"message"`
	Prompts       []string `json:"prompts"`
}

func Logout() {
	url := UrlBase + UrlLogout
	data, _ := IbGet(url)
	fmt.Println(string(data))
}

func Status() {
	url := UrlBase + UrlStatus
	data, _ := IbGet(url)
	fmt.Println(string(data))
}

func (s IbStatus) Print() {
	// use like func
	// red := color.FgRed.Render
	// kgreen := color.FgGreen.Render
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("Connection: %s \n", red(s.Connected))
	fmt.Printf("This is a %s and this is %s.\n", red("warning"), red("error"))

}
