package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 返回将所有字母都转为对应的大小写版本的拷贝。
func ToLowerAndToUpper() {
	fmt.Println(strings.ToLower("Gopher")) // gopher
	fmt.Println(strings.ToUpper("Gopher")) // GOPHER
}

// 判断字符串是否有前缀字符串prefix或后缀字符串suffix，大小写敏感。
func HasPrefixOrSuffix() {
	// 前缀
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	// 后缀
	fmt.Println(strings.HasSuffix("Gopher", "er")) // true
}

// Contains相关的一些方法
func Contains() {
	// 判断字符串s是否包含子串substr。
	fmt.Println("---判断字符串s是否包含子串substr----")
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
	fmt.Println("----------------------------------")

	// 判断字符串s是否包含字符串chars中的任一字符。
	fmt.Println("---判断字符串s是否包含字符串chars中的任一字符----")
	fmt.Println(strings.ContainsAny("team", "a"))     //true
	fmt.Println(strings.ContainsAny("team", "ea"))    //true
	fmt.Println(strings.ContainsAny("team ", " "))    //true
	fmt.Println(strings.ContainsAny("team", "e & i")) //true
	fmt.Println(strings.ContainsAny("team", "u & i")) //false
	fmt.Println(strings.ContainsAny("team", ""))      //false
	fmt.Println(strings.ContainsAny("team", "i"))     //false
	fmt.Println(strings.ContainsAny(" ", ""))         //false
	fmt.Println(strings.ContainsAny("", " "))         //false
	fmt.Println("----------------------------------------------")
}

// 返回字符串s中有几个不重复的sep子串。
func Count() {
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // 5
	fmt.Println(strings.Count("", ""))        // 1
}

// strings.Replace(s, old, new string, n int)
// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
func Replace() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
}

// Index相关的方法
func Index() {
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "k"))   // 4
	fmt.Println(strings.Index("chicken", "chi")) // 0
	fmt.Println(strings.Index("chicken", "dmr")) // -1
	fmt.Println(strings.Index("chicken", ""))    // 0
	fmt.Println("----------------------------------")

	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'c')) // 0
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1
	fmt.Println("----------------------------------")

	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1
	fmt.Println("----------------------------------")

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
	fmt.Println("----------------------------------")

	fmt.Println(strings.Index("go gopher", "go"))         // 0
	fmt.Println(strings.LastIndex("go gopher", "go "))    // 0
	fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1
	fmt.Println(strings.LastIndex("go gopher", ""))       // 9
	fmt.Println("----------------------------------")

	fmt.Println(strings.LastIndexAny("chicken", "aeiouy")) // 5
	fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))   // -1
	fmt.Println("----------------------------------")

	fmt.Println(strings.LastIndexFunc("Hello, world", f)) // -1
	fmt.Println(strings.LastIndexFunc("Hello, 世界", f))    // 10
	fmt.Println("----------------------------------")
}

// Trim相关方法
func Trim() {
	// 返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	// Trim(s string, cutset string) string
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! ")) // ["Achtung! Achtung"]
	fmt.Println("")
	fmt.Println("-----------------------")

	// 返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) // a lone gopher
	fmt.Println("-----------------------")

	// 返回将s前后端所有满足f的unicode码值都去掉的字符串。
	f := func(r rune) bool {
		return unicode.IsSpace(r)
	}
	fmt.Println(strings.TrimFunc(" \t\n a lone gopher \n\t\r\n", f)) // a lone gopher
	fmt.Println("-----------------------")

	// 返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung! Achtung! !!! ", "! ")) // ["Achtung! Achtung! !!! "]
	fmt.Println("")
	fmt.Println("-----------------------")

	// 返回去除s可能的前缀prefix的字符串。
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, "God")
	fmt.Println(s)
	s = strings.TrimPrefix(s, "Goodbye,")
	fmt.Println("Hello" + s) // Hello, world!
	fmt.Println("-----------------------")

}

// 字符串分割的相关方法
func Split() {
	fmt.Println("-----------------------")
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   ")) // Fields are: ["foo" "bar" "baz"]
	fmt.Printf("Fields are: %q\n", strings.Fields("     "))             // Fields are: []
	fmt.Println("-----------------------")

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	// Fields are: ["foo1" "bar2" "baz3"]
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	fmt.Println("-----------------------")

	// Split用分隔符切割 SplitAfter保留分隔符
	// ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	// ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	// [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	// [""]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println("-----------------------")

	// 用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片。
	// 每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割
	// 如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	// 参数n决定返回的切片的数目：
	// n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	// n == 0: 返回nil
	// n < 0 : 返回所有的子字符串组成的切片
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))   // ["a" "b,c"]
	fmt.Printf("%q\n", strings.SplitN("a,b,,c", ",", -1)) // ["a" "b" "" "c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)
	fmt.Println("-----------------------")
}

// 将一系列字符串连接为一个字符串，之间用sep来分隔。
func Join() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo, bar, baz
}
func main() {
	// ToLowerAndToUpper()
	// HasPrefixOrSuffix()
	// Contains()
	// Count()
	// Replace()
	// Index()
	// Trim()
	// Split()
	Join()
}
