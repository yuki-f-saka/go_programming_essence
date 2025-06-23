package main

import (
	"fmt"
	"time"
)

/*
	問題②：一定時間後に強制終了させる（select＋time.After）
	目的：selectによるタイムアウト処理の理解

	// generatorから永遠にメッセージが送られてくる。
	// ただしmain関数では3秒経過したら「タイムアウト」と表示して終了。
*/

func generator() <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", "log", i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	ch := generator()
	timeout := time.After(3 * time.Second)
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("タイムアウト")
			return
		}
	}
}
