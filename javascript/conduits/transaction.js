export class Transactions {
    static transactionsAwaitingApproval = [];

    static add(txn) {
        this.transactionsAwaitingApproval.push(txn);
    }

    static remove(txn) {
        this.transactionsAwaitingApproval = this.transactionsAwaitingApproval.filter(
            tx => JSON.stringify(tx) !== JSON.stringify(txn)
        );
    }

    static getPendingTransactions() {
        return [...this.transactionsAwaitingApproval];
    }
}
