package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	//"os"
	"nhooyr.io/websocket"
	"syscall/js"
)

var conn *websocket.Conn = nil

func makeConnection() error {
	go func() {
		u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/ws-init"}

		// Create websocket connection
		// var err error
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		// ctx := context.TODO()
		fmt.Println("Dialing", u.String(), "...")
		conn, _, _ = websocket.Dial(ctx, u.String(), nil)
		conn.Write(context.TODO(), websocket.MessageText, []byte("Hello World"))
		fmt.Println("After dial")
		// if err != nil {
		// 	fmt.Println("dial:", err)
		// 	return err
		// }
		fmt.Println("returning nil from makeConnection")
	}()
	return nil
}

func sendMsg(msg string) error {
	// err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
	err := conn.Write(context.TODO(), websocket.MessageText, []byte(msg))
	return err
}

func sendMsgWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		err := sendMsg(args[0].String())
		if err != nil {
			fmt.Println("Got error sending message:", err)
			return fmt.Sprintf("Got error while writing: %s", err)
		}
		return "Message sent"
	})
	return f
}

func makeConnectionWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 0 {
			return "Invalid number of arguments passed"
		}
		err := makeConnection()
		if err != nil {
			fmt.Println("Got error making connection:", err)
			return "Error making connection"
		}
		return "Connection established"
	})
	return f
}

func main() {
	// Expose functions to JS
	js.Global().Set("makeConnection", makeConnectionWrapper())
	js.Global().Set("sendMsg", sendMsgWrapper())

	// Keep the gocode running
	<-make(chan bool)
}
