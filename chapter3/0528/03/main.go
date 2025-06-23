package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	// HTTPサーバーからダウンロード
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot downloadCSV: ", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b // main関数にコンテンツを送信
	}
}

func main() {
	urls := []string{
		"https://my-server.com/data01.csv",
		"https://my-server.com/data02.csv",
		"https://my-server.com/data03.csv",
	}

	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	// goroutineからコンテンツを受け取る
	for _, b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			// レコードの登録
			insertRecords(records)
		}

	}
	wg.Wait()
}
