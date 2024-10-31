package sdk

import (
	"github.com/pwrlabs/pwrgo/rpc"
	"fmt"
	"encoding/hex"
	"strings"
)

func getVmDataActive() {
	startBlock := 843500
    endBlock := 843750
    vmId := 123

    // fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
	transactions := rpc.GetVmDataTransactions(startBlock, endBlock, vmId)

	for _, tx := range transactions {
        sender := tx.Sender
        data := tx.Data

        // Remove the '0x' prefix and decode the hexadecimal data to bytes data
        decodedData, err := hex.DecodeString(data[2:])
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

func getVmData() {
	startBlock := 843500
    endBlock := 843750
    vmId := 123

    // fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
	transactions := rpc.GetVmDataTransactions(startBlock, endBlock, vmId)

	for _, tx := range transactions {
		fmt.Println("Data:", tx.Data)
	}
}

func getBlock() {
    // the block number we want fetch
	blockNumber := 10

    // get the block by number
	block := rpc.GetBlockByNumber(blockNumber)

	for i, transaction := range block.Transactions {
		fmt.Printf("Sender %d: %s\n", i, transaction.Sender)
	}
}

func account() {
	address := "0xA4710E3D79E1ED973AF58E0F269E9B21DD11BC64"

    // get balance of address
	balance := rpc.GetBalanceOfAddress(address)
	fmt.Println("Balance:", balance)

    // get nonce of address
	nonce := rpc.GetNonceOfAddress(address)
	fmt.Println("Nonce:", nonce)
}

func ReadData() {
	// account()
	// getBlock()
	// getVmData()
	// decoding()
	getVmDataActive()
}