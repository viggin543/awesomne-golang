package __structs_interfaces_embedding

import (
	"fmt"
)

type notifier interface {
	notify() // [idea tip] find implementation cmd+alt+B
}

type admin struct {
	user  // embedding  user is called the "Inner type" and admin the "Outer type"
	level string
}

func (a *admin) notify() { // overriding the embedded method
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

//  a polymorphic function ( accepts different implementations of notifier )
func sendNotification(n notifier) { // references vs values
	n.notify()
}
