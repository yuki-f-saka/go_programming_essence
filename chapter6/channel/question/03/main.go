/*
🔹 問題③：fan-out（1つのチャネルを複数goroutineで処理）
目的：複数のworkerへの分散処理

// 1つのジョブチャネルから複数のworkerが処理を受け取る。
// 例えば1〜20までの数字を送って、2倍にして結果チャネルに送る。
*/

/*
✅ fan-outが使われる代表的な場面
	•	並列APIアクセス
	•	ファイル・画像処理
	•	ジョブキュー処理
	•	ストリーム / イベント処理

✅ 主なメリット
	•	並列化による高速化
	•	負荷分散による安定化
	•	柔軟なスケーリングと構造のシンプルさ
*/

package main

import (
	"fmt"
	"time"
)

// 1秒ごとにカウントアップで1から20までの数値を送るgenerator
func generator() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // 終了を明示
	}()
	return ch
}

// generatorで詰め込んだ数字を取り出すためのチャネルを引数にし、2倍にして新たなチャネルに詰めて返却
func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d received %d\n", id, job)
		result <- 2 * job
	}
}

func main() {
	jobs := generator()
	results := make(chan int)

	// 3つのworkerを起動(fanout)
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 20この結果を受信して表示
	for i := 0; i < 20; i++ {
		fmt.Println(<-results)
	}
}
