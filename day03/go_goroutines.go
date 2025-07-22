// implementing the go routines

package main

import (
	"fmt"
	"sync"
)

// implemetation of waitgroups

func tasks(id int,wg *sync.WaitGroup){
	defer wg.Done()
fmt.Println("indexes with values ",id)

}

func main(){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		// creating a anonomous function
		go tasks(i,&wg)
	}

	wg.Wait()

}