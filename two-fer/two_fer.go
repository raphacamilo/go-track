// package for share cookie with friends
package twofer

import "fmt"

// ShareWith returns a string saying who you will share the cookie with
func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("One for %s, one for me.", name)
	}

	return "One for you, one for me."
}
