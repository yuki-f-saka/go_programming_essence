package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
}

func main() {
	user := User{
		Name: "Bob",
	}
	tmpl := `<a href="{{.Name}}">link</a>` //
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}

}
