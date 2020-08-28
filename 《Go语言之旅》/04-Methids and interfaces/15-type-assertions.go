package main

import "fmt"

// 类型断言
// t, ok := i.(T)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic: interface conversion: interface {} is string, not float64

	fmt.Println(f)
}