package __generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	opa := Map([]int{1, 2, 3, 4}, func(t int) string {
		return fmt.Sprintf("%d", t*t)
	})
	fmt.Println(opa)
}

func TestReduce(t *testing.T) {
	res := Reduce([]int{1, 2, 3, 4}, func(t int, r float64) float64 {
		return float64(t) + r
	}, 0)
	assert.EqualValues(t, 10, res)
}

func TestSumIntsOrFloats(t *testing.T) {
	floats := SumIntsOrFloats(map[string]int64{
		"a": 1,
		"b": 1,
	})
	assert.EqualValues(t, 2, floats)
}

func TestSumNumbers(t *testing.T) {
	floats := SumNumbers(map[string]int64{
		"a": 1,
		"b": 1,
	})
	assert.EqualValues(t, 2, floats)
}

func TestGenericSlice_Print(t *testing.T) {
	g := GenericSlice[int]{1, 2, 3}
	g.Print()
}

func TestBox_equals(t *testing.T) {
	var x = new(Box[int]) // generic structs must be instantiated ( using new or make, like channels )
	x.Val = 123
	assert.True(t, x.equals(123))
}

func TestProcess(t *testing.T) {
	Process(int16(3))
}

// Example of Type Chaining
// U is defined in terms of T
func ToChan[U ~[]T, T any](t U) <-chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for _, x := range t {
			c <- x
		}
	}()
	return c
}
