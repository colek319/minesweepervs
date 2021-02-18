package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * Put routes here.  If you want to load an HTML file you need to add it
 * explicitly by calling LoadHTMLFiles.
 */
func setupRoutes(r *gin.Engine) *gin.Engine {
	r.LoadHTMLFiles("assets/index.html")

	r.GET("/join/:username", joinLobbyHandler)
	r.GET("/printalllobbies", printAllMatchesHandler)

	// Serve index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	//r.Static("/", "index.html")

	return r
}
