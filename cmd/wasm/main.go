package main

import (
	"fmt"
	"net/url"
	//"os"
	"github.com/gorilla/websocket"
	"syscall/js"
)

var conn *websocket.Conn = nil

func makeConnection() error {
	urlstr := "localhost:9090/ws-init"
	u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/ws-init"}
	fmt.Println("urlstr:", urlstr)
	fmt.Println("url formatter as string:", u.String())
	// Create websocket connection
	fmt.Println("Dialing", u.String(), "...")
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dial:", err)
		return err
	}
	return nil
}

func sendMsg(msg string) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
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
