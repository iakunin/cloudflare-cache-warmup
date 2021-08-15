package main

import (
	"flag"
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"net/http"
	"os"
)

type Item struct {
	Id        string `csv:"id"`
	Title     string `csv:"Title"`
	Permalink string `csv:"Permalink"`
	Status    string `csv:"Status"`
}

func main() {
	maxGoroutines := flag.Int("goroutines", 50, "goroutines count")
	flag.Parse()

	items := filterUnpublished(
		getItems(os.Stdin),
	)
	guard := make(chan struct{}, *maxGoroutines)

	for _, item := range items {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(item *Item) {
			processUrl(item.Permalink)
			<-guard // removes an int from semaphore, allowing another to proceed
		}(item)
	}
}

func filterUnpublished(items []*Item) []*Item {
	var result []*Item

	for _, item := range items {
		if item.Status == "publish" {
			result = append(result, item)
		}
	}

	return result
}

func getItems(reader io.Reader) []*Item {
	var items []*Item

	err := gocsv.Unmarshal(reader, &items)
	if err != nil {
		log.Fatal(err)
	}

	return items
}

func processUrl(url string) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		cacheStatus := resp.Header.Get("CF-Cache-Status")

		log.Printf("Url=%s; CF-Cache-Status=%s\n", url, cacheStatus)
		if cacheStatus == "HIT" {
			break
		}
	}
}
