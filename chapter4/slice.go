package chapter4

// Declare an integer array of five elements.
var array [5]int

func Slices() {
	// Initialize each element with a specific value.
	array = [5]int{10, 20, 30, 40, 50}
	// Capacity is determined based on the number of values initialized.
	//[idea tip -> ctrl+j: peek definition]
	array = [...]int{10, 20, 30, 40, 50}
	//// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	array = [5]int{1: 10, 2: 20}
	// Change the value at index 2.
	array[2] = 35

}

func PointersSlice() {
	// Initialize index 0 and 1 of the array with integer pointers.
	pointers := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*pointers[0] = 10
	*pointers[1] = 20
	// open debug window to view the pointer addresses
}

//[tip -> generate test cmd+shift+t]
func StringSlice() [5]string {
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// Deep copy the values from array2 into array1.
	//Only arrays of the same type can be assigned. [5]string
	array1 = array2
	array2[1] = "banana"
	return array1
}
