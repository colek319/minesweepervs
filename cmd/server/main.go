package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func setupServer() {
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

func main() {
	setupServer()

	router := setupRouter()
	router.Run(":9090")
}
