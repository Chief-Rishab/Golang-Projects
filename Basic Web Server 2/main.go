package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	add := r.FormValue("address")
	fmt.Fprintf(w, "Entered values are Name: %s and Address: %s", name, add)
}

func main() {
	fileServer := http.FileServer(http.Dir("./")) //FileServer returns a handler that serves HTTP requests
	//with the contents of the file system rooted at root.
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at Port 2000\n")
	if err := http.ListenAndServe(":2000", nil); err != nil {
		log.Fatal(err)
	}
}
