package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("./cmd/exp/hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
	}{
		Name: "Susan Smith",
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
