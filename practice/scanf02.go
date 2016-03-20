package main

import "fmt"

func main() {
	var n int

	fmt.Printf("This program can multiply your enter integer by 100.\nPlease Enter a integer.\n>> ")
	fmt.Scanf("%d", &n)

	fmt.Printf("%d * 100 = %d\n", n, n*100)
}
