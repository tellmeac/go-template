package web

import "github.com/gin-gonic/gin"

// Route represents route that should be handled by server.
type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
