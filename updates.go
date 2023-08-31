package republishbot

type Updates interface {
	GetReleases() ([]map[string]string, error)
}
