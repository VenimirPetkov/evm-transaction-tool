package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
}

func GetPrivateKey(pk string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetClient(inf string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(inf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ExtractAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress, nil
}

func GetSuggestedGasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, err
}

func BuildTransaction(client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	return tx, err
}

func GetChainId(client *ethclient.Client) (*big.Int, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	return chainID, err
}

func SendSignedTransaction(client *ethclient.Client, signedTx *types.Transaction) error {
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}
	return nil
}

func SignTransaction(tx *types.Transaction, chainID *big.Int, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, err
}

func GetEVMAccountFromMnemonic(mnemonic string, accountIndex uint64) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := generateEthPrivateKeyFromMnemonic(mnemonic, accountIndex)
	if err != nil {
		return nil, common.Address{}, err
	}
	fromAddress, err := generatePublicKeyFromPrivateKey(privateKey)
	if err != nil {
		return nil, common.Address{}, err
	}
	return privateKey, fromAddress, err
}

func generatePublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress, nil
}

func generateEthPrivateKeyFromMnemonic(mnemonic string, accountIndex uint64) (*ecdsa.PrivateKey, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.FormatUint(accountIndex, 10))
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	return wallet.PrivateKey(account)
}
