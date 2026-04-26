package utils

import (
	"math"
	"strings"
)

// Pagination calculates offset and total pages.
func Pagination(page, perPage int, total int64) (offset int, totalPage int) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	offset = (page - 1) * perPage
	totalPage = int(math.Ceil(float64(total) / float64(perPage)))
	return
}

// SanitizeString trims and lowercases a string.
func SanitizeString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// StringPtr returns a pointer to a string value.
func StringPtr(s string) *string {
	return &s
}

// UintPtr returns a pointer to a uint value.
func UintPtr(u uint) *uint {
	return &u
}
