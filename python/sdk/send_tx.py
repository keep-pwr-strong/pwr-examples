from pwrpy.pwrwallet import Wallet
import time
from datetime import datetime, timedelta
from dotenv import load_dotenv
import os
load_dotenv()

# Setting up your wallet in the SDK
seed_phrase = os.getenv("SEED_PHRASE")
wallet = Wallet.new(seed_phrase)

def remove_guardian():
    # Remove your wallet guardian
    res = wallet.remove_guardian()

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

def set_guardian():
    # Guardian address that will verify your transactions
    guardian = "0x34bfe9c609ca72d5a4661889033a221fa07ef61a"

    # Guardian validity period - 30 minutes
    current_time = datetime.now()
    future_time = current_time + timedelta(minutes=30) # 30 minutes from now
    expiry_date = int(time.mktime(future_time.timetuple()))

    # Set your wallet guardian
    res = wallet.set_guardian(expiry_date, guardian)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

def move_stake():
    from_validator = "FROM_VALIDATOR_ADDRESS"
    to_validator = "TO_VALIDATOR_ADDRESS"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000000000

    # Move stake token from validator to another
    res = wallet.move_stake(amount, from_validator, to_validator)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

def withdraw():
    # Validator address you delegated
    validator_address = "VALIDATOR_ADDRESS"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000000000

    # Withdraw the delegated pwr tokens
    res = wallet.withdraw(amount, validator_address)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

def delegate():
    # Validator address
    validator_address = "VALIDATOR_ADDRESS"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000000000

    # Delegate the validator
    res = wallet.delegate(validator_address, amount)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

def send_payable_data():
    # VM id used to send the transaction to
    vida_id = 123
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 10
    # Buffer data to be included in the transaction
    data = "LFG!".encode()

    # Send the data at vmID 919 and pay 1e3
    res = wallet.send_payable_vida_data(vida_id, data, amount)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)

def send_data():
    # VM id used to send the transaction to
    vida_id = 123
    # Buffer data to be included in the transaction
    data = "Hello World!".encode()

    # Send the data at vmID 123 to the chain
    res = wallet.send_vida_data(vida_id, data)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)


def transfer():
    # Tokens recipient address
    recipient_address = "0x3B3B69093879E7B6F28366FA3C32762590FF547E"
    # Tokens amount - 1 PWR = 1e9 = 1000000000
    amount = 1000
    # Transfer pwr tokens from the wallet
    res = wallet.transfer_pwr(recipient_address, amount)

    # Error handling
    if res.success:
        print("Transaction Hash:", res.hash.hex())
    else:
        print("Error:", res.message)


if __name__ == "__main__":
    transfer()
    send_data()
    send_payable_data()
    # delegate()
    # withdraw()
    # move_stake()
    # set_guardian()
    # remove_guardian()