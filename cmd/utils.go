package cmd

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

type Wallet struct {
	account types.Account
	c       *client.Client
}

func CreateNewWallet(RPCEndpoint string) Wallet {
	newAccount := types.NewAccount()
	data := []byte(newAccount.PrivateKey)
	err := ioutil.WriteFile("data", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return Wallet{
		newAccount,
		client.NewClient(RPCEndpoint),
	}
}

func ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {

	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}
	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func GetBalance(privateKey []byte) (uint64, error) {
	wallet, _ := ImportOldWallet(privateKey, string(rpc.DevnetRPCEndpoint))
	balance, err := wallet.c.GetBalance(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
	)
	if err != nil {
		return 0, nil
	}
	return balance, nil
}

func RequestAirDrop(privateKey []byte, amount uint64) (string, error) {
	wallet, _ := ImportOldWallet(privateKey, rpc.DevnetRPCEndpoint)
	amount = amount * 1e9

	txHash, err := wallet.c.RequestAirdrop(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
		amount,
	)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func Transfer(privateKey []byte, receiver string, amount uint64) (string, error) {
	wallet, _ := ImportOldWallet(privateKey, rpc.DevnetRPCEndpoint)

	response, err := wallet.c.GetRecentBlockhash(context.TODO())
	if err != nil {
		return "", err
	}

	message := types.NewMessage(
		wallet.account.PublicKey,
		[]types.Instruction{
			sysprog.Transfer(
				wallet.account.PublicKey,
				common.PublicKeyFromString(receiver),
				amount,
			),
		},
		response.Blockhash,
	)

	tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account})
	if err != nil {
		return "", err
	}

	txHash, err := wallet.c.SendTransaction2(context.TODO(), tx)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
