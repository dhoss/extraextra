package monitor

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
