package engine

import "fmt"

func displayResults(resultIndex [][]int) {
	docs := getDocuments()
	fmt.Printf("Documents retrieved %d\n", len(docs))

	for _, ids := range resultIndex {
		for _, id := range ids {
			fmt.Printf("result link %d %s\n", id, DB[id])
		}
	}

}
