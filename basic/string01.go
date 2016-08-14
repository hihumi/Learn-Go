package main

import "fmt"

func main() {
	var str01 string = "あいうえお"
	fmt.Println(str01)

	str01 = str01 + "かきくけこ"
	fmt.Println(str01)

	str01 += "さしすせそ"
	fmt.Println(str01)

	/* var str02 string
	   # command-line-arguments
	   ./string01.go:15: str02 declared and not used
	*/
}
