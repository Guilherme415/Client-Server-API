package config

import (
	"github.com/Guilherme415/Client-Server-API/config/dependency"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	e.POST("", dependency.DolarCotationController.SaveDolarCotation)
}
