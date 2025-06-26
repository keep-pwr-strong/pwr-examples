from pwrpy.pwrsdk import PWRPY

# Setting up the rpc api
pwr = PWRPY("https://pwrrpc.pwrlabs.io/")

def get_vida_data_active():
    start_block = 40635
    end_block = 40726
    vida_id = 123
    # fetch the transactions sent from `start_block` to `end_block` in `vida_id`
    transactions = pwr.get_vida_data_transactions(start_block, end_block, vida_id)

    for txs in transactions:
        sender = txs.sender
        data = txs.data

        # Decode the hexadecimal data to bytes data
        decoded_data = bytes.fromhex(data)
        # Convert the bytes data to a UTF-8 string
        string_data = decoded_data.decode('utf-8')

        if string_data.startswith("Hi"):
            word = string_data[3:]
            print(f'{sender}: {word}')
        elif string_data.startswith("Hello"):
            word = string_data[6:]
            print(f'{sender}: {word}')
get_vida_data_active()

def decoding():
    hex_data = "0x48656C6C6F20576F726C6421"

    # Remove the '0x' prefix and decode the hexadecimal data to bytes data
    decoded_data = bytes.fromhex(hex_data[2:])
    # Convert the decoded data to a UTF-8 string
    string_data = decoded_data.decode('utf-8')

    print(f'Outputs: {string_data}') # Outputs: Hello World!
decoding()

def get_vida_data():
    start_block = 40635
    end_block = 40726
    vida_id = 123

    transactions = pwr.get_vida_data_transactions(start_block, end_block, vida_id)
    for txs in transactions:
        print("Data:", txs.data)
get_vida_data()

def get_block():
    # the block number we want fetch
    block_number = 100
    # get the block by number
    block = pwr.get_block_by_number(block_number)
    
    # prints the sender address from every transaction in the block
    for index, txs in enumerate(block.transactions):
        transaction = pwr.get_transaction_by_hash(txs.transaction_hash)
        print(f"Sender {index}: {transaction.sender}")
get_block()

def account():
    address = "0x3b3b69093879e7b6f28366fa3c32762590ff547e"

    # get balance of address
    balance = pwr.get_balance_of_address(address)
    print(f"Balance: {balance}")
    # get nonce of address
    nonce = pwr.get_nonce_of_address(address)
    print(f"Nonce: {nonce}")
account()
