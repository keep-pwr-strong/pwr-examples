use pwr_rs::Wallet;

#[tokio::main]
pub async fn main() {
    let random_wallet = Wallet::random();

    // Get the wallet address
    let address = random_wallet.get_address();
    println!("Address: {address}");

    // Get the wallet's private key
    // let private_key = random_wallet.private_key();
    // println!("PrivateKey: {private_key}");

    // Get the wallet balance
    let balance = random_wallet.get_balance().await;
    println!("Balance: {balance}");

    // Get the wallet's current nonce
    let nonce = random_wallet.get_nonce().await;
    println!("Balance: {nonce}");
}
