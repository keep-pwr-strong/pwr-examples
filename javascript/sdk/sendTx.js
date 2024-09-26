import { PWRWallet } from "@pwrjs/core";
import dotenv from 'dotenv';
dotenv.config();

// Setting up your wallet in the SDK
const privateKey = process.env.PRIVATE_KEY;
const wallet = new PWRWallet(privateKey);

async function sendMessage() {
    const obj = { message: "please send me pwr" };
    const data = Buffer.from(JSON.stringify(obj), 'utf8');
    const vmId = 123;

    const res = await wallet.sendVMDataTxn(vmId, data);
    console.log(res.transactionHash);
}
sendMessage()

async function removeGuardian() {
    // Remove your wallet guardian
    const txHash = await wallet.removeGuardian();

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
removeGuardian()

async function setGuardian() {
    // Guardian address that will verify your transactions
    const guardian = "0x34bfe9c609ca72d5a4661889033a221fa07ef61a";

    // Guardian validity period - 30 minutes
    const futureDate = new Date();
    futureDate.setDate(futureDate.getMinutes() + 30); // 30 minutes from now
    const expiryDate = Math.floor(futureDate.getTime() / 1000);
    
    // Set your wallet guardian
    const txHash = await wallet.setGuardian(guardian, expiryDate);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
setGuardian()

async function moveStake() {
    const fromValidator = "FROM_VALIDATOR_ADDRESS";
    const toValidator = "TO_VALIDATOR_ADDRESS";
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    const amount = 1e9;
    
    // Move stake token from validator to another
    const txHash = await wallet.moveStake(amount, fromValidator, toValidator);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
moveStake()

async function withdraw() {
    // Validator address you delegated
    const validator = "VALIDATOR_ADDRESS";
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    const amount = 1e9;
    
    // Withdraw the delegated pwr tokens
    const txHash = await wallet.withdraw(validator, amount);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
withdraw()

async function delegate() {
    // Validator address
    const validator = "VALIDATOR_ADDRESS";
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    const amount = 1e9;
    
    // Delegate the validator
    const txHash = await wallet.delegate(validator, amount);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
delegate()

async function sendPayableData() {
    // VM id used to send the transaction to
    const vmId = 919;
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    const amount = 10;
    // Buffer data to be included in the transaction
    const data = Buffer.from('Hello World!');
    
    // Send the data at vmID 919 and pay 1e3
    const txHash = await wallet.sendPayableVmDataTransaction(vmId, amount, data);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
sendPayableData()

async function sendData() {
    // VM id used to send the transaction to
    const vmId = 123;
    // Buffer data to be included in the transaction
    const data = Buffer.from('TO THE MOOOOOOON!');
    
    // Send the data at vmID 123 to the chain
    const txHash = await wallet.sendVMDataTxn(vmId, data);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
sendData()

async function transfer() {
    // Tokens recipient address
    const recipientAddress = "RECIPIENT_ADDRESS";
    // Tokens amount - 1 PWR = 1e9 = 1000000000
    const amount = 1e3;
    // Transfer pwr tokens from the wallet
    const txHash = await wallet.transferPWR(recipientAddress, amount);
    
    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
transfer();
