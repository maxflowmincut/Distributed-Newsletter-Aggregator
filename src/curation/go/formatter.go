package curation

import (
    "newsletter-aggregator/src/db"
    "strings"
)

func FormatArticlesForEmail(articles []db.Article) string {
    var formattedStrings []string

    for _, article := range articles {
        formattedArticle := FormatArticle(article)
        formattedStrings = append(formattedStrings, formattedArticle)
    }

    return strings.Join(formattedStrings, "\n\n")
}

func FormatArticle(article db.Article) string {
    return "Title: " + article.Title + "\nLink: " + article.Link + "\nDescription: " + article.Description
}