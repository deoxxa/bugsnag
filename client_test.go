package bugsnag

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	t.Parallel()

	c := NewClient(os.Getenv("BUGSNAG_KEY"))

	c.Notify([]Event{
		{
			Exceptions: []Exception{
				{
					ErrorClass: "what",
					Message:    "this is bad",
					Stacktrace: []StackFrame{
						{
							File:         "test.go",
							LineNumber:   5,
							ColumnNumber: 5,
							Method:       "testing",
							InProject:    true,
						},
					},
				},
			},
			App: &App{
				Commit: "be0c8adea19ce0f6cbcf90934c1a181a7bd98caa",
			},
		},
	})
}
