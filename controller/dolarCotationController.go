package controller

import (
	"fmt"
	"net/http"

	"github.com/Guilherme415/Client-Server-API/business"
	"github.com/gin-gonic/gin"
)

type IDolarCotationController interface {
	SaveDolarCotation(c *gin.Context)
}

type DolarCotationController struct {
	dolarCotationBusiness business.IDolarCotationBusiness
}

func NewDolarControllerController(dolarCotationBusiness business.IDolarCotationBusiness) IDolarCotationController {
	return &DolarCotationController{dolarCotationBusiness}
}

func (d *DolarCotationController) SaveDolarCotation(c *gin.Context) {
	dolarCotation, err := d.dolarCotationBusiness.DolarCotationMonitoring()
	if err != nil {
		errorMessage := fmt.Sprintf("An internval server error has ocurred, details: %s", err.Error())

		c.JSON(http.StatusBadRequest, errorMessage)
		return

	}

	c.JSON(http.StatusOK, dolarCotation)
}
