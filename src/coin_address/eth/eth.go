package eth

import "github.com/ethereum/go-ethereum/crypto"

type Eth struct{}

func (e *Eth) GetKeyPair() (PriKey, PubKey string, err error) {
	priv, err := crypto.GenerateKey()
	
	return PriKey, PubKey, nil
}

func (e *Eth) GetAddress(PriKey string, PubKey string) (Address string, err error) {

	return Address, nil
}

func (e *Eth) BuildTransaction() (UnsignedTransaction string, err error) {

	return UnsignedTransaction, nil
}

func (e *Eth) SignTransaction() (SignedTransaction string, err error) {

	return SignedTransaction, nil
}

func (e *Eth) BroadCast(SignedTransaction string) (txid string, err error) {

	return txid, nil
}

func (e *Eth) DecodeTransaction() (err error) {

	return nil
}
