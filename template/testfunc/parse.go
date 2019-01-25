package testfunc

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func TestParse() {
	t, err := template.New("test").Parse("name:{{.Name}}, age:{{.Age}}")
	if err != nil {
		panic(err)
	}
	// err = t.Execute(os.Stdout, "testParse string")
	// err = t.Execute(os.Stdout, 1000)
	err = t.Execute(os.Stdout, Person{Name: "test", Age: 18})
	if err != nil {
		panic(err)
	}
}

func TestParseFiles() {

	t, err := template.New("a.txt").ParseFiles("./testfunc/a.txt", "./testfunc/b.txt", "./testfunc/dir/c.txt")
	if err != nil {
		panic(err)
	}
	// ts := t.Templates()
	// fmt.Println(ts)
	err = t.Execute(os.Stdout, "Execute string\r\n")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "a.txt", "ExecuteTemplate string\r\n")
	// err = t.Execute(os.Stdout, 1000)
	// err = t.Execute(os.Stdout, Person{Name: "test", Age: 18})
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "c.txt", "ExecuteTemplate string\r\n")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "b.txt", 18)
	if err != nil {
		panic(err)
	}
}
