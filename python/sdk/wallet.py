from pwrpy.pwrwallet import PWRWallet

random_wallet = PWRWallet()

# Get the wallet address
address = random_wallet.get_address()
print("Address:", address)

# Get the wallet's private key
private_key = random_wallet.get_private_key()
print("PrivateKey:", private_key)

# Get the wallet balance
balance = random_wallet.get_balance()
print("Balance:", balance)

# Get the wallet's current nonce
nonce = random_wallet.get_nonce().data
print("Nonce:", nonce)

# Create a wallet from an existing private key (String || ByteArray || Int)
# in this example we will store the private key as a string
private_key = "0x04828e90065864c111871769c601d7de2246570b39dd37c19ccac16c14b18f72";
wallet = PWRWallet(private_key)
print("Address:", wallet.get_address())