package main

import "fmt"

func main() {
	fmt.Println("true && true is", true && true)
	fmt.Println("true && false is", true && false)
	fmt.Println("false && true is", false && true)
	fmt.Println("false && false is", false && false)

	fmt.Println("true || true is", true || true)
	fmt.Println("true || false is", true || false)
	fmt.Println("false || true is", false || true)
	fmt.Println("false || false is", false || false)

	fmt.Println("!true is", !true)
	fmt.Println("!!true is", !!true)
	fmt.Println("!!!true is", !!!true)

	fmt.Println("!false is", !false)
	fmt.Println("!!false is", !!false)
	fmt.Println("!!!false is", !!!false)
}
