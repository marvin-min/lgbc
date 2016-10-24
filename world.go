package go_web

import "fmt"

const KEY = "89adfl"

func Say(something string) {
	name, _ := something, "asdfsa"
	fmt.Printf("您好: %s\n", name);
	name = "asdfa"
	c := []byte(name)
	c[3] = '9'
	name = "=" + name
	fmt.Println(string(name))
	rating := map[string]float32{"c":5, "GO":4.5, "python":6, "c++":2}
	csharp, ok := rating["c#"]
	if ok {
		fmt.Println("the rating of c# is:", csharp)
	}
	delete(rating, "c")
	x := 3
	fmt.Println(add(&x))
	fmt.Println(x)

	slice := []int {1,2,3,4,5,6,7,8,9}
	odd := filter(slice,isOdd)
	fmt.Println(odd)

}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}
func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	} else {
		return false
	}
}

func filter(slice []int, f testInt)[]int  {
	var result []int
	for _, value := range slice{
		if f(value){
			result = append(result,value)
		}
	}
	return result
}


func add(a *int) int {
	defer fmt.Println("disconnection from server")
	defer fmt.Println("close connection")
	*a = *a + 1
	fmt.Println("=====before return=====")
	return *a
}

func computerRs(args ...int) (x int, y int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
	//return args[1],args[0];
	args[1] = args[1] + 1
	return args[1], args[0]
}
