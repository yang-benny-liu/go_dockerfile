package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// API routes

	// Serve files from static folder
	//	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/", http.FileServer(http.Dir("~/")))

	// Serve api /hi
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":8080"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}
