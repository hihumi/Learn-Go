package main

import "fmt"

func main() {
	fmt.Println("1 == 1 is", 1 == 1)
	fmt.Println("1 == 2 is", 1 == 2)

	fmt.Println("1 != 1 is", 1 != 1)
	fmt.Println("1 != 2 is ", 1 != 2)

	fmt.Println("1 > 2 is", 1 > 2)
	fmt.Println("1 < 2 is", 1 < 2)

	fmt.Println("100 >= 99 is", 100 >= 99)
	fmt.Println("100 <= 99 is", 100 <= 99)
	fmt.Println("100 >= 100 is", 100 >= 100)
	fmt.Println("100 <= 100 is", 100 <= 100)

	fmt.Println("\"aaa\" == \"aaa\" is", "aaa" == "aaa")
	fmt.Println("\"aaa\" == \"aab\" is", "aaa" == "aab")
}
