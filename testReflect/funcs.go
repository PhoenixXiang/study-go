package testReflect

import (
	"fmt"
	"reflect"
)

func BaseFuncs(i interface{}) {
	// TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
	fmt.Println("type: ", reflect.TypeOf(i))
	// ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
	fmt.Println("value: ", reflect.ValueOf(i))

}


