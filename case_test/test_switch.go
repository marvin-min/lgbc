package main

import (
	"fmt"
	//"strconv"
	"strconv"
)

type Element interface{}
type List[] Element

type Person struct {
	name string
	age int
}



func (p Person)String() string{
	return "name = " + p.name + "; age = " + strconv.Itoa(p.age)
}

func main() {
	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Sam",23}

	for index, element := range list {
		switch value := element.(type) { //element.(type) 不能在switch外的任何逻辑调用,如果需要在switch外调用则必须使用 if value, ok := men.(Employee) 形式
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n",index,value)
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %s\n",index,value)
		default:
			fmt.Printf("list[%d] is off a is %s\n",index,value)
		}
	}
}
