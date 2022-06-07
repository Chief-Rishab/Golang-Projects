package main

import (
    "fmt"
    "log"
    "net/http"
)
//The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an http.Request as its arguments.
func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //Fprintf is printing on the stream, r.URL.Path[1:] will omit / from the Url path
}

//r.URL.Path is the path component of the request URL.

//The main function begins with a call to http.HandleFunc, 
//which tells the http package to handle all requests to the web root ("/") with handler.

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":3000", nil)) //It then calls http.ListenAndServe, specifying that it should listen on port 3000 on any interface (":3000").
	//This function will block until the program is terminated.
}


// If you run this program and access the URL:

// http://localhost:3000/monkeys
// the program would present a page containing:

// Hi there, I love monkeys!