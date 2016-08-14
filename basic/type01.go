package main

import "fmt"

func main() {
	type typeInteger int
	//fmt.Println(typeInteger) // error

	var i typeInteger = 1234567890
	fmt.Println(i)

	i += 1
	fmt.Println(i)

	i -= 1
	fmt.Println(i)

	i = i + 1
	fmt.Println(i)

	i = i - 1
	fmt.Println(i)

	i *= 2
	fmt.Println(i)

	i /= 2
	fmt.Println(i)

	i = i * 2
	fmt.Println(i)

	i = i / 2
	fmt.Println(i)

}
