package conduits

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"log"
	"time"
	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/rpc"
	"github.com/pwrlabs/pwrgo/wallet"
	"github.com/pwrlabs/pwrgo/encode"
)

func Sync() {
	// Setting up your wallet in the SDK
    godotenv.Load()
    privateKey := os.Getenv("PRIVATE_KEY")
    wallet := wallet.FromPrivateKey(privateKey)

	startingBlock := 876040 // Adjust starting block as needed
	vmId := 123

	// Defining an asynchronous loop function that fetches and processes new transactions
	go func() {
		for {
			// Fetching the latest block number from the blockchain via the RPC API
			latestBlock := rpc.GetLatestBlockNumber()
			// Defining the effective block range for the next batch of transactions, limiting to 1000 blocks at a time
			effectiveLatestBlock := latestBlock
			if latestBlock > startingBlock+1000 {
				effectiveLatestBlock = startingBlock + 1000
			}

			// Checking if there are new blocks to process
			if effectiveLatestBlock > startingBlock {
				// Fetching VM data transactions between the starting block and the effective latest block for a given VM ID
				txns := rpc.GetVmDataTransactions(startingBlock, effectiveLatestBlock, vmId)
				// Looping through the transactions fetched from the blockchain
				for _, txn := range txns {
					sender := txn.Sender
					dataHex := txn.Data
					nonce := wallet.GetNonce()

					// Converting the hex data to a buffer and then to a UTF-8 string
					dataBytes, _ := hex.DecodeString(dataHex[2:])
					var obj map[string]interface{}
					if err := json.Unmarshal(dataBytes, &obj); err != nil {
						log.Println("Error parsing JSON:", err)
						continue
					}

					// Iterating over each key in the object to check for specific conditions
					if message, ok := obj["message"].(string); ok && message == "please send me pwr" {
						var buffer []byte
						// Building a transfer transaction to send PWR tokens
						buffer, err := encode.TransferTxBytes(100, sender, nonce)
						if err != nil {
							log.Println("Error encoding transaction:", err)
							continue
						}
						// Adding the transaction to the Transactions struct
						PendingTransactions.Add(buffer)
						// Logging the message and the sender to the console
						fmt.Printf("\nMessage from %s: %s\n", sender, message)
					}
				}
				// Updating the starting block number for the next loop iteration
				startingBlock = effectiveLatestBlock + 1
			}
			time.Sleep(1 * time.Second) // Wait 1 second before the next loop
		}
	}()
}