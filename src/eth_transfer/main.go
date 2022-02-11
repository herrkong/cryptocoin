package eth_transfer

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/getamis/sirius/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/sha3"

	pubUtils "threshold/src/public/utils"
)

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

var configFile string

var Cmd = &cobra.Command{
	Use:   "eth_transfer",
	Short: "ERC20 chain transfer",
	Long:  `eth from public key, msg, r and s.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := initService(cmd)
		if err != nil {
			log.Crit("Failed to init", "err", err)
		}

		c, err := readTransferConfigFile(configFile)
		if err != nil {
			log.Crit("Failed to read config file", "configFile", configFile, "err", err)
		}

		// 读配置
		pubKey, err := pubUtils.PubKeyFromXYString(c.Pubkey.X, c.Pubkey.Y)
		if err != nil {
			log.Error("Cannot convert config to public key", "config", c.Pubkey, "err", err)
			return nil
		}
		log.Info("Read config", "pubKey", pubKey)

		compressedPKHexString := pubUtils.PubKeyToCompressedHexString(pubKey)
		log.Info("Read config", "compressedPKHexString", compressedPKHexString)

		recoveredPK, _ := pubUtils.PubKeyFromHexString(compressedPKHexString)
		recoveredPK.Curve = pubUtils.GetCurve()
		log.Info("Read config", "recoveredPK", recoveredPK)

		pubKey = recoveredPK

		fromAddress := crypto.PubkeyToAddress(*pubKey)
		log.Info("Read config", "Fm", fromAddress.String())

		toAddress := common.HexToAddress(c.To) // 目的地址 0x71EfEC872FA6f790CfC9190a804889Df6D81be10 / 0x9f5c661f1c759478d594a0f9d193f9f44ec5938d
		log.Info("Read config", "To", toAddress.String())

		// 转的额度
		value, ok := new(big.Int).SetString(c.Value, 10)
		if !ok {
			log.Error("Cannot convert string to big int", "value", c.Value)
			return nil
		}
		log.Info("Read config", "value", value)

		defaultGasPrice, ok := new(big.Int).SetString(c.DefaultGasPrice, 10)
		if !ok {
			log.Error("Cannot convert string to big int", "default_gas_price", c.DefaultGasPrice)
			return nil
		}
		log.Info("Read config", "defaultGasPrice", defaultGasPrice)

		defaultGasLimit, err := strconv.ParseUint(c.DefaultGasLimit, 10, 64)
		if err != nil {
			log.Error("Cannot convert string to unit64", "default_gas_limit", c.DefaultGasLimit)
			return nil
		}
		log.Info("Read config", "defaultGasPrice", defaultGasLimit)

		var contractAddress common.Address
		if c.Contract != "" {
			contractAddress = common.HexToAddress(c.Contract) // erc20合约地址
			log.Info("Read config", "Tk", contractAddress.String())
		}

		/* ETH转账 start */

		// 转账
		//client, err := ethclient.Dial("https://mainnet.infura.io")
		client, err := ethclient.Dial("https://rinkeby.infura.io/v3/394965ce035a43faa213dc628059593b")
		if err != nil {
			log.Error("eth client Dial error", "error", err)
			return nil
		}

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Error("nonce error", "error", err)
			return nil
		}
		log.Info("PendingNonceAt", `nonce`, nonce)

		var data []byte

		if c.Contract == "" {
			log.Info("ETH Transfer")
		} else {
			log.Info("ERC20 Transfer")
			transferFnSignature := []byte("transfer(address,uint256)")
			erc20hash := sha3.NewLegacyKeccak256()
			erc20hash.Write(transferFnSignature)
			methodID := erc20hash.Sum(nil)[:4]
			log.Info("methodID", "methodID", methodID)

			paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
			log.Info("paddedAddress", "paddedAddress", paddedAddress)

			paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
			value.SetInt64(0) //ERC20转账，这个value要设置成0
			log.Info("paddedAmount", "paddedAmount", paddedAmount)

			data = append(data, methodID...)
			data = append(data, paddedAddress...)
			data = append(data, paddedAmount...)

			// 统一to地址
			toAddress = common.HexToAddress(contractAddress.String())
		}

		gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  fromAddress,
			To:    &toAddress,
			Value: value,
			Data:  data,
		})
		if err != nil {
			log.Error("EstimateGas error, set as default gas limit", "error", err)
			gasLimit = defaultGasLimit
		} else {
			log.Info("EstimateGas success")
		}
		log.Info("gasLimit", "gasLimit", gasLimit)

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Error("SuggestGasPrice error, set as default gas price", "error", err)
			gasPrice = defaultGasPrice
		} else {
			log.Info("SuggestGasPrice success")
		}
		log.Info("gasPrice", "gasPrice", gasPrice)

		// 组装tx
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       &toAddress,
			Value:    value,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Data:     data,
		})

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Error("NetworkID error", "error", err)
			return nil
		}

		signer := types.NewEIP155Signer(chainID)

		//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		//if err != nil {
		//	log.Fatal(err)
		//}

		// 门限签
		hash := signer.Hash(tx)
		digest := hash[:]
		digestString := hex.EncodeToString(digest)
		log.Info("Tx Hash", "Tx digest", digestString)

		// 这里用门限签把digest签名得到r、s
		var i string
		fmt.Printf("r_s_v:")
		_, _ = fmt.Scanln(&i)
		tmp := strings.Split(i, "_")
		if len(tmp) != 3 {
			log.Crit("Sign", "err", "invalid signer result")
		}
		R, _ := new(big.Int).SetString(tmp[0], 10)
		S, _ := new(big.Int).SetString(tmp[1], 10)
		V, _ := strconv.ParseUint(tmp[2], 10, 8)

		// 修复erc20签名半高
		if S.Cmp(secp256k1halfN) > 0 {
			// 1. Modify s to 0 < s < N /2
			// ref: condition 283 in https://ethereum.github.io/yellowpaper/paper.pdf
			S = new(big.Int).Sub(secp256k1N, S)
			// 2. Calculate recovery id
			// https://ethereum.stackexchange.com/questions/42455/during-ecdsa-signing-how-do-i-generate-the-recovery-id
			V = V ^ 1
		}

		// rs转成常规签名
		sign := make([]byte, 65)
		copy(sign[0:32], common.LeftPadBytes(R.Bytes(), 32))
		copy(sign[32:64], common.LeftPadBytes(S.Bytes(), 32))
		sign[64] = byte(V)

		// 自验签
		ret := crypto.VerifySignature(pubUtils.CompressPubKey(pubKey), digest, sign[:64])
		log.Info("Off line verify sign", "result", ret)
		if ret == false {
			return nil
		}

		// 附上签名
		signedTx, err := tx.WithSignature(signer, sign)
		if err != nil {
			log.Error("sign err", "sign err", err)
		} else {
			log.Info("sign success", "signTx", signedTx)
		}

		// 发起转账
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Error("SendTransaction err", "SendTransaction err", err)
			return nil
		}

		log.Info("SendTransaction success", "tx hash", signedTx.Hash().Hex())
		log.Info("SendTransaction success", "explorer", "https://rinkeby.etherscan.io/tx/"+signedTx.Hash().Hex())

		return nil
	},
}

func init() {
	Cmd.Flags().String("config", "", "transfer config file path")
}

func initService(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	configFile = viper.GetString("config")

	return nil
}
