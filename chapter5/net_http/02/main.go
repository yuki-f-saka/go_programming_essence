package main

import (
	"database/sql"
	"net/http"
)

type MyContext struct {
	db *sql.DB
}

func (m *MyContext) handle(w http.ResponseWriter, r *http.Request) {
	// m.dbを使った処理
}

// func main() {
// 	myctx := NewMyContext()
// 	http.HandleFunc("/", myctx.handle)
// 	//
// }

// ここのmyctx.handleのようにメソッドは引数を別にその場で以下のように書かなくてもいい
// myctx.handle(w, r)
