package main

import "fmt"

func main() {
	var i int = 1234567890
	fmt.Println(i)

	var ui32 uint32 = uint32(i)
	fmt.Println(ui32)

	var f32 float32 = float32(ui32)
	fmt.Println(f32)

	var s string = string(i)
	fmt.Println(s)

	var b []byte = []byte("xyz")
	fmt.Println(b)

	var bJaFirst []byte = []byte("あ")
	fmt.Println(bJaFirst)

	var bJaSecond []byte = []byte("あいう")
	fmt.Println(bJaSecond)
}
