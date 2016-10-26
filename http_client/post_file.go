package main

import (
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"strconv"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	filename = "/Users/marvinmin/Downloads/RubyMine-2016.2.4.dmg"

	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	fh, err := os.Open(filename)

	if err != nil {
		fmt.Println("error open file")
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("error copy file")
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main() {
	url := "http://localhost:8080/upload"

	for i := 0; i < 100; i++ {
		go fmt.Println(postFile("file" + strconv.Itoa(i),url))
	}
}
