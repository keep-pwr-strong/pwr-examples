package messagesDapp

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
	"encoding/hex"
	"github.com/pwrlabs/pwrgo/rpc"
)

func Sync() {
	startingBlock := 880920
	vmId := 1234

	loop := func() {
		for {
			latestBlock := rpc.GetLatestBlockNumber()

			effectiveLatestBlock := latestBlock
			if latestBlock > startingBlock+1000 {
				effectiveLatestBlock = startingBlock + 1000
			}

			if effectiveLatestBlock > startingBlock {
				// fetch the transactions in `vmId = 1234`
				transactions := rpc.GetVmDataTransactions(startingBlock, effectiveLatestBlock, vmId)

				for _, txn := range transactions {
					sender := txn.Sender
					dataHex := txn.Data
					// Remove the '0x' prefix and decode the hexadecimal data to bytes data
					dataBytes, _ := hex.DecodeString(dataHex[2:])
					// convert the bytes data to UTF-8 string as json
					var obj map[string]interface{}
					if err := json.Unmarshal(dataBytes, &obj); err != nil {
						log.Println("Error parsing JSON:", err)
						continue
					}

					if message, ok := obj["message"]; ok {
						fmt.Printf("\nMessage from %s: %s\n", sender, message)
					} else {
						// Handle other data fields if needed
					}
				}

				startingBlock = effectiveLatestBlock + 1
			}

			time.Sleep(1 * time.Second) // Wait 1 second before the next loop
		}
	}

	go loop()
}
