// Package monitor -
package monitor

// Updates -
type Updates interface {
	GetReleases() ([]map[string]string, error)
}
