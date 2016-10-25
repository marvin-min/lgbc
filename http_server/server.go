package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("ListenAndServer: ",err)
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request)  {
	request.ParseForm()
	fmt.Fprintf(writer, "Hello:"+ request.FormValue("name"))
}
