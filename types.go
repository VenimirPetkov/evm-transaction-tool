package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	INFURA_KEY  string `json:"INFURA_KEY"`
	PRIVATE_KEY string `json:"PRIVATE_KEY"`
	MNEMONIC    string `json:"MNEMONIC"`
}

func LoadConfig() (*Config, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
func (c *Config) GetMnemonic() string {
	return c.MNEMONIC
}
func (c *Config) GetInfura() string {
	return c.INFURA_KEY
}
