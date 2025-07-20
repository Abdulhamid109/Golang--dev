// Learning about the constants
// constant grouping
package main

import "fmt"

const name string = "golang"
func main(){
	const(
		port = 5000
		localhost="http://localhost:5000"
	)
	fmt.Println(name,port,localhost)
}