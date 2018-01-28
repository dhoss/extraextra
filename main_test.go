package main

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/dhoss/mendicantbias/monitor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

type stubConfig struct {
	Config Config
}

type badStubConfig struct {
	Config Config
}

type badStubFeed struct {
	Monitor Monitor
}

func (config stubConfig) ReadFile(configFile string) (io.Reader, error) {
	return strings.NewReader(`test,url`), nil
}

func (config badStubFeed) FeedList() ([][]string, error) {
	err := errors.New("Hurrr")
	return nil, err
}

func (suite *MainTestSuite) TestMainSuccess() {
	config = stubConfig{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "test.csv"}
	reader, _ := config.ReadFile(os.Args[1])
	feed = monitor.Feed{FeedURLReader: reader}

	assert.NotPanics(suite.T(), func() { main() })
}

func (suite *MainTestSuite) TestMainNoConfigFileFailsProperly() {
	config = stubConfig{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", ""}
	reader, _ := config.ReadFile(os.Args[1])
	feed = monitor.Feed{FeedURLReader: reader}
	assert.Panics(suite.T(), func() { main() })
}

func (suite *MainTestSuite) TestMainFeedListErrorFailsProperly() {
	config = stubConfig{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "test"}
	reader, _ := config.ReadFile(os.Args[1])
	feed = badStubFeed.Feed{FeedURLReader: reader}
	assert.Panics(suite.T(), func() { main() })
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
