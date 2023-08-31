// Package rss -
package rss

import (
	"errors"
	"log/slog"
	"reflect"

	"github.com/mmcdole/gofeed"
)

var ErrNoStore = errors.New("no store supplied cannot continue")
var ErrNoURL = errors.New("no urls supplied cannot continue")

// watched structure for handle parsing of RSS/Atom feeds.
type watched struct {
	Feeds []feedData
	Store Storage
}

type feedData struct {
	URL  string
	Feed *gofeed.Feed
}

// check that the concrete instance passed in as a Storage is not nil.
func isNilFixed(instance Storage) bool {
	if instance == nil {
		return true
	}

	switch reflect.TypeOf(instance).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(instance).IsNil()
	}

	return false
}

// NewWatched - create a new watched instance.
func NewWatched(store Storage, urls ...string) (*watched, error) {
	if isNilFixed(store) {
		return nil, ErrNoStore
	}

	if len(urls) < 1 {
		return nil, ErrNoURL
	}

	newWatched := &watched{Store: store}

	for _, u := range urls {
		newWatched.Feeds = append(newWatched.Feeds, struct {
			URL  string
			Feed *gofeed.Feed
		}{URL: u})
	}

	return newWatched, nil
}

func (w *watched) GetReleases(c chan map[string]string) {
	w.GetUnseenReleases(c)
}

// GetUnseenReleases - get previously unseen releases.
func (w *watched) GetUnseenReleases(unseen chan map[string]string) {
	slog.Debug("Getting unseen")

	//nolint:gomnd
	// The buffer length is arbirtrary.
	feeds := make(chan feedData, 10)

	go w.update(feeds)

	for feed := range feeds {
		if feed.Feed == nil {
			continue
		}

		slog.Debug("Feed", "value", feed.URL)
		slog.Debug("Feed Item", "value", feed.Feed.Items)

		for idx := range feed.Feed.Items {
			title := feed.Feed.Items[idx].Title

			seen, err := w.Store.CheckExists(feed.URL, title)
			if err != nil {
				slog.Warn("unable to check %q with error %v", title, err)

				continue
			}

			if !seen {
				content := feed.Feed.Items[idx].Content
				link := feed.Feed.Items[idx].Link

				if err := w.Store.CreateItem(feed.URL, title, content, link); err != nil {
					slog.Warn("unable to create %q with error %v", title, err)

					continue
				}

				unseen <- map[string]string{"title": title, "content": content, "link": link}
			}
		}
	}

	slog.Debug("Unseen count", "value", len(unseen))
	close(unseen)
}

// Update - fetch all items for all feeds.
func (w *watched) update(feeds chan feedData) {
	slog.Debug("Fetching items for feeds")

	fp := gofeed.NewParser()

	for idx := range w.Feeds {
		slog.Debug("URL", "value", w.Feeds[idx].URL)

		url := w.Feeds[idx].URL

		feed, err := fp.ParseURL(url)
		if err != nil {
			continue
		}

		feeds <- feedData{URL: url, Feed: feed}
	}

	close(feeds)
}

//nolint:lll
// FetchURL fetches the feed URL and also fakes the user-agent to be able
// to retrieve data from sites like reddit.
// func (w *watched) fetchURL(fp *gofeed.Parser, url string) (feed *gofeed.Feed, err error) {
//	slog.Debug("fetching URL", "value", url)
//
//	client := &http.Client{}
//
//	slog.Debug("Preparing request")
//
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
//
//	slog.Debug("Making Request")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if resp != nil {
//		defer func() {
//			ce := resp.Body.Close()
//			if ce != nil {
//				err = ce
//			}
//		}()
//	}
//
//	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
//		return nil, fmt.Errorf("failed to get url %v, %v", resp.StatusCode, resp.Status)
//	}
//
//	slog.Debug("returning body")
//	return fp.Parse(resp.Body)
//}
