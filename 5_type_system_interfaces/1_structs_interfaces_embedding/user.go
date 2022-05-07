package __structs_interfaces_embedding

import (
	"fmt"
)

// a struct is a custom data type.
// what is the difference between data structs and  software "objects" ?
type user struct { // private struct
	name  string
	email string
}

// method with a non pointer receiver
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
	// functions have no receivers
	// methods have receivers
}

// a pointer receiver. ( almost always used in real world )
func (u *user) changeEmail(email string) {
	u.email = email
}
