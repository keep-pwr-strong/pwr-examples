from pwrpy.pwrwallet import Wallet
from dotenv import load_dotenv
import json
import os
load_dotenv()

# Setting up your wallet in the SDK
seed_phrase = os.getenv("SEED_PHRASE")
wallet = Wallet.new(seed_phrase)

def send_message():
    vida_id = 123
    obj = {"message": "please send me pwr"}
    data = json.dumps(obj).encode('utf-8') # Serialize to JSON bytes
    fee_per_byte = wallet.get_rpc().get_fee_per_byte()

    # Sending the VIDA data transaction
    res = wallet.send_vida_data(vida_id, data, fee_per_byte)
    if res.success:
        print(f"Transaction hash: 0x{res.hash.hex()}")
    else:
        print(f"Transaction failed: {res.error}")
send_message()