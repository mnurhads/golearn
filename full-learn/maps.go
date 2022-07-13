package main

import (
	"fmt"
)

// func new
func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}

// end
func main() {
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	m["one hundred"] = 100
	fmt.Println(m)

	// maps literal
	fmt.Println("----------------------------------")
	var angka = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, 
	}

	fmt.Println(angka)

	fmt.Println("----------------------------")
	// maps key
	var tinderMatch = make(map[string]string)

	// Adding keys to a map
	tinderMatch["budi"] = "Angelina"
	tinderMatch["James"] = "Sophia"
	tinderMatch["David"] = "Emma"

	fmt.Println(tinderMatch)

	// match
	tinderMatch["budi"] = "Jennifer"
	fmt.Println(tinderMatch)

	// map exist
	fmt.Println("----------------------------------")
	var employees = map[int]string{
		1001: "Muhammad",
		1002: "Nur",
		1003: "Hadi",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}