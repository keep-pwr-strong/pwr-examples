use pwr_rs::RPC;
use hex;

async fn get_vm_data_active() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();
    let start_block = 843500;
    let end_block = 843750;
    let vm_id = 123;

    // fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
    let transactions = rpc.vm_data_transactions(start_block, end_block, vm_id).await.unwrap();

    for txs in transactions {
        let sender = txs.sender;
        let data = txs.data;

        // Convert the bytes data to a UTF-8 string
        let string_data = String::from_utf8(data.clone()).expect("Invalid UTF-8");

        if string_data.starts_with("Hi") {
            let word = &string_data[3..];
            println!("{}: {}", sender, word);
        }
        else if string_data.starts_with("Hello") {
            let word = &string_data[6..];
            println!("{}: {}", sender, word);
        }
    }
}

fn decoding() {
    let hex_data = "0x48656C6C6F20576F726C6421";

    // Remove the '0x' prefix and decode the hexadecimal data to bytes data
    let decoded_data = hex::decode(&hex_data[2..]).expect("Decoding failed");
    // Convert the decoded data to a UTF-8 string
    let string_data = String::from_utf8(decoded_data).expect("Invalid UTF-8");

    println!("Outputs: {}", string_data); // Outputs: Hello World!
}

async fn get_vm_data() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();

    let start_block = 843500;
    let end_block = 843750;
    let vm_id = 123;

    // fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
    let transactions = rpc.vm_data_transactions(start_block, end_block, vm_id).await.unwrap();
    // prints the trasnactions data
    for txs in transactions {
        println!("Data: {:?}", txs.data);
    }
}

async fn get_block() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();
    // the block number we want fetch
    let block_number = 20000;
    // get the block by number
    let block = rpc.block_by_number(block_number).await.unwrap();

    // prints the sender address from every transaction in the block
    for (index, txs) in block.transactions.iter().enumerate() {
        println!("Sender {}: {}", index, txs.sender);
    }
}

async fn account() {
    // Setting up the rpc api
    let rpc = RPC::new("https://pwrrpc.pwrlabs.io/").await.unwrap();
    let address = "0x3b3b69093879e7b6f28366fa3c32762590ff547e";
    // get balance of address
    let balance = rpc.balance_of_address(address).await.unwrap();
    println!("Balance: {balance}");
    // get nonce of address
    let nonce = rpc.nonce_of_address(address).await.unwrap();
    println!("Nonce: {nonce}");
}

#[tokio::main]
pub async fn main() {
    account().await;
    get_block().await;
    get_vm_data().await;
    decoding();
    get_vm_data_active().await;
}