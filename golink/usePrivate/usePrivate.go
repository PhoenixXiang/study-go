package usePrivate

import (
	"fmt"

	_ "github.com/PhoenixXiang/study-go/golink/private"
)

func usePrivateFunc()

var usePrivateStr string

func Call() {
	usePrivateFunc()
	fmt.Printf("private str = %s", usePrivateStr)

}
