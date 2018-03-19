package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/getlantern/systray"
)

// Response struct which contains
// last market price
type Price struct {
	MarketName string `json:"market"`
	Last       string `json:"last"`
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("assets/Raindropmemory-Legendora-Coin.ico"))
	systray.SetTooltip("Última cotação do Bitcoin na Braziliex")

	go func() {
		for {
			systray.SetTitle(getBitcoinPriceBraziliex())
			time.Sleep(60 * time.Second)
		}
	}()

}

func onExit() {
	// Clean stuffs here
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

func getBitcoinPriceBraziliex() string {
	response, err := http.Get("https://braziliex.com/api/v1/public/ticker/btc_brl")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Price
	json.Unmarshal(responseData, &responseObject)

	return `R$ ` + responseObject.Last[0:8]
}
