package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// ブロッキングが起こりうるのは、http.Get(u)やio.ReadAll(resp.Body)やinsertRecordsなど

// 指定した複数のurlから並行でCSVを取得して内容をチャンネルに送る関数
func DownloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	// HTTPサーバーからのダウンロード
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV: ", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b // チャンネルにコンテンツを送信
	}
}

func main() {
	urls := []string{
		"http://my-server.com/data01.csv",
		"http://my-server.com/data02.csv",
		"http://my-server.com/data03.csv",
	}

	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go DownloadCSV(&wg, urls, ch)

	// goroutineからコンテンツを受け取る
	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			// レコードの登録
			// insertRecords(records)
			fmt.Println(records)
		}
	}
	wg.Wait()
}
