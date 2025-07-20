// switch statement
package main

import (
	"fmt"
	"time"
)


func main()  {
	i:=5

	switch i{
	case 1:
		fmt.Print("woe")
	case 3:
		fmt.Print("joe")
	default:
		fmt.Print("Found ypu")
	}

	// multiple cases in the switch
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Print("weekend")
	case time.Monday,time.Tuesday,time.Wednesday,time.Thursday,time.Friday:
		fmt.Print("working days")
	default :
	fmt.Print("don't know")
	}

}