from pwrpy.pwrwallet import PWRWallet
from sync_messages import sync
from dotenv import load_dotenv
import json
import threading
import os
load_dotenv()

# Setting up your wallet in the SDK
private_key = os.getenv("PRIVATE_KEY")
wallet = PWRWallet(private_key)
vm_id = 1234

def main():
    threading.Thread(target=sync, daemon=True).start()

    while True:
        message = input("")
        obj = {"message": message}
        data = json.dumps(obj).encode('utf-8')
        
        response = wallet.send_vm_data_transaction(vm_id, data)
        if response.success==False:
            print('FAILED!')
main()
