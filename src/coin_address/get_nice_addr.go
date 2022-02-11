package main

import (
	btchelper "cryptocoin/src/coin_address/btc"
	"encoding/hex"
	"fmt"
)

func GetSingleBtcResult() {
	btc := new(btchelper.Btc)
	privateKey, publicKey, err := btc.GetBtcKeyPair()
	if err != nil {
		fmt.Printf("Get btc key pair error: %s\n", err)
	}

	prikey := hex.EncodeToString(privateKey.Serialize())
	pubkey := hex.EncodeToString(publicKey.SerializeCompressed())

	fmt.Printf("prikey= %s\n,pubkey= %s\n", prikey, pubkey)

	chains := []string{"btc", "btctest", "btc_p2wpkh", "btctest_p2wpkh", "btc_p2sh-p2wpkh", "btctest_p2sh-p2wpkh"}
	chain := chains[0]
	address, err := btc.GetBitcoinAddress(publicKey, chain)
	if err != nil {
		fmt.Printf("Get btc address error: %s\n", err)
	}
	fmt.Printf("address= %s\n", address)
	btc_reuslt := &btchelper.BtcResult{
		Prikey:  prikey,
		Pubkey:  pubkey,
		Address: address,
	}
	btc.WriteJosnToFile(btc_reuslt)
}

func main1() {
	for i := 0; i < 100; i++ {
		GetSingleBtcResult()
	}
}
