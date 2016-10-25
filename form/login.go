package main

import (
	"net/http"
	"github.com/alecthomas/template"
	"fmt"
	"log"
	"reflect"
)

func main() {
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k,v := range r.Form {
		fmt.Printf("%s = %v\n",k,v)
	}
	if r.Method == "GET" {
		t,_ := template.ParseFiles("/Users/marvinmin/projects/go/src/go-web/form/login.html")
		t.Execute(w,nil)
	}else{
		fmt.Fprintf(w, "欢迎您：%s",r.Form["name"])
		fmt.Print(reflect.TypeOf(r.Form["name"]))
	}
}
