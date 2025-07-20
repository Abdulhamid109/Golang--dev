package main

import "fmt"

func data() func() int {
	var count int = 0

	return func() int {
		count+=1
		return count
	}
}
func main() {
	increment := data()
	fmt.Println(increment())
}
