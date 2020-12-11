// Closer will help you manage errors in io.Closer used in defer consistently
package closer

import "io"

// CloseErrorHandler defines the interface to use to handle io.Closer errors
type CloseErrorHandler interface {
	Close(c io.Closer)
}

// CloseErrorHandlerRequired can be used to compose your own struct to use a CloseErrorHandler, or fallback on DiscardErrorHandler if required
type CloseErrorHandlerRequired struct {
	Handler CloseErrorHandler
}

// CloseHandler will return your io.Closer handler, or a DiscardErrorHandler if no handler was defined
func (c CloseErrorHandlerRequired) CloseHandler() CloseErrorHandler {
	if nil == c.Handler {
		return DiscardErrorHandler{}
	}
	return c.Handler
}
