/******

*****/
package accounts

import (
	"fmt"
	"io/ioutil"
	// "log"
	"crypto/tls"
	"encoding/json"
	"github.com/fatih/color"
	"net/http"
)

func setPolicy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func IbGet(url) ([]byte, error) {
	setPolicy()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

func IbPost(url, reqjson) ([]byte, error) {
	setPolicy()
	resp, err := http.Post(url, "application/json", reqjson)
	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

func (a IbAccunts) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}
