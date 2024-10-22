package implementing

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type CreditCard struct {
	amount float64
}
type PayPal struct {
	amount float64
}

func (c *CreditCard) ProcessPayment(amount float64) {
	c.amount = amount
	fmt.Printf("Processing credit card payment of %.2f\n", c.amount)
}
func (p *PayPal) ProcessPayment(amount float64) {
	p.amount = amount
	fmt.Printf("Processing PayPal payment of %.2f\n", p.amount)
}

func Task3() {
	fmt.Println("Task 3  ****************")
	var creditCard PaymentProcessor
	var payPal PaymentProcessor
	creditCard = &CreditCard{}
	payPal = &PayPal{}
	creditCard.ProcessPayment(100)
	payPal.ProcessPayment(75.5)
}
