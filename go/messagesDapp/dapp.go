package messagesDapp

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/wallet"
)

func DApp() {
	// Setting up your wallet in the SDK
	godotenv.Load()
	privateKey := os.Getenv("PRIVATE_KEY")
	wallet := wallet.FromPrivateKey(privateKey)
	vmId := 1234

	go Sync()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your message: ")
		message, _ := reader.ReadString('\n')
		message = message[:len(message)-1]

		object := map[string]string{"message": message}
		jsonData, err := json.Marshal(object)
		if err != nil {
			fmt.Println("Failed to encode message:", err)
			continue
		}

		// Send the VM data
		tx := wallet.SendVMData(vmId, jsonData)
		if tx.Success {
			fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
		} else {
			fmt.Println("Error:", tx.Error)
		}
	}
}