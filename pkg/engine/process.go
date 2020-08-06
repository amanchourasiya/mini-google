package engine

/*
This file will process the xml document and load in memory

*/

import  (
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title`
	URL   string `xml:"url`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([] Document, error) {
	f, err := os.Open(path)
	if err != nil{
		return nil, err
	}
	defer f.Close()
	
	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil{
		return nil, err
	}
	docs := dump.Documents
	for i := range docs{
		docs[i].ID = i
	}
	return docs, nil
}
