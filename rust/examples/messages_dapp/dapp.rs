use pwr_rs::Wallet;
use tokio::io::{self, AsyncBufReadExt, BufReader};
use tokio::spawn;
use crate::messages_dapp::sync_messages::sync;
use dotenvy::dotenv;
use std::env;

#[tokio::main]
pub async fn homain() -> Result<(), Box<dyn std::error::Error>> {
    dotenv().ok();
    // Setting up your wallet in the SDK
    let seed_phrase = env::var("SEED_PHRASE").unwrap();
    let wallet = Wallet::new(&seed_phrase);
    let vida_id: u64 = 1234;
    let stdin = io::stdin();
    let reader = BufReader::new(stdin);
    let mut lines = reader.lines();
    
    spawn(sync());
    
    while let Some(message) = lines.next_line().await? {
        let obj = serde_json::json!({ "message": message });
        let data = serde_json::to_vec(&obj)?;
        let fee_per_byte = (wallet.get_rpc().await).get_fee_per_byte().await.unwrap();

        // Send the VIDA data
        wallet.send_vida_data(vida_id, data, fee_per_byte).await;
    }
    Ok(())
}