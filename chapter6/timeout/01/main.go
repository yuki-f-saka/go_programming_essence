package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := generator("Hello")
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case s := <-ch:
// 			fmt.Println(s)
// 		case <-time.After(time.Second):
// 			fmt.Println("waited too long")
// 			return
// 		}
// 	}
// }

func main() {
	ch := generator("Hello")
	timeout := time.After(5 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("5s timeout")
			return
		}
	}
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}
