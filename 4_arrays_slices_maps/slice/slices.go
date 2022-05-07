package slice

import (
	"fmt"
	"github.com/samber/lo"
)

func SlicesCanGrow() []string {
	// Contains a length and capacity of 5 elements.
	slice := make([]string, 5)
	// Contains a length of 3 and has a capacity of 5 elements.
	slice = make([]string, 3, 5)
	return slice
}

func SliceLiteral() []int {
	slice := []int{1, 2, 3}
	return slice
}

func SliceLiteralWithCapacity() []string {
	return []string{99: ""}
}

func NilSlice() []int {
	var slice1 []int                  // slice1 is nil
	var a [2]int                      //can an array be nil ?
	return append(slice1, a[0], a[1]) // growing a slice
	// a slice is smart dynamically growing
}

func EmptySlice() []int {
	slice := make([]int, 0)
	slice = []int{}
	return slice
}

func SlicingASlice() []int {
	slice := []int{10, 20, 30, 40, 50}
	ints := slice[2:4] // 30,40
	ints[0] = -1       // what slice[2] is ?
	return ints
}

func IterateASlice() {
	for index, value := range []int{10, 20, 30, 40} {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}
	// OMG, GENERICS !
	// [idea tip] -> peek implementation ( ctrl + shift+ i )
	lo.ForEach([]string{"a", "b", "c", "d"}, func(t string, i int) {
		fmt.Printf("Index: %d  Value: %s\n", i, t)
	})
}

func IterationTrap() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		//value is a copy of slice[i]. its address is the same every iteration !!! don't ever return the address of the for loop value variable !
		fmt.Printf("Value: %d  Value-Addr: %X  ElemAddr: %X\n",
			value, &value, &slice[index])
	}
}

func GoodOldForLoop() {
	for index := 2; index < len([]int{10, 20, 30, 40}); index++ {
		fmt.Printf("Index: %d  Value: %d\n", index, []int{10, 20, 30, 40}[index])
	}
}

func PassingASliceToAFunciton() {
	a := []int{1e6: 9}
	foo(a)
	fmt.Println(a[0])
}

func foo(ints []int) { // a slice is a reference type, it always passed by reference, and its zero value is nil
	ints[0] = -1
	fmt.Println(ints)
}

// prefer arrays to slices !
