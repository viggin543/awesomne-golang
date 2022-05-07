package __structs_interfaces_embedding

import "testing"

func TestName(t *testing.T) {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// The embedded inner type's implementation of the
	// interface is NOT "promoted" to the outer type.
	sendNotification(&ad)

	// We can access the inner type's method directly.
	ad.user.notify() // composition vs inheritance

	// The inner type's method is NOT promoted.
	ad.notify() // you can sort of call this "overriding"
	// what would this call if we remove the embedding of a user in an admin ?
}
