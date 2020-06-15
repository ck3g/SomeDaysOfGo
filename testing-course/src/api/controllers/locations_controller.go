package controllers

import (
	"net/http"

	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/services"
	"github.com/gin-gonic/gin"
)

// GetCountry fetches a country by :country_id query param
func GetCountry(c *gin.Context) {
	country, err := services.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, country)
}
