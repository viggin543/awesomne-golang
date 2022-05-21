package entities

// what are cyclic dependencies ?

type user struct {
	Name   string
	Email  string
	banana string // a private struct field
}

type Admin struct {
	user   // The embedded type is unexported. outside the package no one will know it exists.
	Rights int
}
