package main

import (
	"net/http"
	"github.com/alecthomas/template"
	"fmt"
	"log"
	"strconv"
)

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Printf("%s = %v\n", k, v)
	}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/Users/marvinmin/projects/go/src/go-web/form/login.html")
		t.Execute(w, nil)
	} else {
		name := r.Form["name"][0]
		if len(name) == 0 {
			fmt.Fprintf(w, "欢迎您：%s", )
		}
		//fmt.Println(w, reflect.TypeOf(name))

		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Fprintf(w, "年龄输入不正确")
			return
		} else if age > 60 || age < 18 {
			fmt.Fprintf(w, "年龄必须在18~60岁之间")
			return
		}
		province := r.Form.Get("province")
		if ok := validProvince(province);ok {
			fmt.Fprintf(w, "欢迎您：%s, 您来自 %s 省", template.HTMLEscapeString(name),province)
		} else {
			fmt.Fprintf(w, "无效省份")
		}

	}
}

func validProvince(province string) bool {
	slice := []string{"陕西", "北京"}
	for _, v := range slice {
		if v == province {
			return true
		}
	}
	return false
}
