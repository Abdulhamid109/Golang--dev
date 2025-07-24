package main

import (
	"fmt"
	"time"
)

func dummy(dum chan int){
	data:=<-dum //--data receiving from the channels
	fmt.Println("This is the data came between the go route =>",data)
}

func main() {
	fmt.Println("Demonstration of go channels....")
	newChannel:=make(chan int);
	go dummy(newChannel)
	newChannel<-12 //--data sending to the channels
	time.Sleep(time.Second*2)

}