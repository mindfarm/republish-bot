// Package main -
package main

import (
	"log"
	"os"
	"time"

	"github.com/mindfarm/republish-bot/monitor/rss"
	postgresstore "github.com/mindfarm/republish-bot/monitor/rss/repository/postgres"

	republishbot "github.com/mindfarm/republish-bot"
	"github.com/mindfarm/republish-bot/platform/reddit"
	"github.com/mindfarm/republish-bot/platform/twitter"
)

func main() {
	// Collect Twitter credentials
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPISecret := os.Getenv("TWITTER_API_SECRET")
	twitterAccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

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
	db, err := postgresstore.NewPGStore(pgDB)
	if err != nil {
		log.Fatalf("Unable to create postgres connection with error %v", err)
	}

	// Retrieve content
	releaseURLs := []string{"https://github.com/golang/tools/releases.atom", "https://github.com/golang/go/releases.atom"}
	w, err := rss.NewWatched(db, releaseURLs...)
	if err != nil {
		log.Fatalf("Unable to create new watched instance with error %v", err)
	}

	republishbot.Republish(platforms, w)
}
