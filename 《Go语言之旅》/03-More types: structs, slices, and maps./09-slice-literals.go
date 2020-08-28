package main

import "fmt"

// [3]bool{true, true, false} ==> 这是一个数组文法：

// []bool{true, true, false}  ==> 这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

func main() {
	q := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(q)

	r := []bool {true, false, true, false, true, false,true}
	fmt.Println(r)

	s := [] struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}