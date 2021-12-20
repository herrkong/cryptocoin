package btc

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

type Btc struct {
}

type BtcResult struct {
	Prikey  string `json:"prikey"`
	Pubkey  string `json:"pubkey"`
	Address string `json:"address"`
}

func (b *Btc) GetBtcKeyPair() (prikey *btcec.PrivateKey, pubkey *btcec.PublicKey, err error) {
	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, nil, err
	}
	publicKey := privateKey.PubKey()
	return privateKey, publicKey, nil
}

// getnew bitcoin address
func (b *Btc) GetBitcoinAddress(pubkey *btcec.PublicKey, chain string) (address string, err error) {
	publickey := pubkey.ToECDSA()
	address, err = b.PublicKeyToAddress(publickey, chain)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (b *Btc) PublicKeyToAddress(pubKey *ecdsa.PublicKey, chain string) (address string, err error) {
	btcNetParamsMap := map[string]chaincfg.Params{
		`btc`:                 chaincfg.MainNetParams,
		`btctest`:             chaincfg.TestNet3Params,
		`btc_p2wpkh`:          chaincfg.MainNetParams,
		`btctest_p2wpkh`:      chaincfg.TestNet3Params,
		`btc_p2sh-p2wpkh`:     chaincfg.MainNetParams,
		`btctest_p2sh-p2wpkh`: chaincfg.TestNet3Params,
	}

	if netParams, ok := btcNetParamsMap[chain]; ok {
		sp := strings.Split(chain, `_`)
		if len(sp) == 1 {
			// 普通非压缩地址
			fromAddressPubKey, e := btcutil.NewAddressPubKey((*btcec.PublicKey)(pubKey).SerializeUncompressed(), &netParams)
			if e != nil {
				err = e
				return
			}
			address = fromAddressPubKey.EncodeAddress()
		} else {
			if sp[1] == `p2wpkh` {
				// P2WPKH
				p2wpkh, _ := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160((*btcec.PublicKey)(pubKey).SerializeCompressed()), &netParams)
				address = p2wpkh.String()
			} else if sp[1] == `p2sh-p2wpkh` {
				// P2SH-P2WPKH
				p2wpkScript, _ := txscript.NewScriptBuilder().
					AddOp(txscript.OP_0).
					AddData(btcutil.Hash160((*btcec.PublicKey)(pubKey).SerializeCompressed())).
					Script()
				p2wpkhNested, _ := btcutil.NewAddressScriptHash(p2wpkScript, &netParams)
				address = p2wpkhNested.String()
			} else {
				return "", fmt.Errorf("not support chain: %s", chain)
			}
		}
	}
	return address, nil
}

func (b *Btc) GetNiceAddress() (address string, times int64, err error) {
	flag := false
	times = 0
	for !flag {
		times++
		privateKey, publicKey, err := b.GetBtcKeyPair()
		if err != nil {
			fmt.Printf("Get btc key pair error: %s\n", err)
		}

		prikey := hex.EncodeToString(privateKey.Serialize())
		pubkey := hex.EncodeToString(publicKey.SerializeCompressed())

		fmt.Printf("prikey= %s\n,pubkey= %s\n", prikey, pubkey)

		chains := []string{"btc", "btctest", "btc_p2wpkh", "btctest_p2wpkh", "btc_p2sh-p2wpkh", "btctest_p2sh-p2wpkh"}
		chain := chains[0]
		address, err = b.GetBitcoinAddress(publicKey, chain)
		if err != nil {
			fmt.Printf("Get btc address error: %s\n", err)
		}
		//fmt.Printf("address= %s\n", address)
		fmt.Printf("address[0:5]= %s\n", address[0:5])
		if address[:5] == "1kong" || address[:5] == "1KONG" {
			//if  strings.EqualFold(address[:5],"1kong"){
			flag = true
		}
	}
	return address, times, nil
}

func (b *Btc) WriteFile(privkey string, pubkey string, address string) {
	result_file, _ := os.OpenFile("result.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	defer result_file.Close()
	result_file_writer := bufio.NewWriter(result_file)
	defer result_file_writer.Flush()
	result_file_writer.Write([]byte(fmt.Sprintf("private_key= %s,pubkey= %s,address= %s\n", privkey, pubkey, address)))
}

// 输出json
func (b *Btc) WriteJosnToFile(btcresult *BtcResult) {
	result_file, _ := os.OpenFile("result.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	defer result_file.Close()
	result_file_writer := bufio.NewWriter(result_file)
	defer result_file_writer.Flush()

	jsonbyte,err := json.Marshal(btcresult)
	if err != nil{
		fmt.Printf("marshal btc result err=%s\n",err)
	}
	result_file_writer.Write(jsonbyte)	
}
