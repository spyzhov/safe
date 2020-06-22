package safe

import (
	"fmt"
	"net/http"
)

func ExampleClose() {
	Printf = func(format string, v ...interface{}) {
		fmt.Printf(format, v)
	}
	resp, err := http.DefaultClient.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer Close(resp.Body, "response body")
	// ...
}
