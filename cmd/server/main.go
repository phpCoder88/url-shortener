package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Everything is working")
	})

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		os.Exit(1)
	}
}
