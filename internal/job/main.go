package job

import (
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/go-co-op/gocron"
	"github.com/mmcdole/gofeed"
)

var feedsTag = "feed"
var gofeedParser = gofeed.NewParser()
var Scheduler = gocron.NewScheduler(time.UTC)

// Calculate the delay between each job base on the number of feed
// Each feed need to be check once an hour
func getDelay(numberOfFeed int) int {
	return int(math.Round(float64(60 / numberOfFeed)))
}

func Main(client otto.Client) {
	// Get All Feeds
	feeds := client.Feed.ListAll(true)

	jobs, err := Scheduler.FindJobsByTag(feedsTag)
	// No job found but we have feeds
	// OR if we have more or less feed than before
	if (err != nil && len(feeds) > 0) || (len(feeds) != len(jobs)) {
		err := Scheduler.RemoveByTag(feedsTag)
		if err != nil {
			fmt.Println("Couldn't reset feed job")
			return
		}

		n := 1
		for _, feed := range feeds {
			feed := feed
			url, _ := url.Parse(feed.Url)

			// Start at different time to avoid parsing all feed at the same time
			when := getDelay(len(feeds)) * n

			fmt.Println("FeedJob - Adding Job -> " + feed.Url)
			_, err := Scheduler.Every(1).
				Hour().
				Tag(feedsTag).
				StartAt(time.Now().Add(time.Duration(when)*time.Minute)).
				Do(parse, client, feed)

			if err != nil {
				fmt.Println("FeedJob - Error initiate the cron for: " + url.Host + " - " + err.Error())
			}

			n += 1
		}
	}
}
