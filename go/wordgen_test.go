package wordgen_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/romance-dev/word-gen/go"
)

type Meaning struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Definitions  []struct {
		Definition string   `json:"definition"`
		Synonyms   []string `json:"synonyms"`
		Antonyms   []string `json:"antonyms"`
		Example    string   `json:"example"`
	} `json:"definitions"`
	Synonyms []string `json:"synonyms"`
	Antonyms []string `json:"antonyms"`
}

type Response struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl"`
	} `json:"phonetics"`
	Meanings []Meaning `json:"meanings"`
	Origin   string    `json:"origin"`
}

func TestNoun(t *testing.T) {
	g := wordgen.New("n")

	N := 1000

	for i := 0; i < N; i++ {
		word := g.String()

		res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		res.Body.Close()

		var r []Response
		err = json.Unmarshal(body, &r)
		if err != nil {
			if strings.Contains(string(body), "error code: 1015") {
				continue
			}

			if strings.Contains(string(body), "No Definitions Found") {
				t.Errorf("%d: not a word: %s", i, word)
				continue
			}
			panic(err)
		}

		if len(r) == 0 {
			t.Errorf("%d: not a word: %s", i, word)
		}

		isNoun := false
	OUTER:
		for _, rr := range r {
			for _, m := range rr.Meanings {
				if m.PartOfSpeech == "noun" {
					isNoun = true
					break OUTER
				}
			}
		}

		if !isNoun {
			t.Errorf("%d: not a noun: %s", i, word)
		}
	}
}

func TestVerb(t *testing.T) {
	g := wordgen.New("v")

	N := 2000

	for i := 0; i < N; i++ {
		word := g.String()

		res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		res.Body.Close()

		var r []Response
		err = json.Unmarshal(body, &r)
		if err != nil {
			if strings.Contains(string(body), "error code: 1015") {
				continue
			}

			if strings.Contains(string(body), "No Definitions Found") {
				t.Errorf("%d: not a word: %s", i, word)
				continue
			}
			panic(err)
		}

		if len(r) == 0 {
			t.Errorf("%d: not a word: %s", i, word)
		}

		isVerb := false
	OUTER:
		for _, rr := range r {
			for _, m := range rr.Meanings {
				if m.PartOfSpeech == "verb" {
					isVerb = true
					break OUTER
				}
			}
		}

		if !isVerb {
			t.Errorf("%d: not a verb: %s", i, word)
		}
	}
}

func TestAdjective(t *testing.T) {
	g := wordgen.New("a")

	N := 1000

	for i := 0; i < N; i++ {
		word := g.String()

		res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		res.Body.Close()

		var r []Response
		err = json.Unmarshal(body, &r)
		if err != nil {
			if strings.Contains(string(body), "error code: 1015") {
				continue
			}

			if strings.Contains(string(body), "No Definitions Found") {
				t.Errorf("%d: not a word: %s", i, word)
				continue
			}
			panic(err)
		}

		if len(r) == 0 {
			t.Errorf("%d: not a word: %s", i, word)
		}

		isAdj := false
	OUTER:
		for _, rr := range r {
			for _, m := range rr.Meanings {
				if m.PartOfSpeech == "adjective" {
					isAdj = true
					break OUTER
				}
			}
		}

		if !isAdj {
			t.Errorf("%d: not an adjective: %s", i, word)
		}
	}
}
