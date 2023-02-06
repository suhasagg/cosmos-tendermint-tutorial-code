package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	aave "github.com/aave/go-aave/contracts/go"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Aave contract on the Ethereum Mainnet
	aaveAddress := common.HexToAddress("0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9")
	aaveContract, err := aave.NewAave(aaveAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Define the address that will interact with the contract
	auth := bind.NewKeyedTransactor(privateKey)

	// Check the current Aave Token (LEND) balance for the defined address
	balance, err := aaveContract.BalanceOf(nil, auth.From)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Aave Token (LEND) balance for address %s: %s\n", auth.From.Hex(), balance)

	// Approve the Aave contract to transfer LEND from the defined address
	approvalAmount := big.NewInt(1000)
	_, err = aaveContract.Approve(auth, aaveAddress, approvalAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Approved Aave contract to transfer %d LEND from address %s\n", approvalAmount, auth.From.Hex())

	// Check the current allowance for the Aave contract to transfer LEND from the defined address
	allowance, err := aaveContract.Allowance(nil, auth.From, aaveAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Allowance for Aave contract to transfer LEND from address %s: %s\n", auth.From.Hex(), allowance)

	// Deposit LEND into the Aave protocol
	depositAmount := big.NewInt(1000)
	_, err = aaveContract.Deposit(auth, depositAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deposited %d LEND into the Aave protocol from address %s\n", depositAmount, auth.From.Hex())

	// Check the current LEND balance for the defined address in the Aave protocol
	balance, err = aaveContract.BalanceOf(nil, auth.From)
}
