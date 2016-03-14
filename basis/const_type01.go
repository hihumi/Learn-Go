package main

import "fmt"

const C1 = 123
const C2 int = 456

func main() {
	var a int = C1 + C2
	fmt.Println("123 + 456 =", a)

	var b int32 = 789
	var c int64 = 1
	fmt.Println("789 + 1 =", b+int32(c))
	fmt.Println("789 + 1 =", int64(b)+c)
}
