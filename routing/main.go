package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"]   // the page number slug

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}

/*
	Methods: Restrict the request handler to specific HTTP methods.
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")


	Hostnames & Subdomains: Restrict the request handler to specific hostnames or subdomains.
	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	Schemes: Restrict the request handler to http/https.
	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	Path Prefixes & Subrouters: Restrict the request handler to specific path prefixes.
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)
*/
