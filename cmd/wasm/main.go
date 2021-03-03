package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	var cvs, doc, ctx js.Value

	// This line will print in console
	fmt.Println("Hello, WebAssembly!")
	fmt.Println("Hello, WebAssembly!")

	width := js.Global().Get("innerWidth").Int()
	height := js.Global().Get("innerHeight").Int()
	doc = js.Global().Get("document")
	cvs = doc.Call("getElementById", "cvs")
	fmt.Println(width, height)
	cvs.Call("setAttribute", "width", 100)
	cvs.Call("setAttribute", "height", 100)
	cvs.Set("tabIndex", 0) // Not sure if this is needed
	ctx = cvs.Call("getContext", "2d")

	ctx.Call("beginPath")
	ctx.Call("rect", 20, 20, 10, 10)
	// ctx.Set("fillStyle", "red")
	ctx.Call("closePath")

	<-make(chan bool)
}
