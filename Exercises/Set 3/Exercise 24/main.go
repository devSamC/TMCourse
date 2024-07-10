package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomInt := rand.Intn(250)

	fmt.Println("randomint is: ", randomInt)
	if randomInt <= 100 {
		fmt.Println("Between 0 and 100")
	} else if randomInt > 100 && randomInt < 201 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}
}
