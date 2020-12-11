package closer

import "io"

// FuncErrorHandler register a function to call when closing an io.Closer
type FuncErrorHandler struct {
	fn func(io.Closer)
}

// Close call the registered function with the passed io.Closer
func (f FuncErrorHandler) Close(closer io.Closer) {
	f.fn(closer)
}

// CloseErrorHandlerFunc is a convenience method to create a FuncErrorHandler
func CloseErrorHandlerFunc(fn func(closer io.Closer)) *FuncErrorHandler {
	return &FuncErrorHandler{
		fn: fn,
	}
}
