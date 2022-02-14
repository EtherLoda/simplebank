package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "pong!"})
}
