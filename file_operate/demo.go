package main

import (
	"os"
	"fmt"
)

func main()  {
	os.Remove("hello.text")
	readFile()
}

func readFile(){
	userFile := "hello.text"
	file, err := os.Open(userFile)
	defer file.Close()
	if err!=nil {
		fmt.Println(userFile,err)
		return
	}
	buf := make([]byte,1024)
	for  {
		n,_ := file.Read(buf)
		if 0 ==n{
			break
		}
		os.Stdout.Write(buf[:n])
	}
}


func writeFile(){
	userFile := "hello.text"
	fout,err:= os.Create(userFile)
	defer  fout.Close()
	if err != nil {
		fmt.Println(userFile,err)
	}
	for i :=0;i<10; i++{
		fout.WriteString("Hello。这是一个中文\r\n")
		fout.WriteString("YUUJJ¥¥¥¥∆˚˚˚˚\r\n")
	}
}

func dirOp()  {
	os.Mkdir("hola",0755)
	os.MkdirAll("hola/test1/test2",0775)
	err := os.Remove("hola")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("hola")
}