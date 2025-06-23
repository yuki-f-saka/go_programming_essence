package main

import "fmt"

func main() {
	// Goのpanicはシステムの強制停止に使うため、ユーザーに正しいエラーメッセージを伝えるのには不向き

	// 復旧可能な例外と、復旧不可能な例外があり、以下は復旧可能な例

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T: %[1]s\n", e)
			// panic関数 -> string: my error
			// ランタイムのpanic -> runtime.boundsError: runtime error: index out of range [2] with length 2
		}
	}()

	/*
		上記のrecoverがないとpanicになる
		panic: runtime error: index out of range [2] with length 2

		goroutine 1 [running]:
		main.main()
		        /Users/funesaka/work/go/go_programming_essence/chapter3/0529/01/main.go:17 +0x24
		exit status 2
	*/

	// 配列の境界外アクセス
	// var a [2]int
	// n := 2
	// println(a[n])

	// // ゼロ除算
	// var n int
	// println(1 / n)

	// // nilポインタのでリファレンス
	// var p *int
	// println(*p)

	panic("my error")
}
