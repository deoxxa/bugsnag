package bugsnag

import (
	"fmt"

	"github.com/facebookgo/stack"
	"github.com/facebookgo/stackerr"
)

func convertError(err error, current stack.Stack) Event {
	ev := Event{}

	if serr, ok := err.(*stackerr.Error); ok {
		err = serr.Underlying()
		ev.Exceptions = convertMultiStack(serr.MultiStack())
	}

	ev.Exceptions = append([]Exception{Exception{
		ErrorClass: fmt.Sprintf("%T", err),
		Message:    err.Error(),
		Stacktrace: convertStack(current),
	}}, ev.Exceptions...)

	return ev
}

func convertStack(s stack.Stack) []StackFrame {
	l := make([]StackFrame, len(s))

	for i, f := range s {
		l[i] = StackFrame{
			File:       f.File,
			LineNumber: f.Line,
			Method:     f.Name,
		}
	}

	return l
}

func convertMultiStack(m *stack.Multi) []Exception {
	l := make([]Exception, len(m.Stacks()))

	for i, s := range m.Stacks() {
		l[(len(l)-1)-i] = Exception{
			ErrorClass: "*stackerr.Error",
			Message:    fmt.Sprintf("relayed at line %d of %q", s[0].Line, s[0].File),
			Stacktrace: convertStack(s),
		}
	}

	return l
}
