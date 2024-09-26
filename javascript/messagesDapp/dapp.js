import { PWRWallet } from "@pwrjs/core";
import { sync } from "./syncMessages.js";
import readline from "readline";
import dotenv from 'dotenv';
dotenv.config();

// Setting up your wallet in the SDK
const privateKey = process.env.PRIVATE_KEY2;
const wallet = new PWRWallet(privateKey);

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

async function main() {
    const vmId = 1234;
    await sync();

    const messageLoop = () => {
        rl.question("", async (message) => {
            const object = { message };

            const response = await wallet.sendVMDataTxn(vmId, Buffer.from(JSON.stringify(object), 'utf8'));
            !response.success && console.log('FAILED!');
            messageLoop(); // Recursively ask for the next message
        });
    };

    messageLoop();
}
main();
