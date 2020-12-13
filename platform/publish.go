// Package publish - provide an interface that needs to be satisfied in order to
// publish to a given platform.
package publish

// Platform -
type Platform interface {
	PublishContent(string) error
}
