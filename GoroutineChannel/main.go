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
	defer close(ch) // 終わったら閉じる(5)

	// HTTPサーバからダウンロード
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannet download CSV: ", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannet read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b // main関数にコンテンツを通信(3)
	}
}

func main() {
	urls := []string{
		"http://my-server.com/data1.csv",
		"http://my-server.com/data2.csv",
		"http://my-server.com/data3.csv",
	}

	// バイト列を転送するためのchannelを作成(1)
	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch) // (2)

	// goroutineからのコンテンツを受け取る(4)
	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		r.Comma = ','
		r.LazyQuotes = true
		for {
			records, err := r.Read()
			if err != nil {
				if err == io.EOF { // EOFエラーを特別に処理
					break // ループを抜ける
				}
				log.Fatal(err)
			}
			insertRecords(records)
		}
	}
	wg.Wait()
}

type Record struct {
	Name string
}

func insertRecords(records []string) {
}
