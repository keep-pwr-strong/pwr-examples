package sdk

import (
	"github.com/pwrlabs/pwrgo/wallet"
	"fmt"
	"time"
)

var privateKey = "0x9D4428C6E0638331B4866B70C831F8BA51C11B031F4B55EED4087BBB8EF0151F"

func removeGuardian() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)

	// Remove your wallet guardian
	tx := wallet.RemoveGuardian()

    // Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func setGuardian() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)

	// Guardian address that will verify your transactions
	guardian := "0x34bfe9c609ca72d5a4661889033a221fa07ef61a"
	// Guardian validity period - 30 minutes
    futureDate := time.Now().Add(30 * time.Minute) // 30 minutes from now
    expiryDate := int(futureDate.Unix()) // Get the Unix timestamp in seconds

	// Set your wallet guardian
    tx := wallet.SetGuardian(guardian, expiryDate)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func moveStake() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)

	fromValidator := "FROM_VALIDATOR_ADDRESS"
    toValidator := "TO_VALIDATOR_ADDRESS"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e9)

	// Move stake token from validator to another
	tx := wallet.MoveStake(amount, fromValidator, toValidator)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func withdraw() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)

	// Validator address
	validator := "VALIDATOR_ADDRESS"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e9)

	// Withdraw the delegated pwr tokens
	tx := wallet.Withdraw(validator, amount)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func delegate() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)

	// Validator address
	validator := "VALIDATOR_ADDRESS"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e9)

	// Delegate the validator
	tx := wallet.Delegate(validator, amount)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func sendPayableData() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)
	// VM id used to send the transaction to
	vmId := 919
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := 10
	// Buffer data to be included in the transaction
	data := []byte("Hello world")

	// Send the data at vmID 919 and pay 1e3
	tx := wallet.SendPayableVMData(vmId, amount, data)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func sendData() {
	// Setting up your wallet in the SDK
	wallet := wallet.FromPrivateKey(privateKey)
	// VM id used to send the transaction to
	vmId := 123
	// Buffer data to be included in the transaction
	data := []byte("Hello world")

	// Send the data at vmID 123 to the chain
	tx := wallet.SendVMData(vmId, data)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func transfer() {
	wallet := wallet.FromPrivateKey(privateKey)

	// Tokens recipient address
	recipientAddress := "0x3B3b69093879e7B6F28366Fa3c32762590Ff547e"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e3)
	// Transfer pwr tokens from the wallet
	tx := wallet.TransferPWR(recipientAddress, amount)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.TxHash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func SendTx() {
	// transfer()
	// sendData()
	// sendPayableData()
	// delegate()
	// withdraw()
	// moveStake()
	// setGuardian()
	removeGuardian()
}