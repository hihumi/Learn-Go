package main

import "fmt"

func main() {
	var n1, n2 int
	var product int

	fmt.Printf("This program can calculate the product of two integers.\nPlease Enter two integers.\n")

	fmt.Printf("First integer >> ")
	fmt.Scanf("%d", &n1)

	fmt.Printf("Second integer >> ")
	fmt.Scanf("%d", &n2)

	product = n1 * n2
	fmt.Printf("The product of %d * %d = %d\n", n1, n2, product)
}
