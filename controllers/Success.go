package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
