package utils

import (
	"fmt"
	"strings"
)

// NormalizeString normalizes a string by trimming spaces and converting it to lowercase
func NormalizeString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

// PrintError prints an error message
func PrintError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
