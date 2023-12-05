package job

import (
	"testing"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/mock"
)

var (
	ottoFeed = otto.Feed{
		Id:  "1",
		Url: "https://test.com/feed",
	}

	item = &gofeed.Item{
		Title:      "Super Title for an Article",
		Published:  "2023-06-04",
		Link:       "https://test.com/article-1",
		Categories: []string{"tag2", "BTC"},
		Authors: []*gofeed.Person{
			{
				Name: "Victor",
			},
		},
	}

	feed = &gofeed.Feed{
		Title: "Super Test",
		Items: []*gofeed.Item{
			item,
		},
	}
)

func TestExecuteWithUpperCaseCategory(t *testing.T) {
	oldParseUrl := parseUrl
	defer func() { parseUrl = oldParseUrl }()

	parseUrl = func(url string) (*gofeed.Feed, error) {
		return feed, nil
	}

	var client otto.Client

	// client.Feed = new(MockClient)
}

type MockClient struct {
	mock.Mock
}
