package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s.\r\n", h.name, h.phone)
}
func (h *Employee) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s, I work for %s.\r\n", h.name, h.phone, h.company)
}

func main() {
	mark := Student{Human{"mark", 25, "2222-2222-2222"}, "TingsHua"}
	sam := Employee{Human{"Sam", 32, "1111-1111-1111"}, "Google"}
	mark.SayHi()
	sam.SayHi()
}
