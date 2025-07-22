package main

import "fmt"


func refernceData(d *int){
	*d=2001
	fmt.Println("Reference => ",*d)
}
func main()  {
	fmt.Println("Lets get Started with Pointers")
	data:=101
	refernceData(&data)
	fmt.Println("Value Changes here also => ",data)
}