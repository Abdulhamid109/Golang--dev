package main

import "fmt"

func main() {
	fmt.Print("Conditonal statement")
	// else if
	age:=18
	if age>=18{
		fmt.Println("person is adult")
	}else if age>=12{
		fmt.Println("teenager")
	}else{
		fmt.Print("feeling sleepy")
	}

	
	// different feature i.e we can declare the variable inside the if condition
	if weigth:=20; weigth<20{
		fmt.Println("UnderWeigth",weigth)
	}else{
		fmt.Println("Over-Weight")
	}
}