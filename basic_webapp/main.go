package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v", err)
	}

	fmt.Fprintf(w, "POST request succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "Not found, look elsewhere", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "not found, look elsewhere", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server on http://localhost:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {

		log.Fatal(err)
	}
}
