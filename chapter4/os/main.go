package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// _, err = io.Copy(os.Stdout, f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)

	fmt.Println(content)

}
