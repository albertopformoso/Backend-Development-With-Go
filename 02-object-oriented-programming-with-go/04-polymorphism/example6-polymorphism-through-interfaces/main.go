package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {}

func (p Paypal) Pay() {
	fmt.Println("Payed with Paypal")
}

type Cash struct {}

func (c Cash) Pay() {
	fmt.Println("Payed with cash")
}

type CreditCard struct {}

func (c CreditCard) Pay() {
	fmt.Println("Payed with credit card")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Type a payment method number")
	fmt.Println("\t1: Paypal \n\t2: Cash \n\t3: Credit Card")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Type a valid method")
	}

	if method > 3 {
		panic("Type a valid method")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}