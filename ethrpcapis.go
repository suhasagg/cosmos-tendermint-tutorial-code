package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Connect to the Ethereum network using an Ethereum client
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
    if err != nil {
        log.Fatal(err)
    }

    // Get the balance of an Ethereum address
    address := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
    balance, err := client.BalanceAt(context.Background(), address, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance of address %s: %s\n", address.Hex(), balance.String())

    // Get the transaction count of an Ethereum address
    nonce, err := client.NonceAt(context.Background(), address, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Nonce of address %s: %d\n", address.Hex(), nonce)

    // Get the latest block number
    blockNumber, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Latest block number: %d\n", blockNumber)

    // Get the hash of the latest block
    block, err := client.BlockByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Hash of latest block: %s\n", block.Hash().Hex())

    // Get the balance of the latest block coinbase address
    coinbase := block.Coinbase()
    balance, err = client.BalanceAt(context.Background(), coinbase, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance of coinbase address %s: %s\n", coinbase.Hex(), balance.String())

    // Get the transaction count of the latest block coinbase address
    nonce, err = client.NonceAt(context.Background(), coinbase, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Nonce of coinbase address %s: %d\n", coinbase.Hex(), nonce)

    // Estimate the gas cost of a transaction
    fromAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
    toAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
    value := big.NewInt(1000000000000000000) // 1 ETH
    gasLimit := uint64(21000)
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
    gasCost, err := client.EstimateGas(context.Background(), tx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Estimated gas cost: %d\n", gasCost)

    // Send a transaction
    privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }
    fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
    tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
    signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
    if err != nil {
        log.Fatal(err)
    }
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Sent transaction with hash: %s\n", signedTx.Hash().Hex())
}
