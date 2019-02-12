package main

import (
	"fmt"
)

func main() {
	// 1.
	fmt.Printf("%T\n\r", change(Token{id: 1}))

	// 2.
	t, ok := change(Token{id: 2}).(Token)
	fmt.Println(ok)
	if ok {
		t.print()
	}

	// 3.
	var test interface{} = Token{id: 3}
	// test = "test"
	switch v := change(test).(type) {
	case Token:
		v.print()
	default:
		fmt.Println(v)
	}

}

type Token struct {
	id int
}

func (t Token) print() {
	fmt.Println(t.id)
}

func change(t interface{}) interface{} {
	return t
}
