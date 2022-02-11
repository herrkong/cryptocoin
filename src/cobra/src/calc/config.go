package calc

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Num1 string `yaml:"num1"`
	Num2 string `yaml:"num2"`
}

func readConfigFile(filePath string) (*Config, error) {
	c := &Config{}
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
