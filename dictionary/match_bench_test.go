package dictionary

import (
	"crossword/matcher"
	"regexp"
	"testing"
)

func BenchmarkMatchRegex(b *testing.B) {
	const benchmarkWord = "basse"

	b.Run("regex", func(b *testing.B) {
		r := regexp.MustCompile("^.A.S.$")
		for i := 0; i < b.N; i++ {
			r.MatchString(benchmarkWord)
		}
	})

	b.Run("pattern", func(b *testing.B) {
		pattern := []rune("_A_S_")
		for i := 0; i < b.N; i++ {
			matcher.Pattern(benchmarkWord, pattern)
		}
	})
}
