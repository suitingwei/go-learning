package main

import (
	"./crawl"
	"fmt"
)

func main() {

	fmt.Println("This is a small program to crawl the awesome-go project")

	saver := crawl.NewFileSaver("./crawl-result")

	downloader := crawl.NewDownloader(saver, false)

	downloader.Serve(":8888")
}
