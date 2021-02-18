package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server on port 9090")
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("./assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		os.Exit(1)
	}
}
