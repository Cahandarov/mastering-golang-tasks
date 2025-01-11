package main

import (
	"fmt"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}
type CardPayment struct {
	CardNumber string
}

type WalletPayment struct {
	WalletID string
}

type BankTransferPayment struct {
	AccountNumber string
}

func (p CardPayment) ProcessPayment(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("cannot pay with negative amount")
	} else {
		return nil
	}
}
func (p WalletPayment) ProcessPayment(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("cannot pay with negative amount")
	} else {
		return nil
	}
}
func (p BankTransferPayment) ProcessPayment(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("cannot pay with negative amount")
	} else {
		return nil
	}
}

func main() {
	var processor PaymentProcessor

	fmt.Println("Select payment method: 1-Card, 2-Wallet, 3-Bank Transfer")
	var choice int
	fmt.Scan(&choice)

	fmt.Println("Enter payment amount:")
	var amount float64
	fmt.Scan(&amount)

	switch choice {
	case 1:
		processor = CardPayment{CardNumber: "1234-5678-9012-3456"}
	case 2:
		processor = WalletPayment{WalletID: "wallet123"}
	case 3:
		processor = BankTransferPayment{AccountNumber: "AZ1234567890"}
	default:
		fmt.Println("Invalid choice")
		return
	}

	err := processor.ProcessPayment(amount)
	if err != nil {
		fmt.Printf("Payment failed: %s\n", err)
	} else {
		fmt.Println("Payment successful!")
	}
}
