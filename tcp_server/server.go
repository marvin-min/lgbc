package main

import (
	"net"
	"go-web/utils"
	"time"
	"fmt"
	"runtime"
	"strconv"
)

var count int = 0
func main() {
	service := ":7777"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)
	utils.CheckErr(err)
	listener, err := net.ListenTCP("tcp",tcpAddr)
	utils.CheckErr(err)
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	fmt.Println("total" ,cpus, "run")
	for {
		conn,err := listener.Accept()
		if err!=nil {
			fmt.Println(err)
			continue
		}
		conn.SetReadDeadline(time.Now().Add(time.Second * 30))
		go hanndleConn(conn)

	}
}
func hanndleConn(conn net.Conn)  {
	count = count + 1
	fmt.Println(count)
	//time.Sleep(1000)
	daytime := time.Now().String() +"\r\n" + strconv.Itoa(count)
	fmt.Println(daytime)
	conn.Write([]byte(daytime))
	defer conn.Close()
}