package say

import "fmt"

var Greeting = "Hello"

func Hello(name string) {
	fmt.Printf("%s %s\n", Greeting, name)
}
