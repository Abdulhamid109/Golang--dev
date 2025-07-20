// conditional statement -- only for loop is present

package main

import "fmt"


func main(){
	i:=1
	// while loop implementation using for loop
	for i<=3{
		fmt.Println(i);
		i++
	}

	// infinate loop
	for{
		if i>5 {
			break
		}
		fmt.Println(i);
		i++
	}
	// classic for loop
	for i := i; i < 20; i++ {
		fmt.Println(i);

	}

	for i:=range(2){
		fmt.Println((i))
	}
}