package closer

import "io"

// DiscardErrorHandler will discard errors returned by io.Closer if any
type DiscardErrorHandler struct {
}

// Close will discard errors returned by io.Closer if any
func (r DiscardErrorHandler) Close(c io.Closer) {
	_ = c.Close()
}
