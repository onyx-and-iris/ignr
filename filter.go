package main

import (
	"strings"
)

// filterFunc returns a function that filters templates based on the specified filter type.
func filterFunc(templates []string, filterType string) func(input string, index int) bool {
	switch filterType {
	case "contains":
		return func(input string, index int) bool {
			return strings.Contains(strings.ToLower(templates[index]), strings.ToLower(input))
		}
	default:
		return func(input string, index int) bool {
			return strings.HasPrefix(strings.ToLower(templates[index]), strings.ToLower(input))
		}
	}
}
