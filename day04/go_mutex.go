// its an sync. tech where we lock the go routines...(for more updates visit book , docs)
package main

import (
	"fmt"
	"sync"
)


type view struct{
	views int
	mu sync.Mutex
}

func (obj *view) increment(wg *sync.WaitGroup){
	defer func ()  {
		wg.Done()
		obj.mu.Unlock()
	}()

	obj.mu.Lock()
	obj.views+=1;
}

func main(){
	var wg sync.WaitGroup;
	mydata:=view{
		views: 0,
	}
	for i:=0;i<100;i++{
		wg.Add(1)
		go mydata.increment(&wg)
	}
	wg.Wait()

	fmt.Println(mydata.views)
}