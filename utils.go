package main

import "strings"

func FormatDescription(description string) string {
	description = strings.ReplaceAll(description, "\n", "")
	if len(description) > 159 {
		return description[:159] + "..."
	}
	return description
}
