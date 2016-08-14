package main

import "fmt"

func main() {
	var b bool
	fmt.Println(b)

	b = true
	fmt.Println(b)

	b = false
	fmt.Println(b)

	b = true || false
	fmt.Println(b)
}
