from pwrpy.pwrsdk import PWRPY
from pwrpy.models.Transaction import VidaDataTransaction
import json
import time

# Setting up the rpc api
rpc = PWRPY("https://pwrrpc.pwrlabs.io/")

vida_id = 1234 # Replace with your VIDA's ID
starting_block = rpc.get_latest_block_number()

def handler_messages(txn: VidaDataTransaction):
    try:
        sender = txn.sender
        data_hex = txn.data
        # Decode the hexadecimal data to bytes data
        data_bytes = bytes.fromhex(data_hex)
        # Convert the bytes data to UTF-8 string as json
        obj = json.loads(data_bytes.decode('utf-8'))
        if 'message' in obj:
            print(f"\nMessage from {sender}: {obj['message']}")
    except Exception as e:
        print('Error in sync:', e)
        time.sleep(1)

def sync():
    rpc.subscribe_to_vida_transactions(vida_id, starting_block, handler=handler_messages)