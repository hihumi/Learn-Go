package main

import "fmt"

func main() {
	fmt.Println("1 + 2 =", 1+2)
	fmt.Println("2 - 1 =", 2-1)
	fmt.Println("2 * 2 =", 2*2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("-10 / 3 =", -10/3)
	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("3 & 6 =", 3&6)
	fmt.Println("3 | 6 =", 3|6)
	fmt.Println("3 &^ 6 =", 3&^6)
	fmt.Println("2 << 1 =", 2<<1)
	fmt.Println("2 >> 1 =", 2>>1)
	fmt.Println("\"abc\" + \"def\" =", "abc"+"def")
}
