package main

import (
	"fmt"
)

func testRange() {

	for _, arg := range []int{1, 2, 3} {
		fmt.Printf("%v\n", arg)
	}

	s := []int{1, 2}
	fmt.Printf("%v\n", &s[0])
	fmt.Printf("%v\n", &s[1])

	for key, value := range s {
		fmt.Printf("key:%v\n", key)
		fmt.Printf("value:%v\n", &value)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("key:%v\n", i)
		fmt.Printf("value:%v\n", &s[i])
	}
}

func main() {

	for i := 0; i < 5; i++ {
		defer func(x int){
			fmt.Printf("%d ", x)
			fmt.Printf("%d\n", i)
		}(i)
		// defer fmt.Printf("%d ", i)
	}

}
