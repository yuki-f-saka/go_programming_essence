package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// path := "/Users/funesaka/Documents/confirm.pdf"
	// println(filepath.Base(path))

	// path1 := "/Users/funesaka/Documents/../../confirm.pdf"
	// println(filepath.Clean(path1))

	// println(filepath.Dir(path))

	// println(filepath.Ext(path))

	// println(filepath.IsAbs(path))

	// path2 := "/Users/funesaka/"
	// path3 := "Documents/confirm.pdf"
	// println(filepath.Join(path2, path3))

	// println("path/filepath")
	// absolute, err := filepath.Rel("/Users/funesaka", "/Users/funesaka/Documents/confirm.pdf")
	// if err == nil {
	// 	println(absolute)
	// }

	// println(filepath.VolumeName(path))

	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
}
