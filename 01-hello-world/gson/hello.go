package hello

import "fmt"

func Hello(name string) string {
	if len(name) == 0 {
		return "Hello, world"
	}
	return fmt.Sprintf("Hello, %s", name)
}
