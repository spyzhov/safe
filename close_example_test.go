package safe

import (
	"net/http"
)

func ExampleClose() {
	resp, err := http.DefaultClient.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer Close(resp.Body, "response body")
	// ...
}
