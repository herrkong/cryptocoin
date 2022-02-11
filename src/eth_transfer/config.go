package eth_transfer

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"threshold/src/tss_tools/common/config"
)

type TransferConfig struct {
	Pubkey          config.Pubkey `yaml:"pubkey"`
	To              string        `yaml:"to"`
	Value           string        `yaml:"value"`
	Contract        string        `yaml:"contract"`
	DefaultGasPrice string        `yaml:"default_gas_price"`
	DefaultGasLimit string        `yaml:"default_gas_limit"`
}

func readTransferConfigFile(filePath string) (*TransferConfig, error) {
	c := &TransferConfig{}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
