package messagesDapp

import (
    "fmt"
    "log"
    "encoding/json"
    "encoding/hex"
    "github.com/pwrlabs/pwrgo/rpc"
)


func handlerMessages(transaction rpc.VidaDataTransaction) {
    sender := transaction.Sender
    dataHex := transaction.Data
    // Decode the hexadecimal data to bytes data
    dataBytes, _ := hex.DecodeString(dataHex)
    // convert the bytes data to UTF-8 string as json
    var obj map[string]interface{}
    if err := json.Unmarshal(dataBytes, &obj); err != nil {
        log.Println("Error parsing JSON:", err)
    }

    if message, ok := obj["message"]; ok {
        fmt.Printf("\nMessage from %s: %s\n", sender, message)
    } else {
        // Handle other data fields if needed
    }
}

func Sync() {
    rpc := rpc.SetRpcNodeUrl("https://pwrrpc.pwrlabs.io")
    startingBlock := rpc.GetLatestBlockNumber()
    vidaId := 1234 // Replace with your VIDA's ID

    _ = rpc.SubscribeToVidaTransactions(
        vidaId,
        startingBlock,
        handlerMessages,
    )
}