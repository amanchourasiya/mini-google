package engine

import "fmt"

func (idx index) Search(term string) [][]int {

	var r [][]int
	for _, token := range analyze(term) {
		if ids, ok := idx[token]; ok {
			r = append(r, ids)
		}
	}
	fmt.Printf("Search results %v\n", r)
	return r
}
