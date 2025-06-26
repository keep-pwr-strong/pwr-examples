from pwrpy.pwrwallet import Wallet

random_wallet = Wallet.new_random(12)

# Get the wallet address
address = random_wallet.get_address()
print("Address:", address)

# Get the wallet's private key
seed_phrase = random_wallet.get_seed_phrase()
print("PrivateKey:", seed_phrase)

# Get the wallet's public key
public_key = random_wallet.get_public_key()
print("Public Key:", public_key.hex())

# Get the wallet's private key
private_key = random_wallet.get_private_key()
print("Private Key:", private_key.hex())

# Get the wallet balance
balance = random_wallet.get_balance()
print("Balance:", balance)

# Get the wallet's current nonce
nonce = random_wallet.get_nonce()
print("Nonce:", nonce)

# Create a wallet from an existing seed phrase (String)
# in this example we will store the seed phrase
seed_phrase = "badge drive deputy afraid siren always green about certain stuff play surround"
wallet = Wallet.new(seed_phrase)
print("Address:", wallet.get_address())