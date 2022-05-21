package search

import (
	"encoding/json"
	"github.com/viggin543/awesomne-golang/2_example_program/sample/data"
	"os"
)

// compile time variable
const dataFile = "data/data.json"

// Feed a struct (new data type) with reflection tags
// structs zero values
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// defer serves a similar cause as a finally clause/
	defer file.Close()

	var feeds []*Feed // why []*Feed and not []Feed ?
	// json unmarshalling ( string\bytes to data type )
	//[tip] -> json to struct https://mholt.github.io/json-to-go/
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}

func RetreiveEmeddedFeeds() (feeds []*Feed, err error) { // named return variables
	// modern way (sine 1.16) to embed data in golang projects
	// why this is useful ?
	err = json.Unmarshal(data.Feeds, &feeds)
	return // naked return
}
