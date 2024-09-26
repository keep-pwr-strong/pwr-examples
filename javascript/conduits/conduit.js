import { PWRJS, PWRWallet, TransactionBuilder } from "@pwrjs/core";
import { Transactions } from "./transaction.js"
import dotenv from 'dotenv';
dotenv.config();

// Setting up your wallet in the SDK
const privateKey = process.env.PRIVATE_KEY;
const wallet = new PWRWallet(privateKey);
// Setting up the rpc api
const rpc = new PWRJS("https://pwrrpc.pwrlabs.io/");

export async function sync() {
    let startingBlock = 876040; // Adjust starting block as needed
    const vmId = 123;

    // Defining an asynchronous loop function that fetches and processes new transactions
    const loop = async () => {
        // Fetching the latest block number from the blockchain via the RPC API
        const latestBlock = await rpc.getLatestBlockNumber();
        // Defining the effective block range for the next batch of transactions, limiting to 1000 blocks at a time
        let effectiveLatestBlock = latestBlock > startingBlock + 1000 ? startingBlock + 1000 : latestBlock;

        // Checking if there are new blocks to process
        if (effectiveLatestBlock > startingBlock) {
            // Fetching VM data transactions between the starting block and the effective latest block for a given VM ID
            const txns = await rpc.getVMDataTransactions(startingBlock, effectiveLatestBlock, vmId);
            // Looping through the transactions fetched from the blockchain
            for (let txn of txns) {
                const sender = txn.sender;
                const dataHex = txn.data;
                let nonce = await wallet.getNonce();
                // Converting the hex data to a buffer and then to a UTF-8 string
                const data = Buffer.from(dataHex.substring(2), 'hex');
                const object = JSON.parse(data.toString('utf8'));

                // Iterating over each key in the object to check for specific conditions
                Object.keys(object).forEach(async (key) => {
                    if (key.toLowerCase() === "message" && object[key].toLowerCase() === "please send me pwr") {
                        // Building a transfer transaction to send PWR tokens
                        const transferTxn = TransactionBuilder.getTransferPwrTransaction(
                            rpc.getChainId(), nonce, 100, sender
                        );
                        // Adding the transaction to the Transactions class
                        Transactions.add(transferTxn)
                        // Logging the message and the sender to the console
                        console.log(`\nMessage from ${sender}: ${object[key]}`);
                    }
                });
            }
            // Updating the starting block number for the next loop iteration
            startingBlock = effectiveLatestBlock + 1;
        }
        setTimeout(loop, 1000); // Wait 1 second before the next loop
    }
    loop();
}