package main

import "github.com/dimitrius-ion/go_basics/crypto/api"


func main() {
	rate, err := api.GetRate("BTC")
	println(rate, err)
}
