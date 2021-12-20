package main

import (
	btchelper "cryptocoin/src/coin_address/btc"
	"fmt"
)


func main() {
	btc := new(btchelper.Btc)
	
	address, times, err := btc.GetNiceAddress()
	if err != nil {
		fmt.Printf("Get btc nice address failed error: %s\n", err)
	}
	fmt.Printf("nice address= %s,times= %d\n", address, times)

}

// prikey= e9c2bcde05d8fe15a4189f15e46d0a66ce5557fc5b9ce1a8115a6c1ec3ab5d43
// ,pubkey= 037dd75890832bf6f38538c6bac9ef5195d5b9f62e0b7bd06bae1b6da0befff47b
// address[0:5]= 1KoNg
// nice address= 1KoNgu9pztbGaJXseQVp9gtMH5mUqtVscj,times= 193617

