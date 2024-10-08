use pwr_rs::Wallet;
use chrono::{Duration, Utc};
use dotenvy::dotenv;
use std::env;

async fn remove_guardian() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // Remove your wallet guardian
    let tx_hash = wallet.remove_guardian().await;

    println!("Transaction Hash: {tx_hash}");
}

async fn set_guardian() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();
    println!("ADDRESS: {}", wallet.get_address());

    // Guardian address that will verify your transactions
    let guardian = "0x34bfe9c609ca72d5a4661889033a221fa07ef61a".to_string();

    // Guardian validity period - 30 minutes
    let current_time = Utc::now();
    let future_time = current_time + Duration::minutes(10); // 30 minutes from now
    let expiry_date = future_time.timestamp() as u64;

    // Set your wallet guardian
    let tx_hash = wallet.set_guardian(guardian, expiry_date).await;

    println!("Transaction Hash: {tx_hash}");
}

async fn move_stake() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    let from_validator = "FROM_VALIDATOR_ADDRESS".to_string();
    let to_validator = "TO_VALIDATOR_ADDRESS".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000000000;

    // Move stake token from validator to another
    let tx_hash = wallet.move_stake(amount, from_validator, to_validator).await;

    println!("Transaction Hash: {tx_hash}");
}

async fn withdraw() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // Validator address you delegated
    let validator = "0x3b3b69093879e7b6f28366fa3c32762590ff547e".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000000000;

    // Delegate the validator
    let tx_hash = wallet.withdraw(validator, amount).await;

    println!("Transaction Hash: {tx_hash}");
}

async fn delegate() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // Validator address
    let validator = "VALIDATOR_ADDRESS".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000000000;

    // Delegate the validator
    let tx_hash = wallet.delegate(validator, amount).await;

    println!("Transaction Hash: {tx_hash}");
}

async fn send_payable_data() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // VM id used to send the transaction to
    let vm_id = 919;
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 10;
    // Buffer data to be included in the transaction
    let data = vec!["Hello World!"];
    let data_as_bytes: Vec<u8> = data.into_iter().flat_map(|s| s.as_bytes().to_vec()).collect();

    // Send the data at vmID 919 and pay 1e3
    let tx_hash = wallet.send_payable_vm_data(vm_id, amount, data_as_bytes).await;

    println!("Transaction Hash: {tx_hash}");
}

async fn send_data() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // VM id used to send the transaction to
    let vm_id = 123;
    // Buffer data to be included in the transaction
    let data = vec!["Hello World!"];
    let data_as_bytes: Vec<u8> = data.into_iter().flat_map(|s| s.as_bytes().to_vec()).collect();

    // Send the data at vmID 123 to the chain
    let tx_hash = wallet.send_vm_data(vm_id, data_as_bytes).await;

    println!("Transaction Hash: {tx_hash}");
}

async fn transfer() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    // Tokens recipient address
    let recipient_address = "0x3B3B69093879E7B6F28366FA3C32762590FF547E".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000;

    // Transfer pwr tokens from the wallet
    let tx_hash = wallet.transfer_pwr(recipient_address, amount).await;

    println!("Transaction Hash: {tx_hash}");
}

#[tokio::main]
pub async fn main() {
    transfer().await;
    send_data().await;
    send_payable_data().await;
    delegate().await;
    withdraw().await;
    move_stake().await;
    set_guardian().await;
    remove_guardian().await;
}


