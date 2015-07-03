package bugsnag

import (
	"encoding/json"
	"testing"

	"github.com/facebookgo/stackerr"
)

func level1() error { return level2() }
func level2() error { return stackerr.Wrap(level3()) }
func level3() error { return stackerr.Newf("this is a test error") }

func TestStackerr(t *testing.T) {
	t.Parallel()

	err := stackerr.Wrap(level1())

	ev := ConvertStackerr(err)

	d, err := json.MarshalIndent(ev, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(d))

	c := NewClient("f8d383a38649f0a460b0c11cefc00661")

	if err := c.Notify([]Event{ev}); err != nil {
		t.Fatal(err)
	}
}
