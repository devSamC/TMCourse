package main

import "fmt"

func main() {
	var array1 [5]int

	for i := range array1 {
		array1[i] = i
	}

	fmt.Println(array1)
	fmt.Printf("%v is of type %T", array1, array1)

}
