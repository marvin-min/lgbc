package main

import (
	"runtime"
	"time"
	"fmt"
)

var num = 14
var cnum chan int

func main() {
	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)

	cnum = make(chan int, num)

	for i := 0; i < num; i++ {
		go Printer(i)
	}

	for i:=0; i< num; i++{
		<-cnum
	}

	fmt.Println("=======Done========")

}

func Printer(i int)  {
	time.Sleep(20*time.Second)
	fmt.Printf("I am %d\r\n",i)
	cnum <- 1
}
