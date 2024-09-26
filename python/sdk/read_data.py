from pwrpy.pwrsdk import PWRPY

# Setting up the rpc api
pwr = PWRPY()

def get_vm_data_active():
    start_block = 843500
    end_block = 843750
    vm_id = 123
    # fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
    transactions = pwr.get_vm_data_txns(start_block, end_block, vm_id)

    for txs in transactions:
        sender = txs.sender
        data = txs.data

        # Remove the '0x' prefix and decode the hexadecimal data to bytes data
        decoded_data = bytes.fromhex(data[2:])
        # Convert the bytes data to a UTF-8 string
        string_data = decoded_data.decode('utf-8')

        if string_data.startswith("Hi"):
            word = string_data[3:]
            print(f'{sender}: {word}')
        elif string_data.startswith("Hello"):
            word = string_data[6:]
            print(f'{sender}: {word}')
get_vm_data_active()

def decoding():
    hex_data = "0x48656C6C6F20576F726C6421"

    # Remove the '0x' prefix and decode the hexadecimal data to bytes data
    decoded_data = bytes.fromhex(hex_data[2:])
    # Convert the decoded data to a UTF-8 string
    string_data = decoded_data.decode('utf-8')

    print(f'Outputs: {string_data}') # Outputs: Hello World!
# decoding()

def get_vm_data():
    start_block = 843500
    end_block = 843750
    vm_id = 123

    transactions = pwr.get_vm_data_txns(start_block, end_block, vm_id)
    for txs in transactions:
        print("Data:", txs.data)
# get_vm_data()

def get_block():
    # the block number we want fetch
    block_number = 20000
    # get the block by number
    block = pwr.get_block_by_number(block_number)
    
    # prints the sender address from every transaction in the block
    for index, txs in enumerate(block.transactions):
        print(f"Sender {index}: {txs.sender}")
# get_block()

def account():
    address = "0x3b3b69093879e7b6f28366fa3c32762590ff547e"

    # get balance of address
    balance = pwr.get_balance_of_address(address)
    print(f"Balance: {balance}")
    # get nonce of address
    nonce = pwr.get_nonce_of_address(address)
    print(f"Nonce: {nonce}")
# account()
