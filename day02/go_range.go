package main

import "fmt"

func main() {
	fmt.Println("Range--> iterating over the data structres")


	// slcie
	data:=[]int{1,2,3}
	for i:=0;i<len(data);i++{
		fmt.Println(data[i])
	}

	// iterating with the help og range 
	for i:=range(len(data)){
		fmt.Println(i,data[i])
	}

	// or
	for i,num:= range(data){
		fmt.Println(num,i)
	}

	// iterating over maps
	m:=map[string]int{"no":1,"no2":2}
	for key,val:=range(m){
		fmt.Println(key,val)
	}
}