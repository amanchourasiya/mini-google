package engine

import (
	"strings"
)

func Search(docs []Document, term string) ([]Document){
	var r []Document
	for _ , doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}