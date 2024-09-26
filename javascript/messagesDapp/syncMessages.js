import { PWRJS } from "@pwrjs/core";

// Setting up your wallet in the SDK
const rpc = new PWRJS("https://pwrrpc.pwrlabs.io/");

export async function sync() {
    let startingBlock = 880920;
    const vmId = 1234;

    const loop = async () => {
        try {
            const latestBlock = await rpc.getLatestBlockNumber();
            let effectiveLatestBlock = latestBlock > startingBlock + 1000 ? startingBlock + 1000 : latestBlock;

            if (effectiveLatestBlock > startingBlock) {
                // Fetch the transactions in `vmId = 1234`
                const txns = await rpc.getVMDataTransactions(startingBlock, effectiveLatestBlock, vmId);

                for (let txn of txns) {
                    const sender = txn.sender;
                    const dataHex = txn.data;
                    // Remove the '0x' prefix and decode the hexadecimal data to bytes data
                    const data = Buffer.from(dataHex.substring(2), 'hex');
                    // Convert the bytes data to UTF-8 string as json
                    const object = JSON.parse(data.toString('utf8'));

                    Object.keys(object).forEach(key => {
                        if (key.toLowerCase() === "message") {
                            console.log(`\nMessage from ${sender}: ${object[key]}`);
                        } else {
                            // Handle other data fields if needed
                        }
                    });
                }

                startingBlock = effectiveLatestBlock + 1;
            }
            setTimeout(loop, 1000); // Wait 1 second before the next loop
        } catch (e) {
            console.error(e);
        }
    };
    loop();
}
// sync();
