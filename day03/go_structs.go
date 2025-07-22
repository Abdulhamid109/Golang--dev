// here we will be creating an invoice template using go

package main

import (
	"fmt"
	"time"
)


type invoice struct{
	id int
	productName string
	quantity int
	price float64
	status string
	time time.Time
}

// we can attach the struct to a function 
// when you want to modify the struct then use the pointers in that case
func (obj *invoice) changeStatus(status string){
	obj.status = status
}

// we can create the constructor using functions
func Newinvoice(id int,productName string,quantity int,price float64,status string) * invoice{
	myinvoice:= invoice{
		id: id,
		productName: productName,
		quantity: quantity,
		price: price,
		status: status,
	}
	return &myinvoice;
}
func main(){
	
	myinvoice := Newinvoice(1,"milk",2,100,"paid-online")
	myinvoice.time = time.Now()
	myinvoice.changeStatus("paid-cash")

	fmt.Println(myinvoice)
}