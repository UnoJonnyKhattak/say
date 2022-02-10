// Package say provides a function to say a phrase to someone
package say

import "fmt"

// Hello returns a string saying "Hello" to the name of
// the person that has been provided.
func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func Bye(name string) string {
	return fmt.Sprintf("Bye %s", name)
}
