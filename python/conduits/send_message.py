from pwrpy.pwrwallet import PWRWallet
from dotenv import load_dotenv
import json
import os
load_dotenv()

# Setting up your wallet in the SDK
private_key = os.getenv("PRIVATE_KEY")
wallet = PWRWallet(private_key)

def send_message():
    obj = {"message": "please send me pwr"}
    data = json.dumps(obj).encode('utf-8')
    vm_id = 123

    # Sending the VM data transaction
    res = wallet.send_vm_data_transaction(vm_id, data)
    print(res.data)
send_message()