package main

import "fmt"

type Fruit int
type Animal int

const (
	Apple Fruit = iota
	Orange
	Banana
)

const (
	Monkey Animal = iota
	Elephant
	Pig
)

// func main() {
// 	// fmt.Println(Apple)
// 	// fmt.Println(Orange)
// 	// fmt.Println(Banana)

// 	// fmt.Println(Monkey)
// 	// fmt.Println(Elephant)
// 	// fmt.Println(Pig)

// 	// var fruit Fruit = Apple
// 	// fmt.Println(fruit)

// 	// fruit = Elephant // コンパイルエラー
// 	// fmt.Println(fruit)

// 	// --------------------------------------------------------------------------------------------------------------

// 	// s := "hoge"
// 	// for i := 0; i < len(s); i++ {
// 	// 	fmt.Println(s[i])
// 	// }

// 	// arr := []int{1, 2, 3, 4, 5}
// 	// i := 0
// 	// for i < 3 {
// 	// 	fmt.Println(arr[i])
// 	// 	i++
// 	// }

// 	// --------------------------------------------------------------------------------------------------------------

// 	// Labeled Break
// finished:
// 	for i := 0; i < 100; i++ {
// 		for j := 0; j < 100; j++ {
// 			fmt.Printf("i: %v\n", i)
// 			fmt.Printf("j: %v\n", j)
// 			if check(i, j) {
// 				break finished
// 			}
// 		}
// 	}
// }

// func check(i, j int) bool {
// 	return (i+j)%3 == 0
// }

// --------------------------------------------------------------------------------------------------------------

// func main() {
// 	a := make([]int, 0, 100)
// 	for i := 0; i < 100; i++ {
// 		a = append(a, i)
// 	}
// 	fmt.Println(a)
// }

// --------------------------------------------------------------------------------------------------------------

// スライスから要素を削除する

func main() {
	a := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}

	fmt.Println(a)

	// 新しく同じ型のスライスを用意して詰め直す
	// a2 := make([]int, 0, len(a)/2)
	// for i := 0; i < len(a); i++ {
	// 	if i%2 == 0 {
	// 		a2 = append(a2, a[i])
	// 	}
	// }
	// a = a2
	// fmt.Println(a)

	// appendを使う
	// n := 50
	// a = append(a[:n], a[n+1:]...) // 50をaから削除する
	// fmt.Println(a)

	// 部分参照とコピー
	m := 50
	a = a[:m+copy(a[m:], a[m+1:])]
	fmt.Println(a)
}
