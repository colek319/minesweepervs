package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	//"os"
	"nhooyr.io/websocket"
	"syscall/js"
)

const TimeoutDuration = time.Second * 60

var conn *websocket.Conn = nil

func makeConnection() error {
	u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/ws-init"}

	// Create websocket connection
	// var err error
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutDuration)
	defer cancel()
	// ctx := context.TODO()
	fmt.Println("Dialing", u.String(), "...")
	conn, _, _ = websocket.Dial(ctx, u.String(), nil)
	fmt.Println("After dial")
	// if err != nil {
	// 	fmt.Println("dial:", err)
	// 	return err
	// }
	fmt.Println("returning nil from makeConnection")
	return nil
}

func makeConnectionWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 0 {
			return "Invalid number of arguments passed"
		}
		var err error
		go func() {
			err = makeConnection()
		}()
		if err != nil {
			fmt.Println("Got error making connection:", err)
			return "Error making connection"
		}
		return "Connection established"
	})
	return f
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

func readMsg() (string, error) {
	fmt.Println("Reading message")
	_, msg, err := conn.Read(context.TODO())
	fmt.Println(string(msg))
	if err != nil {
		return "", errors.New("error reading from server")
	}
	return string(msg), nil
}

func readMsgWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 0 {
			return "invalid number of arguments"
		}
		var msg string
		var err error
		go func() {
			msg, err = readMsg()
		}()
		if err != nil {
			fmt.Println("error while reading message: ", err)
			return fmt.Sprintf("error while reading message: %s", err)
		}
		return string(msg)
	})
	return f
}

func main() {
	// Expose functions to JS
	js.Global().Set("makeConnection", makeConnectionWrapper())
	js.Global().Set("sendMsg", sendMsgWrapper())
	js.Global().Set("readMsg", readMsgWrapper())

	// Keep the gocode running
	<-make(chan bool)
}
