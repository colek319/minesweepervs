package main

import (
	"context"
	"fmt"
	"net/url"
	"nhooyr.io/websocket"
	"time"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/ws-init"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	fmt.Println("dialing...")
	c, _, err := websocket.Dial(ctx, u.String(), nil)
	defer c.Close(websocket.StatusNormalClosure, "Closing")
	if err != nil {
		fmt.Println("Got error:", err)
	}
	c.Write(ctx, websocket.MessageText, []byte("Hi this is a message"))
}
