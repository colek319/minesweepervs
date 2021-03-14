package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

const mAXPLAYERS = 2
const mAXMATCHES = 2

// Lobby is the data for a lobby of Minesweepervs
type Lobby struct {
	Players map[string]int
	NextID  int
}

func makeNewLobby() Lobby {
	return Lobby{Players: make(map[string]int), NextID: 0}
}

var lobbies []Lobby

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	return setupRoutes(r)
}

func setupServerState() {
	// Make a match and append to matches list
	lobbies = append(lobbies, makeNewLobby())
}

func joinLobby(name string) bool {
	for i := range lobbies {
		m := &lobbies[i]
		if len(m.Players) < mAXPLAYERS {
			m.Players[name] = m.NextID
			m.NextID++
			return true
		}
	}

	// If we are full, we should not add the player
	if len(lobbies) >= mAXMATCHES {
		return false
	}

	// We now need to make a new match to hold the player
	lobbies = append(lobbies, makeNewLobby())
	m := &lobbies[len(lobbies)-1]
	m.Players[name] = m.NextID
	m.NextID++
	return true
}

// "/join/:username"
func joinLobbyHandler(c *gin.Context) {
	name := c.Param("username")
	if joinLobby(name) {
		fmt.Println(name, " joined a match")
		c.String(http.StatusOK, "Joined match as %s", name)
	} else {
		fmt.Println(name, "could not join a match")
		c.String(http.StatusOK, "Could not join match :-(")
	}
}

// "/printallmatches"
func printAllMatchesHandler(c *gin.Context) {
	fmt.Println("lobbies has form:", lobbies)
	c.JSON(http.StatusOK, lobbies)
}

// TODO(will) - figure out what our read and write buffer sizes should be
// var wsUpgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in wsHandler")
	conn, err := websocket.Accept(w, r, nil)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	fmt.Println("entering loop")
	for {
		fmt.Println("Starting read")
		_, msg, err := conn.Read(ctx)
		fmt.Println("After read")
		if err != nil {
			fmt.Println("error when reading:", err)
			conn.Close(websocket.StatusInternalError, "oops, closing...")
			fmt.Println("Connection closed")
			break
		}
		fmt.Println("Got message:", string(msg))

		m := "You sent: " + string(msg)
		err = conn.Write(r.Context(), websocket.MessageText, []byte(m))
		if err != nil {
			fmt.Println("error when writing:", err)
			break
		}
	}
	fmt.Println("Closing connection")
	conn.Close(websocket.StatusNormalClosure, "Done")

}

func main() {
	setupServerState()

	router := setupRouter()
	router.Run(":9090")
}
