package main 

import "fmt"

// for 循环的 range 形式可遍历切片或映射

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128 }

func main() {
	for index, value := range pow {
		fmt.Printf(" 2 ^ %d = %d \n", index, value)
	}
}