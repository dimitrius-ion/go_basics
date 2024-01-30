package main

import (
	"github.com/dimitrius-ion/go_basics/crypto/api"
	"fmt"
	"time"
)


func main() {
	go getCurencyData("BTC")
	go getCurencyData("ETH")
	time.Sleep(5 * time.Second)
}

func getCurencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("Rate for %v: %.2f\n", rate.Currency, rate.Price)
	}
}