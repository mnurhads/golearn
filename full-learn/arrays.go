package main

import "fmt"

func main() {
	// simple array
	var x [5]int // An array of 5 integers
	fmt.Println(x)

	var y [8]string // An array of 8 strings
	fmt.Println(y)

	var z [3]complex128 // An array of 3 complex numbers
	fmt.Println(z)

	// array index
	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 105

	fmt.Printf("x[0] = %d, x[1] = %d, x[2] = %d\n", x[0], x[1], x[2])
	fmt.Println("x = ", x)

	// array values
	a1 := [5]string{"English", "Japan", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)

	// array baru
	// Iterating over and Array and printing its elements
	names := [3]string{"Mark Zuckerberg", "Bill Gates", "Larrt Page"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Finding the Sum of an Array
	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}

	fmt.Printf("Sum of all the elements in array  %v = %f\n", a, sum)

	// array range
	// Iterating over an array using range form of for loop
	daysOfWeek := [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggi"}

	for mulai, values := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", mulai, values)
	}

	// Finding the sum of an array
	data := [4]float64{3.5, 7.2, 4.8, 9.5}
	tambah := float64(0)

	for _, values := range data {
		tambah = tambah + values
	}

	fmt.Printf("Sum of all the elements in array %v = %f\n", data, tambah)

	// array multi dimensi
	list := [2][2]int{
		{3, 5},
		{7, 9}, // This comma is necessary
	}
	fmt.Println(list)

	list_data := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(list_data)
}