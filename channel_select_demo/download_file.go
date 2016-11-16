package main

import (
	"runtime"
	"time"
	"fmt"
)

var num = 10
var cnum chan int

func main() {
	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)

	cnum = make(chan int, num)

	for i := 0; i < num; i++ {
		go Printer(i)
	}

	for i:=0; i< num; i++{
		value, status := <-cnum
		fmt.Printf("Finish to deal %d \r\n",value)
		fmt.Println("status",status)
	}

	fmt.Println("=======Done========")

}

func Printer(i int)  {
	time.Sleep(3*time.Second)
	fmt.Printf("I am %d\r\n",i)
	cnum <- i
}
