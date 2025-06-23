package main

import (
	"fmt"
	"time"
)

/*
	generatorのようなチャンネルを返す関数の書き方は、
	非同期で値を生成し続ける処理や、
	ストリーム処理を簡単に実装したいときに便利です。

	どんなときに嬉しいか
		バックグラウンドで値を生成し続けたいとき
		例：センサー値の取得、ログの監視、イベントの購読など

		複数の消費者（受信側）に値を渡したいとき
		例：ワーカープールやファンアウト処理

		イテレータのように使いたいとき
		例：forループで順次値を受け取りたい場合

		非同期で重い処理を実行し、その結果を受け取りたいとき
		例：APIリクエストやバッチ処理の結果を順次受け取る

	メリット
	呼び出し側は「値を受け取るだけ」でよく、生成側の実装を意識しなくてよい
	並行処理が簡単に書ける
	Goらしいシンプルな非同期設計ができる
*/

// 戻り値の<-chan stringはチャンネルが文字列の受信専用ですよということを表している
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
	ch := generator("Hello!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
