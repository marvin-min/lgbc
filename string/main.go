package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(strings.Contains("seafood","food"))
	s := []string{"foo","bar","baz"}
	fmt.Println(strings.Join(s,";"))
	fmt.Println(strings.Index("hello","l"))
	fmt.Println(strings.Repeat("*",100))
	fmt.Println(strings.Replace("hello","l","i",-1))
	fmt.Printf("%q\n",strings.Split("hello-world","-"))
	fmt.Printf("%q\n",strings.Split("hello-world",""))
	fmt.Printf("%q\n",strings.Trim("!helllo ! world !","!"))
	fmt.Println(strings.Fields(" foo bar zaaa "))

	str := make([]byte,0,100)
	str = strconv.AppendInt(str,4567,10)
	str = strconv.AppendBool(str,false)
	str = strconv.AppendQuote(str,"^^^^^^^GGGGGHHH˙˙˙˙©©©ƒƒƒƒƒ")
	str = strconv.AppendQuoteRune(str,'¶')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(1234.1234,'g',12,64)
	c := strconv.FormatInt(1234,10)
	fmt.Println(a,b,c)


}
