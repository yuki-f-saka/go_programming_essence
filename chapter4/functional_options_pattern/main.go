package main

import (
	"log"
	"os"
	"time"
)

// Goには関数のデフォルト引数がない
// Goにも可変引数はあるが、型が固定になってしまう。

/*
	func doSomething(args ...string) {
	}
*/

// anyを使って型のことなる可変個引数を使うこともできるが、関数が扱える型をコンパイル時に固定できない
// functional options patternを使用すると、開発者が限定した型のみを可変個で渡せる関数を提供できる

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func New(host string, port int, options ...Option) *Server {
	srv := &Server{
		host: host,
		port: port,
	}
	for _, opt := range options {
		opt(srv)
	}
	return srv
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

type Option func(*Server)

func (s *Server) Start() error {
	// do something
	return nil
}

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	srv := New("localhost", 8080,
		WithTimeout(time.Minute),
		WithLogger(logger),
	)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}

	// ...Optionに関しては以下のように必須ではないので、以前のコードを壊さない。
	srv2 := New("localhost", 3000)
	if err := srv2.Start(); err != nil {
		log.Fatal(err)
	}
}
