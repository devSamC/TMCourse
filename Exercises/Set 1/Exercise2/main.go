package main 

import (
	"fmt"
)

func main() {

	adams := 42 
	fmt.Printf("42 as a hexadecimal, %x \n", adams)
	fmt.Printf("42 as binary, %b \n", adams )

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%v \t as hex: %#x \t as binary: %b\n",a,a,a)
	fmt.Printf("%v \t as hex: %#x \t as binary: %b\n",b,b,b)
	fmt.Printf("%v \t as hex: %#x \t as binary: %b\n",c,c,c)
	fmt.Printf("%v \t as hex: %#x \t as binary: %b\n",d,d,d)
	fmt.Printf("%v \t as hex: %#x \t as binary: %b\n",e,e,e)
	fmt.Printf("%v \t as hex: %#x \t as binary: %b\n",f,f,f)
}