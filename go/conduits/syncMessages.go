package conduits

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "github.com/pwrlabs/pwrgo/rpc"
    "github.com/pwrlabs/pwrgo/wallet"
    "github.com/pwrlabs/pwrgo/encode"
)

func handlerMessages(transaction rpc.VidaDataTransaction) {
    godotenv.Load()
    seedPhrase := os.Getenv("SEED_PHRASE")
    wallet, err := wallet.New(seedPhrase)
    if err != nil {
        fmt.Println("Error getting the wallet:", err)
        return
    }

    sender := transaction.Sender
    dataHex := transaction.Data
    nonce := wallet.GetNonce()

    // Converting the hex data to a buffer and then to a UTF-8 string
    dataBytes, _ := hex.DecodeString(dataHex)
    var obj map[string]interface{}
    if err := json.Unmarshal(dataBytes, &obj); err != nil {
        fmt.Println("Error parsing JSON:", err)
    }

    // Iterating over each key in the object to check for specific conditions
    if message, ok := obj["message"].(string); ok && message == "please send me pwr" {
        var buffer []byte
        // Building a transfer transaction to send PWR tokens
        buffer, err := encode.TransferTxBytes(1000000000, sender, nonce, wallet.Address, wallet.GetRpc().GetFeePerByte())
        if err != nil {
            fmt.Println("Error encoding transaction:", err)
        }
        // Adding the transaction to the Transactions struct
        PendingTransactions.Add(buffer)
        // Logging the message and the sender to the console
        fmt.Printf("\nMessage from 0x%s: %s\n", sender, message)
    }
}

func Sync() {
    rpc := rpc.SetRpcNodeUrl("https://pwrrpc.pwrlabs.io")
    startingBlock := rpc.GetLatestBlockNumber()
    vidaId := 123

    rpc.SubscribeToVidaTransactions(vidaId, startingBlock, handlerMessages)
}