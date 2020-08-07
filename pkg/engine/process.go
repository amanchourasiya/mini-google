package engine

/*
This file will process the xml document and load in memory

*/

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

type Document struct {
	Title string `xml:"title`
	URL   string `xml:"url`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	fmt.Print("Processing xml document")
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}
	fmt.Print("Processed xml document")
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	fmt.Print("Assigned ID to all documents")
	return docs, nil
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number
		return !unicode.IsLetter(r) && !unicode.IsNumber((r))
	})
}

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

var stopwords = map[string]struct{}{ // I wish Go had built-in sets.
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordsFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordsFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
