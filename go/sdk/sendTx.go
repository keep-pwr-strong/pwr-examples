package sdk

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pwrlabs/pwrgo/wallet"
)

func removeGuardian() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Remove your wallet guardian
	tx := wallet.RemoveGuardian(feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func setGuardian() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Guardian address that will verify your transactions
	guardian := "0x34bfe9c609ca72d5a4661889033a221fa07ef61a"
	// Guardian validity period - 30 minutes
	futureDate := time.Now().Add(30 * time.Minute) // 30 minutes from now
	expiryDate := int(futureDate.Unix())           // Get the Unix timestamp in seconds

	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Set your wallet guardian
	tx := wallet.SetGuardian(expiryDate, guardian, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func moveStake() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fromValidator := "FROM_VALIDATOR_ADDRESS"
	toValidator := "TO_VALIDATOR_ADDRESS"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int64(1e9)
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Move stake token from validator to another
	tx := wallet.MoveStake(amount, fromValidator, toValidator, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func withdraw() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Validator address
	validator := "VALIDATOR_ADDRESS"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e9)
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Withdraw the delegated pwr tokens
	tx := wallet.Withdraw(amount, validator, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func delegate() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Validator address
	validator := "VALIDATOR_ADDRESS"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e9)
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Delegate the validator
	tx := wallet.Delegate(validator, amount, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func sendPayableData() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// VM id used to send the transaction to
	vidaId := 919
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int64(10)
	// Buffer data to be included in the transaction
	data := []byte("Hello world")
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Send the data at vmID 919 and pay 1e3
	tx := wallet.SendPayableVidaData(vidaId, data, amount, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func sendData() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// VM id used to send the transaction to
	vidaId := 123
	// Buffer data to be included in the transaction
	data := []byte("Hello world")
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Send the data at vmID 123 to the chain
	tx := wallet.SendVidaData(vidaId, data, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
	} else {
		fmt.Println("Error:", tx.Error)
	}
}

func transfer() {
	godotenv.Load()
	seedPhrase := os.Getenv("SEED_PHRASE")
	// Setting up your wallet in the SDK
	wallet, err := wallet.New(seedPhrase)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Tokens recipient address
	recipientAddress := "0x3B3b69093879e7B6F28366Fa3c32762590Ff547e"
	// Tokens amount - 1 PWR = 1e9 = 1000000000
	amount := int(1e3)
	feePerByte := wallet.GetRpc().GetFeePerByte()

	// Transfer pwr tokens from the wallet
	tx := wallet.TransferPWR(recipientAddress, amount, feePerByte)

	// Error handling
	if tx.Success {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash)
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
