package handler

import (
	"github.com/gin-gonic/gin"
)

func ShowTop(c *gin.Context) {
	c.HTML(200, "top.html", nil)
}
