from pwrpy.pwrsdk import PWRPY
from pwrpy.pwrwallet import Wallet
from pwrpy.models.Transaction import VidaDataTransaction
from pwrpy.TransactionBuilder import TransactionBuilder
from transaction import Transactions
from dotenv import load_dotenv
import json
import time
import os
load_dotenv()

# Setting up your wallet in the SDK
seed_phrase = os.getenv("SEED_PHRASE")
wallet = Wallet.new(seed_phrase)
# Setting up the rpc api
rpc = PWRPY("https://pwrrpc.pwrlabs.io/")

def handler_messages(txn: VidaDataTransaction):
    try:
        sender = txn.sender
        data_hex = txn.data
        # Assuming data_hex is a hex string starting with '0x'
        data_bytes = bytes.fromhex(data_hex)
        # Converting the hex data to a buffer and then to a UTF-8 string
        data_str = data_bytes.decode('utf-8')
        # Parsing the JSON string into a Python dictionary
        obj = json.loads(data_str)
        # Checking if the transaction contains a "message" field with the specified value
        if 'message' in obj and obj['message'].lower() == "please send me pwr":
            fee_per_byte = rpc.get_fee_per_byte()
            # Building a transfer transaction to send PWR tokens
            transfer_txn = TransactionBuilder.get_transfer_pwr_transaction(
                sender, 1000000000, wallet.get_nonce(), rpc.get_chainId(), wallet.get_address(), fee_per_byte
            )
            # Adding the transfer transaction to the Transactions list for later execution
            Transactions.add(transfer_txn)
            # Printing a message to the console showing the sender and the message content
            print(f"\nMessage from 0x{sender}: {obj['message']}")
    except Exception as e:
        print('Error in sync:', e)
        time.sleep(1)

def sync():
    starting_block = rpc.get_latest_block_number()
    vida_id = 123
    rpc.subscribe_to_vida_transactions(vida_id, starting_block, handler=handler_messages)