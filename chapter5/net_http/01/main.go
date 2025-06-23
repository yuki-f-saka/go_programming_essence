package main

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			// GETの処理
// 		default:
// 			//
// 		}
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		f, err := os.Open("/path/th/content.txt")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer f.Close()
// 		io.Copy(w, f)
// 	})
// }

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world")
// }

// func main() {
// 	http.HandleFunc("/", myHandler)
// 	http.ListenAndServe(":8080", nil)
// }
