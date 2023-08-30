// Package main -
package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/mindfarm/republish-bot/monitor/rss"
	postgresstore "github.com/mindfarm/republish-bot/monitor/rss/repository/postgres"

	republishbot "github.com/mindfarm/republish-bot"
)

func main() {
	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelInfo)

	// Create the Twitter client
	tc, err := twitter.NewTwitterClient(twitterAPIKey, twitterAPISecret, twitterAccessToken, twitterAccessTokenSecret)
	if err != nil {
		log.Fatalf("Unable to create twitter client with error %v", err)
	}

	platforms := map[string]republishbot.Platform{}
	platforms["twitter"] = tc

	// Collect Reddit credentials
	redditClientID := os.Getenv("REDDIT_CLIENT_ID")
	redditClientSecret := os.Getenv("REDDIT_CLIENT_SECRET")
	redditUsername := os.Getenv("REDDIT_USERNAME")
	redditPassword := os.Getenv("REDDIT_PASSWORD")

	// Create the Reddit client
	rc, err := reddit.NewRedditClient(redditUsername, redditPassword, redditClientID, redditClientSecret)
	if err != nil {
		log.Fatalf("Unable to create reddit client with error %v", err)
	}

	platforms["reddit"] = rc

	//TODO Create clients for other platforms (eg. Usenet, Slack, Direct Mail lists)

	// Connect to the github specific data store
	pgDB := os.Getenv("POSTGRES_DATABASE")
	slog.Debug("PG", pgDB)
	db, err := postgresstore.NewPGStore(pgDB)
	if err != nil {
		log.Fatalf("Unable to create postgres connection with error %v", err)
	}

	// Retrieve content
	w, err := rss.NewWatched(db, releaseURLs...)
	if err != nil {
		log.Fatalf("Unable to create new watched instance with error %v", err)
	}

	for {
		slog.Debug("Start of infinite loop")
		releases, err := w.GetReleases()
		if err != nil {
			log.Printf("WARNING GetReleases() returned error %v", err)
		}
		for idx := range releases {
			slog.Debug("Publishing ", "value", item["title"])
			for k := range platforms {
				if err := platforms[k].PublishContent(releases[idx]); err != nil {
					log.Printf("WARNING PublishContent() returned error %v for %s", err, releases[idx])
				}
			}
		}
		// Sleep and check again in 30 seconds
		time.Sleep(time.Second * 30)
	}
}
