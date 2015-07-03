package bugsnag

// Event is an event that goes to bugsnag.
//
// Note: the `PayloadVersion` field is a special type that always encodes to
// `"2"` in JSON, so it doesn't need to be set for the request to work. In fact,
// whatever you set it to will be ignored. It just needs to exist for bugsnag.
type Event struct {
	PayloadVersion payloadVersion                    `json:"payloadVersion"`
	Exceptions     []Exception                       `json:"exceptions,omitempty"`
	Threads        []Thread                          `json:"threads,omitempty"`
	Context        string                            `json:"context,omitempty"`
	GroupingHash   string                            `json:"groupingHash,omitempty"`
	Severity       string                            `json:"severity,omitempty"`
	User           *User                             `json:"user,omitempty"`
	App            *App                              `json:"app,omitempty"`
	Device         *Device                           `json:"device,omitempty"`
	MetaData       map[string]map[string]interface{} `json:"metaData,omitempty"`
}

// Exception represents a single exception.
type Exception struct {
	ErrorClass string       `json:"errorClass"`
	Message    string       `json:"message"`
	Stacktrace []StackFrame `json:"stacktrace"`
}

// Thread represents a (background?) thread. Threads should be in an order that
// makes sense for your application.
type Thread struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Stacktrace []StackFrame `json:"stacktrace"`
}

// StackFrame represents a single stack frame, commonly shown as one line in
// an exception.
type StackFrame struct {
	File         string            `json:"file"`
	LineNumber   int               `json:"lineNumber"`
	ColumnNumber int               `json:"columnNumber"`
	Method       string            `json:"method"`
	InProject    bool              `json:"inProject"`
	Code         map[string]string `json:"code,omitempty"`
}

// User represents information about the user affected by the crash or error.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// App represents some information about the application that crashed or
// encountered an error.
//
// Note that the `Commit` field is non-standard - bugsnag does _not_ actively
// support this field, so don't annoy them if it doesn't do what you want.
type App struct {
	Version      string `json:"version"`
	ReleaseStage string `json:"releaseStage"`
	Commit       string `json:"commit"`
}

// Device represents the device that was accessing the application at the time
// it crashed, or an error occurred.
type Device struct {
	OSVersion string `json:"osVersion"`
	Hostname  string `json:"hostname"`
}

type notifier struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	URL     string `json:"url"`
}

type notifyRequest struct {
	APIKey   string   `json:"apiKey"`
	Notifier notifier `json:"notifier"`
	Events   []Event  `json:"events"`
}

type payloadVersion string

func (p payloadVersion) MarshalJSON() ([]byte, error) {
	return []byte("\"2\""), nil
}

func (p *payloadVersion) UnmarshalJSON(b []byte) error {
	*p = "2"

	return nil
}

func (p payloadVersion) String() string {
	return "2"
}
