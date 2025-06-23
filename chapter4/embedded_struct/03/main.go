package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

func (a Attr) String() string {
	return fmt.Sprintf("%s(%d)", a.Name, a.Age)
}

type AttrEx struct {
	Name string
}

func (a AttrEx) String() string {
	return fmt.Sprintf("(a.k.a %s)", a.Name)
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

	fmt.Println(teacher.Attr.String(), teacher.AttrEx.String())
}
