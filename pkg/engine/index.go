package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type index map[string][]int

func (idx index) add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func CreateIndex() index {
	idx := make(index)
	docs, err := LoadDocuments("/root/enwiki-latest-abstract1.xml")
	if err != nil {
		fmt.Printf("Failed to load documents %v\n", err)
		return nil
	}
	idx.add(docs)
	return idx
}

func SaveIndex(idx index) {
	fmt.Println("Saving Index")
	jsonData, err := json.MarshalIndent(idx, "", " ")
	if err != nil {
		fmt.Printf("Error marshaling index to json %v\n", err)
	}
	err = ioutil.WriteFile("/root/index.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error saving index json in json file %v\n", err)
	}

}

func LoadIndex(idx index) index {

}
