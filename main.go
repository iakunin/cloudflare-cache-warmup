package main

import (
	"fmt"
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
	items := getItems(os.Stdin)
	for i, item := range items {
		if item.Status == "publish" {
			fmt.Printf(
				"(%d of %d) Processing url = %s \n",
				i,
				len(items),
				item.Permalink,
			)

			processUrl(item.Permalink)
		}
	}
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

		fmt.Println("CF-Cache-Status = ", cacheStatus)
		if cacheStatus == "HIT" {
			break
		}
	}
}
