package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func uploadFile(writer http.ResponseWriter, request *http.Request) {

	t, _ := template.ParseFiles("/Users/marvinmin/projects/go/src/go-web/http_server/upload.gtpl")
	method := request.Method
	if method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t.Execute(writer, token)
	} else if method == "POST" {
		request.ParseMultipartForm(32 << 20)
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(writer, "%v", handler.Header)
		name := strings.Split(handler.Filename, "/")
		f, err := os.OpenFile("/Users/marvinmin/projects/go/src/go-web/http_server/"+name[len(name)-1], os.O_WRONLY | os.O_CREATE, 066)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)

	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Fprintf(writer, "Hello:" + request.FormValue("name"))
}
