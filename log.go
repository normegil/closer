package closer

import (
	"io"
	"log"
)

// LogErrorHandler will log errors returned by io.Closer
type LogErrorHandler struct {
	// Fatal should be set to true if the logging should happen on fatal level
	Fatal bool
}

// Close will log errors returned by io.Closer
func (r LogErrorHandler) Close(c io.Closer) {
	if err := c.Close(); nil != err {
		if r.Fatal {
			log.Fatal(err)
		} else {
			log.Println(err)
		}
	}
}
