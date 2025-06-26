pub mod sdk;
pub mod messages_dapp;
pub mod conduits;

fn main() {
    // sdk
    // sdk::wallet::main();
    // sdk::send_tx::main();
    // sdk::claim_vida_id::main();
    // sdk::read_data::main();

    // messages dapp
    // let _ = messages_dapp::dapp::main();

    // conduits
    conduits::app::main();
}
