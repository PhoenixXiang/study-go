package testReflect

import "testing"

func TestBaseFuncs(t *testing.T) {
	var num float64
	BaseFuncs(num)
	num = 1.2345
	BaseFuncs(num)
	BaseFuncs(nil)

	var u = User{1, "xyt", 123.4}
	BaseFuncs(u)
}
