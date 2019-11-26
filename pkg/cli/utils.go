package cli

import (
	"fmt"
	"strings"
)

func ask4confirm(query string) bool {
	var s string

	fmt.Printf("%s (y/N): ", query)
	_, err := fmt.Scanln(&s)
	if err != nil {
		return false
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}
