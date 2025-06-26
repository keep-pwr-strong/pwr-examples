use pwr_rs::Wallet;
use dotenvy::dotenv;
use std::env;

async fn claim() {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);

    // Add a unique VM ID
    let vida_id = 102032;
    let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

    // Claim the VM ID
    let res = wallet.claim_vida_id(vida_id, fee_per_byte).await;

    if res.success {
        println!("Transaction Hash: {}", res.data.unwrap());
    } else {
        println!("Error: {}", res.error);
    }
}

#[tokio::main]
pub async fn main() {
    claim().await;
}