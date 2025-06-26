package sdk

import (
	"fmt"

	"github.com/pwrlabs/pwrgo/wallet"
)

func Wallet() {
	// Create a new wallet
	random, err := wallet.NewRandom(12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Get the wallet address
	fmt.Println("Address:", random.GetAddress())

	// Get the wallet's seed phrase
	seedPhrase := random.GetSeedPhrase()
	fmt.Println("SeedPhrase:", seedPhrase)

	// Get the wallet's public key
	publicKey := random.GetPublicKey()
	fmt.Println("PublicKey:", publicKey)

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
	seedPhraseString := "badge drive deputy afraid siren always green about certain stuff play surround"
	wallet, err := wallet.New(seedPhraseString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Address:", wallet.GetAddress())
}
