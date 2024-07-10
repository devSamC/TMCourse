package main

import "fmt"

const banana string = "banana"

var pi float64 = 3.14

func main() {
	cranberry := "cranberries"

	fmt.Printf("The type of %v is %T\n", banana, banana)
	fmt.Printf("The type of %v is %T\n", pi, pi)
	fmt.Printf("The type of %v is %T\n", cranberry, cranberry)

}
