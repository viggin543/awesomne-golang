package __structs_interfaces_embedding

import "testing"

func TestName(t *testing.T) {
	adminush := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	// The embedded inner type's implementation of the
	// interface is NOT "promoted" to the outer type. ( admins implementation of notify 'overrides' embedded implementation )
	sendNotification(&adminush) // why passing a reference here ?

	// We can access the inner type's method directly.
	adminush.user.notify() // composition vs inheritance

	// The inner type's method is NOT promoted.
	adminush.notify() // you can sort of call this "overriding"
	// what would this call if admin did not implement a notifier interface ?
	// would admin be a notifier then ?
	// this is called "inner type promotion" in golang
}
