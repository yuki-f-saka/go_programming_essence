// 🔹 問題④：バッファ付きチャネル vs 非バッファチャネルの違いを確認
// 目的：チャネルの詰まり、ブロッキングの理解

// バッファなし・バッファ付きでメッセージ送信→受信の挙動を観察してみる。
// 違いをprintで可視化する。

// 🎯 この問題の目的
// 	•	バッファ付きチャネル（make(chan T, n)）と非バッファチャネル（make(chan T)）の違いを理解する
// 	•	「送信/受信がブロックされる条件」を確認する
// 	•	実行順や動作の違いを観察して、goroutineが止まる・止まらない理由を理解する

package main

import "fmt"

func main() {
	fmt.Println("=== 非バッファチャネル ===")
	nonBuffered := make(chan int)
	go func() {
		nonBuffered <- 1
		fmt.Println("nonBufferedチャネルに1を送信した。受信がないとここで止まる")
	}()
	fmt.Println(<-nonBuffered)
	fmt.Println("受信があったので↑ブロック解除")

	fmt.Println("=== バッファ付きチャネル ===")
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println("バッファ2あるので、ここまではブロッキングされずにチャネルに値を送信できた。")
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
}
