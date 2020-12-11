# Closer

Go library providing utilities to handle [io.Closer](https://pkg.go.dev/io#Closer) in defer statements.

## Install

`go get github.com/normegil/closer`

## Documentation

See [godoc](https://pkg.go.dev/github.com/normegil/closer) for documentation.

You can compose your own struct with [CloseErrorHandlerRequired](https://pkg.go.dev/github.com/normegil/closer#CloseErrorHandlerRequired) 
or use directly the [CloseErrorHandler](https://pkg.go.dev/github.com/normegil/closer#CloseErrorHandlerRequired) interface.

Available handlers:
* [DiscardErrorHandler](https://pkg.go.dev/github.com/normegil/closer#DiscardErrorHandler)
* [LogErrorHandler](https://pkg.go.dev/github.com/normegil/closer#LogErrorHandler)
* [FuncErrorHandler](https://pkg.go.dev/github.com/normegil/closer#FuncErrorHandler)