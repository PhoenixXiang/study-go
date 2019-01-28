package testAction

import (
	"os"
	"text/template"
)

func Output() {
	t, err := template.New("examples.txt").ParseFiles("./testAction/examples.txt")
	if err != nil {
		panic(err)
	}
	// err = t.Execute(os.Stdout, "testParse string")
	// err = t.Execute(os.Stdout, 1000)
	err = t.Execute(os.Stdout, "output")
	if err != nil {
		panic(err)
	}
}