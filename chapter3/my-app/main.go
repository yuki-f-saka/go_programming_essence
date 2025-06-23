package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	// 定数で宣言すると使われ方によって型が自動で決定されるためコンパイルエラーにならない。
	// const x = 1
	// y := 1
	// f := 1.2 + (x+2)*float64(y)
	// fmt.Println(f)

	// a := 1
	// b := 1
	// g := 1.2 + (a+2)*float64(b)  これだとコンパイルエラー

	// ------------------------------------------------------------------------------------------------

	// iotaについて

	// const (
	// 	Apple = iota
	// 	Orange
	// 	Banana
	// )

	// fmt.Println(Apple)  // 0
	// fmt.Println(Orange) // 1
	// fmt.Println(Banana) // 2

	// ---------------------------------------------

	// const (
	// 	Apple  = iota + iota // 0 + 0 = 0
	// 	Orange               // 1 + 1 = 2
	// 	Banana = iota + 3    // 2 + 3 = 5
	// )

	// fmt.Println(Apple)  // 0
	// fmt.Println(Orange) // 2
	// fmt.Println(Banana) // 5

	// ------------------------------------------------------------------------------------------------
}
