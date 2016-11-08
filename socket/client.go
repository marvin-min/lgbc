package main

import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func parseIP() {
	if (len(os.Args) != 2) {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr)
	}
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
