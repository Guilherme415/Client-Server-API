package config

import (
	"fmt"

	"github.com/Guilherme415/Client-Server-API/db"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	err := db.ConnectDB()
	if err != nil {
		panic(fmt.Errorf("can't connet DB, err: %s", err.Error()))
	}

	return r
}
