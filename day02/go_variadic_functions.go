package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {
	nums:=[]int{1, 2, 3, 3, 4, 4, 5, 5, 5, 67}
	result := sum(nums...)
	fmt.Println(result)

}