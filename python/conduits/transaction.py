class Transactions:
    transactions_awaiting_approval = []

    @classmethod
    def add(cls, txn):
        cls.transactions_awaiting_approval.append(txn)

    @classmethod
    def remove(cls, txn):
        cls.transactions_awaiting_approval = [
            tx for tx in cls.transactions_awaiting_approval 
            if tx != txn
        ]

    @classmethod
    def get_pending_transactions(cls):
        return cls.transactions_awaiting_approval.copy()
