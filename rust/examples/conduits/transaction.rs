use std::sync::{Mutex, Arc};
use serde_json::Value;
use lazy_static::lazy_static;

lazy_static! {
    static ref TRANSACTIONS: Arc<Mutex<Vec<Value>>> = Arc::new(Mutex::new(Vec::new()));
}

pub struct Transactions;

impl Transactions {
    pub fn new() -> Self {
        Transactions
    }

    pub fn add(txn: Value) {
        let mut txns = TRANSACTIONS.lock().unwrap();
        txns.push(txn);
    }

    pub fn remove(txn: &Value) {
        let mut txns = TRANSACTIONS.lock().unwrap();
        *txns = txns.iter()
            .filter(|&tx| tx != txn)
            .cloned()
            .collect();
    }

    pub fn get_pending_transactions() -> Vec<Value> {
        let txns = TRANSACTIONS.lock().unwrap();
        txns.clone()
    }
}