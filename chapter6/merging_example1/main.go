package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 疑似的なニュース構造体
type News struct {
	Source string
	Title  string
}

// 疑似GoogleニュースAPI
func googleNews() <-chan News {
	ch := make(chan News)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
			ch <- News{Source: "Google", Title: fmt.Sprintf("Google News #%d", i)}
		}
		close(ch)
	}()
	return ch
}

// 疑似YahooニュースAPI
func yahooNews() <-chan News {
	ch := make(chan News)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
			ch <- News{Source: "Yahoo", Title: fmt.Sprintf("Yahoo News #%d", i)}
		}
		close(ch)
	}()
	return ch
}

// fanIn: 複数のニュースソースを1つのchannelに合流
func fanIn(channels ...<-chan News) <-chan News {
	out := make(chan News)

	for _, ch := range channels {
		ch := ch // goroutineでクロージャがキャプチャする値を固定化
		go func() {
			for n := range ch {
				out <- n
			}
		}()
	}

	return out
}

func main() {
	rand.New(rand.NewSource((time.Now().UnixNano()))) // ランダムなsleepのために必要

	// ニュースソース2つを合流
	google := googleNews()
	yahoo := yahooNews()
	merged := fanIn(google, yahoo)

	// すべてのニュースを受信（最大10件）
	for i := 0; i < 10; i++ {
		n := <-merged
		fmt.Printf("[%s] %s\n", n.Source, n.Title)
	}
}
