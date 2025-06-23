package main

import "fmt"

// channel
func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {
	var s string

	ch := make(chan string)
	go server(ch)

	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)

}
