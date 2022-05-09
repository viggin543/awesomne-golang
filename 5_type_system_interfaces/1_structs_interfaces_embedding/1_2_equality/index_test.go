package __2_equality

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestEquality(t *testing.T) {

	type Person struct {
		Name string
	}

	a := Person{"Bill DeRose"}
	b := Person{"Bill DeRose"}

	assert.True(t, a == b)

	//--------
	type Person2 struct {
		Friend *Person
	}

	a1 := Person2{Friend: &Person{}}
	b1 := Person2{Friend: &Person{}}
	assert.True(t, a1 == b1)

	reflect.DeepEqual(a, b) // true
	//recursive deep equal ( follow pointers )
}
