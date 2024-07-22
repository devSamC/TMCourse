package main

import "fmt"

func main() {
	slice := make([][]string, 0)
	james := []string{"James", "bond", "Shaken, not stirred"}
	moneypenny := []string{"Miss", "Moneypenny", "Im 008"}

	slice = append(slice, james, moneypenny)

	fmt.Println(slice)

	for _, v := range slice {
		fmt.Println(v)
		for index, value := range v {
			fmt.Printf("At index %v: %v\n", index, value)
		}
	}
}
