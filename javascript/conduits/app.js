import express from "express";
import { Transactions } from "./transaction.js"
import { sync } from "./conduit.js"

// Initialize the Express application by creating an app object
const app = express();

// Add sync to fetch messages and add it to the pending txs
sync();

// Define an HTTP GET route at '/pendingVmTransactions'
// When accessed, this route will return the list of pending transactions
app.get('/pendingVmTransactions', (req, res) => {
    // Set the response header to ensure the response is sent as JSON data
    res.header("Content-Type", "application/json");
    // Retrieve the list of pending transactions using the getPendingTransactions method
    const pendingTransactions = Transactions.getPendingTransactions();
     // Map through each transaction in the pendingTransactions array
    const array = pendingTransactions.map(txn => {
        // Convert each transaction (assumed to be a Buffer or Uint8Array) to a hexadecimal string
        const hexString = '0x' + Array.from(txn, byte => byte.toString(16).padStart(2, '0')).join(''); // Convert Buffer to hex string
        // Remove the transaction from the pending transactions list after processing
        Transactions.remove(txn);
        // Return the hexadecimal representation of the transaction
        return hexString;
    });
    // Send the resulting array of hex strings as a JSON response
    res.json(array);
})
// Set the port number for the server to listen on
const port = 8000;
// Start the Express server and listen for connections on the specified port
app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
})