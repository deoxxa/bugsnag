package bugsnag

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/facebookgo/stackerr"
)

var normalError = errors.New("this is a normal test error")
var panicError = errors.New("this is a panic test error")

func exampleParseContentsInner(ch chan<- error) {
	ch <- stackerr.Wrap(normalError)
}

func exampleParseContentsOuter(ch chan<- error) {
	c := make(chan error)
	go exampleParseContentsInner(c)
	ch <- stackerr.Wrap(<-c)
}

func exampleReadAndParse(ch chan<- error) {
	c := make(chan error)
	go exampleParseContentsOuter(c)
	ch <- stackerr.Wrap(<-c)
}

func TestStackerr(t *testing.T) {
	a := assert.New(t)

	t.Parallel()

	c := NewClient(apiKey)

	ch := make(chan error)

	go exampleReadAndParse(ch)

	a.NoError(c.Errors(<-ch, normalError))
	a.Equal(1, c.notifications)
}

func TestPanic(t *testing.T) {
	a := assert.New(t)

	c := NewClient(apiKey)

	a.Panics(func() {
		defer c.ReportPanic()
		panic(panicError)
	})
	a.Equal(1, c.notifications)

	a.NotPanics(func() {
		defer c.ReportPanic()
	})
	a.Equal(1, c.notifications)
}
