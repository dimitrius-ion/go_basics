package main

import (
	"github.com/dimitrius-ion/go_basics/crypto/api"
	"fmt"
)


func main() {
	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("Rate for %v: %.2f\n", rate.Currency, rate.Price)
	}
}
