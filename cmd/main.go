// Package main -
package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/mindfarm/republish-bot/monitor/rss"
	postgresstore "github.com/mindfarm/republish-bot/monitor/rss/repository/postgres"
	publish "github.com/mindfarm/republish-bot/platform"
	"github.com/mindfarm/republish-bot/platform/twitter"
)

func main() {
	// Collect twitter credentials
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPISecret := os.Getenv("TWITTER_API_SECRET")
	twitterAccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	// Create the twitter client
	tc, err := twitter.NewTwitterClient(twitterAPIKey, twitterAPISecret, twitterAccessToken, twitterAccessTokenSecret)
	if err != nil {
		log.Fatalf("Unable to create twitter client with error %v", err)
	}
	platforms := map[string]publish.Platform{}
	platforms["twitter"] = tc

	//TODO Create clients for other platforms (eg. Usenet, Slack, Direct Mail lists)

	// Connect to the github specific data store
	pgDB := os.Getenv("POSTGRES_DATABASE")
	db, err := postgresstore.NewPGStore(pgDB)
	if err != nil {
		log.Fatalf("Unable to create postgres connection with error %v", err)
	}

	// Retrieve content
	releaseURL := os.Getenv("RELEASE_URL")
	w, err := rss.NewWatched(db, releaseURL)
	if err != nil {
		log.Fatalf("Unable to create new watched instance with error %v", err)
	}

	for {
		releases, err := w.GetReleases()
		if err != nil {
			log.Printf("WARNING GetReleases() returned error %v", err)
		}
		for idx := range releases {
			for k := range platforms {
				var releaseInf string
				if k == "twitter" {
					// Drop the content
					releaseInf = strings.Join([]string{releases[idx][0], " is now available.", "\n\n", "Further information can be found at: ", releases[idx][2]}, "")
				} else {

					releaseInf = strings.Join(releases[idx], "\n\n")
				}
				if err := platforms[k].PublishContent(releaseInf); err != nil {
					log.Printf("WARNING PublishContent() returned error %v for %s", err, releaseInf)
				}
			}
		}
		// Sleep and check again in 30 seconds
		time.Sleep(time.Second * 30)
	}
}
