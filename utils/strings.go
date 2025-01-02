package utils

import (
	"regexp"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func RuneLen(s string) int {
	return utf8.RuneCountInString(s)
}

func RemoveAccent(r rune) rune {
	// Normalize the string to decomposed form (NFD)
	decomposed := norm.NFD.String(string(r))

	// Filter out combining marks
	result := make([]rune, 0, len(decomposed))
	for _, char := range decomposed {
		if !unicode.Is(unicode.Mn, char) { // Mn: Non-spacing marks (accents)
			result = append(result, char)
		}
	}

	return result[0]
}

func RemoveAccentString(s string) string {
	res := make([]rune, RuneLen(s))
	for i, c := range []rune(s) {
		res[i] = RemoveAccent(c)
	}

	return string(res)
}

func MatchRegex(s string, regex string) bool {
	r := regexp.MustCompile(regex)
	return r.MatchString(RemoveAccentString(s))
}
