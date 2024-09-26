from pwrpy.pwrsdk import PWRPY
import json
import time

# Setting up the rpc api
pwr = PWRPY()

def sync():
    starting_block = 880920 
    vm_id = 1234

    while True:
        try:
            latest_block = pwr.get_latest_block_number()
            effective_latest_block = min(latest_block, starting_block + 1000)

            if effective_latest_block >= starting_block:
                # Fetch the transactions in `vmId = 1234`
                txns = pwr.get_vm_data_txns(starting_block, effective_latest_block, vm_id)
                for txn in txns:
                    sender = txn.sender
                    data_hex = txn.data
                    # Remove the '0x' prefix and decode the hexadecimal data to bytes data
                    data_bytes = bytes.fromhex(data_hex[2:])
                    # Convert the bytes data to UTF-8 string as json
                    obj = json.loads(data_bytes.decode('utf-8'))
                    if 'message' in obj:
                        print(f"\nMessage from {sender}: {obj['message']}")

                starting_block = effective_latest_block + 1
            time.sleep(1) # Wait 1 second before the next loop
        except Exception as e:
            print('Error in sync:', e)
            time.sleep(1)
# sync()