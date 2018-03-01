// Provide convert functions (string => snake_case, chain-case, camelCase, PascalCase).
package casee

import (
	"github.com/fatih/camelcase"
	"strings"
	"unicode"
)

// Convert argument to snake_case style string.
// If argument is empty, return itself.
func ToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "_")
}

// If argument is snake_case style string, return true.
func IsSnakeCase(s string) bool {
	if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	} else {
		return isMadeByLowerAndDigit(s)
	}
}

// Convert argument to chain-case style string.
// If argument is empty, return itself.
func ToChainCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "-")
}

// If argument is chain-case style string, return true.
func IsChainCase(s string) bool {
	if strings.Contains(s, "-") {
		fields := strings.Split(s, "-")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	} else {
		return isMadeByLowerAndDigit(s)
	}
}

// Convert argument to camelCase style string
// If argument is empty, return itself
func ToCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	for i, f := range fields {
		if i != 0 {
			fields[i] = toUpperFirstRune(f)
		}
	}
	return strings.Join(fields, "")
}

// If argument is camelCase style string, return true.
// If first character is digit, always returns false
func IsCamelCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	} else {
		return isMadeByAlphanumeric(s) && isFirstRuneLower(s)
	}
}

// Convert argument to PascalCase style string
// If argument is empty, return itself
func ToPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	for i, f := range fields {
		fields[i] = toUpperFirstRune(f)
	}
	return strings.Join(fields, "")
}

// If argument is PascalCase style string, return true.
// If first character is digit, always returns false
func IsPascalCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	} else {
		return isMadeByAlphanumeric(s) && isFirstRuneUpper(s)
	}
}

func isMadeByLowerAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByAlphanumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isFirstRuneUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsUpper(getRuneAt(s, 0))
}

func isFirstRuneLower(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsLower(getRuneAt(s, 0))
}

func isFirstRuneDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsDigit(getRuneAt(s, 0))
}

func getRuneAt(s string, i int) rune {
	if len(s) == 0 {
		return 0
	}

	rs := []rune(s)
	return rs[0]
}

func splitToLowerFields(s string) []string {
	defaultCap := len([]rune(s)) / 3
	fields := make([]string, 0, defaultCap)

	for _, sf := range strings.Fields(s) {
		for _, su := range strings.Split(sf, "_") {
			for _, sh := range strings.Split(su, "-") {
				for _, sc := range camelcase.Split(sh) {
					fields = append(fields, strings.ToLower(sc))
				}
			}
		}
	}
	return fields
}

func toUpperFirstRune(s string) string {
	rs := []rune(s)
	return strings.ToUpper(string(rs[0])) + string(rs[1:])
}
