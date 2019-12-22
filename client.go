package bugsnag // import "fknsrs.biz/p/bugsnag"

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
	ProjectPackages []string
	notifications   int
}

// NewClient creates a new Client with the given API key
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:          apiKey,
		URL:             "https://notify.bugsnag.com/",
		NotifierName:    "fknsrs.biz/p/bugsnag",
		NotifierVersion: "0.0.0",
		NotifierURL:     "http://www.fknsrs.biz/p/bugsnag",
	}
}

// Notify sends a set of (maybe just one) events off to bugsnag
func (c *Client) Notify(events []Event) error {
	c.notifications = c.notifications + 1

	v := notifyRequest{
		APIKey: c.APIKey,
		Notifier: notifier{
			Name:    c.NotifierName,
			Version: c.NotifierVersion,
			URL:     c.NotifierURL,
		},
		Events: events,
	}

	d, err := json.Marshal(v)
	if err != nil {
		return err
	}

	r, err := http.Post(c.URL, "application/json", bytes.NewReader(d))
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return fmt.Errorf("invalid status code; expected 200 but got %d", r.StatusCode)
	}

	return nil
}

func (c *Client) Notifications() int {
	if c == nil {
		return 0
	}

	return c.notifications
}

func (c *Client) ReportPanic() {
	if e := recover(); e != nil {
		if err, ok := e.(error); ok {
			c.errors(3, err)
		}

		panic(e)
	}
}

func (c *Client) Errors(errs ...error) error {
	return c.errors(1, errs...)
}

func (c *Client) errors(skip int, errs ...error) error {
	l := make([]Event, len(errs))
	for i, e := range errs {
		// we add one to skip because we want to skip errors() as well
		l[i] = convertError(e)
	}

	return c.Notify(l)
}
