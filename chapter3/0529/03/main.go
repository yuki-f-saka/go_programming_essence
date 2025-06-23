package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisor must not be zero")
	}
	return a / b, nil
}

func exitIf(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", message, err)
		os.Exit(1)
	}
}

func main() {
	// 引数の数
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Number of argument must be 2")
		os.Exit(1)
	}
	// 引数a
	a, err := strconv.ParseFloat(os.Args[1], 64)
	exitIf(err, "First argument must be float value")
	// 引数b
	b, err := strconv.ParseFloat(os.Args[2], 64)
	exitIf(err, "Second argument must be float value")
	// やっと計算
	result, err := divide(a, b)
	exitIf(err, "Invalid argument")
	fmt.Println(result)

	/*
		Goはこのようにエラーを明示的にハンドリングすることの重要さを言語使用として伝えています。
		エラーが起きうる関数ではerrorを返すべき。
	*/
}
