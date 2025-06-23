package main

import (
	"fmt"
	"time"
)

func sendMessage(msg string) {
	println(msg)
}

// race condition
func main() {
	message := "hi"
	fmt.Println(message)
	// main関数の中でgoroutineを使うと、main関数の終了と共に強制終了してしまう
	// main関数がgoroutineの終了を待つためにはsyncパッケージを使う必要がある
	go func() {
		sendMessage(message)
	}()
	message = "ho"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
}
