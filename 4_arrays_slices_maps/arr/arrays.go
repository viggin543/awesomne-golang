package arr

import (
	"encoding/json"
	"fmt"
)

// Declare an integer array of five elements.
// arrays are always fix size and can't grow !!!
var array [5]int

func ArraysAreFixedInSize() {
	array = [5]int{10, 20, 30, 40, 50}
	// Capacity is determined based on the number of values initialized.
	//[idea tip -> ctrl+j: peek definition ( array variable ) ]
	array = [...]int{10, 20, 30, 40, 50}
	//// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	array = [5]int{1: 10, 2: 20}
	array[2] = 35

}

func ArrayPointers() {
	// Initialize index 0 and 1 of the array with integer pointers.
	pointers := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*pointers[0] = 10
	*pointers[1] = 20
	// open debug window to view the pointer addresses
}

// [tip -> generate test/jump to test cmd+shift+t]

func StringArrays() [5]string {
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// Deep copy the values from array2 into array1.
	//Only arrays of the same type can be assigned. [5]string
	array1 = array2
	array2[1] = "banana"
	return array1
}

func StringPointers() bool {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"
	// Copying an array of pointers copies the pointer values and not the values that the pointers are pointing to
	array1 = array2
	*array2[1] = "banana"

	return *array1[1] == "banana"
}

func MultiDimArr() {
	var array [4][2]int
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	array[0][0] = 10 // static typing allows the compiler to know each array size in ram and a fast random access
	marshal, _ := json.Marshal(array)
	fmt.Println(string(marshal))

}

func PassingArraysToFunctions() {
	//  of 8 megabytes.
	var arr [1e6]int
	foo(arr)
	fmt.Println(arr[0]) // arrays are passed by value.

}

func foo(array [1e6]int) { // this will copy 8 megabytes on every func call. ouch...
	array[0] = 1
	fmt.Println(len(array))
}
