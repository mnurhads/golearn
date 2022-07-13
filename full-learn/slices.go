package main

import (
	"fmt"
)

func main() {
	// varian
	var a = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	// Creating a slice from the array
	var s []string = a[1:4]

	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", s)

	slice1 := a[1:4]
	slice2 := a[:3]
	slice3 := a[2:]
	slice4 := a[:]

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)

	//slice other
	fmt.Println("----------------------------------------")
	cities := []string{"Jermani", "Sundong", "Prindavan", "Beijing", "Jogja", "Malang", "Bangalore", "Hyderabad", "Hong Kong"}

	asianCities := cities[3:]
	indoCities := asianCities[1:3]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
	fmt.Println("indoCities = ", indoCities)

	// SLice Program
	fmt.Println("---------------------------------------")
	bahasa1 := []string{"C", "C++", "Java"}
	bahasa2 := append(bahasa1, "Python", "Ruby", "Go")

	fmt.Printf("bahasa1 = %v, len = %d, cap = %d\n", bahasa1, len(bahasa1), cap(bahasa1))
	fmt.Printf("bahasa2 = %v, len = %d, cap = %d\n", bahasa2, len(bahasa2), cap(bahasa2))

	slice1[0] = "C#"
	fmt.Println("\nbahasa1 = ", bahasa1)
	fmt.Println("bahasa2 = ", bahasa2)
}