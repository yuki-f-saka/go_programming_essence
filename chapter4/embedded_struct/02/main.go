package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

type AttrEx struct {
	Name string
}

type Teacher struct {
	Attr
	AttrEx
	Subject string
}

func main() {
	teacher := &Teacher{
		Attr: Attr{
			Name: "Ben",
			Age:  39,
		},
		AttrEx: AttrEx{
			Name: "bbben",
		},
		Subject: "hisotory",
	}
	fmt.Println(teacher.Attr.Name)
	fmt.Println(teacher.AttrEx.Name)
}
