package main

import (
	"fmt"
)

func main() {
	// 1.
	fmt.Println("1-------")
	fmt.Printf("%T\n\r", change(Token{id: 1}))

	// 2.
	fmt.Println("2-------")
	t, ok := change(Token{id: 2}).(Token)
	fmt.Println(ok)
	if ok {
		t.print()
	}

	// 3.
	fmt.Println("3-------")
	var test interface{} = Token{id: 3}
	// test = "test"
	switch v := change(test).(type) {
	case Token:
		v.print()
	default:
		fmt.Println(v)
	}

	fmt.Println("4-------")
	t1 = Token{id: 4}
	fmt.Println(t1)

}

// 用于验证Token是否实现了Test接口
// 1.内存开销小
var t1 Test = (*Token)(nil)
// 2.
var t2 Test = &Token{}

type Test interface {
	print()
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
