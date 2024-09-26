from pwrpy.pwrsdk import PWRPY
from pwrpy.pwrwallet import PWRWallet
from pwrpy.TransactionBuilder import TransactionBuilder
from transaction import Transactions
from dotenv import load_dotenv
import json
import time
import os
load_dotenv()

# Setting up your wallet in the SDK
private_key = os.getenv("PRIVATE_KEY")
wallet = PWRWallet(private_key)
# Setting up the rpc api
pwr = PWRPY()

def sync():
    starting_block = 876040 # Adjust starting block as needed
    vm_id = 123

    # Starting an infinite loop to continuously fetch and process transactions
    while True:
        # Fetching the latest block number from the blockchain via the RPC API
        latest_block = pwr.get_latest_block_number()
        # Setting the effective block range to fetch, with a limit of 1000 blocks per iteration
        effective_latest_block = min(latest_block, starting_block + 1000)

        # Checking if there are new blocks to process
        if effective_latest_block >= starting_block:
            # Fetching the latest block number from the blockchain
            txns = pwr.get_vm_data_txns(starting_block, effective_latest_block, vm_id)
            # Looping through the transactions fetched from the blockchain
            for txn in txns:
                sender = txn.sender
                data_hex = txn.data
                # Assuming data_hex is a hex string starting with '0x'
                data_bytes = bytes.fromhex(data_hex[2:])
                # Converting the hex data to a buffer and then to a UTF-8 string
                data_str = data_bytes.decode('utf-8')
                # Parsing the JSON string into a Python dictionary
                obj = json.loads(data_str)
                # Checking if the transaction contains a "message" field with the specified value
                if 'message' in obj and obj['message'].lower() == "please send me pwr":
                    # Building a transfer transaction to send PWR tokens
                    transfer_txn = TransactionBuilder.get_transfer_pwr_transaction(
                        sender, 100, wallet.get_nonce(), pwr.get_chainId()
                    )
                    # Adding the transfer transaction to the Transactions list for later execution
                    Transactions.add(transfer_txn)
                    # Printing a message to the console showing the sender and the message content
                    print(f"\nMessage from {sender}: {obj['message']}")
            # Updating the starting block number for the next loop iteration
            starting_block = effective_latest_block + 1
        time.sleep(1) # Wait 1 second before the next loop