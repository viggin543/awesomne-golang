package __structs_interfaces_embedding

import (
	"fmt"
)

type notifier interface {
	notify()
}

type admin struct {
	user  // embedding
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// sendNotification accepts values that implement the notifier interface
func sendNotification(n notifier) {
	n.notify()
}
