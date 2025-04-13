package main

import (
	"fmt"
)

func main() {
	fmt.Println("=================== Array ===================")

	// There are two types of array in programming languages which is given below
	// 1. Static Array
	// 2. Dynamic Array

	// Static Array
	// Static array is a fixed size array that is declared with a specific size.
	// The size of the array is determined at the time of declaration.

	// Dynamic Array
	// Dynamic array is a variable size array that is declared with a specific size.
	// The size of the array is determined at the time of declaration.

	// Note: Array can be one dimensional and multi dimensional

	// We can perform the following operations on the array
	// 1. Insertion
	// 2. Deletion
	// 3. Search
	// 4. Update
	// 5. Traverse

	// Static array:
	// var array [5]int = [5]int{1, 2, 3, 4, 5} // Static array
	// fmt.Println(array)

	// Dynamic array: Dynamic array in golang is called slice
	// var dynamicArray []int = []int{1, 2, 3, 4, 5}
	// fmt.Println(dynamicArray)

	// Insertion and updation:

	// var newValue int
	// var indexForInsertion int
	// fmt.Println("Enter the new value to be inserted: ")
	// fmt.Scanln(&newValue)
	// fmt.Println("Enter the index to be inserted for static array: ")
	// fmt.Scanln(&indexForInsertion)

	// array[4] = newValue // We can't use append in the static array as we are doing with the dynamic array in the following code
	// fmt.Println("After insertion in the static array: ")
	// array[indexForInsertion] = newValue
	// fmt.Println(array)

	// dynamicArray = append(dynamicArray, newValue)
	// fmt.Println("After insertion in the dynamic array: ")
	// fmt.Println(dynamicArray)

	// Deletion:
	// - Dynamic Array:
	// deleteFromDynamicArray(&dynamicArray)
	// fmt.Println(dynamicArray)

	// - Static Array:
	// deleteFromStaticArray(&array, 5)
	// fmt.Println(array)

	// Search:
	// - Static Array:
	// var value int
	// fmt.Println("Enter the value to be searched in the static array: ")
	// fmt.Scanln(&value)
	// index := searchFromStaticArray(&array, value)
	// if index == -1 {
	// 	fmt.Println("Value not found")
	// } else {
	// 	fmt.Println("Value found at index: ", index)
	// }

	// - Dynamic Array:
	// fmt.Println("Enter the value to be searched in the dynamic array: ")
	// fmt.Scanln(&value)
	// index = searchFromDynamicArray(&dynamicArray, value)
	// if index == -1 {
	// 	fmt.Println("Value not found")
	// } else {
	// 	fmt.Println("Value found at index: ", index)
	// }

	// Insert at Index:

	testSlice := make([]int, 20, 50)
	insertAtIndex(&testSlice, 9)
	fmt.Println(testSlice)
}

// func deleteFromStaticArray(array *[5]int, size int) {
// 	var index int
// 	fmt.Println("Enter the index to the value to be deleted: ")
// 	fmt.Scanln(&index)

// 	// if the current index is less than the index for the deletion then we do nothing,
// 	// if the index is equal to the index then we will shift the array to the left by one position

// 	var i = 0
// 	for ; i < size; i++ {
// 		if i >= index && i != size-1 {
// 			(*array)[i] = (*array)[i+1]
// 		}
// 		if i == size-1 {
// 			(*array)[i] = 0
// 		}
// 	}

// }

// func deleteFromDynamicArray(array *[]int) {
// 	var index int
// 	fmt.Println("Enter the index to the value to be deleted: ")
// 	fmt.Scanln(&index)

// 	*array = append((*array)[:index], (*array)[index+1:]...)
// 	// Moder approach:
// 	// *array = slices.Delete(*array, index, index+1)
// }

// // This function will return the first occurence of the value or nil if not found:
// func searchFromStaticArray(array *[5]int, value int) int {
// 	var i int = 0
// 	for ; i < len(*array); i++ {
// 		if (*array)[i] == value {
// 			return i
// 		}
// 	}
// 	return -1
// }

// // This function will return the first occurence of the value or nil if not found:
// func searchFromDynamicArray(array *[]int, value int) int {
// 	var i int = 0
// 	for ; i < len(*array); i++ {
// 		if (*array)[i] == value {
// 			return i
// 		}
// 	}
// 	return -1
// }

func insertAtIndex(array *[]int, value int) {
	var index int
	fmt.Println("Enter the index to be inserted: ")
	fmt.Scanln(&index)

	// Insert at starting: if index is 0 we have to shift other elements to the right
	// Insert in between: If in between we have to shift the element to the right of the index

	// Insert at last
	if index == 0 {
		var i = len(*array) - 1
		for ; i > 0; i-- {
			(*array)[i] = (*array)[i-1]
		}
		(*array)[0] = value
	} else if index > 0 && index < len(*array) {
		var i = len(*array) - 1
		for ; i > index; i-- {
			(*array)[i] = (*array)[i-1]
		}
		(*array)[index] = value
	}

}
