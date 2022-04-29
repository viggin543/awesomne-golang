package search

import (
	"encoding/json"
	"github.com/viggin543/awesomne-golang/code/chapter2/sample/data"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetreiveEmeddedFeeds() (feeds []*Feed, err error) {
	err = json.Unmarshal(data.Feeds, &feeds)
	return
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
