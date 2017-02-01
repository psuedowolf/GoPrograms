package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Smruthi")
}
func main() {
	http.HandleFunc("/", sayHelloName) // say hi
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen and Server error:", err)
	}
	fmt.Println("vim-go")
}
