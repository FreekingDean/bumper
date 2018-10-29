package parser

import (
	"regexp"
	"strings"
)

var SPLIT_REGEXP = regexp.MustCompile(`[\.\[\]\ ]`)

func matchTitle(original, query string) bool {
	normalizedOriginal := normalizeString(original)
	normalizedQuery := normalizeString(query)
	i := 0
	for _, node := range normalizedOriginal {
		if node == normalizedQuery[i] {
			i += 1
		} else {
			i = 0
		}
		if i == len(normalizedQuery) {
			return true
		}
	}
	return false
}

func normalizeString(original string) []string {
	normalized := strings.ToLower(original)
	return SPLIT_REGEXP.Split(normalized, -1)
}
