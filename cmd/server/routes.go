package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

/*
 * Put routes here.  If you want to load an HTML file you need to add it
 * explicitly by calling LoadHTMLFiles.
 */
func setupRoutes(r *gin.Engine) *gin.Engine {
	r.LoadHTMLFiles("assets/index.html")

	r.GET("/join/:username", joinLobbyHandler)
	r.GET("/printalllobbies", printAllMatchesHandler)
	r.GET("/ws-init", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})
	r.Use(static.Serve("/assets", static.LocalFile("./assets", false)))

	// Serve index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}
