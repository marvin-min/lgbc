package main

import (
	"net"
	"go-web/utils"
	"time"
	"fmt"
	"runtime"
	"log"
	"os"
	"strings"
)

var count int = 0

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	utils.CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckErr(err)
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	fmt.Println("total", cpus, "run")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn.SetReadDeadline(time.Now().Add(time.Second * 30))
		go hanndleConn(conn)

	}
}
func hanndleConn(conn net.Conn) {
	var buf = make([]byte, 10)
	log.Println("start to read from conn")
	var text int;
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		text = n
		log.Printf("read %d bytes, content is %s\n", text, string(buf[:text]))
		if strings.EqualFold(string(buf[:text]),"bye\r\n") {
			conn.Write([]byte("Bye"))
			conn.Close()
			os.Exit(1)
		}else{
			conn.Write((buf[:text]))
		}
	}

}