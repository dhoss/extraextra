package monitor

import (
	"encoding/csv"
	"io"
	"log"
)

// Monitor is an interface for a Monitor
type Monitor interface {
	FeedList() [][]string
}

// Feed is a struct for feed information
type Feed struct {
	FeedURLReader io.Reader
}

// FeedList returns a list of feed URLs from a file
func (f Feed) FeedList() ([][]string, error) {
	var feeds [][]string
	csvReader := csv.NewReader(f.FeedURLReader)
	for {
		feed, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
