// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// Given a name, return string with message "One for {name}, one for me"
// no name, return "One for you, one for me"
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %v, one for me.", name)
	} else {
		return "One for you, one for me."
	}
}
