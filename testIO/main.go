package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	s, err := ioutil.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	s, err = ioutil.ReadFile("./dir/b.txt")
	// s, err = ioutil.ReadFile("dir/b.txt") // 同上
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func main() {
	readFile()
}
