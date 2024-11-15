package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a slice by make function
	var gpaSlice []float64
	gpaSlice = make([]float64, 3)
	gpaSlice[0] = 4.1
	gpaSlice[1] = 3.9
	gpaSlice[2] = 4.23
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slice litaral
	// gpasSlice := []float64{4.1, 3.9, 4.23} // slice litaral
	// fmt.Println(gpasSlice, reflect.TypeOf(gpasSlice))

	// Create a slice by slicing an existing array
	// gpas := [5]float64{3.5, 4.1, 3.9, 4.23} // array := array litaral
	// fmt.Println(gpas, reflect.TypeOf(gpas))
	// gpasSlice := gpas[1:4] // slice := slicing array
	// fmt.Println(gpasSlice, reflect.TypeOf(gpasSlice))

}
