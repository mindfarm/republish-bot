package republishbot

// Platform - provide an interface that needs to be satisfied in order to
// publish to a given platform.
type Platform interface {
	PublishContent(map[string]string) error
}
