use pwr_rs::RPC;
use pwr_rs::transaction::types::VidaDataTransaction;
use std::sync::Arc;

fn handler_messages(txn: VidaDataTransaction) {
    // Get the address of the transaction sender
    let sender = txn.sender;
    // Get the data sent in the transaction (In Hex Format)
    let data = txn.data;
    // Convert data string to bytes
    let data_str = String::from_utf8(data).unwrap();
    let object: serde_json::Value = serde_json::from_str(&data_str).unwrap();
    let obj_map = object.as_object().unwrap();

    // Check the action and execute the necessary code
    for (key, value) in obj_map {
        if key.to_lowercase() == "message" {
            let message_str = value.as_str().unwrap();
            println!("\nMessage from {}: {}", sender, message_str);
        } else {
            // Handle other data fields if needed
        }
    }
}

pub async fn sync() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();
    let rpc = Arc::new(rpc);

    let starting_block = rpc.get_latest_block_number().await.unwrap();
    let vida_id: u64 = 1234;

    rpc.subscribe_to_vida_transactions(vida_id, starting_block, handler_messages);
}