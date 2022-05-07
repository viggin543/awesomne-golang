package __embedding

import (
	"encoding/json"
	"fmt"
	"github.com/viggin543/awesomne-golang/chapter5/1_structs_interfaces_embedding/1_embedding/entities"
)

// why this main is not executable ?
func main() {
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	marshal, _ := json.Marshal(a)
	fmt.Printf("%v\n", a)
	fmt.Println("")
	fmt.Println(string(marshal)) // json marshaling handles the embedding as you would expect
}
