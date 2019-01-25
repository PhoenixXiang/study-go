package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	p, err := filepath.Abs("./a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

