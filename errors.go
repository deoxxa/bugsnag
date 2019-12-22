package bugsnag

import (
	"fmt"
)

type ErrorConverter interface {
	CanConvertError(err error) bool
	ConvertError(err error) Event
}

var errorConverters []ErrorConverter

func RegisterErrorConverter(c ErrorConverter) {
	errorConverters = append(errorConverters, c)
}

func convertError(err error) Event {
	for _, c := range errorConverters {
		if c.CanConvertError(err) {
			return c.ConvertError(err)
		}
	}

	return defaultConvertError(err)
}

func defaultConvertError(err error) Event {
	return Event{
		Exceptions: []Exception{
			Exception{
				ErrorClass: fmt.Sprintf("%T", err),
				Message:    err.Error(),
			},
		},
	}
}
