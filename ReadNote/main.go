package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chief-Rishab/mymodule/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
