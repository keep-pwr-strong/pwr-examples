use pwr_rs::{RPC, Wallet, transaction::NewTransactionData};
use std::time::Duration;
use tokio::time::sleep;
use serde_json::{Value, json};
use crate::conduits::transaction::Transactions;
use dotenvy::dotenv;
use std::env;

pub async fn sync(transactions: &Transactions) {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();
    
    let mut starting_block: u64 = 876040; // Adjust starting block as needed
    let vm_id: u64 = 123;

    // Starting an infinite loop to continuously fetch and process transactions
    loop {
        // Fetch the latest block number without explicit error handling
        let latest_block = rpc.lates_block_number().await.unwrap();
        // Defining the effective block range for the next batch of transactions, limiting to 1000 blocks at a time
        let effective_latest_block = if latest_block > starting_block + 1000 {
            starting_block + 1000
        } else {
            latest_block
        };
        // Checking if there are new blocks to process
        if effective_latest_block >= starting_block {
            // Fetching VM data transactions between the starting block and the effective latest block for a given VM ID
            let txns = rpc.vm_data_transactions(starting_block, effective_latest_block, vm_id).await.unwrap();
            // Iterating through each transaction
            for txn in txns {
                let sender = txn.sender;
                let data = txn.data; // Assuming txn.data is Vec<u8>
                // Convert data bytes to UTF-8 string without explicit error handling
                let data_str = String::from_utf8(data).unwrap();
                // Parse JSON data without explicit error handling
                let object: Value = serde_json::from_str(&data_str).unwrap();
                // Converting the parsed JSON object into a map of key-value pairs
                let obj_map = object.as_object().unwrap();
                // Iterating through the key-value pairs in the parsed JSON object
                for (key, value) in obj_map {
                    // Converting the value to a string (assuming it's a string value)
                    let message_str = value.as_str().unwrap();
                    if key.to_lowercase() == "message" && message_str == "please send me pwr" {
                        // Constructing a transfer transaction using the NewTransactionData::Transfer variant
                        let transfer_tx = NewTransactionData::Transfer { 
                            amount: 100, recipient: sender[2..].to_string()
                        }.serialize_for_broadcast(wallet.get_nonce().await, rpc.chain_id, &wallet);
                        // Checking if the transaction was successfully serialized, then adding it to the transactions
                        if let Ok(txn_hex) = transfer_tx.map_err(|e| e.to_string()) {
                            let hex_string = format!("0x{}", txn_hex.iter().map(|byte| format!("{:02x}", byte)).collect::<String>());
                            transactions.add(json!(hex_string));
                        }
                        // Printing the sender and message to the console
                        println!("\nMessage from {}: {}", sender, message_str);
                    }
                }
            }
            // Updating the starting block number for the next loop iteration
            starting_block = effective_latest_block + 1;
        }
        sleep(Duration::from_secs(1)).await; // Wait 1 second before the next loop
    }
}