from pwrpy.pwrwallet import PWRWallet
from dotenv import load_dotenv
import os
load_dotenv()

# Setting up your wallet in the SDK
private_key = os.getenv("PRIVATE_KEY")
wallet = PWRWallet(private_key)

def claim():
    # Add a unique VM ID
    vm_id = 102030

    # Claim the VM ID
    tx_hash = wallet.claim_vm_id(vm_id)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

if __name__ == "__main__":
    claim()