package main

// Learnings in this project

/*
Handling the json documents bw server and client

Using Gorilla Mux, a special type handler for functions for the web servers.
Using Gorilla Mux we can ensure that the Method is as per the requirements of the User.
We can specify GET, PUT, DEL.. methods before hand

Generated this movies API and tested the end points with Postman, also learnt how to get parameters from the request body

Also used the http.ResponseWriter function to Set the header of the response (to json elements) which will be received by the client

Decoding the JSON document(string) from client into struct variables

*/

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

/*
Suppose that the server will send some JSON.
By setting the Content-Type header, the server can inform the client that JSON data is being sent.
Through the ResponseWriter, a handler function can add a header as follows:
w.Header().Set("Content-Type", "application/json; charset=utf-8")
*/

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // gets the parameter passed in the request fromt the client
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Syntax 2 for append:  append(slice, another Slice ...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)

	//loop over movies slice
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	//delete the movie with the id you've sent
	//add a new movie- the move that we send in the body of postman
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "458454", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

/* ----------- Reading for working with JSON and structs(JSON is encoded and decoded in []byte slices for transfer bw server and clients) in Golang -----------------------
Working with JSON
If you need to work with JSON objects, you will be pleased to hear that Go's standard library includes everything you need to parse and encode JSON through the encoding/json package.
Default types
When encoding or decoding a JSON object in Go, the following types are used:
bool for JSON booleans,
float64 for JSON numbers,
string for JSON strings,
nil for JSON null,
map[string]interface{} for JSON objects, and
[]interface{} for JSON arrays.
Encoding
To encode a data structure as JSON, the json.Marshal function is used. Here's an example:
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    FirstName string
    LastName  string
    Age       int
    email     string
}

func main() {
    p := Person{
        FirstName: "Abraham",
        LastName:  "Freeman",
        Age:       100,
        email:     "abraham.freeman@hey.com",
    }

    json, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(json))
}

In the above program, we have a Person struct with four different fields. In the main function, an instance of Person is created with all the fields initialized. The json.Marshal method is then used to convert the p structure to JSON. This method returns a slice of bytes or an error, which we have to handle before accessing the JSON data.
To convert a slice of bytes to a string in Go, we need to perform type conversion, as demonstrated above. Running this program will produce the following output:
{"FirstName":"Abraham","LastName":"Freeman","Age":100}

As you can see, we get a valid JSON object that can be used in any way we want. Note that the email field is left out of the result. This is because it is not exported from the Person object by virtue of starting with a lowercase letter.
By default, Go uses the same property names in the struct as field names in the resulting JSON object. However, this can be changed through the use of struct field tags.
type Person struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age"`
    email     string `json:"email"`
}

The struct field tags above specify that the JSON encoder should map the FirstName property in the struct to a first_name field in the JSON object and so on. This change in the previous example produces the following output:
{"first_name":"Abraham","last_name":"Freeman","age":100}


Decoding
The json.Unmarshal function is used for decoding a JSON object into a Go struct. It has the following signature:
func Unmarshal(data []byte, v interface{}) error

It accepts a byte slice of JSON data and a place to store the decoded data. If the decoding is successful, the error returned will be nil.
Assuming we have the following JSON object,
json := "{"first_name":"John","last_name":"Smith","age":35, "place_of_birth": "London", gender:"male"}"

We can decode it to an instance of the Person struct, as shown below:


func main() {
    b := `{"first_name":"John","last_name":"Smith","age":35, "place_of_birth": "London", "gender":"male", "email": "john.smith@hmail.com"}`

    var p Person
    err := json.Unmarshal([]byte(b), &p)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%+v\n", p)
}


And you get the following output:
{FirstName:John LastName:Smith Age:35 email:}

Unmarshal only decodes fields that are found in the destination type. In this case, place_of_birth and gender are ignored since they do not map to any struct field in Person. This behavior can be leveraged to pick only a few specific fields out of a large JSON object. As before, unexported fields in the destination struct are unaffected even if they have a corresponding field in the JSON object. That's why email remains an empty string in the output even though it is present in the JSON object.

*/
