package main

import "fmt"

func main() {
	var n1 int
	var quit string = "q"
	var yes string = "y"

	fmt.Println("This program is the judgement of Even or Odd.\nPlease Enter an integer.\nRepeat key is y. Quit key is q.")
	for {
		fmt.Printf(">> ")
		fmt.Scanf("%d", &n1)

		if n1%2 == 0 {
			fmt.Printf("an even\n")
		} else {
			fmt.Printf("an odd\n")
		}

		fmt.Printf("Plaese Enter y or q key >> ")
		fmt.Scanln(&yes)
		if yes == "y" {
			continue
		}
		fmt.Scanln(&quit)
		if quit == "q" {
			break
		}
	}
	fmt.Printf("the end\n")
}
