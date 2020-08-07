package main

import (
	"github.com/amanchourasiya/mini-google/pkg/engine"
)

func main() {
	idx := engine.CreateIndex()
	idx.Search("Small WIld cat")
}
