package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)
	fmt.Println(&i)  // memory location of i, aka reference
	fmt.Println(*&i) // dereferencing the memory location of i, aka getting the actual value
	*&i = 2          // setting a new value in the same memory location
	fmt.Println(i)
	fmt.Println(&i)

}
