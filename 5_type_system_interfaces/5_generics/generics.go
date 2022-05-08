package __generics

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Map[T any, R any](t []T, f func(t T) R) (ret []R) {
	for _, v := range t {
		ret = append(ret, f(v))
	}
	return
}

func Reduce[T any, R any](t []T, f func(t T, r R) R, zero R) R {
	panic("implement me")
}

// [idea tip] -> step into comparable implementation ( cmd + b )
// notice "int64 | float64" a union type

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Number interface to be a type constraint
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

//---
// creating a custom generic type

type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() { // a generic method
	for _, v := range g {
		fmt.Println(v)
	}
}

// a generic struct

type Box[T constraints.Ordered] struct { //constraints package have useful stuf
	Val T
}

// we can't define type params in methods, only functions
func (b *Box[T]) equals(t T) bool {
	return b.Val == t // this is possible since T is comparable
	// what will happen if T becomes any ?
}

type Signal int16 // type override, type alias

func Process[T ~int16](value T) T { // try removing the `~` operator
	return value - 1
}

// ~ tells the go compiler, any type override that can be reduced to  int16
// generics introduced type constraints which brake the duck typing paradigm.
// there is no type erasure in golang, unlike java
