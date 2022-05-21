package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlicesCanGrow(t *testing.T) {
	SlicesCanGrow()

}

func TestPassingASSliceToAFFunction(t *testing.T) {
	PassingASliceToAFunciton()
}

func TestGoodOldForLoop(t *testing.T) {
	GoodOldForLoop()
}

func TestIterationTrap(t *testing.T) {
	IterationTrap()
}

func TestIterateASlice(t *testing.T) {
	IterateASlice()
}

func TestSlicingASlice(t *testing.T) {
	SlicingASlice()
}

func TestEmptySlice(t *testing.T) {
	EmptySlice()
}

func TestNilSlice(t *testing.T) {
	NilSlice()
}

func TestSliceLitter_withCap(t *testing.T) {
	SliceLiteralWithCapacity()
}

func TestSliceLiteral(t *testing.T) {
	assert.EqualValues(t, []int{1, 2, 3}, SliceLiteral())

}
