package webservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
