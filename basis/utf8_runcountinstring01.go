package main

import "fmt"
import "unicode/utf8"

func main() {
	var ja string = "日本"
	fmt.Println(ja, "len :", utf8.RuneCountInString(ja))
}
