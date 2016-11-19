package main

import (
	"crypto/sha256"
	"io"
	"fmt"
	"crypto/sha1"
	"crypto/md5"
	"encoding/base64"
)

func main() {
	hello := `您好啊啊  啊啊啊 ！`
	debyte := base64Encode([]byte(hello))
	rs, _ := base64Decode(debyte)
	fmt.Println(string(rs))
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte)([]byte, error)  {
	return base64.StdEncoding.DecodeString(string(src))
}

func basic()  {
	h := sha256.New()
	io.WriteString(h,"hello")
	//第二个参数相当于是盐值。用来混淆
	fmt.Printf("%x\r\n",h.Sum(nil))

	x := sha1.New()
	io.WriteString(x,"adfadsfsadfdsf")
	fmt.Printf("%x\r\n",x.Sum(nil))

	md5 := md5.New()
	io.WriteString(md5,"adfadsfsadfdsf")
	fmt.Printf("%x\r\n",md5.Sum(nil))
}
