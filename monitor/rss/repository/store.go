// Package store -
package store

// Storage -
type Storage interface {
	CheckExists(title string) (bool, error)
	CreateItem(title, content, link string) error
}
