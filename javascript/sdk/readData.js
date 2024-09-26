import { PWRJS } from "@pwrjs/core";

// Setting up the rpc api
const rpc = new PWRJS("https://pwrrpc.pwrlabs.io/");

async function getVmDataActive() {
    const startBlock = 843500;
    const endBlock = 843750;
    const vmId = 123;

    // fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
    const transactions = await rpc.getVMDataTransactions(startBlock, endBlock, vmId);

    for (let txs of transactions) {
        const sender = txs.sender;
        const data = txs.data;

        // Remove the '0x' prefix and decode the hexadecimal data to bytes data
        const decodedData = Buffer.from(data.slice(2), 'hex');
        // Convert the bytes data to a UTF-8 string
        const stringData = decodedData.toString('utf8');
        
        if (stringData.startsWith("Hi")) {
            const word = stringData.substring(3);
            console.log(`${sender}: ${word}`);
        }
        else if (stringData.startsWith("Hello")) {
            const word = stringData.substring(6);
            console.log(`${sender}: ${word}`)
        }
    }
}
getVmDataActive()

async function decoding() {
    const hexData = "0x48656C6C6F20576F726C6421";

    // Remove the '0x' prefix and decode the hexadecimal data to bytes data
    const decodedData = Buffer.from(hexData.slice(2), 'hex');
    // Convert the decoded data to a UTF-8 string
    const stringData = decodedData.toString('utf8');

    console.log(`Outputs: ${stringData}`); // Outputs: Hello World!
}
// decoding()

async function getVmData() {
    const startBlock = 843500;
    const endBlock = 843750;
    const vmId = 123;

    // fetch the transactions sent from `startBlock` to `endBlock` in `vmId`
    const transactions = await rpc.getVMDataTransactions(startBlock, endBlock, vmId);

    // prints the trasnactions data
    for (let txs of transactions) {
        console.log("Data:", txs.data);
    }
}
// getVmData()

async function getBlock() {
    // the block number we want fetch
    const blockNumber = 20000;
    // get the block by number
    const block = await rpc.getBlockByNumber(blockNumber);

    for (let i in block.transactions) {
        console.log(`Sender ${i}: ${block.transactions[i].sender}`);
    }
}
// getBlock()

async function account() {
    const address = "0x3b3b69093879e7b6f28366fa3c32762590ff547e";

    // get balance of address
    const balance = await rpc.getBalanceOfAddress(address);
    console.log(`Balance: ${balance}`);
    // get nonce of address
    const nonce = await rpc.getNonceOfAddress(address);
    console.log(`Nonce: ${nonce}`);
}
// account()

async function vmData() {
    // const transactions = await rpc.getVMDataTransactions(836599, 836600, 69);
    // for (let i=0; i<transactions.length; i++) {
    //     console.log(transactions[i].sender);
    // }
    // console.log(transactions.length)
}
// vmData()
