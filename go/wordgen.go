package wordgen

import (
	"context"
	"strconv"
	"strings"

	"github.com/bytedance/gopkg/lang/fastrand"
)

// Generator produces strings containing random words that conform to a provided pattern.
type Generator struct {
	parts  []any
	minLen int
}

// New returns a Generator that is designed to be used repeatedly for a provided pattern.
//
// The pattern can contain '#' (base-10 digit), 'h'/'H' (hexadecimal digit), 'n'/'N'/'Ñ' (nouns), 'a'/'A'/'Ã' (adjectives),
// 'v'/'V'/'Ṽ' (verbs) and '*' (any word). '\' can be used to escape the special characters. Upper-case characters
// generate upper-case words. Tilde characters generate title case words.
//
// Example
//
//	`学中文:n-v-a-n` could return "学中文:dog-ran-red-banana"
func New(pattern string) Generator {
	g := Generator{}

	var str []rune
	var escape bool
	for _, r := range pattern {
		switch r {
		case '*', 'n', 'N', 'Ñ', 'v', 'V', 'Ṽ', 'a', 'A', 'Ã', '#', 'h', 'H':
			if escape {
				escape = false
				str = append(str, r)
			} else {
				if len(str) > 0 {
					g.parts = append(g.parts, str)
					str = nil
				}
				g.parts = append(g.parts, map[rune]int{'*': 0, 'n': 1, 'N': 2, 'Ñ': 3, 'v': 4, 'V': 5, 'Ṽ': 6, 'a': 7, 'A': 8, 'Ã': 9, '#': 10, 'h': 11, 'H': 12}[r])
			}
		case '\\':
			if escape {
				escape = false
				str = append(str, r)
			} else {
				escape = true
			}
		default:
			str = append(str, r)
		}
	}

	if len(str) > 0 {
		g.parts = append(g.parts, str)
		str = nil
	}

	// Calculate minimum length of string
	for _, v := range g.parts {
		switch v := v.(type) {
		case int:
			g.minLen = g.minLen + 1
		case []rune:
			g.minLen = g.minLen + len(string(v))
		}
	}

	return g
}

func findWord(words *map[string]struct{}, list []string) string {
	for {
		word := list[fastrand.Intn(len(list))]
		if _, exists := (*words)[word]; !exists {
			(*words)[word] = struct{}{}
			return word
		}
	}
}

func (g Generator) string() string {
	words := map[string]struct{}{}
	var b strings.Builder
	for _, v := range g.parts {
		switch v := v.(type) {
		case int:
			if v == 0 {
				var word string
				c := len(nouns) + len(verbs) + len(adjectives)
				for {
					r := fastrand.Intn(c)
					if r < len(nouns) {
						word = nouns[r]
					} else if r < len(nouns)+len(verbs) {
						word = verbs[r-len(nouns)]
					} else {
						word = adjectives[r-len(nouns)-len(verbs)]
					}

					if _, exists := words[word]; !exists {
						break
					}
				}
				words[word] = struct{}{}
				b.WriteString(word)
			} else if v == 1 {
				b.WriteString(findWord(&words, nouns))
			} else if v == 2 {
				b.WriteString(strings.ToUpper(findWord(&words, nouns)))
			} else if v == 3 {
				w := findWord(&words, nouns)
				b.WriteString(strings.ToUpper(string(w[0])) + w[1:])
			} else if v == 4 {
				b.WriteString(findWord(&words, verbs))
			} else if v == 5 {
				b.WriteString(strings.ToUpper(findWord(&words, verbs)))
			} else if v == 6 {
				w := findWord(&words, verbs)
				b.WriteString(strings.ToUpper(string(w[0])) + w[1:])
			} else if v == 7 {
				b.WriteString(findWord(&words, adjectives))
			} else if v == 8 {
				b.WriteString(strings.ToUpper(findWord(&words, adjectives)))
			} else if v == 9 {
				w := findWord(&words, adjectives)
				b.WriteString(strings.ToUpper(string(w[0])) + w[1:])
			} else if v == 10 {
				b.WriteString(strconv.Itoa(fastrand.Intn(10)))
			} else if v == 11 || v == 12 {
				charset := "0123456789abcdef"
				n := string(charset[fastrand.Intn(len(charset))])
				if v == 12 {
					n = strings.ToUpper(n)
				}
				b.WriteString(n)
			}
		case []rune:
			b.WriteString(string(v))
		}
	}
	return b.String()
}

// String generates a random string that conforms to the provided pattern.
//
// NOTE: When setting an optional max string length, it should be substantially
// larger than the theoretical minimum to prevent a potential infinite loop. A
// context can be supplied as the first argument to provide a deadline.
// If the deadline is reached, an empty string is returned.
func (g Generator) String(max ...any) string {
	var ctx context.Context
	_max := 0

	for _, v := range max {
		if v, ok := v.(context.Context); ok {
			ctx = v
			continue
		}
		_max = v.(int)
		break
	}

	if len(max) == 0 || _max == 0 {
		return g.string()
	}

	if _max <= g.minLen {
		panic("max must be greater than " + strconv.Itoa(g.minLen) + " (min length)")
	}

	for {
		if ctx != nil && ctx.Err() != nil {
			return ""
		}
		s := g.string()
		if len(s) <= _max {
			return s
		}
	}

	panic("should not reach here")
}

// Error must not be used! It is only defined so that the fmt
// package can "pretty print" the Generator since the Stringer
// interface can't be implemented.
//
// NOTE: Do not use Generator as an error.
func (g Generator) Error() string {
	var b strings.Builder
	for _, v := range g.parts {
		switch v := v.(type) {
		case int:
			b.WriteString("[" + string([]rune{'*', 'n', 'N', 'Ñ', 'v', 'V', 'Ṽ', 'a', 'A', 'Ã', '#', 'h', 'H'}[v]) + "]")
		case []rune:
			b.WriteString(string(v))
		}
	}

	return b.String() + " (min: " + strconv.Itoa(g.minLen) + " bytes)"
}
