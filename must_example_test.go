package safe

import (
	"fmt"
	"os"
)

func ExampleMust() {
	defer func() {
		fmt.Println(recover())
	}()

	fmt.Println(Must(fmt.Println("Example")))

	file := Must(os.Open("doc.go")).(*os.File)
	Close(file, "doc.go")

	Must(os.Open("***NOT_EXISTS***"))
	// Output:
	// Example
	// 8
	// open ***NOT_EXISTS***: no such file or directory
}
