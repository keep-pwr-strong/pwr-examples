use pwr_rs::RPC;
use std::time::Duration;
use tokio::time::sleep;
use serde_json::Value;

pub async fn sync() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();

    let mut starting_block: u64 = 880920;
    let vm_id: u64 = 1234;

    loop {
        let latest_block = rpc.lates_block_number().await.unwrap();
        let effective_latest_block = if latest_block > starting_block + 1000 {
            starting_block + 1000
        } else {
            latest_block
        };

        if effective_latest_block >= starting_block {
            // Fetch the transactions in `vmId = 1234`
            let txns = rpc.vm_data_transactions(starting_block, effective_latest_block, vm_id).await.unwrap();
            for txn in txns {
                let sender = txn.sender;
                let data = txn.data; // txn.data is Vec<u8>
                // Convert data bytes to UTF-8 string without explicit error handling
                let data_str = String::from_utf8(data).unwrap();
                // Parse JSON data without explicit error handling
                let object: Value = serde_json::from_str(&data_str).unwrap();
                let obj_map = object.as_object().unwrap();

                for (key, value) in obj_map {
                    if key.to_lowercase() == "message" {
                        let message_str = value.as_str().unwrap();
                        println!("\nMessage from {}: {}", sender, message_str);
                    } else {
                        // Handle other data fields if needed
                    }
                }
            }
            starting_block = effective_latest_block + 1;
        }
        sleep(Duration::from_secs(1)).await; // Wait 1 second before the next loop
    }
}


#[tokio::main]
pub async fn main() {
    sync().await;
}
