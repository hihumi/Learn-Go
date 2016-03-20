package main

import "fmt"

func main() {
	var n int

	fmt.Printf("This program can multiply your number by 100.\nPlease Enter a numeric number.\n>> ")
	fmt.Scanf("%d", &n)

	fmt.Printf("%d * 100 = %d\n", n, n*100)
}
