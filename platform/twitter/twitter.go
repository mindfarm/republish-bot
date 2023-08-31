// Package twitter -
package twitter

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type twitterClient struct {
	*twitter.Client
}

var ErrMissingCfg = errors.New("missing consumerKey, consumerSecret, accessToken, or accessSecret cannot continue")

func NewTwitterClient(consumerKey, consumerSecret, accessToken, accessSecret string) (*twitterClient, error) {
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return nil, ErrMissingCfg
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return &twitterClient{twitter.NewClient(httpClient)}, nil
}

// Max length of a single tweet.
const tweetLen = 280

func (t *twitterClient) chunkContent(content string) []string {
	// break the content up into chunks that are tweetLen or less characters
	// if the tweetLen character isn't whitespace, then look at the tweetLen -1,
	// and so on until a space is found. The next chunk then starts at that
	// position.
	chunks := []string{}
	bottom := 0
	runeContent := []rune(content)
	top := tweetLen

	l := len(runeContent)
	for bottom < l {
		if top >= l {
			top = l
			chunks = append(chunks, string(runeContent[bottom:top]))

			break
		}

		chunk := runeContent[bottom:top]
		// adjust the top value down to the first whitespace char
		for i := len(chunk) - 1; i > 0; i-- {
			r := chunk[i]
			if unicode.IsSpace(r) {
				top = bottom + i

				break
			}
		}

		chunks = append(chunks, string(runeContent[bottom:top]))
		bottom = top + 1
		top += tweetLen
	}

	return chunks
}

func (t *twitterClient) PublishContent(content map[string]string) error {
	params := &twitter.StatusUpdateParams{}

	title := content["title"]
	// Twitter should not publish pre-releases (only possible where pre has been
	// included in the title of the release, as is the case with go-pls)
	if strings.ContainsAny(title, "pre") {
		log.Printf("Twitter: Ignoring a pre-release %s", content["title"])

		return nil
	}

	// Go project has a weird title structure
	// [release-branch.go1.15] go1.15.2
	tmp := strings.Split(title, "]")
	if len(tmp) > 1 {
		title = strings.TrimSpace(tmp[1])
	}

	//nolint:lll
	c := strings.Join([]string{title, " is now available.", "\n\n", "Further information can be found at: ", content["link"]}, "")
	for _, snippet := range t.chunkContent(c) {
		tweet, resp, err := t.Statuses.Update(snippet, params)
		if resp.StatusCode != http.StatusOK {
			log.Printf("http return status was %d, with %s", resp.StatusCode, resp.Status)
			log.Printf("accompanied with error: %v", err)
			err = fmt.Errorf("bad status code returned with error %w", err)

			return err
		}

		params.InReplyToStatusID = tweet.ID
	}

	return nil
}
