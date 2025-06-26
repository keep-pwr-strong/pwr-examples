from pwrpy.pwrwallet import Wallet
from sync_messages import sync
from dotenv import load_dotenv
import json
import threading
import os
load_dotenv()

# Setting up your wallet in the SDK
seed_phrase = os.getenv("SEED_PHRASE")
wallet = Wallet.new(seed_phrase)
vida_id = 1234

def main():
    threading.Thread(target=sync, daemon=True).start()

    while True:
        message = input("")
        obj = {"message": message}
        data = json.dumps(obj).encode('utf-8')
        
        # Send the VIDA data
        response = wallet.send_vida_data(vida_id, data)
        if response.success==False:
            print('FAILED!')
main()