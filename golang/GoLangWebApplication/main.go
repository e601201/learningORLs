package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[2:])
		}),
	)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}