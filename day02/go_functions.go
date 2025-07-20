package main

import "fmt"

func add(a int ,b int)int{
	return a+b;
}

func data(fn func(a int)int){
	fn(1)
}

func main() {
	fn:=func (a int) int {
		return 2
	}
	data(fn)
	fmt.Println("Fuunctions in golang")
	fmt.Println("addition==>",add(20,10))
	// im go the functions are first class citizen functions
	// i.e you can assign it to a variable, can pass as args in another functions
}