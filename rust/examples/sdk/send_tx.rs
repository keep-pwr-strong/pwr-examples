use pwr_rs::Wallet;
use chrono::{Duration, Utc};
use dotenvy::dotenv;
use std::env;

async fn remove_guardian() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Remove your wallet guardian
    let res = wallet.remove_guardian(fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn set_guardian() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);
    println!("ADDRESS: {}", wallet.get_address());

    // Guardian address that will verify your transactions
    let guardian = "0x34bfe9c609ca72d5a4661889033a221fa07ef61a".to_string();

    // Guardian validity period - 30 minutes
    let current_time = Utc::now();
    let future_time = current_time + Duration::minutes(10); // 30 minutes from now
    let expiry_date = future_time.timestamp() as u64;

    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Set your wallet guardian
    let res = wallet.set_guardian(expiry_date, guardian, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn move_stake() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    let from_validator = "FROM_VALIDATOR_ADDRESS".to_string();
    let to_validator = "TO_VALIDATOR_ADDRESS".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000000000;
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Move stake token from validator to another
    let res = wallet.move_stake(amount, from_validator, to_validator, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn withdraw() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // Validator address you delegated
    let validator = "0x3b3b69093879e7b6f28366fa3c32762590ff547e".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000000000;
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Delegate the validator
    let res = wallet.withdraw(amount, validator, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn delegate() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // Validator address
    let validator = "VALIDATOR_ADDRESS".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000000000;
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Delegate the validator
    let res = wallet.delegate(validator, amount, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn send_payable_data() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // VM id used to send the transaction to
    let vida_id = 919;
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 10;
    // Buffer data to be included in the transaction
    let data = vec!["Hello World!"];
    let data_as_bytes: Vec<u8> = data.into_iter().flat_map(|s| s.as_bytes().to_vec()).collect();
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Send the data at vmID 919 and pay 1e3
    let res = wallet.send_payable_vida_data(vida_id, data_as_bytes, amount, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn send_data() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // VM id used to send the transaction to
    let vida_id = 123;
    // Buffer data to be included in the transaction
    let data = vec!["Hello World!"];
    let data_as_bytes: Vec<u8> = data.into_iter().flat_map(|s| s.as_bytes().to_vec()).collect();
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Send the data at vmID 123 to the chain
    let res = wallet.send_vida_data(vida_id, data_as_bytes, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

async fn transfer() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // Tokens recipient address
    let recipient_address = "0x3B3B69093879E7B6F28366FA3C32762590FF547E".to_string();
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    let amount = 1000;
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Transfer pwr tokens from the wallet
    let res = wallet.transfer_pwr(recipient_address, amount, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
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


