package main

import (
	"html/template"
	"os"
)

type Phone struct {
	Brand string
}

type Person struct {
	UserName string
	email string
	Phones []*Phone
}

func main()  {
	t := template.New("hello")
	t, _ = t.Parse(`hello: {{.UserName}}"
	 {{with .Phones }}
	 {{range .}}
	 	likes Phone brand:{{.Brand}}
	 {{end}}
	 {{end}}
	`)
	hw := Phone{"HUAWEI"}
	vivo := Phone{"VIVO"}
	p := Person{UserName:"¥¥∆µµµµ",email:"minxiao@aa.com",
		Phones: []*Phone{&hw,&vivo} }

	t.Execute(os.Stdout,p)
}
