package main

import "fmt"

func main() {
	// same : // var a, b = 100, 200
	a, b := 100, 200
	fmt.Println(a, b)

	a, b, c := 300, 400, 500
	fmt.Println(a, b, c)
}
