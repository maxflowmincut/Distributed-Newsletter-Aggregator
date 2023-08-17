package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	
	"newsletter-aggregator/src/cmd"
	"newsletter-aggregator/src/config"
	"newsletter-aggregator/src/db/sqlite"
	"newsletter-aggregator/src/rss"
	"newsletter-aggregator/src/curation/go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	appConfig := config.LoadConfig()

	db, err := sqlite.NewSQLiteDB(appConfig.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

    rss.StartRSSFetcherScheduler(db, appConfig)
    curation.StartDispatchScheduler(db)
	cmd.Execute(db)

	// Wait for termination signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Program terminated!")
}