package main

import (
	"strings"

	"github.com/spf13/viper"
)

// filterFunc returns a function that filters templates based on the specified filter type.
func filterFunc(templates []string) func(input string, index int) bool {
	switch viper.GetString("filter") {
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
