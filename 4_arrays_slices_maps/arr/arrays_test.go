package arr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceMe(t *testing.T) {
	ArraysAreFixedInSize()
}

func TestPtrSlice(t *testing.T) {
	ArrayPointers()
}

func TestStringSlice(t *testing.T) {
	StringArrays()
}

func TestStringPointers(t *testing.T) {
	assert.True(t, StringPointers())
}

func TestMultiDimArr(t *testing.T) {
	MultiDimArr()
}

// go benchmark tool is awesome !!!
func BenchmarkPassingLargeArrByValue(b *testing.B) {
	PassingArraysToFunctions()
}
