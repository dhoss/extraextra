package monitor

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MonitorTestSuite struct {
	suite.Suite
}

/*
type stubConfig struct {
	Config Config
}


func (config stubConfig) ReadFile(configFile string) (io.Reader, error) {
	return strings.NewReader(expectedCsvLine), nil
}
*/

var expectedCsvLine = `first_name,last_name`

func (suite *MonitorTestSuite) TestFeedList() {
	var expectedFeedList [][]string
	var expectedFeedItem []string
	expectedFeedItem = append(expectedFeedItem, "first_name")
	expectedFeedItem = append(expectedFeedItem, "last_name")
	expectedFeedList = append(expectedFeedList, expectedFeedItem)

	feed := Feed{FeedURLReader: strings.NewReader(expectedCsvLine)}

	feedListActual, err := feed.FeedList()
	assert.Equal(suite.T(), expectedFeedList, feedListActual)
	assert.Nil(suite.T(), err)
}

func (suite *MonitorTestSuite) TestFeedListFails() {
	var badCsv = `bad,more,badstuff
	badrow`

	feed := Feed{FeedURLReader: strings.NewReader(badCsv)}
	_, err := feed.FeedList()
	assert.Error(suite.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMonitorTestSuite(t *testing.T) {
	suite.Run(t, new(MonitorTestSuite))
}
