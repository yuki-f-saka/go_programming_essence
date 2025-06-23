package main

import (
	"fmt"
	"time"
)

/*
問題①：2つのgoroutineから同時にメッセージを送って受け取る
目的：goroutineとchannelの基本

// 2つのgenerator関数を作って、それぞれ違うメッセージを1秒ごとに送る。
// main関数では10回分のメッセージを受信して出力する。
*/
func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	ch1 := generator("hello")
	ch2 := generator("bye")
	for i := 0; i < 10; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
