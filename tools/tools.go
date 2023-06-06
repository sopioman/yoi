package tools

import (
	"strings"
)

// Gets the text within a specified set of characters
func SplitBetween(str, bef, aft string) string {
	sa := strings.SplitN(str, bef, 2)
	if len(sa) == 1 {
		return ""
	}
	sa = strings.SplitN(sa[1], aft, 2)
	if len(sa) == 1 {
		return ""
	}
	return sa[0]
}

// Get substring after a string
func AfterChar(value string, aft string) string {
	pos := strings.LastIndex(value, aft)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(aft)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
