package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Day02==> Maps")
	// its associative datastructure
	// 1.creating map
	m:=make(map[string]string)
	// 2.addign the eelments
	m["name"]="abdulhamid"
	m["role"] = "golang developer"
	m["working"] = "hexaware"


	// 3.deleting the element from the mao
	delete(m,"working")
	fmt.Println(m)

	// another way
	mapp:=map[string]int{"salary":40000000,"shift_time":1}
	mapp2:=map[string]int{"salary":40000000,"shift_time":1}
	fmt.Println(mapp)
	

	// to checck if the elemnet is present or not
	// use simple if else
	value,data:=mapp["salary"];
	if data{
		fmt.Println("data is preseent with the value of ",value)
	}else{
		fmt.Println("no data found.")
	}

	// to check if both the maps are equal
	fmt.Println(maps.Equal(mapp,mapp2))

}