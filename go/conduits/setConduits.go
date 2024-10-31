package conduits

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/wallet"
)

func SetConduits() {
	// Setting up your wallet in the SDK
    godotenv.Load()
    privateKey := os.Getenv("PRIVATE_KEY")
    wallet := wallet.FromPrivateKey(privateKey)

	vmIds := 9999
	conduits := []string{"0x7EbFBd2BABA5F68F720C059d62eFc4aaCFA66513"}

	tx := wallet.SetConduits(vmIds, conduits)

	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}
