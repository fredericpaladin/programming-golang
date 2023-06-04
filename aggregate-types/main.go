package main

import "fmt"

type Car struct {
	Make       string
	Model      string
	Year       int
	NumOfTires int
	IsLuxury   bool
}

func main() {

	var myStrings [3]string
	myStrings[0] = "Cat"
	myStrings[1] = "Dog"
	myStrings[2] = "Fish"

	fmt.Println("First element is", myStrings[0])

	// var myCar Car
	// myCar.Make = "BMW"
	// myCar.Model = "335"
	// myCar.Year = 2014
	// myCar.NumOfTires = 4
	// myCar.IsLuxury = false
	myCar := Car{
		Make:       "BMW",
		Model:      "335",
		Year:       2014,
		NumOfTires: 4,
		IsLuxury:   false,
	}

	fmt.Printf("My car is a %d %s %s.\n", myCar.Year, myCar.Make, myCar.Model)
}
