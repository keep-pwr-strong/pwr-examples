use std::sync::{Mutex, Arc};
use serde_json::Value;

pub struct Transactions {
    transactions_awaiting_approval: Arc<Mutex<Vec<Value>>>,
}

impl Transactions {
    pub fn new() -> Self {
        Transactions {
            transactions_awaiting_approval: Arc::new(Mutex::new(Vec::new())),
        }
    }

    pub fn add(&self, txn: Value) {
        let mut txns = self.transactions_awaiting_approval.lock().unwrap();
        txns.push(txn);
    }

    pub fn remove(&self, txn: &Value) {
        let mut txns = self.transactions_awaiting_approval.lock().unwrap();
        *txns = txns.iter()
            .filter(|&tx| tx != txn) // Comparing JSON values directly
            .cloned()
            .collect();
    }

    pub fn get_pending_transactions(&self) -> Vec<Value> {
        let txns = self.transactions_awaiting_approval.lock().unwrap();
        txns.clone()
    }
}

