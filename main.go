package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// Monitor is an interface for a Monitor
type Monitor interface {
	FeedList() [][]string
}

// Feed is a struct for feed information
type Feed struct {
	config string
}

func main() {
	configFile := os.Args[1]
	var feed Monitor = Feed{config: configFile}
	var feedUrls = feed.FeedList()
	log.Println("Starting up")
	log.Println("Reading config from", configFile)
	log.Println("Read in", len(feedUrls), "feeds")

}

// FeedList returns a list of feed URLs from a file
func (f Feed) FeedList() [][]string {
	var feeds [][]string
	filePath := f.config
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	csvReader := csv.NewReader(reader)
	for {
		feed, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		feeds = append(feeds, feed)
	}

	return feeds
}
