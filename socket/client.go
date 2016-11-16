package main

import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
	"runtime"
	"strconv"
)

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	runtime.GOMAXPROCS(runtime.NumCPU()-4)
	 god()

}

func god()  {
	for i := 0; i < 10; i++ {
		fmt.Println("send", i, "request")
		go sendRequest(i)

	}
	select {}
	os.Exit(0)
}

func sendRequest(i int) {
	service := "localhost:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	defer conn.Close()
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println(string(result))
	writeFile(strconv.Itoa(i),string(result))
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
func writeFile(file string, content string) {
	userFile := "/Users/marvinmin/projects/go/src/" +file + ".text"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
	}
	fout.WriteString(content)
}
