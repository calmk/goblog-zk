package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello\n")
}

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hello zk</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/zk", hello_world)
	http.ListenAndServe(":7000", nil)
}
