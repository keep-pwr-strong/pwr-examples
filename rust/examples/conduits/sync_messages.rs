use pwr_rs::{RPC, Wallet, transaction::{NewTransactionData, VidaDataTransaction}};
use serde_json::{Value, json};
use crate::conduits::transaction::Transactions;
use dotenvy::dotenv;
use std::env;
use std::sync::Arc;

const FEE_PER_BYTE: u64 = 1000;
const CHAIN_ID: u8 = 0;

fn handler_messages(txn: VidaDataTransaction) {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // Spawn a new async task to handle the async operations
    tokio::spawn(async move {
        let sender = txn.sender;
        let data = txn.data;
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
                    amount: 1000000000, receiver: sender[2..].to_string()
                }.serialize_for_broadcast(wallet.get_nonce().await, CHAIN_ID, FEE_PER_BYTE, &wallet);
                // Checking if the transaction was successfully serialized, then adding it to the transactions
                if let Ok(txn_hex) = transfer_tx.map_err(|e| e.to_string()) {
                    // Convert each transaction (assumed to be a Buffer or Uint8Array) to a hexadecimal string
                    let hex_string = format!("0x{}", txn_hex.iter().map(|byte| format!("{:02x}", byte)).collect::<String>());
                    Transactions::add(json!(hex_string));
                }
                // Printing the sender and message to the console
                println!("\nMessage from 0x{}: {}", sender, message_str);
            }
        }
    });
}

pub async fn sync() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();
    let rpc = Arc::new(rpc);

    let starting_block: u64 = rpc.get_latest_block_number().await.unwrap();
    let vida_id: u64 = 123;

    rpc.subscribe_to_vida_transactions(vida_id, starting_block, handler_messages);
}