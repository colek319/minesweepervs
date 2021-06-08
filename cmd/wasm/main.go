package main

import (
	"context"
	"errors"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"

	//"os"
	"syscall/js"

	"nhooyr.io/websocket"
)

const TimeoutDuration = time.Second * 10

var (
	conn       *websocket.Conn    = nil
	connCtx    context.Context    = nil
	connCancel context.CancelFunc = nil
	connAlive  chan bool          = make(chan bool)
)

func makeConnection(ctx context.Context) error {
	u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/ws-init"}
	alive := make(chan bool)
	connCtx = context.WithValue(context.Background(), "alive", alive)
	connCtx, connCancel = context.WithCancel(connCtx)
	// Create websocket connection
	var err error
	log.Infof("dialing: %s", u.String())
	conn, _, err = websocket.Dial(connCtx, u.String(), nil)
	if err != nil {
		log.Errorf("failed to dial server at %s: %v", u.String(), err)
		return err
	}
websocketLoop:
	for {
		select {
		case <-time.After(TimeoutDuration):
			connCtx = nil
			connCancel()
			log.Info("connection timed out")
			break websocketLoop
		case <-alive:
			log.Debug("connection reset")
		case <-connCtx.Done():
			connCtx = nil
			log.Info("connection cancalled")
			break
		}
	}
	return nil
}

func makeConnectionWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 0 {
			log.Fatal("failed to make connection: invalid number of arguments passed")
			return ""
		}
		var err error
		go func() {
			// if make connection
			err = makeConnection(connCtx)
		}()
		if err != nil {
			log.Errorf("error making connection: %v", err)
			return ""
		}
		return ""
	})
	return f
}

func sendMsg(ctx context.Context, msg string) error {
	if ctx == nil {
		return errors.New("nil connection")
	}
	writeCtx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	log.Debugf("sending message: %s", msg)
	err := conn.Write(writeCtx, websocket.MessageText, []byte(msg))
	ctx.Value("alive").(chan bool) <- true
	return err
}

func sendMsgWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			log.Fatal("failed to send message: invalid number of arguments passed")
			return ""
		}
		var err error
		c := make(chan bool)
		go func() {
			err = sendMsg(connCtx, args[0].String())
			c <- err != nil
		}()
		if <-c {
			log.Errorf("error while writing server: %v", err)
			return ""
		}
		return ""
	})
	return f
}

func readMsg(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errors.New("nil connection")
	}
	readCtx, cancel := context.WithTimeout(connCtx, time.Millisecond*500)
	defer cancel()

	_, msg, err := conn.Read(readCtx)
	if err != nil {
		return "", err
	}
	log.Debugf("message received: %s", msg)
	ctx.Value("alive").(chan bool) <- true
	return string(msg), nil
}

func readMsgWrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 0 {
			return "invalid number of arguments"
		}
		var msg string
		var err error
		c := make(chan bool)
		go func() {
			msg, err = readMsg(connCtx)
			c <- err != nil
		}()
		if <-c {
			log.Errorf("error while reading: %v", err)
			return msg
		}
		return msg
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
