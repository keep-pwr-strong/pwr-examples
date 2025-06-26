package sdk

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/wallet"
)

func ClaimVidaId() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")

	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Add a unique VM ID
	vidaId := 102030
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Claim the VM ID
	tx := wallet.ClaimVidaId(vidaId, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}
