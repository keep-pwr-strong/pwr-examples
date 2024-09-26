import { PWRWallet } from "@pwrjs/core";
import dotenv from 'dotenv';
dotenv.config();

// Setting up your wallet in the SDK
const privateKey = process.env.PRIVATE_KEY;
const wallet = new PWRWallet(privateKey);

async function claim() {
    // Add a unique VM ID
    const vmId = 102030;

    // Claim the VM ID
    const txHash = await wallet.claimVmId(vmId);

    // Error handling
    if (txHash.success) {
        console.log("Transaction Hash:", txHash.transactionHash);
    } else {
        console.log("Error:", txHash.message);
    }
}
claim()