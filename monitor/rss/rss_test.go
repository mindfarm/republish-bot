package rss

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeStorage struct{}

var ceBool bool
var ceErr, ciErr error

func (f *fakeStorage) CheckExists(title string) (bool, error) {
	return ceBool, ceErr
}
func (f *fakeStorage) CreateItem(title, content, link string) error {

	return ciErr
}

func TestNewRSS(t *testing.T) {

	testcases := map[string]struct {
		urls   []string
		output *watched
		err    error
		store  *fakeStorage
	}{
		"New RSS": {
			urls:   []string{"http://example.com"},
			output: &watched{},
			store:  &fakeStorage{},
		},
		"No storage": {
			urls:   []string{"http://example.com"},
			output: &watched{},
			err:    fmt.Errorf("no store supplied cannot continue"),
		},
		"No urls": {
			output: &watched{},
			store:  &fakeStorage{},
			err:    fmt.Errorf("no urls supplied cannot continue"),
		},
		"Multiple URLS": {
			urls:   []string{"http://example.com", "https://example2.com"},
			output: &watched{},
			store:  &fakeStorage{},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			fakeStore := tc.store
			output, err := NewWatched(fakeStore, tc.urls...)
			if tc.err == nil {
				assert.Nilf(t, err, "newRSS generated an unexpected error %v", err)
				assert.IsType(t, tc.output, output, "Not an RSS instance")
				assert.Equal(t, len(tc.urls), len(output.Feeds), "output has incorrect number of feeds")
				for idx := range tc.urls {
					assert.Equalf(t, tc.urls[idx], output.Feeds[idx].URL, "feed %d has different url", idx)
				}
			} else {
				assert.NotNil(t, err, "newRSS didn't generate an expected error ")
				assert.EqualError(t, err, tc.err.Error(), "newRSS generated a different error", err)
			}
		})
	}
}
