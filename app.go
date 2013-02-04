package main

import (
	// "encoding/json"
	"fmt"
	// "html/template"
	"net/http"
	// "net/url"
	"os"
	// "strings"
	r "github.com/christopherhesse/rethinkgo"
)

type Board struct {
	Name string
}

func main() {
	r.Connect(os.Getenv("RETHINKDB_HOST"), "kanban")
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
