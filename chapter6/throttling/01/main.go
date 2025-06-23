package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	// download JSON
	println(u)
	time.Sleep(time.Second * 6)
}

// 600秒かかる処理
// func main() {
// 	before := time.Now()
// 	for i := 1; i < 100; i++ {
// 		u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
// 		downloadJSON(u)
// 	}

// 	fmt.Printf("%v\n", time.Since(before))
// }

// ↓　goroutineで並列にする

// ver2
/*
func main() {
	before := time.Now()

	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			downloadJSON(u)
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}
*/

// throttlingを導入して同時に20とする
// 20個ずつ処理して、1ブロックで6秒なので合計30秒かかる
func main() {
	before := time.Now()

	limit := make(chan struct{}, 20) // サイズ20
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		i := i

		go func() {
			// limitが20よりも多いとブロックする
			limit <- struct{}{}

			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			downloadJSON(u)

			// limitから抜き出す
			<-limit
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}
