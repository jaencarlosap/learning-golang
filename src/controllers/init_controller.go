package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to api")
}
