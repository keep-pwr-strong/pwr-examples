use crate::conduits::{transaction::Transactions, conduit::sync};
use warp::Filter;
use tokio;
use std::sync::Arc;

#[tokio::main]
pub async fn main() {
    let txs = Arc::new(Transactions::new());

    // Add sync to fetch messages and add it to the pending txs
    tokio::spawn({
        let txs = txs.clone();
        async move {
            sync(&txs).await;
        }
    });

    // Define an HTTP GET route at '/pendingVmTransactions'
    // When accessed, this route will return the list of pending transactions
    let user_route = warp::path("pendingVmTransactions")
        .map(move || {
            // Retrieve the list of pending transactions using the getPendingTransactions method
            let tx = txs.get_pending_transactions();
            // Map through each transaction in the pendingTransactions array
            for txn in tx.clone() {
                // Return the hexadecimal representation of the transaction
                txs.remove(&txn);
            }
            // Send the resulting array of hex strings as a JSON response
            warp::reply::json(&tx)
        });
    
    // Set the port number for the server to listen on
    // Start the Warp server and listen for connections on the specified port
    warp::serve(user_route)
        .run(([127, 0, 0, 1], 8000)) // Bind to localhost:8000
        .await;
}