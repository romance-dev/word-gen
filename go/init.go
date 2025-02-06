package wordgen

import (
	"bufio"
	"bytes"
	_ "embed"
	"runtime"
)

//go:embed nouns.list
var _nouns []byte

//go:embed verbs.list
var _verbs []byte

//go:embed adjectives.list
var _ads []byte

var (
	nouns      []string // 1666
	verbs      []string // 5492
	adjectives []string // 2397
)

func parse(raw *[]byte, parsed *[]string) {
	scanner := bufio.NewScanner(bytes.NewReader(*raw))
	for scanner.Scan() {
		*parsed = append(*parsed, scanner.Text())
	}
	*raw = nil
}

func init() {
	parse(&_nouns, &nouns)
	parse(&_verbs, &verbs)
	parse(&_ads, &adjectives)
	runtime.GC()
}
