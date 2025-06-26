package messagesDapp

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
    "github.com/pwrlabs/pwrgo/wallet"
    "github.com/joho/godotenv"
)

func main() {
    // Setting up your wallet in the SDK
    godotenv.Load()
    seedPhrase := os.Getenv("SEED_PHRASE")
    wallet, err := wallet.New(seedPhrase)
    if err != nil {
        fmt.Println("Error getting the wallet:", err)
        return
    }
    vidaId := 1234 // Replace with your VIDA's ID

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
        feePerByte := wallet.GetRpc().GetFeePerByte()

        // Send the VIDA data
        tx := wallet.SendVidaData(vidaId, jsonData, feePerByte)
        if tx.Success {
            fmt.Printf("Transaction Hash: %s\n", tx.Hash)
        } else {
            fmt.Println("Error:", tx.Error)
        }
    }
}