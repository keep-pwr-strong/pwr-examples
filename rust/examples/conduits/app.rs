use crate::conduits::{transaction::Transactions, sync_messages::sync};
use warp::Filter;
use tokio;

#[tokio::main]
pub async fn main() {
    // Add sync to fetch messages and add it to the pending txs
    tokio::spawn(async move {
        sync().await;
    });

    // Define an HTTP GET route at '/pending-vida-transactions'
    // When accessed, this route will return the list of pending transactions
    let user_route = warp::path("pending-vida-transactions")
        .map(move || {
            // Retrieve the list of pending transactions using the getPendingTransactions method
            let tx = Transactions::get_pending_transactions();
            // Map through each transaction in the pendingTransactions array
            for txn in tx.clone() {
                // Return the hexadecimal representation of the transaction
                Transactions::remove(&txn);
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