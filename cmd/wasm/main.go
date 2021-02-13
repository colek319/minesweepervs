package main

import "fmt"

func main() {
	// This line will print in console
	fmt.Println("Hello, WebAssembly!")

	<-make(chan bool)
}
