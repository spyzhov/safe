package safe

import "fmt"

type tp struct{}

func (t *tp) String() string {
	if t == nil {
		return "<nil>"
	}
	return "string"
}

func ExampleIsNil() {
	var p *tp                    // p = (*T,nil)
	var s fmt.Stringer = p       // s = (*T,nil)
	fmt.Println(s == nil)        // false: (*T,nil) != (nil,nil)
	fmt.Println(s == (*tp)(nil)) // true : (*T,nil) == (*T,nil)
	fmt.Println(s.(*tp) == nil)  // true : nil == nil
	fmt.Println(IsNil(p))
	fmt.Println(IsNil(s))
	// Output:
	// false
	// true
	// true
	// true
	// true
}
