package main

import "fmt"

func main() {
	// openFile()
	doSomething1()
	doSomething2()
}

// BAD example
// deferはループ毎には発行されない
// func openFile() error {
// 	for i := 0; i < 100; i++ {
// 		f, err := os.Open("data.txt")
// 		if err != nil {
// 			return err
// 		}
// 		defer f.Close()
// 	}
// 	return nil
// }

func doSomething1() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()

	n = 2
}

func doSomething2() {
	var n = 1
	defer fmt.Println(n)

	n = 2
}
