import { PWRWallet } from "@pwrjs/core";
const randomWallet = new PWRWallet();

// Get the wallet address
const address = randomWallet.getAddress();
console.log(`Address: ${address}`);

// Get the wallet's private key
// const privateKey = randomWallet.getPrivateKey();
// console.log(`PrivateKey: ${privateKey}`);

// Get the wallet balance
randomWallet.getBalance()
    .then(balance => console.log(`Balance: ${balance}`));

// Get the wallet's current nonce
randomWallet.getNonce()
    .then(nonce => console.log(`Nonce: ${nonce}`));

// Create a wallet from an existing private key (String || ByteArray || Int)
// in this example we will store the private key as a string
const privateKey = "0x04828e90065864c111871769c601d7de2246570b39dd37c19ccac16c14b18f72";
const wallet = new PWRWallet(privateKey);
console.log(wallet.getAddress());
