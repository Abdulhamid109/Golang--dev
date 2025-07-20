// arrays
package main

import "fmt"

func main() {
	var nums[5]int 
	fmt.Print("length =>",len(nums))
	nums[0] = 1
	fmt.Println(nums[0])
	for i:=range(nums){
		fmt.Print(nums[i])
	}
	fmt.Println()
	// intializing the array
	num :=[5]int{1,2,3,4,5}
	fmt.Println(num)

	// 2d arrays
	// var twoDarrays[2][2]int

	// intializing the 2d arrays
	twoDarray:=[2][2]int{{1,2},{3,4}}
	fmt.Println(twoDarray)

	// in go arrays are used when you prehand know the size==> used for memory optimization+constant time access
	// in go slicies are used when you don't know the size ==> the data is dynamic
}