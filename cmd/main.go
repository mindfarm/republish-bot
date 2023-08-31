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
	irc "github.com/mindfarm/republish-bot/platform/irc-freenode"
)

func main() {
	// slog level management.
	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelInfo)

	// Collect IRC credentials
	ircUsername := os.Getenv("IRC_USER")
	ircPassword := os.Getenv("IRC_PASSWORD")

	// Create the IRC client
	tc, err := irc.NewFreenodeClient(ircUsername, ircPassword)
	if err != nil {
		log.Fatalf("Unable to create irc client with error %v", err)
	}

	platforms := map[string]republishbot.Platform{}
	platforms["irc"] = tc

	// Connect to the github specific data store
	pgDB := os.Getenv("POSTGRES_DATABASE")

	db, err := postgresstore.NewPGStore(pgDB)
	if err != nil {
		log.Fatalf("Unable to create postgres connection with error %v", err)
	}

	// Retrieve content
	watched, err := rss.NewWatched(db, releaseURLs...)
	if err != nil {
		log.Fatalf("Unable to create new watched instance with error %v", err)
	}

	for {
		//nolint:gomnd
		// buffer size is arbitrarily set to '5'
		releases := make(chan map[string]string, 5)

		slog.Debug("Start of infinite loop")

		go watched.GetReleases(releases)

		for item := range releases {
			slog.Debug("Publishing ", "value", item["title"])

			for k := range platforms {
				if err := platforms[k].PublishContent(item); err != nil {
					log.Printf("WARNING PublishContent() returned error %v for %s", err, item)
				}
			}
		}
		//nolint:gomnd
		// Sleep and check again in 30 seconds
		time.Sleep(time.Second * 30)
	}
}
