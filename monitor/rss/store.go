package rss

// Storage -
type Storage interface {
	CheckExists(feed, title string) (bool, error)
	CreateItem(feed, title, content, link string) error
}
