package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dhoss/mendicantbias/monitor"
)

var feed monitor.Feed
var feeds [][]string

// Config is an interface to provide configuration
type Config interface {
	ReadFile(string) (io.Reader, error)
}

// FeedConfig takes a string path to a csv file with URLs
type FeedConfig struct {
	Config Config
}

var config Config = FeedConfig{}

var feed Feed

func main() {
	// Would be great to just pass config and have Feed do the right thing
	configFile := os.Args[1]
	feedReader, err := config.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
	}

	feed = monitor.Feed{FeedURLReader: feedReader}

	feeds, err := feed.FeedList()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting up")
	log.Println("Reading URLs from", configFile)
	log.Println("Read in", len(feeds), "feeds")
}

// ReadFile reads in a file of URLs and their descriptions
// and returns an io.Reader
func (config FeedConfig) ReadFile(configFile string) (io.Reader, error) {
	urls, err := ioutil.ReadFile(configFile)

	return strings.NewReader(string(urls[:])), err
}
