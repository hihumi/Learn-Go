package main

import "fmt"

func main() {
	var en string = "japanese"
	fmt.Println(en)
	fmt.Println(en, "len :", len(en))

	var ja string = "日本・日本語・日本人"
	fmt.Println(ja)
	fmt.Println(ja, "len :", len(ja))
}
