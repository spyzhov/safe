package safe

import (
	"encoding/json"
	"fmt"
)

func ExampleWrap() {
	var value interface{}
	err := Wrap(json.Unmarshal([]byte(`[broken`), &value), "broken json")
	fmt.Println(err)
	// Output:
	// broken json: invalid character 'b' looking for beginning of value
}
