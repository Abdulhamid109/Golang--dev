package main

import "fmt"

func printSlice[T any](data []T){
	for _,d:= range data{
		fmt.Println(d)
	}
}

func main()  {
	fmt.Println("Implementation of generics.....")
	// stringData:=[]string{"hello","this","is","golang","programming"}
	numData :=[]int{1,2,3,4}
	printSlice(numData)
}