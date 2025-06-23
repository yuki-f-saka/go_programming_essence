package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var v interface{}
	// v = 1
	// n := v.(int)
	// fmt.Println(n)

	// v = "こんにちは世界"
	// s := v.(string)
	// fmt.Println(s)

	// --------------------------------------------------------------------------------------------------------------

	// var v interface{}

	// v = "こんにちは世界"

	// if s, ok := v.(string); !ok {
	// 	fmt.Printf("vはstringではない")
	// } else {
	// 	fmt.Printf("vは文字列 %q です", s)
	// }

	// --------------------------------------------------------------------------------------------------------------

	var v V = 123

	PrintDetai(v)
}

type V int

// どんな値でも引数として受け取れ、処理できる関数を実装するには型スイッチをつかう
func PrintDetai(v any) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/int32/int64 型:", rt)
	case reflect.String:
		fmt.Println("string型: ", rt)
	default:
		fmt.Println("知らない型")
	}
}
