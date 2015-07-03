package bugsnag

import (
	"testing"
)

func TestClient(t *testing.T) {
	t.Parallel()

	c := NewClient("f8d383a38649f0a460b0c11cefc00661")

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
