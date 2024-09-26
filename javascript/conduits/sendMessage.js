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
sendMessage();