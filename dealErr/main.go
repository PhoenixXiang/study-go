package main

import "fmt"

func main() {
	/*
		111 Helloworld
		33333
		Recover from r :  Panic 1!
		22222
		11111
		333 Helloworld
	*/

	defer_call()
	fmt.Println("333 Helloworld")
}

func defer_call() {
	defer func() {
		fmt.Println("11111")
	}()
	defer func() {
		fmt.Println("22222")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from r : ", r)
		}
	}()

	defer func() {
		fmt.Println("33333")
	}()

	fmt.Println("111 Helloworld")

	// panic("Panic 1!")

	// panic("Panic 2!")

	fmt.Println("222 Helloworld")
}
