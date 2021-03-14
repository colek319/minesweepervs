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
	fmt.Println(u.String())
	var err error
	_ = err

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, _, err = websocket.Dial(ctx, u.String(), nil)
	defer conn.Close(websocket.StatusInternalError, "")

	conn.SetReadLimit(65536)
	for i := 0; i < 10; i++ {
		err = conn.Write(ctx, websocket.MessageText, []byte("1"))
	}

	err = conn.Close(websocket.StatusNormalClosure, "")
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
