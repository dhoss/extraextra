package monitor

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MonitorTestSuite struct {
	suite.Suite
	Feed Feed
}

var expectedCsvLine = `first_name,last_name`

/*"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`*/

func (suite *MonitorTestSuite) SetUpTests() {

	suite.Feed = Feed{}
}

func (suite *MonitorTestSuite) TestFeedList() {
	var expectedFeedList [][]string
	var expectedFeedItem []string
	expectedFeedItem = append(expectedFeedItem, "first_name")
	expectedFeedItem = append(expectedFeedItem, "last_name")
	expectedFeedList = append(expectedFeedList, expectedFeedItem)

	var feedListActual = suite.Feed.FeedList(strings.NewReader(expectedCsvLine))
	suite.Equal(expectedFeedList, feedListActual)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMonitorTestSuite(t *testing.T) {
	suite.Run(t, new(MonitorTestSuite))
}
