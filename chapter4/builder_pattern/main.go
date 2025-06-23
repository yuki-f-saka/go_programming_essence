package main

import (
	"log"
	"os"
	"time"

	"builder_pattern/server"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.NewBuilder("localhost", 8080).
		Timeout(time.Minute).
		Logger(logger).
		Build()
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
