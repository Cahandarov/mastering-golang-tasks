package implementing

import "fmt"

type PaymentProcessor interface {
	processPayment(amount float64)
}
type CreditCard struct {
	amount float64
}
type PayPal struct {
	amount float64
}

func (c *CreditCard) processPayment(amount float64) {
	c.amount = amount
	fmt.Printf("Processing credit card payment of %.2f\n", c.amount)
}
func (p *PayPal) processPayment(amount float64) {
	p.amount = amount
	fmt.Printf("Processing PayPal payment of %.2f\n", p.amount)
}
func Task3() {
	fmt.Println("Task 3  ****************")

	creditCard := &CreditCard{}
	payPal := &PayPal{}

	creditCard.processPayment(100)
	payPal.processPayment(75.5)
}
