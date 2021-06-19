package util

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func DoHttpError(c *gin.Context, code int, err error) {
	DoError(err)
	errParts := strings.Split(err.Error(), "|")
	c.String(code, errParts[len(errParts)-1])
}

func DoError(err error) {
	log.Println(err)
}
