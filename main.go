package main

import "fmt"

func main() {
	x := 5
	fmt.Println(&x)

	ptr := &x
	fmt.Println(*ptr)
	fmt.Println("Hello, World! sarasa")
}
