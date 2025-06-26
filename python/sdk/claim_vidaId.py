from pwrpy.pwrwallet import Wallet
from dotenv import load_dotenv
import os
load_dotenv()

# Setting up your wallet in the SDK
seed_phrase = os.getenv("SEED_PHRASE")
wallet = Wallet.new(seed_phrase)

def claim():
    # Add a unique VM ID
    vida_id = 102030

    # Claim the VM ID
    res = wallet.claim_vida_id(vida_id)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

if __name__ == "__main__":
    claim()