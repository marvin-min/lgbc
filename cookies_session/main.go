package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	cookie := http.Cookie{Name:"hello",Value:"world"}
	http.SetCookie(w,&cookie)
	for _, cookie := range r.Cookies(){
		fmt.Fprintln(w,cookie.Name+"="+cookie.Value)
	}
}