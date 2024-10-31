package conduits

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/wallet"
)

func SendMessage() {
	// Setting up your wallet in the SDK
    godotenv.Load()
    privateKey := os.Getenv("PRIVATE_KEY")
    wallet := wallet.FromPrivateKey(privateKey)

	vmId := 123
	data, _ := json.Marshal(map[string]string{"message": "please send me pwr"})

	// Sending the VM data transaction
	tx := wallet.SendVMData(vmId, data)

	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Printf("Failed to send transaction: %s\n", tx.Error)
	}
}