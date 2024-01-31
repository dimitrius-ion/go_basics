package main

import (
	"github.com/dimitrius-ion/go_basics/crypto/api"
	"fmt"
	 "sync"
)


func main() {
	currencies := []string{"BTC", "ETH",  "BCH"}

	// like a counter
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)

		// lambda function runs in a separate thread
		go func(currency string) {
			defer wg.Done()
			getCurencyData(currency)
		}(currency)
	}
	wg.Wait()
}

func getCurencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("Rate for %v: %.2f\n", rate.Currency, rate.Price)
	}
}