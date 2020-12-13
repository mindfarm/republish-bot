// Package monitor -
package monitor

// Updates -
type Updates interface {
	GetReleases() (proverb, translation, explanation string, err error)
	// GetLatestCommit() (placename, translation, explanation string, err error)
}
