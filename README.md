# safe

[![Build Status](https://travis-ci.com/spyzhov/safe.svg?token=swf7VyTzTWuHyiC9QzT4&branch=master)](https://travis-ci.com/spyzhov/safe)
[![Go Report Card](https://goreportcard.com/badge/github.com/spyzhov/safe)](https://goreportcard.com/report/github.com/spyzhov/safe)
[![GoDoc](https://godoc.org/github.com/spyzhov/safe?status.svg)](https://godoc.org/github.com/spyzhov/safe)
[![Coverage Status](https://coveralls.io/repos/github/spyzhov/safe/badge.svg?branch=master)](https://coveralls.io/github/spyzhov/safe?branch=master)

Package safe is a minimal project with safe methods to write clean code.

## safe.Close

`Close` will close any `io.Closer` object if it's not `nil`, any error will be only logged.

```go
	resp, _ := http.DefaultClient.Get("http://example.com")
	defer safe.Close(resp.Body, "GET http://example.com")
	// ...
```

## safe.Must

`Must` will return the first argument as the result, if the error is `nil`, otherwise will call `panic`.

```go
	resp := safe.Must(http.DefaultClient.Get("http://example.com")).(http.*Response)
	// ...
```

## safe.IsNil

`IsNil` is a safe and fast method to check is current `interface{}` value is `nil`.

Reason is incident, that `nil != nil` in some cases.

Full description of the problem in article ["Go-tcha: When nil != nil"](https://dev.to/pauljlucas/go-tcha-when-nil--nil-hic).

```go
func Foo(str fmt.Stringer) {
    if !IsNil(str) {
        // ...
    }
}
```

# License

Licensed under the MIT license, see [LICENSE](LICENSE) for details.
