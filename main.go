package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dhoss/mendicantbias/monitor"
)

func main() {
	configFile := os.Args[1]
	feedListContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Can't open", configFile, ":", err)
	}
	feedListReader := strings.NewReader(string(feedListContent[:]))
	var feed monitor.Monitor = monitor.Feed{}
	var feedUrls = feed.FeedList(feedListReader)
	log.Println("Starting up")
	log.Println("Reading config from", configFile)
	log.Println("Read in", len(feedUrls), "feeds")
}
