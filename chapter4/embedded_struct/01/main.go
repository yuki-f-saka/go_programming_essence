package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

type Teacher struct {
	Attr
	Subject string
}

type Student struct {
	Attr
	Score int
}

func main() {

	teacher := &Teacher{
		Attr: Attr{
			Name: "Ben",
			Age:  39,
		},
		Subject: "history",
	}

	student := &Student{
		Attr: Attr{
			Name: "Mia",
			Age:  17,
		},
		Score: 89,
	}

	fmt.Println(teacher.Name, teacher.Subject)
	fmt.Println(student.Name, student.Score)
}
