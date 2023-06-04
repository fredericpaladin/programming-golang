package main

import (
	"fmt"
	"sort"
)

func main() {
	x := 40
	pointer := &x

	fmt.Println("The value of x is", x)
	fmt.Println("The address of pointer is", pointer)
	fmt.Println("The value of pointer is", *pointer)

	*pointer = 156

	fmt.Println("The new value of x is", x)
	fmt.Println("The new address of pointer is", pointer)
	fmt.Println("The new value of pointer is", *pointer)

	changeValue(&x) // Pass the reference of x
	fmt.Println("The new new value of x is", x)

	// ---------------------------------------------------

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

	//------------------------------
	// maps are reference types, so they are always passed by reference. You don't need a pointer.
	// maps are keys and values, and you must use the make function when creating
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// range through the map and print out key and value
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete a value from a map by key
	delete(intMap, "four")

	// check to see if an item exists in a map by key
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is in not map")
	}

	// update a value in the map
	intMap["two"] = 4
}

func changeValue(num *int) {
	*num = 25
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // Copy the last element into 'i'
	a = a[:len(a)-1]   // Truncate

	return a
}
