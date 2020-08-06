package main

import (
	"fmt"

	"github.com/amanchourasiya/mini-google/pkg/engine"
)

func main() {
	docs, err := engine.LoadDocuments("/root/enwiki-latest-abstract1.xml")
	if err != nil {
		fmt.Printf("Error loading xml document %s", err)
	}
	_ = engine.Search(docs, "cat")
	fmt.Printf("Results found")
}
