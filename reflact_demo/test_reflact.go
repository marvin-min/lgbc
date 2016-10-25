package main

import (
	"reflect"
	"fmt"
)

type Person struct {
	name string
	age int

}
func main() {
	p := Person{"Marvin",29}
	t := reflect.TypeOf(3)
	v := reflect.ValueOf(&p)
	tag := v.Elem().Field(0)
	fmt.Print(tag)
	fmt.Print(t,v)
}
