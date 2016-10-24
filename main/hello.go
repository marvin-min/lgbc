package main

import (
//"go-web"
)
import (
	"fmt"
)

func main() {
	dog := Dog{}
	duck := Duck{}
	dog.can()
	duck.can()
}

type Dog struct {
}

type Duck struct {

}

func (dog Dog)can()   {
	fmt.Println("Dog can catch mouse!")
}
func (duck Duck)can()   {
	fmt.Print("Duck can catch fish!")
}