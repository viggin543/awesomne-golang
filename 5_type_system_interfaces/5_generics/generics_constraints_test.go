package __generics

import "testing"

func TestGenerticConstrains(t *testing.T) {

	category := Category{
		ID:   1,
		Name: "Go Generics",
		Slug: "go-generics",
	}
	cc := New[Category]()
	cc.Set(category.Slug, category)

	post := Post{
		ID: 1,
		Categories: []Category{
			{ID: 1, Name: "Go Generics", Slug: "go-generics"},
		},
		Title: "Generics in Golang structs",
		Slug:  "generics-in-golang-structs",
	}
	cp := New[Post]()
	cp.Set(post.Slug, post)

}
