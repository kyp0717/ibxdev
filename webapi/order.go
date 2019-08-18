package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
)

type IbMessage struct {
	id      string   `json:id`
	message []string `json:message`
}

type IbMessages struct {
	Messages []IbMessage `json:messages`
}

func PlaceOrder(fname string) ([]byte, error) {
	// get the account num
	acctnum := GetAcctNum()
	// read the yaml file from cmd line
	GetOrder(fname)
	reqJson, err := json.Marshal(map[string]string{
		"ticker": viper.GetString("ticker"),
		"tif":    viper.GetString("tif"),
		"price":  viper.GetString("price"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	url := UrlBase + UrlAccount + acctnum + "/order"
	data, err := IbPost(url, bytes.NewBuffer(reqJson))
	return data, err

}

func GetOrder(fname string) {
	viper.SetConfigName(fname)                 // name of config file (without extension)
	viper.AddConfigPath("$HOME/go/src/ibxdev") // call multiple times to add many search paths
	viper.AddConfigPath(".")                   // optionally look for config in the working directory
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}
