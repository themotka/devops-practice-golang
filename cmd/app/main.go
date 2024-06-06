package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprintf(w, "Sample local server %d", addition(1, 0))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func addition(a, b int) int {
	return a + b
}
