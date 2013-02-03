package main

import (
	"fmt"
	"net/http"
	"os"
	r "github.com/christopherhesse/rethinkgo"
)

func main() {
	// r.connect(os.Getenv("RETHINKDB_HOST"))
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
	  panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Hello, World")
}
