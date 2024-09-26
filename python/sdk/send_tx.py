from pwrpy.pwrwallet import PWRWallet
import time
from datetime import datetime, timedelta
from dotenv import load_dotenv
import os
load_dotenv()

# Setting up your wallet in the SDK
private_key = os.getenv("PRIVATE_KEY")
wallet = PWRWallet(private_key)

def remove_guardian():
    # Remove your wallet guardian
    tx_hash = wallet.remove_guardian()

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

def set_guardian():
    # Guardian address that will verify your transactions
    guardian = "0x34bfe9c609ca72d5a4661889033a221fa07ef61a"

    # Guardian validity period - 30 minutes
    current_time = datetime.now()
    future_time = current_time + timedelta(minutes=30) # 30 minutes from now
    expiry_date = int(time.mktime(future_time.timetuple()))

    # Set your wallet guardian
    tx_hash = wallet.set_guardian(guardian, expiry_date)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

def move_stake():
    from_validator = "FROM_VALIDATOR_ADDRESS"
    to_validator = "TO_VALIDATOR_ADDRESS"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000000000

    # Move stake token from validator to another
    tx_hash = wallet.move_stake(amount, from_validator, to_validator)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

def withdraw():
    # Validator address you delegated
    validator = "VALIDATOR_ADDRESS"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000000000

    # Withdraw the delegated pwr tokens
    tx_hash = wallet.withdraw(validator, amount)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

def delegate():
    # Validator address
    validator = "VALIDATOR_ADDRESS"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000000000

    # Delegate the validator
    tx_hash = wallet.delegate(validator, amount)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

def send_payable_data():
    # VM id used to send the transaction to
    vm_id = 919
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 10
    # Buffer data to be included in the transaction
    data = "Hello World!".encode()

    # Send the data at vmID 919 and pay 1e3
    tx_hash = wallet.send_payable_vm_data_transaction(vm_id, amount, data)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)

def send_data():
    # VM id used to send the transaction to
    vm_id = 123
    # Buffer data to be included in the transaction
    data = "Hello World!".encode()

    # Send the data at vmID 123 to the chain
    tx_hash = wallet.send_vm_data_transaction(vm_id, data)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)


def transfer():
    # Tokens recipient address
    recipient_address = "0x3B3B69093879E7B6F28366FA3C32762590FF547E"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000
    # Transfer pwr tokens from the wallet
    tx_hash = wallet.transfer_pwr(recipient_address, amount)

    # Error handling
    if tx_hash.success:
        print("Transaction Hash:", tx_hash.data)
    else:
        print("Error:", tx_hash.message)


if __name__ == "__main__":
    transfer()
    send_data()
    send_payable_data()
    delegate()
    withdraw()
    move_stake()
    set_guardian()
    remove_guardian()