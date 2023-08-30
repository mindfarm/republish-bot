// Package rss -
package rss

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"reflect"
	"time"

	"github.com/mmcdole/gofeed"
)

// watched structure for handle parsing of RSS/Atom feeds
type watched struct {
	Feeds []struct {
		URL  string
		Feed *gofeed.Feed
	}
	Store Storage
}

// check that the concrete instance passed in as a Storage is not nil
func isNilFixed(i Storage) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// NewWatched - create a new watched instance
//
//nolint:golint,revive
func NewWatched(store Storage, urls ...string) (*watched, error) {
	if isNilFixed(store) {
		return nil, fmt.Errorf("no store supplied cannot continue")
	}
	if len(urls) < 1 {
		return nil, fmt.Errorf("no urls supplied cannot continue")
	}
	r := &watched{Store: store}
	for _, u := range urls {
		// slog.Debug("Adding URL: ", u)
		r.Feeds = append(r.Feeds, struct {
			URL  string
			Feed *gofeed.Feed
		}{URL: u})
	}
	return r, nil
}

func (w *watched) GetReleases() ([]map[string]string, error) {
	return w.GetUnseenReleases()
}

// GetUnseenReleases - get previously unseen releases
func (w *watched) GetUnseenReleases() ([]map[string]string, error) {
	// Keep trying for up to 30 seconds (In case there is an upstream site issue)
	slog.Debug("Getting unseen")
	maxWait := 6
	for maxWait > 0 {
		err := w.update()
		if err != nil {
			time.Sleep(time.Second * 5)
			maxWait--
		} else {
			break
		}
	}
	unseen := []map[string]string{}
	for i := range w.Feeds {
		if w.Feeds[i].Feed == nil {
			continue
		}
		slog.Debug("Feed", "value", w.Feeds[i].URL)
		slog.Debug("Feed Item", "value", w.Feeds[i].Feed.Items)
		for j := range w.Feeds[i].Feed.Items {
			title := w.Feeds[i].Feed.Items[j].Title
			seen, err := w.Store.CheckExists(title)
			if err != nil {
				log.Printf("unable to check %q with error %v", title, err)
				return nil, fmt.Errorf("unable to check store with error %w", err)
			}
			if !seen {
				content := w.Feeds[i].Feed.Items[j].Content
				link := w.Feeds[i].Feed.Items[j].Link
				if err := w.Store.CreateItem(title, content, link); err != nil {
					log.Printf("unable to create %q with error %v", title, err)
					return nil, fmt.Errorf("unable to create item with error %w", err)
				}
				unseen = append(unseen, map[string]string{"title": title, "content": content, "link": link})
			}
		}
	}
	slog.Debug("Unseen count", "value", len(unseen))
	return unseen, nil
}

// Update - fetch all items for all feeds
func (w *watched) update() error {
	slog.Debug("Fetching items for feeds")
	fp := gofeed.NewParser()

	// slog.Debug("Feeds: ", w.Feeds)
	for idx := range w.Feeds {
		slog.Debug("URL", "value", w.Feeds[idx].URL)
		url := w.Feeds[idx].URL
		feed, err := fp.ParseURL(url)
		if err != nil {
			// slog.Warn("URL Fetch Error", err.Error())
			continue
		}
		// feed, err := w.fetchURL(fp, url)
		// if err != nil {
		// 	slog.Debug("URL Fetch Error", err.Error())
		// 	return fmt.Errorf("error fetching url: %s, err: %v", url, err)
		// }
		// slog.Debug("Adding feed", feed)
		w.Feeds[idx].Feed = feed
	}

	return nil
}

// FetchURL fetches the feed URL and also fakes the user-agent to be able
// to retrieve data from sites like reddit.
func (w *watched) fetchURL(fp *gofeed.Parser, url string) (feed *gofeed.Feed, err error) {
	slog.Debug("fetching URL", "value", url)
	client := &http.Client{}

	slog.Debug("Preparing request")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	slog.Debug("Making Request")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer func() {
			ce := resp.Body.Close()
			if ce != nil {
				err = ce
			}
		}()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to get url %v, %v", resp.StatusCode, resp.Status)
	}
	slog.Debug("returning body")
	return fp.Parse(resp.Body)
}
