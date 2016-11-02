package main

import (
	"fmt"
	"time"
	"runtime"
)

func forever() {
	for {
		fmt.Println(".")
	}
}

func show() {
	for {
		fmt.Println("-")
	}
}

func main() {
	n := runtime.NumCPU()
	fmt.Print("total",n,"cpu")
	runtime.GOMAXPROCS(n)
	go show()
	go forever()
	for {
		time.Sleep(1000)
	}
}





func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c<-x
		x,y = y, x+y
	}
	close(c)
}

//func main() {
//	c := make(chan int,10)
//	go fibonacci(cap(c),c)
//	for i := range c {
//		fmt.Println(i)
//	}
//}

func channelDemo() {
	var a []int = make([]int, 100000000)

	for i := 0; i < 100000000; i++ {
		a[i] += i
	}

	b := [15]int{10, 2, 3, 4, 5, 6, 7, 8, 90}
	c := make(chan int, 2)
	fmt.Println(time.Now())
	go sum(a, c)
	go sum(b[:], c)
	x, y := <-c, <-c
	fmt.Println(time.Now())
	fmt.Println(x, y)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

