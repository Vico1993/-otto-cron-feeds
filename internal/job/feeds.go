package job

import (
	"fmt"
	"strings"

	textrank "github.com/DavidBelicza/TextRank/v2"
	"github.com/Vico1993/Otto-client/otto"
	"github.com/mmcdole/gofeed"
)

var (
	// Default Rule for parsing.
	rule = textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language = textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef = textrank.NewDefaultAlgorithm()
)

// Extracting the Parsing url to a new var
var parseUrl = func(url string) (*gofeed.Feed, error) {
	return gofeedParser.ParseURL(url)
}

// Parse one Rss Feed
func parse(client otto.Client, oFeed otto.Feed) {
	fmt.Println("FeedJob - Start : " + oFeed.Url)

	feed, err := parseUrl(oFeed.Url)
	if err != nil {
		fmt.Println("Couldn't parsed " + oFeed.Url + ": " + err.Error())
		return
	}

	for _, item := range feed.Items {
		item := item

		itemTags := item.Categories
		if len(itemTags) == 0 {
			itemTags = findTagFromTitle(item.Title)
		}

		if isArticleAlreadyAdded(client, item.Title, oFeed) {
			continue
		}

		author := "Unknown"
		if len(item.Authors) > 0 {
			author = item.Authors[0].Name
		}

		client.Article.Create(
			oFeed.Id,
			item.Title,
			feed.Link,
			author,
			item.Link,
			cleanCategories(itemTags),
		)
	}

	fmt.Println("FeedJob - End : " + oFeed.Url)
}

// Clean categories from rss feed before going in db
func cleanCategories(categories []string) []string {
	cats := []string{}

	for _, category := range categories {
		category := strings.ToLower(category)
		if strings.Contains(category, " ") {
			cats = append(cats, strings.Split(category, " ")...)
		} else if strings.Contains(category, ",") {
			cats = append(cats, strings.Split(category, ",")...)
		} else {
			cats = append(cats, category)
		}
	}

	return cats
}

// Extract important word from the title
func findTagFromTitle(title string) []string {
	// TextRank object
	tr := textrank.NewTextRank()
	// Add text.
	tr.Populate(title, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all words order by weight.
	words := textrank.FindSingleWords(tr)

	var tags []string
	for _, word := range words {
		tags = append(tags, strings.ToLower(word.Word))
	}

	return tags
}

// isArticleAlreadyAdded look into the DB to find if it's a new article...
func isArticleAlreadyAdded(client otto.Client, title string, feed otto.Feed) bool {
	found := false
	for _, feed := range client.Feed.ListArticles(feed.Id) {
		if feed.Title == title {
			found = true
			break
		}
	}

	return found
}
