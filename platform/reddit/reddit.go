// Package reddit -
package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type client struct {
	ClientID                  string
	clientSecret              string
	AccessToken               string
	AccessTokenExpirationTime time.Time
	Username                  string
	password                  string

	httpClient *http.Client
	UserAgent  string

	lastRequestTime time.Time
}

var ErrCreateRedditClient = errors.New("unable to create reddit client")

func NewRedditClient(username, password, clientID, clientSecret string) (*client, error) {
	if username == "" {
		return nil, fmt.Errorf("%w no username supplied", ErrCreateRedditClient)
	}

	if password == "" {
		return nil, fmt.Errorf("%w no password supplied", ErrCreateRedditClient)
	}

	if clientID == "" {
		return nil, fmt.Errorf("%w no clientID supplied", ErrCreateRedditClient)
	}

	if clientSecret == "" {
		return nil, fmt.Errorf("%w no clientSecret supplied", ErrCreateRedditClient)
	}

	redditClient := &client{
		ClientID:     clientID,
		clientSecret: clientSecret,
		UserAgent:    "golang:mindfarm_bot:v0.0.1 (by /u/announce_bot)",
		httpClient:   &http.Client{},
		Username:     username,
		password:     password,
	}

	return redditClient, nil
}

func (c *client) doRequest(request *http.Request) (*http.Response, error) {
	request.Header.Set("User-Agent", c.UserAgent)

	//nolint:gomnd
	waitLength := time.Duration(2) * time.Second
	elapsedTime := time.Since(c.lastRequestTime)

	if elapsedTime < waitLength {
		time.Sleep(waitLength - elapsedTime)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	c.lastRequestTime = time.Now()

	return response, nil
}

func (c *client) authorize() error {
	// See
	// https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example.
	if c.AccessToken != "" && time.Now().Before(c.AccessTokenExpirationTime) {
		return nil
	}

	form := url.Values{
		"grant_type": {"password"},
		"username":   {c.Username},
		"password":   {c.password},
	}

	endpointURL := "https://www.reddit.com/api/v1/access_token"

	req, err := http.NewRequest("POST", endpointURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.ClientID, c.clientSecret)

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}

	defer func() {
		if bErr := resp.Body.Close(); bErr != nil {
			slog.Warn("unable to close response body", "value", bErr)
		}
	}()

	type TokenStruct struct {
		Scope       string
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	tokenStruct := TokenStruct{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&tokenStruct)
	if err != nil {
		return err
	}

	c.AccessToken = tokenStruct.AccessToken
	c.AccessTokenExpirationTime = time.Now().Add(time.Duration(tokenStruct.ExpiresIn) * time.Second)

	return nil
}

func (c *client) Post(resourceURL string, values url.Values) (*http.Response, error) {
	err := c.authorize()
	if err != nil {
		return nil, err
	}

	endpointURL := "https://oauth.reddit.com" + resourceURL

	req, err := http.NewRequest("POST", endpointURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.doRequest(req)
}

func (c *client) PublishContent(content map[string]string) error {
	title := content["title"]
	// Reddit should not publish pre-releases (only possible where pre has been
	// included in the title of the release, as is the case with go-pls)
	if strings.ContainsAny(title, "pre") {
		slog.Warn("REDDIT: Ignoring a pre-release %s", "value", content["title"])

		return nil
	}
	// Go project has a weird title structure
	// [release-branch.go1.15] go1.15.2
	tmp := strings.Split(title, "]")
	if len(tmp) > 1 {
		title = strings.TrimSpace(tmp[1])
	}

	resourceURL := "/api/submit"
	values := url.Values{
		"kind":  {"self"},
		"sr":    {"golang"},
		"title": {title},
		"text": {fmt.Sprintf(`
Further information can be found at %s

`, content["link"])}}

	_, err := c.Post(resourceURL, values)

	return err
}

/* Unused but worth keeping for example reasons
func (c *client) Get(resourceURL string) (*http.Response, error) {
	err := c.authorize()
	if err != nil {
		return nil, err
	}
	fmt.Printf("C: %#v\n", c)

	endpointURL := "https://oauth.reddit.com/" + resourceURL
	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "bearer "+c.AccessToken)

	return c.doRequest(req)
}
*/

/* Unused but worth keeping for example purposes
func (c *client) SubmitComment(parentID string, text string) error {
	values := url.Values{
		"api_type": {"json"},
		"text":     {text},
		"thing_id": {parentID},
	}

	response, err := c.Post("/api/comment.json", values)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//  TODO: Parse response and actually make sure submission was successful

	return nil
}
*/
