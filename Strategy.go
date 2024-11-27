package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Pagando $%.2f con tarjeta de crédito.\n", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Pagando $%.2f con PayPal.\n", amount)
}

type BankTransferPayment struct{}

func (b *BankTransferPayment) Pay(amount float64) {
	fmt.Printf("Pagando $%.2f con transferencia bancaria.\n", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	if p.strategy == nil {
		fmt.Println("Por favor, selecciona un método de pago.")
		return
	}
	p.strategy.Pay(amount)
}

func main() {
	payment := &PaymentContext{}

	payment.SetStrategy(&CreditCardPayment{})
	payment.Pay(100.0)

	payment.SetStrategy(&PayPalPayment{})
	payment.Pay(200.0)

	payment.SetStrategy(&BankTransferPayment{})
	payment.Pay(300.0)
}
