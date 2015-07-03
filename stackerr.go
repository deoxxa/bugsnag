package bugsnag

import (
	"fmt"

	"github.com/facebookgo/stack"
)

// ConvertStackerr converts a stackerr.Error (or anything that looks like one)
// into an Event that can be submitted to bugsnag. It iterates through the
// stacks kept in the stackerr object and submits each of them as separate
// exceptions in order to populate the "caused" area of the bugsnag UI.
func ConvertStackerr(err error) Event {
	serr, ok := err.(LikeStackerr)
	if !ok {
		return Event{}
	}

	ev := Event{
		Exceptions: convertMultiStack(serr.MultiStack()),
	}

	ev.Exceptions = append(ev.Exceptions, Exception{
		ErrorClass: fmt.Sprintf("%T", serr.Underlying()),
		Message:    serr.Underlying().Error(),
		Stacktrace: []StackFrame{
			{
				File:   "dummy.go",
				Method: "original_error_site",
			},
		},
	})

	return ev
}

// LikeStackerr represents the behaviour that we rely on from a stackerr.Error.
// It's defined as an interface so that forks of stackerr will (should?) still
// work.
type LikeStackerr interface {
	MultiStack() *stack.Multi
	Underlying() error
}

func convertMultiStack(m *stack.Multi) []Exception {
	var l []Exception

	for _, s := range m.Stacks() {
		ex := Exception{
			ErrorClass: "*stackerr.Error",
			Message:    fmt.Sprintf("relayed at line %d of %q", s[0].Line, s[0].File),
		}

		for _, f := range s {
			ex.Stacktrace = append(ex.Stacktrace, StackFrame{
				File:       f.File,
				LineNumber: f.Line,
				Method:     f.Name,
			})
		}

		l = append([]Exception{ex}, l...)
	}

	return l
}
