package main

import (
	MyEth "cryptocoin/src/coin_address/eth"

	"github.com/google/martian/log"
)

func main() {
	OneEth := new(MyEth.Eth)
	PriKey, Pubkey, err := OneEth.GetKeyPair()
	if err != nil {
		log.Errorf("get eth key pair err= %s\n", err)
	}
	log.Infof("PriKey = %s,PubKey =%s\n", PriKey, Pubkey)

	Address, err := OneEth.GetAddress(PriKey, Pubkey)
	if err != nil {
		log.Errorf("get address err= %s\n", err)
	}

	log.Infof("address = %s\n", Address)

}
