package curation

import (
    "log"
    "strings"
    "time"
    "net/smtp"

    "newsletter-aggregator/src/config"
    "newsletter-aggregator/src/db/sqlite"

    "github.com/robfig/cron/v3"
)

func sendEmail(config *config.AppConfig, toEmail, subject, body string) error {
    smtpServer := config.SMTPServer
    smtpUser := config.SMTPUser
    smtpPass := config.SMTPPassword

    serverName := strings.Split(smtpServer, ":")[0]

    auth := smtp.PlainAuth("", smtpUser, smtpPass, serverName)

    msg := "From: " + smtpUser + "\n" +
        "To: " + toEmail + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    err := smtp.SendMail(smtpServer, auth, smtpUser, []string{toEmail}, []byte(msg))
    if err != nil {
        log.Println("Failed to send the email:", err)
        return err
    }

    return nil
}

func Dispatch(db *sqlite.SQLiteDB) {
    appConfig := config.LoadConfig()
    articleCount := appConfig.ArticleSendLimit

    users, err := db.GetAllUsers()
    if err != nil {
        log.Printf("Error fetching users: %v", err)
        return
    }

    currentTime := time.Now().UTC().Format("15:04")

    for _, user := range users {
        if user.SendTime+":00" == currentTime {
            articles, err := db.GetArticlesMatchingPreferences(user.Preferences, articleCount)
            if err != nil {
                log.Printf("Error fetching articles for user %s: %v", user.Email, err)
                continue
            }

            if len(articles) < articleCount {
                additionalArticles, err := db.GetRandomArticles(articleCount - len(articles))
                if err != nil {
                    log.Printf("Error fetching random articles for user %s: %v", user.Email, err)
                    continue
                }
                articles = append(articles, additionalArticles...)
            }

            formattedContent := FormatArticlesForEmail(articles)

            emailSubject := "Your Daily Newsletter"
            err = sendEmail(appConfig, user.Email, emailSubject, formattedContent)
            if err != nil {
                log.Printf("Failed to send the email to %s: %v", user.Email, err)
            }
        }
    }
}

func GetPreferencesFromString(prefString string) []string {
    return strings.Split(prefString, ",")
}

func StartDispatchScheduler(db *sqlite.SQLiteDB) {
    c := cron.New(cron.WithLocation(time.UTC))

    _, err := c.AddFunc("0 * * * *", func() {
        Dispatch(db)
    })

    if err != nil {
        log.Fatalf("Failed to schedule dispatcher: %v", err)
    }

    c.Start()
}