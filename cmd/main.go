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
