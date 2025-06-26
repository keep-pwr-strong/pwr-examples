package sdk

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pwrlabs/pwrgo/rpc"
)

// Setting up the rpc api
var pwr = rpc.SetRpcNodeUrl("https://pwrrpc.pwrlabs.io/")

func getVidaDataActive() {
	startBlock := 40635
	endBlock := 40726
	vidaId := 123

	// fetch the transactions sent from `startBlock` to `endBlock` in `vidaId`
	transactions := pwr.GetVidaDataTransactions(startBlock, endBlock, vidaId)

	for _, tx := range transactions {
		sender := tx.Sender
		data := tx.Data

		// Decode the hexadecimal data to bytes data
		decodedData, err := hex.DecodeString(data)
		if err != nil {
			fmt.Println("Error decoding hex:", err)
			continue
		}

		// Convert the decoded data to a UTF-8 string
		stringData := string(decodedData)

		if strings.HasPrefix(stringData, "Hi") {
			word := stringData[2:]
			fmt.Printf("%s: %s\n", sender, word)
		} else if strings.HasPrefix(stringData, "Hello") {
			word := stringData[5:]
			fmt.Printf("%s: %s\n", sender, word)
		}
	}
}

func decoding() {
	hexData := "0x48656C6C6F20576F726C6421" // hex data

	// Remove the '0x' prefix and decode the hexadecimal data to bytes data
	decodedData, _ := hex.DecodeString(hexData[2:])
	// Convert the decoded data to a UTF-8 string
	stringData := string(decodedData)

	fmt.Printf("Outputs: %s\n", stringData) // Outputs: Hello World!
}

func getVidaData() {
	startBlock := 843500
	endBlock := 843750
	vidaId := 123

	// fetch the transactions sent from `startBlock` to `endBlock` in `vidaId`
	transactions := pwr.GetVidaDataTransactions(startBlock, endBlock, vidaId)

	for _, tx := range transactions {
		fmt.Println("Data:", tx.Data)
	}
}

func getBlock() {
	// the block number we want fetch
	blockNumber := 100

	// get the block by number
	block := pwr.GetBlockByNumber(blockNumber)

	for i, txs := range block.Transactions {
		transaction := pwr.GetTransactionByHash(txs.TransactionHash)
		fmt.Printf("Sender %d: %s\n", i, transaction.Sender)
	}
}

func account() {
	address := "0xA4710E3D79E1ED973AF58E0F269E9B21DD11BC64"

	// get balance of address
	balance := pwr.GetBalanceOfAddress(address)
	fmt.Println("Balance:", balance)

	// get nonce of address
	nonce := pwr.GetNonceOfAddress(address)
	fmt.Println("Nonce:", nonce)
}

func ReadData() {
	// account()
	// getBlock()
	// getVidaData()
	// decoding()
	getVidaDataActive()
}
