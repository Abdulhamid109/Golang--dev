// slicies --> dynamic arrays

package main

import (
	"fmt"
	"slices"
)

func main(){
	// var num[]int               //uninitialized slice --> nil
	// var nums = make([]int,2);
	// third parameter --> capacity-> maximum number of element can fit
	// fmt.Println(nums)
	// appeding the elements to the slice
	// nums = append(nums, 1)
	// num:=[]int{10}
	// fmt.Print(num)

	
	// var nums2 = make([]int,(len(nums)+len(num)))
	// copy(num,nums);
	// copy+ function in slice
	// copy(nums2,num)

	// fmt.Println(nums2)


	// ---slice operator
	nums:=[]int{1,3,4,6,7}
	nums2:=[]int{1,3,4,6,7}

	fmt.Println(nums[1:3])   //--exclusive on second index and inclussive in 1st index

	// comparison
	fmt.Println(slices.Equal(nums,nums2))

	// 2d slices
	numgs:=[][]int{{123},{234}}
	fmt.Println(numgs)

	
	
}