package closer

import "io"

// PanicErrorHandler will panic on errors returned by io.Closer if any
type PanicErrorHandler struct {
}

// Close will panic on errors returned by io.Closer if any
func (r PanicErrorHandler) Close(c io.Closer) {
	if err := c.Close(); nil != err {
		panic(err)
	}
}
