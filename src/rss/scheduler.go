package rss

import (
    "log"
    "time"

    "newsletter-aggregator/src/config"
    "newsletter-aggregator/src/db/sqlite"
    "newsletter-aggregator/src/curation"

    "github.com/robfig/cron/v3"
)

func StartRSSFetcherScheduler(db *sqlite.SQLiteDB, conf *config.AppConfig) {
    c := cron.New(cron.WithLocation(time.UTC))

    // Run at midnight every day in GMT+0
    _, err := c.AddFunc("0 0 * * *", func() {
        log.Println("Scheduled task started!")
        err := db.ClearPreviousArticles()
        if err != nil {
            log.Printf("Error clearing previous articles: %v", err)
            return
        }

        articles, err := FetchRSSArticles(conf.RSSFetchLimit)
        if err != nil {
            log.Printf("Error fetching articles: %v", err)
            return
        }

        for _, article := range articles {
            categories := curation.Categorize(article.Description)
            article.Category = sqlite.CategoriesToStrings(categories)

            _, err := db.SaveArticle(article)
            if err != nil {
                log.Printf("Error saving article %s: %v", article.Title, err)
            }
        }
        log.Println("Scheduled task ended!")
    })

    if err != nil {
        log.Fatalf("Failed to schedule RSS fetcher: %v", err)
    }

    c.Start()
}