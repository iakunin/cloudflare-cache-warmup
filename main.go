package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type WpPost struct {
	Id        string `csv:"id"`
	Title     string `csv:"Title"`
	Permalink string `csv:"Permalink"`
	Status    string `csv:"Status"`
}

func main() {
	path := os.Args[1]
	posts := getWpPosts(path)

	for i, post := range posts {
		if post.Status == "publish" {
			fmt.Printf(
				"(%d of %d) Processing url = %s \n",
				i,
				len(posts),
				post.Permalink,
			)

			processUrl(post.Permalink)
		}
	}
}

func getWpPosts(path string) []*WpPost {
	in, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func(in *os.File) {
		_ = in.Close()
	}(in)

	var posts []*WpPost

	err = gocsv.UnmarshalFile(in, &posts)
	if err != nil {
		log.Fatal(err)
	}

	return posts
}

func processUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	for {
		cacheStatus := resp.Header.Get("CF-Cache-Status")
		fmt.Println("CF-Cache-Status = ", cacheStatus)
		if cacheStatus == "HIT" {
			break
		}
	}
}
