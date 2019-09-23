package main

import "github.com/PhoenixXiang/study-go/golink/usePrivate"

/*
Notes：
1. //go:linkname 不能通过 main 直接访问（试了不行，不清楚为什么）
2. //go:linkname 所在文件要 import  _ "unsafe"
3. 函数签名文件要 import 私有函数所在包
4. 函数签名文件夹要建立后缀为 .s 的空文件，原因在于Go在编译的时候会启用 -complete 编译器flag，它要求所有的函数必需包含函数体，创建一个空的汇编语言文件绕过这个限制。
*/
func main() {
	usePrivate.Call()
}
