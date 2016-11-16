package main

import (
	"sync"
	"runtime"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func main() {
	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)
	for i:= 0 ; i<100;i++ {
		wg.Add(1)
		go Printer(i)
	}
	wg.Wait()
	fmt.Println("DONE!")
}

func Printer(i int)  {
	time.Sleep(20*time.Second)
	fmt.Printf("I am %d\r\n",i)
	defer wg.Done()
}
