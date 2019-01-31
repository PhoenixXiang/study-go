package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Abs() {
	p, err := filepath.Abs("./a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

func IsAbs() {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// true
	fmt.Println(filepath.IsAbs(filepath.Join(d, "a.txt")))
	// false
	fmt.Println(filepath.IsAbs("./a.txt"))
	// false
	fmt.Println(filepath.IsAbs("a.txt"))
}

func Join() {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	// E:\GoPath\src\github.com\PhoenixXiang\study-go\testPath\a.txt
	fmt.Println(filepath.Join(d, "a.txt"))
	fmt.Println(filepath.Join(d, "./a.txt"))
	fmt.Println(filepath.Join(d, "..\\dir2"))
}

func Rel() {
	// Rel函数返回一个相对路径，将basepath和该路径用路径分隔符连起来的新路径在词法上等价于targpath。也就是说，Join(basepath, Rel(basepath, targpath))等价于targpath本身。如果成功执行，返回值总是相对于basepath的，即使basepath和targpath没有共享的路径元素。如果两个参数一个是相对路径而另一个是绝对路径，或者targpath无法表示为相对于basepath的路径，将返回错误。

	basepath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("basepath:" + basepath)

	p, err := filepath.Rel(basepath, filepath.Join(basepath, "dir1/c.txt"))
	if err != nil {
		panic(err)
	}
	// dir1\c.txt
	fmt.Println(p)

	// filepath.Rel第一个参数应该为目录
	p, err = filepath.Rel(filepath.Join(basepath, "dir1/c.txt"), filepath.Join(basepath, "dir2/b.txt"))
	if err != nil {
		panic(err)
	}
	// ..\..\dir2\b.txt
	fmt.Println(p)
	// E:\GoPath\src\github.com\PhoenixXiang\study-go\testPath\dir2\b.txt
	fmt.Println(filepath.Join(filepath.Join(basepath, "dir1/c.txt"), p))

	p, err = filepath.Rel(filepath.Join(basepath, "dir1/"), filepath.Join(basepath, "dir2/b.txt"))
	if err != nil {
		panic(err)
	}
	// ..\..\dir2\b.txt
	fmt.Println(p)
	// E:\GoPath\src\github.com\PhoenixXiang\study-go\testPath\dir2\b.txt
	fmt.Println(filepath.Join(filepath.Join(basepath, "dir1/"), p))

	// E:\GoPath\src\github.com\PhoenixXiang\study-go\testPath\dir1
	// fmt.Println(filepath.Join(basepath, "dir1"))
	// E:\GoPath\src\github.com\PhoenixXiang\study-go\testPath\dir2
	// fmt.Println(filepath.Join(basepath, "dir2"))
	p, err = filepath.Rel(filepath.Join(basepath, "dir1"), filepath.Join(basepath, "dir2"))
	if err != nil {
		panic(err)
	}
	// ..\dir2
	fmt.Println(p)

	// E:\GoPath\src\github.com\PhoenixXiang\study-go\testPath\dir2
	fmt.Println(filepath.Join(filepath.Join(basepath, "dir1"), p))
}

func Walk() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	extension := ".txt"
	// fmt.Println(dir)
	// Walk函数会遍历root指定的目录下的文件树，对每一个该文件树中的目录和文件都会调用walkFn，包括root自身。所有访问文件/目录时遇到的错误都会传递给walkFn过滤。文件是按词法顺序遍历的，这让输出更漂亮，但也导致处理非常大的目录时效率会降低。Walk函数不会遍历文件树中的符号链接（快捷方式）文件包含的路径。
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
		} else {
			rel, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			fmt.Println(rel)
			ext := filepath.Ext(path)
			if ext == extension {

				buf, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				contents := string(buf)
				fmt.Println(contents)

				// ToSlash函数将path中的路径分隔符替换为斜杠（'/'）并返回替换结果，多个路径分隔符会替换为多个斜杠。
				name := filepath.ToSlash(rel)
				fmt.Println(name)
				
			}

		}
		return nil
	})
}

func main() {
	// Abs()
	// IsAbs()
	// Join()
	// Rel()
	Walk()

}
