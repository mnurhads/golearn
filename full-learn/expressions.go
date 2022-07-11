package main
import "fmt"

func main() {
	var berat = 21.0
	if berat < 18.5 {
		fmt.Println("You are underweight");
	} else if berat >= 18.5 && berat < 25.0 {
		fmt.Println("Your weight is normal");
	} else if berat >= 25.0 && berat < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}

	// switch case
	var dayOfWeek = 6
	switch dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
		default: fmt.Println("Invalid day")
	}

	// loop
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", num)
			break
		}
	}

	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue;
		}
		fmt.Printf("%d ", num)
	}
}