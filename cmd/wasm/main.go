package main

import (
	"context"
	"fmt"
	"net/url"
	"time"
	//"os"
	"syscall/js"

	"nhooyr.io/websocket"
)

var conn *websocket.Conn = nil

func makeConnection() error {
	u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/ws-init"}

	// Create websocket connection
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// ctx := context.TODO()
	fmt.Println("Dialing", u.String(), "...")
	conn, _, err = websocket.Dial(ctx, u.String(), nil)
	fmt.Println("After dial")
	if err != nil {
		fmt.Println("dial:", err)
		return err
	}
	fmt.Println("returning nil from makeConnection")
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
	//ws := js.Global().Get("WebSocket").New("ws://localhost:9090/ws-init")
	// Expose functions to JS
	js.Global().Set("makeConnection", makeConnectionWrapper())
	js.Global().Set("sendMsg", sendMsgWrapper())

	// Keep the gocode running
	<-make(chan bool)
}
