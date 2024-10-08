use pwr_rs::Wallet;
use dotenvy::dotenv;
use std::env;

async fn conduits() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let private_key = env::var("PRIVATE_KEY").unwrap();
    let wallet = Wallet::from_hex(&private_key).unwrap();

    let conduits: Vec<String> = vec![
        "0x7882A8b75d2128708F58F0945e684A2679929Eef".to_string(),
    ];
    let vm_id: u64 = 9990;

    let res = wallet.set_conduits(vm_id, conduits).await;
    println!("{}", res);
}

#[tokio::main]
pub async fn main() {
    conduits().await;
}