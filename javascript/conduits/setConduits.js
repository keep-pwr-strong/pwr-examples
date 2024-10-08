import { PWRWallet } from "@pwrjs/core";
import dotenv from 'dotenv';
dotenv.config();

// Setting up your wallet in the SDK
const privateKey = process.env.PRIVATE_KEY;
const wallet = new PWRWallet(privateKey);

async function conduits() {
    const conduits = [
        Buffer.from("conduit_node_address", "hex"),
    ];
    const vmId = 9999;

    const res = await wallet.setConduits(vmId, conduits);
    console.log(res.transactionHash);
}
conduits();
