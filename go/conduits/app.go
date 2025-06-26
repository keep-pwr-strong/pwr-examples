package conduits

import (
    "fmt"
    "log"
    "encoding/hex"
    "encoding/json"
    "net/http"
)

func main() {
    // Add sync to fetch messages and add it to the pending txs
    go Sync()

    // Define an HTTP GET route at '/pending-vida-transaction'
    // When accessed, this route will return the list of pending transactions
    http.HandleFunc("/pending-vida-transactions", func(w http.ResponseWriter, r *http.Request) {
        // Set the response header to ensure the response is sent as JSON data
        w.Header().Set("Content-Type", "application/json")
        // Retrieve the list of pending transactions using the getPendingTransactions method
        pending := PendingTransactions.GetPendingTransactions()

        var hexStrings []string
        // Map through each transaction in the pendingTransactions array
        for _, txn := range pending {
            // Convert each transaction (assumed to be a Buffer or Uint8Array) to a hexadecimal string
            hexString := "0x" + hex.EncodeToString(txn)
            hexStrings = append(hexStrings, hexString)
            // Remove the transaction from the pending transactions list after processing
            PendingTransactions.Remove(txn)
        }

        // Send the resulting array of hex strings as a JSON response
        if err := json.NewEncoder(w).Encode(hexStrings); err != nil {
            http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
            return
        }
    })

    // Set the port number for the server to listen on
    port := ":8000"
    fmt.Printf("Server running on http://localhost%s\n", port)
    // Start the HTTP server and listen for connections on the specified port
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}