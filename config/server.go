package config

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	r := gin.Default()

	return r
}
