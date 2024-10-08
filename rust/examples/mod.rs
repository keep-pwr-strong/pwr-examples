pub mod sdk;
pub mod messages_dapp;
pub mod conduits;

fn main() {
    // sdk
    // sdk::wallet::main();
    // sdk::send_tx::main();
    // sdk::claim_vm_id::main();
    // sdk::read_data::main();

    // messages dapp
    // messages_dapp::send_message::main();
    // messages_dapp::sync_messages::main();
    // let _ = messages_dapp::dapp::main();

    // conduits
    // conduits::main();
    // conduits::app::main();
    conduits::set_conduits::main();
}
