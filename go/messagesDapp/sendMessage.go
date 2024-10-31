package messagesDapp

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/wallet"
)

// var PrivateKey = "0x9D4428C6E0638331B4866B70C831F8BA51C11B031F4B55EED4087BBB8EF0151F"

func SendMessage() {
	// Setting up your wallet in the SDK
	godotenv.Load()
	privateKey := os.Getenv("PRIVATE_KEY")
	wallet := wallet.FromPrivateKey(privateKey)

	vmId := 123
	data, _ := json.Marshal(map[string]string{"message": "Hello World!"})

	tx := wallet.SendVMData(vmId, data)

	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Printf("Failed to send transaction: %s\n", tx.Error)
	}
}
