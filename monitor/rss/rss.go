// Package rss -
package rss

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	store "github.com/mindfarm/republish-bot/monitor/rss/repository"
	"github.com/mmcdole/gofeed"
)

// watched structure for handle parsing of RSS/Atom feeds
type watched struct {
	Feeds []struct {
		URL  string
		Feed *gofeed.Feed
	}
	Store store.Storage
}

// check that the concrete instance passed in as a store.Storage is not nil
func isNilFixed(i store.Storage) bool {
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
//nolint:golint
func NewWatched(store store.Storage, urls ...string) (*watched, error) {
	if isNilFixed(store) {
		return nil, fmt.Errorf("no store supplied cannot continue")
	}
	if len(urls) < 1 {
		return nil, fmt.Errorf("no urls supplied cannot continue")
	}
	r := &watched{Store: store}
	for _, u := range urls {
		r.Feeds = append(r.Feeds, struct {
			URL  string
			Feed *gofeed.Feed
		}{URL: u})
	}
	return r, nil
}

func (w *watched) GetReleases() ([][]string, error) {
	return w.GetUnseenReleases()
}

// GetUnseenReleases - get previously unseen releases
func (w *watched) GetUnseenReleases() ([][]string, error) {
	w.update()
	unseen := [][]string{}
	for i := range w.Feeds {
		for j := range w.Feeds[i].Feed.Items {
			title := w.Feeds[i].Feed.Items[j].Title
			seen, err := w.Store.CheckExists(title)
			if err != nil {
				log.Printf("unable to check %q with error %v", title, err)
				return [][]string{}, fmt.Errorf("unable to check store with error %w", err)
			}
			if !seen {
				content := w.Feeds[i].Feed.Items[j].Content
				link := w.Feeds[i].Feed.Items[j].Link
				if err := w.Store.CreateItem(title, content, link); err != nil {
					log.Printf("unable to create %q with error %v", title, err)
					return [][]string{}, fmt.Errorf("unable to create item with error %w", err)
				}
				unseen = append(unseen, []string{title, content, link})
			}
		}
	}
	return unseen, nil
}

// Update - fetch all items for all feeds
func (w *watched) update() {
	fp := gofeed.NewParser()
	for idx := range w.Feeds {
		url := w.Feeds[idx].URL
		feed, err := w.fetchURL(fp, url)
		if err != nil {
			log.Printf("error fetching url: %s, err: %v", url, err)
		}
		w.Feeds[idx].Feed = feed
	}
}

// FetchURL fetches the feed URL and also fakes the user-agent to be able
// to retrieve data from sites like reddit.
func (w *watched) fetchURL(fp *gofeed.Parser, url string) (feed *gofeed.Feed, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
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

	return fp.Parse(resp.Body)
}
