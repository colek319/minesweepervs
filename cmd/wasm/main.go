package main

import (
	"fmt"
	"syscall/js"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/htmlcanvas"
)

func onClickHandler() {

}

func onClickHandlerWrapper() {

}

func main() {
	// This line will print in console
	fmt.Println("Hello, WebAssembly!")

	cvs := js.Global().Get("document").Call("getElementById", "cvs")
	c := htmlcanvas.New(cvs, 100, 100, 5.0)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(canvas.Black)
	ctx.DrawPath(10, 10, canvas.Rectangle(10, 10))

	<-make(chan bool)
}
