// lets create the enums for order status

package main

import "fmt"

type orderStatus int

const (
	Received orderStatus = iota
	Failed
	Confirmed
	Delivered
)

func changeOrderStatus(stat orderStatus){
	fmt.Println("Changed the order status to == > ",stat)
}

func main(){
	changeOrderStatus(Confirmed)
}