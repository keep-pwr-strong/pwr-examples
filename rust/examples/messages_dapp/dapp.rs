use pwr_rs::Wallet;
use tokio::io::{self, AsyncBufReadExt, BufReader};
use tokio::spawn;
use crate::messages_dapp::sync_messages::sync;
use dotenvy::dotenv;
use std::env;

#[tokio::main]
pub async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();
    let vm_id: u64 = 1234;
    let stdin = io::stdin();
    let reader = BufReader::new(stdin);
    let mut lines = reader.lines();
    
    spawn(sync());
    
    while let Some(message) = lines.next_line().await? {
        let obj = serde_json::json!({ "message": message });
        let data = serde_json::to_vec(&obj)?;

        // Send the VM data and get the transaction hash
        wallet.send_vm_data(vm_id, data).await;
    }
    Ok(())
}
