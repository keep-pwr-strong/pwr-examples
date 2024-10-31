package sdk

import (
	"github.com/pwrlabs/pwrgo/wallet"
	"fmt"
)

func Wallet() {
	// Create a new wallet
	random := wallet.NewWallet()
	// Get the wallet address
	fmt.Println("Address:", random.GetAddress())

	// Get the wallet's private key
    privateKeys := random.GetPrivateKey()
    fmt.Println("PrivateKey:", privateKeys)

	// Get the wallet balance
	balance := random.GetBalance()
	fmt.Println("Balance:", balance)

	// Get the wallet's current nonce
	nonce := random.GetNonce()
	fmt.Println("Nonce:", nonce)
	
	// Create a wallet from an existing private key
	// in this example we will store the private key as a string
	privateKey := "0x04828e90065864c111871769c601d7de2246570b39dd37c19ccac16c14b18f72"
	wallet := wallet.FromPrivateKey(privateKey)
	fmt.Println("Address:", wallet.GetAddress())
}
