package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestMain(t *testing.T) {
	main()
}

func TestETHTransfer(t *testing.T) {
}

func TestGetEVMAccountFromMnemonic(t *testing.T) {
	conf, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	accountIndex := 0
	pk, addr, err := GetEVMAccountFromMnemonic(conf.GetMnemonic(), uint64(accountIndex))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("public key", addr)
	fmt.Printf("priv key %x\n", crypto.FromECDSA(pk))
}
