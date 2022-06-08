package main

import (
	"fmt"
	"log"
	"net/http"
)

//The first argument to the helloHandler function is a value of the type http.ResponseWriter. 
//This is the mechanism used for sending responses to any connected HTTP clients.
//It's also how response headers are set. The second argument is a pointer to an http.Request. 
//It's how data is retrieved from the web request. For example, the details from a form submission can be accessed through the request pointer.

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
/*
Go's ServeMux does not have any special way to specify the allowed methods for a route,
so you have to check the request method yourself in the HTTP handler.

Let's make sure that the indexHandler function returns an error if a non-GET request is made to the root route.

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    if r.Method == "GET" {
        w.Write([]byte("<h1>Welcome to my web server!</h1>"))
    } else {
        http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
    }
}
*/

/*
The http.ResponseWriter interface has a Write method which accepts a byte slice and writes the data to the connection as part of an HTTP response. 
Converting a string to a byte slice is as easy as using []byte(str), and that's how we're able to respond to HTTP requests.

example usage::: w.Write([]byte("<h1>Welcome to my web server!</h1>"))
*/

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
	/*
	Serving static assets
	To serve static files from a web server, the http.FileServer() method can be utilized.

	func main() {
    	staticHandler := http.FileServer(http.Dir("./assets"))
    	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))
	}
	
	The FileServer() function returns an http.Handler that responds to all HTTP requests with the contents of the provided filesystem.
	In the above example, the filesystem is given as the assets directory relative to the application.
	
	For other details go through- 
	http.Redirect, http.NotFound, 
	https://www.honeybadger.io/blog/go-web-services/
	
	*/
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
/*
ListenAndServe starts an HTTP server with a given address and handler. 
The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
*/
