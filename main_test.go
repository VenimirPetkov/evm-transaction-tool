package main

import (
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestMain(t *testing.T) {
	main()
}

func TestETHTransfer(t *testing.T) {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/YOUR-INFURA-KEY-HERE")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("YOUR-PRIVATE-KEY-HERE")
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

	gasPrice := GetSuggestedGasPrice(client)

	fromAddress := ExtractAddressFromPrivateKey(privateKey)
	tx := BuildTransaction(client, fromAddress, toAddress, value, gasLimit, gasPrice)

	chainID, err := getChainId(client)

	signedTx := SignTransaction(tx, chainID, privateKey)

	err = SendSignedTransaction(client, signedTx)

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func TestGetEVMAccountFromMnemonic(t *testing.T) {
	conf, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	accountIndex := 0
	pk, addr := GetEVMAccountFromMnemonic(conf.GetMnemonic(), uint64(accountIndex))
	fmt.Println("public key", addr)
	fmt.Printf("priv key %x\n", crypto.FromECDSA(pk))
}
