use pwr_rs::Wallet;
use dotenvy::dotenv;
use std::env;
use serde_json::json;

async fn send_message() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    let obj = json!({ "message": "lfg" });
    let data = serde_json::to_vec(&obj).unwrap(); // Serialize to JSON bytes
    let vm_id = 1234;
    let res = wallet.send_vm_data(vm_id, data).await;
    println!("{}", res);
}

#[tokio::main]
pub async fn main() {
    send_message().await;
}