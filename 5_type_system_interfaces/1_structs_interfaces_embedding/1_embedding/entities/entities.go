package entities

// what are cyclic dependencies ?

type user struct {
	Name   string
	Email  string
	banana string // a private struct field
}

type Admin struct {
	user   // The embedded type is unexported.
	Rights int
}
