package sdk

import (
	"github.com/pwrlabs/pwrgo/wallet"
	"fmt"
)

// var privateKey = "0x9D4428C6E0638331B4866B70C831F8BA51C11B031F4B55EED4087BBB8EF0151F"

func ClaimVmId() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)

	// Add a unique VM ID
	vmId := 102030

	// Claim the VM ID
	tx := wallet.ClaimVMId(vmId)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}
