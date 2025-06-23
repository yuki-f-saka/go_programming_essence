package main

import (
	"fmt"
	"math/rand"
)

// func main() {
// 	quit := make(chan bool)
// 	ch := generator("Hello", quit)
// 	for i := rand.Intn(50); i >= 0; i-- {
// 		fmt.Println(<-ch, i)
// 	}
// 	quit <- true
// }

func main() {
	quit := make(chan string)
	ch := generator("Hello", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "bye"
	fmt.Printf("generator says %s", <-quit)
}

func generator(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				quit <- "see you"
				return
			}
		}
	}()
	return ch
}
