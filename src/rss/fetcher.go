package rss

import (
    "log"
    "sync"

    "github.com/mmcdole/gofeed"
    "newsletter-aggregator/src/db"
)

func FetchRSSArticles(limit int) ([]db.Article, error) {
    var wg sync.WaitGroup
    fp := gofeed.NewParser()
    articleChan := make(chan []db.Article)
    var articles []db.Article

    for _, url := range FeedURLs {
        wg.Add(1)

        go func(rssURL string) {
            defer wg.Done()

            var localArticles []db.Article
            feed, err := fp.ParseURL(rssURL)
            if err != nil {
                log.Printf("Error fetching the feed from %s: %v", rssURL, err)
                return
            }

            for i, item := range feed.Items {
                if i >= limit {
                    break
                }

                article := db.Article{
                    Title:       item.Title,
                    Description: item.Description,
                    Link:        item.Link,
                }
                localArticles = append(localArticles, article)
            }
            articleChan <- localArticles
        }(url)
    }

    go func() {
        wg.Wait()
        close(articleChan)
    }()

    for a := range articleChan {
        articles = append(articles, a...)
    }

    return articles, nil
}