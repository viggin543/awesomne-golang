package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlicesCanGrow(t *testing.T) {
	SlicesCanGrow()
	assert.EqualValues(t, []int{1, 2, 3}, SliceLiteral())
	SliceLiteralWithCapacity()
	NilSlice()
	EmptySlice()
	SlicingASlice()
	IterateASlice()
	IterationTrap()
	GoodOldForLoop()
	PassingASliceToAFunciton()
}
