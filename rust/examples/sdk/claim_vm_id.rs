use pwr_rs::Wallet;
use dotenvy::dotenv;
use std::env;

async fn claim() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // Add a unique VM ID
    let vm_id = 102032;

    // Claim the VM ID
    let tx_hash = wallet.claim_vm_id(vm_id).await;

    println!("Transaction Hash: {tx_hash}")
}

#[tokio::main]
pub async fn main() {
    claim().await;
}