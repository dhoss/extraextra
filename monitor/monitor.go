package monitor

import (
	"encoding/csv"
	"io"
	"log"
)

// Monitor is an interface for a Monitor
type Monitor interface {
	FeedList(io.Reader) [][]string
}

// Feed is a struct for feed information
type Feed struct {
}

// FeedList returns a list of feed URLs from a file
func (f Feed) FeedList(feedListReader io.Reader) [][]string {
	var feeds [][]string
	csvReader := csv.NewReader(feedListReader)
	for {
		feed, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error parsing CSV", err)
		}

		feeds = append(feeds, feed)
	}

	return feeds
}
