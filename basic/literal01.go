package main

import "fmt"

func main() {
	iA := 0
	fmt.Println(iA)

	iB := 1234
	fmt.Println(iB)

	iC := -1234
	fmt.Println(iC)

	// octal
	oct := 0123
	fmt.Println(oct)

	// hex
	hex := 0xFF
	fmt.Println(hex)

	fA := 0.0
	fmt.Println(fA)

	fB := -123.45
	fmt.Println(fB)

	fC := 0.123e+3
	fmt.Println(fC)

	fD := 0.12345e5
	fmt.Println(fD)

	fE := 12345e-3
	fmt.Println(fE)

	fF := 12345e-5
	fmt.Println(fF)

	compA := 0i
	fmt.Println(compA)

	compB := -123.45i
	fmt.Println(compB)

	compC := 0.12345e+5i
	fmt.Println(compC)

	// complex()
	compD := complex(1, 2)
	fmt.Println(compD)
}
