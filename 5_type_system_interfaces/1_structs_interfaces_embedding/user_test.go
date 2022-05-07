package __structs_interfaces_embedding

import "testing"

func TestUser(t *testing.T) {
	// declare a variable using the `var` keyword only to indicate its initialized to its zero value
	bill := user{"Bill", "bill@email.com"} //  implicit struct key names
	bill.notify()
	// [idea tip] => ctrl+j to view lisa definition
	lisa := &user{name: "Lisa", email: "lisa@email.com"} // explicit key names
	lisa.notify()

	// (&bill).changeEmail() behind the scene [ go compiler ]
	bill.changeEmail("bill@newdomain.com")
	// what would have happened if changeEmail had a non pointer receiver ?
	bill.notify()
	// (*lisa).notify()  behind the scene [ go compiler ]
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
