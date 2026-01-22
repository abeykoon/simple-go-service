package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received, sleeping for 10 seconds...")
		time.Sleep(10 * time.Second)
		fmt.Fprintln(w, "Hello! I slept for 10 seconds.")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test endpoint")
	})

	fmt.Println("Server starting on :9095")
	http.ListenAndServe(":9095", nil)
}
