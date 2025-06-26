use pwr_rs::Wallet;
use dotenvy::dotenv;
use std::env;
use serde_json::json;

pub async fn send_message() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    let vida_id = 123;
    let obj = json!({ "message": "please send me pwr" });
    let data = serde_json::to_vec(&obj).unwrap(); // Serialize to JSON bytes
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Sending the VIDA data transaction
    let res = wallet.send_vida_data(vida_id, data, fee_per_byte).await;
    if res.success {
        println!("Transaction hash: {:?}", res.data.unwrap());
    } else {
        println!("Transaction failed: {:?}", res.error);
    }
}

#[tokio::main]
pub async fn main() {
    send_message().await;
}