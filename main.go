package main

import (
	"fmt"
)

func Names() (string, string) {
	return "Foo", "Bar"
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)

	// The underscore character (_) is used to ignore returned value
	n3, _ := Names()
	fmt.Println(n3)
}
