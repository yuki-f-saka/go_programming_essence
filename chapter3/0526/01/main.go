package main

import (
	"fmt"
	"sort"
)

func main() {
	// s := "hello"
	// fmt.Printf("%c", s[0])

	// // 内容を書き換えるにはバイト列に変換する
	// b := []byte(s)
	// b[0] = 'H'
	// s = string(b)
	// fmt.Printf("%c", s[0])

	// ----------------------------------------------------------------------------------------------------

	// s := "こんにちわ世界"
	// rs := []rune(s)
	// rs[4] = 'は'
	// s = string(rs)
	// fmt.Println(s)

	// ----------------------------------------------------------------------------------------------------

	// var b [4]byte
	// b2 := b[:]
	// // 固定長の配列は上記のようにスライスに変換できる

	// ----------------------------------------------------------------------------------------------------

	// var content = `複数行の
	// 	文章から
	// 	なっている
	// `
	// fmt.Println(content)

	// ----------------------------------------------------------------------------------------------------

	// map

	m := make(map[string]int, 10)
	m["John"] = 3
	m["Emi"] = 4
	m["Johnny"] = 6

	fmt.Println(m)

	delete(m, "Emi")

	fmt.Println(m)

	// # mapは順序を保証しない
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, m[k])
	}

}
