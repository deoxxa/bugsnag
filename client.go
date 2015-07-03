package bugsnag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client holds information to include in requests to bugsnag.
type Client struct {
	APIKey          string
	URL             string
	NotifierName    string
	NotifierVersion string
	NotifierURL     string
}

// NewClient creates a new Client with the given API key
func NewClient(apiKey string) Client {
	return Client{
		APIKey:          apiKey,
		URL:             "https://notify.bugsnag.com/",
		NotifierName:    "fknsrs.biz/p/bugsnag",
		NotifierVersion: "0.0.0",
		NotifierURL:     "http://www.fknsrs.biz/p/bugsnag",
	}
}

// Notify sends a set of (maybe just one) events off to bugsnag
func (c Client) Notify(events []Event) error {
	n := notifier{
		Name:    c.NotifierName,
		Version: c.NotifierVersion,
		URL:     c.NotifierURL,
	}

	v := notifyRequest{
		APIKey:   c.APIKey,
		Notifier: n,
		Events:   events,
	}

	d, err := json.Marshal(v)
	if err != nil {
		return err
	}

	r, err := http.Post(c.URL, "application/json", bytes.NewReader(d))
	if r.StatusCode != 200 {
		return fmt.Errorf("invalid status code; expected 200 but got %d", r.StatusCode)
	}

	return nil
}
