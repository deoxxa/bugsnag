package bugsnag

import (
	"os"
)

var apiKey = os.Getenv("BUGSNAG_KEY")
