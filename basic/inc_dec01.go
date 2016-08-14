package main

import "fmt"

func main() {
	var inc int = 0
	var dec int = 0

	inc++
	dec--

	fmt.Println("inc++ is", inc)
	fmt.Println("dec-- is", dec)

	inc++
	dec--

	fmt.Println("inc++ is", inc)
	fmt.Println("dec-- is", dec)
}
