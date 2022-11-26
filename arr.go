package main

import (
	"fmt"
)

func main() {

	// Creating an array
	arr := [7]string{"one", "two", "three", "four",
		"five"}

	// Display array
	fmt.Println("Array:", arr)

	// Creating a slice
	myslice := arr[0:1]
	myslice = arr[0:2]

	// Display slice
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
