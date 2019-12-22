package bugsnag // import "fknsrs.biz/p/bugsnag/bugsnagstackerr"

import (
	"fmt"

	"github.com/facebookgo/stack"
	"github.com/facebookgo/stackerr"

	"fknsrs.biz/p/bugsnag"
)

func init() {
	bugsnag.RegisterErrorConverter(&errorConverter{})
}

type errorConverter struct{}

func (errorConverter) CanConvertError(err error) bool {
	_, ok := err.(*stackerr.Error)
	return ok
}

func (errorConverter) ConvertError(err error) bugsnag.Event {
	ev := bugsnag.Event{}

	if serr, ok := err.(*stackerr.Error); ok {
		err = serr.Underlying()
		ev.Exceptions = convertMultiStack(serr.MultiStack())
	}

	ev.Exceptions = append([]bugsnag.Exception{{
		ErrorClass: fmt.Sprintf("%T", err),
		Message:    err.Error(),
	}}, ev.Exceptions...)

	return ev
}

func convertStack(s stack.Stack) []bugsnag.StackFrame {
	l := make([]bugsnag.StackFrame, len(s))

	for i, f := range s {
		l[i] = bugsnag.StackFrame{
			File:       f.File,
			LineNumber: f.Line,
			Method:     f.Name,
		}
	}

	return l
}

func convertMultiStack(m *stack.Multi) []bugsnag.Exception {
	l := make([]bugsnag.Exception, len(m.Stacks()))

	for i, s := range m.Stacks() {
		l[(len(l)-1)-i] = bugsnag.Exception{
			ErrorClass: "*stackerr.Error",
			Message:    fmt.Sprintf("relayed at line %d of %q", s[0].Line, s[0].File),
			Stacktrace: convertStack(s),
		}
	}

	return l
}
