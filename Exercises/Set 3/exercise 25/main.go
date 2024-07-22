package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where intitialisation for my program occurs")
}

func main() {
	randomInt := rand.Intn(250)

	// fmt.Println("randomint is: ", randomInt)
	// if randomInt <= 100 {
	// 	fmt.Println("Between 0 and 100")
	// } else if randomInt > 100 && randomInt < 201 {
	// 	fmt.Println("Between 101 and 200")
	// } else {
	// 	fmt.Println("Between 201 and 250")
	// }

	switch {
	case randomInt <= 100:
		fmt.Printf("%v is between 0 and 100\n", randomInt)
	case randomInt <= 200:
		fmt.Printf("%v is between 101 and 200\n", randomInt)
	default:
		fmt.Printf("%v is between 201 and 250\n", randomInt)
	}

	// if y != 5 {
	// 	if x < 4 && y < 4 {
	// 		fmt.Println("Both X and Y are less than  4")
	// 	} else if x > 6 && y > 6 {
	// 		fmt.Println("Both x and y are greater than 6 ")
	// 	} else if x >= 4 && x <= 6 {
	// 		fmt.Println("X is between 4 and 6 inclusive :)")
	// 	}
	// } else if y == 5 {
	// 	fmt.Println("Y is 5")
	// 	if x >= 4 && x <= 6 {
	// 		fmt.Println("X is between 4 and 6 inclusive :)")
	// 	}
	// } else {
	// 	fmt.Println("None of the conditions were met")
	// }
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("X is : %v\n", x)
		fmt.Printf("Y is : %v\n", y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both X and Y are less than  4")
		case x > 6 && y > 6:
			fmt.Println("Both x and y are greater than 6 ")
		case y == 5:
			fmt.Println("Y is 5")
		case x >= 4 && x <= 6:
			fmt.Println("X is between 4 and 6 inclusive :)")
		default:
			fmt.Println("None of the conditions were met")
		}
	}
}
