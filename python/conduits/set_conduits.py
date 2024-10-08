from pwrpy.pwrwallet import PWRWallet
from dotenv import load_dotenv
import os
load_dotenv()

# Setting up your wallet in the SDK
private_key = os.getenv("PRIVATE_KEY")
wallet = PWRWallet(private_key)

def conduits():
    conduits = [
        bytes.fromhex("conduit_node_address"),
    ]
    vm_id = 9990

    # Sending the VM data transaction
    res = wallet.set_conduits(vm_id, conduits)
    print(res.data)

conduits()