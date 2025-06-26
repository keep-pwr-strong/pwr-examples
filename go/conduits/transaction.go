package conduits

import (
    "sync"
)

type Transactions struct {
    transactionsAwaitingApproval [][]byte
    mu                           sync.Mutex
}

var PendingTransactions = &Transactions{}

// Add adds a transaction in []byte format to the list of awaiting approvals.
func (t *Transactions) Add(txn []byte) {
    t.mu.Lock()
    defer t.mu.Unlock()
    t.transactionsAwaitingApproval = append(t.transactionsAwaitingApproval, txn)
}

// Remove removes a transaction in []byte format from the list of awaiting approvals.
func (t *Transactions) Remove(txn []byte) {
    t.mu.Lock()
    defer t.mu.Unlock()
    newList := [][]byte{}
    for _, tx := range t.transactionsAwaitingApproval {
        if string(tx) != string(txn) {
            newList = append(newList, tx)
        }
    }
    t.transactionsAwaitingApproval = newList
}

// GetPendingTransactions retrieves a copy of the transactions awaiting approval.
func (t *Transactions) GetPendingTransactions() [][]byte {
    t.mu.Lock()
    defer t.mu.Unlock()
    return append([][]byte{}, t.transactionsAwaitingApproval...)
}