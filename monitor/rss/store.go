package rss

// Storage - API for storing rss data.
type Storage interface {
	CheckExists(feed, title string) (bool, error)
	CreateItem(feed, title, content, link string) error
}
