package  main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi()  {
	fmt.Printf("Hi ,I am %s you can call me on %s\n",h.name,h.phone)
}

func (h Human)Sing(lyrics string)  {
	fmt.Printf("La la la ...%s\n",lyrics)
}

func(e Employee) SayHi(){
	fmt.Printf("Hi ,I am %s you can call me on %s, i work for %s \n",e.name,e.phone,e.company)
}
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mark := Student{Human{"mark", 25, "2222-2222-2222"}, "TingsHua",0.01}
	sam := Employee{Human{"Sam", 32, "1111-1111-1111"}, "Google",1111}

	var men Men
	men = mark
	men.SayHi()
	men.Sing("yoyoyo")
	men = sam
	men.SayHi()
	men.Sing("hohoho")
}
