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
    seedPhrase := os.Getenv("SEED_PHRASE")
    wallet, _ := wallet.New(seedPhrase)

    vidaId := 123
    data, _ := json.Marshal(map[string]string{"message": "please send me pwr"})
    feePerByte := wallet.GetRpc().GetFeePerByte()

    // Sending the VIDA data transaction
    tx := wallet.SendVidaData(vidaId, data, feePerByte)

    if tx.Success {
        fmt.Printf("Transaction Hash: %s\n", tx.Hash)
    } else {
        fmt.Printf("Failed to send transaction: %s\n", tx.Error)
    }
}