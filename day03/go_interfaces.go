// implementation of payment gateway using interfaces

package main

import "fmt"

type Paymenter interface{
	Pay(amount float32)
}

type payment struct{
	gateway Paymenter
}

func (obj payment) makePayment(amount float32){
	obj.gateway.Pay(amount);
}

// different gateways
type razorpay struct{}

func (o razorpay) Pay(amount float32){
	fmt.Println("Transaction done with razorpay of ",amount," $")
}

type paypal struct{}

func (p paypal) Pay(amount float32){
	fmt.Println("Transaction done with paypal of ",amount,"$")
}

func main() {
	fmt.Println("Here we will be implementating the interfaces in go.....")
	// raxorpay instance
	// razorpayinstance:=razorpay{}
	paypalinstance:=paypal{}
	newpayment:=payment{
		gateway: paypalinstance,
	}
	newpayment.makePayment(100);

}