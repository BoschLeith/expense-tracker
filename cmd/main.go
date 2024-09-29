package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8000", nil)
}
