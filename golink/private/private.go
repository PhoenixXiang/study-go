package a

import (
	"fmt"
	_ "unsafe"
)

//go:linkname privateStr github.com/PhoenixXiang/study-go/golink/usePrivate.usePrivateStr
var privateStr = "this is privateStr\n"

//go:linkname privateFunc github.com/PhoenixXiang/study-go/golink/usePrivate.usePrivateFunc
func privateFunc() {
	fmt.Print("this is privateFunc\n")
}
