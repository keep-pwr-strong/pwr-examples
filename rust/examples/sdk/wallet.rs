use pwr_rs::Wallet;

#[tokio::main]
pub async fn main() {
    let random_wallet = Wallet::new_random(12);

    // Get the wallet address
    let address = random_wallet.get_address();
    println!("Address: {address}");

    // Get the wallet's private key
    let seed_phrase = random_wallet.get_seed_phrase();
    println!("Seed Phrase: {seed_phrase}");

    // Get the wallet's public key
    let public_key = random_wallet.get_public_key();
    println!("Public Key: {:?}", hex::encode(public_key));

    // Get the wallet's private key
    let private_key = random_wallet.get_private_key();
    println!("Private Key: {:?}", hex::encode(private_key));

    // Get the wallet balance
    let balance = random_wallet.get_balance().await;
    println!("Balance: {balance}");

    // Get the wallet's current nonce
    let nonce = random_wallet.get_nonce().await;
    println!("Nonce: {nonce}");

    // Create a wallet from an existing seed phrase (String)
    // in this example we will store the seed phrase
    let seed_phrase = "badge drive deputy afraid siren always green about certain stuff play surround";
    let wallet = Wallet::new(seed_phrase);

    // Get the wallet's address
    let address = wallet.get_address();
    println!("Address: {address}");
}
