package republishbot

// Updates -
type Updates interface {
	GetReleases() ([]map[string]string, error)
}
